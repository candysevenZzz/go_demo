package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var (
	path = "http://www.xbiquge.la/0/15/12961.html" //基础路径
)

//获取命令行输入
func getStdIn() (start, end int, ok bool){
	ok = true

	fmt.Println("请输入爬虫的起始页：")
	_, _ = fmt.Scan(&start)
	fmt.Println("请输入爬虫的终止页：")
	_, _ = fmt.Scan(&end)
	if end < start || start < 0 {
		fmt.Println("爬虫的页数范围错误：起始页必须大于等于0且小于等于中止页")
		ok = false
		return
	}

	var kw string
	fmt.Println("请输入文章名称：")
	_, _ = fmt.Scan(&kw)
	return
}

//处理任务
func DoWork(start, end int)  {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go crawling(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

//爬取数据
func crawling(i int, page chan<- int)  {
	path = fmt.Sprintf(path, i)
	resp, err := http.Get(path)
	if !errorCheck(err, "http.Get") {
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4*1024)
	result := ""
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			errorCheck(err, "resp.Body.Read")
			break
		}
		result += string(buf[:n])
	}

	//内容过滤
	//res := filter(result, "<h1>(?s:(.*?))</h1>")
	Store(result, strconv.Itoa(i), ".html")
	page <- i
}

//错误检测
func errorCheck(err error, msg string) (ok bool) {
	ok = true
	if err != nil {
		ok = false
		fmt.Printf("%s error: %v\n", msg,  err)
	}
	return
}

//内容过滤
func filter(content, regex string) (res string) {
	reg := regexp.MustCompile(regex)
	str := reg.FindAllStringSubmatch(content, -1)
	for _, data := range str{
		res += data[1]
	}
	return
}

//写入文件
func Store(content string, fileName, fileType string)  {
	dir, _ := os.Getwd()
	storePath := dir + "\\store\\"
	if _, err := os.Stat(storePath); err != nil {
		_ = os.MkdirAll(storePath, os.ModePerm)
		_ = os.Chmod(storePath, os.ModePerm)
	}

	f, err := os.OpenFile(storePath + fileName + fileType, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if !errorCheck(err, "os.Create") {
		return
	}
	_, _ = f.WriteString(content)
	f.Close()
}

func main() {
	start, end, ok  := getStdIn()
	if !ok {
		return
	}
	DoWork(start, end)

	//str := "<div class=\"content_read\">\n            <div class=\"box_con\">\n                <div class=\"con_top\">\n                    <script>textselect();</script>\n                    <a href=\"/\">新笔趣阁</a> &gt; <a href=\"/fenlei/2_1.html\">修真小说</a> &gt; <a href=\"http://www.xbiquge.la/7/7004/\">遮天</a> &gt;  第一章 星空中的青铜巨棺                </div>\n                <div class=\"bookname\">\n                    <h1> 第一章 星空中的青铜巨棺</h1>\n                    <div class=\"bottem1\">\n                        <a href=\"javascript:;\" onclick=\"showpop_vote(7004);\">投推荐票</a> <a href=\"/7/7004/\">上一章</a> &larr; <a href=\"http://www.xbiquge.la/7/7004/\">章节目录</a> &rarr; <a href=\"/7/7004/3246382.html\">下一章</a> <a href=\"javascript:;\" onclick=\"showpop_addmark(7004,3246381);\">加入书签</a>\n                    </div>\n                                        <div id=\"listtj\">&nbsp;推荐阅读：<a href=\"/9/9785/\" target=\"_blank\">剑卒过河</a>、<a href=\"/0/10/\" target=\"_blank\">武炼巅峰</a>、<a href=\"/26/26874/\" target=\"_blank\">沧元图</a>、<a href=\"/15/15409/\" target=\"_blank\">牧神记</a>、<a href=\"/26/26511/\" target=\"_blank\">剑徒之路</a>、<a href=\"/13/13959/\" target=\"_blank\">圣墟</a>、<a href=\"/2/2029/\" target=\"_blank\">极品透视</a>、<a href=\"/32/32626/\" target=\"_blank\">叶辰孙怡夏若雪</a>、<a href=\"/0/951/\" target=\"_blank\">伏天氏</a>、<a href=\"/1/1988/\" target=\"_blank\">龙城</a>、<a href=\"/30/30581/\" target=\"_blank\">顶级神豪</a>、<a href=\"/0/656/\" target=\"_blank\">莽荒纪</a>、<a href=\"/7/7552/\" target=\"_blank\">万古神帝</a>、<a href=\"/20/20948/\" target=\"_blank\">最佳女婿</a>、<a href=\"/2/2210/\" target=\"_blank\">全职法师</a></div>\n                </div>\n                <table style=\"width:100%; text-align:center;\"><tr><td><script>read_1_1();</script></td><td><script>read_1_2();</script></td><td><script>read_1_3();</script></td></tr></table>\n                <div id=\"content\">&nbsp;&nbsp;&nbsp;&nbsp;第一章星空中的青铜巨棺\n<br />\n<br />&nbsp;&nbsp;&nbsp;&nbsp;生命是世间最伟大的奇迹。\n<br />\n<br />&nbsp;&nbsp;&nbsp;&nbsp;四方上下曰宇。宇虽有实，而无定处可求。往古来今曰宙。宙虽有增长，不知其始之所至。</div>"
	//res := filter(str, "<h1>(?s:(.*?))</h1>")
	//fmt.Println(res)
}

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

//贴吧地址 "https://tieba.baidu.com/f?kw=%E8%B5%9B%E5%8D%9A%E6%9C%8B%E5%85%8B2077&ie=utf-8&pn=50"
var (
	basePath = "https://tieba.baidu.com/f?kw=%s&ie=utf-8&pn=" //基础路径
)

//获取命令行输入
func getStdIn() (start, end int, path string, ok bool){
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
	fmt.Println("请输入爬虫的贴吧关键词：")
	_, _ = fmt.Scan(&kw)
	path = formatUrl(kw)
	return
}

//预处理地址
func formatUrl(kw string)(path string){
	kw = url.QueryEscape(kw)
	path = fmt.Sprintf(basePath, kw)
	return
}

//处理任务
func DoWork(start, end int, path string)  {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go crawling(path, i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

//爬取数据
func crawling(path string, i int, page chan<- int)  {
	path += strconv.Itoa((i-1) * 50)
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

//写入文件
func Store(content string, fileName, fileType string)  {
	storePath := "D:\\practice\\go\\store\\"
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

func Spider() {
	start, end, path, ok := getStdIn()
	if !ok {
		return
	}

	DoWork(start, end, path)
}

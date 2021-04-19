package main

import (
	"net"
	"strings"
	"time"
	"pratice/src/util"
)

type Client struct {
	Ch 	 chan string //当前用户对应的数据管道
	Name string 	 //姓名
	Addr string 	 //IP
}

//在线用户集合
var OnLine map[string]Client
//群发消息管道
var message = make(chan string)

var isQuit = make(chan bool)
var hasData = make(chan bool)

//格式话化消息
func MakeMessage(client Client, msg string) string {
	return  "[" + client.Addr +"]"+ client.Name + ": "+ msg +"\n"
}

//处理连接
func HandleConnect(conn net.Conn)  {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	client := Client{Ch: make(chan string), Name: addr, Addr:addr}
	OnLine[addr] = client //保存当前连接用户信息进在线用户集合中

	//新建协程，为当前用户发送信息
	go SendMsg(client, conn)

	message <- MakeMessage(client, "Has Login!!!")//上线通知

	//新开一协程，持续读取用户发送数据
	go ReadData(client, conn)

	// 做个死循环，不要让方法结束
	for {
		//处理事件
		select {
			case <- isQuit :
				message <- MakeMessage(client, client.Name + ": Has Logout")//广播用户已退出
				delete(OnLine, addr)
				return
			case <- hasData:
			case <- time.After(10 * time.Second):
				message <- MakeMessage(client, "Time Out")
				delete(OnLine, addr)
				return
		}
	}
}

//持续读取用户发送数据
func ReadData(client Client, conn net.Conn)  {
	for{
		buf := make([]byte, 4*1024)
		n, err := conn.Read(buf)
		if err != nil {
			util.LogAndPrintln("Read", "Accept err", err)
			return
		}
		if n == 0 {
			isQuit <- true
			util.LogAndPrintln("Read", "Accept err", err)
			return
		}

		msg := string(buf[:n])
		msg = strings.Trim(msg,"\n")//去除nc工具的末尾\n
		msg = strings.TrimSpace(msg)
		
		//读取用户输入，根据输入，做出相应处理
		data := strings.ToLower(msg)
		switch data {
			case "exit":
				isQuit <- true
				return
			case "rename":
				client.Name = data
				OnLine[client.Addr] = client
				_, _ = conn.Write([]byte("rename success:" + data))
			case "who":
				for _, cli := range OnLine{
					_, _ = conn.Write([]byte(cli.Addr + "：" + cli.Name + "\n"))
				}
			default:
				message <- MakeMessage(client, msg) //读到数据 就写入消息管道， 广播协程从消息管道读取到数据，就会向在线用户群体发送消息
		}
		hasData <- true
	}
}

//向客户端发送消息
func SendMsg(client Client, conn net.Conn) {
	for msg := range client.Ch{ //当前用户的数据通道 读取数据，阻塞 直到获取到数据才开始处理
		_, _ = conn.Write([]byte(msg))
	}
}

//只要有消息来了，遍历map，给map中的每个成员广播此消息
func Broadcast()  {
	//给map分配空间
	OnLine = make(map[string]Client)

	for {
		msg := <- message
		for  _, client := range OnLine { //向在线用户的管道传入消息
			client.Ch <- msg
		}
	}
}

func main() {
	//建立端口监听
	listener, err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		util.LogAndPrintln("error", "Listen err", err)
		return
	}
	defer listener.Close()

	//为全部成员发送广播消息==》群发
	go Broadcast()

	//无限循环
	for {
		//阻塞等待用户连接
		conn, err := listener.Accept()
		if err != nil {
			util.LogAndPrintln("error", "Accept err", err)
			continue
		}

		//处理连接，每个连接 对应一个协程去处理
		go HandleConnect(conn)
	}
}


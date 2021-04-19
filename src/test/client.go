package test

import (
	"fmt"
	"net"
	"strings"
)

const (
	TCP = "tcp"
	TCP4 = "tcp4"
 	TCP6 = "tcp6"

 	ADDR = "127.0.0.1:"
 	PORT = "8088"
)


func main3() {
	//启动监听
	listener, err :=net.Listen(TCP, ADDR + PORT)
	if err != nil {
		fmt.Println("Listen监听连接错误：", err)
		return
	}
	fmt.Printf("服务端[%s]开始监听连接......\n", listener.Addr())
	fmt.Println("-----------------------------------------------")

	//关闭监听
	defer listener.Close()

	//接收多个用户请求连接
	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Accept接收连接错误：", err)
			//跳出本次错误，继续监听其他请求
			continue
		}

		//新建协程处理用户请求
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn)  {
	//关闭连接
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println("==============================================")
	fmt.Printf("服务端接收到来自[%s]的请求连接\n", addr)

	//读取接收到的数据
	buf := make([]byte, 1024)
	for {//死循环持续读取用户输入数据
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("Read服务端读取数据错误：", err)
			return
		}
		//转换内容为字符串
		content := string(buf[:n])
		content = strings.Trim(content,"\n")
		content = strings.TrimSpace(content)
		fmt.Printf("服务端接收到[%s]的内容是：%s\n", addr, content)

		//数据转换
		res := strings.ToUpper(content) + "\n"

		//将数据发送给用户
		_, err = conn.Write([]byte(res))
		if err != nil{
			fmt.Println("Write服务端回写数据错误：", err)
			return
		}

		if "exit" == strings.ToLower(content) {
			fmt.Printf("Close客户端[%s]请求关闭连接\n", addr)
			return
		}
	}

}
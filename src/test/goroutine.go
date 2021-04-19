package test

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type Study struct{
	Id int
	Name string
	Data []string
}

func main6() {
	fmt.Println("这是主协程")

	//申明管道
	ch := make(chan Study, 3)
	go func() {
		//关闭通道
		defer close(ch)

		for i := 5; i>0; i-- {
			fmt.Println("这是子协程......")

			data := Study{i,"HaHa"+strconv.Itoa(i), []string{"qq","dd","bb"}}
			ch <- data

			time.Sleep(time.Second)
		}
		runtime.Goexit()

	}()

	for{
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("主协程在等待数据......")

		//从通道读取数据
		data, ok := <- ch
		fmt.Println("主协程在阻塞等待数据......")
		if ok != false {
			fmt.Println("主协程在接收数据......")
			fmt.Printf("主协程接收到的数据: %+v\n", data)
		} else {
			return
		}
	}
}

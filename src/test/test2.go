package test

import "fmt"

//import操作 具有.(.操作后续使用包中的方法可不就包名)和_(_操作只调用该包中的init方法)
func init()  {
	fmt.Println("这是第一步")
}

func main14() {
	fmt.Println("这是第二步")
}

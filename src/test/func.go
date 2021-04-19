package test
import "fmt"

// 声明一个函数类型
type cb func(int) int

func main5() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
}

//函数作为参数
func testCallBack(x int, f cb) {
	res := f(x)
	fmt.Println("res = ", res)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

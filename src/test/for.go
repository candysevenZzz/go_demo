package test

import (
	"fmt"
	"time"
)

func main4() {
/*	sum := 1
	for ; sum <= 10; {
		fmt.Println("相加前：sum=", sum)
		sum += sum
		fmt.Println("相加后：sum=",sum)
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10{
		sum += sum
	}
	fmt.Println(sum)

	 a:= 0
	//无限循环
	for{
		if a < 100{
			a ++
		} else {
			break
		}
	}
	fmt.Println("i的值：", i)*/

	/* 定义局部变量 */
	//var i, j int
	start:= time.Now()

	fmt.Println("start: ", start)


	//for i=2; i < 1000000; i++ {
	//	for j=2; j <= (i/j); j++ {
	//		if i%j==0 {
	//			break; // 如果发现因子，则不是素数
	//		}
	//	}
	//	if j > (i/j) {
	//		fmt.Printf("%d  是素数\n", i);
	//	}
	//}
}

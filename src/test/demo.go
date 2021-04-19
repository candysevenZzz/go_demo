package test

import (
	"fmt"
)

func test2()  {
	var a = 100
		fmt.Printf("a= %d\n", a)
		fmt.Printf("a= %v\n", &a)

		var b *int
		b = &a
		fmt.Printf("b= %v\n", b)
		*b = 10
		fmt.Printf("b= %v, a= %v\n", b, a)
}

func Test () {
	var i *int
	fmt.Printf("i= %v\n", i)
	fmt.Println("i=", i)
}
package test

import "fmt"

func main9()  {
	var a int = 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */
	ip = &a  /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	p := *ip
	p1 := ip
	fmt.Printf("p 变量的值: %d\n", p )
	fmt.Printf("p1 变量的指针地址: %v\n", p1 )


	k := 1
	var l *int
	var ll **int

	l = &k
	ll = &l

	fmt.Println("k的值：", k)
	fmt.Println("k的内存地址：", &k)
	fmt.Println("*l的值：", *l)
	fmt.Println("l的指针地址：", l)
	fmt.Println("*ll的值：", **ll)
	fmt.Println("ll的指针地址：", *ll)
}

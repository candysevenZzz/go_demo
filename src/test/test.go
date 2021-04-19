package test

import "fmt"

var i int

func main12() {
/*	for{
		func(){
			i += 1
		}()

		if i > 10 {
			break
		} else {
			fmt.Println("i的值是：", i)
		}
	}*/


	k := 1
	var l *int
	// l := &k
	//var ll **int
	ll := &l
	var l1 *int
	// l := &k

	l = &k
	ll = &l
	l1 = &k
	fmt.Println("k的值：", k)
	fmt.Println("k的内存地址：", &k)
	fmt.Println("*l的值：", *l)
	fmt.Println("l的指针地址：", l)
	fmt.Println("*ll的值：", **ll)
	fmt.Println("ll的指针地址：", *ll)

	a := add(l, l1)
	fmt.Println("a=", a)
	fmt.Printf("l = %d, l1 = %d\n", *l  , *l1)
}

func add(a *int, b *int) int {
	fmt.Printf("a = %d, b = %d\n", *a  , *b)
	*b = 3
	fmt.Printf("a = %d, b = %d\n", *a  , *b)
	return *a + *b
}

package test

import "fmt"

type st struct {
	id   int
	name string
}

func main13() {
	d := st{1, "Jo"}
	fmt.Println(d, "值传递前")
	fValue(d)
	fmt.Println(d, "值传递后，外层值不变")
	fPoint(&d)
	fmt.Println(d, "指针传递后，外层值变化")
}

func fValue(s st) {
	s.id++
	s.name = "of"
	fmt.Println(s, "值传递函数")
}

func fPoint(s *st) {
	s.id++
	s.name = "of"
	fmt.Println(*s, "指针传递函数")
}
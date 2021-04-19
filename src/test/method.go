package test

import "fmt"

type test struct {
	id int
	name string
}

func (t test) add(param int) int {
	t.id = 3
	return t.id + param
}

func main8() {
	t := test{1,"aaa"}
	res := t.add(2)
	fmt.Println("res", res)
	fmt.Println("t", t.id)
}

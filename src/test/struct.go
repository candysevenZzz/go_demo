package test

import "fmt"

type book struct {
	id int
	name string
	page int
}

type student struct {
	book
	class string
}

func main11() {
	c1 := book{1,"西游记", 500}
	c2 := book{2,"红楼梦", 1000}

	fmt.Println("c1 == c2", c1 == c2)

	c3 := new(book)
	fmt.Println("c3", c3)

	s := student{}
	s.id = 1
	fmt.Println("s", s)

	//顺序初始化
	s1 := student{book{1,"haha",1},"hah"}
	fmt.Printf("%+v\n", s1)// %+v : 自动推导类型，显示更多的信息

	//指定初始化
	s2 :=  student{book:book{id: 1}}
	fmt.Printf("%+v\n", s2)
	fmt.Println(s2.page,s2.class,s2.id)


}
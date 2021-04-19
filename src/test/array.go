package test

import "fmt"

func main1() {
	//array:数组的申明方式 [n]type{} 其中长度n可以省略，用"..."来代替，golang会自动根据元素的个数来计算长度
	arr := [3]int{}
	var arr1 [6]int
	fmt.Println("arr = ", arr)
	fmt.Println("arr1 = ", arr1)
	arr[0] = 2
	arr[1] = 22
	arr[2] = 222
	fmt.Println("arr =", arr)

	//slice:切片 []type{}
	var s []int
	s1 := []int{1,2,3,4,5,6,7}
	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)

	//map:集合/映射/字典 var m map[key_type]value_type ;key_type:key的类型，value_type：值的类型
	//var m map[int]string
	m := make(map[int]string)
	m[1] = "haha"
	m[2] = "hello"
	m[3] = "1"
	for i:=0;i<10;i++ {
		m[i] =  fmt.Sprintf("%d", i)
	}
	fmt.Println("m = ", m)

	var s4 int
	s3 := [5]byte{'a','b'}
	fmt.Println("s3", s3)
	s4 = len(s3)
	fmt.Println("s4", s4)
}

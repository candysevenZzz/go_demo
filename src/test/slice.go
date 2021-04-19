package test

import "fmt"

func main10()  {
	var arr = [10]int {1,2,3,4,5,6,7,8,9,10}
	s := arr[2:5:6]
	s1 := append(s, 33,44,55,66,77,88,99)
	fmt.Println("s1 = ", s1)

	s2 := new(map[int]string)
	fmt.Println("s2 ", s2)
}

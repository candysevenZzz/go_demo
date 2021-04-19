package test

import (
	"fmt"
	"sync"
)
var count int

func sum(i int, wg *sync.WaitGroup)  {
	count += i
	fmt.Printf("第%d次计算时，count = %d\n", i, count)
	wg.Done()
}

func WaitGroup() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go sum(i, &wg)
	}
	wg.Wait()
	fmt.Println("count的总和是：", count)
}

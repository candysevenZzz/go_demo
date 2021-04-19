package test

import (
	"fmt"
	"runtime"
	"time"
)

func run(i, n uint64 , ch chan uint64){
	count := uint64(0)
	for ; i < n; i++ {
		count = count +i
	}
	ch <- count
}

func MultiCpuAdd(num uint64){
	t1 := time.Now()
	NCPU := runtime.NumCPU()
	fmt.Printf("CPU数量:%d\n", NCPU)
	runtime.GOMAXPROCS(NCPU)
	chs := make([] chan uint64, NCPU)
	for i := uint64(0); i < uint64(NCPU); i++{
		chs[i] = make(chan uint64)
		n := num /uint64(NCPU)
		go run(i*n, (i+1)*n, chs[i])
	}

	count := uint64(0)
	for i :=0; i < NCPU; i++{
		t := <- chs[i]
		count = count +t
	}

	t2 := time.Now()
	fmt.Printf("cpu num:%d,用时:%ds,count:%d\n", NCPU, t2.Sub(t1)/1000000000, count)
}
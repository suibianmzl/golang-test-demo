package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 两个 goroutine 同时操作的变量，竞争变量
	counter int
	waitGroup sync.WaitGroup
)

func incCount(int)  {
	defer waitGroup.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// 当期的goroutine 主动让出资源，从线程退出，并返会队列中
		runtime.Gosched()
		value++
		counter = value
	}
}

func main()  {
	runtime.GOMAXPROCS(1)
	waitGroup.Add(2)

	go incCount(1)
	go incCount(2)

	waitGroup.Wait()
	fmt.Println(counter)
}
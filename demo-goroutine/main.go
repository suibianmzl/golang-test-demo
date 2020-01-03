package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {

	// 1 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.NumCPU())

	// 2 设置等待器
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	fmt.Println("=========start=======")

	// 3 创建第一个 goroutine
	go func() {
		defer waitGroup.Done()

		// 打印3遍字母表
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	// 4 创建第二个 goroutine
	go func() {
		defer waitGroup.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	// 5 阻塞 main goroutine
	waitGroup.Wait()
	fmt.Println("========end========")
}
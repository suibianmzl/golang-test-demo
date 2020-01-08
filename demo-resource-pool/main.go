package main

import (
	"github.com/suibianmzl/golang-demos/demo-resource-pool/db"
	"github.com/suibianmzl/golang-demos/demo-resource-pool/pool"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (

	// 要使用的goroutines的数量
	maxGoroutines  = 5

	// 池中的资源的数量
	pooledResources = 2
)

func performQuery(query int, p *pool.Pool)  {
	// 1 获取连接
	conn, err := p.Acquire()

	if err != nil {
		log.Println("acquire conn error,", err)
		return
	}

	// 使用结束后，释放连接
	defer p.Release(conn)

	// 该 log 模拟对连接的使用
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d] \n", query, conn.(*db.DBConnection))

}

func main()  {
	var waitGroup sync.WaitGroup
	waitGroup.Add(maxGoroutines)

	// 1 创建一个pool
	p, err := pool.New(db.CreateConn, pooledResources)
	if err != nil {
		log.Println("create Pool error")
	}

	//2 开启 goroutine 执行任务
	for query := 0; query < maxGoroutines; query++ {
		// 每个goroutine需要自己复制一份要、查询值的副本，
		// 不然所有的查询会共享同一个查询变量，即所有的 goroutine 最后的 query 值都是3
		go func(q int) {
			performQuery(q, p)
			waitGroup.Done()
		}(query)
		//time.Sleep(1000*time.Millisecond) // 用于测试从 resources channel 中获取资源
	}

	// 3. 关闭连接池
	waitGroup.Wait()
	p.Close()
	log.Println("pool closed - main")
}
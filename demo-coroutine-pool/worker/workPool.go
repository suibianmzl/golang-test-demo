package worker

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	wg sync.WaitGroup
	// 工作池
	taskPool chan Worker
}

func New(maxGoroutineNum int) *Pool  {

	// 1. 初始化一个pool
	p := Pool{
		taskPool: make(chan Worker),
	}

	p.wg.Add(maxGoroutineNum)
	// 2. 创建 maxGoroutineNum 个 goroutine，并发的从 taskPool 中获取任务
	for i := 0; i < maxGoroutineNum; i ++ {
		go func() {
			for task := range p.taskPool { // 阻塞获取，一旦没有任务，阻塞在这里 - 无缓冲 channel
				// 3. 执行任务
				task.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

//  提交任务到worker池中
func (p *Pool) Run(worker Worker)()  {
	p.taskPool <- worker
}

// 关闭释放资源
func (p *Pool) Shutdown()  {
	close(p.taskPool)
	p.wg.Wait()
}
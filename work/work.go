/*
 * work 无缓冲的Pool
 * usage:
 *	nameWorker 是实现Worker接口的指针对象
 *	p := work.New(2)
 *	p.Run(&nameWorker)
 *	p.Shutdown()
 */
package main

import "sync"

// 声明一个worker 接口
type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{work: make(chan Worker)}

	p.wg.Add(maxGoroutines)

	// 使用 maxGoRoutines 个goroutine 来运行任务
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	// 返回取到的链接池对象-指针
	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

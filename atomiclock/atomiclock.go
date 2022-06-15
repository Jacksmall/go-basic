/*
 * 锁住共享资源-atomic 原子函数
 * atomic.AddInt64(&counter, 1)
 * 原子函数，原子操作，锁住共享资源
 * StoreInt64() 和 LoadInt64() 可以创建同步标识&&读取同步标识
 * 这个标识可以向程序的多个goroutine 通知某个特殊状态
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// counter int64
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	// runtime.GOMAXPROCS(1)
	// wg.Add(2)

	// go incCounter(1)
	// go incCounter(2)

	// wg.Wait()
	// fmt.Printf("Finished Counter: %d\n", counter)
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// 给定 goroutine 执行时间 1s
	time.Sleep(1 * time.Second)

	// g该停止工作了，安全的设置 shutdown 标识
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
	fmt.Println("Finished")
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// 要停止工作了吗？读取shutdown标识
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s down\n", name)
			break
		}
	}
}

// func incCounter(n int) {
// 	defer wg.Done()

// 	for i := 0; i < 2; i++ {
// 		// 原子函数，原子操作，锁住共享资源
// 		atomic.AddInt64(&counter, 1)
// 		// 当前goroutine退出，并回到队列中等待
// 		runtime.Gosched()
// 	}
// }

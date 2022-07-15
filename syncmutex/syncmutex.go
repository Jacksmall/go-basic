/*
 * 互斥锁 - sync.mutex
 * 同一时刻只允许一个goroutine 进入锁住的区域- 临界区
 * mutex.Lock()
 * 释放锁，允许其他goroutine 进入临界区
 * mutex.Unlock()
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {

	// runtime.GOMAXPROCS(1)

	wg.Add(2)

	go Increment(1)
	go Increment(2)

	fmt.Println("Waiting")
	wg.Wait()
	fmt.Printf("Finished counter:%d\n", counter)
}

func Increment(n int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		// 同一时刻只允许一个goroutine进入 这个临界区
		mutex.Lock()
		{
			value := counter
			// 将当前goroutine从线程中退出，并放回到队列
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}

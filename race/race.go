/*
 * 竞争状态 - race candition
 * 多个goroutine 在没有互相同步的情况下，访问共享的资源, 会相互竞争
 * 这种竞争状态极其容易出问题，所以要对共享资源进行读写操作必须是原子性的
 * 可以通过加锁来保证只有一个 goroutine 进行读写，完成后释放该锁，其他goroutine才能进行操作
 * 该例子模拟了 调度器 只能使用一个数量的 逻辑处理器，对共享资源counter进行读写
 * 本来预期结果counter 应该变为4 但是结果为2
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 共享的资源counter
	counter int
	wg      sync.WaitGroup
)

func main() {

	runtime.GOMAXPROCS(1)

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
		value := counter

		// 将当前goroutine从线程中退出，并放回到队列
		runtime.Gosched()

		value++

		counter = value
	}
}

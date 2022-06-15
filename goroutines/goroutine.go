/*
 * 通过模拟调度器只能使用一个逻辑处理器来执行goroutine
 * runtime.GOMAXPROCS()
 * sync.WaitGroup{}
 * go func() {}()
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 配置使用1个逻辑处理器来处理，底层会进行队列排队，使用 csp 理论 Goroutine 2 先执行；后 Goroutine 1 执行
	runtime.GOMAXPROCS(1)
	// 计数信号量
	wg := sync.WaitGroup{}
	// 计数器 2
	wg.Add(2)
	// 开始goroutine
	go func() {
		// 计数器 -1
		defer wg.Done()
		for count := 0; count < 2; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// 开始goroutine
	go func() {
		// 计数器 -1
		defer wg.Done()
		for count := 0; count < 2; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("等待程序完成...")
	// 等待计数器为0，则程序完成
	wg.Wait()

	fmt.Println("程序完成")
}

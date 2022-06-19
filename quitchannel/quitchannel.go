package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// func work(wg *sync.WaitGroup, quit chan bool) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		default:
// 			fmt.Println("Hello,world!")
// 		case <-quit:
// 			return
// 		}
// 	}
// }

// func main() {
// 	quit := make(chan bool)
// 	wg := sync.WaitGroup{}

// 	// 10个goroutines同时执行, quit 一旦关闭，通知所有的goroutines关闭退出
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go work(&wg, quit)
// 	}
// 	// 让 goroutines 奔跑1s
// 	time.Sleep(1 * time.Second)
// 	close(quit)
// 	wg.Wait()
// }

func work(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("Hello,world")
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(ctx, &wg)
	}

	time.Sleep(time.Second)
	wg.Wait()
}

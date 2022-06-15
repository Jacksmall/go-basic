package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 无缓冲的通道接力棒
	baton := make(chan int)

	// 为最后一位跑步者计数加1
	wg.Add(1)

	// 进入goroutine，第一位跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	// 等待比赛结束
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	// 等待接力棒
	runner := <-baton

	fmt.Printf("Runner %d Running With Baton\n", runner)

	// 创建下一个跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑步
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一个跑步者
	fmt.Printf("Runner %d Exchange baton To Runner %d\n", runner, newRunner)
	baton <- newRunner
}

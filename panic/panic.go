package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始了")
	c := boring()
	fmt.Println("boring()中任务完成了")
	// 接收到了通道事务完成的信号，接收通道里的信号并结束
	// dosomething
	<-c // receiver 等待boring 中的任务处理完成，discard the value
}

func boring() <-chan int {
	c := make(chan int)

	go func() {
		fmt.Println("goroutine中开始干活了")
		time.Sleep(1 * time.Second)
		res := 5
		fmt.Println("goroutine中活干完了")
		c <- res // sender 发送一个标识，值无所谓
	}()

	return c
}

// func boring(i int) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Printf("下面发生了错误，这里recover了, err: %v", err)
// 		}
// 	}()
// 	if i == 2 {
// 		panic("测试recovery")
// 	}
// 	fmt.Println(i)
// }

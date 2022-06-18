package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(6, ch)

	go Consumer(ch)
	// 通过运行时间无法保证稳定的输出结果
	// time.Sleep(5 * time.Second)
	// 通过用户Inturrept来保证稳定输出结果
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	fmt.Printf("quit(%v) \n", <-sig)
}

/*
 * goroutine 交替运行的例子 - 5000以内的素数
 * 当goroutine 1 执行需要耗费很多的时间，调度器会将该goroutine 等待
 * 给goroutine 2 执行的机会
 * PrintPrime 使用了 label next 写法，第一次见这种写法，要谨记 和 php 的过时goto 关键字类似
 */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// wg.Add(2)

	// go PrintPrime("A")
	// go PrintPrime("B")

	// fmt.Println("等待程序运行完成...")

	// wg.Wait()

	// fmt.Println("程序运行完成!")
	ch := GeneralizeNumber()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i, prime)
		ch = PrimeFilter(ch, prime)
	}
}

func PrintPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Printf("Completed %s\n", prefix)
}

func GeneralizeNumber() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)

	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()

	return out
}

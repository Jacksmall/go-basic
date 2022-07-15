package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(12)
	// T1:全是6
	for i := 0; i < 6; i++ {
		go func() {
			fmt.Println("T1:", i)
			wg.Done()
		}()
	}
	// T2:随机顺序0-5
	for i := 0; i < 6; i++ {
		go func(i int) {
			fmt.Println("T2:", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

package main

import "fmt"

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

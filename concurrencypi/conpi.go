package main

import (
	"fmt"
	"math"
)

func pi(n int) float64 {
	ch := make(chan float64)

	for i := 0; i < n; i++ {
		go term(ch, float64(i))
	}

	f := 0.0

	for j := 0; j < n; j++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, v float64) {
	ch <- 4 * math.Pow(-1, v) / (2*v + 1)
}

func main() {
	fmt.Println(pi(5000))
}

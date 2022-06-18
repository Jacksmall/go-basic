package main

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

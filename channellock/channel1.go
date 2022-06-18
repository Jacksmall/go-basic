package main

import "fmt"

var msg string
var done = make(chan bool)

func hello() {
	msg = "你好，世界"
	close(done)
}

func next() {
	go hello()
	<-done
	fmt.Println(msg)
}

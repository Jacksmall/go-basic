package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

var timeout = 3 * time.Second

func main() {

	// a := 2
	// b := &a // 0xc000018088
	// *b = 3
	// fmt.Println(b, a) // 0xc000018090 3

	// r := runner.New(timeout)

	// r.Add(createTest(), createTest(), createTest())

	// if err := r.Start(); err != nil {
	// 	switch err {
	// 	case runner.ErrTimeout:
	// 		log.Println("Terminal due to timeout")
	// 		os.Exit(1)
	// 	case runner.ErrInterrupt:
	// 		log.Println("Terminal due to interrupt")
	// 		os.Exit(2)
	// 	}
	// }

	// log.Println("Process End.")

	fmt.Println(os.Args[0] + " [flag]")

	fmt.Println(url.Parse("http://www.baidu.com/api/getAdd"))
}

func createTest() func(int) {
	return func(id int) {
		log.Printf("Task Process #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

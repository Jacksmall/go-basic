package main

import (
	"log"
	"os"
	"time"

	"github.com/Jacksmall/go-basic/runner"
)

var timeout = 3 * time.Second

func main() {
	r := runner.New(timeout)

	r.Add(createTest(), createTest(), createTest())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminal due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminal due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process End.")
}

func createTest() func(int) {
	return func(id int) {
		log.Printf("Task Process #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

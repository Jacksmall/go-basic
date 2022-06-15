package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./example <url>")
		os.Exit(-1)
	}
}

func main() {
	// curl
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}
	if r.StatusCode != http.StatusOK {
		log.Fatalf("Http status code: %d\n", r.StatusCode)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err.Error())
	}
}

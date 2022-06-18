package main

import (
	"fmt"
	"log"

	"github.com/Jacksmall/go-basic/greetings"
)

func main() {

	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(messages)

	message, err := greetings.Hello("Gradys")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("你好 %s\n", message)
}

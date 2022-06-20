package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/Jacksmall/go-basic/protobuf"
	"google.golang.org/protobuf/proto"
)

func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "Name:", p.Name)

	if p.Email != "" {
		fmt.Fprintln(w, "Email:", p.Email)
	}

	for _, phone := range p.Phones {
		switch phone.Type {
		case pb.Person_HOME:
			fmt.Fprint(w, " Home phone #:")
		case pb.Person_MOBILE:
			fmt.Fprint(w, " Mobile phone #:")
		case pb.Person_WORK:
			fmt.Fprint(w, " Work phone #:")
		}
		fmt.Fprintln(w, phone.Number)
	}
}

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("用法: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}

	fname := os.Args[1]
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("读取文件错误", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("格式化addressbook失败", err)
	}

	listPeople(os.Stdout, book)
}

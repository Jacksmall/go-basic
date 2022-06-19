package client

import (
	"fmt"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

const HelloServiceName = "path/to/pkg.HelloService"

var _HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, nil
	}
	return &HelloServiceClient{c}, nil
}

func (p *HelloServiceClient) Hello(msg string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", msg, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal("Calling:", err)
	}
	fmt.Println(reply)
}

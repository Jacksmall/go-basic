package helloservice

import (
	"log"
	"net"
	"net/rpc"
)

// 服务名称
const HelloServiceName = "path/to/pkg.HelloService"

// 服务接口
type HelloServiceInterface = interface {
	Hello(msg string, reply *string) error
}

// 注册服务(接口)
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

// 对于提供服务的rpc函数，第二个参数是结果回复-必须是指针
// GO语言RPC规则 : 方法只能有两个可序列化的参数， 并返回一个错误类型
func (s *HelloService) Hello(msg string, reply *string) error {
	*reply = "hello:" + msg
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("ListenTCP server: %v", err)
	}

	// 多个tcp连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("AcceptTCP server: %v", err)
		}

		go rpc.ServeConn(conn)
	}
}

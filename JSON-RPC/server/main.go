/*
 JSON-RPC(Remote Procedure Call)
 원격 프로시저(함수) 호출
네트워크를 통해 다른 컴퓨터에 있는 함수를 호출.
*/

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculator struct{}
type Args struct {
	A, B int
}

func (c *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
func (c *Calculator) Sub(args *Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}
func (c *Calculator) Mul(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func main() {
	calc := new(Calculator)

	rpc.Register(calc) //rpc에 등록
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error", err)
	}

	for { //err가 발생했을때 프로그램이 꺼지지 않고 계속 실행
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go jsonrpc.ServeConn((conn))
		fmt.Println("new client connected")
	}

}

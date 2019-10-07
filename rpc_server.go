package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"../trpc/arith"
)

var srv_addr = flag.String("l", "0.0.0.0:8888", "server url ,eg:0.0.0.0:8888 ")

func main() {
	flag.Parse()
	startServer()
}


func startServer() {
	l, err := net.Listen("tcp", *srv_addr)
	if err != nil {
		panic(err)
	}
	log.Printf("NewServer test RPC server listening on %s\n",*srv_addr)
	newServer := rpc.NewServer()
	arith.RegisterArithService(newServer,new(ArithImpl))

	newServer.Accept(l)
	//newServer.HandleHTTP("/foo", "/bar")
}

type ArithImpl int

func (t *ArithImpl)Add(a, b int) (result int, err error){
	result= a + b
	fmt.Printf("%d+%d=%d\n",a,b,result)
	return
}

func (t *ArithImpl) Mul(a, b int) (result int, err error){
	result =a * b
	fmt.Printf("%d*%d=%d\n",a,b,result)
	return
}

func (t *ArithImpl) Div(a, b int) (result int, err error){
	if b == 0 {
		err= errors.New("divide by zero")
		return
	}
	result = a / b
	fmt.Printf("%d/%d=%d\n",a,b,result)
	return
}
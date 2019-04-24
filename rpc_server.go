package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"trpc/arith"
)

func main() {
	startServer()
}


func startServer() {
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	log.Println("NewServer test RPC server listening on 127.0.0.1:8888")
	newServer := rpc.NewServer()
	arith.RegisterArithService(newServer,new(ArithImpl))

	newServer.Accept(l)
	//newServer.HandleHTTP("/foo", "/bar")
}

type ArithImpl int

func (t *ArithImpl)Add(a, b int) (result int, err error){
	result= a + b
	return
}

func (t *ArithImpl) Mul(a, b int) (result int, err error){
	result =a * b
	return
}

func (t *ArithImpl) Div(a, b int) (result int, err error){
	if b == 0 {
		err= errors.New("divide by zero")
		return
	}
	result = a / b
	return
}
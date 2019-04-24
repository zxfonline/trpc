package main

import (
	"fmt"
	"trpc/arith"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err!=nil{
		panic(err)
	}
	arithClient:=arith.NewArithClient(client)
	defer arithClient.Close()
	result,err:=arithClient.Add(1,2)
	if err!=nil{
		panic(err)
	}
	fmt.Println(result)

	result,err=arithClient.Mul(4,2)
	if err!=nil{
		panic(err)
	}
	fmt.Println(result)

	result,err=arithClient.Div(4,2)
	if err!=nil{
		panic(err)
	}
	fmt.Println(result)
}

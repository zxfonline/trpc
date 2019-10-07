package main

import (
	"flag"
	"fmt"
	"../trpc/arith"
	"net/rpc"
)
var addr = flag.String("l", "127.0.0.1:8888", "server url ,eg:127.0.0.1:8888 ")

func main() {
	flag.Parse()
	client, err := rpc.Dial("tcp", *addr)
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

package rpctest

import (
	"fmt"
	"log"
	"net/rpc"
)

const serverAddress = "localhost"

func RpcClientMain() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("rpc dial http error: ", err)
	}
	args := &Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("client call Args.Multiply error:", err)
	}
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)
}

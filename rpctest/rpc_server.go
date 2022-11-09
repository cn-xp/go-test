package rpctest

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	N, M int
}

func (a *Args) Multiply(args *Args, reply *int) error {
	*reply = args.M * args.N
	return nil
}

func RpcServerMain() {
	calc := new(Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Starting rpctest server error:", err)
	}
	go http.Serve(listener, nil)
	time.Sleep(5 * time.Second)
}

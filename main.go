//服务端
package main

import (
	_"fmt"
	"net/rpc"
	"net/http"
	"net"
	"sync"
	"go-based-rpc-demo/server"
)

func main() {
	//定义对象
	rpcArchi := new(server.RpcArchi)
	//注册对象
	rpc.Register(rpcArchi)
	//注册Http Handle
	rpc.HandleHTTP()
	//监听端口8899
	l, e := net.Listen("tcp", ":8899")
	if e != nil {
		panic(e)
	}

	go http.Serve(l, nil)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
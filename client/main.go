//客户端程序
package main

import (
	_"net/http"
	"fmt"
	"net/rpc"
	"go-based-rpc-demo/rpcr"
)

func main(){
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("DialHTTP出错")
		panic(err)
	}
	args := rpcr.InArgs{A:10}
	var reply rpcr.OutArgs
	err = client.Call("RpcArchi.Multiply", &args, &reply)
	if err != nil {
		fmt.Println("Call出错")
		panic(err)
	}
	fmt.Printf("结果:%d",reply.B)
}
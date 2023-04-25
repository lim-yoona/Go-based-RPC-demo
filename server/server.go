package server

import (
	_"fmt"
	"go-based-rpc-demo/rpcr"
)

//定义一个对象

type RpcArchi struct{}

func (self * RpcArchi) Multiply(x *rpcr.InArgs, y *rpcr.OutArgs) error {
	y.B = x.A * x.A
	return nil
}

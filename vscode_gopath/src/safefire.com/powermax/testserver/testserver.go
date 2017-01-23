package main

import (
	"fmt"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
	"safefire.com/powermax/rpc"
)

const (
	NetworkAddr = "127.0.0.1:19090"
)

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) Ping() (r *rpc.SimpleReturn, err error) {
	fmt.Println("调用Ping成功")
	r = rpc.NewSimpleReturn()
	r.Code = "200"
	r.Msg = "调用Ping成功"
	return r,nil
}

func (this *RpcServiceImpl) SendDoubleData(doubleValue []*rpc.DoubleValue) (r *rpc.SimpleReturn, err error){
	fmt.Println("调用SendDoubleData成功")
	r = rpc.NewSimpleReturn()
	r.Code = "200"
	r.Msg = "调用SendDoubleData成功"
	return r,nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := rpc.NewRpcServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}

package main

import (
	"flag"
	"os"

	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	_ "thriftExample/rpc_benchmark/thrift/proto/gen-go/genpack"
	"thriftExample/rpc_benchmark/thrift/proto/gen-go/genpack"
	"context"
)

type Hello struct{}

func (t *Hello) Say(ctx context.Context,name *genpack.BenchmarkMessage) (r *genpack.BenchmarkMessage, err error) {
	s := "OK"

	var i int32 = 100
	name.Field1 = s
	name.Field2 = i
	return name, nil
}

var host = flag.String("s", "0.0.0.0:8972", "listened ip and port")

func main() {
	flag.Parse()
	//var waitc = make(chan string,1)
	
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(*host)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &Hello{}
	processor := genpack.NewGreeterProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", *host)
	server.Serve()

}
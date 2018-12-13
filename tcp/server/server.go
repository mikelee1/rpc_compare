package main

import (
	"flag"
	"net/rpc"
	"net"
	"log"
	"thriftExample/rpc_benchmark/http/models"
)

type Hello struct{}

func (t *Hello) Say( args *models.BenchmarkMessage,reply *models.BenchmarkMessage) (err error) {
	s := "OK"
	var i int32 = 100
	reply.Field1 = s
	reply.Field2 = i
	return nil
}

func main() {
	flag.Parse()

	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)


	hello := new(Hello)
	rpc.Register(hello)
	rpc.Accept(inbound)
}
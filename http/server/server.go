package main

import (
	"flag"
	"net/rpc"
	"net/http"
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
	hello := new(Hello)
	rpc.Register(hello)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
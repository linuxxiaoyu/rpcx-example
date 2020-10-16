package main

import (
	"flag"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("reuseport", *addr)
}

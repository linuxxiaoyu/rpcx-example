package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith struct{}

// the second paramter is not a pointer
func (a *Arith) Mul(ctx context.Context, args Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Println("C=", reply.C)
	return nil
}

func main() {
	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"flag"
	"os"

	example "github.com/rpcxio/rpcx-examples"

	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "./rpcx.socket", "server address")

type Arith struct{}

func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	os.Remove(*addr)
	s := server.NewServer()
	if err := s.Register(new(Arith), ""); err != nil {
		panic(err)
	}

	err := s.Serve("unix", *addr)
	if err != nil {
		panic(err)
	}
}

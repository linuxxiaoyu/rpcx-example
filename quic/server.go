package main

import (
	"crypto/tls"
	"flag"
	"log"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Print(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	s := server.NewServer(server.WithTLSConfig(config))
	s.RegisterName("Arith", new(example.Arith), "")

	err = s.Serve("quic", *addr)
	if err != nil {
		panic(err)
	}
}

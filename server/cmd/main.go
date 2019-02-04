package main

import (
	"flag"
	"github.com/capnpVSgrpc/server/capnp"
	"github.com/capnpVSgrpc/server/grpc"
	"log"
)

func main() {

	cmd := flag.String("cmd", "grpc", "Protocol name grpc or capnp")

	flag.Parse()

	switch *cmd {
	case "grpc":
		gs := grpc.GrpcServer{}
		gs.Start()
	case "capnp":
		cs := capnp.CapnpServer{}
		cs.Start()
	default:
		log.Println("Error : Please insert the inputs!")
	}

}

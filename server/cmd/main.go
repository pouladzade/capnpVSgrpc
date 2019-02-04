package main

import (
	"github.com/capnpVSgrpc/server/capnp"
	"github.com/capnpVSgrpc/server/grpc"
)

func main() {

	gs := grpc.GrpcServer{}
	go gs.Start()

	cs := capnp.CapnpServer{}

	cs.Start()
}

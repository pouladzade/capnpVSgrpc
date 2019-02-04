package grpc

import (
	"context"
	"github.com/capnpVSgrpc/config"
	gb "github.com/capnpVSgrpc/protobuff"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	address string
}

func (gs *GrpcServer) EchoAccount(ctx context.Context, in *gb.Account) (*gb.Account, error) {
	return in, nil
}

func (gs *GrpcServer) Start() {
	log.Printf("Grpc Server is running on %s ...", config.GrpcAddress)
	lis, err := net.Listen(config.Network, config.GrpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gb.RegisterCoreBankingServer(s, &GrpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

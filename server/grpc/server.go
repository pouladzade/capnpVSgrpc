package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"context"
	gb "github.com/capnpVSgrpc/grpcbenchmark"
)

const(
	code = 123456
	message = "It's just a benchmark between gocapnp proto and gogrpc"
)

type grpcServer struct{
	address string
	network  string
	acc gb.Account

}

func (gs *grpcServer) SetAddress( add string){
	gs.address = add
}

func (gs *grpcServer) SetNetwork(network string){
	gs.network = network
}

func(gs *grpcServer) SetAccount( acc *gb.Account){
	gs.acc = *acc
}

func (gs *grpcServer) CreateAccount(ctx context.Context, in *gb.Account) (*gb.Result, error) {

	gs.SetAccount(in)
	return &gb.Result{}, nil
}

func (gs *grpcServer) GetAccountInfo(ctx context.Context, in *gb.AccountId) (*gb.Response, error) {

	return &gb.Response{}, nil
}

func (gs *grpcServer) runGrpcServer(){
	lis, err := net.Listen(gs.network, gs.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gb.RegisterCoreBankingServer(s, &grpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

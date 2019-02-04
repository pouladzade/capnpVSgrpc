package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/capnpVSgrpc/grpcbenchmark"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)
type grpcClient struct {}


func (gc *grpcClient) sendRequest() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCoreBankingClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateAccount(ctx, &pb.Account{Name: ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

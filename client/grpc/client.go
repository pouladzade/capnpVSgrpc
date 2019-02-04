package grpc

import (
	"context"
	"encoding/hex"
	"github.com/capnpVSgrpc/config"
	"log"
	"time"

	pb "github.com/capnpVSgrpc/protobuff"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	address string
}

func (gc *GrpcClient) Run() {

	var acc pb.Account

	acc.AccountId = &pb.AccountId{}
	acc.Name = config.AccountName
	acc.Balance = config.Balance
	acc.AccountId.AccountId = make([]byte, 32)
	acc.AccountId.AccountId, _ = hex.DecodeString(config.AccountId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	conn, err := grpc.Dial(config.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cb := pb.NewCoreBankingClient(conn)

	defer cancel()
	res, err := cb.EchoAccount(ctx, &acc)
	if err != nil {
		log.Fatalf("could not greet: %v, %v", err, res)
	}
	//log.Printf("\n\nGRPC Client Result :\n",hex.EncodeToString(res.AccountId.AccountId))

}

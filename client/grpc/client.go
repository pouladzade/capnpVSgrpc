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

func (gc *GrpcClient) Run(rnd int) int64 {

	var acc pb.Account

	acc.AccountId = &pb.AccountId{}
	acc.Name = config.AccountName
	acc.Balance = config.Balance
	acc.AccountId.AccountId = make([]byte, 32)
	acc.AccountId.AccountId, _ = hex.DecodeString(config.AccountId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	conn, err := grpc.Dial(config.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cb := pb.NewCoreBankingClient(conn)

	defer cancel()
	var nano int64
	for i := 0; i < rnd; i++ {
		t1 := time.Now()
		res, err := cb.EchoAccount(ctx, &acc)
		if err != nil {
			log.Fatalf("could not greet: %v, %v", err, res)
		}
		//log.Printf("\n\nGRPC Client Result :\n", hex.EncodeToString(res.AccountId.AccountId))
		t2 := time.Now()
		nano += t2.UnixNano() - t1.UnixNano()
	}

	return nano
}

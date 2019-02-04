package capnp

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/capnpVSgrpc/config"
	"log"
	"net"
	"time"
	"zombiezen.com/go/capnproto2/rpc"

	pb "github.com/capnpVSgrpc/capnpproto"
)

type CapnpClient struct {
}

func (gc *CapnpClient) Run(rnd int) int64 {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	conn, err := net.Dial(config.Network, config.CapnpAddress)

	if err != nil {
		log.Println(err.Error())
	}
	rpcconn := rpc.NewConn(rpc.StreamTransport(conn))
	if rpcconn == nil {
		log.Println(fmt.Errorf("Can not create RPC connection!"))
	}
	defer rpcconn.Close()
	cb := pb.CoreBanking{Client: rpcconn.Bootstrap(ctx)}
	var nano int64

	for i := 0; i < rnd; i++ {
		t1 := time.Now()
		res, err := cb.EchoAccount(ctx, func(p pb.CoreBanking_echoAccount_Params) error {
			acc, _ := p.NewAcc1()
			acc.SetName(config.AccountName)
			acc.SetBalance(config.Balance)
			accid, _ := hex.DecodeString(config.AccountId)
			acc.SetAccountId(accid)
			return nil
		}).Struct()

		if err != nil {
			log.Printf("\nResult : %v,\nError :%v", res, err)
		}
		t2 := time.Now()
		nano += t2.UnixNano() - t1.UnixNano()
		/*
			ac, _ := res.Acc2()
			acid, _ := ac.AccountId()
			log.Printf("\n\n CAPNP Client Result :\n", hex.EncodeToString(acid))
		*/
	}

	return nano
}

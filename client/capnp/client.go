package capnp

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/capnpVSgrpc/config"
	"log"
	"net"
	"zombiezen.com/go/capnproto2/rpc"

	pb "github.com/capnpVSgrpc/capnpbenchmark"
)

type CapnpClient struct {
}

func (gc *CapnpClient) Run() {

	ctx := context.Background()
	conn, err := net.Dial(config.Network, config.CapnpAddress)
	if err != nil {
		log.Println(err.Error())
	}
	rpcconn := rpc.NewConn(rpc.StreamTransport(conn))
	if rpcconn == nil {
		log.Println(fmt.Errorf("Can not create RPC connection!"))
	}
	cb := pb.CoreBanking{Client: rpcconn.Bootstrap(ctx)}

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
	//ac,_ := res.Acc2()
	//acid,_ := ac.AccountId()
	//log.Printf("\n\n CAPNP Client Result :\n",hex.EncodeToString(acid))

}

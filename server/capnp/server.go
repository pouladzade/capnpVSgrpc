package capnp

import (
	"fmt"
	crb "github.com/capnpVSgrpc/capnpproto"
	"github.com/capnpVSgrpc/config"
	"log"
	"net"
	"zombiezen.com/go/capnproto2/rpc"
)

type CapnpServer struct {
	acc crb.Account
}

func (cs *CapnpServer) EchoAccount(call crb.CoreBanking_echoAccount) error {
	acc, _ := call.Params.Acc1()
	call.Results.SetAcc2(acc)
	return nil
}

func (gs *CapnpServer) Start() {

	l, err := net.Listen(config.Network, config.CapnpAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Printf("Capnp Server is running on %s ...", config.CapnpAddress)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(conn, gs)
	}
}

func serve(c net.Conn, service *CapnpServer) {
	main := crb.CoreBanking_ServerToClient(service)
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
	err := conn.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
}

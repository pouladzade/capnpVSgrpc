package main

import (
	"flag"
	cc "github.com/capnpVSgrpc/client/capnp"
	gc "github.com/capnpVSgrpc/client/grpc"
	"log"
)

func main() {

	cmd := flag.String("cmd", "grpc", "Protocol name grpc or capnp")
	round := flag.Int("rnd", 1, "The number of repeats (Rounds) ")

	flag.Parse()

	switch *cmd {
	case "grpc":
		runGrpc(*round)
	case "capnp":
		runCapnp(*round)
	default:
		log.Println("Error : Please insert the inputs!")
	}

}

func runGrpc(rnd int) {
	// GRPC Benchmark
	clnt := gc.GrpcClient{}
	nano := clnt.Run(rnd)
	log.Printf("\nGRPC done in %d Round in Average Time  : %d Nano Second", rnd, nano/int64(rnd))
}

func runCapnp(rnd int) {

	// CAPNP Benchmark
	clnt := cc.CapnpClient{}
	nano := clnt.Run(rnd)
	log.Printf("\nCAPNP done in %d Round in Average Time : %d Nano Second", rnd, nano/int64(rnd))

}

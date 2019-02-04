package main

import (
	"flag"
	cc "github.com/capnpVSgrpc/client/capnp"
	gc "github.com/capnpVSgrpc/client/grpc"
	"log"
	"time"
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

func runGrpc(round int) {
	// GRPC Benchmark
	var nano int64
	gclnt := gc.GrpcClient{}
	for i := 0; i < round; i++ {
		t1 := time.Now()
		gclnt.Run()
		t2 := time.Now()
		nano += t2.UnixNano() - t1.UnixNano()
	}

	log.Printf("\nGRPC done in %d Round in Average Time  : %d Nano Second", round, nano/int64(round))

}

func runCapnp(round int) {

	// CAPNP Benchmark
	var nano int64
	cclnt := cc.CapnpClient{}
	for i := 0; i < round; i++ {
		t1 := time.Now()
		cclnt.Run()
		t2 := time.Now()
		nano += t2.UnixNano() - t1.UnixNano()
	}

	log.Printf("\nCAPNP done in %d Round in Average Time : %d Nano Second", round, nano/int64(round))

}

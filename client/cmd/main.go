package main

import (
	cc "github.com/capnpVSgrpc/client/capnp"
	gc "github.com/capnpVSgrpc/client/grpc"
	"github.com/capnpVSgrpc/config"
	"log"
	"time"
)

func main() {

	// GRPC Benchmark
	var nano1, nano2 int64
	gclnt := gc.GrpcClient{}
	for i := 0; i < config.SampleRate; i++ {
		t1 := time.Now()
		gclnt.Run()
		t2 := time.Now()
		nano1 += t2.UnixNano() - t1.UnixNano()
	}

	log.Printf("\nGRPC done in %d Round in Average Time  : %d Nano Second", config.SampleRate, nano1/config.SampleRate)

	// CAPNP Benchmark

	cclnt := cc.CapnpClient{}
	for i := 0; i < config.SampleRate; i++ {
		t3 := time.Now()
		cclnt.Run()
		t4 := time.Now()
		nano2 += t4.UnixNano() - t3.UnixNano()
	}

	log.Printf("\nCAPNP done in %d Round in Average Time : %d Nano Second", config.SampleRate, nano2/config.SampleRate)

}

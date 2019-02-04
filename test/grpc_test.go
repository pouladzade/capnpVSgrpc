package test

import (
	cc "github.com/capnpVSgrpc/client/capnp"
	gc "github.com/capnpVSgrpc/client/grpc"
	"testing"
)

func BenchmarkGrpc(b *testing.B) {
	client := gc.GrpcClient{}
	for i := 0; i < b.N; i++ {
		client.Run()
	}

}

func BenchmarkCapnp(b *testing.B) {
	client := cc.CapnpClient{}
	for i := 0; i < b.N; i++ {
		client.Run()
	}
}

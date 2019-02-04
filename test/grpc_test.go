package test

import (
	gc "github.com/capnpVSgrpc/client/grpc"
	"testing"
)

func BenchmarkGrpc(b *testing.B) {
	client := gc.GrpcClient{}
	for i := 0; i < b.N; i++ {
		client.Run()
	}

}

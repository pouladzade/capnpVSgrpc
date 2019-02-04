package test

import (
	cc "github.com/capnpVSgrpc/client/capnp"
	"testing"
)

func BenchmarkCapnp(b *testing.B) {
	client := cc.CapnpClient{}
	for i := 0; i < b.N; i++ {
		client.Run()
	}
}

```
BenchmarkCapnp-8   	   10000	    108113 ns/op
BenchmarkCapnp-8   	   10000	    103253 ns/op
```

```
You can run Server :
ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/server/build$ ./server --cmd="grpc"
ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/server/build$ ./server --cmd="capnp"

And then run the Client for each protocol and set the round for benchmark:

ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/client/build$ ./client --cmd="grpc" --rnd=100000
2019/02/04 23:57:18 
GRPC done in 100000 Round in Average Time  : 102678 Nano Second

ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/client/build$ ./client --cmd="capnp" --rnd=100000
2019/02/04 23:57:35 
CAPNP done in 100000 Round in Average Time : 99580 Nano Second
```

### As you can see The Capnp rpc can be just a little bit faster!!!!!!

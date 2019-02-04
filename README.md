## Compiling the code

You need to install [Go](https://golang.org/) (version 1.10.1 or higher)
<p>And <a href="https://capnproto.org/install.html" rel="nofollow">The Cap'n Proto tools</a>
<p>And  <a href="https://github.com/protocolbuffers/protobuf/releases" rel="nofollow"> the protobuf compiler (protoc)</a>.

After installing them, you can follow these steps to compile and build the cryptobank project:

```bash
mkdir -p $GOPATH/src/github.com/capnpVSgrpc
cd $GOPATH/src/github.com/capnpVSgrpc
git clone https://github.com/pouladzade/capnpVSgrpc.git
make
```
```
Two executable file will be created:

'client' in this directory:
"$GOPATH/src/github.com/capnpVSgrpc/client/build"

 and 'server' in this directory:
 $GOPATH/src/github.com/capnpVSgrpc/server/build
```

```
BenchmarkCapnp-8   	   10000	    108113 ns/op
BenchmarkCapnp-8   	   10000	    103253 ns/op
```

Run Server for each of them  in a separate terminal:
```

ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/server/build$ ./server --cmd="grpc"
ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/server/build$ ./server --cmd="capnp"
```

And then run the Client for each protocol and set the round for benchmark:
```
ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/client/build$ ./client --cmd="grpc" --rnd=100000
2019/02/04 23:57:18 
GRPC done in 100000 Round in Average Time  : 102678 Nano Second

ahmad@ahmad:~/go/src/github.com/capnpVSgrpc/client/build$ ./client --cmd="capnp" --rnd=100000
2019/02/04 23:57:35 
CAPNP done in 100000 Round in Average Time : 99580 Nano Second
```

### As you can see The Capnp rpc can be just a little bit faster!!!!!!

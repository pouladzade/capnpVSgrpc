
GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	gopkg.in/alecthomas/gometalinter.v2 \
	zombiezen.com/go/capnproto2/... \
	github.com/syndtr/goleveldb/leveldb \
	github.com/gogo/protobuf/protoc-gen-gogo\
	google.golang.org/grpc

PACKAGES=$(shell go list ./... | grep -v '/vendor/')
SERVER=-tags 'server'
CLIENT=-tags 'client'
LDFLAGS= -ldflags
CAPNPPROTO=$(GOPATH)/src/zombiezen.com/go/capnproto2/std
INCLUDE = -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/capnpVSgrpc/protobuff
all: tools fmt deps gen build

########################################
### Tools & dependencies
tools:
	@capnp --version || (echo "Install Capn'p first; https://capnproto.org/install.html"; false)
	@protoc --version || (echo "Install Protoc first; https://github.com/protocolbuffers/protobuf"; false)
	@echo "Installing tools"
	go get $(GOTOOLS)
	@gometalinter.v2 --install

deps:
	@echo "Cleaning vendors..."
	rm -rf vendor/
	@echo "Running dep..."
	dep ensure -v

########################################
### Build cryptobank
build:
	go build $(SERVER) -o server/build/server ./server/cmd/
	go build $(CLIENT) -o client/build/client ./client/cmd/

########################################
### Testing
test:
	go test $(PACKAGES)

########################################
### Generating capnp go files
gen:
	capnp compile -I $(CAPNPPROTO) -ogo ./capnpproto/capnpbenchmark.capnp
	--protoc $(INCLUDE) --gogo_out=plugins=grpc:. ./protobuff/grpcbenchmark.proto

########################################
### Formatting, linting, and vetting
fmt:
	@go fmt ./...

metalinter:
	@echo "--> Running linter"
	@gometalinter.v2 --vendor --deadline=600s --disable-all  \
		--enable=deadcode \
		--enable=gosimple \
	 	--enable=misspell \
		--enable=safesql \
		./...
		#--enable=gas \
		#--enable=maligned \
		#--enable=dupl \
		#--enable=errcheck \
		#--enable=goconst \
		#--enable=gocyclo \
		#--enable=goimports \
		#--enable=golint \ <== comments on anything exported
		#--enable=gotype \
	 	#--enable=ineffassign \
	   	#--enable=interfacer \
	   	#--enable=megacheck \
	   	#--enable=staticcheck \
	   	#--enable=structcheck \
	   	#--enable=unconvert \
	   	#--enable=unparam \
		#--enable=unused \
	   	#--enable=varcheck \
		#--enable=vet \
		#--enable=vetshadow \



.PHONY: tools fmt deps capnp build test
.PHONY: gen
.PHONY: build
.PHONY: tools deps
.PHONY: fmt metalinter
.PHONY: all


// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpcbenchmark/grpcbenchmark.proto

package grpcbenchmark

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AccountId struct {
	AccountId            []byte   `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountId) Reset()         { *m = AccountId{} }
func (m *AccountId) String() string { return proto.CompactTextString(m) }
func (*AccountId) ProtoMessage()    {}
func (*AccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed47da540bb4cb53, []int{0}
}
func (m *AccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountId.Unmarshal(m, b)
}
func (m *AccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountId.Marshal(b, m, deterministic)
}
func (m *AccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountId.Merge(m, src)
}
func (m *AccountId) XXX_Size() int {
	return xxx_messageInfo_AccountId.Size(m)
}
func (m *AccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountId.DiscardUnknown(m)
}

var xxx_messageInfo_AccountId proto.InternalMessageInfo

func (m *AccountId) GetAccountId() []byte {
	if m != nil {
		return m.AccountId
	}
	return nil
}

type Account struct {
	AccountId            *AccountId `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Balance              uint64     `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed47da540bb4cb53, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAccountId() *AccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountId)(nil), "grpcbenchmark.AccountId")
	proto.RegisterType((*Account)(nil), "grpcbenchmark.Account")
}

func init() { proto.RegisterFile("grpcbenchmark/grpcbenchmark.proto", fileDescriptor_ed47da540bb4cb53) }

var fileDescriptor_ed47da540bb4cb53 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2f, 0x2a, 0x48,
	0x4e, 0x4a, 0xcd, 0x4b, 0xce, 0xc8, 0x4d, 0x2c, 0xca, 0xd6, 0x47, 0xe1, 0xe9, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0xf1, 0xa2, 0x08, 0x4a, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x55, 0x25, 0x95, 0xa6, 0x81, 0x79, 0x60,
	0x0e, 0x98, 0x05, 0xd1, 0xad, 0xa4, 0xc9, 0xc5, 0xe9, 0x98, 0x9c, 0x9c, 0x5f, 0x9a, 0x57, 0xe2,
	0x99, 0x22, 0x24, 0xc3, 0xc5, 0x99, 0x08, 0xe3, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x21,
	0x04, 0x94, 0xf2, 0xb9, 0xd8, 0xa1, 0x4a, 0x85, 0xcc, 0xd0, 0x15, 0x72, 0x1b, 0x49, 0xe8, 0xa1,
	0x3a, 0x0e, 0x6e, 0x2a, 0x92, 0x11, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x52, 0x62, 0x4e, 0x62, 0x5e,
	0x72, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x8c, 0x6b, 0xe4, 0xc3, 0xc5, 0xed, 0x9c,
	0x5f, 0x94, 0xea, 0x94, 0x98, 0x97, 0x9d, 0x99, 0x97, 0x2e, 0x64, 0xcb, 0xc5, 0xed, 0x9a, 0x9c,
	0x91, 0x0f, 0x73, 0x83, 0x18, 0x76, 0x0b, 0xa5, 0x70, 0x88, 0x3b, 0x09, 0x46, 0xf1, 0xeb, 0xa1,
	0x06, 0x60, 0x12, 0x1b, 0x38, 0x0c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xec, 0x4e, 0x1a,
	0x01, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoreBankingClient is the client API for CoreBanking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreBankingClient interface {
	EchoAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
}

type coreBankingClient struct {
	cc *grpc.ClientConn
}

func NewCoreBankingClient(cc *grpc.ClientConn) CoreBankingClient {
	return &coreBankingClient{cc}
}

func (c *coreBankingClient) EchoAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/grpcbenchmark.CoreBanking/EchoAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreBankingServer is the server API for CoreBanking service.
type CoreBankingServer interface {
	EchoAccount(context.Context, *Account) (*Account, error)
}

func RegisterCoreBankingServer(s *grpc.Server, srv CoreBankingServer) {
	s.RegisterService(&_CoreBanking_serviceDesc, srv)
}

func _CoreBanking_EchoAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreBankingServer).EchoAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcbenchmark.CoreBanking/EchoAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreBankingServer).EchoAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoreBanking_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcbenchmark.CoreBanking",
	HandlerType: (*CoreBankingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoAccount",
			Handler:    _CoreBanking_EchoAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcbenchmark/grpcbenchmark.proto",
}

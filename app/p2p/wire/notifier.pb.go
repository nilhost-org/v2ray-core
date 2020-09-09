package wire

import (
	"context"
	"fmt"
	"math"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c0fc606bc4470de, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "wire.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "wire.HelloReply")
}

func init() { proto.RegisterFile("notifier.proto", fileDescriptor_1c0fc606bc4470de) }

var fileDescriptor_1c0fc606bc4470de = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x5d, 0x58, 0x74, 0x1d, 0x44, 0x24, 0x5e, 0x96, 0x9e, 0x24, 0x07, 0x11, 0xc1, 0x04,
	0xba, 0x0f, 0xe0, 0xd5, 0x83, 0x78, 0xd8, 0x3e, 0x41, 0xda, 0x4e, 0x63, 0x20, 0xed, 0xc4, 0x24,
	0x55, 0xfa, 0xf6, 0xd2, 0x69, 0x0b, 0x7b, 0x9b, 0x2f, 0xf9, 0xe6, 0x9f, 0x1f, 0xee, 0x07, 0xca,
	0xae, 0x73, 0x18, 0x55, 0x88, 0x94, 0x49, 0xec, 0xff, 0x5c, 0xc4, 0xe2, 0xcd, 0xba, 0xfc, 0x3d,
	0xd6, 0xaa, 0xa1, 0x5e, 0x5b, 0xb2, 0xa4, 0xf9, 0xb3, 0x1e, 0x3b, 0x26, 0x06, 0x9e, 0x96, 0xa5,
	0xe2, 0xf5, 0xb7, 0x8c, 0x66, 0x62, 0xbb, 0xa1, 0x88, 0xda, 0x84, 0xa0, 0x43, 0x19, 0xf4, 0x9c,
	0xa5, 0x13, 0x62, 0xeb, 0x5d, 0xca, 0x8b, 0x2b, 0x25, 0xdc, 0x7d, 0xa0, 0xf7, 0x74, 0xc6, 0x9f,
	0x11, 0x53, 0x16, 0x02, 0xf6, 0x83, 0xe9, 0xf1, 0xb8, 0x7b, 0xda, 0xbd, 0xdc, 0x9e, 0x79, 0x96,
	0xcf, 0x00, 0xab, 0x13, 0xfc, 0x24, 0x8e, 0x70, 0xd3, 0x63, 0x4a, 0xc6, 0x6e, 0xd2, 0x86, 0xe5,
	0x3b, 0x1c, 0xbe, 0xd6, 0xfa, 0xe2, 0x04, 0x87, 0xca, 0x4c, 0xbc, 0x26, 0x1e, 0xd5, 0x7c, 0x59,
	0x31, 0x54, 0x88, 0xed, 0xa7, 0x4b, 0xb9, 0x78, 0xb8, 0x78, 0xe4, 0x60, 0x79, 0x55, 0x5f, 0x73,
	0xa7, 0xd3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x0d, 0x8e, 0x32, 0x06, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifierClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloSeedList, opts ...grpc.CallOption) (*HelloReply, error)
}

type notifierClient struct {
	cc *grpc.ClientConn
}

func NewNotifierClient(cc *grpc.ClientConn) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) SayHello(ctx context.Context, in *HelloSeedList, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/wire.Notifier/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServer is the server API for Notifier service.
type NotifierServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloSeedList) (*HelloReply, error)
}

func RegisterNotifierServer(s *grpc.Server, srv NotifierServer) {
	s.RegisterService(&_Notifier_serviceDesc, srv)
}

func _Notifier_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloSeedList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.Notifier/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).SayHello(ctx, req.(*HelloSeedList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wire.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Notifier_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifier.proto",
}

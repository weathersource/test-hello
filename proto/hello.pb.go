// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SayHelloResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e10c42a6c052d6, []int{0}
}

func (m *SayHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloResponse.Unmarshal(m, b)
}
func (m *SayHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloResponse.Marshal(b, m, deterministic)
}
func (m *SayHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResponse.Merge(m, src)
}
func (m *SayHelloResponse) XXX_Size() int {
	return xxx_messageInfo_SayHelloResponse.Size(m)
}
func (m *SayHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResponse proto.InternalMessageInfo

func (m *SayHelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloResponse)(nil), "wxsrc.protobuf.SayHelloResponse")
}

func init() { proto.RegisterFile("hello/hello.proto", fileDescriptor_49e10c42a6c052d6) }

var fileDescriptor_49e10c42a6c052d6 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x7c, 0xe5, 0x15, 0xc5, 0x45,
	0xc9, 0x10, 0x4e, 0x52, 0x69, 0x9a, 0x94, 0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x4c,
	0x40, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0x22, 0xaf, 0xa4, 0xc2, 0x25, 0x10, 0x9c, 0x58, 0xe9,
	0x01, 0xd2, 0x1e, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc0, 0xc5, 0x9c, 0x5b,
	0x9c, 0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x1a, 0x85, 0x71, 0xf1, 0x80, 0x95,
	0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xb9, 0x71, 0x71, 0xc0, 0x74, 0x09, 0x89, 0xe9,
	0x41, 0xcc, 0x87, 0x5b, 0xa8, 0xe7, 0x0a, 0x32, 0x5f, 0x4a, 0x41, 0x0f, 0xd5, 0x1d, 0x7a, 0xe8,
	0xf6, 0x28, 0x31, 0x38, 0xe9, 0x46, 0x69, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x97, 0xa7, 0x26, 0x96, 0x64, 0xa4, 0x16, 0x15, 0xe7, 0x97, 0x16, 0x25, 0x43, 0x9d,
	0x5b, 0x94, 0x5a, 0x90, 0xaf, 0x9f, 0x0e, 0xf5, 0x5f, 0x12, 0x1b, 0x58, 0xcc, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x61, 0x8b, 0x8a, 0xeb, 0xf5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// SayHello implements hello.HelloService.
	SayHello(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/wxsrc.protobuf.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// SayHello implements hello.HelloService.
	SayHello(context.Context, *empty.Empty) (*SayHelloResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wxsrc.protobuf.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wxsrc.protobuf.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/hello.proto",
}

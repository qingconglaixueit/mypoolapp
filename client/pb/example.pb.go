// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestReq struct {
	Message              []byte   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReq) Reset()         { *m = TestReq{} }
func (m *TestReq) String() string { return proto.CompactTextString(m) }
func (*TestReq) ProtoMessage()    {}
func (*TestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_aecddcb727a5c294, []int{0}
}
func (m *TestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReq.Unmarshal(m, b)
}
func (m *TestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReq.Marshal(b, m, deterministic)
}
func (dst *TestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReq.Merge(dst, src)
}
func (m *TestReq) XXX_Size() int {
	return xxx_messageInfo_TestReq.Size(m)
}
func (m *TestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReq.DiscardUnknown(m)
}

var xxx_messageInfo_TestReq proto.InternalMessageInfo

func (m *TestReq) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type TestRsp struct {
	Message              []byte   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRsp) Reset()         { *m = TestRsp{} }
func (m *TestRsp) String() string { return proto.CompactTextString(m) }
func (*TestRsp) ProtoMessage()    {}
func (*TestRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_aecddcb727a5c294, []int{1}
}
func (m *TestRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRsp.Unmarshal(m, b)
}
func (m *TestRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRsp.Marshal(b, m, deterministic)
}
func (dst *TestRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRsp.Merge(dst, src)
}
func (m *TestRsp) XXX_Size() int {
	return xxx_messageInfo_TestRsp.Size(m)
}
func (m *TestRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRsp.DiscardUnknown(m)
}

var xxx_messageInfo_TestRsp proto.InternalMessageInfo

func (m *TestRsp) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*TestReq)(nil), "pb.testReq")
	proto.RegisterType((*TestRsp)(nil), "pb.testRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestsvrClient is the client API for Testsvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestsvrClient interface {
	// Say is simple request.
	Say(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestRsp, error)
}

type testsvrClient struct {
	cc *grpc.ClientConn
}

func NewTestsvrClient(cc *grpc.ClientConn) TestsvrClient {
	return &testsvrClient{cc}
}

func (c *testsvrClient) Say(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestRsp, error) {
	out := new(TestRsp)
	err := c.cc.Invoke(ctx, "/pb.testsvr/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestsvrServer is the server API for Testsvr service.
type TestsvrServer interface {
	// Say is simple request.
	Say(context.Context, *TestReq) (*TestRsp, error)
}

func RegisterTestsvrServer(s *grpc.Server, srv TestsvrServer) {
	s.RegisterService(&_Testsvr_serviceDesc, srv)
}

func _Testsvr_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestsvrServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.testsvr/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestsvrServer).Say(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Testsvr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.testsvr",
	HandlerType: (*TestsvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Testsvr_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_example_aecddcb727a5c294) }

var fileDescriptor_example_aecddcb727a5c294 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe6,
	0x62, 0x2f, 0x49, 0x2d, 0x2e, 0x09, 0x4a, 0x2d, 0x14, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71, 0xe1, 0x8a, 0x8a, 0x0b,
	0x70, 0x2b, 0x32, 0xd2, 0x81, 0x28, 0x2a, 0x2e, 0x2b, 0x12, 0x52, 0xe4, 0x62, 0x0e, 0x4e, 0xac,
	0x14, 0xe2, 0xd6, 0x2b, 0x48, 0xd2, 0x83, 0x9a, 0x2e, 0x85, 0xe0, 0x14, 0x17, 0x28, 0x31, 0x24,
	0xb1, 0x81, 0x9d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xed, 0xfd, 0x3a, 0x10, 0x93, 0x00,
	0x00, 0x00,
}

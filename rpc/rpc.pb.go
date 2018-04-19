// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	Engine
	Empty
*/
package rpc

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

type Engine struct {
	Binary         string   `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
	Listen         string   `protobuf:"bytes,2,opt,name=listen" json:"listen,omitempty"`
	Replicas       []string `protobuf:"bytes,3,rep,name=replicas" json:"replicas,omitempty"`
	EnableBackends []string `protobuf:"bytes,4,rep,name=enable_backends,json=enableBackends" json:"enable_backends,omitempty"`
}

func (m *Engine) Reset()                    { *m = Engine{} }
func (m *Engine) String() string            { return proto.CompactTextString(m) }
func (*Engine) ProtoMessage()               {}
func (*Engine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Engine) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *Engine) GetListen() string {
	if m != nil {
		return m.Listen
	}
	return ""
}

func (m *Engine) GetReplicas() []string {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *Engine) GetEnableBackends() []string {
	if m != nil {
		return m.EnableBackends
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Engine)(nil), "Engine")
	proto.RegisterType((*Empty)(nil), "Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LonghornLauncherService service

type LonghornLauncherServiceClient interface {
	UpgradeEngine(ctx context.Context, in *Engine, opts ...grpc.CallOption) (*Empty, error)
}

type longhornLauncherServiceClient struct {
	cc *grpc.ClientConn
}

func NewLonghornLauncherServiceClient(cc *grpc.ClientConn) LonghornLauncherServiceClient {
	return &longhornLauncherServiceClient{cc}
}

func (c *longhornLauncherServiceClient) UpgradeEngine(ctx context.Context, in *Engine, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/LonghornLauncherService/UpgradeEngine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LonghornLauncherService service

type LonghornLauncherServiceServer interface {
	UpgradeEngine(context.Context, *Engine) (*Empty, error)
}

func RegisterLonghornLauncherServiceServer(s *grpc.Server, srv LonghornLauncherServiceServer) {
	s.RegisterService(&_LonghornLauncherService_serviceDesc, srv)
}

func _LonghornLauncherService_UpgradeEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Engine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LonghornLauncherServiceServer).UpgradeEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LonghornLauncherService/UpgradeEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LonghornLauncherServiceServer).UpgradeEngine(ctx, req.(*Engine))
	}
	return interceptor(ctx, in, info, handler)
}

var _LonghornLauncherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LonghornLauncherService",
	HandlerType: (*LonghornLauncherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpgradeEngine",
			Handler:    _LonghornLauncherService_UpgradeEngine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xb1, 0x4b, 0xc7, 0x30,
	0x10, 0x85, 0xfd, 0x59, 0x4d, 0xed, 0x81, 0x0a, 0x19, 0x34, 0x74, 0x2a, 0x59, 0xec, 0xd4, 0x41,
	0x67, 0x17, 0xa1, 0x5b, 0xa7, 0x8a, 0xb3, 0x24, 0xe9, 0xd1, 0x06, 0xeb, 0x25, 0x5c, 0xab, 0xd0,
	0xc1, 0xff, 0x5d, 0x6c, 0x8a, 0xdb, 0x7d, 0x1f, 0xf7, 0xe0, 0x3d, 0x28, 0x38, 0xba, 0x26, 0x72,
	0x58, 0x83, 0xfe, 0x01, 0xd1, 0xd2, 0xe8, 0x09, 0xe5, 0x1d, 0x08, 0xeb, 0xc9, 0xf0, 0xa6, 0x4e,
	0xd5, 0xa9, 0x2e, 0xfa, 0x83, 0xfe, 0xfc, 0xec, 0x97, 0x15, 0x49, 0x9d, 0x27, 0x9f, 0x48, 0x96,
	0x70, 0xc5, 0x18, 0x67, 0xef, 0xcc, 0xa2, 0xb2, 0x2a, 0xab, 0x8b, 0xfe, 0x9f, 0xe5, 0x03, 0xdc,
	0x22, 0x19, 0x3b, 0xe3, 0xbb, 0x35, 0xee, 0x03, 0x69, 0x58, 0xd4, 0xc5, 0xfe, 0x72, 0x93, 0xf4,
	0xcb, 0x61, 0x75, 0x0e, 0x97, 0xed, 0x67, 0x5c, 0xb7, 0xc7, 0x67, 0xb8, 0xef, 0x02, 0x8d, 0x53,
	0x60, 0xea, 0xcc, 0x17, 0xb9, 0x09, 0xf9, 0x15, 0xf9, 0xdb, 0x3b, 0x94, 0x1a, 0xae, 0xdf, 0xe2,
	0xc8, 0x66, 0xc0, 0xa3, 0x69, 0xde, 0xa4, 0xa3, 0x14, 0xcd, 0x1e, 0xd6, 0x67, 0x56, 0xec, 0x6b,
	0x9e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xac, 0xea, 0x27, 0xda, 0x00, 0x00, 0x00,
}
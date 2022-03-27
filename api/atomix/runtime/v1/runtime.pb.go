// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/runtime/v1/runtime.proto

package v1

import (
	context "context"
	fmt "fmt"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("atomix/runtime/v1/runtime.proto", fileDescriptor_d426e124a0fc8e61) }

var fileDescriptor_d426e124a0fc8e61 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x2a, 0xcd, 0x2b, 0xc9, 0xcc, 0x4d, 0xd5, 0x2f, 0x33, 0x84, 0x31, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x21, 0x0a, 0xf4, 0x60, 0xa2, 0x65, 0x86, 0x46, 0x9c,
	0x5c, 0xec, 0x41, 0x10, 0x9e, 0x93, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x24, 0xb1, 0x81, 0xb5, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xfe, 0xb9, 0xc0, 0x61,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeClient is the client API for Runtime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeClient interface {
}

type runtimeClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeClient(cc *grpc.ClientConn) RuntimeClient {
	return &runtimeClient{cc}
}

// RuntimeServer is the server API for Runtime service.
type RuntimeServer interface {
}

// UnimplementedRuntimeServer can be embedded to have forward compatible implementations.
type UnimplementedRuntimeServer struct {
}

func RegisterRuntimeServer(s *grpc.Server, srv RuntimeServer) {
	s.RegisterService(&_Runtime_serviceDesc, srv)
}

var _Runtime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.runtime.v1.Runtime",
	HandlerType: (*RuntimeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "atomix/runtime/v1/runtime.proto",
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/runtime/v1/runtime.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Primitive struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Primitive) Reset()         { *m = Primitive{} }
func (m *Primitive) String() string { return proto.CompactTextString(m) }
func (*Primitive) ProtoMessage()    {}
func (*Primitive) Descriptor() ([]byte, []int) {
	return fileDescriptor_d426e124a0fc8e61, []int{0}
}
func (m *Primitive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Primitive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Primitive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Primitive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Primitive.Merge(m, src)
}
func (m *Primitive) XXX_Size() int {
	return m.Size()
}
func (m *Primitive) XXX_DiscardUnknown() {
	xxx_messageInfo_Primitive.DiscardUnknown(m)
}

var xxx_messageInfo_Primitive proto.InternalMessageInfo

func (m *Primitive) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreatePrimitiveRequest struct {
	Primitive Primitive `protobuf:"bytes,1,opt,name=primitive,proto3" json:"primitive"`
}

func (m *CreatePrimitiveRequest) Reset()         { *m = CreatePrimitiveRequest{} }
func (m *CreatePrimitiveRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePrimitiveRequest) ProtoMessage()    {}
func (*CreatePrimitiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d426e124a0fc8e61, []int{1}
}
func (m *CreatePrimitiveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePrimitiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePrimitiveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePrimitiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrimitiveRequest.Merge(m, src)
}
func (m *CreatePrimitiveRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreatePrimitiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrimitiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrimitiveRequest proto.InternalMessageInfo

func (m *CreatePrimitiveRequest) GetPrimitive() Primitive {
	if m != nil {
		return m.Primitive
	}
	return Primitive{}
}

type CreatePrimitiveResponse struct {
}

func (m *CreatePrimitiveResponse) Reset()         { *m = CreatePrimitiveResponse{} }
func (m *CreatePrimitiveResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePrimitiveResponse) ProtoMessage()    {}
func (*CreatePrimitiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d426e124a0fc8e61, []int{2}
}
func (m *CreatePrimitiveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePrimitiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePrimitiveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePrimitiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrimitiveResponse.Merge(m, src)
}
func (m *CreatePrimitiveResponse) XXX_Size() int {
	return m.Size()
}
func (m *CreatePrimitiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrimitiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrimitiveResponse proto.InternalMessageInfo

type DeletePrimitiveRequest struct {
	Primitive Primitive `protobuf:"bytes,1,opt,name=primitive,proto3" json:"primitive"`
}

func (m *DeletePrimitiveRequest) Reset()         { *m = DeletePrimitiveRequest{} }
func (m *DeletePrimitiveRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePrimitiveRequest) ProtoMessage()    {}
func (*DeletePrimitiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d426e124a0fc8e61, []int{3}
}
func (m *DeletePrimitiveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeletePrimitiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeletePrimitiveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeletePrimitiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePrimitiveRequest.Merge(m, src)
}
func (m *DeletePrimitiveRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeletePrimitiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePrimitiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePrimitiveRequest proto.InternalMessageInfo

func (m *DeletePrimitiveRequest) GetPrimitive() Primitive {
	if m != nil {
		return m.Primitive
	}
	return Primitive{}
}

type DeletePrimitiveResponse struct {
}

func (m *DeletePrimitiveResponse) Reset()         { *m = DeletePrimitiveResponse{} }
func (m *DeletePrimitiveResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePrimitiveResponse) ProtoMessage()    {}
func (*DeletePrimitiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d426e124a0fc8e61, []int{4}
}
func (m *DeletePrimitiveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeletePrimitiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeletePrimitiveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeletePrimitiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePrimitiveResponse.Merge(m, src)
}
func (m *DeletePrimitiveResponse) XXX_Size() int {
	return m.Size()
}
func (m *DeletePrimitiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePrimitiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePrimitiveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Primitive)(nil), "atomix.runtime.v1.Primitive")
	proto.RegisterType((*CreatePrimitiveRequest)(nil), "atomix.runtime.v1.CreatePrimitiveRequest")
	proto.RegisterType((*CreatePrimitiveResponse)(nil), "atomix.runtime.v1.CreatePrimitiveResponse")
	proto.RegisterType((*DeletePrimitiveRequest)(nil), "atomix.runtime.v1.DeletePrimitiveRequest")
	proto.RegisterType((*DeletePrimitiveResponse)(nil), "atomix.runtime.v1.DeletePrimitiveResponse")
}

func init() { proto.RegisterFile("atomix/runtime/v1/runtime.proto", fileDescriptor_d426e124a0fc8e61) }

var fileDescriptor_d426e124a0fc8e61 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x2a, 0xcd, 0x2b, 0xc9, 0xcc, 0x4d, 0xd5, 0x2f, 0x33, 0x84, 0x31, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x21, 0x0a, 0xf4, 0x60, 0xa2, 0x65, 0x86, 0x52, 0x22,
	0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0x49, 0x9e, 0x8b, 0x33, 0xa0,
	0x28, 0x33, 0x37, 0xb3, 0x24, 0xb3, 0x2c, 0x55, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x8a, 0xe2, 0x12, 0x73, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x85, 0x2b, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x72, 0xe0, 0xe2, 0x2c,
	0x80, 0x89, 0x81, 0xb5, 0x70, 0x1b, 0xc9, 0xe8, 0x61, 0xd8, 0xab, 0x07, 0xd7, 0xe7, 0xc4, 0x72,
	0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x42, 0x93, 0x92, 0x24, 0x97, 0x38, 0x86, 0xd9, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0x60, 0x6b, 0x5d, 0x52, 0x73, 0x52, 0x69, 0x65, 0x2d, 0x86, 0xd9, 0x10, 0x6b, 0x8d,
	0xee, 0x32, 0x72, 0xb1, 0x07, 0x41, 0x0c, 0x11, 0xca, 0xe0, 0xe2, 0x47, 0x73, 0x9d, 0x90, 0x26,
	0x16, 0x8b, 0xb0, 0x87, 0x8e, 0x94, 0x16, 0x31, 0x4a, 0x21, 0xb6, 0x82, 0x6c, 0x42, 0x73, 0x10,
	0x56, 0x9b, 0xb0, 0x07, 0x08, 0x56, 0x9b, 0x70, 0xf8, 0xcf, 0x49, 0xe2, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0xe9, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xc0, 0x17, 0x71, 0x29, 0x5b, 0x02, 0x00, 0x00,
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
	CreatePrimitive(ctx context.Context, in *CreatePrimitiveRequest, opts ...grpc.CallOption) (*CreatePrimitiveResponse, error)
	DeletePrimitive(ctx context.Context, in *DeletePrimitiveRequest, opts ...grpc.CallOption) (*DeletePrimitiveResponse, error)
}

type runtimeClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeClient(cc *grpc.ClientConn) RuntimeClient {
	return &runtimeClient{cc}
}

func (c *runtimeClient) CreatePrimitive(ctx context.Context, in *CreatePrimitiveRequest, opts ...grpc.CallOption) (*CreatePrimitiveResponse, error) {
	out := new(CreatePrimitiveResponse)
	err := c.cc.Invoke(ctx, "/atomix.runtime.v1.Runtime/CreatePrimitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeClient) DeletePrimitive(ctx context.Context, in *DeletePrimitiveRequest, opts ...grpc.CallOption) (*DeletePrimitiveResponse, error) {
	out := new(DeletePrimitiveResponse)
	err := c.cc.Invoke(ctx, "/atomix.runtime.v1.Runtime/DeletePrimitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeServer is the server API for Runtime service.
type RuntimeServer interface {
	CreatePrimitive(context.Context, *CreatePrimitiveRequest) (*CreatePrimitiveResponse, error)
	DeletePrimitive(context.Context, *DeletePrimitiveRequest) (*DeletePrimitiveResponse, error)
}

// UnimplementedRuntimeServer can be embedded to have forward compatible implementations.
type UnimplementedRuntimeServer struct {
}

func (*UnimplementedRuntimeServer) CreatePrimitive(ctx context.Context, req *CreatePrimitiveRequest) (*CreatePrimitiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrimitive not implemented")
}
func (*UnimplementedRuntimeServer) DeletePrimitive(ctx context.Context, req *DeletePrimitiveRequest) (*DeletePrimitiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePrimitive not implemented")
}

func RegisterRuntimeServer(s *grpc.Server, srv RuntimeServer) {
	s.RegisterService(&_Runtime_serviceDesc, srv)
}

func _Runtime_CreatePrimitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrimitiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).CreatePrimitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.runtime.v1.Runtime/CreatePrimitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).CreatePrimitive(ctx, req.(*CreatePrimitiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runtime_DeletePrimitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePrimitiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).DeletePrimitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.runtime.v1.Runtime/DeletePrimitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).DeletePrimitive(ctx, req.(*DeletePrimitiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Runtime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.runtime.v1.Runtime",
	HandlerType: (*RuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrimitive",
			Handler:    _Runtime_CreatePrimitive_Handler,
		},
		{
			MethodName: "DeletePrimitive",
			Handler:    _Runtime_DeletePrimitive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atomix/runtime/v1/runtime.proto",
}

func (m *Primitive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Primitive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Primitive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRuntime(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreatePrimitiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePrimitiveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePrimitiveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Primitive.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRuntime(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CreatePrimitiveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePrimitiveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePrimitiveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *DeletePrimitiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeletePrimitiveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeletePrimitiveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Primitive.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRuntime(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DeletePrimitiveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeletePrimitiveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeletePrimitiveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintRuntime(dAtA []byte, offset int, v uint64) int {
	offset -= sovRuntime(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Primitive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRuntime(uint64(l))
	}
	return n
}

func (m *CreatePrimitiveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Primitive.Size()
	n += 1 + l + sovRuntime(uint64(l))
	return n
}

func (m *CreatePrimitiveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DeletePrimitiveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Primitive.Size()
	n += 1 + l + sovRuntime(uint64(l))
	return n
}

func (m *DeletePrimitiveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovRuntime(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRuntime(x uint64) (n int) {
	return sovRuntime(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Primitive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRuntime
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Primitive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Primitive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRuntime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRuntime
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRuntime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRuntime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRuntime
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreatePrimitiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRuntime
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatePrimitiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePrimitiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Primitive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRuntime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRuntime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRuntime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Primitive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRuntime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRuntime
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreatePrimitiveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRuntime
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreatePrimitiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePrimitiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRuntime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRuntime
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeletePrimitiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRuntime
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeletePrimitiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeletePrimitiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Primitive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRuntime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRuntime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRuntime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Primitive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRuntime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRuntime
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeletePrimitiveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRuntime
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeletePrimitiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeletePrimitiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRuntime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRuntime
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRuntime(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRuntime
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRuntime
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRuntime
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRuntime
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRuntime
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRuntime
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRuntime        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRuntime          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRuntime = fmt.Errorf("proto: unexpected end of group")
)

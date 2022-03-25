// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/runtime/v1/agent.proto

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

type ConfigureAgentRequest struct {
}

func (m *ConfigureAgentRequest) Reset()         { *m = ConfigureAgentRequest{} }
func (m *ConfigureAgentRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigureAgentRequest) ProtoMessage()    {}
func (*ConfigureAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7013fe0e18fc06de, []int{0}
}
func (m *ConfigureAgentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigureAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigureAgentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigureAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureAgentRequest.Merge(m, src)
}
func (m *ConfigureAgentRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigureAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureAgentRequest proto.InternalMessageInfo

type ConfigureAgentResponse struct {
}

func (m *ConfigureAgentResponse) Reset()         { *m = ConfigureAgentResponse{} }
func (m *ConfigureAgentResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigureAgentResponse) ProtoMessage()    {}
func (*ConfigureAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7013fe0e18fc06de, []int{1}
}
func (m *ConfigureAgentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigureAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigureAgentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigureAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureAgentResponse.Merge(m, src)
}
func (m *ConfigureAgentResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConfigureAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureAgentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigureAgentRequest)(nil), "atomix.runtime.v1.ConfigureAgentRequest")
	proto.RegisterType((*ConfigureAgentResponse)(nil), "atomix.runtime.v1.ConfigureAgentResponse")
}

func init() { proto.RegisterFile("atomix/runtime/v1/agent.proto", fileDescriptor_7013fe0e18fc06de) }

var fileDescriptor_7013fe0e18fc06de = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x2a, 0xcd, 0x2b, 0xc9, 0xcc, 0x4d, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0x4c,
	0x4f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x48, 0xeb, 0x41, 0xa5,
	0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44, 0xa1,
	0x92, 0x38, 0x97, 0xa8, 0x73, 0x7e, 0x5e, 0x5a, 0x66, 0x7a, 0x69, 0x51, 0xaa, 0x23, 0xc8, 0x80,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x09, 0x2e, 0x31, 0x74, 0x89, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0xa3, 0x4c, 0x2e, 0x56, 0xb0, 0x80, 0x50, 0x02, 0x17, 0x27, 0x5c, 0x89, 0x90,
	0x86, 0x1e, 0x86, 0x95, 0x7a, 0x58, 0x4d, 0x96, 0xd2, 0x24, 0x42, 0x25, 0xc4, 0x2a, 0x27, 0x89,
	0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39,
	0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x3b, 0xdf, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x63, 0xd2, 0x6b, 0x0b, 0x08, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	Configure(ctx context.Context, in *ConfigureAgentRequest, opts ...grpc.CallOption) (*ConfigureAgentResponse, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Configure(ctx context.Context, in *ConfigureAgentRequest, opts ...grpc.CallOption) (*ConfigureAgentResponse, error) {
	out := new(ConfigureAgentResponse)
	err := c.cc.Invoke(ctx, "/atomix.runtime.v1.Agent/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	Configure(context.Context, *ConfigureAgentRequest) (*ConfigureAgentResponse, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) Configure(ctx context.Context, req *ConfigureAgentRequest) (*ConfigureAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.runtime.v1.Agent/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Configure(ctx, req.(*ConfigureAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.runtime.v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _Agent_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atomix/runtime/v1/agent.proto",
}

func (m *ConfigureAgentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigureAgentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigureAgentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ConfigureAgentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigureAgentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigureAgentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintAgent(dAtA []byte, offset int, v uint64) int {
	offset -= sovAgent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfigureAgentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ConfigureAgentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovAgent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigureAgentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
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
			return fmt.Errorf("proto: ConfigureAgentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigureAgentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAgent
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
func (m *ConfigureAgentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
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
			return fmt.Errorf("proto: ConfigureAgentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigureAgentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAgent
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
func skipAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
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
					return 0, ErrIntOverflowAgent
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
					return 0, ErrIntOverflowAgent
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
				return 0, ErrInvalidLengthAgent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAgent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAgent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAgent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAgent = fmt.Errorf("proto: unexpected end of group")
)

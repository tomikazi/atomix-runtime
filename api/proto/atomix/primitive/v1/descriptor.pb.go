// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/primitive/v1/descriptor.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// PartitionStrategy is an enum for indicating the strategy used to partition a primitive
type PartitionStrategy int32

const (
	PartitionStrategy_NONE        PartitionStrategy = 0
	PartitionStrategy_HASH        PartitionStrategy = 1
	PartitionStrategy_RANGE       PartitionStrategy = 2
	PartitionStrategy_RANDOM      PartitionStrategy = 3
	PartitionStrategy_ROUND_ROBIN PartitionStrategy = 4
)

var PartitionStrategy_name = map[int32]string{
	0: "NONE",
	1: "HASH",
	2: "RANGE",
	3: "RANDOM",
	4: "ROUND_ROBIN",
}

var PartitionStrategy_value = map[string]int32{
	"NONE":        0,
	"HASH":        1,
	"RANGE":       2,
	"RANDOM":      3,
	"ROUND_ROBIN": 4,
}

func (x PartitionStrategy) String() string {
	return proto.EnumName(PartitionStrategy_name, int32(x))
}

func (PartitionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a92601157dbc5a3a, []int{0}
}

// OperationType is an enum for specifying the type of operation
type OperationType int32

const (
	OperationType_COMMAND OperationType = 0
	OperationType_QUERY   OperationType = 1
	OperationType_CREATE  OperationType = 2
	OperationType_CLOSE   OperationType = 3
)

var OperationType_name = map[int32]string{
	0: "COMMAND",
	1: "QUERY",
	2: "CREATE",
	3: "CLOSE",
}

var OperationType_value = map[string]int32{
	"COMMAND": 0,
	"QUERY":   1,
	"CREATE":  2,
	"CLOSE":   3,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a92601157dbc5a3a, []int{1}
}

// AggregateStrategy is an enum for indicating the strategy used to aggregate a field
type AggregateStrategy int32

const (
	AggregateStrategy_CHOOSE_FIRST AggregateStrategy = 0
	AggregateStrategy_APPEND       AggregateStrategy = 1
	AggregateStrategy_SUM          AggregateStrategy = 2
)

var AggregateStrategy_name = map[int32]string{
	0: "CHOOSE_FIRST",
	1: "APPEND",
	2: "SUM",
}

var AggregateStrategy_value = map[string]int32{
	"CHOOSE_FIRST": 0,
	"APPEND":       1,
	"SUM":          2,
}

func (x AggregateStrategy) String() string {
	return proto.EnumName(AggregateStrategy_name, int32(x))
}

func (AggregateStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a92601157dbc5a3a, []int{2}
}

var E_PrimitiveType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "atomix.primitive.v1.primitive_type",
	Tag:           "bytes,50000,opt,name=primitive_type",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_Partitioned = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50001,
	Name:          "atomix.primitive.v1.partitioned",
	Tag:           "varint,50001,opt,name=partitioned",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_PartitionStrategy = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*PartitionStrategy)(nil),
	Field:         71000,
	Name:          "atomix.primitive.v1.partition_strategy",
	Tag:           "varint,71000,opt,name=partition_strategy,enum=atomix.primitive.v1.PartitionStrategy",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_PartitionKey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         72000,
	Name:          "atomix.primitive.v1.partition_key",
	Tag:           "varint,72000,opt,name=partition_key",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_PartitionRange = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         72001,
	Name:          "atomix.primitive.v1.partition_range",
	Tag:           "varint,72001,opt,name=partition_range",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_OperationName = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         61000,
	Name:          "atomix.primitive.v1.operation_name",
	Tag:           "bytes,61000,opt,name=operation_name",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_OperationType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*OperationType)(nil),
	Field:         61001,
	Name:          "atomix.primitive.v1.operation_type",
	Tag:           "varint,61001,opt,name=operation_type,enum=atomix.primitive.v1.OperationType",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_OperationAsync = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         61002,
	Name:          "atomix.primitive.v1.operation_async",
	Tag:           "varint,61002,opt,name=operation_async",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_OperationId = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         61003,
	Name:          "atomix.primitive.v1.operation_id",
	Tag:           "varint,61003,opt,name=operation_id",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

var E_OperationAggregate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*AggregateStrategy)(nil),
	Field:         62001,
	Name:          "atomix.primitive.v1.operation_aggregate",
	Tag:           "varint,62001,opt,name=operation_aggregate,enum=atomix.primitive.v1.AggregateStrategy",
	Filename:      "atomix/primitive/v1/descriptor.proto",
}

func init() {
	proto.RegisterEnum("atomix.primitive.v1.PartitionStrategy", PartitionStrategy_name, PartitionStrategy_value)
	proto.RegisterEnum("atomix.primitive.v1.OperationType", OperationType_name, OperationType_value)
	proto.RegisterEnum("atomix.primitive.v1.AggregateStrategy", AggregateStrategy_name, AggregateStrategy_value)
	proto.RegisterExtension(E_PrimitiveType)
	proto.RegisterExtension(E_Partitioned)
	proto.RegisterExtension(E_PartitionStrategy)
	proto.RegisterExtension(E_PartitionKey)
	proto.RegisterExtension(E_PartitionRange)
	proto.RegisterExtension(E_OperationName)
	proto.RegisterExtension(E_OperationType)
	proto.RegisterExtension(E_OperationAsync)
	proto.RegisterExtension(E_OperationId)
	proto.RegisterExtension(E_OperationAggregate)
}

func init() {
	proto.RegisterFile("atomix/primitive/v1/descriptor.proto", fileDescriptor_a92601157dbc5a3a)
}

var fileDescriptor_a92601157dbc5a3a = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcb, 0x6e, 0xd3, 0x4e,
	0x14, 0xc6, 0xed, 0x26, 0xbd, 0xe4, 0xe4, 0x36, 0x71, 0x37, 0xd5, 0x5f, 0xfa, 0x9b, 0x08, 0x21,
	0x54, 0x65, 0xe1, 0xa8, 0x65, 0xe7, 0x05, 0x92, 0xeb, 0xb8, 0x4d, 0x04, 0xf1, 0x84, 0x71, 0xb2,
	0x60, 0x15, 0xb9, 0xc9, 0x60, 0xac, 0x36, 0xb1, 0xe5, 0x0c, 0xa1, 0x79, 0x01, 0xd6, 0x3c, 0x07,
	0x3b, 0x78, 0x02, 0xd8, 0x15, 0xd8, 0x84, 0x5d, 0x17, 0x5d, 0xa0, 0x84, 0x07, 0x41, 0xe3, 0xd4,
	0x76, 0x20, 0x15, 0xde, 0xf9, 0x72, 0xbe, 0xdf, 0xf7, 0xcd, 0x39, 0x67, 0xe0, 0x91, 0xcd, 0xbc,
	0x91, 0x7b, 0x55, 0xf7, 0x03, 0x77, 0xe4, 0x32, 0x77, 0x4a, 0xeb, 0xd3, 0xa3, 0xfa, 0x90, 0x4e,
	0x06, 0x81, 0xeb, 0x33, 0x2f, 0x50, 0xfc, 0xc0, 0x63, 0x9e, 0xb4, 0xbf, 0xaa, 0x52, 0xe2, 0x2a,
	0x65, 0x7a, 0xf4, 0x5f, 0xd5, 0xf1, 0x3c, 0xe7, 0x92, 0xd6, 0xc3, 0x92, 0xf3, 0x37, 0xaf, 0x36,
	0x64, 0x35, 0x0c, 0x95, 0x8e, 0x1d, 0x30, 0x97, 0xb9, 0xde, 0xd8, 0x62, 0x81, 0xcd, 0xa8, 0x33,
	0x93, 0xf6, 0x20, 0x6b, 0x62, 0xd3, 0x40, 0x02, 0x7f, 0x6a, 0x6a, 0x56, 0x13, 0x89, 0x52, 0x0e,
	0xb6, 0x89, 0x66, 0x9e, 0x19, 0x68, 0x4b, 0x02, 0xd8, 0x21, 0x9a, 0xd9, 0xc0, 0x6d, 0x94, 0x91,
	0xca, 0x90, 0x27, 0xb8, 0x67, 0x36, 0xfa, 0x04, 0x9f, 0xb4, 0x4c, 0x94, 0xad, 0x3d, 0x85, 0x22,
	0xf6, 0x69, 0x60, 0x73, 0x60, 0x77, 0xe6, 0x53, 0x29, 0x0f, 0xbb, 0x3a, 0x6e, 0xb7, 0x35, 0xb3,
	0x81, 0x04, 0x4e, 0x79, 0xd1, 0x33, 0xc8, 0x4b, 0x24, 0x72, 0x8a, 0x4e, 0x0c, 0xad, 0xcb, 0x89,
	0x39, 0xd8, 0xd6, 0x9f, 0x63, 0xcb, 0x40, 0x99, 0x9a, 0x0a, 0x15, 0xcd, 0x71, 0x02, 0xea, 0xd8,
	0x8c, 0xc6, 0x81, 0x10, 0x14, 0xf4, 0x26, 0xc6, 0x96, 0xd1, 0x3f, 0x6d, 0x11, 0xab, 0x8b, 0x04,
	0xae, 0xd6, 0x3a, 0x1d, 0xc3, 0x6c, 0x20, 0x51, 0xda, 0x85, 0x8c, 0xd5, 0x6b, 0xa3, 0x2d, 0xb5,
	0x09, 0xa5, 0xf8, 0xf8, 0x7d, 0xc6, 0xcd, 0x1f, 0x28, 0xab, 0x0e, 0x28, 0x51, 0x07, 0x14, 0x8b,
	0x06, 0x53, 0x77, 0x40, 0xb1, 0xcf, 0x03, 0x4e, 0x0e, 0xe6, 0xef, 0x32, 0x55, 0xf1, 0x30, 0x47,
	0x8a, 0xb1, 0x90, 0x87, 0x56, 0x75, 0xc8, 0xfb, 0x51, 0x5b, 0xe8, 0x30, 0x1d, 0xf3, 0x23, 0xc4,
	0xec, 0x91, 0x75, 0x95, 0xfa, 0x16, 0xa4, 0xf8, 0xb5, 0x3f, 0x89, 0xce, 0x22, 0x6f, 0xb0, 0xda,
	0x94, 0xbd, 0xf6, 0x86, 0x11, 0xea, 0xe6, 0x43, 0xb6, 0x2a, 0x1e, 0x96, 0x8e, 0x1f, 0x2b, 0xf7,
	0x4c, 0x54, 0xd9, 0x18, 0x16, 0xa9, 0xf8, 0x7f, 0x7f, 0x52, 0x1b, 0x50, 0x4c, 0x8c, 0x2f, 0xe8,
	0x4c, 0xfa, 0x7f, 0xc3, 0xf3, 0xd4, 0xa5, 0x97, 0xb1, 0xe5, 0xe7, 0x4f, 0xd9, 0x30, 0x7d, 0x21,
	0x56, 0x3d, 0xa3, 0x33, 0xb5, 0x09, 0xe5, 0x84, 0x12, 0xd8, 0x63, 0x87, 0xa6, 0x71, 0xbe, 0xdc,
	0x71, 0x4a, 0xb1, 0x8e, 0x70, 0x99, 0x7a, 0x06, 0x25, 0x2f, 0xda, 0x89, 0xfe, 0xd8, 0x1e, 0xd1,
	0xd4, 0x26, 0x5c, 0xdf, 0xde, 0x8d, 0x25, 0xd6, 0x99, 0xf6, 0x88, 0xaa, 0x17, 0xeb, 0xa0, 0x70,
	0xc0, 0x69, 0xa0, 0xaf, 0x21, 0xa8, 0x74, 0xfc, 0xf0, 0xde, 0x6e, 0xfe, 0xb1, 0xa9, 0x6b, 0x66,
	0xe1, 0x0e, 0xb4, 0xa0, 0x9c, 0x98, 0xd9, 0x93, 0xd9, 0x78, 0x90, 0xea, 0xf6, 0xed, 0x76, 0xb5,
	0x06, 0x49, 0x4a, 0x8d, 0xeb, 0x54, 0x1d, 0x0a, 0x09, 0xca, 0x1d, 0xa6, 0x72, 0xbe, 0x87, 0x9c,
	0x22, 0xc9, 0xc7, 0xaa, 0xd6, 0x50, 0xbd, 0x82, 0xfd, 0xb5, 0x3c, 0xd1, 0x1d, 0x49, 0x9b, 0xc9,
	0xc7, 0x5f, 0x99, 0x7f, 0xac, 0xd3, 0xc6, 0x55, 0x23, 0x52, 0x12, 0x3d, 0xfa, 0x77, 0x72, 0x70,
	0xbd, 0x90, 0xc5, 0xf9, 0x42, 0x16, 0x7f, 0x2e, 0x64, 0xf1, 0xfd, 0x52, 0x16, 0xe6, 0x4b, 0x59,
	0xb8, 0x59, 0xca, 0xc2, 0xf9, 0x4e, 0x68, 0xfa, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0xe8, 0x28, 0x9a, 0xa4, 0x04, 0x00, 0x00,
}

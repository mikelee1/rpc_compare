// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	BenchmarkMessage
*/
package main

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

type BenchmarkMessage struct {
	Field1   string   `protobuf:"bytes,1,opt,name=field1" json:"field1,omitempty"`
	Field9   string   `protobuf:"bytes,9,opt,name=field9" json:"field9,omitempty"`
	Field18  string   `protobuf:"bytes,18,opt,name=field18" json:"field18,omitempty"`
	Field80  bool     `protobuf:"varint,80,opt,name=field80" json:"field80,omitempty"`
	Field81  bool     `protobuf:"varint,81,opt,name=field81" json:"field81,omitempty"`
	Field2   int32    `protobuf:"varint,2,opt,name=field2" json:"field2,omitempty"`
	Field3   int32    `protobuf:"varint,3,opt,name=field3" json:"field3,omitempty"`
	Field280 int32    `protobuf:"varint,280,opt,name=field280" json:"field280,omitempty"`
	Field6   int32    `protobuf:"varint,6,opt,name=field6" json:"field6,omitempty"`
	Field22  int64    `protobuf:"varint,22,opt,name=field22" json:"field22,omitempty"`
	Field4   string   `protobuf:"bytes,4,opt,name=field4" json:"field4,omitempty"`
	Field5   []uint64 `protobuf:"fixed64,5,rep,packed,name=field5" json:"field5,omitempty"`
	Field59  bool     `protobuf:"varint,59,opt,name=field59" json:"field59,omitempty"`
	Field7   string   `protobuf:"bytes,7,opt,name=field7" json:"field7,omitempty"`
	Field16  int32    `protobuf:"varint,16,opt,name=field16" json:"field16,omitempty"`
	Field130 int32    `protobuf:"varint,130,opt,name=field130" json:"field130,omitempty"`
	Field12  bool     `protobuf:"varint,12,opt,name=field12" json:"field12,omitempty"`
	Field17  bool     `protobuf:"varint,17,opt,name=field17" json:"field17,omitempty"`
	Field13  bool     `protobuf:"varint,13,opt,name=field13" json:"field13,omitempty"`
	Field14  bool     `protobuf:"varint,14,opt,name=field14" json:"field14,omitempty"`
	Field104 int32    `protobuf:"varint,104,opt,name=field104" json:"field104,omitempty"`
	Field100 int32    `protobuf:"varint,100,opt,name=field100" json:"field100,omitempty"`
	Field101 int32    `protobuf:"varint,101,opt,name=field101" json:"field101,omitempty"`
	Field102 string   `protobuf:"bytes,102,opt,name=field102" json:"field102,omitempty"`
	Field103 string   `protobuf:"bytes,103,opt,name=field103" json:"field103,omitempty"`
	Field29  int32    `protobuf:"varint,29,opt,name=field29" json:"field29,omitempty"`
	Field30  bool     `protobuf:"varint,30,opt,name=field30" json:"field30,omitempty"`
	Field60  int32    `protobuf:"varint,60,opt,name=field60" json:"field60,omitempty"`
	Field271 int32    `protobuf:"varint,271,opt,name=field271" json:"field271,omitempty"`
	Field272 int32    `protobuf:"varint,272,opt,name=field272" json:"field272,omitempty"`
	Field150 int32    `protobuf:"varint,150,opt,name=field150" json:"field150,omitempty"`
	Field23  int32    `protobuf:"varint,23,opt,name=field23" json:"field23,omitempty"`
	Field24  bool     `protobuf:"varint,24,opt,name=field24" json:"field24,omitempty"`
	Field25  int32    `protobuf:"varint,25,opt,name=field25" json:"field25,omitempty"`
	Field78  bool     `protobuf:"varint,78,opt,name=field78" json:"field78,omitempty"`
	Field67  int32    `protobuf:"varint,67,opt,name=field67" json:"field67,omitempty"`
	Field68  int32    `protobuf:"varint,68,opt,name=field68" json:"field68,omitempty"`
	Field128 int32    `protobuf:"varint,128,opt,name=field128" json:"field128,omitempty"`
	Field129 string   `protobuf:"bytes,129,opt,name=field129" json:"field129,omitempty"`
	Field131 int32    `protobuf:"varint,131,opt,name=field131" json:"field131,omitempty"`
}

func (m *BenchmarkMessage) Reset()                    { *m = BenchmarkMessage{} }
func (m *BenchmarkMessage) String() string            { return proto.CompactTextString(m) }
func (*BenchmarkMessage) ProtoMessage()               {}
func (*BenchmarkMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BenchmarkMessage) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func (m *BenchmarkMessage) GetField9() string {
	if m != nil {
		return m.Field9
	}
	return ""
}

func (m *BenchmarkMessage) GetField18() string {
	if m != nil {
		return m.Field18
	}
	return ""
}

func (m *BenchmarkMessage) GetField80() bool {
	if m != nil {
		return m.Field80
	}
	return false
}

func (m *BenchmarkMessage) GetField81() bool {
	if m != nil {
		return m.Field81
	}
	return false
}

func (m *BenchmarkMessage) GetField2() int32 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *BenchmarkMessage) GetField3() int32 {
	if m != nil {
		return m.Field3
	}
	return 0
}

func (m *BenchmarkMessage) GetField280() int32 {
	if m != nil {
		return m.Field280
	}
	return 0
}

func (m *BenchmarkMessage) GetField6() int32 {
	if m != nil {
		return m.Field6
	}
	return 0
}

func (m *BenchmarkMessage) GetField22() int64 {
	if m != nil {
		return m.Field22
	}
	return 0
}

func (m *BenchmarkMessage) GetField4() string {
	if m != nil {
		return m.Field4
	}
	return ""
}

func (m *BenchmarkMessage) GetField5() []uint64 {
	if m != nil {
		return m.Field5
	}
	return nil
}

func (m *BenchmarkMessage) GetField59() bool {
	if m != nil {
		return m.Field59
	}
	return false
}

func (m *BenchmarkMessage) GetField7() string {
	if m != nil {
		return m.Field7
	}
	return ""
}

func (m *BenchmarkMessage) GetField16() int32 {
	if m != nil {
		return m.Field16
	}
	return 0
}

func (m *BenchmarkMessage) GetField130() int32 {
	if m != nil {
		return m.Field130
	}
	return 0
}

func (m *BenchmarkMessage) GetField12() bool {
	if m != nil {
		return m.Field12
	}
	return false
}

func (m *BenchmarkMessage) GetField17() bool {
	if m != nil {
		return m.Field17
	}
	return false
}

func (m *BenchmarkMessage) GetField13() bool {
	if m != nil {
		return m.Field13
	}
	return false
}

func (m *BenchmarkMessage) GetField14() bool {
	if m != nil {
		return m.Field14
	}
	return false
}

func (m *BenchmarkMessage) GetField104() int32 {
	if m != nil {
		return m.Field104
	}
	return 0
}

func (m *BenchmarkMessage) GetField100() int32 {
	if m != nil {
		return m.Field100
	}
	return 0
}

func (m *BenchmarkMessage) GetField101() int32 {
	if m != nil {
		return m.Field101
	}
	return 0
}

func (m *BenchmarkMessage) GetField102() string {
	if m != nil {
		return m.Field102
	}
	return ""
}

func (m *BenchmarkMessage) GetField103() string {
	if m != nil {
		return m.Field103
	}
	return ""
}

func (m *BenchmarkMessage) GetField29() int32 {
	if m != nil {
		return m.Field29
	}
	return 0
}

func (m *BenchmarkMessage) GetField30() bool {
	if m != nil {
		return m.Field30
	}
	return false
}

func (m *BenchmarkMessage) GetField60() int32 {
	if m != nil {
		return m.Field60
	}
	return 0
}

func (m *BenchmarkMessage) GetField271() int32 {
	if m != nil {
		return m.Field271
	}
	return 0
}

func (m *BenchmarkMessage) GetField272() int32 {
	if m != nil {
		return m.Field272
	}
	return 0
}

func (m *BenchmarkMessage) GetField150() int32 {
	if m != nil {
		return m.Field150
	}
	return 0
}

func (m *BenchmarkMessage) GetField23() int32 {
	if m != nil {
		return m.Field23
	}
	return 0
}

func (m *BenchmarkMessage) GetField24() bool {
	if m != nil {
		return m.Field24
	}
	return false
}

func (m *BenchmarkMessage) GetField25() int32 {
	if m != nil {
		return m.Field25
	}
	return 0
}

func (m *BenchmarkMessage) GetField78() bool {
	if m != nil {
		return m.Field78
	}
	return false
}

func (m *BenchmarkMessage) GetField67() int32 {
	if m != nil {
		return m.Field67
	}
	return 0
}

func (m *BenchmarkMessage) GetField68() int32 {
	if m != nil {
		return m.Field68
	}
	return 0
}

func (m *BenchmarkMessage) GetField128() int32 {
	if m != nil {
		return m.Field128
	}
	return 0
}

func (m *BenchmarkMessage) GetField129() string {
	if m != nil {
		return m.Field129
	}
	return ""
}

func (m *BenchmarkMessage) GetField131() int32 {
	if m != nil {
		return m.Field131
	}
	return 0
}

func init() {
	proto.RegisterType((*BenchmarkMessage)(nil), "main.BenchmarkMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	// Sends a greeting
	Say(ctx context.Context, in *BenchmarkMessage, opts ...grpc.CallOption) (*BenchmarkMessage, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Say(ctx context.Context, in *BenchmarkMessage, opts ...grpc.CallOption) (*BenchmarkMessage, error) {
	out := new(BenchmarkMessage)
	err := grpc.Invoke(ctx, "/main.Hello/Say", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	// Sends a greeting
	Say(context.Context, *BenchmarkMessage) (*BenchmarkMessage, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BenchmarkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Hello/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Say(ctx, req.(*BenchmarkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Hello_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x31, 0x69, 0xd2, 0xd4, 0x02, 0x54, 0x7c, 0x08, 0x1f, 0x45, 0x20, 0xab, 0x27, 0x9f,
	0xa2, 0xdd, 0x59, 0xdb, 0x6b, 0x0b, 0x0e, 0x08, 0x38, 0x70, 0x01, 0x41, 0x78, 0x82, 0x94, 0xba,
	0x7f, 0x44, 0xda, 0x54, 0x0d, 0x17, 0x6e, 0xfc, 0x79, 0x00, 0x38, 0x21, 0x1e, 0x84, 0x07, 0x44,
	0x24, 0xf1, 0xce, 0xb7, 0x48, 0x1c, 0x67, 0x7e, 0xfe, 0x76, 0xbe, 0x99, 0x1d, 0x6f, 0x3a, 0xbe,
	0x3a, 0x9a, 0x5e, 0x5d, 0x2f, 0x3f, 0x2e, 0xb3, 0x9d, 0x8b, 0xf9, 0xf9, 0xe5, 0xe1, 0xef, 0x71,
	0xba, 0xff, 0xac, 0xbb, 0x7c, 0x7f, 0x76, 0x31, 0xbf, 0xfe, 0xf0, 0xaa, 0x5b, 0xad, 0xe6, 0xa7,
	0x5d, 0x36, 0x49, 0x47, 0x27, 0xe7, 0xdd, 0xe2, 0xd8, 0x22, 0xc9, 0x93, 0x62, 0x6f, 0xb6, 0x8d,
	0x42, 0xbe, 0xc5, 0x1e, 0xe5, 0xdb, 0x0c, 0xe9, 0xee, 0xe6, 0x8b, 0x06, 0xd9, 0x1a, 0xf4, 0x61,
	0x20, 0x8d, 0xc1, 0x9b, 0x3c, 0x29, 0xc6, 0xb3, 0x3e, 0x54, 0x62, 0xf1, 0x96, 0x89, 0x56, 0x11,
	0xdc, 0xcc, 0x93, 0x62, 0xb8, 0xad, 0x22, 0x21, 0xef, 0x30, 0xa0, 0xbc, 0xcb, 0x1e, 0xa4, 0xe3,
	0xcd, 0x17, 0x8d, 0xc1, 0xaf, 0x8d, 0x24, 0x24, 0x82, 0xa8, 0xc6, 0x88, 0x44, 0x75, 0x28, 0x2f,
	0x82, 0x49, 0x9e, 0x14, 0x83, 0x59, 0x1f, 0x06, 0x45, 0x89, 0x1d, 0x6a, 0xb2, 0x0c, 0xf9, 0x0a,
	0xc3, 0x7c, 0x50, 0x8c, 0xb6, 0xf9, 0x2a, 0x9c, 0x54, 0xb5, 0x78, 0x4c, 0x8d, 0x54, 0x6d, 0x50,
	0x78, 0xec, 0xd2, 0x49, 0x5e, 0xc7, 0x55, 0x63, 0x7f, 0x6d, 0xaa, 0x0f, 0x43, 0x2b, 0xd6, 0x19,
	0x7c, 0x4d, 0xa8, 0x15, 0xeb, 0x74, 0x62, 0x56, 0x70, 0x8b, 0x0a, 0x59, 0x51, 0xe2, 0x71, 0x97,
	0x09, 0x95, 0x72, 0xb8, 0xcd, 0xc4, 0x29, 0x29, 0x71, 0x87, 0x49, 0x99, 0x1d, 0xf4, 0x26, 0x4c,
	0x89, 0x33, 0xf6, 0x60, 0x98, 0x19, 0x1c, 0x47, 0xcc, 0x10, 0xb3, 0xe8, 0x22, 0x66, 0x89, 0x09,
	0x4e, 0xd6, 0xc3, 0x08, 0x31, 0x31, 0x87, 0xd3, 0x88, 0xa9, 0x4b, 0x69, 0xf1, 0x90, 0x46, 0x25,
	0xba, 0x73, 0xce, 0xe0, 0x11, 0xf9, 0xa7, 0x39, 0xd5, 0x06, 0x4f, 0x48, 0x53, 0x1b, 0xdd, 0x14,
	0x6f, 0xf1, 0x3d, 0xda, 0x14, 0x6f, 0x09, 0x0a, 0x7e, 0xc4, 0x50, 0xf4, 0x62, 0x2a, 0x83, 0x9f,
	0xd1, 0xc5, 0x54, 0x5a, 0x50, 0x1c, 0xee, 0xb1, 0x49, 0xb2, 0x5f, 0x02, 0x64, 0x52, 0x4a, 0x25,
	0x15, 0xee, 0xb3, 0x46, 0xf7, 0xc9, 0x37, 0x78, 0x4d, 0x1a, 0xaf, 0x3f, 0x53, 0xed, 0xf1, 0x9c,
	0x1b, 0xd3, 0x6b, 0xae, 0x1b, 0xbc, 0x60, 0xd2, 0xa8, 0x71, 0x69, 0xf0, 0x39, 0x32, 0x2e, 0x0c,
	0x5b, 0x7c, 0x49, 0x78, 0xf4, 0xd2, 0xd2, 0x2e, 0x5a, 0x7c, 0x8b, 0x77, 0xd1, 0xca, 0xd3, 0x74,
	0xf8, 0xb2, 0x5b, 0x2c, 0x96, 0x99, 0x4f, 0x07, 0xef, 0xe6, 0x9f, 0xb2, 0xc9, 0xf4, 0xef, 0x6b,
	0x32, 0xfd, 0xf7, 0x25, 0x39, 0xf8, 0x4f, 0xfe, 0xf0, 0xc6, 0xd1, 0x68, 0xfd, 0x0a, 0xb9, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x58, 0x9f, 0xeb, 0x0c, 0x91, 0x04, 0x00, 0x00,
}

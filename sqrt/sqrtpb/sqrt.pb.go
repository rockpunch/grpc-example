// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sqrt.proto

package sqrtpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SqrtRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtRequest) Reset()         { *m = SqrtRequest{} }
func (m *SqrtRequest) String() string { return proto.CompactTextString(m) }
func (*SqrtRequest) ProtoMessage()    {}
func (*SqrtRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_746f45ae1873c6cc, []int{0}
}

func (m *SqrtRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtRequest.Unmarshal(m, b)
}
func (m *SqrtRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtRequest.Marshal(b, m, deterministic)
}
func (m *SqrtRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtRequest.Merge(m, src)
}
func (m *SqrtRequest) XXX_Size() int {
	return xxx_messageInfo_SqrtRequest.Size(m)
}
func (m *SqrtRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtRequest proto.InternalMessageInfo

func (m *SqrtRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SqrtResponse struct {
	Sqrt                 int32    `protobuf:"varint,1,opt,name=sqrt,proto3" json:"sqrt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtResponse) Reset()         { *m = SqrtResponse{} }
func (m *SqrtResponse) String() string { return proto.CompactTextString(m) }
func (*SqrtResponse) ProtoMessage()    {}
func (*SqrtResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_746f45ae1873c6cc, []int{1}
}

func (m *SqrtResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtResponse.Unmarshal(m, b)
}
func (m *SqrtResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtResponse.Marshal(b, m, deterministic)
}
func (m *SqrtResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtResponse.Merge(m, src)
}
func (m *SqrtResponse) XXX_Size() int {
	return xxx_messageInfo_SqrtResponse.Size(m)
}
func (m *SqrtResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtResponse proto.InternalMessageInfo

func (m *SqrtResponse) GetSqrt() int32 {
	if m != nil {
		return m.Sqrt
	}
	return 0
}

func init() {
	proto.RegisterType((*SqrtRequest)(nil), "sqrt.SqrtRequest")
	proto.RegisterType((*SqrtResponse)(nil), "sqrt.SqrtResponse")
}

func init() {
	proto.RegisterFile("sqrt.proto", fileDescriptor_746f45ae1873c6cc)
}

var fileDescriptor_746f45ae1873c6cc = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2e, 0x2c, 0x2a,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x54, 0xb9, 0xb8, 0x83, 0x0b,
	0x8b, 0x4a, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0xf2, 0x4a, 0x73,
	0x93, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x25, 0x25, 0x2e, 0x1e,
	0x88, 0xb2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x21, 0x2e, 0xb0, 0x76, 0xa8, 0x2a, 0x30,
	0xdb, 0xc8, 0x86, 0x8b, 0x33, 0x39, 0x31, 0x27, 0x39, 0x1e, 0xc4, 0x11, 0xd2, 0xe7, 0x62, 0x71,
	0x4e, 0xcc, 0x49, 0x16, 0x12, 0xd4, 0x03, 0x5b, 0x89, 0x64, 0x87, 0x94, 0x10, 0xb2, 0x10, 0xc4,
	0x3c, 0x25, 0x06, 0x27, 0x8e, 0x28, 0x36, 0x90, 0x70, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x7d, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x46, 0xd4, 0xbc, 0xad, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcSqrtClient is the client API for CalcSqrt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcSqrtClient interface {
	Calc(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type calcSqrtClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcSqrtClient(cc grpc.ClientConnInterface) CalcSqrtClient {
	return &calcSqrtClient{cc}
}

func (c *calcSqrtClient) Calc(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/sqrt.calc_sqrt/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcSqrtServer is the server API for CalcSqrt service.
type CalcSqrtServer interface {
	Calc(context.Context, *SqrtRequest) (*SqrtResponse, error)
}

// UnimplementedCalcSqrtServer can be embedded to have forward compatible implementations.
type UnimplementedCalcSqrtServer struct {
}

func (*UnimplementedCalcSqrtServer) Calc(ctx context.Context, req *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calc not implemented")
}

func RegisterCalcSqrtServer(s *grpc.Server, srv CalcSqrtServer) {
	s.RegisterService(&_CalcSqrt_serviceDesc, srv)
}

func _CalcSqrt_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcSqrtServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sqrt.calc_sqrt/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcSqrtServer).Calc(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalcSqrt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sqrt.calc_sqrt",
	HandlerType: (*CalcSqrtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _CalcSqrt_Calc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sqrt.proto",
}
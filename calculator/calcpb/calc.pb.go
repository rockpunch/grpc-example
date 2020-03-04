// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

package calc

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

type PndcRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PndcRequest) Reset()         { *m = PndcRequest{} }
func (m *PndcRequest) String() string { return proto.CompactTextString(m) }
func (*PndcRequest) ProtoMessage()    {}
func (*PndcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{0}
}

func (m *PndcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PndcRequest.Unmarshal(m, b)
}
func (m *PndcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PndcRequest.Marshal(b, m, deterministic)
}
func (m *PndcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PndcRequest.Merge(m, src)
}
func (m *PndcRequest) XXX_Size() int {
	return xxx_messageInfo_PndcRequest.Size(m)
}
func (m *PndcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PndcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PndcRequest proto.InternalMessageInfo

func (m *PndcRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PndcResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PndcResponse) Reset()         { *m = PndcResponse{} }
func (m *PndcResponse) String() string { return proto.CompactTextString(m) }
func (*PndcResponse) ProtoMessage()    {}
func (*PndcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2b9900dc883ea68, []int{1}
}

func (m *PndcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PndcResponse.Unmarshal(m, b)
}
func (m *PndcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PndcResponse.Marshal(b, m, deterministic)
}
func (m *PndcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PndcResponse.Merge(m, src)
}
func (m *PndcResponse) XXX_Size() int {
	return xxx_messageInfo_PndcResponse.Size(m)
}
func (m *PndcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PndcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PndcResponse proto.InternalMessageInfo

func (m *PndcResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*PndcRequest)(nil), "PndcRequest")
	proto.RegisterType((*PndcResponse)(nil), "PndcResponse")
}

func init() {
	proto.RegisterFile("calc.proto", fileDescriptor_a2b9900dc883ea68)
}

var fileDescriptor_a2b9900dc883ea68 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe5, 0xe2, 0x0e, 0xc8, 0x4b, 0x49, 0x0e, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0xcb, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x92,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2, 0x94, 0xd4, 0xb8, 0x78, 0x20, 0xca, 0x8a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x41, 0xea, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x60, 0xea, 0x20, 0x3c,
	0x23, 0x33, 0x88, 0x71, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xea, 0x5c, 0x2c, 0x20,
	0xae, 0x10, 0x8f, 0x1e, 0x92, 0x25, 0x52, 0xbc, 0x7a, 0xc8, 0x66, 0x29, 0x31, 0x18, 0x30, 0x26,
	0xb1, 0x81, 0x5d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x23, 0xd8, 0x65, 0xb7, 0x9b, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PndcServiceClient is the client API for PndcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PndcServiceClient interface {
	Pndc(ctx context.Context, in *PndcRequest, opts ...grpc.CallOption) (PndcService_PndcClient, error)
}

type pndcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPndcServiceClient(cc grpc.ClientConnInterface) PndcServiceClient {
	return &pndcServiceClient{cc}
}

func (c *pndcServiceClient) Pndc(ctx context.Context, in *PndcRequest, opts ...grpc.CallOption) (PndcService_PndcClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PndcService_serviceDesc.Streams[0], "/PndcService/Pndc", opts...)
	if err != nil {
		return nil, err
	}
	x := &pndcServicePndcClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PndcService_PndcClient interface {
	Recv() (*PndcResponse, error)
	grpc.ClientStream
}

type pndcServicePndcClient struct {
	grpc.ClientStream
}

func (x *pndcServicePndcClient) Recv() (*PndcResponse, error) {
	m := new(PndcResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PndcServiceServer is the server API for PndcService service.
type PndcServiceServer interface {
	Pndc(*PndcRequest, PndcService_PndcServer) error
}

// UnimplementedPndcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPndcServiceServer struct {
}

func (*UnimplementedPndcServiceServer) Pndc(req *PndcRequest, srv PndcService_PndcServer) error {
	return status.Errorf(codes.Unimplemented, "method Pndc not implemented")
}

func RegisterPndcServiceServer(s *grpc.Server, srv PndcServiceServer) {
	s.RegisterService(&_PndcService_serviceDesc, srv)
}

func _PndcService_Pndc_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PndcRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PndcServiceServer).Pndc(m, &pndcServicePndcServer{stream})
}

type PndcService_PndcServer interface {
	Send(*PndcResponse) error
	grpc.ServerStream
}

type pndcServicePndcServer struct {
	grpc.ServerStream
}

func (x *pndcServicePndcServer) Send(m *PndcResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PndcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PndcService",
	HandlerType: (*PndcServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pndc",
			Handler:       _PndcService_Pndc_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calc.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qtrago.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageRequest struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5bb0006e5595bda, []int{0}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type MessageReply struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReply) Reset()         { *m = MessageReply{} }
func (m *MessageReply) String() string { return proto.CompactTextString(m) }
func (*MessageReply) ProtoMessage()    {}
func (*MessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5bb0006e5595bda, []int{1}
}

func (m *MessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageReply.Unmarshal(m, b)
}
func (m *MessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageReply.Marshal(b, m, deterministic)
}
func (m *MessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReply.Merge(m, src)
}
func (m *MessageReply) XXX_Size() int {
	return xxx_messageInfo_MessageReply.Size(m)
}
func (m *MessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReply proto.InternalMessageInfo

func (m *MessageReply) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type TradeRequest struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeRequest) Reset()         { *m = TradeRequest{} }
func (m *TradeRequest) String() string { return proto.CompactTextString(m) }
func (*TradeRequest) ProtoMessage()    {}
func (*TradeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5bb0006e5595bda, []int{2}
}

func (m *TradeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeRequest.Unmarshal(m, b)
}
func (m *TradeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeRequest.Marshal(b, m, deterministic)
}
func (m *TradeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeRequest.Merge(m, src)
}
func (m *TradeRequest) XXX_Size() int {
	return xxx_messageInfo_TradeRequest.Size(m)
}
func (m *TradeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeRequest proto.InternalMessageInfo

func (m *TradeRequest) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type TradeReply struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeReply) Reset()         { *m = TradeReply{} }
func (m *TradeReply) String() string { return proto.CompactTextString(m) }
func (*TradeReply) ProtoMessage()    {}
func (*TradeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5bb0006e5595bda, []int{3}
}

func (m *TradeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeReply.Unmarshal(m, b)
}
func (m *TradeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeReply.Marshal(b, m, deterministic)
}
func (m *TradeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeReply.Merge(m, src)
}
func (m *TradeReply) XXX_Size() int {
	return xxx_messageInfo_TradeReply.Size(m)
}
func (m *TradeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeReply.DiscardUnknown(m)
}

var xxx_messageInfo_TradeReply proto.InternalMessageInfo

func (m *TradeReply) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "proto.MessageRequest")
	proto.RegisterType((*MessageReply)(nil), "proto.MessageReply")
	proto.RegisterType((*TradeRequest)(nil), "proto.TradeRequest")
	proto.RegisterType((*TradeReply)(nil), "proto.TradeReply")
}

func init() { proto.RegisterFile("qtrago.proto", fileDescriptor_d5bb0006e5595bda) }

var fileDescriptor_d5bb0006e5595bda = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2c, 0x29, 0x4a,
	0x4c, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x2a, 0x5c, 0x7c,
	0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x59, 0xc5, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x12, 0x17, 0x0f, 0x5c, 0x55, 0x41, 0x4e, 0x25, 0x2e, 0x35, 0x21, 0x45, 0x89, 0x29, 0x78, 0xcd,
	0x51, 0xe0, 0xe2, 0x82, 0xaa, 0xc1, 0x61, 0x8a, 0x51, 0x2b, 0x23, 0x97, 0x70, 0x60, 0x69, 0x62,
	0x5e, 0x49, 0x66, 0x49, 0x62, 0x49, 0x66, 0x59, 0x2a, 0x48, 0x79, 0x66, 0x5e, 0xba, 0x90, 0x35,
	0x17, 0x37, 0xd4, 0x05, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0xa2, 0x10, 0x5f, 0xe8, 0xa1, 0xba, 0x5d,
	0x4a, 0x18, 0x5d, 0xb8, 0x20, 0xa7, 0x52, 0x89, 0x41, 0xc8, 0x90, 0x8b, 0x15, 0x6c, 0xad, 0x10,
	0x4c, 0x1e, 0xd9, 0xa1, 0x52, 0x82, 0xa8, 0x82, 0x60, 0x2d, 0x4e, 0x66, 0x5c, 0x12, 0xc9, 0xf9,
	0xb9, 0x7a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xe9, 0xf9, 0x29, 0xc9, 0xf9, 0x79, 0xe9,
	0x10, 0x85, 0x4e, 0x12, 0x58, 0x1c, 0x18, 0x00, 0x92, 0x09, 0x60, 0x4c, 0x62, 0x03, 0x2b, 0x31,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x18, 0xbe, 0xd6, 0x6d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuantitativeTradingClient is the client API for QuantitativeTrading service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuantitativeTradingClient interface {
	MessageInfo(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
	Trade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*TradeReply, error)
}

type quantitativeTradingClient struct {
	cc *grpc.ClientConn
}

func NewQuantitativeTradingClient(cc *grpc.ClientConn) QuantitativeTradingClient {
	return &quantitativeTradingClient{cc}
}

func (c *quantitativeTradingClient) MessageInfo(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/proto.QuantitativeTrading/MessageInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quantitativeTradingClient) Trade(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*TradeReply, error) {
	out := new(TradeReply)
	err := c.cc.Invoke(ctx, "/proto.QuantitativeTrading/Trade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuantitativeTradingServer is the server API for QuantitativeTrading service.
type QuantitativeTradingServer interface {
	MessageInfo(context.Context, *MessageRequest) (*MessageReply, error)
	Trade(context.Context, *TradeRequest) (*TradeReply, error)
}

func RegisterQuantitativeTradingServer(s *grpc.Server, srv QuantitativeTradingServer) {
	s.RegisterService(&_QuantitativeTrading_serviceDesc, srv)
}

func _QuantitativeTrading_MessageInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantitativeTradingServer).MessageInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuantitativeTrading/MessageInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantitativeTradingServer).MessageInfo(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuantitativeTrading_Trade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantitativeTradingServer).Trade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuantitativeTrading/Trade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantitativeTradingServer).Trade(ctx, req.(*TradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuantitativeTrading_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.QuantitativeTrading",
	HandlerType: (*QuantitativeTradingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageInfo",
			Handler:    _QuantitativeTrading_MessageInfo_Handler,
		},
		{
			MethodName: "Trade",
			Handler:    _QuantitativeTrading_Trade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qtrago.proto",
}
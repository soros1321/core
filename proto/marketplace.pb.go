// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace.proto

package sonm

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

type GetOrdersRequest struct {
	// Order keeps slot and type for searching
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
	// Count describe how namy results must be returned (order by price)
	Count uint64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *GetOrdersRequest) Reset()                    { *m = GetOrdersRequest{} }
func (m *GetOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersRequest) ProtoMessage()               {}
func (*GetOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *GetOrdersRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *GetOrdersRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetOrdersReply struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *GetOrdersReply) Reset()                    { *m = GetOrdersReply{} }
func (m *GetOrdersReply) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersReply) ProtoMessage()               {}
func (*GetOrdersReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *GetOrdersReply) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

// GetProcessingReply represent Node's local
// orders processing status
type GetProcessingReply struct {
	Orders map[string]*GetProcessingReply_ProcessedOrder `protobuf:"bytes,2,rep,name=orders" json:"orders,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GetProcessingReply) Reset()                    { *m = GetProcessingReply{} }
func (m *GetProcessingReply) String() string            { return proto.CompactTextString(m) }
func (*GetProcessingReply) ProtoMessage()               {}
func (*GetProcessingReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *GetProcessingReply) GetOrders() map[string]*GetProcessingReply_ProcessedOrder {
	if m != nil {
		return m.Orders
	}
	return nil
}

type GetProcessingReply_ProcessedOrder struct {
	Id        string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Status    uint32     `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Timestamp *Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Extra     string     `protobuf:"bytes,4,opt,name=extra" json:"extra,omitempty"`
}

func (m *GetProcessingReply_ProcessedOrder) Reset()         { *m = GetProcessingReply_ProcessedOrder{} }
func (m *GetProcessingReply_ProcessedOrder) String() string { return proto.CompactTextString(m) }
func (*GetProcessingReply_ProcessedOrder) ProtoMessage()    {}
func (*GetProcessingReply_ProcessedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{2, 0}
}

func (m *GetProcessingReply_ProcessedOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetProcessingReply_ProcessedOrder) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetProcessingReply_ProcessedOrder) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *GetProcessingReply_ProcessedOrder) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOrdersRequest)(nil), "sonm.GetOrdersRequest")
	proto.RegisterType((*GetOrdersReply)(nil), "sonm.GetOrdersReply")
	proto.RegisterType((*GetProcessingReply)(nil), "sonm.GetProcessingReply")
	proto.RegisterType((*GetProcessingReply_ProcessedOrder)(nil), "sonm.GetProcessingReply.ProcessedOrder")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Market service

type MarketClient interface {
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error)
	GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error)
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	CancelOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Empty, error)
	// GetProcessing returns currently processing orders
	// not required in Marketplace service
	// must be implemented on Node
	GetProcessing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProcessingReply, error)
}

type marketClient struct {
	cc *grpc.ClientConn
}

func NewMarketClient(cc *grpc.ClientConn) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error) {
	out := new(GetOrdersReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrderByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CancelOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetProcessing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProcessingReply, error) {
	out := new(GetProcessingReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetProcessing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Market service

type MarketServer interface {
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersReply, error)
	GetOrderByID(context.Context, *ID) (*Order, error)
	CreateOrder(context.Context, *Order) (*Order, error)
	CancelOrder(context.Context, *Order) (*Empty, error)
	// GetProcessing returns currently processing orders
	// not required in Marketplace service
	// must be implemented on Node
	GetProcessing(context.Context, *Empty) (*GetProcessingReply, error)
}

func RegisterMarketServer(s *grpc.Server, srv MarketServer) {
	s.RegisterService(&_Market_serviceDesc, srv)
}

func _Market_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrderByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CancelOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetProcessing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetProcessing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetProcessing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetProcessing(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Market_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _Market_GetOrders_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Market_GetOrderByID_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Market_CreateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Market_CancelOrder_Handler,
		},
		{
			MethodName: "GetProcessing",
			Handler:    _Market_GetProcessing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace.proto",
}

func init() { proto.RegisterFile("marketplace.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5d, 0xaf, 0xd2, 0x40,
	0x10, 0xa5, 0x05, 0x1a, 0x3b, 0x95, 0x0f, 0x27, 0x84, 0x34, 0x7d, 0xc2, 0x6a, 0x02, 0x3e, 0xc8,
	0x03, 0xc6, 0x84, 0xf8, 0xf1, 0xa2, 0x10, 0x42, 0x8c, 0xd1, 0x34, 0xfe, 0x81, 0xa5, 0x9d, 0x98,
	0x86, 0x7e, 0xb9, 0xbb, 0x35, 0xf6, 0xc1, 0x9f, 0xe5, 0xdf, 0x33, 0x37, 0xdd, 0x2d, 0xdc, 0x72,
	0xb9, 0xbc, 0x75, 0xce, 0x9c, 0x73, 0xa6, 0x73, 0x66, 0xe1, 0x59, 0xca, 0xf8, 0x91, 0x64, 0x91,
	0xb0, 0x90, 0x96, 0x05, 0xcf, 0x65, 0x8e, 0x3d, 0x91, 0x67, 0xa9, 0x67, 0x1f, 0xe2, 0x48, 0x03,
	0xde, 0x28, 0xce, 0x6a, 0x28, 0x8b, 0x99, 0x06, 0xfc, 0x2f, 0x30, 0xde, 0x91, 0xfc, 0xc6, 0x23,
	0xe2, 0x22, 0xa0, 0x5f, 0x25, 0x09, 0x89, 0xcf, 0xa1, 0x9f, 0xd7, 0x80, 0x6b, 0xcc, 0x8c, 0x85,
	0xb3, 0x72, 0x96, 0xb5, 0x64, 0xa9, 0x38, 0x81, 0xee, 0xe0, 0x04, 0xfa, 0x61, 0x5e, 0x66, 0xd2,
	0x35, 0x67, 0xc6, 0xa2, 0x17, 0xe8, 0xc2, 0x7f, 0x0b, 0xc3, 0x96, 0x59, 0x91, 0x54, 0xf8, 0x02,
	0x2c, 0x25, 0x10, 0xae, 0x31, 0xeb, 0x3e, 0xf4, 0x6a, 0x5a, 0xfe, 0x3f, 0x13, 0x70, 0x47, 0xf2,
	0x3b, 0xcf, 0x43, 0x12, 0x22, 0xce, 0x7e, 0x6a, 0xed, 0x87, 0xb3, 0xd6, 0x54, 0xda, 0x97, 0x5a,
	0x7b, 0xcd, 0xd4, 0x76, 0x62, 0x9b, 0x49, 0x5e, 0x9d, 0x4c, 0xbd, 0xbf, 0x30, 0x6c, 0x68, 0x14,
	0xa9, 0x3e, 0x0e, 0xc1, 0x8c, 0x23, 0xb5, 0x93, 0x1d, 0x98, 0x71, 0x84, 0x53, 0xb0, 0x84, 0x64,
	0xb2, 0x14, 0x6a, 0x89, 0x41, 0xd0, 0x54, 0xf8, 0x1a, 0x6c, 0x19, 0xa7, 0x24, 0x24, 0x4b, 0x0b,
	0xb7, 0xab, 0x22, 0x18, 0xe9, 0xd1, 0x3f, 0x4e, 0x70, 0x70, 0xcf, 0xa8, 0xa3, 0xa0, 0x3f, 0x92,
	0x33, 0xb7, 0xa7, 0x9c, 0x75, 0xe1, 0x1d, 0xc0, 0x69, 0xfd, 0x15, 0x8e, 0xa1, 0x7b, 0xa4, 0xaa,
	0x19, 0x5e, 0x7f, 0xe2, 0x47, 0xe8, 0xff, 0x66, 0x49, 0x49, 0x6a, 0xb8, 0xb3, 0x9a, 0xdf, 0x5c,
	0xee, 0x72, 0x8b, 0x40, 0xab, 0xde, 0x99, 0x6b, 0x63, 0xf5, 0xdf, 0x00, 0xeb, 0xab, 0xba, 0x39,
	0xbe, 0x07, 0xfb, 0x9c, 0x3c, 0x4e, 0xcf, 0x5e, 0x17, 0x77, 0xf5, 0x26, 0x57, 0x78, 0x91, 0x54,
	0x7e, 0x07, 0xe7, 0xf0, 0xf4, 0x84, 0x7d, 0xaa, 0xf6, 0x1b, 0x7c, 0xa2, 0x79, 0xfb, 0x8d, 0xd7,
	0x3e, 0x97, 0xdf, 0xc1, 0x57, 0xe0, 0x7c, 0xe6, 0xc4, 0x24, 0xe9, 0x40, 0xdb, 0xdd, 0xc7, 0xa8,
	0x2c, 0x0b, 0x29, 0xb9, 0x4d, 0xdd, 0xa6, 0x85, 0xac, 0xc7, 0xaf, 0x61, 0x70, 0xb1, 0x36, 0xb6,
	0xfb, 0x9e, 0x7b, 0x2b, 0x18, 0xbf, 0x73, 0xb0, 0xd4, 0x1b, 0x7e, 0x73, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x38, 0x2d, 0xda, 0xe1, 0xfa, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: miner.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// grpccmd imports
import (
	"github.com/nathanielc/grpccmd"
	"github.com/spf13/cobra"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NATType int32

const (
	NATType_NONE                   NATType = 0
	NATType_BLOCKED                NATType = 1
	NATType_FULL                   NATType = 2
	NATType_SYMMETRIC              NATType = 3
	NATType_RESTRICTED             NATType = 4
	NATType_PORT_RESTRICTED        NATType = 5
	NATType_SYMMETRIC_UDP_FIREWALL NATType = 6
	NATType_UNKNOWN                NATType = 7
)

var NATType_name = map[int32]string{
	0: "NONE",
	1: "BLOCKED",
	2: "FULL",
	3: "SYMMETRIC",
	4: "RESTRICTED",
	5: "PORT_RESTRICTED",
	6: "SYMMETRIC_UDP_FIREWALL",
	7: "UNKNOWN",
}
var NATType_value = map[string]int32{
	"NONE":                   0,
	"BLOCKED":                1,
	"FULL":                   2,
	"SYMMETRIC":              3,
	"RESTRICTED":             4,
	"PORT_RESTRICTED":        5,
	"SYMMETRIC_UDP_FIREWALL": 6,
	"UNKNOWN":                7,
}

func (x NATType) String() string {
	return proto.EnumName(NATType_name, int32(x))
}
func (NATType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type MinerHandshakeRequest struct {
	Hub   string      `protobuf:"bytes,1,opt,name=hub" json:"hub,omitempty"`
	Tasks []*TaskInfo `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *MinerHandshakeRequest) Reset()                    { *m = MinerHandshakeRequest{} }
func (m *MinerHandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeRequest) ProtoMessage()               {}
func (*MinerHandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *MinerHandshakeRequest) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *MinerHandshakeRequest) GetTasks() []*TaskInfo {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type MinerHandshakeReply struct {
	Miner        string        `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
	Capabilities *Capabilities `protobuf:"bytes,2,opt,name=capabilities" json:"capabilities,omitempty"`
	NatType      NATType       `protobuf:"varint,3,opt,name=natType,enum=sonm.NATType" json:"natType,omitempty"`
}

func (m *MinerHandshakeReply) Reset()                    { *m = MinerHandshakeReply{} }
func (m *MinerHandshakeReply) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeReply) ProtoMessage()               {}
func (*MinerHandshakeReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *MinerHandshakeReply) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *MinerHandshakeReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func (m *MinerHandshakeReply) GetNatType() NATType {
	if m != nil {
		return m.NatType
	}
	return NATType_NONE
}

type MinerStartRequest struct {
	Id            string                    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Registry      string                    `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string                    `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string                    `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	RestartPolicy *ContainerRestartPolicy   `protobuf:"bytes,5,opt,name=restartPolicy" json:"restartPolicy,omitempty"`
	Resources     *TaskResourceRequirements `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
	PublicKeyData string                    `protobuf:"bytes,7,opt,name=PublicKeyData" json:"PublicKeyData,omitempty"`
	CommitOnStop  bool                      `protobuf:"varint,8,opt,name=commitOnStop" json:"commitOnStop,omitempty"`
	Env           map[string]string         `protobuf:"bytes,9,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// OrderId describes an unique order identifier.
	// It is here for proper resource allocation and limitation.
	OrderId string `protobuf:"bytes,10,opt,name=orderId" json:"orderId,omitempty"`
}

func (m *MinerStartRequest) Reset()                    { *m = MinerStartRequest{} }
func (m *MinerStartRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStartRequest) ProtoMessage()               {}
func (*MinerStartRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *MinerStartRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MinerStartRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *MinerStartRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *MinerStartRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *MinerStartRequest) GetRestartPolicy() *ContainerRestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return nil
}

func (m *MinerStartRequest) GetResources() *TaskResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *MinerStartRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

func (m *MinerStartRequest) GetCommitOnStop() bool {
	if m != nil {
		return m.CommitOnStop
	}
	return false
}

func (m *MinerStartRequest) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *MinerStartRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type SocketAddr struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	//
	// Actually an `uint16` here. Protobuf is so clear and handy.
	Port uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *SocketAddr) Reset()                    { *m = SocketAddr{} }
func (m *SocketAddr) String() string            { return proto.CompactTextString(m) }
func (*SocketAddr) ProtoMessage()               {}
func (*SocketAddr) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *SocketAddr) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SocketAddr) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type MinerStartReply struct {
	Container string   `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
	Routes    []*Route `protobuf:"bytes,2,rep,name=routes" json:"routes,omitempty"`
}

func (m *MinerStartReply) Reset()                    { *m = MinerStartReply{} }
func (m *MinerStartReply) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReply) ProtoMessage()               {}
func (*MinerStartReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *MinerStartReply) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *MinerStartReply) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type TaskInfo struct {
	Request *MinerStartRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Reply   *MinerStartReply   `protobuf:"bytes,2,opt,name=reply" json:"reply,omitempty"`
}

func (m *TaskInfo) Reset()                    { *m = TaskInfo{} }
func (m *TaskInfo) String() string            { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()               {}
func (*TaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *TaskInfo) GetRequest() *MinerStartRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *TaskInfo) GetReply() *MinerStartReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

type Route struct {
	Port     string      `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
	Endpoint *SocketAddr `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *Route) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Route) GetEndpoint() *SocketAddr {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type MinerStatusMapRequest struct {
}

func (m *MinerStatusMapRequest) Reset()                    { *m = MinerStatusMapRequest{} }
func (m *MinerStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStatusMapRequest) ProtoMessage()               {}
func (*MinerStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func init() {
	proto.RegisterType((*MinerHandshakeRequest)(nil), "sonm.MinerHandshakeRequest")
	proto.RegisterType((*MinerHandshakeReply)(nil), "sonm.MinerHandshakeReply")
	proto.RegisterType((*MinerStartRequest)(nil), "sonm.MinerStartRequest")
	proto.RegisterType((*SocketAddr)(nil), "sonm.SocketAddr")
	proto.RegisterType((*MinerStartReply)(nil), "sonm.MinerStartReply")
	proto.RegisterType((*TaskInfo)(nil), "sonm.TaskInfo")
	proto.RegisterType((*Route)(nil), "sonm.Route")
	proto.RegisterType((*MinerStatusMapRequest)(nil), "sonm.MinerStatusMapRequest")
	proto.RegisterEnum("sonm.NATType", NATType_name, NATType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Miner service

type MinerClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingReply, error)
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoReply, error)
	Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error)
	Load(ctx context.Context, opts ...grpc.CallOption) (Miner_LoadClient, error)
	Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error)
	Stop(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error)
	TaskDetails(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskStatusReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error)
	DiscoverHub(ctx context.Context, in *DiscoverHubRequest, opts ...grpc.CallOption) (*Empty, error)
}

type minerClient struct {
	cc *grpc.ClientConn
}

func NewMinerClient(cc *grpc.ClientConn) MinerClient {
	return &minerClient{cc}
}

func (c *minerClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error) {
	out := new(MinerHandshakeReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Handshake", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Load(ctx context.Context, opts ...grpc.CallOption) (Miner_LoadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[0], c.cc, "/sonm.Miner/Load", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerLoadClient{stream}
	return x, nil
}

type Miner_LoadClient interface {
	Send(*Chunk) error
	Recv() (*Progress, error)
	grpc.ClientStream
}

type minerLoadClient struct {
	grpc.ClientStream
}

func (x *minerLoadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *minerLoadClient) Recv() (*Progress, error) {
	m := new(Progress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error) {
	out := new(MinerStartReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Stop(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Miner/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[1], c.cc, "/sonm.Miner/TasksStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTasksStatusClient{stream}
	return x, nil
}

type Miner_TasksStatusClient interface {
	Send(*MinerStatusMapRequest) error
	Recv() (*StatusMapReply, error)
	grpc.ClientStream
}

type minerTasksStatusClient struct {
	grpc.ClientStream
}

func (x *minerTasksStatusClient) Send(m *MinerStatusMapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *minerTasksStatusClient) Recv() (*StatusMapReply, error) {
	m := new(StatusMapReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) TaskDetails(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/TaskDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[2], c.cc, "/sonm.Miner/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Miner_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type minerTaskLogsClient struct {
	grpc.ClientStream
}

func (x *minerTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) DiscoverHub(ctx context.Context, in *DiscoverHubRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Miner/DiscoverHub", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Miner service

type MinerServer interface {
	Ping(context.Context, *Empty) (*PingReply, error)
	Info(context.Context, *Empty) (*InfoReply, error)
	Handshake(context.Context, *MinerHandshakeRequest) (*MinerHandshakeReply, error)
	Load(Miner_LoadServer) error
	Start(context.Context, *MinerStartRequest) (*MinerStartReply, error)
	Stop(context.Context, *ID) (*Empty, error)
	TasksStatus(Miner_TasksStatusServer) error
	TaskDetails(context.Context, *ID) (*TaskStatusReply, error)
	TaskLogs(*TaskLogsRequest, Miner_TaskLogsServer) error
	DiscoverHub(context.Context, *DiscoverHubRequest) (*Empty, error)
}

func RegisterMinerServer(s *grpc.Server, srv MinerServer) {
	s.RegisterService(&_Miner_serviceDesc, srv)
}

func _Miner_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerHandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Handshake(ctx, req.(*MinerHandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Load_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MinerServer).Load(&minerLoadServer{stream})
}

type Miner_LoadServer interface {
	Send(*Progress) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type minerLoadServer struct {
	grpc.ServerStream
}

func (x *minerLoadServer) Send(m *Progress) error {
	return x.ServerStream.SendMsg(m)
}

func (x *minerLoadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Miner_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Start(ctx, req.(*MinerStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Stop(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TasksStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MinerServer).TasksStatus(&minerTasksStatusServer{stream})
}

type Miner_TasksStatusServer interface {
	Send(*StatusMapReply) error
	Recv() (*MinerStatusMapRequest, error)
	grpc.ServerStream
}

type minerTasksStatusServer struct {
	grpc.ServerStream
}

func (x *minerTasksStatusServer) Send(m *StatusMapReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *minerTasksStatusServer) Recv() (*MinerStatusMapRequest, error) {
	m := new(MinerStatusMapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Miner_TaskDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).TaskDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/TaskDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).TaskDetails(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MinerServer).TaskLogs(m, &minerTaskLogsServer{stream})
}

type Miner_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type minerTaskLogsServer struct {
	grpc.ServerStream
}

func (x *minerTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Miner_DiscoverHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).DiscoverHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/DiscoverHub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).DiscoverHub(ctx, req.(*DiscoverHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Miner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Miner",
	HandlerType: (*MinerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Miner_Ping_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Miner_Info_Handler,
		},
		{
			MethodName: "Handshake",
			Handler:    _Miner_Handshake_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Miner_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Miner_Stop_Handler,
		},
		{
			MethodName: "TaskDetails",
			Handler:    _Miner_TaskDetails_Handler,
		},
		{
			MethodName: "DiscoverHub",
			Handler:    _Miner_DiscoverHub_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Load",
			Handler:       _Miner_Load_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TasksStatus",
			Handler:       _Miner_TasksStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TaskLogs",
			Handler:       _Miner_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "miner.proto",
}

// Begin grpccmd
var _ = grpccmd.RunE

// Miner
var _MinerCmd = &cobra.Command{
	Use:   "miner [method]",
	Short: "Subcommand for the Miner service.",
}

var _Miner_PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Make the Ping method call, input-type: sonm.Empty output-type: sonm.PingReply",
	RunE: grpccmd.RunE(
		"Ping",
		"sonm.Empty",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Make the Info method call, input-type: sonm.Empty output-type: sonm.InfoReply",
	RunE: grpccmd.RunE(
		"Info",
		"sonm.Empty",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_HandshakeCmd = &cobra.Command{
	Use:   "handshake",
	Short: "Make the Handshake method call, input-type: sonm.MinerHandshakeRequest output-type: sonm.MinerHandshakeReply",
	RunE: grpccmd.RunE(
		"Handshake",
		"sonm.MinerHandshakeRequest",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_LoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Make the Load method call, input-type: sonm.Chunk output-type: sonm.Progress",
	RunE: grpccmd.RunE(
		"Load",
		"sonm.Chunk",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Make the Start method call, input-type: sonm.MinerStartRequest output-type: sonm.MinerStartReply",
	RunE: grpccmd.RunE(
		"Start",
		"sonm.MinerStartRequest",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Make the Stop method call, input-type: sonm.ID output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"Stop",
		"sonm.ID",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_TasksStatusCmd = &cobra.Command{
	Use:   "tasksStatus",
	Short: "Make the TasksStatus method call, input-type: sonm.MinerStatusMapRequest output-type: sonm.StatusMapReply",
	RunE: grpccmd.RunE(
		"TasksStatus",
		"sonm.MinerStatusMapRequest",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_TaskDetailsCmd = &cobra.Command{
	Use:   "taskDetails",
	Short: "Make the TaskDetails method call, input-type: sonm.ID output-type: sonm.TaskStatusReply",
	RunE: grpccmd.RunE(
		"TaskDetails",
		"sonm.ID",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_TaskLogsCmd = &cobra.Command{
	Use:   "taskLogs",
	Short: "Make the TaskLogs method call, input-type: sonm.TaskLogsRequest output-type: sonm.TaskLogsChunk",
	RunE: grpccmd.RunE(
		"TaskLogs",
		"sonm.TaskLogsRequest",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

var _Miner_DiscoverHubCmd = &cobra.Command{
	Use:   "discoverHub",
	Short: "Make the DiscoverHub method call, input-type: sonm.DiscoverHubRequest output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"DiscoverHub",
		"sonm.DiscoverHubRequest",
		func(c *grpc.ClientConn) interface{} {
			return NewMinerClient(c)
		},
	),
}

// Register commands with the root command and service command
func init() {
	grpccmd.RegisterServiceCmd(_MinerCmd)
	_MinerCmd.AddCommand(
		_Miner_PingCmd,
		_Miner_InfoCmd,
		_Miner_HandshakeCmd,
		_Miner_LoadCmd,
		_Miner_StartCmd,
		_Miner_StopCmd,
		_Miner_TasksStatusCmd,
		_Miner_TaskDetailsCmd,
		_Miner_TaskLogsCmd,
		_Miner_DiscoverHubCmd,
	)
}

// End grpccmd

func init() { proto.RegisterFile("miner.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 878 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc6, 0xfc, 0x04, 0x38, 0x2c, 0x09, 0x3d, 0xd9, 0x74, 0x5d, 0x77, 0xd5, 0x22, 0x77, 0xd5,
	0xd2, 0x1f, 0x45, 0x29, 0xad, 0xa2, 0x76, 0xd5, 0x9b, 0x6c, 0x70, 0xb4, 0x28, 0x04, 0xd0, 0x40,
	0xb4, 0xea, 0x55, 0x34, 0xe0, 0x29, 0x71, 0x01, 0xdb, 0x9d, 0x19, 0x47, 0xe2, 0x09, 0x7a, 0xd3,
	0xbb, 0x3e, 0x4e, 0x5f, 0xae, 0x9a, 0x19, 0x1b, 0x4c, 0x36, 0xb9, 0x62, 0xce, 0x37, 0xe7, 0xe7,
	0x3b, 0xdf, 0x39, 0x1e, 0xa0, 0xb1, 0x0e, 0x42, 0xc6, 0x4f, 0x63, 0x1e, 0xc9, 0x08, 0xcb, 0x22,
	0x0a, 0xd7, 0xce, 0x51, 0x10, 0xaa, 0xdf, 0x30, 0xa0, 0x06, 0x76, 0x70, 0x4e, 0x63, 0x3a, 0x0b,
	0x56, 0x81, 0x0c, 0x98, 0x30, 0x98, 0x3b, 0x82, 0x93, 0x1b, 0x15, 0xf9, 0x9e, 0x86, 0xbe, 0xb8,
	0xa7, 0x4b, 0x46, 0xd8, 0x5f, 0x09, 0x13, 0x12, 0x5b, 0x50, 0xba, 0x4f, 0x66, 0xb6, 0xd5, 0xb6,
	0x3a, 0x75, 0xa2, 0x8e, 0xf8, 0x06, 0x2a, 0x92, 0x8a, 0xa5, 0xb0, 0x8b, 0xed, 0x52, 0xa7, 0xd1,
	0x3d, 0x3c, 0x55, 0xd9, 0x4f, 0xa7, 0x54, 0x2c, 0xfb, 0xe1, 0x1f, 0x11, 0x31, 0x97, 0xee, 0x3f,
	0x16, 0x1c, 0x3f, 0xce, 0x18, 0xaf, 0x36, 0xf8, 0x12, 0x2a, 0x9a, 0x62, 0x9a, 0xd1, 0x18, 0x78,
	0x0e, 0x2f, 0xf2, 0xa4, 0xec, 0x62, 0xdb, 0xea, 0x34, 0xba, 0x68, 0x52, 0x5f, 0xe6, 0x6e, 0xc8,
	0x9e, 0x1f, 0x7e, 0x03, 0xd5, 0x90, 0xca, 0xe9, 0x26, 0x66, 0x76, 0xa9, 0x6d, 0x75, 0x0e, 0xbb,
	0x4d, 0x13, 0x32, 0xbc, 0x98, 0x2a, 0x90, 0x64, 0xb7, 0xee, 0x7f, 0x25, 0xf8, 0x44, 0xd3, 0x99,
	0x48, 0xca, 0x65, 0xd6, 0xdc, 0x21, 0x14, 0x03, 0x3f, 0x65, 0x52, 0x0c, 0x7c, 0x74, 0xa0, 0xc6,
	0xd9, 0x22, 0x10, 0x92, 0x6f, 0x34, 0x85, 0x3a, 0xd9, 0xda, 0x8a, 0x78, 0xb0, 0xa6, 0x0b, 0x53,
	0xa8, 0x4e, 0x8c, 0x81, 0x08, 0x65, 0x9a, 0xc8, 0x7b, 0xbb, 0xac, 0x41, 0x7d, 0xc6, 0x77, 0xd0,
	0xe4, 0x4c, 0xa8, 0x3a, 0xe3, 0x68, 0x15, 0xcc, 0x37, 0x76, 0x45, 0x77, 0xf3, 0x3a, 0xed, 0x26,
	0x0a, 0x25, 0x55, 0x4c, 0x48, 0xde, 0x87, 0xec, 0x87, 0xe0, 0x6f, 0x50, 0xe7, 0x4c, 0x44, 0x09,
	0x9f, 0x33, 0x61, 0x1f, 0xe8, 0xf8, 0x2f, 0x76, 0x42, 0x93, 0xf4, 0x4a, 0xf5, 0x11, 0x70, 0xb6,
	0x66, 0xa1, 0x14, 0x64, 0x17, 0x80, 0x6f, 0xa0, 0x39, 0x4e, 0x66, 0xab, 0x60, 0x7e, 0xcd, 0x36,
	0x3d, 0x2a, 0xa9, 0x5d, 0xd5, 0xf4, 0xf6, 0x41, 0x74, 0xe1, 0xc5, 0x3c, 0x5a, 0xaf, 0x03, 0x39,
	0x0a, 0x27, 0x32, 0x8a, 0xed, 0x5a, 0xdb, 0xea, 0xd4, 0xc8, 0x1e, 0x86, 0x5d, 0x28, 0xb1, 0xf0,
	0xc1, 0xae, 0xeb, 0x51, 0xb7, 0x0d, 0x83, 0x8f, 0x74, 0x3c, 0xf5, 0xc2, 0x07, 0x2f, 0x94, 0x7c,
	0x43, 0x94, 0x33, 0xda, 0x50, 0x8d, 0xb8, 0xcf, 0x78, 0xdf, 0xb7, 0x41, 0xd7, 0xcd, 0x4c, 0xe7,
	0x1c, 0x6a, 0x99, 0xab, 0x5a, 0xac, 0x25, 0xdb, 0x64, 0x8b, 0xb5, 0x64, 0x5a, 0xe1, 0x07, 0xba,
	0x4a, 0x58, 0x2a, 0xbd, 0x31, 0xde, 0x16, 0x7f, 0xb1, 0xdc, 0x9f, 0x01, 0x26, 0xd1, 0x7c, 0xc9,
	0xe4, 0x85, 0xef, 0x73, 0xad, 0xb9, 0xef, 0x67, 0x1b, 0xa4, 0xcf, 0x0a, 0x8b, 0x23, 0x2e, 0x75,
	0x68, 0x93, 0xe8, 0xb3, 0x3b, 0x85, 0xa3, 0x3c, 0x55, 0xb5, 0x7d, 0xaf, 0xa1, 0x3e, 0xcf, 0xf4,
	0x4f, 0xe3, 0x77, 0x00, 0x7e, 0x05, 0x07, 0x3c, 0x4a, 0x24, 0xcb, 0x56, 0xbb, 0x61, 0xfa, 0x25,
	0x0a, 0x23, 0xe9, 0x95, 0xfb, 0x27, 0xd4, 0xb2, 0x5d, 0xc7, 0x1f, 0xa1, 0xca, 0x8d, 0x04, 0x3a,
	0x59, 0xa3, 0xfb, 0xea, 0x19, 0x85, 0x48, 0xe6, 0x87, 0xdf, 0x43, 0x85, 0x2b, 0x2a, 0xe9, 0x8a,
	0x9f, 0x7c, 0x1c, 0x10, 0xaf, 0x36, 0xc4, 0xf8, 0xb8, 0x7d, 0xa8, 0xe8, 0xe2, 0xdb, 0xf6, 0xd2,
	0x96, 0xd5, 0x19, 0x7f, 0x80, 0x1a, 0x0b, 0xfd, 0x38, 0x0a, 0x42, 0x99, 0x26, 0x6b, 0x99, 0x64,
	0x3b, 0xa9, 0xc8, 0xd6, 0xc3, 0x7d, 0x95, 0x7e, 0xe0, 0x13, 0x49, 0x65, 0x22, 0x6e, 0x68, 0x9c,
	0x32, 0xfb, 0xee, 0x6f, 0x0b, 0xaa, 0xe9, 0xe7, 0x82, 0x35, 0x28, 0x0f, 0x47, 0x43, 0xaf, 0x55,
	0xc0, 0x06, 0x54, 0xdf, 0x0d, 0x46, 0x97, 0xd7, 0x5e, 0xaf, 0x65, 0x29, 0xf8, 0xea, 0x76, 0x30,
	0x68, 0x15, 0xb1, 0x09, 0xf5, 0xc9, 0xef, 0x37, 0x37, 0xde, 0x94, 0xf4, 0x2f, 0x5b, 0x25, 0x3c,
	0x04, 0x20, 0xde, 0x44, 0x19, 0x53, 0xaf, 0xd7, 0x2a, 0xe3, 0x31, 0x1c, 0x8d, 0x47, 0x64, 0x7a,
	0x97, 0x03, 0x2b, 0xe8, 0xc0, 0xa7, 0xdb, 0x98, 0xbb, 0xdb, 0xde, 0xf8, 0xee, 0xaa, 0x4f, 0xbc,
	0x0f, 0x17, 0x83, 0x41, 0xeb, 0x40, 0x95, 0xb9, 0x1d, 0x5e, 0x0f, 0x47, 0x1f, 0x86, 0xad, 0x6a,
	0xf7, 0xdf, 0x32, 0x54, 0x34, 0x47, 0xfc, 0x1a, 0xca, 0xe3, 0x20, 0x5c, 0x60, 0x3a, 0x00, 0x6f,
	0x1d, 0xcb, 0x8d, 0x73, 0x64, 0x0c, 0x75, 0xa1, 0x45, 0x72, 0x0b, 0xca, 0x4f, 0xcf, 0xe1, 0x29,
	0x3f, 0xfd, 0x18, 0xa5, 0x7e, 0x1e, 0xd4, 0xb7, 0xcf, 0x10, 0x7e, 0x9e, 0x93, 0xfc, 0xf1, 0x73,
	0xe7, 0x7c, 0xf6, 0xf4, 0xa5, 0x49, 0xf3, 0x2d, 0x94, 0x07, 0x11, 0xf5, 0xb3, 0x72, 0x97, 0xf7,
	0x49, 0xb8, 0x74, 0xd2, 0xf7, 0x6f, 0xcc, 0xa3, 0x05, 0x67, 0x42, 0xb8, 0x85, 0x8e, 0x75, 0x66,
	0xe1, 0xaf, 0x50, 0xd1, 0xe3, 0xc4, 0xe7, 0x36, 0xc2, 0x79, 0x7a, 0xf2, 0x6e, 0x01, 0xbf, 0x84,
	0xb2, 0xfe, 0xf4, 0x6a, 0x69, 0x1f, 0x3d, 0x27, 0xdf, 0x9e, 0x5b, 0xc0, 0x2b, 0x68, 0xa8, 0x0d,
	0x14, 0x66, 0x94, 0x7b, 0xfd, 0x3c, 0x9e, 0xae, 0xf3, 0x32, 0x5d, 0x89, 0x1d, 0xae, 0x8b, 0x68,
	0x8e, 0x67, 0x26, 0x4f, 0x8f, 0x49, 0x1a, 0xac, 0x44, 0xae, 0xde, 0xc9, 0xee, 0xa5, 0x31, 0x81,
	0x19, 0xb5, 0xb7, 0x66, 0xf7, 0x07, 0xd1, 0x42, 0x60, 0xce, 0x49, 0xd9, 0x59, 0xc1, 0xe3, 0x7d,
	0x58, 0x6b, 0xe4, 0x16, 0xce, 0x2c, 0x3c, 0x87, 0x46, 0x2f, 0x10, 0xf3, 0xe8, 0x81, 0xf1, 0xf7,
	0xc9, 0x0c, 0x6d, 0xe3, 0x97, 0x83, 0xb2, 0x0c, 0xfb, 0xdd, 0xce, 0x0e, 0xf4, 0x1f, 0xd4, 0x4f,
	0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x5f, 0x94, 0xa4, 0xda, 0x06, 0x00, 0x00,
}

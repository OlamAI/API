// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simulation.proto

package pb

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SpawnAgentMessage struct {
	X                    int32    `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnAgentMessage) Reset()         { *m = SpawnAgentMessage{} }
func (m *SpawnAgentMessage) String() string { return proto.CompactTextString(m) }
func (*SpawnAgentMessage) ProtoMessage()    {}
func (*SpawnAgentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{0}
}

func (m *SpawnAgentMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnAgentMessage.Unmarshal(m, b)
}
func (m *SpawnAgentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnAgentMessage.Marshal(b, m, deterministic)
}
func (m *SpawnAgentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnAgentMessage.Merge(m, src)
}
func (m *SpawnAgentMessage) XXX_Size() int {
	return xxx_messageInfo_SpawnAgentMessage.Size(m)
}
func (m *SpawnAgentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnAgentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnAgentMessage proto.InternalMessageInfo

func (m *SpawnAgentMessage) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *SpawnAgentMessage) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type SpawnAgentResultMessage struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnAgentResultMessage) Reset()         { *m = SpawnAgentResultMessage{} }
func (m *SpawnAgentResultMessage) String() string { return proto.CompactTextString(m) }
func (*SpawnAgentResultMessage) ProtoMessage()    {}
func (*SpawnAgentResultMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{1}
}

func (m *SpawnAgentResultMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnAgentResultMessage.Unmarshal(m, b)
}
func (m *SpawnAgentResultMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnAgentResultMessage.Marshal(b, m, deterministic)
}
func (m *SpawnAgentResultMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnAgentResultMessage.Merge(m, src)
}
func (m *SpawnAgentResultMessage) XXX_Size() int {
	return xxx_messageInfo_SpawnAgentResultMessage.Size(m)
}
func (m *SpawnAgentResultMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnAgentResultMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnAgentResultMessage proto.InternalMessageInfo

func (m *SpawnAgentResultMessage) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// TODO - implement this function
type ObserveRequestMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObserveRequestMessage) Reset()         { *m = ObserveRequestMessage{} }
func (m *ObserveRequestMessage) String() string { return proto.CompactTextString(m) }
func (*ObserveRequestMessage) ProtoMessage()    {}
func (*ObserveRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{2}
}

func (m *ObserveRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObserveRequestMessage.Unmarshal(m, b)
}
func (m *ObserveRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObserveRequestMessage.Marshal(b, m, deterministic)
}
func (m *ObserveRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserveRequestMessage.Merge(m, src)
}
func (m *ObserveRequestMessage) XXX_Size() int {
	return xxx_messageInfo_ObserveRequestMessage.Size(m)
}
func (m *ObserveRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserveRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ObserveRequestMessage proto.InternalMessageInfo

type EntityMessage struct {
	Class                string   `protobuf:"bytes,1,opt,name=Class,proto3" json:"Class,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityMessage) Reset()         { *m = EntityMessage{} }
func (m *EntityMessage) String() string { return proto.CompactTextString(m) }
func (*EntityMessage) ProtoMessage()    {}
func (*EntityMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{3}
}

func (m *EntityMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityMessage.Unmarshal(m, b)
}
func (m *EntityMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityMessage.Marshal(b, m, deterministic)
}
func (m *EntityMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityMessage.Merge(m, src)
}
func (m *EntityMessage) XXX_Size() int {
	return xxx_messageInfo_EntityMessage.Size(m)
}
func (m *EntityMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EntityMessage proto.InternalMessageInfo

func (m *EntityMessage) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *EntityMessage) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *EntityMessage) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*SpawnAgentMessage)(nil), "pb.SpawnAgentMessage")
	proto.RegisterType((*SpawnAgentResultMessage)(nil), "pb.SpawnAgentResultMessage")
	proto.RegisterType((*ObserveRequestMessage)(nil), "pb.ObserveRequestMessage")
	proto.RegisterType((*EntityMessage)(nil), "pb.EntityMessage")
}

func init() { proto.RegisterFile("simulation.proto", fileDescriptor_961a558581160483) }

var fileDescriptor_961a558581160483 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xce, 0xcc, 0x2d,
	0xcd, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0xd2, 0xe7, 0x12, 0x0c, 0x2e, 0x48, 0x2c, 0xcf, 0x73, 0x4c, 0x4f, 0xcd, 0x2b, 0xf1, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0xe2, 0xe1, 0x62, 0x8c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0d, 0x62, 0x8c, 0x00, 0xf1, 0x22, 0x25, 0x98, 0x20, 0xbc, 0x48, 0x25, 0x43, 0x2e, 0x71, 0x84,
	0x86, 0xa0, 0xd4, 0xe2, 0xd2, 0x1c, 0xb8, 0x36, 0x31, 0x2e, 0xb6, 0xe0, 0x92, 0xc4, 0x92, 0xd2,
	0x62, 0xb0, 0x5e, 0xce, 0x20, 0x28, 0x4f, 0x49, 0x9c, 0x4b, 0xd4, 0x3f, 0xa9, 0x38, 0xb5, 0xa8,
	0x2c, 0x35, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0x18, 0xa6, 0x41, 0xc9, 0x91, 0x8b, 0xd7, 0x35, 0xaf,
	0x24, 0xb3, 0xa4, 0x12, 0x66, 0x82, 0x08, 0x17, 0xab, 0x73, 0x4e, 0x62, 0x31, 0xcc, 0x00, 0x08,
	0x07, 0xe2, 0x1c, 0x26, 0x14, 0xe7, 0x30, 0x43, 0x9d, 0x63, 0xd4, 0xcb, 0xc8, 0xc5, 0x15, 0x0c,
	0xf7, 0x98, 0x90, 0x13, 0x17, 0x17, 0xc2, 0x75, 0x42, 0xa2, 0x7a, 0x05, 0x49, 0x7a, 0x18, 0xde,
	0x93, 0x92, 0x46, 0x15, 0x46, 0xf1, 0x84, 0x12, 0x83, 0x90, 0x35, 0x17, 0x3b, 0xd4, 0xb9, 0x42,
	0x92, 0x20, 0x95, 0x58, 0xdd, 0x2e, 0x25, 0x08, 0x92, 0x42, 0x71, 0xbd, 0x12, 0x83, 0x01, 0x63,
	0x12, 0x1b, 0x38, 0x68, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0xf4, 0xdb, 0xc0, 0x6e,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimulationClient is the client API for Simulation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimulationClient interface {
	// Execute something on data
	SpawnAgent(ctx context.Context, in *SpawnAgentMessage, opts ...grpc.CallOption) (*SpawnAgentResultMessage, error)
	// Observe simulation
	Observe(ctx context.Context, in *ObserveRequestMessage, opts ...grpc.CallOption) (Simulation_ObserveClient, error)
}

type simulationClient struct {
	cc *grpc.ClientConn
}

func NewSimulationClient(cc *grpc.ClientConn) SimulationClient {
	return &simulationClient{cc}
}

func (c *simulationClient) SpawnAgent(ctx context.Context, in *SpawnAgentMessage, opts ...grpc.CallOption) (*SpawnAgentResultMessage, error) {
	out := new(SpawnAgentResultMessage)
	err := c.cc.Invoke(ctx, "/pb.Simulation/SpawnAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) Observe(ctx context.Context, in *ObserveRequestMessage, opts ...grpc.CallOption) (Simulation_ObserveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Simulation_serviceDesc.Streams[0], "/pb.Simulation/Observe", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulationObserveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Simulation_ObserveClient interface {
	Recv() (*EntityMessage, error)
	grpc.ClientStream
}

type simulationObserveClient struct {
	grpc.ClientStream
}

func (x *simulationObserveClient) Recv() (*EntityMessage, error) {
	m := new(EntityMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimulationServer is the server API for Simulation service.
type SimulationServer interface {
	// Execute something on data
	SpawnAgent(context.Context, *SpawnAgentMessage) (*SpawnAgentResultMessage, error)
	// Observe simulation
	Observe(*ObserveRequestMessage, Simulation_ObserveServer) error
}

func RegisterSimulationServer(s *grpc.Server, srv SimulationServer) {
	s.RegisterService(&_Simulation_serviceDesc, srv)
}

func _Simulation_SpawnAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpawnAgentMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationServer).SpawnAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Simulation/SpawnAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationServer).SpawnAgent(ctx, req.(*SpawnAgentMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simulation_Observe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObserveRequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulationServer).Observe(m, &simulationObserveServer{stream})
}

type Simulation_ObserveServer interface {
	Send(*EntityMessage) error
	grpc.ServerStream
}

type simulationObserveServer struct {
	grpc.ServerStream
}

func (x *simulationObserveServer) Send(m *EntityMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _Simulation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Simulation",
	HandlerType: (*SimulationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SpawnAgent",
			Handler:    _Simulation_SpawnAgent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Observe",
			Handler:       _Simulation_Observe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "simulation.proto",
}
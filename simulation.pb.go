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

type SpawnAgentRequest struct {
	X                    int32    `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnAgentRequest) Reset()         { *m = SpawnAgentRequest{} }
func (m *SpawnAgentRequest) String() string { return proto.CompactTextString(m) }
func (*SpawnAgentRequest) ProtoMessage()    {}
func (*SpawnAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{0}
}

func (m *SpawnAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnAgentRequest.Unmarshal(m, b)
}
func (m *SpawnAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnAgentRequest.Marshal(b, m, deterministic)
}
func (m *SpawnAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnAgentRequest.Merge(m, src)
}
func (m *SpawnAgentRequest) XXX_Size() int {
	return xxx_messageInfo_SpawnAgentRequest.Size(m)
}
func (m *SpawnAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnAgentRequest proto.InternalMessageInfo

func (m *SpawnAgentRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *SpawnAgentRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type SpawnAgentResult struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnAgentResult) Reset()         { *m = SpawnAgentResult{} }
func (m *SpawnAgentResult) String() string { return proto.CompactTextString(m) }
func (*SpawnAgentResult) ProtoMessage()    {}
func (*SpawnAgentResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{1}
}

func (m *SpawnAgentResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnAgentResult.Unmarshal(m, b)
}
func (m *SpawnAgentResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnAgentResult.Marshal(b, m, deterministic)
}
func (m *SpawnAgentResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnAgentResult.Merge(m, src)
}
func (m *SpawnAgentResult) XXX_Size() int {
	return xxx_messageInfo_SpawnAgentResult.Size(m)
}
func (m *SpawnAgentResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnAgentResult.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnAgentResult proto.InternalMessageInfo

func (m *SpawnAgentResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AgentObservationRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentObservationRequest) Reset()         { *m = AgentObservationRequest{} }
func (m *AgentObservationRequest) String() string { return proto.CompactTextString(m) }
func (*AgentObservationRequest) ProtoMessage()    {}
func (*AgentObservationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{2}
}

func (m *AgentObservationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentObservationRequest.Unmarshal(m, b)
}
func (m *AgentObservationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentObservationRequest.Marshal(b, m, deterministic)
}
func (m *AgentObservationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentObservationRequest.Merge(m, src)
}
func (m *AgentObservationRequest) XXX_Size() int {
	return xxx_messageInfo_AgentObservationRequest.Size(m)
}
func (m *AgentObservationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentObservationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentObservationRequest proto.InternalMessageInfo

func (m *AgentObservationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AgentObservationResult struct {
	Cells                []string `protobuf:"bytes,1,rep,name=Cells,proto3" json:"Cells,omitempty"`
	Energy               int32    `protobuf:"varint,2,opt,name=Energy,proto3" json:"Energy,omitempty"`
	Health               int32    `protobuf:"varint,3,opt,name=Health,proto3" json:"Health,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentObservationResult) Reset()         { *m = AgentObservationResult{} }
func (m *AgentObservationResult) String() string { return proto.CompactTextString(m) }
func (*AgentObservationResult) ProtoMessage()    {}
func (*AgentObservationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{3}
}

func (m *AgentObservationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentObservationResult.Unmarshal(m, b)
}
func (m *AgentObservationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentObservationResult.Marshal(b, m, deterministic)
}
func (m *AgentObservationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentObservationResult.Merge(m, src)
}
func (m *AgentObservationResult) XXX_Size() int {
	return xxx_messageInfo_AgentObservationResult.Size(m)
}
func (m *AgentObservationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentObservationResult.DiscardUnknown(m)
}

var xxx_messageInfo_AgentObservationResult proto.InternalMessageInfo

func (m *AgentObservationResult) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *AgentObservationResult) GetEnergy() int32 {
	if m != nil {
		return m.Energy
	}
	return 0
}

func (m *AgentObservationResult) GetHealth() int32 {
	if m != nil {
		return m.Health
	}
	return 0
}

type AgentActionRequest struct {
	// TODO - switch to enum instead of string for Action
	// enum ActionType {
	//   UP = 0;
	//   DOWN = 1;
	//   LEFT = 2;
	//   RIGHT = 3;
	// }
	// Id of the agent
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// The action you would like to perform
	Action string `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`
	// The direction the action should go
	Direction            string   `protobuf:"bytes,3,opt,name=Direction,proto3" json:"Direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentActionRequest) Reset()         { *m = AgentActionRequest{} }
func (m *AgentActionRequest) String() string { return proto.CompactTextString(m) }
func (*AgentActionRequest) ProtoMessage()    {}
func (*AgentActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{4}
}

func (m *AgentActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentActionRequest.Unmarshal(m, b)
}
func (m *AgentActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentActionRequest.Marshal(b, m, deterministic)
}
func (m *AgentActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentActionRequest.Merge(m, src)
}
func (m *AgentActionRequest) XXX_Size() int {
	return xxx_messageInfo_AgentActionRequest.Size(m)
}
func (m *AgentActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentActionRequest proto.InternalMessageInfo

func (m *AgentActionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AgentActionRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AgentActionRequest) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

type AgentActionResult struct {
	// Whether or not the action was succesful
	// Could be unsuccesful if attempting to interract with an entity
	//  that is not in range, already consumed, etc.
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentActionResult) Reset()         { *m = AgentActionResult{} }
func (m *AgentActionResult) String() string { return proto.CompactTextString(m) }
func (*AgentActionResult) ProtoMessage()    {}
func (*AgentActionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{5}
}

func (m *AgentActionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentActionResult.Unmarshal(m, b)
}
func (m *AgentActionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentActionResult.Marshal(b, m, deterministic)
}
func (m *AgentActionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentActionResult.Merge(m, src)
}
func (m *AgentActionResult) XXX_Size() int {
	return xxx_messageInfo_AgentActionResult.Size(m)
}
func (m *AgentActionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentActionResult.DiscardUnknown(m)
}

var xxx_messageInfo_AgentActionResult proto.InternalMessageInfo

func (m *AgentActionResult) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

// TODO - implement this
type SpectateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpectateRequest) Reset()         { *m = SpectateRequest{} }
func (m *SpectateRequest) String() string { return proto.CompactTextString(m) }
func (*SpectateRequest) ProtoMessage()    {}
func (*SpectateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{6}
}

func (m *SpectateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpectateRequest.Unmarshal(m, b)
}
func (m *SpectateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpectateRequest.Marshal(b, m, deterministic)
}
func (m *SpectateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpectateRequest.Merge(m, src)
}
func (m *SpectateRequest) XXX_Size() int {
	return xxx_messageInfo_SpectateRequest.Size(m)
}
func (m *SpectateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpectateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpectateRequest proto.InternalMessageInfo

type CellUpdate struct {
	X                    int32    `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Occupant             string   `protobuf:"bytes,3,opt,name=Occupant,proto3" json:"Occupant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CellUpdate) Reset()         { *m = CellUpdate{} }
func (m *CellUpdate) String() string { return proto.CompactTextString(m) }
func (*CellUpdate) ProtoMessage()    {}
func (*CellUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_961a558581160483, []int{7}
}

func (m *CellUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellUpdate.Unmarshal(m, b)
}
func (m *CellUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellUpdate.Marshal(b, m, deterministic)
}
func (m *CellUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellUpdate.Merge(m, src)
}
func (m *CellUpdate) XXX_Size() int {
	return xxx_messageInfo_CellUpdate.Size(m)
}
func (m *CellUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_CellUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_CellUpdate proto.InternalMessageInfo

func (m *CellUpdate) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CellUpdate) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CellUpdate) GetOccupant() string {
	if m != nil {
		return m.Occupant
	}
	return ""
}

func init() {
	proto.RegisterType((*SpawnAgentRequest)(nil), "pb.SpawnAgentRequest")
	proto.RegisterType((*SpawnAgentResult)(nil), "pb.SpawnAgentResult")
	proto.RegisterType((*AgentObservationRequest)(nil), "pb.AgentObservationRequest")
	proto.RegisterType((*AgentObservationResult)(nil), "pb.AgentObservationResult")
	proto.RegisterType((*AgentActionRequest)(nil), "pb.AgentActionRequest")
	proto.RegisterType((*AgentActionResult)(nil), "pb.AgentActionResult")
	proto.RegisterType((*SpectateRequest)(nil), "pb.SpectateRequest")
	proto.RegisterType((*CellUpdate)(nil), "pb.CellUpdate")
}

func init() { proto.RegisterFile("simulation.proto", fileDescriptor_961a558581160483) }

var fileDescriptor_961a558581160483 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0xa5, 0x25, 0x10, 0x7a, 0xdf, 0x0b, 0x0f, 0xee, 0x03, 0x6c, 0xaa, 0x31, 0x64, 0x56, 0xb8,
	0x41, 0x23, 0x4b, 0x13, 0x13, 0x22, 0x26, 0xb2, 0x30, 0x24, 0x6d, 0x4c, 0xc0, 0x85, 0x49, 0x5b,
	0x46, 0x24, 0xa9, 0x65, 0xec, 0x4c, 0x35, 0xfe, 0x08, 0xff, 0xb3, 0x99, 0x0f, 0xbe, 0x95, 0xe5,
	0x39, 0xf7, 0xde, 0xd3, 0x33, 0xe7, 0x14, 0x6a, 0x7c, 0xfe, 0x9a, 0x27, 0xa1, 0x98, 0x2f, 0xd2,
	0x2e, 0xcb, 0x16, 0x62, 0x81, 0x36, 0x8b, 0xc8, 0x39, 0xd4, 0x03, 0x16, 0x7e, 0xa4, 0xfd, 0x19,
	0x4d, 0x85, 0x4f, 0xdf, 0x72, 0xca, 0x05, 0xfe, 0x05, 0x6b, 0xec, 0x5a, 0x6d, 0xab, 0x53, 0xf2,
	0xad, 0xb1, 0x44, 0x13, 0xd7, 0xd6, 0x68, 0x42, 0x08, 0xd4, 0x36, 0x0f, 0x78, 0x9e, 0x08, 0xac,
	0x82, 0x3d, 0x9c, 0xaa, 0x03, 0xc7, 0xb7, 0x87, 0x53, 0x72, 0x06, 0x47, 0x6a, 0x3c, 0x8a, 0x38,
	0xcd, 0xde, 0xd5, 0x27, 0x97, 0xd2, 0xbb, 0xab, 0x4f, 0xd0, 0xda, 0x5f, 0x55, 0xa2, 0x0d, 0x28,
	0xdd, 0xd0, 0x24, 0xe1, 0xae, 0xd5, 0x2e, 0x76, 0x1c, 0x5f, 0x03, 0x6c, 0x41, 0xf9, 0x36, 0xa5,
	0xd9, 0xec, 0xd3, 0x38, 0x32, 0x48, 0xf2, 0x77, 0x34, 0x4c, 0xc4, 0x8b, 0x5b, 0xd4, 0xbc, 0x46,
	0xe4, 0x11, 0x50, 0xe9, 0xf7, 0xe3, 0x03, 0x2e, 0xe4, 0xb5, 0x5e, 0x50, 0xaa, 0x8e, 0x6f, 0x10,
	0x9e, 0x80, 0x33, 0x98, 0x67, 0x54, 0x8f, 0x8a, 0x6a, 0xb4, 0x26, 0x48, 0x0f, 0xea, 0x5b, 0xda,
	0xca, 0xf6, 0x29, 0x40, 0x90, 0xc7, 0x31, 0xe5, 0xfc, 0x39, 0x4f, 0xd4, 0x27, 0x2a, 0xfe, 0x06,
	0x43, 0xea, 0xf0, 0x2f, 0x60, 0x34, 0x16, 0xa1, 0xa0, 0xc6, 0x0d, 0x19, 0x00, 0xc8, 0xc7, 0x3d,
	0xb0, 0x69, 0x28, 0xe8, 0xa1, 0xf0, 0xd1, 0x83, 0xca, 0x28, 0x8e, 0x73, 0x16, 0xa6, 0xc2, 0xd8,
	0x59, 0xe1, 0xcb, 0x2f, 0x1b, 0x20, 0x58, 0x55, 0x8c, 0x57, 0x00, 0xeb, 0x9e, 0xb0, 0xd9, 0x65,
	0x51, 0x77, 0xaf, 0x68, 0xaf, 0xb1, 0x4b, 0xcb, 0x27, 0x90, 0x02, 0xde, 0x43, 0x6d, 0xb7, 0x15,
	0x3c, 0x96, 0xbb, 0xbf, 0xd4, 0xea, 0x79, 0x3f, 0x0f, 0x8d, 0xdc, 0x35, 0xfc, 0xd9, 0x08, 0x0a,
	0x5b, 0xab, 0xe5, 0xad, 0x56, 0xbc, 0xe6, 0x1e, 0x6f, 0xee, 0x7b, 0x50, 0x59, 0x66, 0x86, 0xff,
	0xb5, 0xe5, 0xad, 0x04, 0xbd, 0xaa, 0x24, 0xd7, 0x19, 0x92, 0xc2, 0x85, 0x15, 0x95, 0xd5, 0x4f,
	0xde, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xa7, 0xdf, 0x97, 0xf8, 0x02, 0x00, 0x00,
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
	// Spawn a new agent
	SpawnAgent(ctx context.Context, in *SpawnAgentRequest, opts ...grpc.CallOption) (*SpawnAgentResult, error)
	// Request observation data from an agent
	AgentObservation(ctx context.Context, in *AgentObservationRequest, opts ...grpc.CallOption) (*AgentObservationResult, error)
	// Perform an action on behalf of an agent
	AgentAction(ctx context.Context, in *AgentActionRequest, opts ...grpc.CallOption) (*AgentActionResult, error)
	// Spectate simulation
	Spectate(ctx context.Context, in *SpectateRequest, opts ...grpc.CallOption) (Simulation_SpectateClient, error)
}

type simulationClient struct {
	cc *grpc.ClientConn
}

func NewSimulationClient(cc *grpc.ClientConn) SimulationClient {
	return &simulationClient{cc}
}

func (c *simulationClient) SpawnAgent(ctx context.Context, in *SpawnAgentRequest, opts ...grpc.CallOption) (*SpawnAgentResult, error) {
	out := new(SpawnAgentResult)
	err := c.cc.Invoke(ctx, "/pb.Simulation/SpawnAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) AgentObservation(ctx context.Context, in *AgentObservationRequest, opts ...grpc.CallOption) (*AgentObservationResult, error) {
	out := new(AgentObservationResult)
	err := c.cc.Invoke(ctx, "/pb.Simulation/AgentObservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) AgentAction(ctx context.Context, in *AgentActionRequest, opts ...grpc.CallOption) (*AgentActionResult, error) {
	out := new(AgentActionResult)
	err := c.cc.Invoke(ctx, "/pb.Simulation/AgentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) Spectate(ctx context.Context, in *SpectateRequest, opts ...grpc.CallOption) (Simulation_SpectateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Simulation_serviceDesc.Streams[0], "/pb.Simulation/Spectate", opts...)
	if err != nil {
		return nil, err
	}
	x := &simulationSpectateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Simulation_SpectateClient interface {
	Recv() (*CellUpdate, error)
	grpc.ClientStream
}

type simulationSpectateClient struct {
	grpc.ClientStream
}

func (x *simulationSpectateClient) Recv() (*CellUpdate, error) {
	m := new(CellUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimulationServer is the server API for Simulation service.
type SimulationServer interface {
	// Spawn a new agent
	SpawnAgent(context.Context, *SpawnAgentRequest) (*SpawnAgentResult, error)
	// Request observation data from an agent
	AgentObservation(context.Context, *AgentObservationRequest) (*AgentObservationResult, error)
	// Perform an action on behalf of an agent
	AgentAction(context.Context, *AgentActionRequest) (*AgentActionResult, error)
	// Spectate simulation
	Spectate(*SpectateRequest, Simulation_SpectateServer) error
}

func RegisterSimulationServer(s *grpc.Server, srv SimulationServer) {
	s.RegisterService(&_Simulation_serviceDesc, srv)
}

func _Simulation_SpawnAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpawnAgentRequest)
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
		return srv.(SimulationServer).SpawnAgent(ctx, req.(*SpawnAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simulation_AgentObservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentObservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationServer).AgentObservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Simulation/AgentObservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationServer).AgentObservation(ctx, req.(*AgentObservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simulation_AgentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationServer).AgentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Simulation/AgentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationServer).AgentAction(ctx, req.(*AgentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simulation_Spectate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SpectateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimulationServer).Spectate(m, &simulationSpectateServer{stream})
}

type Simulation_SpectateServer interface {
	Send(*CellUpdate) error
	grpc.ServerStream
}

type simulationSpectateServer struct {
	grpc.ServerStream
}

func (x *simulationSpectateServer) Send(m *CellUpdate) error {
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
		{
			MethodName: "AgentObservation",
			Handler:    _Simulation_AgentObservation_Handler,
		},
		{
			MethodName: "AgentAction",
			Handler:    _Simulation_AgentAction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Spectate",
			Handler:       _Simulation_Spectate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "simulation.proto",
}

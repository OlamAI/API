// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simulation/simulation.proto

package olamai_endpoints_simulation

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
	return fileDescriptor_17e12b66aec6c312, []int{0}
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
	return fileDescriptor_17e12b66aec6c312, []int{1}
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
	return fileDescriptor_17e12b66aec6c312, []int{2}
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
	Alive                bool     `protobuf:"varint,1,opt,name=Alive,proto3" json:"Alive,omitempty"`
	Cells                []string `protobuf:"bytes,2,rep,name=Cells,proto3" json:"Cells,omitempty"`
	Energy               int32    `protobuf:"varint,3,opt,name=Energy,proto3" json:"Energy,omitempty"`
	Health               int32    `protobuf:"varint,4,opt,name=Health,proto3" json:"Health,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentObservationResult) Reset()         { *m = AgentObservationResult{} }
func (m *AgentObservationResult) String() string { return proto.CompactTextString(m) }
func (*AgentObservationResult) ProtoMessage()    {}
func (*AgentObservationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e12b66aec6c312, []int{3}
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

func (m *AgentObservationResult) GetAlive() bool {
	if m != nil {
		return m.Alive
	}
	return false
}

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

type ResetWorldRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetWorldRequest) Reset()         { *m = ResetWorldRequest{} }
func (m *ResetWorldRequest) String() string { return proto.CompactTextString(m) }
func (*ResetWorldRequest) ProtoMessage()    {}
func (*ResetWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e12b66aec6c312, []int{4}
}

func (m *ResetWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetWorldRequest.Unmarshal(m, b)
}
func (m *ResetWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetWorldRequest.Marshal(b, m, deterministic)
}
func (m *ResetWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetWorldRequest.Merge(m, src)
}
func (m *ResetWorldRequest) XXX_Size() int {
	return xxx_messageInfo_ResetWorldRequest.Size(m)
}
func (m *ResetWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetWorldRequest proto.InternalMessageInfo

type ResetWorldResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetWorldResult) Reset()         { *m = ResetWorldResult{} }
func (m *ResetWorldResult) String() string { return proto.CompactTextString(m) }
func (*ResetWorldResult) ProtoMessage()    {}
func (*ResetWorldResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e12b66aec6c312, []int{5}
}

func (m *ResetWorldResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetWorldResult.Unmarshal(m, b)
}
func (m *ResetWorldResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetWorldResult.Marshal(b, m, deterministic)
}
func (m *ResetWorldResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetWorldResult.Merge(m, src)
}
func (m *ResetWorldResult) XXX_Size() int {
	return xxx_messageInfo_ResetWorldResult.Size(m)
}
func (m *ResetWorldResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetWorldResult.DiscardUnknown(m)
}

var xxx_messageInfo_ResetWorldResult proto.InternalMessageInfo

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
	return fileDescriptor_17e12b66aec6c312, []int{6}
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
	return fileDescriptor_17e12b66aec6c312, []int{7}
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
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpectateRequest) Reset()         { *m = SpectateRequest{} }
func (m *SpectateRequest) String() string { return proto.CompactTextString(m) }
func (*SpectateRequest) ProtoMessage()    {}
func (*SpectateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e12b66aec6c312, []int{8}
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

func (m *SpectateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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
	return fileDescriptor_17e12b66aec6c312, []int{9}
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

type SubscribeToRegionRequest struct {
	// TODO - Put this in metadata instead of packet
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeToRegionRequest) Reset()         { *m = SubscribeToRegionRequest{} }
func (m *SubscribeToRegionRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeToRegionRequest) ProtoMessage()    {}
func (*SubscribeToRegionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e12b66aec6c312, []int{10}
}

func (m *SubscribeToRegionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeToRegionRequest.Unmarshal(m, b)
}
func (m *SubscribeToRegionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeToRegionRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeToRegionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeToRegionRequest.Merge(m, src)
}
func (m *SubscribeToRegionRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeToRegionRequest.Size(m)
}
func (m *SubscribeToRegionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeToRegionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeToRegionRequest proto.InternalMessageInfo

func (m *SubscribeToRegionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SubscribeToRegionRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *SubscribeToRegionRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type SubscribeToRegionResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeToRegionResult) Reset()         { *m = SubscribeToRegionResult{} }
func (m *SubscribeToRegionResult) String() string { return proto.CompactTextString(m) }
func (*SubscribeToRegionResult) ProtoMessage()    {}
func (*SubscribeToRegionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e12b66aec6c312, []int{11}
}

func (m *SubscribeToRegionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeToRegionResult.Unmarshal(m, b)
}
func (m *SubscribeToRegionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeToRegionResult.Marshal(b, m, deterministic)
}
func (m *SubscribeToRegionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeToRegionResult.Merge(m, src)
}
func (m *SubscribeToRegionResult) XXX_Size() int {
	return xxx_messageInfo_SubscribeToRegionResult.Size(m)
}
func (m *SubscribeToRegionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeToRegionResult.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeToRegionResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SpawnAgentRequest)(nil), "olamai.endpoints.simulation.SpawnAgentRequest")
	proto.RegisterType((*SpawnAgentResult)(nil), "olamai.endpoints.simulation.SpawnAgentResult")
	proto.RegisterType((*AgentObservationRequest)(nil), "olamai.endpoints.simulation.AgentObservationRequest")
	proto.RegisterType((*AgentObservationResult)(nil), "olamai.endpoints.simulation.AgentObservationResult")
	proto.RegisterType((*ResetWorldRequest)(nil), "olamai.endpoints.simulation.ResetWorldRequest")
	proto.RegisterType((*ResetWorldResult)(nil), "olamai.endpoints.simulation.ResetWorldResult")
	proto.RegisterType((*AgentActionRequest)(nil), "olamai.endpoints.simulation.AgentActionRequest")
	proto.RegisterType((*AgentActionResult)(nil), "olamai.endpoints.simulation.AgentActionResult")
	proto.RegisterType((*SpectateRequest)(nil), "olamai.endpoints.simulation.SpectateRequest")
	proto.RegisterType((*CellUpdate)(nil), "olamai.endpoints.simulation.CellUpdate")
	proto.RegisterType((*SubscribeToRegionRequest)(nil), "olamai.endpoints.simulation.SubscribeToRegionRequest")
	proto.RegisterType((*SubscribeToRegionResult)(nil), "olamai.endpoints.simulation.SubscribeToRegionResult")
}

func init() { proto.RegisterFile("simulation/simulation.proto", fileDescriptor_17e12b66aec6c312) }

var fileDescriptor_17e12b66aec6c312 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x63, 0x87, 0x54, 0xc9, 0x80, 0x20, 0x19, 0x50, 0x6b, 0x5c, 0x84, 0xca, 0x5e, 0x28,
	0x12, 0x38, 0x88, 0x94, 0x07, 0x88, 0x28, 0x88, 0x9e, 0x2a, 0xd9, 0x20, 0x5a, 0x6e, 0x1b, 0x7b,
	0x08, 0x96, 0x36, 0xb6, 0xf1, 0xae, 0x8b, 0x38, 0x21, 0xf1, 0x76, 0xbc, 0x15, 0xf2, 0xae, 0x53,
	0x27, 0x35, 0x71, 0x9a, 0x5b, 0xfe, 0xb3, 0xf3, 0x3d, 0x3f, 0x07, 0x0e, 0x65, 0xbc, 0x28, 0x04,
	0x57, 0x71, 0x9a, 0x8c, 0xeb, 0x9f, 0x5e, 0x96, 0xa7, 0x2a, 0xc5, 0xc3, 0x54, 0xf0, 0x05, 0x8f,
	0x3d, 0x4a, 0xa2, 0x2c, 0x8d, 0x13, 0x25, 0xbd, 0xda, 0x85, 0x8d, 0x61, 0x14, 0x64, 0xfc, 0x67,
	0x32, 0x9d, 0x53, 0xa2, 0x7c, 0xfa, 0x51, 0x90, 0x54, 0x78, 0x0f, 0xac, 0x0b, 0xc7, 0x3a, 0xb2,
	0x8e, 0x7b, 0xbe, 0x75, 0x51, 0xaa, 0x4b, 0xc7, 0x36, 0xea, 0x92, 0x31, 0x18, 0xae, 0x06, 0xc8,
	0x42, 0x28, 0xbc, 0x0f, 0xf6, 0x59, 0xa4, 0x03, 0x06, 0xbe, 0x7d, 0x16, 0xb1, 0x17, 0x70, 0xa0,
	0x9f, 0xcf, 0x67, 0x92, 0xf2, 0x2b, 0x5d, 0x68, 0x99, 0xfa, 0xa6, 0xab, 0x82, 0xfd, 0xa6, 0xab,
	0x4e, 0xfa, 0x08, 0x7a, 0x53, 0x11, 0x5f, 0x91, 0x76, 0xee, 0xfb, 0x46, 0x94, 0xd6, 0x77, 0x24,
	0x84, 0x74, 0xec, 0xa3, 0xee, 0xf1, 0xc0, 0x37, 0x02, 0xf7, 0x61, 0xef, 0x7d, 0x42, 0xf9, 0xfc,
	0x97, 0xd3, 0xd5, 0x7d, 0x56, 0xaa, 0xb4, 0x7f, 0x24, 0x2e, 0xd4, 0x77, 0xe7, 0x8e, 0xb1, 0x1b,
	0xc5, 0x1e, 0xc2, 0xc8, 0x27, 0x49, 0xea, 0x4b, 0x9a, 0x8b, 0xa8, 0x6a, 0x8d, 0x21, 0x0c, 0x57,
	0x8d, 0x65, 0x13, 0xec, 0x2b, 0xa0, 0x6e, 0x6f, 0x1a, 0xb6, 0x0c, 0x51, 0x96, 0x31, 0x0e, 0x7a,
	0x4d, 0x03, 0xbf, 0x52, 0xf8, 0x04, 0x06, 0xa7, 0x71, 0x4e, 0xe6, 0xa9, 0xab, 0x9f, 0x6a, 0x03,
	0x9b, 0xc0, 0x68, 0x2d, 0xb7, 0x9e, 0xfa, 0x29, 0x40, 0x50, 0x84, 0x21, 0x49, 0xf9, 0xad, 0x10,
	0xd5, 0xe8, 0x2b, 0x16, 0xf6, 0x0c, 0x1e, 0x04, 0x19, 0x85, 0x8a, 0x2b, 0xda, 0xb4, 0xd2, 0x53,
	0x80, 0x72, 0x2b, 0x9f, 0xb3, 0x88, 0x2b, 0x6a, 0xbb, 0x25, 0xba, 0xd0, 0x3f, 0x0f, 0xc3, 0x22,
	0xe3, 0x89, 0xaa, 0xda, 0xbb, 0xd6, 0xec, 0x03, 0x38, 0x41, 0x31, 0x93, 0x61, 0x1e, 0xcf, 0xe8,
	0x53, 0xea, 0xd3, 0xbc, 0x65, 0x7e, 0x5d, 0xc3, 0x5e, 0xab, 0xd1, 0x5d, 0xf2, 0xf2, 0x18, 0x0e,
	0xfe, 0x93, 0xa7, 0x9c, 0xf5, 0xcd, 0xdf, 0x1e, 0x40, 0x70, 0x8d, 0x22, 0x2e, 0x00, 0x6a, 0xb2,
	0xd0, 0xf3, 0x5a, 0xb0, 0xf5, 0x1a, 0xcc, 0xba, 0xaf, 0x6e, 0xed, 0xaf, 0x0f, 0xdb, 0xc1, 0xdf,
	0x30, 0xbc, 0x49, 0x1e, 0x9e, 0xb4, 0x26, 0xd9, 0xc0, 0xb4, 0x3b, 0xd9, 0x31, 0xaa, 0x6a, 0x20,
	0x83, 0xbb, 0x2b, 0xf7, 0xc7, 0xf1, 0xf6, 0x2c, 0x6b, 0x14, 0xba, 0xde, 0xed, 0x03, 0xaa, 0x8a,
	0x0b, 0x80, 0x9a, 0xf0, 0x2d, 0x1b, 0x6e, 0x7c, 0x1f, 0x5b, 0x36, 0xdc, 0xf8, 0x74, 0x3a, 0x48,
	0xd0, 0x5f, 0xb2, 0x8a, 0x2f, 0xb7, 0x9c, 0x67, 0x0d, 0x69, 0xf7, 0x79, 0xab, 0x77, 0x4d, 0x37,
	0xeb, 0xbc, 0xb6, 0xf0, 0x8f, 0x05, 0xa3, 0x06, 0x62, 0xf8, 0xb6, 0xbd, 0xe0, 0x06, 0xb4, 0xdd,
	0x93, 0x5d, 0xc3, 0xcc, 0xac, 0xb3, 0x3d, 0xfd, 0x5f, 0x3b, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff,
	0x22, 0x0e, 0x0e, 0x1f, 0x8a, 0x05, 0x00, 0x00,
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
	// Reset the world
	ResetWorld(ctx context.Context, in *ResetWorldRequest, opts ...grpc.CallOption) (*ResetWorldResult, error)
	// Spectate simulation
	Spectate(ctx context.Context, in *SpectateRequest, opts ...grpc.CallOption) (Simulation_SpectateClient, error)
	// Subscribe to region
	SubscribeToRegion(ctx context.Context, in *SubscribeToRegionRequest, opts ...grpc.CallOption) (*SubscribeToRegionResult, error)
}

type simulationClient struct {
	cc *grpc.ClientConn
}

func NewSimulationClient(cc *grpc.ClientConn) SimulationClient {
	return &simulationClient{cc}
}

func (c *simulationClient) SpawnAgent(ctx context.Context, in *SpawnAgentRequest, opts ...grpc.CallOption) (*SpawnAgentResult, error) {
	out := new(SpawnAgentResult)
	err := c.cc.Invoke(ctx, "/olamai.endpoints.simulation.Simulation/SpawnAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) AgentObservation(ctx context.Context, in *AgentObservationRequest, opts ...grpc.CallOption) (*AgentObservationResult, error) {
	out := new(AgentObservationResult)
	err := c.cc.Invoke(ctx, "/olamai.endpoints.simulation.Simulation/AgentObservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) AgentAction(ctx context.Context, in *AgentActionRequest, opts ...grpc.CallOption) (*AgentActionResult, error) {
	out := new(AgentActionResult)
	err := c.cc.Invoke(ctx, "/olamai.endpoints.simulation.Simulation/AgentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) ResetWorld(ctx context.Context, in *ResetWorldRequest, opts ...grpc.CallOption) (*ResetWorldResult, error) {
	out := new(ResetWorldResult)
	err := c.cc.Invoke(ctx, "/olamai.endpoints.simulation.Simulation/ResetWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulationClient) Spectate(ctx context.Context, in *SpectateRequest, opts ...grpc.CallOption) (Simulation_SpectateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Simulation_serviceDesc.Streams[0], "/olamai.endpoints.simulation.Simulation/Spectate", opts...)
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

func (c *simulationClient) SubscribeToRegion(ctx context.Context, in *SubscribeToRegionRequest, opts ...grpc.CallOption) (*SubscribeToRegionResult, error) {
	out := new(SubscribeToRegionResult)
	err := c.cc.Invoke(ctx, "/olamai.endpoints.simulation.Simulation/SubscribeToRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimulationServer is the server API for Simulation service.
type SimulationServer interface {
	// Spawn a new agent
	SpawnAgent(context.Context, *SpawnAgentRequest) (*SpawnAgentResult, error)
	// Request observation data from an agent
	AgentObservation(context.Context, *AgentObservationRequest) (*AgentObservationResult, error)
	// Perform an action on behalf of an agent
	AgentAction(context.Context, *AgentActionRequest) (*AgentActionResult, error)
	// Reset the world
	ResetWorld(context.Context, *ResetWorldRequest) (*ResetWorldResult, error)
	// Spectate simulation
	Spectate(*SpectateRequest, Simulation_SpectateServer) error
	// Subscribe to region
	SubscribeToRegion(context.Context, *SubscribeToRegionRequest) (*SubscribeToRegionResult, error)
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
		FullMethod: "/olamai.endpoints.simulation.Simulation/SpawnAgent",
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
		FullMethod: "/olamai.endpoints.simulation.Simulation/AgentObservation",
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
		FullMethod: "/olamai.endpoints.simulation.Simulation/AgentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationServer).AgentAction(ctx, req.(*AgentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simulation_ResetWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationServer).ResetWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/olamai.endpoints.simulation.Simulation/ResetWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationServer).ResetWorld(ctx, req.(*ResetWorldRequest))
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

func _Simulation_SubscribeToRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeToRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulationServer).SubscribeToRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/olamai.endpoints.simulation.Simulation/SubscribeToRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulationServer).SubscribeToRegion(ctx, req.(*SubscribeToRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Simulation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "olamai.endpoints.simulation.Simulation",
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
		{
			MethodName: "ResetWorld",
			Handler:    _Simulation_ResetWorld_Handler,
		},
		{
			MethodName: "SubscribeToRegion",
			Handler:    _Simulation_SubscribeToRegion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Spectate",
			Handler:       _Simulation_Spectate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "simulation/simulation.proto",
}
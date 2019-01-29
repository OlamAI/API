// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datacom.proto

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

type ExecuteMessage struct {
	Execution            string   `protobuf:"bytes,1,opt,name=Execution,proto3" json:"Execution,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteMessage) Reset()         { *m = ExecuteMessage{} }
func (m *ExecuteMessage) String() string { return proto.CompactTextString(m) }
func (*ExecuteMessage) ProtoMessage()    {}
func (*ExecuteMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e946789395f79ab3, []int{0}
}

func (m *ExecuteMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteMessage.Unmarshal(m, b)
}
func (m *ExecuteMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteMessage.Marshal(b, m, deterministic)
}
func (m *ExecuteMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteMessage.Merge(m, src)
}
func (m *ExecuteMessage) XXX_Size() int {
	return xxx_messageInfo_ExecuteMessage.Size(m)
}
func (m *ExecuteMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteMessage proto.InternalMessageInfo

func (m *ExecuteMessage) GetExecution() string {
	if m != nil {
		return m.Execution
	}
	return ""
}

type QueryMessage struct {
	Query                string   `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryMessage) Reset()         { *m = QueryMessage{} }
func (m *QueryMessage) String() string { return proto.CompactTextString(m) }
func (*QueryMessage) ProtoMessage()    {}
func (*QueryMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e946789395f79ab3, []int{1}
}

func (m *QueryMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMessage.Unmarshal(m, b)
}
func (m *QueryMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMessage.Marshal(b, m, deterministic)
}
func (m *QueryMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMessage.Merge(m, src)
}
func (m *QueryMessage) XXX_Size() int {
	return xxx_messageInfo_QueryMessage.Size(m)
}
func (m *QueryMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMessage proto.InternalMessageInfo

func (m *QueryMessage) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type QueryResponseMessage struct {
	Response             string   `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponseMessage) Reset()         { *m = QueryResponseMessage{} }
func (m *QueryResponseMessage) String() string { return proto.CompactTextString(m) }
func (*QueryResponseMessage) ProtoMessage()    {}
func (*QueryResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e946789395f79ab3, []int{2}
}

func (m *QueryResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponseMessage.Unmarshal(m, b)
}
func (m *QueryResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponseMessage.Marshal(b, m, deterministic)
}
func (m *QueryResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponseMessage.Merge(m, src)
}
func (m *QueryResponseMessage) XXX_Size() int {
	return xxx_messageInfo_QueryResponseMessage.Size(m)
}
func (m *QueryResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponseMessage proto.InternalMessageInfo

func (m *QueryResponseMessage) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e946789395f79ab3, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExecuteMessage)(nil), "pb.ExecuteMessage")
	proto.RegisterType((*QueryMessage)(nil), "pb.QueryMessage")
	proto.RegisterType((*QueryResponseMessage)(nil), "pb.QueryResponseMessage")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

func init() { proto.RegisterFile("datacom.proto", fileDescriptor_e946789395f79ab3) }

var fileDescriptor_e946789395f79ab3 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0x2c, 0x49,
	0x4c, 0xce, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe3,
	0xe2, 0x73, 0xad, 0x48, 0x4d, 0x2e, 0x2d, 0x49, 0xf5, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15,
	0x92, 0xe1, 0xe2, 0x84, 0x88, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21,
	0x04, 0x94, 0x54, 0xb8, 0x78, 0x02, 0x4b, 0x53, 0x8b, 0x2a, 0x61, 0xaa, 0x45, 0xb8, 0x58, 0xc1,
	0x7c, 0xa8, 0x4a, 0x08, 0x47, 0xc9, 0x88, 0x4b, 0x04, 0xcc, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf,
	0x2b, 0x86, 0x9b, 0x2d, 0xc5, 0xc5, 0x01, 0x13, 0x82, 0x6a, 0x80, 0xf3, 0x95, 0xd8, 0xb9, 0x58,
	0x5d, 0x73, 0x0b, 0x4a, 0x2a, 0x8d, 0x72, 0xb8, 0xd8, 0x5d, 0x20, 0xee, 0x14, 0xd2, 0xe2, 0x62,
	0x87, 0xba, 0x4e, 0x48, 0x48, 0xaf, 0x20, 0x49, 0x0f, 0xd5, 0xa9, 0x52, 0x9c, 0x60, 0x31, 0x90,
	0x26, 0x25, 0x06, 0x21, 0x53, 0xa8, 0x4b, 0x84, 0x04, 0x40, 0xa2, 0xc8, 0x8e, 0x94, 0x92, 0x80,
	0x8b, 0xa0, 0x39, 0x48, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x16, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x6a, 0x13, 0x90, 0x1c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatacomClient is the client API for Datacom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatacomClient interface {
	// Execute something on data
	Execute(ctx context.Context, in *ExecuteMessage, opts ...grpc.CallOption) (*Empty, error)
	// Query data
	Query(ctx context.Context, in *QueryMessage, opts ...grpc.CallOption) (*QueryResponseMessage, error)
}

type datacomClient struct {
	cc *grpc.ClientConn
}

func NewDatacomClient(cc *grpc.ClientConn) DatacomClient {
	return &datacomClient{cc}
}

func (c *datacomClient) Execute(ctx context.Context, in *ExecuteMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.Datacom/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datacomClient) Query(ctx context.Context, in *QueryMessage, opts ...grpc.CallOption) (*QueryResponseMessage, error) {
	out := new(QueryResponseMessage)
	err := c.cc.Invoke(ctx, "/pb.Datacom/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatacomServer is the server API for Datacom service.
type DatacomServer interface {
	// Execute something on data
	Execute(context.Context, *ExecuteMessage) (*Empty, error)
	// Query data
	Query(context.Context, *QueryMessage) (*QueryResponseMessage, error)
}

func RegisterDatacomServer(s *grpc.Server, srv DatacomServer) {
	s.RegisterService(&_Datacom_serviceDesc, srv)
}

func _Datacom_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatacomServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Datacom/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatacomServer).Execute(ctx, req.(*ExecuteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datacom_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatacomServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Datacom/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatacomServer).Query(ctx, req.(*QueryMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Datacom_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Datacom",
	HandlerType: (*DatacomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Datacom_Execute_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Datacom_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datacom.proto",
}
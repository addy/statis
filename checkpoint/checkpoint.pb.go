// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkpoint.proto

package checkpoint

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

type CheckpointRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointRequest) Reset()         { *m = CheckpointRequest{} }
func (m *CheckpointRequest) String() string { return proto.CompactTextString(m) }
func (*CheckpointRequest) ProtoMessage()    {}
func (*CheckpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bab050ffa824783, []int{0}
}

func (m *CheckpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointRequest.Unmarshal(m, b)
}
func (m *CheckpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointRequest.Marshal(b, m, deterministic)
}
func (m *CheckpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointRequest.Merge(m, src)
}
func (m *CheckpointRequest) XXX_Size() int {
	return xxx_messageInfo_CheckpointRequest.Size(m)
}
func (m *CheckpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointRequest proto.InternalMessageInfo

func (m *CheckpointRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CheckpointRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CheckpointRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type CheckpointReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointReply) Reset()         { *m = CheckpointReply{} }
func (m *CheckpointReply) String() string { return proto.CompactTextString(m) }
func (*CheckpointReply) ProtoMessage()    {}
func (*CheckpointReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bab050ffa824783, []int{1}
}

func (m *CheckpointReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointReply.Unmarshal(m, b)
}
func (m *CheckpointReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointReply.Marshal(b, m, deterministic)
}
func (m *CheckpointReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointReply.Merge(m, src)
}
func (m *CheckpointReply) XXX_Size() int {
	return xxx_messageInfo_CheckpointReply.Size(m)
}
func (m *CheckpointReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointReply.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointReply proto.InternalMessageInfo

func (m *CheckpointReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*CheckpointRequest)(nil), "checkpoint.CheckpointRequest")
	proto.RegisterType((*CheckpointReply)(nil), "checkpoint.CheckpointReply")
}

func init() { proto.RegisterFile("checkpoint.proto", fileDescriptor_9bab050ffa824783) }

var fileDescriptor_9bab050ffa824783 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0x48, 0x4d,
	0xce, 0x2e, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x45, 0x73, 0x09, 0x3a, 0xc3, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90,
	0x08, 0x17, 0x6b, 0x71, 0x49, 0x62, 0x49, 0xaa, 0x04, 0x13, 0x58, 0x10, 0xc2, 0x11, 0x92, 0xe1,
	0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x56, 0x60, 0xd4, 0x60,
	0x0e, 0x42, 0x08, 0x28, 0x69, 0x72, 0xf1, 0x23, 0x1b, 0x5e, 0x90, 0x53, 0x29, 0x24, 0xc6, 0xc5,
	0x06, 0xd2, 0x59, 0x5a, 0x0c, 0x36, 0x9c, 0x35, 0x08, 0xca, 0x33, 0x8a, 0xe3, 0xe2, 0x42, 0x28,
	0x15, 0x0a, 0xe0, 0x12, 0x70, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x45, 0x12, 0x93, 0xd5, 0x43, 0xf2,
	0x08, 0x86, 0x9b, 0xa5, 0xa4, 0x71, 0x49, 0x17, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0xbd,
	0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xae, 0xb2, 0xbd, 0x64, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CheckpointClient is the client API for Checkpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckpointClient interface {
	CreateCheckpoint(ctx context.Context, in *CheckpointRequest, opts ...grpc.CallOption) (*CheckpointReply, error)
}

type checkpointClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckpointClient(cc grpc.ClientConnInterface) CheckpointClient {
	return &checkpointClient{cc}
}

func (c *checkpointClient) CreateCheckpoint(ctx context.Context, in *CheckpointRequest, opts ...grpc.CallOption) (*CheckpointReply, error) {
	out := new(CheckpointReply)
	err := c.cc.Invoke(ctx, "/checkpoint.Checkpoint/CreateCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckpointServer is the server API for Checkpoint service.
type CheckpointServer interface {
	CreateCheckpoint(context.Context, *CheckpointRequest) (*CheckpointReply, error)
}

// UnimplementedCheckpointServer can be embedded to have forward compatible implementations.
type UnimplementedCheckpointServer struct {
}

func (*UnimplementedCheckpointServer) CreateCheckpoint(ctx context.Context, req *CheckpointRequest) (*CheckpointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCheckpoint not implemented")
}

func RegisterCheckpointServer(s *grpc.Server, srv CheckpointServer) {
	s.RegisterService(&_Checkpoint_serviceDesc, srv)
}

func _Checkpoint_CreateCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckpointServer).CreateCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkpoint.Checkpoint/CreateCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckpointServer).CreateCheckpoint(ctx, req.(*CheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Checkpoint_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checkpoint.Checkpoint",
	HandlerType: (*CheckpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCheckpoint",
			Handler:    _Checkpoint_CreateCheckpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkpoint.proto",
}
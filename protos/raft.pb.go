// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.7.1
// source: raft.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RaftMsg_Type int32

const (
	RaftMsg_UNDEFINED RaftMsg_Type = 0
	RaftMsg_ERROR     RaftMsg_Type = 1
	RaftMsg_RESPONSE  RaftMsg_Type = 2
	RaftMsg_REDIRECT  RaftMsg_Type = 3
	RaftMsg_HEARTBEAT RaftMsg_Type = 4
	RaftMsg_VOTE      RaftMsg_Type = 5
	RaftMsg_GET       RaftMsg_Type = 6
	RaftMsg_SET       RaftMsg_Type = 7
)

// Enum value maps for RaftMsg_Type.
var (
	RaftMsg_Type_name = map[int32]string{
		0: "UNDEFINED",
		1: "ERROR",
		2: "RESPONSE",
		3: "REDIRECT",
		4: "HEARTBEAT",
		5: "VOTE",
		6: "GET",
		7: "SET",
	}
	RaftMsg_Type_value = map[string]int32{
		"UNDEFINED": 0,
		"ERROR":     1,
		"RESPONSE":  2,
		"REDIRECT":  3,
		"HEARTBEAT": 4,
		"VOTE":      5,
		"GET":       6,
		"SET":       7,
	}
)

func (x RaftMsg_Type) Enum() *RaftMsg_Type {
	p := new(RaftMsg_Type)
	*p = x
	return p
}

func (x RaftMsg_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RaftMsg_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_raft_proto_enumTypes[0].Descriptor()
}

func (RaftMsg_Type) Type() protoreflect.EnumType {
	return &file_raft_proto_enumTypes[0]
}

func (x RaftMsg_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RaftMsg_Type.Descriptor instead.
func (RaftMsg_Type) EnumDescriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{3, 0}
}

type KVPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term  uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVPair) Reset() {
	*x = KVPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVPair) ProtoMessage() {}

func (x *KVPair) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVPair.ProtoReflect.Descriptor instead.
func (*KVPair) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{0}
}

func (x *KVPair) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *KVPair) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *KVPair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVPair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term           uint64    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	SubmittedIndex uint64    `protobuf:"varint,2,opt,name=submittedIndex,proto3" json:"submittedIndex,omitempty"`
	Length         uint32    `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	KvPair         []*KVPair `protobuf:"bytes,4,rep,name=kvPair,proto3" json:"kvPair,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{1}
}

func (x *Heartbeat) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Heartbeat) GetSubmittedIndex() uint64 {
	if x != nil {
		return x.SubmittedIndex
	}
	return 0
}

func (x *Heartbeat) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Heartbeat) GetKvPair() []*KVPair {
	if x != nil {
		return x.KvPair
	}
	return nil
}

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{2}
}

func (x *Vote) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

type RaftMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    RaftMsg_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.RaftMsg_Type" json:"type,omitempty"`
	Payload []byte       `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RaftMsg) Reset() {
	*x = RaftMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftMsg) ProtoMessage() {}

func (x *RaftMsg) ProtoReflect() protoreflect.Message {
	mi := &file_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftMsg.ProtoReflect.Descriptor instead.
func (*RaftMsg) Descriptor() ([]byte, []int) {
	return file_raft_proto_rawDescGZIP(), []int{3}
}

func (x *RaftMsg) GetType() RaftMsg_Type {
	if x != nil {
		return x.Type
	}
	return RaftMsg_UNDEFINED
}

func (x *RaftMsg) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_raft_proto protoreflect.FileDescriptor

var file_raft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x5a, 0x0a, 0x06, 0x4b, 0x56, 0x50, 0x61, 0x69, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x87, 0x01, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x26, 0x0a, 0x06, 0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4b, 0x56, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x06, 0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x22, 0x1a, 0x0a, 0x04, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0xb6, 0x01, 0x0a, 0x07, 0x52, 0x61, 0x66, 0x74, 0x4d,
	0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x4d, 0x73,
	0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x67, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41,
	0x54, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4f, 0x54, 0x45, 0x10, 0x05, 0x12, 0x07, 0x0a,
	0x03, 0x47, 0x45, 0x54, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x07, 0x32,
	0x3a, 0x0a, 0x0b, 0x52, 0x61, 0x66, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2b,
	0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x52, 0x61, 0x66, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x52, 0x61, 0x66, 0x74, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_raft_proto_rawDescOnce sync.Once
	file_raft_proto_rawDescData = file_raft_proto_rawDesc
)

func file_raft_proto_rawDescGZIP() []byte {
	file_raft_proto_rawDescOnce.Do(func() {
		file_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_raft_proto_rawDescData)
	})
	return file_raft_proto_rawDescData
}

var file_raft_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_raft_proto_goTypes = []interface{}{
	(RaftMsg_Type)(0), // 0: protos.RaftMsg.Type
	(*KVPair)(nil),    // 1: protos.KVPair
	(*Heartbeat)(nil), // 2: protos.Heartbeat
	(*Vote)(nil),      // 3: protos.Vote
	(*RaftMsg)(nil),   // 4: protos.RaftMsg
}
var file_raft_proto_depIdxs = []int32{
	1, // 0: protos.Heartbeat.kvPair:type_name -> protos.KVPair
	0, // 1: protos.RaftMsg.type:type_name -> protos.RaftMsg.Type
	4, // 2: protos.RaftSupport.Agent:input_type -> protos.RaftMsg
	4, // 3: protos.RaftSupport.Agent:output_type -> protos.RaftMsg
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_raft_proto_init() }
func file_raft_proto_init() {
	if File_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVPair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raft_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raft_proto_goTypes,
		DependencyIndexes: file_raft_proto_depIdxs,
		EnumInfos:         file_raft_proto_enumTypes,
		MessageInfos:      file_raft_proto_msgTypes,
	}.Build()
	File_raft_proto = out.File
	file_raft_proto_rawDesc = nil
	file_raft_proto_goTypes = nil
	file_raft_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RaftSupportClient is the client API for RaftSupport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftSupportClient interface {
	Agent(ctx context.Context, in *RaftMsg, opts ...grpc.CallOption) (*RaftMsg, error)
}

type raftSupportClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftSupportClient(cc grpc.ClientConnInterface) RaftSupportClient {
	return &raftSupportClient{cc}
}

func (c *raftSupportClient) Agent(ctx context.Context, in *RaftMsg, opts ...grpc.CallOption) (*RaftMsg, error) {
	out := new(RaftMsg)
	err := c.cc.Invoke(ctx, "/protos.RaftSupport/Agent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftSupportServer is the server API for RaftSupport service.
type RaftSupportServer interface {
	Agent(context.Context, *RaftMsg) (*RaftMsg, error)
}

// UnimplementedRaftSupportServer can be embedded to have forward compatible implementations.
type UnimplementedRaftSupportServer struct {
}

func (*UnimplementedRaftSupportServer) Agent(context.Context, *RaftMsg) (*RaftMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Agent not implemented")
}

func RegisterRaftSupportServer(s *grpc.Server, srv RaftSupportServer) {
	s.RegisterService(&_RaftSupport_serviceDesc, srv)
}

func _RaftSupport_Agent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftSupportServer).Agent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.RaftSupport/Agent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftSupportServer).Agent(ctx, req.(*RaftMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.RaftSupport",
	HandlerType: (*RaftSupportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Agent",
			Handler:    _RaftSupport_Agent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}

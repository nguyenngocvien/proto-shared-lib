// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.2
// source: group/rpc_get_group_members.proto

package group

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetGroupMembersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupId       string                 `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGroupMembersRequest) Reset() {
	*x = GetGroupMembersRequest{}
	mi := &file_group_rpc_get_group_members_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupMembersRequest) ProtoMessage() {}

func (x *GetGroupMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_get_group_members_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupMembersRequest.ProtoReflect.Descriptor instead.
func (*GetGroupMembersRequest) Descriptor() ([]byte, []int) {
	return file_group_rpc_get_group_members_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupMembersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetGroupMembersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        []string               `protobuf:"bytes,1,rep,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGroupMembersResponse) Reset() {
	*x = GetGroupMembersResponse{}
	mi := &file_group_rpc_get_group_members_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupMembersResponse) ProtoMessage() {}

func (x *GetGroupMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_get_group_members_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupMembersResponse.ProtoReflect.Descriptor instead.
func (*GetGroupMembersResponse) Descriptor() ([]byte, []int) {
	return file_group_rpc_get_group_members_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupMembersResponse) GetUserId() []string {
	if x != nil {
		return x.UserId
	}
	return nil
}

var File_group_rpc_get_group_members_proto protoreflect.FileDescriptor

const file_group_rpc_get_group_members_proto_rawDesc = "" +
	"\n" +
	"!group/rpc_get_group_members.proto\x12\x05group\"3\n" +
	"\x16GetGroupMembersRequest\x12\x19\n" +
	"\bgroup_id\x18\x01 \x01(\tR\agroupId\"2\n" +
	"\x17GetGroupMembersResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x03(\tR\x06userIdB6Z4github.com/nguyenngocvien/proto-shared-lib/gen/groupb\x06proto3"

var (
	file_group_rpc_get_group_members_proto_rawDescOnce sync.Once
	file_group_rpc_get_group_members_proto_rawDescData []byte
)

func file_group_rpc_get_group_members_proto_rawDescGZIP() []byte {
	file_group_rpc_get_group_members_proto_rawDescOnce.Do(func() {
		file_group_rpc_get_group_members_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_group_rpc_get_group_members_proto_rawDesc), len(file_group_rpc_get_group_members_proto_rawDesc)))
	})
	return file_group_rpc_get_group_members_proto_rawDescData
}

var file_group_rpc_get_group_members_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_group_rpc_get_group_members_proto_goTypes = []any{
	(*GetGroupMembersRequest)(nil),  // 0: group.GetGroupMembersRequest
	(*GetGroupMembersResponse)(nil), // 1: group.GetGroupMembersResponse
}
var file_group_rpc_get_group_members_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_group_rpc_get_group_members_proto_init() }
func file_group_rpc_get_group_members_proto_init() {
	if File_group_rpc_get_group_members_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_group_rpc_get_group_members_proto_rawDesc), len(file_group_rpc_get_group_members_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_group_rpc_get_group_members_proto_goTypes,
		DependencyIndexes: file_group_rpc_get_group_members_proto_depIdxs,
		MessageInfos:      file_group_rpc_get_group_members_proto_msgTypes,
	}.Build()
	File_group_rpc_get_group_members_proto = out.File
	file_group_rpc_get_group_members_proto_goTypes = nil
	file_group_rpc_get_group_members_proto_depIdxs = nil
}

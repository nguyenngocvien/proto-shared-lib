// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.2
// source: chat/rpc_send_mesage.proto

package chat

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

type SendMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *ChatMessage           `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	mi := &file_chat_rpc_send_mesage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_rpc_send_mesage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_rpc_send_mesage_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetMessage() *ChatMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	mi := &file_chat_rpc_send_mesage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_rpc_send_mesage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_chat_rpc_send_mesage_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_chat_rpc_send_mesage_proto protoreflect.FileDescriptor

const file_chat_rpc_send_mesage_proto_rawDesc = "" +
	"\n" +
	"\x1achat/rpc_send_mesage.proto\x12\x04chat\x1a\x13chat/rpc_chat.proto\"A\n" +
	"\x12SendMessageRequest\x12+\n" +
	"\amessage\x18\x01 \x01(\v2\x11.chat.ChatMessageR\amessage\"/\n" +
	"\x13SendMessageResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccessB5Z3github.com/nguyenngocvien/proto-shared-lib/gen/chatb\x06proto3"

var (
	file_chat_rpc_send_mesage_proto_rawDescOnce sync.Once
	file_chat_rpc_send_mesage_proto_rawDescData []byte
)

func file_chat_rpc_send_mesage_proto_rawDescGZIP() []byte {
	file_chat_rpc_send_mesage_proto_rawDescOnce.Do(func() {
		file_chat_rpc_send_mesage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_rpc_send_mesage_proto_rawDesc), len(file_chat_rpc_send_mesage_proto_rawDesc)))
	})
	return file_chat_rpc_send_mesage_proto_rawDescData
}

var file_chat_rpc_send_mesage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_rpc_send_mesage_proto_goTypes = []any{
	(*SendMessageRequest)(nil),  // 0: chat.SendMessageRequest
	(*SendMessageResponse)(nil), // 1: chat.SendMessageResponse
	(*ChatMessage)(nil),         // 2: chat.ChatMessage
}
var file_chat_rpc_send_mesage_proto_depIdxs = []int32{
	2, // 0: chat.SendMessageRequest.message:type_name -> chat.ChatMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_rpc_send_mesage_proto_init() }
func file_chat_rpc_send_mesage_proto_init() {
	if File_chat_rpc_send_mesage_proto != nil {
		return
	}
	file_chat_rpc_chat_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_rpc_send_mesage_proto_rawDesc), len(file_chat_rpc_send_mesage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_rpc_send_mesage_proto_goTypes,
		DependencyIndexes: file_chat_rpc_send_mesage_proto_depIdxs,
		MessageInfos:      file_chat_rpc_send_mesage_proto_msgTypes,
	}.Build()
	File_chat_rpc_send_mesage_proto = out.File
	file_chat_rpc_send_mesage_proto_goTypes = nil
	file_chat_rpc_send_mesage_proto_depIdxs = nil
}

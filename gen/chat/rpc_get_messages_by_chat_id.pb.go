// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.2
// source: chat/rpc_get_messages_by_chat_id.proto

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

type GetMessagesByChatIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"` // From path
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                // From query
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`              // From query
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMessagesByChatIDRequest) Reset() {
	*x = GetMessagesByChatIDRequest{}
	mi := &file_chat_rpc_get_messages_by_chat_id_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessagesByChatIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesByChatIDRequest) ProtoMessage() {}

func (x *GetMessagesByChatIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_rpc_get_messages_by_chat_id_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesByChatIDRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesByChatIDRequest) Descriptor() ([]byte, []int) {
	return file_chat_rpc_get_messages_by_chat_id_proto_rawDescGZIP(), []int{0}
}

func (x *GetMessagesByChatIDRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *GetMessagesByChatIDRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetMessagesByChatIDRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatId        string                 `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	SenderId      string                 `protobuf:"bytes,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Text          string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp     string                 `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // ISO 8601 string
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_chat_rpc_get_messages_by_chat_id_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_rpc_get_messages_by_chat_id_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_rpc_get_messages_by_chat_id_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type GetMessagesByChatIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Message             `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMessagesByChatIDResponse) Reset() {
	*x = GetMessagesByChatIDResponse{}
	mi := &file_chat_rpc_get_messages_by_chat_id_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessagesByChatIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesByChatIDResponse) ProtoMessage() {}

func (x *GetMessagesByChatIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_rpc_get_messages_by_chat_id_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesByChatIDResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesByChatIDResponse) Descriptor() ([]byte, []int) {
	return file_chat_rpc_get_messages_by_chat_id_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessagesByChatIDResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *GetMessagesByChatIDResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetMessagesByChatIDResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_chat_rpc_get_messages_by_chat_id_proto protoreflect.FileDescriptor

const file_chat_rpc_get_messages_by_chat_id_proto_rawDesc = "" +
	"\n" +
	"&chat/rpc_get_messages_by_chat_id.proto\x12\x04chat\"c\n" +
	"\x1aGetMessagesByChatIDRequest\x12\x17\n" +
	"\achat_id\x18\x01 \x01(\tR\x06chatId\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x03 \x01(\x05R\x06offset\"\x81\x01\n" +
	"\aMessage\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\achat_id\x18\x02 \x01(\tR\x06chatId\x12\x1b\n" +
	"\tsender_id\x18\x03 \x01(\tR\bsenderId\x12\x12\n" +
	"\x04text\x18\x04 \x01(\tR\x04text\x12\x1c\n" +
	"\ttimestamp\x18\x05 \x01(\tR\ttimestamp\"v\n" +
	"\x1bGetMessagesByChatIDResponse\x12)\n" +
	"\bmessages\x18\x01 \x03(\v2\r.chat.MessageR\bmessages\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x03 \x01(\x05R\x06offsetB5Z3github.com/nguyenngocvien/proto-shared-lib/gen/chatb\x06proto3"

var (
	file_chat_rpc_get_messages_by_chat_id_proto_rawDescOnce sync.Once
	file_chat_rpc_get_messages_by_chat_id_proto_rawDescData []byte
)

func file_chat_rpc_get_messages_by_chat_id_proto_rawDescGZIP() []byte {
	file_chat_rpc_get_messages_by_chat_id_proto_rawDescOnce.Do(func() {
		file_chat_rpc_get_messages_by_chat_id_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_rpc_get_messages_by_chat_id_proto_rawDesc), len(file_chat_rpc_get_messages_by_chat_id_proto_rawDesc)))
	})
	return file_chat_rpc_get_messages_by_chat_id_proto_rawDescData
}

var file_chat_rpc_get_messages_by_chat_id_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_rpc_get_messages_by_chat_id_proto_goTypes = []any{
	(*GetMessagesByChatIDRequest)(nil),  // 0: chat.GetMessagesByChatIDRequest
	(*Message)(nil),                     // 1: chat.Message
	(*GetMessagesByChatIDResponse)(nil), // 2: chat.GetMessagesByChatIDResponse
}
var file_chat_rpc_get_messages_by_chat_id_proto_depIdxs = []int32{
	1, // 0: chat.GetMessagesByChatIDResponse.messages:type_name -> chat.Message
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_rpc_get_messages_by_chat_id_proto_init() }
func file_chat_rpc_get_messages_by_chat_id_proto_init() {
	if File_chat_rpc_get_messages_by_chat_id_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_rpc_get_messages_by_chat_id_proto_rawDesc), len(file_chat_rpc_get_messages_by_chat_id_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_rpc_get_messages_by_chat_id_proto_goTypes,
		DependencyIndexes: file_chat_rpc_get_messages_by_chat_id_proto_depIdxs,
		MessageInfos:      file_chat_rpc_get_messages_by_chat_id_proto_msgTypes,
	}.Build()
	File_chat_rpc_get_messages_by_chat_id_proto = out.File
	file_chat_rpc_get_messages_by_chat_id_proto_goTypes = nil
	file_chat_rpc_get_messages_by_chat_id_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.2
// source: workflow/rpc_cancel_workflow.proto

package worflow

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

type CancelWorkflowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkflowId    string                 `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelWorkflowRequest) Reset() {
	*x = CancelWorkflowRequest{}
	mi := &file_workflow_rpc_cancel_workflow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelWorkflowRequest) ProtoMessage() {}

func (x *CancelWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_rpc_cancel_workflow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelWorkflowRequest.ProtoReflect.Descriptor instead.
func (*CancelWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_workflow_rpc_cancel_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *CancelWorkflowRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

type CancelWorkflowResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelWorkflowResponse) Reset() {
	*x = CancelWorkflowResponse{}
	mi := &file_workflow_rpc_cancel_workflow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelWorkflowResponse) ProtoMessage() {}

func (x *CancelWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_rpc_cancel_workflow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelWorkflowResponse.ProtoReflect.Descriptor instead.
func (*CancelWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_workflow_rpc_cancel_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *CancelWorkflowResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_workflow_rpc_cancel_workflow_proto protoreflect.FileDescriptor

const file_workflow_rpc_cancel_workflow_proto_rawDesc = "" +
	"\n" +
	"\"workflow/rpc_cancel_workflow.proto\x12\bworkflow\"8\n" +
	"\x15CancelWorkflowRequest\x12\x1f\n" +
	"\vworkflow_id\x18\x01 \x01(\tR\n" +
	"workflowId\"0\n" +
	"\x16CancelWorkflowResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06statusB8Z6github.com/nguyenngocvien/proto-shared-lib/gen/worflowb\x06proto3"

var (
	file_workflow_rpc_cancel_workflow_proto_rawDescOnce sync.Once
	file_workflow_rpc_cancel_workflow_proto_rawDescData []byte
)

func file_workflow_rpc_cancel_workflow_proto_rawDescGZIP() []byte {
	file_workflow_rpc_cancel_workflow_proto_rawDescOnce.Do(func() {
		file_workflow_rpc_cancel_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_workflow_rpc_cancel_workflow_proto_rawDesc), len(file_workflow_rpc_cancel_workflow_proto_rawDesc)))
	})
	return file_workflow_rpc_cancel_workflow_proto_rawDescData
}

var file_workflow_rpc_cancel_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_workflow_rpc_cancel_workflow_proto_goTypes = []any{
	(*CancelWorkflowRequest)(nil),  // 0: workflow.CancelWorkflowRequest
	(*CancelWorkflowResponse)(nil), // 1: workflow.CancelWorkflowResponse
}
var file_workflow_rpc_cancel_workflow_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_workflow_rpc_cancel_workflow_proto_init() }
func file_workflow_rpc_cancel_workflow_proto_init() {
	if File_workflow_rpc_cancel_workflow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_workflow_rpc_cancel_workflow_proto_rawDesc), len(file_workflow_rpc_cancel_workflow_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflow_rpc_cancel_workflow_proto_goTypes,
		DependencyIndexes: file_workflow_rpc_cancel_workflow_proto_depIdxs,
		MessageInfos:      file_workflow_rpc_cancel_workflow_proto_msgTypes,
	}.Build()
	File_workflow_rpc_cancel_workflow_proto = out.File
	file_workflow_rpc_cancel_workflow_proto_goTypes = nil
	file_workflow_rpc_cancel_workflow_proto_depIdxs = nil
}

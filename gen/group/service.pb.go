// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.2
// source: group/service.proto

package group

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_group_service_proto protoreflect.FileDescriptor

const file_group_service_proto_rawDesc = "" +
	"\n" +
	"\x13group/service.proto\x12\x05group\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1dgroup/rpc_search_groups.proto\x1a\x1cgroup/rpc_create_group.proto\x1a\x1cgroup/rpc_delete_group.proto\x1a\x1bgroup/rpc_add_members.proto\x1a!group/rpc_get_group_members.proto2\xa0\a\n" +
	"\fGroupService\x12\x9f\x01\n" +
	"\tGetGroups\x12\x1a.group.SearchGroupsRequest\x1a\x1b.group.SearchGroupsResponse\"Y\x92AD\n" +
	"\x05Group\x12\n" +
	"Get groups\x1a\x1dRetrieve a list of all groupsb\x10\n" +
	"\x0e\n" +
	"\n" +
	"BearerAuth\x12\x00\x82\xd3\xe4\x93\x02\f\x12\n" +
	"/v1/groups\x12\x99\x01\n" +
	"\vCreateGroup\x12\x19.group.CreateGroupRequest\x1a\x1a.group.CreateGroupResponse\"S\x92A;\n" +
	"\x05Group\x12\fCreate group\x1a\x12Create a new groupb\x10\n" +
	"\x0e\n" +
	"\n" +
	"BearerAuth\x12\x00\x82\xd3\xe4\x93\x02\x0f:\x01*\"\n" +
	"/v1/groups\x12\xa1\x01\n" +
	"\vDeleteGroup\x12\x19.group.DeleteGroupRequest\x1a\x1a.group.DeleteGroupResponse\"[\x92AA\n" +
	"\x05Group\x12\fDelete group\x1a\x18Delete a group by its IDb\x10\n" +
	"\x0e\n" +
	"\n" +
	"BearerAuth\x12\x00\x82\xd3\xe4\x93\x02\x11*\x0f/v1/groups/{id}\x12\xd2\x01\n" +
	"\x11AddMembersToGroup\x12\x18.group.AddMembersRequest\x1a\x19.group.AddMembersResponse\"\x87\x01\x92A\\\n" +
	"\x05Group\x12\x14Add members to group\x1a+Add one or more members to a specific groupb\x10\n" +
	"\x0e\n" +
	"\n" +
	"BearerAuth\x12\x00\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/v1/groups/{group_id}/members\x12\xd8\x01\n" +
	"\x0fGetGroupMembers\x12\x1d.group.GetGroupMembersRequest\x1a\x1e.group.GetGroupMembersResponse\"\x85\x01\x92A]\n" +
	"\x05Group\x12\x11Get group members\x1a/Retrieve a list of members for a specific groupb\x10\n" +
	"\x0e\n" +
	"\n" +
	"BearerAuth\x12\x00\x82\xd3\xe4\x93\x02\x1f\x12\x1d/v1/groups/{group_id}/membersB6Z4github.com/nguyenngocvien/proto-shared-lib/gen/groupb\x06proto3"

var file_group_service_proto_goTypes = []any{
	(*SearchGroupsRequest)(nil),     // 0: group.SearchGroupsRequest
	(*CreateGroupRequest)(nil),      // 1: group.CreateGroupRequest
	(*DeleteGroupRequest)(nil),      // 2: group.DeleteGroupRequest
	(*AddMembersRequest)(nil),       // 3: group.AddMembersRequest
	(*GetGroupMembersRequest)(nil),  // 4: group.GetGroupMembersRequest
	(*SearchGroupsResponse)(nil),    // 5: group.SearchGroupsResponse
	(*CreateGroupResponse)(nil),     // 6: group.CreateGroupResponse
	(*DeleteGroupResponse)(nil),     // 7: group.DeleteGroupResponse
	(*AddMembersResponse)(nil),      // 8: group.AddMembersResponse
	(*GetGroupMembersResponse)(nil), // 9: group.GetGroupMembersResponse
}
var file_group_service_proto_depIdxs = []int32{
	0, // 0: group.GroupService.GetGroups:input_type -> group.SearchGroupsRequest
	1, // 1: group.GroupService.CreateGroup:input_type -> group.CreateGroupRequest
	2, // 2: group.GroupService.DeleteGroup:input_type -> group.DeleteGroupRequest
	3, // 3: group.GroupService.AddMembersToGroup:input_type -> group.AddMembersRequest
	4, // 4: group.GroupService.GetGroupMembers:input_type -> group.GetGroupMembersRequest
	5, // 5: group.GroupService.GetGroups:output_type -> group.SearchGroupsResponse
	6, // 6: group.GroupService.CreateGroup:output_type -> group.CreateGroupResponse
	7, // 7: group.GroupService.DeleteGroup:output_type -> group.DeleteGroupResponse
	8, // 8: group.GroupService.AddMembersToGroup:output_type -> group.AddMembersResponse
	9, // 9: group.GroupService.GetGroupMembers:output_type -> group.GetGroupMembersResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_group_service_proto_init() }
func file_group_service_proto_init() {
	if File_group_service_proto != nil {
		return
	}
	file_group_rpc_search_groups_proto_init()
	file_group_rpc_create_group_proto_init()
	file_group_rpc_delete_group_proto_init()
	file_group_rpc_add_members_proto_init()
	file_group_rpc_get_group_members_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_group_service_proto_rawDesc), len(file_group_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_service_proto_goTypes,
		DependencyIndexes: file_group_service_proto_depIdxs,
	}.Build()
	File_group_service_proto = out.File
	file_group_service_proto_goTypes = nil
	file_group_service_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.5.1
// source: admin.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa8, 0x03,
	0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x31, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x12, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x37, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x12, 0x14, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0a, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_admin_proto_goTypes = []interface{}{
	(*LoginReq)(nil),        // 0: messages.LoginReq
	(*LogoutReq)(nil),       // 1: messages.LogoutReq
	(*UserAddReq)(nil),      // 2: messages.UserAddReq
	(*UserUpdateReq)(nil),   // 3: messages.UserUpdateReq
	(*UserQueryReq)(nil),    // 4: messages.UserQueryReq
	(*UserListReq)(nil),     // 5: messages.UserListReq
	(*UserDeleteReq)(nil),   // 6: messages.UserDeleteReq
	(*LoginReply)(nil),      // 7: messages.LoginReply
	(*LogoutReply)(nil),     // 8: messages.LogoutReply
	(*UserAddReply)(nil),    // 9: messages.UserAddReply
	(*UserUpdateReply)(nil), // 10: messages.UserUpdateReply
	(*UserQueryReply)(nil),  // 11: messages.UserQueryReply
	(*UserListReply)(nil),   // 12: messages.UserListReply
	(*UserDeleteReply)(nil), // 13: messages.UserDeleteReply
}
var file_admin_proto_depIdxs = []int32{
	0,  // 0: admin.Admin.Login:input_type -> messages.LoginReq
	1,  // 1: admin.Admin.Logout:input_type -> messages.LogoutReq
	2,  // 2: admin.Admin.UserAdd:input_type -> messages.UserAddReq
	3,  // 3: admin.Admin.UserUpdate:input_type -> messages.UserUpdateReq
	4,  // 4: admin.Admin.UserQuery:input_type -> messages.UserQueryReq
	5,  // 5: admin.Admin.UserList:input_type -> messages.UserListReq
	6,  // 6: admin.Admin.UserDelete:input_type -> messages.UserDeleteReq
	7,  // 7: admin.Admin.Login:output_type -> messages.LoginReply
	8,  // 8: admin.Admin.Logout:output_type -> messages.LogoutReply
	9,  // 9: admin.Admin.UserAdd:output_type -> messages.UserAddReply
	10, // 10: admin.Admin.UserUpdate:output_type -> messages.UserUpdateReply
	11, // 11: admin.Admin.UserQuery:output_type -> messages.UserQueryReply
	12, // 12: admin.Admin.UserList:output_type -> messages.UserListReply
	13, // 13: admin.Admin.UserDelete:output_type -> messages.UserDeleteReply
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	file_message_login_proto_init()
	file_message_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}

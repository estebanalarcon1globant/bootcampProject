// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: grpc/user.proto

package grpc

import (
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

type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PwdHash string `protobuf:"bytes,1,opt,name=pwd_hash,json=pwdHash,proto3" json:"pwd_hash,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_grpc_user_proto_rawDescGZIP(), []int{0}
}

func (x *NewUser) GetPwdHash() string {
	if x != nil {
		return x.PwdHash
	}
	return ""
}

func (x *NewUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewUser) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PwdHash string `protobuf:"bytes,2,opt,name=pwd_hash,json=pwdHash,proto3" json:"pwd_hash,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_grpc_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetPwdHash() string {
	if x != nil {
		return x.PwdHash
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type GetUsersParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetUsersParams) Reset() {
	*x = GetUsersParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersParams) ProtoMessage() {}

func (x *GetUsersParams) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersParams.ProtoReflect.Descriptor instead.
func (*GetUsersParams) Descriptor() ([]byte, []int) {
	return file_grpc_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersParams) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUsersParams) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_grpc_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserList) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_grpc_user_proto protoreflect.FileDescriptor

var file_grpc_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x4a, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x77, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x77, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x77, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x77, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x6c, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x62, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_user_proto_rawDescOnce sync.Once
	file_grpc_user_proto_rawDescData = file_grpc_user_proto_rawDesc
)

func file_grpc_user_proto_rawDescGZIP() []byte {
	file_grpc_user_proto_rawDescOnce.Do(func() {
		file_grpc_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_user_proto_rawDescData)
	})
	return file_grpc_user_proto_rawDescData
}

var file_grpc_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_user_proto_goTypes = []interface{}{
	(*NewUser)(nil),        // 0: grpc.NewUser
	(*User)(nil),           // 1: grpc.User
	(*GetUsersParams)(nil), // 2: grpc.GetUsersParams
	(*UserList)(nil),       // 3: grpc.UserList
}
var file_grpc_user_proto_depIdxs = []int32{
	1, // 0: grpc.UserList.users:type_name -> grpc.User
	0, // 1: grpc.UserService.CreateUser:input_type -> grpc.NewUser
	2, // 2: grpc.UserService.GetUsers:input_type -> grpc.GetUsersParams
	1, // 3: grpc.UserService.CreateUser:output_type -> grpc.User
	3, // 4: grpc.UserService.GetUsers:output_type -> grpc.UserList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_user_proto_init() }
func file_grpc_user_proto_init() {
	if File_grpc_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
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
		file_grpc_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_grpc_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersParams); i {
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
		file_grpc_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
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
			RawDescriptor: file_grpc_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_user_proto_goTypes,
		DependencyIndexes: file_grpc_user_proto_depIdxs,
		MessageInfos:      file_grpc_user_proto_msgTypes,
	}.Build()
	File_grpc_user_proto = out.File
	file_grpc_user_proto_rawDesc = nil
	file_grpc_user_proto_goTypes = nil
	file_grpc_user_proto_depIdxs = nil
}

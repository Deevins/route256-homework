// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: api/proto/server/users.proto

package pb

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

type UserDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserDTO) Reset() {
	*x = UserDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDTO) ProtoMessage() {}

func (x *UserDTO) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDTO.ProtoReflect.Descriptor instead.
func (*UserDTO) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{0}
}

func (x *UserDTO) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserDTO) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateUserRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateUserRequestV1) Reset() {
	*x = CreateUserRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequestV1) ProtoMessage() {}

func (x *CreateUserRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequestV1.ProtoReflect.Descriptor instead.
func (*CreateUserRequestV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequestV1) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequestV1) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateUserResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateUserResponseV1) Reset() {
	*x = CreateUserResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponseV1) ProtoMessage() {}

func (x *CreateUserResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponseV1.ProtoReflect.Descriptor instead.
func (*CreateUserResponseV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponseV1) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListUserRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUserRequestV1) Reset() {
	*x = ListUserRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequestV1) ProtoMessage() {}

func (x *ListUserRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequestV1.ProtoReflect.Descriptor instead.
func (*ListUserRequestV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{3}
}

type ListUserResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserDTO `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *ListUserResponseV1) Reset() {
	*x = ListUserResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponseV1) ProtoMessage() {}

func (x *ListUserResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponseV1.ProtoReflect.Descriptor instead.
func (*ListUserResponseV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserResponseV1) GetUsers() []*UserDTO {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserRequestV1) Reset() {
	*x = GetUserRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequestV1) ProtoMessage() {}

func (x *GetUserRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequestV1.ProtoReflect.Descriptor instead.
func (*GetUserRequestV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserRequestV1) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserDTO `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponseV1) Reset() {
	*x = GetUserResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponseV1) ProtoMessage() {}

func (x *GetUserResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponseV1.ProtoReflect.Descriptor instead.
func (*GetUserResponseV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserResponseV1) GetUser() *UserDTO {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateEmailRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UpdateEmailRequestV1) Reset() {
	*x = UpdateEmailRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailRequestV1) ProtoMessage() {}

func (x *UpdateEmailRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailRequestV1.ProtoReflect.Descriptor instead.
func (*UpdateEmailRequestV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEmailRequestV1) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateEmailRequestV1) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdateEmailResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk bool `protobuf:"varint,1,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
}

func (x *UpdateEmailResponseV1) Reset() {
	*x = UpdateEmailResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailResponseV1) ProtoMessage() {}

func (x *UpdateEmailResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailResponseV1.ProtoReflect.Descriptor instead.
func (*UpdateEmailResponseV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateEmailResponseV1) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

type DeleteUserRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserRequestV1) Reset() {
	*x = DeleteUserRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequestV1) ProtoMessage() {}

func (x *DeleteUserRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequestV1.ProtoReflect.Descriptor instead.
func (*DeleteUserRequestV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserRequestV1) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteUserResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk bool `protobuf:"varint,1,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
}

func (x *DeleteUserResponseV1) Reset() {
	*x = DeleteUserResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_users_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponseV1) ProtoMessage() {}

func (x *DeleteUserResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_users_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponseV1.ProtoReflect.Descriptor instead.
func (*DeleteUserResponseV1) Descriptor() ([]byte, []int) {
	return file_api_proto_server_users_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteUserResponseV1) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

var File_api_proto_server_users_proto protoreflect.FileDescriptor

var file_api_proto_server_users_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x54, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x54, 0x4f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x47, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x22, 0x3c, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x31, 0x12, 0x26, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x54, 0x4f, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x24, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x45, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2c, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x31, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x22, 0x2e, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12,
	0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x69, 0x73, 0x4f, 0x6b, 0x32, 0x84, 0x03, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x4b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x31, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56,
	0x31, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x1b, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31,
	0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x1e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x00, 0x32, 0x0f, 0x0a, 0x0d, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_server_users_proto_rawDescOnce sync.Once
	file_api_proto_server_users_proto_rawDescData = file_api_proto_server_users_proto_rawDesc
)

func file_api_proto_server_users_proto_rawDescGZIP() []byte {
	file_api_proto_server_users_proto_rawDescOnce.Do(func() {
		file_api_proto_server_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_server_users_proto_rawDescData)
	})
	return file_api_proto_server_users_proto_rawDescData
}

var file_api_proto_server_users_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_proto_server_users_proto_goTypes = []interface{}{
	(*UserDTO)(nil),               // 0: user.v1.UserDTO
	(*CreateUserRequestV1)(nil),   // 1: user.v1.CreateUserRequestV1
	(*CreateUserResponseV1)(nil),  // 2: user.v1.CreateUserResponseV1
	(*ListUserRequestV1)(nil),     // 3: user.v1.ListUserRequestV1
	(*ListUserResponseV1)(nil),    // 4: user.v1.ListUserResponseV1
	(*GetUserRequestV1)(nil),      // 5: user.v1.GetUserRequestV1
	(*GetUserResponseV1)(nil),     // 6: user.v1.GetUserResponseV1
	(*UpdateEmailRequestV1)(nil),  // 7: user.v1.UpdateEmailRequestV1
	(*UpdateEmailResponseV1)(nil), // 8: user.v1.UpdateEmailResponseV1
	(*DeleteUserRequestV1)(nil),   // 9: user.v1.DeleteUserRequestV1
	(*DeleteUserResponseV1)(nil),  // 10: user.v1.DeleteUserResponseV1
}
var file_api_proto_server_users_proto_depIdxs = []int32{
	0,  // 0: user.v1.ListUserResponseV1.Users:type_name -> user.v1.UserDTO
	0,  // 1: user.v1.GetUserResponseV1.user:type_name -> user.v1.UserDTO
	1,  // 2: user.v1.UserServiceV1.CreateUser:input_type -> user.v1.CreateUserRequestV1
	3,  // 3: user.v1.UserServiceV1.ListUser:input_type -> user.v1.ListUserRequestV1
	5,  // 4: user.v1.UserServiceV1.GetUser:input_type -> user.v1.GetUserRequestV1
	7,  // 5: user.v1.UserServiceV1.UpdateEmail:input_type -> user.v1.UpdateEmailRequestV1
	9,  // 6: user.v1.UserServiceV1.DeleteUser:input_type -> user.v1.DeleteUserRequestV1
	2,  // 7: user.v1.UserServiceV1.CreateUser:output_type -> user.v1.CreateUserResponseV1
	4,  // 8: user.v1.UserServiceV1.ListUser:output_type -> user.v1.ListUserResponseV1
	6,  // 9: user.v1.UserServiceV1.GetUser:output_type -> user.v1.GetUserResponseV1
	8,  // 10: user.v1.UserServiceV1.UpdateEmail:output_type -> user.v1.UpdateEmailResponseV1
	10, // 11: user.v1.UserServiceV1.DeleteUser:output_type -> user.v1.DeleteUserResponseV1
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_server_users_proto_init() }
func file_api_proto_server_users_proto_init() {
	if File_api_proto_server_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_server_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDTO); i {
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
		file_api_proto_server_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequestV1); i {
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
		file_api_proto_server_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponseV1); i {
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
		file_api_proto_server_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequestV1); i {
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
		file_api_proto_server_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserResponseV1); i {
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
		file_api_proto_server_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequestV1); i {
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
		file_api_proto_server_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponseV1); i {
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
		file_api_proto_server_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailRequestV1); i {
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
		file_api_proto_server_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailResponseV1); i {
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
		file_api_proto_server_users_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequestV1); i {
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
		file_api_proto_server_users_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponseV1); i {
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
			RawDescriptor: file_api_proto_server_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_server_users_proto_goTypes,
		DependencyIndexes: file_api_proto_server_users_proto_depIdxs,
		MessageInfos:      file_api_proto_server_users_proto_msgTypes,
	}.Build()
	File_api_proto_server_users_proto = out.File
	file_api_proto_server_users_proto_rawDesc = nil
	file_api_proto_server_users_proto_goTypes = nil
	file_api_proto_server_users_proto_depIdxs = nil
}
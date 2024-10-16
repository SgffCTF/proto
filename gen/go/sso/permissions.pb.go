// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: sso/permissions.proto

package ssov1

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

type CanReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DutyId int64  `protobuf:"varint,1,opt,name=duty_id,json=dutyId,proto3" json:"duty_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CanReadRequest) Reset() {
	*x = CanReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanReadRequest) ProtoMessage() {}

func (x *CanReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanReadRequest.ProtoReflect.Descriptor instead.
func (*CanReadRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *CanReadRequest) GetDutyId() int64 {
	if x != nil {
		return x.DutyId
	}
	return 0
}

func (x *CanReadRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CanReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CanReadResponse) Reset() {
	*x = CanReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanReadResponse) ProtoMessage() {}

func (x *CanReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanReadResponse.ProtoReflect.Descriptor instead.
func (*CanReadResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *CanReadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CanCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CanCreateRequest) Reset() {
	*x = CanCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanCreateRequest) ProtoMessage() {}

func (x *CanCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanCreateRequest.ProtoReflect.Descriptor instead.
func (*CanCreateRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *CanCreateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CanCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CanCreateResponse) Reset() {
	*x = CanCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanCreateResponse) ProtoMessage() {}

func (x *CanCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanCreateResponse.ProtoReflect.Descriptor instead.
func (*CanCreateResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *CanCreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CanUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CanUpdateRequest) Reset() {
	*x = CanUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanUpdateRequest) ProtoMessage() {}

func (x *CanUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanUpdateRequest.ProtoReflect.Descriptor instead.
func (*CanUpdateRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{4}
}

func (x *CanUpdateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CanUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CanUpdateResponse) Reset() {
	*x = CanUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanUpdateResponse) ProtoMessage() {}

func (x *CanUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanUpdateResponse.ProtoReflect.Descriptor instead.
func (*CanUpdateResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{5}
}

func (x *CanUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CanDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DutyId int64  `protobuf:"varint,1,opt,name=duty_id,json=dutyId,proto3" json:"duty_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CanDeleteRequest) Reset() {
	*x = CanDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanDeleteRequest) ProtoMessage() {}

func (x *CanDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanDeleteRequest.ProtoReflect.Descriptor instead.
func (*CanDeleteRequest) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{6}
}

func (x *CanDeleteRequest) GetDutyId() int64 {
	if x != nil {
		return x.DutyId
	}
	return 0
}

func (x *CanDeleteRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CanDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CanDeleteResponse) Reset() {
	*x = CanDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_permissions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanDeleteResponse) ProtoMessage() {}

func (x *CanDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_permissions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanDeleteResponse.ProtoReflect.Descriptor instead.
func (*CanDeleteResponse) Descriptor() ([]byte, []int) {
	return file_sso_permissions_proto_rawDescGZIP(), []int{7}
}

func (x *CanDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_sso_permissions_proto protoreflect.FileDescriptor

var file_sso_permissions_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x73, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3f, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x75, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x75, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x11,
	0x43, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x43,
	0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x41, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x75, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x75, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xbb, 0x02, 0x0a, 0x0f, 0x44, 0x75, 0x74, 0x79, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x07, 0x43, 0x61,
	0x6e, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09,
	0x43, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x63, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_permissions_proto_rawDescOnce sync.Once
	file_sso_permissions_proto_rawDescData = file_sso_permissions_proto_rawDesc
)

func file_sso_permissions_proto_rawDescGZIP() []byte {
	file_sso_permissions_proto_rawDescOnce.Do(func() {
		file_sso_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_permissions_proto_rawDescData)
	})
	return file_sso_permissions_proto_rawDescData
}

var file_sso_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sso_permissions_proto_goTypes = []any{
	(*CanReadRequest)(nil),    // 0: permissions.CanReadRequest
	(*CanReadResponse)(nil),   // 1: permissions.CanReadResponse
	(*CanCreateRequest)(nil),  // 2: permissions.CanCreateRequest
	(*CanCreateResponse)(nil), // 3: permissions.CanCreateResponse
	(*CanUpdateRequest)(nil),  // 4: permissions.CanUpdateRequest
	(*CanUpdateResponse)(nil), // 5: permissions.CanUpdateResponse
	(*CanDeleteRequest)(nil),  // 6: permissions.CanDeleteRequest
	(*CanDeleteResponse)(nil), // 7: permissions.CanDeleteResponse
}
var file_sso_permissions_proto_depIdxs = []int32{
	0, // 0: permissions.DutyPermissions.CanRead:input_type -> permissions.CanReadRequest
	2, // 1: permissions.DutyPermissions.CanCreate:input_type -> permissions.CanCreateRequest
	4, // 2: permissions.DutyPermissions.CanUpdate:input_type -> permissions.CanUpdateRequest
	6, // 3: permissions.DutyPermissions.CanDelete:input_type -> permissions.CanDeleteRequest
	1, // 4: permissions.DutyPermissions.CanRead:output_type -> permissions.CanReadResponse
	3, // 5: permissions.DutyPermissions.CanCreate:output_type -> permissions.CanCreateResponse
	5, // 6: permissions.DutyPermissions.CanUpdate:output_type -> permissions.CanUpdateResponse
	7, // 7: permissions.DutyPermissions.CanDelete:output_type -> permissions.CanDeleteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_permissions_proto_init() }
func file_sso_permissions_proto_init() {
	if File_sso_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_permissions_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CanReadRequest); i {
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
		file_sso_permissions_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CanReadResponse); i {
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
		file_sso_permissions_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CanCreateRequest); i {
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
		file_sso_permissions_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CanCreateResponse); i {
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
		file_sso_permissions_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CanUpdateRequest); i {
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
		file_sso_permissions_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CanUpdateResponse); i {
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
		file_sso_permissions_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CanDeleteRequest); i {
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
		file_sso_permissions_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CanDeleteResponse); i {
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
			RawDescriptor: file_sso_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_permissions_proto_goTypes,
		DependencyIndexes: file_sso_permissions_proto_depIdxs,
		MessageInfos:      file_sso_permissions_proto_msgTypes,
	}.Build()
	File_sso_permissions_proto = out.File
	file_sso_permissions_proto_rawDesc = nil
	file_sso_permissions_proto_goTypes = nil
	file_sso_permissions_proto_depIdxs = nil
}

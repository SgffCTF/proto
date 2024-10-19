// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: duties/duties.proto

package dutiesv1

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

type Duty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Duty) Reset() {
	*x = Duty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Duty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duty) ProtoMessage() {}

func (x *Duty) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duty.ProtoReflect.Descriptor instead.
func (*Duty) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{0}
}

func (x *Duty) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Duty) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Duty) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ReadDutyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadDutyRequest) Reset() {
	*x = ReadDutyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDutyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDutyRequest) ProtoMessage() {}

func (x *ReadDutyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDutyRequest.ProtoReflect.Descriptor instead.
func (*ReadDutyRequest) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{1}
}

func (x *ReadDutyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReadDutyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duty *Duty `protobuf:"bytes,1,opt,name=duty,proto3" json:"duty,omitempty"`
}

func (x *ReadDutyResponse) Reset() {
	*x = ReadDutyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDutyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDutyResponse) ProtoMessage() {}

func (x *ReadDutyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDutyResponse.ProtoReflect.Descriptor instead.
func (*ReadDutyResponse) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{2}
}

func (x *ReadDutyResponse) GetDuty() *Duty {
	if x != nil {
		return x.Duty
	}
	return nil
}

type ReadOwnerDutiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId int64 `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *ReadOwnerDutiesRequest) Reset() {
	*x = ReadOwnerDutiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOwnerDutiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOwnerDutiesRequest) ProtoMessage() {}

func (x *ReadOwnerDutiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOwnerDutiesRequest.ProtoReflect.Descriptor instead.
func (*ReadOwnerDutiesRequest) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{3}
}

func (x *ReadOwnerDutiesRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type ReadTargetDutiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId int64 `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *ReadTargetDutiesRequest) Reset() {
	*x = ReadTargetDutiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadTargetDutiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTargetDutiesRequest) ProtoMessage() {}

func (x *ReadTargetDutiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTargetDutiesRequest.ProtoReflect.Descriptor instead.
func (*ReadTargetDutiesRequest) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{4}
}

func (x *ReadTargetDutiesRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type ReadDutiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duties []*Duty `protobuf:"bytes,1,rep,name=duties,proto3" json:"duties,omitempty"`
}

func (x *ReadDutiesResponse) Reset() {
	*x = ReadDutiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDutiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDutiesResponse) ProtoMessage() {}

func (x *ReadDutiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDutiesResponse.ProtoReflect.Descriptor instead.
func (*ReadDutiesResponse) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{5}
}

func (x *ReadDutiesResponse) GetDuties() []*Duty {
	if x != nil {
		return x.Duties
	}
	return nil
}

type CreateDutyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	OwnerId     int64  `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	TargetId    int64  `protobuf:"varint,4,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
}

func (x *CreateDutyRequest) Reset() {
	*x = CreateDutyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDutyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDutyRequest) ProtoMessage() {}

func (x *CreateDutyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDutyRequest.ProtoReflect.Descriptor instead.
func (*CreateDutyRequest) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDutyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateDutyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateDutyRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *CreateDutyRequest) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

type CreateDutyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDutyResponse) Reset() {
	*x = CreateDutyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDutyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDutyResponse) ProtoMessage() {}

func (x *CreateDutyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDutyResponse.ProtoReflect.Descriptor instead.
func (*CreateDutyResponse) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{7}
}

type UpdateDutyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duty *Duty `protobuf:"bytes,1,opt,name=duty,proto3" json:"duty,omitempty"`
}

func (x *UpdateDutyRequest) Reset() {
	*x = UpdateDutyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDutyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDutyRequest) ProtoMessage() {}

func (x *UpdateDutyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDutyRequest.ProtoReflect.Descriptor instead.
func (*UpdateDutyRequest) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateDutyRequest) GetDuty() *Duty {
	if x != nil {
		return x.Duty
	}
	return nil
}

type UpdateDutyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duty *Duty `protobuf:"bytes,1,opt,name=duty,proto3" json:"duty,omitempty"`
}

func (x *UpdateDutyResponse) Reset() {
	*x = UpdateDutyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDutyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDutyResponse) ProtoMessage() {}

func (x *UpdateDutyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDutyResponse.ProtoReflect.Descriptor instead.
func (*UpdateDutyResponse) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateDutyResponse) GetDuty() *Duty {
	if x != nil {
		return x.Duty
	}
	return nil
}

type DeleteDutyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDutyRequest) Reset() {
	*x = DeleteDutyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDutyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDutyRequest) ProtoMessage() {}

func (x *DeleteDutyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDutyRequest.ProtoReflect.Descriptor instead.
func (*DeleteDutyRequest) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteDutyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDutyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDutyResponse) Reset() {
	*x = DeleteDutyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duties_duties_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDutyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDutyResponse) ProtoMessage() {}

func (x *DeleteDutyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_duties_duties_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDutyResponse.ProtoReflect.Descriptor instead.
func (*DeleteDutyResponse) Descriptor() ([]byte, []int) {
	return file_duties_duties_proto_rawDescGZIP(), []int{11}
}

var File_duties_duties_proto protoreflect.FileDescriptor

var file_duties_duties_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x22, 0x4e, 0x0a,
	0x04, 0x44, 0x75, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x34, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x75, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x74, 0x79,
	0x52, 0x04, 0x64, 0x75, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x17, 0x52,
	0x65, 0x61, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x3a, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x64, 0x75, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x44, 0x75, 0x74, 0x79, 0x52, 0x06, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x22, 0x83, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x04, 0x64, 0x75, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64,
	0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x74, 0x79, 0x52, 0x04, 0x64, 0x75, 0x74, 0x79,
	0x22, 0x36, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x75, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x75,
	0x74, 0x79, 0x52, 0x04, 0x64, 0x75, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xb6, 0x03, 0x0a, 0x06, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3d,
	0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x44, 0x75, 0x74, 0x79, 0x12, 0x17, 0x2e, 0x64, 0x75, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x1e, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x44, 0x75,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x10,
	0x52, 0x65, 0x61, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x1f, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x44, 0x75, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x44,
	0x75, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x64, 0x75,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79,
	0x12, 0x19, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x75,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x75, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x75, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e,
	0x63, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x75, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x64, 0x75, 0x74, 0x69, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_duties_duties_proto_rawDescOnce sync.Once
	file_duties_duties_proto_rawDescData = file_duties_duties_proto_rawDesc
)

func file_duties_duties_proto_rawDescGZIP() []byte {
	file_duties_duties_proto_rawDescOnce.Do(func() {
		file_duties_duties_proto_rawDescData = protoimpl.X.CompressGZIP(file_duties_duties_proto_rawDescData)
	})
	return file_duties_duties_proto_rawDescData
}

var file_duties_duties_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_duties_duties_proto_goTypes = []any{
	(*Duty)(nil),                    // 0: duties.Duty
	(*ReadDutyRequest)(nil),         // 1: duties.ReadDutyRequest
	(*ReadDutyResponse)(nil),        // 2: duties.ReadDutyResponse
	(*ReadOwnerDutiesRequest)(nil),  // 3: duties.ReadOwnerDutiesRequest
	(*ReadTargetDutiesRequest)(nil), // 4: duties.ReadTargetDutiesRequest
	(*ReadDutiesResponse)(nil),      // 5: duties.ReadDutiesResponse
	(*CreateDutyRequest)(nil),       // 6: duties.CreateDutyRequest
	(*CreateDutyResponse)(nil),      // 7: duties.CreateDutyResponse
	(*UpdateDutyRequest)(nil),       // 8: duties.UpdateDutyRequest
	(*UpdateDutyResponse)(nil),      // 9: duties.UpdateDutyResponse
	(*DeleteDutyRequest)(nil),       // 10: duties.DeleteDutyRequest
	(*DeleteDutyResponse)(nil),      // 11: duties.DeleteDutyResponse
}
var file_duties_duties_proto_depIdxs = []int32{
	0,  // 0: duties.ReadDutyResponse.duty:type_name -> duties.Duty
	0,  // 1: duties.ReadDutiesResponse.duties:type_name -> duties.Duty
	0,  // 2: duties.UpdateDutyRequest.duty:type_name -> duties.Duty
	0,  // 3: duties.UpdateDutyResponse.duty:type_name -> duties.Duty
	1,  // 4: duties.Duties.ReadDuty:input_type -> duties.ReadDutyRequest
	3,  // 5: duties.Duties.ReadOwnerDuties:input_type -> duties.ReadOwnerDutiesRequest
	4,  // 6: duties.Duties.ReadTargetDuties:input_type -> duties.ReadTargetDutiesRequest
	6,  // 7: duties.Duties.CreateDuty:input_type -> duties.CreateDutyRequest
	8,  // 8: duties.Duties.UpdateDuty:input_type -> duties.UpdateDutyRequest
	10, // 9: duties.Duties.DeleteDuty:input_type -> duties.DeleteDutyRequest
	2,  // 10: duties.Duties.ReadDuty:output_type -> duties.ReadDutyResponse
	5,  // 11: duties.Duties.ReadOwnerDuties:output_type -> duties.ReadDutiesResponse
	5,  // 12: duties.Duties.ReadTargetDuties:output_type -> duties.ReadDutiesResponse
	7,  // 13: duties.Duties.CreateDuty:output_type -> duties.CreateDutyResponse
	9,  // 14: duties.Duties.UpdateDuty:output_type -> duties.UpdateDutyResponse
	11, // 15: duties.Duties.DeleteDuty:output_type -> duties.DeleteDutyResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_duties_duties_proto_init() }
func file_duties_duties_proto_init() {
	if File_duties_duties_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_duties_duties_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Duty); i {
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
		file_duties_duties_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReadDutyRequest); i {
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
		file_duties_duties_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReadDutyResponse); i {
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
		file_duties_duties_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReadOwnerDutiesRequest); i {
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
		file_duties_duties_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReadTargetDutiesRequest); i {
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
		file_duties_duties_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ReadDutiesResponse); i {
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
		file_duties_duties_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDutyRequest); i {
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
		file_duties_duties_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDutyResponse); i {
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
		file_duties_duties_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateDutyRequest); i {
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
		file_duties_duties_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateDutyResponse); i {
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
		file_duties_duties_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteDutyRequest); i {
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
		file_duties_duties_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteDutyResponse); i {
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
			RawDescriptor: file_duties_duties_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_duties_duties_proto_goTypes,
		DependencyIndexes: file_duties_duties_proto_depIdxs,
		MessageInfos:      file_duties_duties_proto_msgTypes,
	}.Build()
	File_duties_duties_proto = out.File
	file_duties_duties_proto_rawDesc = nil
	file_duties_duties_proto_goTypes = nil
	file_duties_duties_proto_depIdxs = nil
}

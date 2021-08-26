// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: ova-entertainment-api/api.proto

package ova_entertainment_api

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

type CreateEntertainmentV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateEntertainmentV1Request) Reset() {
	*x = CreateEntertainmentV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_entertainment_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntertainmentV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntertainmentV1Request) ProtoMessage() {}

func (x *CreateEntertainmentV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_entertainment_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntertainmentV1Request.ProtoReflect.Descriptor instead.
func (*CreateEntertainmentV1Request) Descriptor() ([]byte, []int) {
	return file_ova_entertainment_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEntertainmentV1Request) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateEntertainmentV1Request) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEntertainmentV1Request) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DescribeEntertainmentV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID      uint64 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DescribeEntertainmentV1Request) Reset() {
	*x = DescribeEntertainmentV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_entertainment_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeEntertainmentV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEntertainmentV1Request) ProtoMessage() {}

func (x *DescribeEntertainmentV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_entertainment_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEntertainmentV1Request.ProtoReflect.Descriptor instead.
func (*DescribeEntertainmentV1Request) Descriptor() ([]byte, []int) {
	return file_ova_entertainment_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeEntertainmentV1Request) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DescribeEntertainmentV1Request) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DescribeEntertainmentV1Request) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DescribeEntertainmentV1Request) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ListEntertainmentV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListEntertainmentV1Request) Reset() {
	*x = ListEntertainmentV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_entertainment_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntertainmentV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntertainmentV1Request) ProtoMessage() {}

func (x *ListEntertainmentV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_entertainment_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntertainmentV1Request.ProtoReflect.Descriptor instead.
func (*ListEntertainmentV1Request) Descriptor() ([]byte, []int) {
	return file_ova_entertainment_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListEntertainmentV1Request) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListEntertainmentV1Request) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type RemoveEntertainmentV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RemoveEntertainmentV1Request) Reset() {
	*x = RemoveEntertainmentV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_entertainment_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveEntertainmentV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveEntertainmentV1Request) ProtoMessage() {}

func (x *RemoveEntertainmentV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_entertainment_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveEntertainmentV1Request.ProtoReflect.Descriptor instead.
func (*RemoveEntertainmentV1Request) Descriptor() ([]byte, []int) {
	return file_ova_entertainment_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveEntertainmentV1Request) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ListEntertainmentsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*EntertainmentV1Response `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListEntertainmentsV1Response) Reset() {
	*x = ListEntertainmentsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_entertainment_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntertainmentsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntertainmentsV1Response) ProtoMessage() {}

func (x *ListEntertainmentsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ova_entertainment_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntertainmentsV1Response.ProtoReflect.Descriptor instead.
func (*ListEntertainmentsV1Response) Descriptor() ([]byte, []int) {
	return file_ova_entertainment_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListEntertainmentsV1Response) GetList() []*EntertainmentV1Response {
	if x != nil {
		return x.List
	}
	return nil
}

type EntertainmentV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID      uint64 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Date        string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *EntertainmentV1Response) Reset() {
	*x = EntertainmentV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_entertainment_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntertainmentV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntertainmentV1Response) ProtoMessage() {}

func (x *EntertainmentV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ova_entertainment_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntertainmentV1Response.ProtoReflect.Descriptor instead.
func (*EntertainmentV1Response) Descriptor() ([]byte, []int) {
	return file_ova_entertainment_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *EntertainmentV1Response) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *EntertainmentV1Response) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *EntertainmentV1Response) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EntertainmentV1Response) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EntertainmentV1Response) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type RemoveV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RemoveV1Response) Reset() {
	*x = RemoveV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_entertainment_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveV1Response) ProtoMessage() {}

func (x *RemoveV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ova_entertainment_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveV1Response.ProtoReflect.Descriptor instead.
func (*RemoveV1Response) Descriptor() ([]byte, []int) {
	return file_ova_entertainment_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveV1Response) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_ova_entertainment_api_api_proto protoreflect.FileDescriptor

var file_ova_entertainment_api_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x76, 0x61, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x6e, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x1e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x1a, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x2e, 0x0a, 0x1c, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x62, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x17,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x86, 0x04, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x12,
	0x7e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61,
	0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x33, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x82, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x35, 0x2e, 0x6f, 0x76,
	0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61,
	0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x12, 0x31, 0x2e,
	0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74,
	0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31,
	0x12, 0x33, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74,
	0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6f, 0x76, 0x61, 0x2d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x76, 0x61, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x74,
	0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ova_entertainment_api_api_proto_rawDescOnce sync.Once
	file_ova_entertainment_api_api_proto_rawDescData = file_ova_entertainment_api_api_proto_rawDesc
)

func file_ova_entertainment_api_api_proto_rawDescGZIP() []byte {
	file_ova_entertainment_api_api_proto_rawDescOnce.Do(func() {
		file_ova_entertainment_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ova_entertainment_api_api_proto_rawDescData)
	})
	return file_ova_entertainment_api_api_proto_rawDescData
}

var file_ova_entertainment_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ova_entertainment_api_api_proto_goTypes = []interface{}{
	(*CreateEntertainmentV1Request)(nil),   // 0: ova.entertainment.api.CreateEntertainmentV1Request
	(*DescribeEntertainmentV1Request)(nil), // 1: ova.entertainment.api.DescribeEntertainmentV1Request
	(*ListEntertainmentV1Request)(nil),     // 2: ova.entertainment.api.ListEntertainmentV1Request
	(*RemoveEntertainmentV1Request)(nil),   // 3: ova.entertainment.api.RemoveEntertainmentV1Request
	(*ListEntertainmentsV1Response)(nil),   // 4: ova.entertainment.api.ListEntertainmentsV1Response
	(*EntertainmentV1Response)(nil),        // 5: ova.entertainment.api.EntertainmentV1Response
	(*RemoveV1Response)(nil),               // 6: ova.entertainment.api.RemoveV1Response
}
var file_ova_entertainment_api_api_proto_depIdxs = []int32{
	5, // 0: ova.entertainment.api.ListEntertainmentsV1Response.list:type_name -> ova.entertainment.api.EntertainmentV1Response
	0, // 1: ova.entertainment.api.api.CreateEntertainmentV1:input_type -> ova.entertainment.api.CreateEntertainmentV1Request
	1, // 2: ova.entertainment.api.api.DescribeEntertainmentV1:input_type -> ova.entertainment.api.DescribeEntertainmentV1Request
	2, // 3: ova.entertainment.api.api.ListEntertainmentsV1:input_type -> ova.entertainment.api.ListEntertainmentV1Request
	3, // 4: ova.entertainment.api.api.RemoveEntertainmentV1:input_type -> ova.entertainment.api.RemoveEntertainmentV1Request
	5, // 5: ova.entertainment.api.api.CreateEntertainmentV1:output_type -> ova.entertainment.api.EntertainmentV1Response
	5, // 6: ova.entertainment.api.api.DescribeEntertainmentV1:output_type -> ova.entertainment.api.EntertainmentV1Response
	4, // 7: ova.entertainment.api.api.ListEntertainmentsV1:output_type -> ova.entertainment.api.ListEntertainmentsV1Response
	6, // 8: ova.entertainment.api.api.RemoveEntertainmentV1:output_type -> ova.entertainment.api.RemoveV1Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ova_entertainment_api_api_proto_init() }
func file_ova_entertainment_api_api_proto_init() {
	if File_ova_entertainment_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ova_entertainment_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntertainmentV1Request); i {
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
		file_ova_entertainment_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeEntertainmentV1Request); i {
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
		file_ova_entertainment_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntertainmentV1Request); i {
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
		file_ova_entertainment_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveEntertainmentV1Request); i {
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
		file_ova_entertainment_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntertainmentsV1Response); i {
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
		file_ova_entertainment_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntertainmentV1Response); i {
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
		file_ova_entertainment_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveV1Response); i {
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
			RawDescriptor: file_ova_entertainment_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ova_entertainment_api_api_proto_goTypes,
		DependencyIndexes: file_ova_entertainment_api_api_proto_depIdxs,
		MessageInfos:      file_ova_entertainment_api_api_proto_msgTypes,
	}.Build()
	File_ova_entertainment_api_api_proto = out.File
	file_ova_entertainment_api_api_proto_rawDesc = nil
	file_ova_entertainment_api_api_proto_goTypes = nil
	file_ova_entertainment_api_api_proto_depIdxs = nil
}

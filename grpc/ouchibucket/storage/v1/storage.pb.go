// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: ouchibucket/storage/v1/storage.proto

package storagev1

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

type GetStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *GetStorageRequest) Reset() {
	*x = GetStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageRequest) ProtoMessage() {}

func (x *GetStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageRequest.ProtoReflect.Descriptor instead.
func (*GetStorageRequest) Descriptor() ([]byte, []int) {
	return file_ouchibucket_storage_v1_storage_proto_rawDescGZIP(), []int{0}
}

func (x *GetStorageRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetStorageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage *Storage `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *GetStorageResponse) Reset() {
	*x = GetStorageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageResponse) ProtoMessage() {}

func (x *GetStorageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageResponse.ProtoReflect.Descriptor instead.
func (*GetStorageResponse) Descriptor() ([]byte, []int) {
	return file_ouchibucket_storage_v1_storage_proto_rawDescGZIP(), []int{1}
}

func (x *GetStorageResponse) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

type CreateStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *CreateStorageRequest) Reset() {
	*x = CreateStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStorageRequest) ProtoMessage() {}

func (x *CreateStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStorageRequest.ProtoReflect.Descriptor instead.
func (*CreateStorageRequest) Descriptor() ([]byte, []int) {
	return file_ouchibucket_storage_v1_storage_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStorageRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type CreateStorageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateStorageResponse) Reset() {
	*x = CreateStorageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStorageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStorageResponse) ProtoMessage() {}

func (x *CreateStorageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStorageResponse.ProtoReflect.Descriptor instead.
func (*CreateStorageResponse) Descriptor() ([]byte, []int) {
	return file_ouchibucket_storage_v1_storage_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStorageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *DeleteStorageRequest) Reset() {
	*x = DeleteStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStorageRequest) ProtoMessage() {}

func (x *DeleteStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStorageRequest.ProtoReflect.Descriptor instead.
func (*DeleteStorageRequest) Descriptor() ([]byte, []int) {
	return file_ouchibucket_storage_v1_storage_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteStorageRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type DeleteStorageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStorageResponse) Reset() {
	*x = DeleteStorageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStorageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStorageResponse) ProtoMessage() {}

func (x *DeleteStorageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStorageResponse.ProtoReflect.Descriptor instead.
func (*DeleteStorageResponse) Descriptor() ([]byte, []int) {
	return file_ouchibucket_storage_v1_storage_proto_rawDescGZIP(), []int{5}
}

type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_storage_v1_storage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_ouchibucket_storage_v1_storage_proto_rawDescGZIP(), []int{6}
}

func (x *Storage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Storage) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

var File_ouchibucket_storage_v1_storage_proto protoreflect.FileDescriptor

var file_ouchibucket_storage_v1_storage_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2e,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22,
	0x31, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x32, 0xd7, 0x02,
	0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x29,
	0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x75, 0x63, 0x68,
	0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xf6, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x69, 0x6d, 0x6f, 0x74, 0x6f, 0x2d, 0x6a, 0x75, 0x6e, 0x79, 0x61,
	0x2f, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x2d, 0x68, 0x75, 0x62, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x53, 0x58, 0xaa, 0x02,
	0x16, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x22, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ouchibucket_storage_v1_storage_proto_rawDescOnce sync.Once
	file_ouchibucket_storage_v1_storage_proto_rawDescData = file_ouchibucket_storage_v1_storage_proto_rawDesc
)

func file_ouchibucket_storage_v1_storage_proto_rawDescGZIP() []byte {
	file_ouchibucket_storage_v1_storage_proto_rawDescOnce.Do(func() {
		file_ouchibucket_storage_v1_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_ouchibucket_storage_v1_storage_proto_rawDescData)
	})
	return file_ouchibucket_storage_v1_storage_proto_rawDescData
}

var file_ouchibucket_storage_v1_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ouchibucket_storage_v1_storage_proto_goTypes = []interface{}{
	(*GetStorageRequest)(nil),     // 0: ouchibucket.storage.v1.GetStorageRequest
	(*GetStorageResponse)(nil),    // 1: ouchibucket.storage.v1.GetStorageResponse
	(*CreateStorageRequest)(nil),  // 2: ouchibucket.storage.v1.CreateStorageRequest
	(*CreateStorageResponse)(nil), // 3: ouchibucket.storage.v1.CreateStorageResponse
	(*DeleteStorageRequest)(nil),  // 4: ouchibucket.storage.v1.DeleteStorageRequest
	(*DeleteStorageResponse)(nil), // 5: ouchibucket.storage.v1.DeleteStorageResponse
	(*Storage)(nil),               // 6: ouchibucket.storage.v1.Storage
}
var file_ouchibucket_storage_v1_storage_proto_depIdxs = []int32{
	6, // 0: ouchibucket.storage.v1.GetStorageResponse.storage:type_name -> ouchibucket.storage.v1.Storage
	0, // 1: ouchibucket.storage.v1.StorageService.GetStorage:input_type -> ouchibucket.storage.v1.GetStorageRequest
	2, // 2: ouchibucket.storage.v1.StorageService.CreateStorage:input_type -> ouchibucket.storage.v1.CreateStorageRequest
	4, // 3: ouchibucket.storage.v1.StorageService.DeleteStorage:input_type -> ouchibucket.storage.v1.DeleteStorageRequest
	1, // 4: ouchibucket.storage.v1.StorageService.GetStorage:output_type -> ouchibucket.storage.v1.GetStorageResponse
	3, // 5: ouchibucket.storage.v1.StorageService.CreateStorage:output_type -> ouchibucket.storage.v1.CreateStorageResponse
	5, // 6: ouchibucket.storage.v1.StorageService.DeleteStorage:output_type -> ouchibucket.storage.v1.DeleteStorageResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ouchibucket_storage_v1_storage_proto_init() }
func file_ouchibucket_storage_v1_storage_proto_init() {
	if File_ouchibucket_storage_v1_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ouchibucket_storage_v1_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageRequest); i {
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
		file_ouchibucket_storage_v1_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageResponse); i {
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
		file_ouchibucket_storage_v1_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStorageRequest); i {
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
		file_ouchibucket_storage_v1_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStorageResponse); i {
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
		file_ouchibucket_storage_v1_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStorageRequest); i {
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
		file_ouchibucket_storage_v1_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStorageResponse); i {
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
		file_ouchibucket_storage_v1_storage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
			RawDescriptor: file_ouchibucket_storage_v1_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ouchibucket_storage_v1_storage_proto_goTypes,
		DependencyIndexes: file_ouchibucket_storage_v1_storage_proto_depIdxs,
		MessageInfos:      file_ouchibucket_storage_v1_storage_proto_msgTypes,
	}.Build()
	File_ouchibucket_storage_v1_storage_proto = out.File
	file_ouchibucket_storage_v1_storage_proto_rawDesc = nil
	file_ouchibucket_storage_v1_storage_proto_goTypes = nil
	file_ouchibucket_storage_v1_storage_proto_depIdxs = nil
}

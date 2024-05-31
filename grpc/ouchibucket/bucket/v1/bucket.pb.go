// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: ouchibucket/bucket/v1/bucket.proto

package bucketv1

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

type CreateBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StorageId string `protobuf:"bytes,2,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
}

func (x *CreateBucketRequest) Reset() {
	*x = CreateBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketRequest) ProtoMessage() {}

func (x *CreateBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketRequest.ProtoReflect.Descriptor instead.
func (*CreateBucketRequest) Descriptor() ([]byte, []int) {
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBucketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBucketRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

type CreateBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateBucketResponse) Reset() {
	*x = CreateBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketResponse) ProtoMessage() {}

func (x *CreateBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketResponse.ProtoReflect.Descriptor instead.
func (*CreateBucketResponse) Descriptor() ([]byte, []int) {
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBucketResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListBucketsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageId string `protobuf:"bytes,1,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
}

func (x *ListBucketsRequest) Reset() {
	*x = ListBucketsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBucketsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBucketsRequest) ProtoMessage() {}

func (x *ListBucketsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBucketsRequest.ProtoReflect.Descriptor instead.
func (*ListBucketsRequest) Descriptor() ([]byte, []int) {
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP(), []int{2}
}

func (x *ListBucketsRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

type ListBucketsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buckets []*Bucket `protobuf:"bytes,1,rep,name=buckets,proto3" json:"buckets,omitempty"`
}

func (x *ListBucketsResponse) Reset() {
	*x = ListBucketsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBucketsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBucketsResponse) ProtoMessage() {}

func (x *ListBucketsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBucketsResponse.ProtoReflect.Descriptor instead.
func (*ListBucketsResponse) Descriptor() ([]byte, []int) {
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP(), []int{3}
}

func (x *ListBucketsResponse) GetBuckets() []*Bucket {
	if x != nil {
		return x.Buckets
	}
	return nil
}

type SyncBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketId     string `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	RelativePath string `protobuf:"bytes,2,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
}

func (x *SyncBucketRequest) Reset() {
	*x = SyncBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncBucketRequest) ProtoMessage() {}

func (x *SyncBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncBucketRequest.ProtoReflect.Descriptor instead.
func (*SyncBucketRequest) Descriptor() ([]byte, []int) {
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP(), []int{4}
}

func (x *SyncBucketRequest) GetBucketId() string {
	if x != nil {
		return x.BucketId
	}
	return ""
}

func (x *SyncBucketRequest) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

type SyncBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncBucketResponse) Reset() {
	*x = SyncBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncBucketResponse) ProtoMessage() {}

func (x *SyncBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncBucketResponse.ProtoReflect.Descriptor instead.
func (*SyncBucketResponse) Descriptor() ([]byte, []int) {
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP(), []int{5}
}

type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RootDirectoryId string `protobuf:"bytes,3,opt,name=root_directory_id,json=rootDirectoryId,proto3" json:"root_directory_id,omitempty"`
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_ouchibucket_bucket_v1_bucket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP(), []int{6}
}

func (x *Bucket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bucket) GetRootDirectoryId() string {
	if x != nil {
		return x.RootDirectoryId
	}
	return ""
}

var File_ouchibucket_bucket_v1_bucket_proto protoreflect.FileDescriptor

var file_ouchibucket_bucket_v1_bucket_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x48, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x22, 0x4e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x75, 0x63,
	0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x22, 0x55, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x79, 0x6e,
	0x63, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x58, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x32, 0xc7, 0x02, 0x0a, 0x0d, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x6f, 0x75,
	0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x29, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63,
	0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x6f,
	0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0xee, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x75, 0x63, 0x68,
	0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x69,
	0x6d, 0x6f, 0x74, 0x6f, 0x2d, 0x6a, 0x75, 0x6e, 0x79, 0x61, 0x2f, 0x6f, 0x75, 0x63, 0x68, 0x69,
	0x2d, 0x68, 0x75, 0x62, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x4f, 0x42, 0x58, 0xaa, 0x02, 0x15, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x15, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x5c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x4f, 0x75, 0x63,
	0x68, 0x69, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ouchibucket_bucket_v1_bucket_proto_rawDescOnce sync.Once
	file_ouchibucket_bucket_v1_bucket_proto_rawDescData = file_ouchibucket_bucket_v1_bucket_proto_rawDesc
)

func file_ouchibucket_bucket_v1_bucket_proto_rawDescGZIP() []byte {
	file_ouchibucket_bucket_v1_bucket_proto_rawDescOnce.Do(func() {
		file_ouchibucket_bucket_v1_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ouchibucket_bucket_v1_bucket_proto_rawDescData)
	})
	return file_ouchibucket_bucket_v1_bucket_proto_rawDescData
}

var file_ouchibucket_bucket_v1_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ouchibucket_bucket_v1_bucket_proto_goTypes = []interface{}{
	(*CreateBucketRequest)(nil),  // 0: ouchibucket.bucket.v1.CreateBucketRequest
	(*CreateBucketResponse)(nil), // 1: ouchibucket.bucket.v1.CreateBucketResponse
	(*ListBucketsRequest)(nil),   // 2: ouchibucket.bucket.v1.ListBucketsRequest
	(*ListBucketsResponse)(nil),  // 3: ouchibucket.bucket.v1.ListBucketsResponse
	(*SyncBucketRequest)(nil),    // 4: ouchibucket.bucket.v1.SyncBucketRequest
	(*SyncBucketResponse)(nil),   // 5: ouchibucket.bucket.v1.SyncBucketResponse
	(*Bucket)(nil),               // 6: ouchibucket.bucket.v1.Bucket
}
var file_ouchibucket_bucket_v1_bucket_proto_depIdxs = []int32{
	6, // 0: ouchibucket.bucket.v1.ListBucketsResponse.buckets:type_name -> ouchibucket.bucket.v1.Bucket
	0, // 1: ouchibucket.bucket.v1.BucketService.CreateBucket:input_type -> ouchibucket.bucket.v1.CreateBucketRequest
	2, // 2: ouchibucket.bucket.v1.BucketService.ListBuckets:input_type -> ouchibucket.bucket.v1.ListBucketsRequest
	4, // 3: ouchibucket.bucket.v1.BucketService.SyncBucket:input_type -> ouchibucket.bucket.v1.SyncBucketRequest
	1, // 4: ouchibucket.bucket.v1.BucketService.CreateBucket:output_type -> ouchibucket.bucket.v1.CreateBucketResponse
	3, // 5: ouchibucket.bucket.v1.BucketService.ListBuckets:output_type -> ouchibucket.bucket.v1.ListBucketsResponse
	5, // 6: ouchibucket.bucket.v1.BucketService.SyncBucket:output_type -> ouchibucket.bucket.v1.SyncBucketResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ouchibucket_bucket_v1_bucket_proto_init() }
func file_ouchibucket_bucket_v1_bucket_proto_init() {
	if File_ouchibucket_bucket_v1_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ouchibucket_bucket_v1_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBucketRequest); i {
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
		file_ouchibucket_bucket_v1_bucket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBucketResponse); i {
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
		file_ouchibucket_bucket_v1_bucket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBucketsRequest); i {
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
		file_ouchibucket_bucket_v1_bucket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBucketsResponse); i {
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
		file_ouchibucket_bucket_v1_bucket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncBucketRequest); i {
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
		file_ouchibucket_bucket_v1_bucket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncBucketResponse); i {
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
		file_ouchibucket_bucket_v1_bucket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucket); i {
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
			RawDescriptor: file_ouchibucket_bucket_v1_bucket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ouchibucket_bucket_v1_bucket_proto_goTypes,
		DependencyIndexes: file_ouchibucket_bucket_v1_bucket_proto_depIdxs,
		MessageInfos:      file_ouchibucket_bucket_v1_bucket_proto_msgTypes,
	}.Build()
	File_ouchibucket_bucket_v1_bucket_proto = out.File
	file_ouchibucket_bucket_v1_bucket_proto_rawDesc = nil
	file_ouchibucket_bucket_v1_bucket_proto_goTypes = nil
	file_ouchibucket_bucket_v1_bucket_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: ouchihub/content/v1/content.proto

package contentv1

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

type CreateContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateContentRequest) Reset() {
	*x = CreateContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_content_v1_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentRequest) ProtoMessage() {}

func (x *CreateContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_content_v1_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentRequest.ProtoReflect.Descriptor instead.
func (*CreateContentRequest) Descriptor() ([]byte, []int) {
	return file_ouchihub_content_v1_content_proto_rawDescGZIP(), []int{0}
}

type CreateContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateContentResponse) Reset() {
	*x = CreateContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_content_v1_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentResponse) ProtoMessage() {}

func (x *CreateContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_content_v1_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentResponse.ProtoReflect.Descriptor instead.
func (*CreateContentResponse) Descriptor() ([]byte, []int) {
	return file_ouchihub_content_v1_content_proto_rawDescGZIP(), []int{1}
}

type ListContentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListContentsRequest) Reset() {
	*x = ListContentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_content_v1_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContentsRequest) ProtoMessage() {}

func (x *ListContentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_content_v1_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContentsRequest.ProtoReflect.Descriptor instead.
func (*ListContentsRequest) Descriptor() ([]byte, []int) {
	return file_ouchihub_content_v1_content_proto_rawDescGZIP(), []int{2}
}

type ListContentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListContentsResponse) Reset() {
	*x = ListContentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_content_v1_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContentsResponse) ProtoMessage() {}

func (x *ListContentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_content_v1_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContentsResponse.ProtoReflect.Descriptor instead.
func (*ListContentsResponse) Descriptor() ([]byte, []int) {
	return file_ouchihub_content_v1_content_proto_rawDescGZIP(), []int{3}
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_content_v1_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_content_v1_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_ouchihub_content_v1_content_proto_rawDescGZIP(), []int{4}
}

var File_ouchihub_content_v1_content_proto protoreflect.FileDescriptor

var file_ouchihub_content_v1_content_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x32, 0xe1, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x65, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x28, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x75, 0x63,
	0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe4, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6b, 0x69, 0x6d, 0x6f, 0x74, 0x6f, 0x2d, 0x6a, 0x75, 0x6e, 0x79, 0x61, 0x2f, 0x6f, 0x75,
	0x63, 0x68, 0x69, 0x2d, 0x68, 0x75, 0x62, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x43, 0x58, 0xaa, 0x02, 0x13, 0x4f, 0x75, 0x63, 0x68, 0x69,
	0x68, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x13, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x5c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75,
	0x62, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ouchihub_content_v1_content_proto_rawDescOnce sync.Once
	file_ouchihub_content_v1_content_proto_rawDescData = file_ouchihub_content_v1_content_proto_rawDesc
)

func file_ouchihub_content_v1_content_proto_rawDescGZIP() []byte {
	file_ouchihub_content_v1_content_proto_rawDescOnce.Do(func() {
		file_ouchihub_content_v1_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_ouchihub_content_v1_content_proto_rawDescData)
	})
	return file_ouchihub_content_v1_content_proto_rawDescData
}

var file_ouchihub_content_v1_content_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ouchihub_content_v1_content_proto_goTypes = []interface{}{
	(*CreateContentRequest)(nil),  // 0: ouchihub.content.v1.CreateContentRequest
	(*CreateContentResponse)(nil), // 1: ouchihub.content.v1.CreateContentResponse
	(*ListContentsRequest)(nil),   // 2: ouchihub.content.v1.ListContentsRequest
	(*ListContentsResponse)(nil),  // 3: ouchihub.content.v1.ListContentsResponse
	(*Content)(nil),               // 4: ouchihub.content.v1.Content
}
var file_ouchihub_content_v1_content_proto_depIdxs = []int32{
	0, // 0: ouchihub.content.v1.ContentService.CreateContent:input_type -> ouchihub.content.v1.CreateContentRequest
	2, // 1: ouchihub.content.v1.ContentService.ListContents:input_type -> ouchihub.content.v1.ListContentsRequest
	1, // 2: ouchihub.content.v1.ContentService.CreateContent:output_type -> ouchihub.content.v1.CreateContentResponse
	3, // 3: ouchihub.content.v1.ContentService.ListContents:output_type -> ouchihub.content.v1.ListContentsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ouchihub_content_v1_content_proto_init() }
func file_ouchihub_content_v1_content_proto_init() {
	if File_ouchihub_content_v1_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ouchihub_content_v1_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentRequest); i {
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
		file_ouchihub_content_v1_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentResponse); i {
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
		file_ouchihub_content_v1_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContentsRequest); i {
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
		file_ouchihub_content_v1_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContentsResponse); i {
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
		file_ouchihub_content_v1_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
			RawDescriptor: file_ouchihub_content_v1_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ouchihub_content_v1_content_proto_goTypes,
		DependencyIndexes: file_ouchihub_content_v1_content_proto_depIdxs,
		MessageInfos:      file_ouchihub_content_v1_content_proto_msgTypes,
	}.Build()
	File_ouchihub_content_v1_content_proto = out.File
	file_ouchihub_content_v1_content_proto_rawDesc = nil
	file_ouchihub_content_v1_content_proto_goTypes = nil
	file_ouchihub_content_v1_content_proto_depIdxs = nil
}
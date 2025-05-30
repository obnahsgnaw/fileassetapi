// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: fileasset_backend_api/index/v1/index.proto

package indexv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v1 "github.com/obnahsgnaw/fileassetapi/gen/fileasset_backend_api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_fileasset_backend_api_index_v1_index_proto_rawDescGZIP(), []int{0}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_fileasset_backend_api_index_v1_index_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type OptionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OptionListRequest) Reset() {
	*x = OptionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionListRequest) ProtoMessage() {}

func (x *OptionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionListRequest.ProtoReflect.Descriptor instead.
func (*OptionListRequest) Descriptor() ([]byte, []int) {
	return file_fileasset_backend_api_index_v1_index_proto_rawDescGZIP(), []int{2}
}

type OptionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableState []*v1.IntOption `protobuf:"bytes,1,rep,name=enable_state,json=enableState,proto3" json:"enable_state,omitempty"`
}

func (x *OptionListResponse) Reset() {
	*x = OptionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionListResponse) ProtoMessage() {}

func (x *OptionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileasset_backend_api_index_v1_index_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionListResponse.ProtoReflect.Descriptor instead.
func (*OptionListResponse) Descriptor() ([]byte, []int) {
	return file_fileasset_backend_api_index_v1_index_proto_rawDescGZIP(), []int{3}
}

func (x *OptionListResponse) GetEnableState() []*v1.IntOption {
	if x != nil {
		return x.EnableState
	}
	return nil
}

var File_fileasset_backend_api_index_v1_index_proto protoreflect.FileDescriptor

var file_fileasset_backend_api_index_v1_index_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x66, 0x69,
	0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x66, 0x69, 0x6c, 0x65,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10,
	0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x40, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0xe7, 0x89, 0x88, 0xe6, 0x9c,
	0xac, 0xe5, 0x8f, 0xb7, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x03, 0x80,
	0x43, 0x01, 0x22, 0x13, 0x0a, 0x11, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7e, 0x0a, 0x12, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a,
	0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x14, 0x92, 0x41, 0x11, 0x32, 0x0f, 0xe5, 0x90, 0xaf, 0xe7, 0xa6, 0x81, 0xe7, 0x94, 0xa8, 0xe7,
	0x8a, 0xb6, 0xe6, 0x80, 0x81, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x3a, 0x03, 0x80, 0x43, 0x01, 0x32, 0xd5, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb7, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x92, 0x41, 0x35, 0x12, 0x09, 0xe7, 0x89, 0x88, 0xe6,
	0x9c, 0xac, 0xe5, 0x8f, 0xb7, 0x1a, 0x1b, 0xe8, 0xbf, 0x94, 0xe5, 0x9b, 0x9e, 0xe6, 0x9c, 0x8d,
	0xe5, 0x8a, 0xa1, 0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe7, 0x89, 0x88, 0xe6, 0x9c, 0xac, 0xe5,
	0x8f, 0xb7, 0x62, 0x0b, 0x0a, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x0b, 0x92, 0x41, 0x08, 0x12, 0x06, 0xe9, 0xbb, 0x98, 0xe8, 0xae, 0xa4, 0x32,
	0xdc, 0x01, 0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xb6, 0x01, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x31, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x92, 0x41, 0x2b, 0x12, 0x06, 0xe9,
	0x80, 0x89, 0xe9, 0xa1, 0xb9, 0x62, 0x0b, 0x0a, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x00, 0x6a, 0x14, 0x0a, 0x07, 0x78, 0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x09, 0x11,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x11, 0x92, 0x41, 0x0e,
	0x12, 0x0c, 0xe9, 0x80, 0x89, 0xe9, 0xa1, 0xb9, 0xe7, 0xae, 0xa1, 0xe7, 0x90, 0x86, 0x42, 0xa2,
	0x09, 0x92, 0x41, 0x8d, 0x07, 0x12, 0x5f, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x20, 0x41, 0x50, 0x49, 0x12, 0x1a, 0x54, 0x68, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x20, 0x61, 0x70, 0x69,
	0x22, 0x2f, 0x0a, 0x04, 0x65, 0x66, 0x6c, 0x79, 0x12, 0x15, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x65, 0x66, 0x6c, 0x79, 0x2d, 0x63, 0x65, 0x74, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x1a,
	0x10, 0x6e, 0x6f, 0x6e, 0x65, 0x40, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x18, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x2d,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x18, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x34, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x34, 0x0a, 0x03, 0x32, 0x30, 0x34, 0x12,
	0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x6b,
	0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x64, 0x0a, 0x2f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x2e, 0x12, 0x31, 0x0a, 0x2f, 0x1a, 0x2d, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x38, 0x0a, 0x03, 0x34,
	0x30, 0x31, 0x12, 0x31, 0x0a, 0x2f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77,
	0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65,
	0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20,
	0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x2e, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x27,
	0x0a, 0x25, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x52, 0x2f, 0x0a, 0x03, 0x34, 0x30, 0x39, 0x12, 0x28,
	0x0a, 0x26, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63,
	0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x2e, 0x52, 0x27, 0x0a, 0x03, 0x34, 0x32, 0x33, 0x12,
	0x20, 0x0a, 0x1e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x2e, 0x52, 0x37, 0x0a, 0x03, 0x34, 0x32, 0x39, 0x12, 0x30, 0x0a, 0x2e, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x20, 0x74, 0x6f, 0x6f, 0x20, 0x6d, 0x61, 0x6e, 0x79, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x20,
	0x75, 0x6e, 0x69, 0x74, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x16, 0x0a, 0x03, 0x35, 0x30,
	0x30, 0x12, 0x0f, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x52, 0x4f, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x44, 0x0a,
	0x0f, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x12, 0x31, 0x0a, 0x2f, 0x1a, 0x2d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5a, 0x3d, 0x0a, 0x17, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x0e,
	0x08, 0x02, 0x1a, 0x08, 0x58, 0x2d, 0x41, 0x70, 0x70, 0x2d, 0x49, 0x64, 0x20, 0x02, 0x0a, 0x22,
	0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x08,
	0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x62, 0x6e, 0x61, 0x68, 0x73, 0x67, 0x6e, 0x61, 0x77, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x49, 0x58, 0xaa, 0x02, 0x1c, 0x46, 0x69, 0x6c, 0x65,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1c, 0x46, 0x69, 0x6c, 0x65, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x5c, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x28, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x5c, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1e, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fileasset_backend_api_index_v1_index_proto_rawDescOnce sync.Once
	file_fileasset_backend_api_index_v1_index_proto_rawDescData = file_fileasset_backend_api_index_v1_index_proto_rawDesc
)

func file_fileasset_backend_api_index_v1_index_proto_rawDescGZIP() []byte {
	file_fileasset_backend_api_index_v1_index_proto_rawDescOnce.Do(func() {
		file_fileasset_backend_api_index_v1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileasset_backend_api_index_v1_index_proto_rawDescData)
	})
	return file_fileasset_backend_api_index_v1_index_proto_rawDescData
}

var file_fileasset_backend_api_index_v1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fileasset_backend_api_index_v1_index_proto_goTypes = []interface{}{
	(*VersionRequest)(nil),     // 0: fileasset_backend_api.index.v1.VersionRequest
	(*VersionResponse)(nil),    // 1: fileasset_backend_api.index.v1.VersionResponse
	(*OptionListRequest)(nil),  // 2: fileasset_backend_api.index.v1.OptionListRequest
	(*OptionListResponse)(nil), // 3: fileasset_backend_api.index.v1.OptionListResponse
	(*v1.IntOption)(nil),       // 4: fileasset_backend_api.common.v1.IntOption
}
var file_fileasset_backend_api_index_v1_index_proto_depIdxs = []int32{
	4, // 0: fileasset_backend_api.index.v1.OptionListResponse.enable_state:type_name -> fileasset_backend_api.common.v1.IntOption
	0, // 1: fileasset_backend_api.index.v1.IndexService.Version:input_type -> fileasset_backend_api.index.v1.VersionRequest
	2, // 2: fileasset_backend_api.index.v1.OptionsService.OptionList:input_type -> fileasset_backend_api.index.v1.OptionListRequest
	1, // 3: fileasset_backend_api.index.v1.IndexService.Version:output_type -> fileasset_backend_api.index.v1.VersionResponse
	3, // 4: fileasset_backend_api.index.v1.OptionsService.OptionList:output_type -> fileasset_backend_api.index.v1.OptionListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fileasset_backend_api_index_v1_index_proto_init() }
func file_fileasset_backend_api_index_v1_index_proto_init() {
	if File_fileasset_backend_api_index_v1_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fileasset_backend_api_index_v1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_fileasset_backend_api_index_v1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_fileasset_backend_api_index_v1_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionListRequest); i {
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
		file_fileasset_backend_api_index_v1_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionListResponse); i {
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
			RawDescriptor: file_fileasset_backend_api_index_v1_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_fileasset_backend_api_index_v1_index_proto_goTypes,
		DependencyIndexes: file_fileasset_backend_api_index_v1_index_proto_depIdxs,
		MessageInfos:      file_fileasset_backend_api_index_v1_index_proto_msgTypes,
	}.Build()
	File_fileasset_backend_api_index_v1_index_proto = out.File
	file_fileasset_backend_api_index_v1_index_proto_rawDesc = nil
	file_fileasset_backend_api_index_v1_index_proto_goTypes = nil
	file_fileasset_backend_api_index_v1_index_proto_depIdxs = nil
}

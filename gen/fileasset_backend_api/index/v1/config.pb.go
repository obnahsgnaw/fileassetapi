// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: fileasset_backend_api/index/v1/config.proto

package indexv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileasset_backend_api_index_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileasset_backend_api_index_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_fileasset_backend_api_index_v1_config_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Debug bool `protobuf:"varint,1,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileasset_backend_api_index_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_fileasset_backend_api_index_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_fileasset_backend_api_index_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

var File_fileasset_backend_api_index_v1_config_proto protoreflect.FileDescriptor

var file_fileasset_backend_api_index_v1_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x66,
	0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe8, 0xb0, 0x83, 0xe8,
	0xaf, 0x95, 0xe6, 0xa8, 0xa1, 0xe5, 0xbc, 0x8f, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x32,
	0xa5, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xc2, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x61, 0x92, 0x41, 0x4c, 0x12, 0x0c, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae,
	0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x1a, 0x1e, 0xe8, 0xbf, 0x94, 0xe5, 0x9b, 0x9e, 0xe6, 0x9c,
	0x8d, 0xe5, 0x8a, 0xa1, 0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae,
	0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x62, 0x1c, 0x0a, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x00, 0x0a, 0x0f, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0xbb, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x26, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x26, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x61, 0x92, 0x41, 0x4c, 0x12, 0x0c, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe9, 0x85,
	0x8d, 0xe7, 0xbd, 0xae, 0x1a, 0x1e, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe6, 0x9c, 0x8d, 0xe5,
	0x8a, 0xa1, 0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0xe4, 0xbf,
	0xa1, 0xe6, 0x81, 0xaf, 0x62, 0x1c, 0x0a, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x00, 0x0a, 0x0f, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x1a, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x11, 0x92, 0x41, 0x0e, 0x12, 0x0c, 0xe5, 0x8a, 0xa8, 0xe6, 0x80,
	0x81, 0xe9, 0x85, 0x8d, 0xe7, 0xbd, 0xae, 0x42, 0x92, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x62, 0x6e, 0x61, 0x68, 0x73,
	0x67, 0x6e, 0x61, 0x77, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46,
	0x49, 0x58, 0xaa, 0x02, 0x1c, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1c, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x5c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x28, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x5c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x46, 0x69,
	0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fileasset_backend_api_index_v1_config_proto_rawDescOnce sync.Once
	file_fileasset_backend_api_index_v1_config_proto_rawDescData = file_fileasset_backend_api_index_v1_config_proto_rawDesc
)

func file_fileasset_backend_api_index_v1_config_proto_rawDescGZIP() []byte {
	file_fileasset_backend_api_index_v1_config_proto_rawDescOnce.Do(func() {
		file_fileasset_backend_api_index_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileasset_backend_api_index_v1_config_proto_rawDescData)
	})
	return file_fileasset_backend_api_index_v1_config_proto_rawDescData
}

var file_fileasset_backend_api_index_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fileasset_backend_api_index_v1_config_proto_goTypes = []interface{}{
	(*ConfigRequest)(nil), // 0: fileasset_backend_api.index.v1.ConfigRequest
	(*Config)(nil),        // 1: fileasset_backend_api.index.v1.Config
}
var file_fileasset_backend_api_index_v1_config_proto_depIdxs = []int32{
	0, // 0: fileasset_backend_api.index.v1.ConfigService.Detail:input_type -> fileasset_backend_api.index.v1.ConfigRequest
	1, // 1: fileasset_backend_api.index.v1.ConfigService.Update:input_type -> fileasset_backend_api.index.v1.Config
	1, // 2: fileasset_backend_api.index.v1.ConfigService.Detail:output_type -> fileasset_backend_api.index.v1.Config
	1, // 3: fileasset_backend_api.index.v1.ConfigService.Update:output_type -> fileasset_backend_api.index.v1.Config
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fileasset_backend_api_index_v1_config_proto_init() }
func file_fileasset_backend_api_index_v1_config_proto_init() {
	if File_fileasset_backend_api_index_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fileasset_backend_api_index_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_fileasset_backend_api_index_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_fileasset_backend_api_index_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fileasset_backend_api_index_v1_config_proto_goTypes,
		DependencyIndexes: file_fileasset_backend_api_index_v1_config_proto_depIdxs,
		MessageInfos:      file_fileasset_backend_api_index_v1_config_proto_msgTypes,
	}.Build()
	File_fileasset_backend_api_index_v1_config_proto = out.File
	file_fileasset_backend_api_index_v1_config_proto_rawDesc = nil
	file_fileasset_backend_api_index_v1_config_proto_goTypes = nil
	file_fileasset_backend_api_index_v1_config_proto_depIdxs = nil
}

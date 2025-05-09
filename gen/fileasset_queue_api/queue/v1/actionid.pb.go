// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: fileasset_queue_api/queue/v1/actionid.proto

package queuev1

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

type ActionId int32

const (
	ActionId_None              ActionId = 0
	ActionId_HandlerConnectReq ActionId = 591 // 客户端连接
	ActionId_HandlerConnectRsp ActionId = 592 // 服务端回放
	ActionId_TaskReq           ActionId = 593 // 服务端下发任务
	ActionId_TaskRsp           ActionId = 594 // 客户端回复收到任务
	ActionId_TaskProgressReq   ActionId = 595 // 客户端上报任务进度
	ActionId_TaskResultReq     ActionId = 596 // 客户端上报任务结果
)

// Enum value maps for ActionId.
var (
	ActionId_name = map[int32]string{
		0:   "None",
		591: "HandlerConnectReq",
		592: "HandlerConnectRsp",
		593: "TaskReq",
		594: "TaskRsp",
		595: "TaskProgressReq",
		596: "TaskResultReq",
	}
	ActionId_value = map[string]int32{
		"None":              0,
		"HandlerConnectReq": 591,
		"HandlerConnectRsp": 592,
		"TaskReq":           593,
		"TaskRsp":           594,
		"TaskProgressReq":   595,
		"TaskResultReq":     596,
	}
)

func (x ActionId) Enum() *ActionId {
	p := new(ActionId)
	*p = x
	return p
}

func (x ActionId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionId) Descriptor() protoreflect.EnumDescriptor {
	return file_fileasset_queue_api_queue_v1_actionid_proto_enumTypes[0].Descriptor()
}

func (ActionId) Type() protoreflect.EnumType {
	return &file_fileasset_queue_api_queue_v1_actionid_proto_enumTypes[0]
}

func (x ActionId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionId.Descriptor instead.
func (ActionId) EnumDescriptor() ([]byte, []int) {
	return file_fileasset_queue_api_queue_v1_actionid_proto_rawDescGZIP(), []int{0}
}

var File_fileasset_queue_api_queue_v1_actionid_proto protoreflect.FileDescriptor

var file_fileasset_queue_api_queue_v1_actionid_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x66,
	0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x8a, 0x01, 0x0a, 0x08,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x10, 0xcf, 0x04, 0x12, 0x16, 0x0a, 0x11, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x73, 0x70, 0x10,
	0xd0, 0x04, 0x12, 0x0c, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x10, 0xd1, 0x04,
	0x12, 0x0c, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x73, 0x70, 0x10, 0xd2, 0x04, 0x12, 0x14,
	0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x10, 0xd3, 0x04, 0x12, 0x12, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x71, 0x10, 0xd4, 0x04, 0x42, 0x82, 0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x66, 0x6c, 0x79, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x51, 0x58, 0xaa, 0x02, 0x1a, 0x46, 0x69,
	0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x41, 0x70, 0x69, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x46, 0x69, 0x6c, 0x65, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x41, 0x70, 0x69, 0x5c, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x26, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x41, 0x70, 0x69, 0x5c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1c, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x73, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fileasset_queue_api_queue_v1_actionid_proto_rawDescOnce sync.Once
	file_fileasset_queue_api_queue_v1_actionid_proto_rawDescData = file_fileasset_queue_api_queue_v1_actionid_proto_rawDesc
)

func file_fileasset_queue_api_queue_v1_actionid_proto_rawDescGZIP() []byte {
	file_fileasset_queue_api_queue_v1_actionid_proto_rawDescOnce.Do(func() {
		file_fileasset_queue_api_queue_v1_actionid_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileasset_queue_api_queue_v1_actionid_proto_rawDescData)
	})
	return file_fileasset_queue_api_queue_v1_actionid_proto_rawDescData
}

var file_fileasset_queue_api_queue_v1_actionid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fileasset_queue_api_queue_v1_actionid_proto_goTypes = []interface{}{
	(ActionId)(0), // 0: fileasset_queue_api.queue.v1.ActionId
}
var file_fileasset_queue_api_queue_v1_actionid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fileasset_queue_api_queue_v1_actionid_proto_init() }
func file_fileasset_queue_api_queue_v1_actionid_proto_init() {
	if File_fileasset_queue_api_queue_v1_actionid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fileasset_queue_api_queue_v1_actionid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fileasset_queue_api_queue_v1_actionid_proto_goTypes,
		DependencyIndexes: file_fileasset_queue_api_queue_v1_actionid_proto_depIdxs,
		EnumInfos:         file_fileasset_queue_api_queue_v1_actionid_proto_enumTypes,
	}.Build()
	File_fileasset_queue_api_queue_v1_actionid_proto = out.File
	file_fileasset_queue_api_queue_v1_actionid_proto_rawDesc = nil
	file_fileasset_queue_api_queue_v1_actionid_proto_goTypes = nil
	file_fileasset_queue_api_queue_v1_actionid_proto_depIdxs = nil
}

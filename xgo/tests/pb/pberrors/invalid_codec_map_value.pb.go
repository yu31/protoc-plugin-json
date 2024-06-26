// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/errors/invalid_codec_map_value.proto

package pberrors

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

// TODO: Add test cases for map value.
type InvalidCodecMapValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InvalidCodecMapValue) Reset() {
	*x = InvalidCodecMapValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_errors_invalid_codec_map_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidCodecMapValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidCodecMapValue) ProtoMessage() {}

func (x *InvalidCodecMapValue) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_errors_invalid_codec_map_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidCodecMapValue.ProtoReflect.Descriptor instead.
func (*InvalidCodecMapValue) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescGZIP(), []int{0}
}

var File_tests_proto_cases_errors_invalid_codec_map_value_proto protoreflect.FileDescriptor

var file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0x16, 0x0a, 0x14, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x63,
	0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x78, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescOnce sync.Once
	file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescData = file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDesc
)

func file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescData)
	})
	return file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDescData
}

var file_tests_proto_cases_errors_invalid_codec_map_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_proto_cases_errors_invalid_codec_map_value_proto_goTypes = []interface{}{
	(*InvalidCodecMapValue)(nil), // 0: errors.InvalidCodecMapValue
}
var file_tests_proto_cases_errors_invalid_codec_map_value_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_errors_invalid_codec_map_value_proto_init() }
func file_tests_proto_cases_errors_invalid_codec_map_value_proto_init() {
	if File_tests_proto_cases_errors_invalid_codec_map_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_errors_invalid_codec_map_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidCodecMapValue); i {
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
			RawDescriptor: file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_errors_invalid_codec_map_value_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_errors_invalid_codec_map_value_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_errors_invalid_codec_map_value_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_errors_invalid_codec_map_value_proto = out.File
	file_tests_proto_cases_errors_invalid_codec_map_value_proto_rawDesc = nil
	file_tests_proto_cases_errors_invalid_codec_map_value_proto_goTypes = nil
	file_tests_proto_cases_errors_invalid_codec_map_value_proto_depIdxs = nil
}

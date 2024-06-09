// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/errors/invalid_syntax_json.proto

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

type InvalidSyntaxJSON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FInt32A int32 `protobuf:"varint,1,opt,name=f_int32a,json=fInt32a,proto3" json:"f_int32a,omitempty"`
	FInt32B int32 `protobuf:"varint,2,opt,name=f_int32b,json=fInt32b,proto3" json:"f_int32b,omitempty"`
}

func (x *InvalidSyntaxJSON) Reset() {
	*x = InvalidSyntaxJSON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_errors_invalid_syntax_json_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidSyntaxJSON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidSyntaxJSON) ProtoMessage() {}

func (x *InvalidSyntaxJSON) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_errors_invalid_syntax_json_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidSyntaxJSON.ProtoReflect.Descriptor instead.
func (*InvalidSyntaxJSON) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescGZIP(), []int{0}
}

func (x *InvalidSyntaxJSON) GetFInt32A() int32 {
	if x != nil {
		return x.FInt32A
	}
	return 0
}

func (x *InvalidSyntaxJSON) GetFInt32B() int32 {
	if x != nil {
		return x.FInt32B
	}
	return 0
}

var File_tests_proto_cases_errors_invalid_syntax_json_proto protoreflect.FileDescriptor

var file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDesc = []byte{
	0x0a, 0x32, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x49, 0x0a, 0x11,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x61, 0x12, 0x19, 0x0a, 0x08,
	0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x62, 0x42, 0x17, 0x5a, 0x15, 0x78, 0x67, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescOnce sync.Once
	file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescData = file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDesc
)

func file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescData)
	})
	return file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDescData
}

var file_tests_proto_cases_errors_invalid_syntax_json_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_proto_cases_errors_invalid_syntax_json_proto_goTypes = []interface{}{
	(*InvalidSyntaxJSON)(nil), // 0: errors.InvalidSyntaxJSON
}
var file_tests_proto_cases_errors_invalid_syntax_json_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_errors_invalid_syntax_json_proto_init() }
func file_tests_proto_cases_errors_invalid_syntax_json_proto_init() {
	if File_tests_proto_cases_errors_invalid_syntax_json_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_errors_invalid_syntax_json_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidSyntaxJSON); i {
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
			RawDescriptor: file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_errors_invalid_syntax_json_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_errors_invalid_syntax_json_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_errors_invalid_syntax_json_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_errors_invalid_syntax_json_proto = out.File
	file_tests_proto_cases_errors_invalid_syntax_json_proto_rawDesc = nil
	file_tests_proto_cases_errors_invalid_syntax_json_proto_goTypes = nil
	file_tests_proto_cases_errors_invalid_syntax_json_proto_depIdxs = nil
}
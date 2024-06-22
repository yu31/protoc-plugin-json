// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/boundary/cycle5.proto

package pbboundary

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
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

type InlineMessageCycle5 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1  string                             `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2  string                             `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
	FMessage1 *InlineMessageCycle5_EmbedMessage1 `protobuf:"bytes,3,opt,name=f_message1,json=fMessage1,proto3" json:"f_message1,omitempty"`
}

func (x *InlineMessageCycle5) Reset() {
	*x = InlineMessageCycle5{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_boundary_cycle5_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InlineMessageCycle5) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InlineMessageCycle5) ProtoMessage() {}

func (x *InlineMessageCycle5) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_boundary_cycle5_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InlineMessageCycle5.ProtoReflect.Descriptor instead.
func (*InlineMessageCycle5) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_boundary_cycle5_proto_rawDescGZIP(), []int{0}
}

func (x *InlineMessageCycle5) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *InlineMessageCycle5) GetFString2() string {
	if x != nil {
		return x.FString2
	}
	return ""
}

func (x *InlineMessageCycle5) GetFMessage1() *InlineMessageCycle5_EmbedMessage1 {
	if x != nil {
		return x.FMessage1
	}
	return nil
}

type InlineMessageCycle5_EmbedMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString3 string `protobuf:"bytes,1,opt,name=f_string3,json=fString3,proto3" json:"f_string3,omitempty"`
	FString4 string `protobuf:"bytes,2,opt,name=f_string4,json=fString4,proto3" json:"f_string4,omitempty"`
	// Types that are assignable to OneType2:
	//	*InlineMessageCycle5_EmbedMessage1_FMessage2
	//	*InlineMessageCycle5_EmbedMessage1_FString5
	//	*InlineMessageCycle5_EmbedMessage1_FString6
	OneType2 isInlineMessageCycle5_EmbedMessage1_OneType2 `protobuf_oneof:"OneType2"`
}

func (x *InlineMessageCycle5_EmbedMessage1) Reset() {
	*x = InlineMessageCycle5_EmbedMessage1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_boundary_cycle5_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InlineMessageCycle5_EmbedMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InlineMessageCycle5_EmbedMessage1) ProtoMessage() {}

func (x *InlineMessageCycle5_EmbedMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_boundary_cycle5_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InlineMessageCycle5_EmbedMessage1.ProtoReflect.Descriptor instead.
func (*InlineMessageCycle5_EmbedMessage1) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_boundary_cycle5_proto_rawDescGZIP(), []int{0, 0}
}

func (x *InlineMessageCycle5_EmbedMessage1) GetFString3() string {
	if x != nil {
		return x.FString3
	}
	return ""
}

func (x *InlineMessageCycle5_EmbedMessage1) GetFString4() string {
	if x != nil {
		return x.FString4
	}
	return ""
}

func (m *InlineMessageCycle5_EmbedMessage1) GetOneType2() isInlineMessageCycle5_EmbedMessage1_OneType2 {
	if m != nil {
		return m.OneType2
	}
	return nil
}

func (x *InlineMessageCycle5_EmbedMessage1) GetFMessage2() *InlineMessageCycle5 {
	if x, ok := x.GetOneType2().(*InlineMessageCycle5_EmbedMessage1_FMessage2); ok {
		return x.FMessage2
	}
	return nil
}

func (x *InlineMessageCycle5_EmbedMessage1) GetFString5() string {
	if x, ok := x.GetOneType2().(*InlineMessageCycle5_EmbedMessage1_FString5); ok {
		return x.FString5
	}
	return ""
}

func (x *InlineMessageCycle5_EmbedMessage1) GetFString6() string {
	if x, ok := x.GetOneType2().(*InlineMessageCycle5_EmbedMessage1_FString6); ok {
		return x.FString6
	}
	return ""
}

type isInlineMessageCycle5_EmbedMessage1_OneType2 interface {
	isInlineMessageCycle5_EmbedMessage1_OneType2()
}

type InlineMessageCycle5_EmbedMessage1_FMessage2 struct {
	FMessage2 *InlineMessageCycle5 `protobuf:"bytes,3,opt,name=f_message2,json=fMessage2,proto3,oneof"`
}

type InlineMessageCycle5_EmbedMessage1_FString5 struct {
	FString5 string `protobuf:"bytes,4,opt,name=f_string5,json=fString5,proto3,oneof"`
}

type InlineMessageCycle5_EmbedMessage1_FString6 struct {
	FString6 string `protobuf:"bytes,5,opt,name=f_string6,json=fString6,proto3,oneof"`
}

func (*InlineMessageCycle5_EmbedMessage1_FMessage2) isInlineMessageCycle5_EmbedMessage1_OneType2() {}

func (*InlineMessageCycle5_EmbedMessage1_FString5) isInlineMessageCycle5_EmbedMessage1_OneType2() {}

func (*InlineMessageCycle5_EmbedMessage1_FString6) isInlineMessageCycle5_EmbedMessage1_OneType2() {}

var File_tests_proto_cases_boundary_cycle5_proto protoreflect.FileDescriptor

var file_tests_proto_cases_boundary_cycle5_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x03, 0x0a, 0x13, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x35, 0x12, 0x21, 0x0a,
	0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31,
	0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x32, 0x12, 0x52, 0x0a, 0x0a, 0x66, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x79, 0x63, 0x6c, 0x65, 0x35, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x42, 0x06, 0x8a, 0xa2, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x09, 0x66, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x1a, 0xfb, 0x01, 0x0a, 0x0d, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2,
	0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x12, 0x21, 0x0a, 0x09,
	0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x12,
	0x46, 0x0a, 0x0a, 0x66, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2e, 0x49,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x35, 0x42, 0x06, 0x8a, 0xa2, 0x1f, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x09, 0x66, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x35, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00,
	0x48, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x35, 0x12, 0x23, 0x0a, 0x09,
	0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x36, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x48, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x36, 0x42, 0x12, 0x0a, 0x08, 0x4f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32, 0x12, 0x06, 0x92,
	0xa2, 0x1f, 0x02, 0x20, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_cases_boundary_cycle5_proto_rawDescOnce sync.Once
	file_tests_proto_cases_boundary_cycle5_proto_rawDescData = file_tests_proto_cases_boundary_cycle5_proto_rawDesc
)

func file_tests_proto_cases_boundary_cycle5_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_boundary_cycle5_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_boundary_cycle5_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_boundary_cycle5_proto_rawDescData)
	})
	return file_tests_proto_cases_boundary_cycle5_proto_rawDescData
}

var file_tests_proto_cases_boundary_cycle5_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tests_proto_cases_boundary_cycle5_proto_goTypes = []interface{}{
	(*InlineMessageCycle5)(nil),               // 0: boundary.InlineMessageCycle5
	(*InlineMessageCycle5_EmbedMessage1)(nil), // 1: boundary.InlineMessageCycle5.EmbedMessage1
}
var file_tests_proto_cases_boundary_cycle5_proto_depIdxs = []int32{
	1, // 0: boundary.InlineMessageCycle5.f_message1:type_name -> boundary.InlineMessageCycle5.EmbedMessage1
	0, // 1: boundary.InlineMessageCycle5.EmbedMessage1.f_message2:type_name -> boundary.InlineMessageCycle5
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_boundary_cycle5_proto_init() }
func file_tests_proto_cases_boundary_cycle5_proto_init() {
	if File_tests_proto_cases_boundary_cycle5_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_boundary_cycle5_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InlineMessageCycle5); i {
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
		file_tests_proto_cases_boundary_cycle5_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InlineMessageCycle5_EmbedMessage1); i {
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
	file_tests_proto_cases_boundary_cycle5_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*InlineMessageCycle5_EmbedMessage1_FMessage2)(nil),
		(*InlineMessageCycle5_EmbedMessage1_FString5)(nil),
		(*InlineMessageCycle5_EmbedMessage1_FString6)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_proto_cases_boundary_cycle5_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_boundary_cycle5_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_boundary_cycle5_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_boundary_cycle5_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_boundary_cycle5_proto = out.File
	file_tests_proto_cases_boundary_cycle5_proto_rawDesc = nil
	file_tests_proto_cases_boundary_cycle5_proto_goTypes = nil
	file_tests_proto_cases_boundary_cycle5_proto_depIdxs = nil
}

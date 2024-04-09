// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/references/type_plain.proto

package pbref

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

type TypePlain1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FInt32A    int32   `protobuf:"varint,1,opt,name=f_int32a,json=fInt32a,proto3" json:"f_int32a,omitempty"`
	FInt32B    int32   `protobuf:"varint,2,opt,name=f_int32b,json=fInt32b,proto3" json:"f_int32b,omitempty"`
	FInt64A    int64   `protobuf:"varint,3,opt,name=f_int64a,json=fInt64a,proto3" json:"f_int64a,omitempty"`
	FInt64B    int64   `protobuf:"varint,4,opt,name=f_int64b,json=fInt64b,proto3" json:"f_int64b,omitempty"`
	FUint32A   uint32  `protobuf:"varint,5,opt,name=f_uint32a,json=fUint32a,proto3" json:"f_uint32a,omitempty"`
	FUint32B   uint32  `protobuf:"varint,6,opt,name=f_uint32b,json=fUint32b,proto3" json:"f_uint32b,omitempty"`
	FUint64A   uint64  `protobuf:"varint,7,opt,name=f_uint64a,json=fUint64a,proto3" json:"f_uint64a,omitempty"`
	FUint64B   uint64  `protobuf:"varint,8,opt,name=f_uint64b,json=fUint64b,proto3" json:"f_uint64b,omitempty"`
	FSint32A   int32   `protobuf:"zigzag32,9,opt,name=f_sint32a,json=fSint32a,proto3" json:"f_sint32a,omitempty"`
	FSint32B   int32   `protobuf:"zigzag32,10,opt,name=f_sint32b,json=fSint32b,proto3" json:"f_sint32b,omitempty"`
	FSint64A   int64   `protobuf:"zigzag64,11,opt,name=f_sint64a,json=fSint64a,proto3" json:"f_sint64a,omitempty"`
	FSint64B   int64   `protobuf:"zigzag64,12,opt,name=f_sint64b,json=fSint64b,proto3" json:"f_sint64b,omitempty"`
	FSfixed32A int32   `protobuf:"fixed32,13,opt,name=f_sfixed32a,json=fSfixed32a,proto3" json:"f_sfixed32a,omitempty"`
	FSfixed32B int32   `protobuf:"fixed32,14,opt,name=f_sfixed32b,json=fSfixed32b,proto3" json:"f_sfixed32b,omitempty"`
	FSfixed64A int64   `protobuf:"fixed64,15,opt,name=f_sfixed64a,json=fSfixed64a,proto3" json:"f_sfixed64a,omitempty"`
	FSfixed64B int64   `protobuf:"fixed64,16,opt,name=f_sfixed64b,json=fSfixed64b,proto3" json:"f_sfixed64b,omitempty"`
	FFixed32A  uint32  `protobuf:"fixed32,17,opt,name=f_fixed32a,json=fFixed32a,proto3" json:"f_fixed32a,omitempty"`
	FFixed32B  uint32  `protobuf:"fixed32,18,opt,name=f_fixed32b,json=fFixed32b,proto3" json:"f_fixed32b,omitempty"`
	FFixed64A  uint64  `protobuf:"fixed64,19,opt,name=f_fixed64a,json=fFixed64a,proto3" json:"f_fixed64a,omitempty"`
	FFixed64B  uint64  `protobuf:"fixed64,20,opt,name=f_fixed64b,json=fFixed64b,proto3" json:"f_fixed64b,omitempty"`
	FFloat1    float32 `protobuf:"fixed32,21,opt,name=f_float1,json=fFloat1,proto3" json:"f_float1,omitempty"`
	FFloat2    float32 `protobuf:"fixed32,22,opt,name=f_float2,json=fFloat2,proto3" json:"f_float2,omitempty"`
	FDouble1   float64 `protobuf:"fixed64,23,opt,name=f_double1,json=fDouble1,proto3" json:"f_double1,omitempty"`
	FDouble2   float64 `protobuf:"fixed64,24,opt,name=f_double2,json=fDouble2,proto3" json:"f_double2,omitempty"`
	FBool1     bool    `protobuf:"varint,25,opt,name=f_bool1,json=fBool1,proto3" json:"f_bool1,omitempty"`
	FBool2     bool    `protobuf:"varint,27,opt,name=f_bool2,json=fBool2,proto3" json:"f_bool2,omitempty"`
}

func (x *TypePlain1) Reset() {
	*x = TypePlain1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_references_type_plain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypePlain1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypePlain1) ProtoMessage() {}

func (x *TypePlain1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_references_type_plain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypePlain1.ProtoReflect.Descriptor instead.
func (*TypePlain1) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_references_type_plain_proto_rawDescGZIP(), []int{0}
}

func (x *TypePlain1) GetFInt32A() int32 {
	if x != nil {
		return x.FInt32A
	}
	return 0
}

func (x *TypePlain1) GetFInt32B() int32 {
	if x != nil {
		return x.FInt32B
	}
	return 0
}

func (x *TypePlain1) GetFInt64A() int64 {
	if x != nil {
		return x.FInt64A
	}
	return 0
}

func (x *TypePlain1) GetFInt64B() int64 {
	if x != nil {
		return x.FInt64B
	}
	return 0
}

func (x *TypePlain1) GetFUint32A() uint32 {
	if x != nil {
		return x.FUint32A
	}
	return 0
}

func (x *TypePlain1) GetFUint32B() uint32 {
	if x != nil {
		return x.FUint32B
	}
	return 0
}

func (x *TypePlain1) GetFUint64A() uint64 {
	if x != nil {
		return x.FUint64A
	}
	return 0
}

func (x *TypePlain1) GetFUint64B() uint64 {
	if x != nil {
		return x.FUint64B
	}
	return 0
}

func (x *TypePlain1) GetFSint32A() int32 {
	if x != nil {
		return x.FSint32A
	}
	return 0
}

func (x *TypePlain1) GetFSint32B() int32 {
	if x != nil {
		return x.FSint32B
	}
	return 0
}

func (x *TypePlain1) GetFSint64A() int64 {
	if x != nil {
		return x.FSint64A
	}
	return 0
}

func (x *TypePlain1) GetFSint64B() int64 {
	if x != nil {
		return x.FSint64B
	}
	return 0
}

func (x *TypePlain1) GetFSfixed32A() int32 {
	if x != nil {
		return x.FSfixed32A
	}
	return 0
}

func (x *TypePlain1) GetFSfixed32B() int32 {
	if x != nil {
		return x.FSfixed32B
	}
	return 0
}

func (x *TypePlain1) GetFSfixed64A() int64 {
	if x != nil {
		return x.FSfixed64A
	}
	return 0
}

func (x *TypePlain1) GetFSfixed64B() int64 {
	if x != nil {
		return x.FSfixed64B
	}
	return 0
}

func (x *TypePlain1) GetFFixed32A() uint32 {
	if x != nil {
		return x.FFixed32A
	}
	return 0
}

func (x *TypePlain1) GetFFixed32B() uint32 {
	if x != nil {
		return x.FFixed32B
	}
	return 0
}

func (x *TypePlain1) GetFFixed64A() uint64 {
	if x != nil {
		return x.FFixed64A
	}
	return 0
}

func (x *TypePlain1) GetFFixed64B() uint64 {
	if x != nil {
		return x.FFixed64B
	}
	return 0
}

func (x *TypePlain1) GetFFloat1() float32 {
	if x != nil {
		return x.FFloat1
	}
	return 0
}

func (x *TypePlain1) GetFFloat2() float32 {
	if x != nil {
		return x.FFloat2
	}
	return 0
}

func (x *TypePlain1) GetFDouble1() float64 {
	if x != nil {
		return x.FDouble1
	}
	return 0
}

func (x *TypePlain1) GetFDouble2() float64 {
	if x != nil {
		return x.FDouble2
	}
	return 0
}

func (x *TypePlain1) GetFBool1() bool {
	if x != nil {
		return x.FBool1
	}
	return false
}

func (x *TypePlain1) GetFBool2() bool {
	if x != nil {
		return x.FBool2
	}
	return false
}

// Only used to test type error.
type TypePlain2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FBytes1  []byte `protobuf:"bytes,2,opt,name=f_bytes1,json=fBytes1,proto3" json:"f_bytes1,omitempty"`
}

func (x *TypePlain2) Reset() {
	*x = TypePlain2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_references_type_plain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypePlain2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypePlain2) ProtoMessage() {}

func (x *TypePlain2) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_references_type_plain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypePlain2.ProtoReflect.Descriptor instead.
func (*TypePlain2) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_references_type_plain_proto_rawDescGZIP(), []int{1}
}

func (x *TypePlain2) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *TypePlain2) GetFBytes1() []byte {
	if x != nil {
		return x.FBytes1
	}
	return nil
}

var File_tests_proto_cases_references_type_plain_proto protoreflect.FileDescriptor

var file_tests_proto_cases_references_type_plain_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x1a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x09,
	0x0a, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x31, 0x12, 0x29, 0x0a, 0x08,
	0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e,
	0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x0a, 0x02, 0x08, 0x01, 0x52, 0x07,
	0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x61, 0x12, 0x29, 0x0a, 0x08, 0x66, 0x5f, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a,
	0x08, 0x0a, 0x06, 0x0a, 0x04, 0x0a, 0x02, 0x08, 0x02, 0x52, 0x07, 0x66, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x62, 0x12, 0x29, 0x0a, 0x08, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04,
	0x12, 0x02, 0x08, 0x01, 0x52, 0x07, 0x66, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x61, 0x12, 0x29, 0x0a,
	0x08, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x12, 0x02, 0x08, 0x02, 0x52,
	0x07, 0x66, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x62, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0x8a, 0xa2, 0x1f,
	0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x3a, 0x02, 0x08, 0x01, 0x52, 0x08, 0x66, 0x55, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x61, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08,
	0x0a, 0x06, 0x0a, 0x04, 0x3a, 0x02, 0x08, 0x02, 0x52, 0x08, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x62, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a,
	0x04, 0x42, 0x02, 0x08, 0x01, 0x52, 0x08, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x61, 0x12,
	0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x62, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x42, 0x02,
	0x08, 0x02, 0x52, 0x08, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x62, 0x12, 0x2b, 0x0a, 0x09,
	0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x11, 0x42,
	0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x1a, 0x02, 0x08, 0x01, 0x52,
	0x08, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x61, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x73,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x42, 0x0e, 0x8a, 0xa2,
	0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x1a, 0x02, 0x08, 0x02, 0x52, 0x08, 0x66, 0x53,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x62, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x12, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a,
	0x08, 0x0a, 0x06, 0x0a, 0x04, 0x22, 0x02, 0x08, 0x01, 0x52, 0x08, 0x66, 0x53, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x61, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x62,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x12, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06,
	0x0a, 0x04, 0x22, 0x02, 0x08, 0x02, 0x52, 0x08, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x62,
	0x12, 0x2f, 0x0a, 0x0b, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x61, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0f, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a,
	0x04, 0x2a, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x61, 0x12, 0x2f, 0x0a, 0x0b, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x62,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0f, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06,
	0x0a, 0x04, 0x2a, 0x02, 0x08, 0x02, 0x52, 0x0a, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x62, 0x12, 0x2f, 0x0a, 0x0b, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x10, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a,
	0x06, 0x0a, 0x04, 0x32, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x61, 0x12, 0x2f, 0x0a, 0x0b, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36,
	0x34, 0x62, 0x18, 0x10, 0x20, 0x01, 0x28, 0x10, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08,
	0x0a, 0x06, 0x0a, 0x04, 0x32, 0x02, 0x08, 0x02, 0x52, 0x0a, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x62, 0x12, 0x2d, 0x0a, 0x0a, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x07, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08,
	0x0a, 0x06, 0x0a, 0x04, 0x4a, 0x02, 0x08, 0x01, 0x52, 0x09, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x61, 0x12, 0x2d, 0x0a, 0x0a, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x62, 0x18, 0x12, 0x20, 0x01, 0x28, 0x07, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a,
	0x06, 0x0a, 0x04, 0x4a, 0x02, 0x08, 0x02, 0x52, 0x09, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x62, 0x12, 0x2d, 0x0a, 0x0a, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x61,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x06, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06,
	0x0a, 0x04, 0x52, 0x02, 0x08, 0x01, 0x52, 0x09, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x61, 0x12, 0x2d, 0x0a, 0x0a, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x62, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x06, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a,
	0x04, 0x52, 0x02, 0x08, 0x02, 0x52, 0x09, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x62,
	0x12, 0x29, 0x0a, 0x08, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x31, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x02, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x5a, 0x02,
	0x08, 0x01, 0x52, 0x07, 0x66, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x31, 0x12, 0x29, 0x0a, 0x08, 0x66,
	0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x32, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0e, 0x8a,
	0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x5a, 0x02, 0x08, 0x02, 0x52, 0x07, 0x66,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x32, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a,
	0x08, 0x0a, 0x06, 0x0a, 0x04, 0x62, 0x02, 0x08, 0x01, 0x52, 0x08, 0x66, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x31, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x32,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06,
	0x0a, 0x04, 0x62, 0x02, 0x08, 0x02, 0x52, 0x08, 0x66, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x32,
	0x12, 0x27, 0x0a, 0x07, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x31, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x6a, 0x02, 0x08,
	0x01, 0x52, 0x06, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x31, 0x12, 0x27, 0x0a, 0x07, 0x66, 0x5f, 0x62,
	0x6f, 0x6f, 0x6c, 0x32, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a,
	0x5a, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x6a, 0x02, 0x08, 0x02, 0x52, 0x06, 0x66, 0x42, 0x6f, 0x6f,
	0x6c, 0x32, 0x22, 0x50, 0x0a, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x32,
	0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x31, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x07, 0x66, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x31, 0x42, 0x14, 0x5a, 0x12, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x72, 0x65, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tests_proto_cases_references_type_plain_proto_rawDescOnce sync.Once
	file_tests_proto_cases_references_type_plain_proto_rawDescData = file_tests_proto_cases_references_type_plain_proto_rawDesc
)

func file_tests_proto_cases_references_type_plain_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_references_type_plain_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_references_type_plain_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_references_type_plain_proto_rawDescData)
	})
	return file_tests_proto_cases_references_type_plain_proto_rawDescData
}

var file_tests_proto_cases_references_type_plain_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tests_proto_cases_references_type_plain_proto_goTypes = []interface{}{
	(*TypePlain1)(nil), // 0: type_plain.TypePlain1
	(*TypePlain2)(nil), // 1: type_plain.TypePlain2
}
var file_tests_proto_cases_references_type_plain_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_references_type_plain_proto_init() }
func file_tests_proto_cases_references_type_plain_proto_init() {
	if File_tests_proto_cases_references_type_plain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_references_type_plain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypePlain1); i {
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
		file_tests_proto_cases_references_type_plain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypePlain2); i {
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
			RawDescriptor: file_tests_proto_cases_references_type_plain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_references_type_plain_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_references_type_plain_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_references_type_plain_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_references_type_plain_proto = out.File
	file_tests_proto_cases_references_type_plain_proto_rawDesc = nil
	file_tests_proto_cases_references_type_plain_proto_goTypes = nil
	file_tests_proto_cases_references_type_plain_proto_depIdxs = nil
}

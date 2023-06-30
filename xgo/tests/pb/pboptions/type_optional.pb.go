// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/options/type_optional.proto

package pboptions

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	pbexternal "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// For test cases:
// 1. custom json key.
// 2. omitempty
type TypeOptional1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FInt32      *int32                 `protobuf:"varint,1,opt,name=f_int32,json=fInt32,proto3,oneof" json:"f_int32,omitempty"`
	FInt64      *int64                 `protobuf:"varint,2,opt,name=f_int64,json=fInt64,proto3,oneof" json:"f_int64,omitempty"`
	FUint32     *uint32                `protobuf:"varint,3,opt,name=f_uint32,json=fUint32,proto3,oneof" json:"f_uint32,omitempty"`
	FUint64     *uint64                `protobuf:"varint,4,opt,name=f_uint64,json=fUint64,proto3,oneof" json:"f_uint64,omitempty"`
	FSint32     *int32                 `protobuf:"zigzag32,5,opt,name=f_sint32,json=fSint32,proto3,oneof" json:"f_sint32,omitempty"`
	FSint64     *int64                 `protobuf:"zigzag64,6,opt,name=f_sint64,json=fSint64,proto3,oneof" json:"f_sint64,omitempty"`
	FSfixed32   *int32                 `protobuf:"fixed32,7,opt,name=f_sfixed32,json=fSfixed32,proto3,oneof" json:"f_sfixed32,omitempty"`
	FSfixed64   *int64                 `protobuf:"fixed64,8,opt,name=f_sfixed64,json=fSfixed64,proto3,oneof" json:"f_sfixed64,omitempty"`
	FFixed32    *uint32                `protobuf:"fixed32,9,opt,name=f_fixed32,json=fFixed32,proto3,oneof" json:"f_fixed32,omitempty"`
	FFixed64    *uint64                `protobuf:"fixed64,10,opt,name=f_fixed64,json=fFixed64,proto3,oneof" json:"f_fixed64,omitempty"`
	FFloat      *float32               `protobuf:"fixed32,11,opt,name=f_float,json=fFloat,proto3,oneof" json:"f_float,omitempty"`
	FDouble     *float64               `protobuf:"fixed64,12,opt,name=f_double,json=fDouble,proto3,oneof" json:"f_double,omitempty"`
	FBool1      *bool                  `protobuf:"varint,21,opt,name=f_bool1,json=fBool1,proto3,oneof" json:"f_bool1,omitempty"`
	FString1    *string                `protobuf:"bytes,22,opt,name=f_string1,json=fString1,proto3,oneof" json:"f_string1,omitempty"`
	FBytes1     []byte                 `protobuf:"bytes,23,opt,name=f_bytes1,json=fBytes1,proto3,oneof" json:"f_bytes1,omitempty"`
	FEnum1      *pbexternal.Enum1      `protobuf:"varint,31,opt,name=f_enum1,json=fEnum1,proto3,enum=external.Enum1,oneof" json:"f_enum1,omitempty"`
	FMessage1   *pbexternal.Message1   `protobuf:"bytes,32,opt,name=f_message1,json=fMessage1,proto3,oneof" json:"f_message1,omitempty"`
	FAny1       *anypb.Any             `protobuf:"bytes,41,opt,name=f_any1,json=fAny1,proto3,oneof" json:"f_any1,omitempty"`
	FDuration1  *durationpb.Duration   `protobuf:"bytes,42,opt,name=f_duration1,json=fDuration1,proto3,oneof" json:"f_duration1,omitempty"`
	FTimestamp1 *timestamppb.Timestamp `protobuf:"bytes,43,opt,name=f_timestamp1,json=fTimestamp1,proto3,oneof" json:"f_timestamp1,omitempty"`
}

func (x *TypeOptional1) Reset() {
	*x = TypeOptional1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_options_type_optional_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeOptional1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeOptional1) ProtoMessage() {}

func (x *TypeOptional1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_options_type_optional_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeOptional1.ProtoReflect.Descriptor instead.
func (*TypeOptional1) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_options_type_optional_proto_rawDescGZIP(), []int{0}
}

func (x *TypeOptional1) GetFInt32() int32 {
	if x != nil && x.FInt32 != nil {
		return *x.FInt32
	}
	return 0
}

func (x *TypeOptional1) GetFInt64() int64 {
	if x != nil && x.FInt64 != nil {
		return *x.FInt64
	}
	return 0
}

func (x *TypeOptional1) GetFUint32() uint32 {
	if x != nil && x.FUint32 != nil {
		return *x.FUint32
	}
	return 0
}

func (x *TypeOptional1) GetFUint64() uint64 {
	if x != nil && x.FUint64 != nil {
		return *x.FUint64
	}
	return 0
}

func (x *TypeOptional1) GetFSint32() int32 {
	if x != nil && x.FSint32 != nil {
		return *x.FSint32
	}
	return 0
}

func (x *TypeOptional1) GetFSint64() int64 {
	if x != nil && x.FSint64 != nil {
		return *x.FSint64
	}
	return 0
}

func (x *TypeOptional1) GetFSfixed32() int32 {
	if x != nil && x.FSfixed32 != nil {
		return *x.FSfixed32
	}
	return 0
}

func (x *TypeOptional1) GetFSfixed64() int64 {
	if x != nil && x.FSfixed64 != nil {
		return *x.FSfixed64
	}
	return 0
}

func (x *TypeOptional1) GetFFixed32() uint32 {
	if x != nil && x.FFixed32 != nil {
		return *x.FFixed32
	}
	return 0
}

func (x *TypeOptional1) GetFFixed64() uint64 {
	if x != nil && x.FFixed64 != nil {
		return *x.FFixed64
	}
	return 0
}

func (x *TypeOptional1) GetFFloat() float32 {
	if x != nil && x.FFloat != nil {
		return *x.FFloat
	}
	return 0
}

func (x *TypeOptional1) GetFDouble() float64 {
	if x != nil && x.FDouble != nil {
		return *x.FDouble
	}
	return 0
}

func (x *TypeOptional1) GetFBool1() bool {
	if x != nil && x.FBool1 != nil {
		return *x.FBool1
	}
	return false
}

func (x *TypeOptional1) GetFString1() string {
	if x != nil && x.FString1 != nil {
		return *x.FString1
	}
	return ""
}

func (x *TypeOptional1) GetFBytes1() []byte {
	if x != nil {
		return x.FBytes1
	}
	return nil
}

func (x *TypeOptional1) GetFEnum1() pbexternal.Enum1 {
	if x != nil && x.FEnum1 != nil {
		return *x.FEnum1
	}
	return pbexternal.Enum1(0)
}

func (x *TypeOptional1) GetFMessage1() *pbexternal.Message1 {
	if x != nil {
		return x.FMessage1
	}
	return nil
}

func (x *TypeOptional1) GetFAny1() *anypb.Any {
	if x != nil {
		return x.FAny1
	}
	return nil
}

func (x *TypeOptional1) GetFDuration1() *durationpb.Duration {
	if x != nil {
		return x.FDuration1
	}
	return nil
}

func (x *TypeOptional1) GetFTimestamp1() *timestamppb.Timestamp {
	if x != nil {
		return x.FTimestamp1
	}
	return nil
}

var File_tests_proto_cases_options_type_optional_proto protoreflect.FileDescriptor

var file_tests_proto_cases_options_type_optional_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x91, 0x0b, 0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x31, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a, 0x07, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x01, 0x48, 0x00, 0x52, 0x06, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x88, 0x01, 0x01,
	0x12, 0x2d, 0x0a, 0x07, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a, 0x07, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x18, 0x01, 0x48, 0x01, 0x52, 0x06, 0x66, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x88, 0x01, 0x01, 0x12,
	0x30, 0x0a, 0x08, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x01, 0x48, 0x02, 0x52, 0x07, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x88, 0x01,
	0x01, 0x12, 0x30, 0x0a, 0x08, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x01, 0x48, 0x03, 0x52, 0x07, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x08, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x11, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x73,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x48, 0x04, 0x52, 0x07, 0x66, 0x53, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x08, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74,
	0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x01, 0x48, 0x05, 0x52, 0x07, 0x66, 0x53, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x0a, 0x66, 0x5f, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0f, 0x42, 0x12, 0x8a, 0xa2, 0x1f,
	0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x01, 0x48,
	0x06, 0x52, 0x09, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x88, 0x01, 0x01, 0x12,
	0x36, 0x0a, 0x0a, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x10, 0x42, 0x12, 0x8a, 0xa2, 0x1f, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x01, 0x48, 0x07, 0x52, 0x09, 0x66, 0x53, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x5f, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x07, 0x42, 0x11, 0x8a, 0xa2, 0x1f, 0x0d,
	0x0a, 0x09, 0x74, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x01, 0x48, 0x08, 0x52,
	0x08, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x09,
	0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x06, 0x42,
	0x11, 0x8a, 0xa2, 0x1f, 0x0d, 0x0a, 0x09, 0x74, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x18, 0x01, 0x48, 0x09, 0x52, 0x08, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x88, 0x01,
	0x01, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x02, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a, 0x07, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x18, 0x01, 0x48, 0x0a, 0x52, 0x06, 0x66, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x30, 0x0a, 0x08, 0x66, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x01, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x48, 0x0b, 0x52, 0x07, 0x66, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x31, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a, 0x07, 0x74, 0x5f, 0x62, 0x6f, 0x6f,
	0x6c, 0x31, 0x18, 0x01, 0x48, 0x0c, 0x52, 0x06, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x31, 0x88, 0x01,
	0x01, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x8a, 0xa2, 0x1f, 0x0d, 0x0a, 0x09, 0x74, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x48, 0x0d, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x31, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x08, 0x66, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08,
	0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x31, 0x18, 0x01, 0x48, 0x0e, 0x52, 0x07, 0x66, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x31, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x07, 0x66, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x31, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b,
	0x0a, 0x07, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x48, 0x0f, 0x52, 0x06, 0x66,
	0x45, 0x6e, 0x75, 0x6d, 0x31, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x0a, 0x66, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31,
	0x42, 0x12, 0x8a, 0xa2, 0x1f, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x18, 0x01, 0x48, 0x10, 0x52, 0x09, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x06, 0x66, 0x5f, 0x61, 0x6e, 0x79, 0x31, 0x18, 0x29,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a,
	0x0a, 0x06, 0x74, 0x5f, 0x61, 0x6e, 0x79, 0x31, 0x18, 0x01, 0x48, 0x11, 0x52, 0x05, 0x66, 0x41,
	0x6e, 0x79, 0x31, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a, 0x0b, 0x66, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x31, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x13, 0x8a, 0xa2, 0x1f, 0x0f, 0x0a, 0x0b, 0x74, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x18, 0x01, 0x48, 0x12, 0x52, 0x0a, 0x66,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x88, 0x01, 0x01, 0x12, 0x58, 0x0a, 0x0c,
	0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x31, 0x18, 0x2b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14,
	0x8a, 0xa2, 0x1f, 0x10, 0x0a, 0x0c, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x31, 0x18, 0x01, 0x48, 0x13, 0x52, 0x0b, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x31, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x5f, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x5f, 0x73,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x5f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c,
	0x31, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x31, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x31, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x5f, 0x61, 0x6e,
	0x79, 0x31, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x31, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x31, 0x42, 0x18, 0x5a, 0x16, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_cases_options_type_optional_proto_rawDescOnce sync.Once
	file_tests_proto_cases_options_type_optional_proto_rawDescData = file_tests_proto_cases_options_type_optional_proto_rawDesc
)

func file_tests_proto_cases_options_type_optional_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_options_type_optional_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_options_type_optional_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_options_type_optional_proto_rawDescData)
	})
	return file_tests_proto_cases_options_type_optional_proto_rawDescData
}

var file_tests_proto_cases_options_type_optional_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_proto_cases_options_type_optional_proto_goTypes = []interface{}{
	(*TypeOptional1)(nil),         // 0: type_optional.TypeOptional1
	(pbexternal.Enum1)(0),         // 1: external.Enum1
	(*pbexternal.Message1)(nil),   // 2: external.Message1
	(*anypb.Any)(nil),             // 3: google.protobuf.Any
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_tests_proto_cases_options_type_optional_proto_depIdxs = []int32{
	1, // 0: type_optional.TypeOptional1.f_enum1:type_name -> external.Enum1
	2, // 1: type_optional.TypeOptional1.f_message1:type_name -> external.Message1
	3, // 2: type_optional.TypeOptional1.f_any1:type_name -> google.protobuf.Any
	4, // 3: type_optional.TypeOptional1.f_duration1:type_name -> google.protobuf.Duration
	5, // 4: type_optional.TypeOptional1.f_timestamp1:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_options_type_optional_proto_init() }
func file_tests_proto_cases_options_type_optional_proto_init() {
	if File_tests_proto_cases_options_type_optional_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_options_type_optional_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeOptional1); i {
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
	file_tests_proto_cases_options_type_optional_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_proto_cases_options_type_optional_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_options_type_optional_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_options_type_optional_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_options_type_optional_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_options_type_optional_proto = out.File
	file_tests_proto_cases_options_type_optional_proto_rawDesc = nil
	file_tests_proto_cases_options_type_optional_proto_goTypes = nil
	file_tests_proto_cases_options_type_optional_proto_depIdxs = nil
}

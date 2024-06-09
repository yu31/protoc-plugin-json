// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/options/type_map.proto

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
type TypeMap1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FInt32      map[int32]int32                   `protobuf:"bytes,1,rep,name=f_int32,json=fInt32,proto3" json:"f_int32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FInt64      map[int64]int64                   `protobuf:"bytes,2,rep,name=f_int64,json=fInt64,proto3" json:"f_int64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FUint32     map[uint32]uint32                 `protobuf:"bytes,3,rep,name=f_uint32,json=fUint32,proto3" json:"f_uint32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FUint64     map[uint64]uint64                 `protobuf:"bytes,4,rep,name=f_uint64,json=fUint64,proto3" json:"f_uint64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FSint32     map[int32]int32                   `protobuf:"bytes,5,rep,name=f_sint32,json=fSint32,proto3" json:"f_sint32,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"zigzag32,2,opt,name=value,proto3"`
	FSint64     map[int64]int64                   `protobuf:"bytes,6,rep,name=f_sint64,json=fSint64,proto3" json:"f_sint64,omitempty" protobuf_key:"zigzag64,1,opt,name=key,proto3" protobuf_val:"zigzag64,2,opt,name=value,proto3"`
	FSfixed32   map[int32]int32                   `protobuf:"bytes,7,rep,name=f_sfixed32,json=fSfixed32,proto3" json:"f_sfixed32,omitempty" protobuf_key:"fixed32,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	FSfixed64   map[int64]int64                   `protobuf:"bytes,8,rep,name=f_sfixed64,json=fSfixed64,proto3" json:"f_sfixed64,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	FFixed32    map[uint32]uint32                 `protobuf:"bytes,9,rep,name=f_fixed32,json=fFixed32,proto3" json:"f_fixed32,omitempty" protobuf_key:"fixed32,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	FFixed64    map[uint64]uint64                 `protobuf:"bytes,10,rep,name=f_fixed64,json=fFixed64,proto3" json:"f_fixed64,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	FFloat      map[string]float32                `protobuf:"bytes,11,rep,name=f_float,json=fFloat,proto3" json:"f_float,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	FDouble     map[string]float64                `protobuf:"bytes,12,rep,name=f_double,json=fDouble,proto3" json:"f_double,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	FBool1      map[bool]bool                     `protobuf:"bytes,21,rep,name=f_bool1,json=fBool1,proto3" json:"f_bool1,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FString1    map[string]string                 `protobuf:"bytes,22,rep,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FBytes1     map[string][]byte                 `protobuf:"bytes,23,rep,name=f_bytes1,json=fBytes1,proto3" json:"f_bytes1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FEnum1      map[string]pbexternal.EnumNum1    `protobuf:"bytes,31,rep,name=f_enum1,json=fEnum1,proto3" json:"f_enum1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=external.EnumNum1"`
	FMessage1   map[string]*pbexternal.Message1   `protobuf:"bytes,32,rep,name=f_message1,json=fMessage1,proto3" json:"f_message1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FAny1       map[string]*anypb.Any             `protobuf:"bytes,41,rep,name=f_any1,json=fAny1,proto3" json:"f_any1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FDuration1  map[string]*durationpb.Duration   `protobuf:"bytes,42,rep,name=f_duration1,json=fDuration1,proto3" json:"f_duration1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FTimestamp1 map[string]*timestamppb.Timestamp `protobuf:"bytes,43,rep,name=f_timestamp1,json=fTimestamp1,proto3" json:"f_timestamp1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TypeMap1) Reset() {
	*x = TypeMap1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_options_type_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeMap1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeMap1) ProtoMessage() {}

func (x *TypeMap1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_options_type_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeMap1.ProtoReflect.Descriptor instead.
func (*TypeMap1) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_options_type_map_proto_rawDescGZIP(), []int{0}
}

func (x *TypeMap1) GetFInt32() map[int32]int32 {
	if x != nil {
		return x.FInt32
	}
	return nil
}

func (x *TypeMap1) GetFInt64() map[int64]int64 {
	if x != nil {
		return x.FInt64
	}
	return nil
}

func (x *TypeMap1) GetFUint32() map[uint32]uint32 {
	if x != nil {
		return x.FUint32
	}
	return nil
}

func (x *TypeMap1) GetFUint64() map[uint64]uint64 {
	if x != nil {
		return x.FUint64
	}
	return nil
}

func (x *TypeMap1) GetFSint32() map[int32]int32 {
	if x != nil {
		return x.FSint32
	}
	return nil
}

func (x *TypeMap1) GetFSint64() map[int64]int64 {
	if x != nil {
		return x.FSint64
	}
	return nil
}

func (x *TypeMap1) GetFSfixed32() map[int32]int32 {
	if x != nil {
		return x.FSfixed32
	}
	return nil
}

func (x *TypeMap1) GetFSfixed64() map[int64]int64 {
	if x != nil {
		return x.FSfixed64
	}
	return nil
}

func (x *TypeMap1) GetFFixed32() map[uint32]uint32 {
	if x != nil {
		return x.FFixed32
	}
	return nil
}

func (x *TypeMap1) GetFFixed64() map[uint64]uint64 {
	if x != nil {
		return x.FFixed64
	}
	return nil
}

func (x *TypeMap1) GetFFloat() map[string]float32 {
	if x != nil {
		return x.FFloat
	}
	return nil
}

func (x *TypeMap1) GetFDouble() map[string]float64 {
	if x != nil {
		return x.FDouble
	}
	return nil
}

func (x *TypeMap1) GetFBool1() map[bool]bool {
	if x != nil {
		return x.FBool1
	}
	return nil
}

func (x *TypeMap1) GetFString1() map[string]string {
	if x != nil {
		return x.FString1
	}
	return nil
}

func (x *TypeMap1) GetFBytes1() map[string][]byte {
	if x != nil {
		return x.FBytes1
	}
	return nil
}

func (x *TypeMap1) GetFEnum1() map[string]pbexternal.EnumNum1 {
	if x != nil {
		return x.FEnum1
	}
	return nil
}

func (x *TypeMap1) GetFMessage1() map[string]*pbexternal.Message1 {
	if x != nil {
		return x.FMessage1
	}
	return nil
}

func (x *TypeMap1) GetFAny1() map[string]*anypb.Any {
	if x != nil {
		return x.FAny1
	}
	return nil
}

func (x *TypeMap1) GetFDuration1() map[string]*durationpb.Duration {
	if x != nil {
		return x.FDuration1
	}
	return nil
}

func (x *TypeMap1) GetFTimestamp1() map[string]*timestamppb.Timestamp {
	if x != nil {
		return x.FTimestamp1
	}
	return nil
}

var File_tests_proto_cases_options_type_map_proto protoreflect.FileDescriptor

var file_tests_proto_cases_options_type_map_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x16, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70,
	0x31, 0x12, 0x47, 0x0a, 0x07, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a, 0x07, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x01, 0x52, 0x06, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x47, 0x0a, 0x07, 0x66, 0x5f,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b,
	0x0a, 0x07, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x01, 0x52, 0x06, 0x66, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x12, 0x4b, 0x0a, 0x08, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x52, 0x07, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x4b, 0x0a, 0x08, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0x01, 0x52, 0x07, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x4b, 0x0a,
	0x08, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61,
	0x70, 0x31, 0x2e, 0x46, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0x01, 0x52, 0x07, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x4b, 0x0a, 0x08, 0x66, 0x5f,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e,
	0x46, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x10, 0x8a, 0xa2,
	0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x01, 0x52, 0x07,
	0x66, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x53, 0x0a, 0x0a, 0x66, 0x5f, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46,
	0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x12, 0x8a,
	0xa2, 0x1f, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18,
	0x01, 0x52, 0x09, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x53, 0x0a, 0x0a,
	0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d,
	0x61, 0x70, 0x31, 0x2e, 0x46, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x12, 0x8a, 0xa2, 0x1f, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x18, 0x01, 0x52, 0x09, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36,
	0x34, 0x12, 0x4f, 0x0a, 0x09, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x11, 0x8a, 0xa2, 0x1f, 0x0d, 0x0a, 0x09, 0x74, 0x5f, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x01, 0x52, 0x08, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x12, 0x4f, 0x0a, 0x09, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36,
	0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x11, 0x8a, 0xa2, 0x1f, 0x0d, 0x0a, 0x09, 0x74, 0x5f,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x01, 0x52, 0x08, 0x66, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x12, 0x47, 0x0a, 0x07, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a, 0x07, 0x74, 0x5f, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x18, 0x01, 0x52, 0x06, 0x66, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x4b, 0x0a, 0x08,
	0x66, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70,
	0x31, 0x2e, 0x46, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x10,
	0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x52, 0x07, 0x66, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x66, 0x5f, 0x62,
	0x6f, 0x6f, 0x6c, 0x31, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x42,
	0x6f, 0x6f, 0x6c, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a,
	0x07, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x31, 0x18, 0x01, 0x52, 0x06, 0x66, 0x42, 0x6f, 0x6f,
	0x6c, 0x31, 0x12, 0x4f, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18,
	0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x11, 0x8a, 0xa2, 0x1f, 0x0d, 0x0a, 0x09, 0x74, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x31, 0x12, 0x4b, 0x0a, 0x08, 0x66, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x31, 0x18,
	0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x42, 0x79, 0x74, 0x65, 0x73, 0x31,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x10, 0x8a, 0xa2, 0x1f, 0x0c, 0x0a, 0x08, 0x74, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x31, 0x18, 0x01, 0x52, 0x07, 0x66, 0x42, 0x79, 0x74, 0x65, 0x73, 0x31,
	0x12, 0x47, 0x0a, 0x07, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x1f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x0a, 0x07, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x31, 0x18,
	0x01, 0x52, 0x06, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x53, 0x0a, 0x0a, 0x66, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31,
	0x2e, 0x46, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x12, 0x8a, 0xa2, 0x1f, 0x0e, 0x0a, 0x0a, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x18, 0x01, 0x52, 0x09, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x43,
	0x0a, 0x06, 0x66, 0x5f, 0x61, 0x6e, 0x79, 0x31, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70,
	0x31, 0x2e, 0x46, 0x41, 0x6e, 0x79, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0e, 0x8a, 0xa2,
	0x1f, 0x0a, 0x0a, 0x06, 0x74, 0x5f, 0x61, 0x6e, 0x79, 0x31, 0x18, 0x01, 0x52, 0x05, 0x66, 0x41,
	0x6e, 0x79, 0x31, 0x12, 0x57, 0x0a, 0x0b, 0x66, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x31, 0x18, 0x2a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x13, 0x8a, 0xa2, 0x1f,
	0x0f, 0x0a, 0x0b, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x18, 0x01,
	0x52, 0x0a, 0x66, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x12, 0x5b, 0x0a, 0x0c,
	0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x31, 0x18, 0x2b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x4d, 0x61, 0x70, 0x31, 0x2e, 0x46, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x14, 0x8a, 0xa2, 0x1f, 0x10, 0x0a, 0x0c, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x31, 0x18, 0x01, 0x52, 0x0b, 0x66, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x31, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3a, 0x0a, 0x0c, 0x46, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46,
	0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x53, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3c, 0x0a, 0x0e, 0x46, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a,
	0x0e, 0x46, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x10, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x10,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x46,
	0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x46, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b,
	0x46, 0x42, 0x6f, 0x6f, 0x6c, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x42, 0x79, 0x74, 0x65, 0x73, 0x31, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x4d, 0x0a, 0x0b, 0x46, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x4e, 0x75, 0x6d, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x50, 0x0a, 0x0e, 0x46, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x4e, 0x0a, 0x0a, 0x46, 0x41, 0x6e, 0x79, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x58, 0x0a, 0x0f, 0x46, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5a, 0x0a, 0x10, 0x46,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x18, 0x5a, 0x16, 0x78, 0x67, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_cases_options_type_map_proto_rawDescOnce sync.Once
	file_tests_proto_cases_options_type_map_proto_rawDescData = file_tests_proto_cases_options_type_map_proto_rawDesc
)

func file_tests_proto_cases_options_type_map_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_options_type_map_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_options_type_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_options_type_map_proto_rawDescData)
	})
	return file_tests_proto_cases_options_type_map_proto_rawDescData
}

var file_tests_proto_cases_options_type_map_proto_msgTypes = make([]protoimpl.MessageInfo, 21)
var file_tests_proto_cases_options_type_map_proto_goTypes = []interface{}{
	(*TypeMap1)(nil),              // 0: options.TypeMap1
	nil,                           // 1: options.TypeMap1.FInt32Entry
	nil,                           // 2: options.TypeMap1.FInt64Entry
	nil,                           // 3: options.TypeMap1.FUint32Entry
	nil,                           // 4: options.TypeMap1.FUint64Entry
	nil,                           // 5: options.TypeMap1.FSint32Entry
	nil,                           // 6: options.TypeMap1.FSint64Entry
	nil,                           // 7: options.TypeMap1.FSfixed32Entry
	nil,                           // 8: options.TypeMap1.FSfixed64Entry
	nil,                           // 9: options.TypeMap1.FFixed32Entry
	nil,                           // 10: options.TypeMap1.FFixed64Entry
	nil,                           // 11: options.TypeMap1.FFloatEntry
	nil,                           // 12: options.TypeMap1.FDoubleEntry
	nil,                           // 13: options.TypeMap1.FBool1Entry
	nil,                           // 14: options.TypeMap1.FString1Entry
	nil,                           // 15: options.TypeMap1.FBytes1Entry
	nil,                           // 16: options.TypeMap1.FEnum1Entry
	nil,                           // 17: options.TypeMap1.FMessage1Entry
	nil,                           // 18: options.TypeMap1.FAny1Entry
	nil,                           // 19: options.TypeMap1.FDuration1Entry
	nil,                           // 20: options.TypeMap1.FTimestamp1Entry
	(pbexternal.EnumNum1)(0),      // 21: external.EnumNum1
	(*pbexternal.Message1)(nil),   // 22: external.Message1
	(*anypb.Any)(nil),             // 23: google.protobuf.Any
	(*durationpb.Duration)(nil),   // 24: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 25: google.protobuf.Timestamp
}
var file_tests_proto_cases_options_type_map_proto_depIdxs = []int32{
	1,  // 0: options.TypeMap1.f_int32:type_name -> options.TypeMap1.FInt32Entry
	2,  // 1: options.TypeMap1.f_int64:type_name -> options.TypeMap1.FInt64Entry
	3,  // 2: options.TypeMap1.f_uint32:type_name -> options.TypeMap1.FUint32Entry
	4,  // 3: options.TypeMap1.f_uint64:type_name -> options.TypeMap1.FUint64Entry
	5,  // 4: options.TypeMap1.f_sint32:type_name -> options.TypeMap1.FSint32Entry
	6,  // 5: options.TypeMap1.f_sint64:type_name -> options.TypeMap1.FSint64Entry
	7,  // 6: options.TypeMap1.f_sfixed32:type_name -> options.TypeMap1.FSfixed32Entry
	8,  // 7: options.TypeMap1.f_sfixed64:type_name -> options.TypeMap1.FSfixed64Entry
	9,  // 8: options.TypeMap1.f_fixed32:type_name -> options.TypeMap1.FFixed32Entry
	10, // 9: options.TypeMap1.f_fixed64:type_name -> options.TypeMap1.FFixed64Entry
	11, // 10: options.TypeMap1.f_float:type_name -> options.TypeMap1.FFloatEntry
	12, // 11: options.TypeMap1.f_double:type_name -> options.TypeMap1.FDoubleEntry
	13, // 12: options.TypeMap1.f_bool1:type_name -> options.TypeMap1.FBool1Entry
	14, // 13: options.TypeMap1.f_string1:type_name -> options.TypeMap1.FString1Entry
	15, // 14: options.TypeMap1.f_bytes1:type_name -> options.TypeMap1.FBytes1Entry
	16, // 15: options.TypeMap1.f_enum1:type_name -> options.TypeMap1.FEnum1Entry
	17, // 16: options.TypeMap1.f_message1:type_name -> options.TypeMap1.FMessage1Entry
	18, // 17: options.TypeMap1.f_any1:type_name -> options.TypeMap1.FAny1Entry
	19, // 18: options.TypeMap1.f_duration1:type_name -> options.TypeMap1.FDuration1Entry
	20, // 19: options.TypeMap1.f_timestamp1:type_name -> options.TypeMap1.FTimestamp1Entry
	21, // 20: options.TypeMap1.FEnum1Entry.value:type_name -> external.EnumNum1
	22, // 21: options.TypeMap1.FMessage1Entry.value:type_name -> external.Message1
	23, // 22: options.TypeMap1.FAny1Entry.value:type_name -> google.protobuf.Any
	24, // 23: options.TypeMap1.FDuration1Entry.value:type_name -> google.protobuf.Duration
	25, // 24: options.TypeMap1.FTimestamp1Entry.value:type_name -> google.protobuf.Timestamp
	25, // [25:25] is the sub-list for method output_type
	25, // [25:25] is the sub-list for method input_type
	25, // [25:25] is the sub-list for extension type_name
	25, // [25:25] is the sub-list for extension extendee
	0,  // [0:25] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_options_type_map_proto_init() }
func file_tests_proto_cases_options_type_map_proto_init() {
	if File_tests_proto_cases_options_type_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_options_type_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeMap1); i {
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
			RawDescriptor: file_tests_proto_cases_options_type_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   21,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_options_type_map_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_options_type_map_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_options_type_map_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_options_type_map_proto = out.File
	file_tests_proto_cases_options_type_map_proto_rawDesc = nil
	file_tests_proto_cases_options_type_map_proto_goTypes = nil
	file_tests_proto_cases_options_type_map_proto_depIdxs = nil
}

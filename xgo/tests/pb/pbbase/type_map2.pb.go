// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/base/type_map2.proto

package pbbase

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

// For test all type of map key.
type TypeMap2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1  map[string]string `protobuf:"bytes,1,rep,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FInt32    map[int32]int32   `protobuf:"bytes,11,rep,name=f_int32,json=fInt32,proto3" json:"f_int32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FInt64    map[int64]int64   `protobuf:"bytes,12,rep,name=f_int64,json=fInt64,proto3" json:"f_int64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FUint32   map[uint32]uint32 `protobuf:"bytes,13,rep,name=f_uint32,json=fUint32,proto3" json:"f_uint32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FUint64   map[uint64]uint64 `protobuf:"bytes,14,rep,name=f_uint64,json=fUint64,proto3" json:"f_uint64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FSint32   map[int32]int32   `protobuf:"bytes,15,rep,name=f_sint32,json=fSint32,proto3" json:"f_sint32,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"zigzag32,2,opt,name=value,proto3"`
	FSint64   map[int64]int64   `protobuf:"bytes,16,rep,name=f_sint64,json=fSint64,proto3" json:"f_sint64,omitempty" protobuf_key:"zigzag64,1,opt,name=key,proto3" protobuf_val:"zigzag64,2,opt,name=value,proto3"`
	FSfixed32 map[int32]int32   `protobuf:"bytes,17,rep,name=f_sfixed32,json=fSfixed32,proto3" json:"f_sfixed32,omitempty" protobuf_key:"fixed32,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	FSfixed64 map[int64]int64   `protobuf:"bytes,18,rep,name=f_sfixed64,json=fSfixed64,proto3" json:"f_sfixed64,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	FFixed32  map[uint32]uint32 `protobuf:"bytes,19,rep,name=f_fixed32,json=fFixed32,proto3" json:"f_fixed32,omitempty" protobuf_key:"fixed32,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	FFixed64  map[uint64]uint64 `protobuf:"bytes,20,rep,name=f_fixed64,json=fFixed64,proto3" json:"f_fixed64,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *TypeMap2) Reset() {
	*x = TypeMap2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_base_type_map2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeMap2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeMap2) ProtoMessage() {}

func (x *TypeMap2) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_base_type_map2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeMap2.ProtoReflect.Descriptor instead.
func (*TypeMap2) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_base_type_map2_proto_rawDescGZIP(), []int{0}
}

func (x *TypeMap2) GetFString1() map[string]string {
	if x != nil {
		return x.FString1
	}
	return nil
}

func (x *TypeMap2) GetFInt32() map[int32]int32 {
	if x != nil {
		return x.FInt32
	}
	return nil
}

func (x *TypeMap2) GetFInt64() map[int64]int64 {
	if x != nil {
		return x.FInt64
	}
	return nil
}

func (x *TypeMap2) GetFUint32() map[uint32]uint32 {
	if x != nil {
		return x.FUint32
	}
	return nil
}

func (x *TypeMap2) GetFUint64() map[uint64]uint64 {
	if x != nil {
		return x.FUint64
	}
	return nil
}

func (x *TypeMap2) GetFSint32() map[int32]int32 {
	if x != nil {
		return x.FSint32
	}
	return nil
}

func (x *TypeMap2) GetFSint64() map[int64]int64 {
	if x != nil {
		return x.FSint64
	}
	return nil
}

func (x *TypeMap2) GetFSfixed32() map[int32]int32 {
	if x != nil {
		return x.FSfixed32
	}
	return nil
}

func (x *TypeMap2) GetFSfixed64() map[int64]int64 {
	if x != nil {
		return x.FSfixed64
	}
	return nil
}

func (x *TypeMap2) GetFFixed32() map[uint32]uint32 {
	if x != nil {
		return x.FFixed32
	}
	return nil
}

func (x *TypeMap2) GetFFixed64() map[uint64]uint64 {
	if x != nil {
		return x.FFixed64
	}
	return nil
}

var File_tests_proto_cases_base_type_map2_proto protoreflect.FileDescriptor

var file_tests_proto_cases_base_type_map2_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x61,
	0x70, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdc, 0x0a, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32, 0x12, 0x3f, 0x0a,
	0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32,
	0x2e, 0x46, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04,
	0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x39,
	0x0a, 0x07, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32, 0x2e,
	0x46, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x8a, 0xa2, 0x1f,
	0x00, 0x52, 0x06, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x39, 0x0a, 0x07, 0x66, 0x5f, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32, 0x2e, 0x46, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x06, 0x66, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x12, 0x3c, 0x0a, 0x08, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x4d, 0x61, 0x70, 0x32, 0x2e, 0x46, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x07, 0x66, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x12, 0x3c, 0x0a, 0x08, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x4d, 0x61, 0x70, 0x32, 0x2e, 0x46, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x07, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x3c, 0x0a, 0x08, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61,
	0x70, 0x32, 0x2e, 0x46, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x07, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x3c,
	0x0a, 0x08, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32,
	0x2e, 0x46, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x8a,
	0xa2, 0x1f, 0x00, 0x52, 0x07, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x42, 0x0a, 0x0a,
	0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32,
	0x2e, 0x46, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x09, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x12, 0x42, 0x0a, 0x0a, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x12,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x4d, 0x61, 0x70, 0x32, 0x2e, 0x46, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x09, 0x66, 0x53, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x12, 0x3f, 0x0a, 0x09, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32, 0x2e, 0x46, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x3f, 0x0a, 0x09, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x32, 0x2e, 0x46, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36,
	0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39,
	0x0a, 0x0b, 0x46, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x55, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a,
	0x0c, 0x46, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x46, 0x53, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x46, 0x53, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x10, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x10, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x07, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x46, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x15, 0x5a, 0x13, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_cases_base_type_map2_proto_rawDescOnce sync.Once
	file_tests_proto_cases_base_type_map2_proto_rawDescData = file_tests_proto_cases_base_type_map2_proto_rawDesc
)

func file_tests_proto_cases_base_type_map2_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_base_type_map2_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_base_type_map2_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_base_type_map2_proto_rawDescData)
	})
	return file_tests_proto_cases_base_type_map2_proto_rawDescData
}

var file_tests_proto_cases_base_type_map2_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_tests_proto_cases_base_type_map2_proto_goTypes = []interface{}{
	(*TypeMap2)(nil), // 0: base.TypeMap2
	nil,              // 1: base.TypeMap2.FString1Entry
	nil,              // 2: base.TypeMap2.FInt32Entry
	nil,              // 3: base.TypeMap2.FInt64Entry
	nil,              // 4: base.TypeMap2.FUint32Entry
	nil,              // 5: base.TypeMap2.FUint64Entry
	nil,              // 6: base.TypeMap2.FSint32Entry
	nil,              // 7: base.TypeMap2.FSint64Entry
	nil,              // 8: base.TypeMap2.FSfixed32Entry
	nil,              // 9: base.TypeMap2.FSfixed64Entry
	nil,              // 10: base.TypeMap2.FFixed32Entry
	nil,              // 11: base.TypeMap2.FFixed64Entry
}
var file_tests_proto_cases_base_type_map2_proto_depIdxs = []int32{
	1,  // 0: base.TypeMap2.f_string1:type_name -> base.TypeMap2.FString1Entry
	2,  // 1: base.TypeMap2.f_int32:type_name -> base.TypeMap2.FInt32Entry
	3,  // 2: base.TypeMap2.f_int64:type_name -> base.TypeMap2.FInt64Entry
	4,  // 3: base.TypeMap2.f_uint32:type_name -> base.TypeMap2.FUint32Entry
	5,  // 4: base.TypeMap2.f_uint64:type_name -> base.TypeMap2.FUint64Entry
	6,  // 5: base.TypeMap2.f_sint32:type_name -> base.TypeMap2.FSint32Entry
	7,  // 6: base.TypeMap2.f_sint64:type_name -> base.TypeMap2.FSint64Entry
	8,  // 7: base.TypeMap2.f_sfixed32:type_name -> base.TypeMap2.FSfixed32Entry
	9,  // 8: base.TypeMap2.f_sfixed64:type_name -> base.TypeMap2.FSfixed64Entry
	10, // 9: base.TypeMap2.f_fixed32:type_name -> base.TypeMap2.FFixed32Entry
	11, // 10: base.TypeMap2.f_fixed64:type_name -> base.TypeMap2.FFixed64Entry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_base_type_map2_proto_init() }
func file_tests_proto_cases_base_type_map2_proto_init() {
	if File_tests_proto_cases_base_type_map2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_base_type_map2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeMap2); i {
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
			RawDescriptor: file_tests_proto_cases_base_type_map2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_base_type_map2_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_base_type_map2_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_base_type_map2_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_base_type_map2_proto = out.File
	file_tests_proto_cases_base_type_map2_proto_rawDesc = nil
	file_tests_proto_cases_base_type_map2_proto_goTypes = nil
	file_tests_proto_cases_base_type_map2_proto_depIdxs = nil
}

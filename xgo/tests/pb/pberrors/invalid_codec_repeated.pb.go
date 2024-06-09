// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/cases/errors/invalid_codec_repeated.proto

package pberrors

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

type InvalidCodecRepeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FInt32Numeric    []int32   `protobuf:"varint,1,rep,packed,name=f_int32_numeric,json=fInt32Numeric,proto3" json:"f_int32_numeric,omitempty"`
	FInt32String     []int32   `protobuf:"varint,2,rep,packed,name=f_int32_string,json=fInt32String,proto3" json:"f_int32_string,omitempty"`
	FInt64Numeric    []int64   `protobuf:"varint,3,rep,packed,name=f_int64_numeric,json=fInt64Numeric,proto3" json:"f_int64_numeric,omitempty"`
	FInt64String     []int64   `protobuf:"varint,4,rep,packed,name=f_int64_string,json=fInt64String,proto3" json:"f_int64_string,omitempty"`
	FUint32Numeric   []uint32  `protobuf:"varint,5,rep,packed,name=f_uint32_numeric,json=fUint32Numeric,proto3" json:"f_uint32_numeric,omitempty"`
	FUint32String    []uint32  `protobuf:"varint,6,rep,packed,name=f_uint32_string,json=fUint32String,proto3" json:"f_uint32_string,omitempty"`
	FUint64Numeric   []uint64  `protobuf:"varint,7,rep,packed,name=f_uint64_numeric,json=fUint64Numeric,proto3" json:"f_uint64_numeric,omitempty"`
	FUint64String    []uint64  `protobuf:"varint,8,rep,packed,name=f_uint64_string,json=fUint64String,proto3" json:"f_uint64_string,omitempty"`
	FSint32Numeric   []int32   `protobuf:"zigzag32,9,rep,packed,name=f_sint32_numeric,json=fSint32Numeric,proto3" json:"f_sint32_numeric,omitempty"`
	FSint32String    []int32   `protobuf:"zigzag32,10,rep,packed,name=f_sint32_string,json=fSint32String,proto3" json:"f_sint32_string,omitempty"`
	FSint64Numeric   []int64   `protobuf:"zigzag64,11,rep,packed,name=f_sint64_numeric,json=fSint64Numeric,proto3" json:"f_sint64_numeric,omitempty"`
	FSint64String    []int64   `protobuf:"zigzag64,12,rep,packed,name=f_sint64_string,json=fSint64String,proto3" json:"f_sint64_string,omitempty"`
	FSfixed32Numeric []int32   `protobuf:"fixed32,13,rep,packed,name=f_sfixed32_numeric,json=fSfixed32Numeric,proto3" json:"f_sfixed32_numeric,omitempty"`
	FSfixed32String  []int32   `protobuf:"fixed32,14,rep,packed,name=f_sfixed32_string,json=fSfixed32String,proto3" json:"f_sfixed32_string,omitempty"`
	FSfixed64Numeric []int64   `protobuf:"fixed64,15,rep,packed,name=f_sfixed64_numeric,json=fSfixed64Numeric,proto3" json:"f_sfixed64_numeric,omitempty"`
	FSfixed64String  []int64   `protobuf:"fixed64,16,rep,packed,name=f_sfixed64_string,json=fSfixed64String,proto3" json:"f_sfixed64_string,omitempty"`
	FFixed32Numeric  []uint32  `protobuf:"fixed32,17,rep,packed,name=f_fixed32_numeric,json=fFixed32Numeric,proto3" json:"f_fixed32_numeric,omitempty"`
	FFixed32String   []uint32  `protobuf:"fixed32,18,rep,packed,name=f_fixed32_string,json=fFixed32String,proto3" json:"f_fixed32_string,omitempty"`
	FFixed64Numeric  []uint64  `protobuf:"fixed64,19,rep,packed,name=f_fixed64_numeric,json=fFixed64Numeric,proto3" json:"f_fixed64_numeric,omitempty"`
	FFixed64String   []uint64  `protobuf:"fixed64,20,rep,packed,name=f_fixed64_string,json=fFixed64String,proto3" json:"f_fixed64_string,omitempty"`
	FFloatNumeric    []float32 `protobuf:"fixed32,21,rep,packed,name=f_float_numeric,json=fFloatNumeric,proto3" json:"f_float_numeric,omitempty"`
	FFloatString     []float32 `protobuf:"fixed32,22,rep,packed,name=f_float_string,json=fFloatString,proto3" json:"f_float_string,omitempty"`
	FDoubleNumeric   []float64 `protobuf:"fixed64,23,rep,packed,name=f_double_numeric,json=fDoubleNumeric,proto3" json:"f_double_numeric,omitempty"`
	FDoubleString    []float64 `protobuf:"fixed64,24,rep,packed,name=f_double_string,json=fDoubleString,proto3" json:"f_double_string,omitempty"`
	FBoolBool        []bool    `protobuf:"varint,25,rep,packed,name=f_bool_bool,json=fBoolBool,proto3" json:"f_bool_bool,omitempty"`
	FBoolString      []bool    `protobuf:"varint,26,rep,packed,name=f_bool_string,json=fBoolString,proto3" json:"f_bool_string,omitempty"`
	FStringNone      []string  `protobuf:"bytes,31,rep,name=f_string_none,json=fStringNone,proto3" json:"f_string_none,omitempty"`
	FBytesNone       [][]byte  `protobuf:"bytes,32,rep,name=f_bytes_none,json=fBytesNone,proto3" json:"f_bytes_none,omitempty"`
}

func (x *InvalidCodecRepeated) Reset() {
	*x = InvalidCodecRepeated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_cases_errors_invalid_codec_repeated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidCodecRepeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidCodecRepeated) ProtoMessage() {}

func (x *InvalidCodecRepeated) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_cases_errors_invalid_codec_repeated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidCodecRepeated.ProtoReflect.Descriptor instead.
func (*InvalidCodecRepeated) Descriptor() ([]byte, []int) {
	return file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescGZIP(), []int{0}
}

func (x *InvalidCodecRepeated) GetFInt32Numeric() []int32 {
	if x != nil {
		return x.FInt32Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFInt32String() []int32 {
	if x != nil {
		return x.FInt32String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFInt64Numeric() []int64 {
	if x != nil {
		return x.FInt64Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFInt64String() []int64 {
	if x != nil {
		return x.FInt64String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFUint32Numeric() []uint32 {
	if x != nil {
		return x.FUint32Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFUint32String() []uint32 {
	if x != nil {
		return x.FUint32String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFUint64Numeric() []uint64 {
	if x != nil {
		return x.FUint64Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFUint64String() []uint64 {
	if x != nil {
		return x.FUint64String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSint32Numeric() []int32 {
	if x != nil {
		return x.FSint32Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSint32String() []int32 {
	if x != nil {
		return x.FSint32String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSint64Numeric() []int64 {
	if x != nil {
		return x.FSint64Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSint64String() []int64 {
	if x != nil {
		return x.FSint64String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSfixed32Numeric() []int32 {
	if x != nil {
		return x.FSfixed32Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSfixed32String() []int32 {
	if x != nil {
		return x.FSfixed32String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSfixed64Numeric() []int64 {
	if x != nil {
		return x.FSfixed64Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFSfixed64String() []int64 {
	if x != nil {
		return x.FSfixed64String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFFixed32Numeric() []uint32 {
	if x != nil {
		return x.FFixed32Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFFixed32String() []uint32 {
	if x != nil {
		return x.FFixed32String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFFixed64Numeric() []uint64 {
	if x != nil {
		return x.FFixed64Numeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFFixed64String() []uint64 {
	if x != nil {
		return x.FFixed64String
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFFloatNumeric() []float32 {
	if x != nil {
		return x.FFloatNumeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFFloatString() []float32 {
	if x != nil {
		return x.FFloatString
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFDoubleNumeric() []float64 {
	if x != nil {
		return x.FDoubleNumeric
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFDoubleString() []float64 {
	if x != nil {
		return x.FDoubleString
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFBoolBool() []bool {
	if x != nil {
		return x.FBoolBool
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFBoolString() []bool {
	if x != nil {
		return x.FBoolString
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFStringNone() []string {
	if x != nil {
		return x.FStringNone
	}
	return nil
}

func (x *InvalidCodecRepeated) GetFBytesNone() [][]byte {
	if x != nil {
		return x.FBytesNone
	}
	return nil
}

var File_tests_proto_cases_errors_invalid_codec_repeated_proto protoreflect.FileDescriptor

var file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a,
	0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdd, 0x0c, 0x0a, 0x14, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x63, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x5f,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04,
	0x0a, 0x02, 0x08, 0x01, 0x52, 0x0d, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x12, 0x35, 0x0a, 0x0e, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x42, 0x0f, 0x8a, 0xa2, 0x1f,
	0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x0a, 0x02, 0x08, 0x02, 0x52, 0x0c, 0x66, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x5f,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x03, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04,
	0x12, 0x02, 0x08, 0x01, 0x52, 0x0d, 0x66, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x12, 0x35, 0x0a, 0x0e, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x42, 0x0f, 0x8a, 0xa2, 0x1f,
	0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x12, 0x02, 0x08, 0x02, 0x52, 0x0c, 0x66, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x10, 0x66, 0x5f,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0d, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a,
	0x04, 0x3a, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x4e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x0f,
	0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x3a, 0x02, 0x08, 0x02, 0x52,
	0x0d, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x39,
	0x0a, 0x10, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x18, 0x07, 0x20, 0x03, 0x28, 0x04, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09,
	0xfa, 0x01, 0x06, 0x0a, 0x04, 0x42, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x66, 0x55, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x04, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x42,
	0x02, 0x08, 0x02, 0x52, 0x0d, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x39, 0x0a, 0x10, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x09, 0x20, 0x03, 0x28, 0x11, 0x42, 0x0f, 0x8a, 0xa2,
	0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x1a, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x66,
	0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x37, 0x0a,
	0x0f, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x11, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01,
	0x06, 0x0a, 0x04, 0x1a, 0x02, 0x08, 0x02, 0x52, 0x0d, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x10, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x12,
	0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x22, 0x02, 0x08,
	0x01, 0x52, 0x0e, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x12, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b,
	0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x22, 0x02, 0x08, 0x02, 0x52, 0x0d, 0x66, 0x53, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x12, 0x66, 0x5f,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0f, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01,
	0x06, 0x0a, 0x04, 0x2a, 0x02, 0x08, 0x01, 0x52, 0x10, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x3b, 0x0a, 0x11, 0x66, 0x5f, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0f, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a,
	0x04, 0x2a, 0x02, 0x08, 0x02, 0x52, 0x0f, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x12, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x10, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x32,
	0x02, 0x08, 0x01, 0x52, 0x10, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x4e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x3b, 0x0a, 0x11, 0x66, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x03, 0x28, 0x10,
	0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x32, 0x02, 0x08,
	0x02, 0x52, 0x0f, 0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x11, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x11, 0x20, 0x03, 0x28, 0x07, 0x42, 0x0f, 0x8a,
	0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x4a, 0x02, 0x08, 0x01, 0x52, 0x0f,
	0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12,
	0x39, 0x0a, 0x10, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x12, 0x20, 0x03, 0x28, 0x07, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a,
	0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x4a, 0x02, 0x08, 0x02, 0x52, 0x0e, 0x66, 0x46, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x11, 0x66, 0x5f,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18,
	0x13, 0x20, 0x03, 0x28, 0x06, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06,
	0x0a, 0x04, 0x52, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x39, 0x0a, 0x10, 0x66, 0x5f, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x06, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x52, 0x02,
	0x08, 0x02, 0x52, 0x0e, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x15, 0x20, 0x03, 0x28, 0x02, 0x42, 0x0f, 0x8a, 0xa2, 0x1f,
	0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x5a, 0x02, 0x08, 0x01, 0x52, 0x0d, 0x66, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x35, 0x0a, 0x0e, 0x66,
	0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x16, 0x20,
	0x03, 0x28, 0x02, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04,
	0x5a, 0x02, 0x08, 0x02, 0x52, 0x0c, 0x66, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x39, 0x0a, 0x10, 0x66, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x18, 0x17, 0x20, 0x03, 0x28, 0x01, 0x42, 0x0f, 0x8a, 0xa2,
	0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x62, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x66,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x37, 0x0a,
	0x0f, 0x66, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x18, 0x20, 0x03, 0x28, 0x01, 0x42, 0x0f, 0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01,
	0x06, 0x0a, 0x04, 0x62, 0x02, 0x08, 0x02, 0x52, 0x0d, 0x66, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x0b, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x19, 0x20, 0x03, 0x28, 0x08, 0x42, 0x0f, 0x8a, 0xa2, 0x1f,
	0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x52, 0x09, 0x66, 0x42,
	0x6f, 0x6f, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x33, 0x0a, 0x0d, 0x66, 0x5f, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x08, 0x42, 0x0f,
	0x8a, 0xa2, 0x1f, 0x0b, 0x5a, 0x09, 0xfa, 0x01, 0x06, 0x0a, 0x04, 0x6a, 0x02, 0x08, 0x02, 0x52,
	0x0b, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x0d,
	0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x18, 0x1f, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x0d, 0x8a, 0xa2, 0x1f, 0x09, 0x5a, 0x07, 0xfa, 0x01, 0x04, 0x0a, 0x02,
	0x7a, 0x00, 0x52, 0x0b, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x6e, 0x65, 0x12,
	0x30, 0x0a, 0x0c, 0x66, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x18,
	0x20, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x0e, 0x8a, 0xa2, 0x1f, 0x0a, 0x5a, 0x08, 0xfa, 0x01, 0x05,
	0x0a, 0x03, 0x82, 0x01, 0x00, 0x52, 0x0a, 0x66, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4e, 0x6f, 0x6e,
	0x65, 0x42, 0x17, 0x5a, 0x15, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x62, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescOnce sync.Once
	file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescData = file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDesc
)

func file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescGZIP() []byte {
	file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescOnce.Do(func() {
		file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescData)
	})
	return file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDescData
}

var file_tests_proto_cases_errors_invalid_codec_repeated_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_proto_cases_errors_invalid_codec_repeated_proto_goTypes = []interface{}{
	(*InvalidCodecRepeated)(nil), // 0: errors.InvalidCodecRepeated
}
var file_tests_proto_cases_errors_invalid_codec_repeated_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_proto_cases_errors_invalid_codec_repeated_proto_init() }
func file_tests_proto_cases_errors_invalid_codec_repeated_proto_init() {
	if File_tests_proto_cases_errors_invalid_codec_repeated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_cases_errors_invalid_codec_repeated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidCodecRepeated); i {
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
			RawDescriptor: file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_cases_errors_invalid_codec_repeated_proto_goTypes,
		DependencyIndexes: file_tests_proto_cases_errors_invalid_codec_repeated_proto_depIdxs,
		MessageInfos:      file_tests_proto_cases_errors_invalid_codec_repeated_proto_msgTypes,
	}.Build()
	File_tests_proto_cases_errors_invalid_codec_repeated_proto = out.File
	file_tests_proto_cases_errors_invalid_codec_repeated_proto_rawDesc = nil
	file_tests_proto_cases_errors_invalid_codec_repeated_proto_goTypes = nil
	file_tests_proto_cases_errors_invalid_codec_repeated_proto_depIdxs = nil
}

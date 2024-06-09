// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/module/external.proto

package pbexternal

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

type EnumNum1 int32

const (
	EnumNum1_Zero  EnumNum1 = 0
	EnumNum1_One   EnumNum1 = 2
	EnumNum1_Two   EnumNum1 = 3
	EnumNum1_Three EnumNum1 = 5
	EnumNum1_Four  EnumNum1 = 6
	EnumNum1_Five  EnumNum1 = 7
	EnumNum1_Six   EnumNum1 = 9
	EnumNum1_Seven EnumNum1 = 11
	EnumNum1_Eight EnumNum1 = 12
	EnumNum1_Nine  EnumNum1 = 15
	EnumNum1_Ten   EnumNum1 = 17
)

// Enum value maps for EnumNum1.
var (
	EnumNum1_name = map[int32]string{
		0:  "Zero",
		2:  "One",
		3:  "Two",
		5:  "Three",
		6:  "Four",
		7:  "Five",
		9:  "Six",
		11: "Seven",
		12: "Eight",
		15: "Nine",
		17: "Ten",
	}
	EnumNum1_value = map[string]int32{
		"Zero":  0,
		"One":   2,
		"Two":   3,
		"Three": 5,
		"Four":  6,
		"Five":  7,
		"Six":   9,
		"Seven": 11,
		"Eight": 12,
		"Nine":  15,
		"Ten":   17,
	}
)

func (x EnumNum1) Enum() *EnumNum1 {
	p := new(EnumNum1)
	*p = x
	return p
}

func (x EnumNum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumNum1) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_proto_module_external_proto_enumTypes[0].Descriptor()
}

func (EnumNum1) Type() protoreflect.EnumType {
	return &file_tests_proto_module_external_proto_enumTypes[0]
}

func (x EnumNum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumNum1.Descriptor instead.
func (EnumNum1) EnumDescriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{0}
}

type Embed_EnumNum1 int32

const (
	Embed_Zero  Embed_EnumNum1 = 0
	Embed_One   Embed_EnumNum1 = 2
	Embed_Two   Embed_EnumNum1 = 3
	Embed_Three Embed_EnumNum1 = 5
	Embed_Four  Embed_EnumNum1 = 6
	Embed_Five  Embed_EnumNum1 = 7
	Embed_Six   Embed_EnumNum1 = 9
	Embed_Seven Embed_EnumNum1 = 11
	Embed_Eight Embed_EnumNum1 = 12
	Embed_Nine  Embed_EnumNum1 = 15
	Embed_Ten   Embed_EnumNum1 = 17
)

// Enum value maps for Embed_EnumNum1.
var (
	Embed_EnumNum1_name = map[int32]string{
		0:  "Zero",
		2:  "One",
		3:  "Two",
		5:  "Three",
		6:  "Four",
		7:  "Five",
		9:  "Six",
		11: "Seven",
		12: "Eight",
		15: "Nine",
		17: "Ten",
	}
	Embed_EnumNum1_value = map[string]int32{
		"Zero":  0,
		"One":   2,
		"Two":   3,
		"Three": 5,
		"Four":  6,
		"Five":  7,
		"Six":   9,
		"Seven": 11,
		"Eight": 12,
		"Nine":  15,
		"Ten":   17,
	}
)

func (x Embed_EnumNum1) Enum() *Embed_EnumNum1 {
	p := new(Embed_EnumNum1)
	*p = x
	return p
}

func (x Embed_EnumNum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Embed_EnumNum1) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_proto_module_external_proto_enumTypes[1].Descriptor()
}

func (Embed_EnumNum1) Type() protoreflect.EnumType {
	return &file_tests_proto_module_external_proto_enumTypes[1]
}

func (x Embed_EnumNum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Embed_EnumNum1.Descriptor instead.
func (Embed_EnumNum1) EnumDescriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{0, 0}
}

type Embed_Message_EnumNum1 int32

const (
	Embed_Message_Zero  Embed_Message_EnumNum1 = 0
	Embed_Message_One   Embed_Message_EnumNum1 = 2
	Embed_Message_Two   Embed_Message_EnumNum1 = 3
	Embed_Message_Three Embed_Message_EnumNum1 = 5
	Embed_Message_Four  Embed_Message_EnumNum1 = 6
	Embed_Message_Five  Embed_Message_EnumNum1 = 7
	Embed_Message_Six   Embed_Message_EnumNum1 = 9
	Embed_Message_Seven Embed_Message_EnumNum1 = 11
	Embed_Message_Eight Embed_Message_EnumNum1 = 12
	Embed_Message_Nine  Embed_Message_EnumNum1 = 15
	Embed_Message_Ten   Embed_Message_EnumNum1 = 17
)

// Enum value maps for Embed_Message_EnumNum1.
var (
	Embed_Message_EnumNum1_name = map[int32]string{
		0:  "Zero",
		2:  "One",
		3:  "Two",
		5:  "Three",
		6:  "Four",
		7:  "Five",
		9:  "Six",
		11: "Seven",
		12: "Eight",
		15: "Nine",
		17: "Ten",
	}
	Embed_Message_EnumNum1_value = map[string]int32{
		"Zero":  0,
		"One":   2,
		"Two":   3,
		"Three": 5,
		"Four":  6,
		"Five":  7,
		"Six":   9,
		"Seven": 11,
		"Eight": 12,
		"Nine":  15,
		"Ten":   17,
	}
)

func (x Embed_Message_EnumNum1) Enum() *Embed_Message_EnumNum1 {
	p := new(Embed_Message_EnumNum1)
	*p = x
	return p
}

func (x Embed_Message_EnumNum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Embed_Message_EnumNum1) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_proto_module_external_proto_enumTypes[2].Descriptor()
}

func (Embed_Message_EnumNum1) Type() protoreflect.EnumType {
	return &file_tests_proto_module_external_proto_enumTypes[2]
}

func (x Embed_Message_EnumNum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Embed_Message_EnumNum1.Descriptor instead.
func (Embed_Message_EnumNum1) EnumDescriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Embed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Embed) Reset() {
	*x = Embed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_module_external_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embed) ProtoMessage() {}

func (x *Embed) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_module_external_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embed.ProtoReflect.Descriptor instead.
func (*Embed) Descriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{0}
}

type Message1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
	FString3 string `protobuf:"bytes,3,opt,name=f_string3,json=fString3,proto3" json:"f_string3,omitempty"`
}

func (x *Message1) Reset() {
	*x = Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_module_external_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1) ProtoMessage() {}

func (x *Message1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_module_external_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1.ProtoReflect.Descriptor instead.
func (*Message1) Descriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{1}
}

func (x *Message1) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *Message1) GetFString2() string {
	if x != nil {
		return x.FString2
	}
	return ""
}

func (x *Message1) GetFString3() string {
	if x != nil {
		return x.FString3
	}
	return ""
}

type Embed_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Embed_Message) Reset() {
	*x = Embed_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_module_external_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embed_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embed_Message) ProtoMessage() {}

func (x *Embed_Message) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_module_external_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embed_Message.ProtoReflect.Descriptor instead.
func (*Embed_Message) Descriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{0, 0}
}

type Message1_Embed1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
	FString3 string `protobuf:"bytes,3,opt,name=f_string3,json=fString3,proto3" json:"f_string3,omitempty"`
}

func (x *Message1_Embed1) Reset() {
	*x = Message1_Embed1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_module_external_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1_Embed1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1_Embed1) ProtoMessage() {}

func (x *Message1_Embed1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_module_external_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1_Embed1.ProtoReflect.Descriptor instead.
func (*Message1_Embed1) Descriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Message1_Embed1) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *Message1_Embed1) GetFString2() string {
	if x != nil {
		return x.FString2
	}
	return ""
}

func (x *Message1_Embed1) GetFString3() string {
	if x != nil {
		return x.FString3
	}
	return ""
}

type Message1_Embed1_Embed2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
	FString3 string `protobuf:"bytes,3,opt,name=f_string3,json=fString3,proto3" json:"f_string3,omitempty"`
}

func (x *Message1_Embed1_Embed2) Reset() {
	*x = Message1_Embed1_Embed2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_module_external_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1_Embed1_Embed2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1_Embed1_Embed2) ProtoMessage() {}

func (x *Message1_Embed1_Embed2) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_module_external_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1_Embed1_Embed2.ProtoReflect.Descriptor instead.
func (*Message1_Embed1_Embed2) Descriptor() ([]byte, []int) {
	return file_tests_proto_module_external_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *Message1_Embed1_Embed2) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (x *Message1_Embed1_Embed2) GetFString2() string {
	if x != nil {
		return x.FString2
	}
	return ""
}

func (x *Message1_Embed1_Embed2) GetFString3() string {
	if x != nil {
		return x.FString3
	}
	return ""
}

var File_tests_proto_module_external_proto protoreflect.FileDescriptor

var file_tests_proto_module_external_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x85, 0x02, 0x0a, 0x05, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x1a, 0x82, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x77, 0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x75, 0x6d,
	0x31, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x65, 0x72, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f,
	0x6e, 0x65, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x54, 0x68, 0x72, 0x65, 0x65, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x75, 0x72,
	0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x76, 0x65, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x69, 0x78, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x10, 0x0b,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x69, 0x67, 0x68, 0x74, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x69, 0x6e, 0x65, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x65, 0x6e, 0x10, 0x11, 0x22, 0x77,
	0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x31, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x65,
	0x72, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x54, 0x77, 0x6f, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x68, 0x72, 0x65, 0x65, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x75, 0x72, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x69, 0x76, 0x65, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x69, 0x78, 0x10, 0x09, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x69, 0x67,
	0x68, 0x74, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x69, 0x6e, 0x65, 0x10, 0x0f, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x65, 0x6e, 0x10, 0x11, 0x22, 0xda, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00,
	0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a,
	0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x1a, 0xe4, 0x01,
	0x0a, 0x06, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x31, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f,
	0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x21, 0x0a, 0x09, 0x66,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x21,
	0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x33, 0x1a, 0x71, 0x0a, 0x06, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x32, 0x12, 0x21, 0x0a, 0x09, 0x66,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x21,
	0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x32, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x33, 0x2a, 0x77, 0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x31,
	0x12, 0x08, 0x0a, 0x04, 0x5a, 0x65, 0x72, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e,
	0x65, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x54, 0x68, 0x72, 0x65, 0x65, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x75, 0x72, 0x10,
	0x06, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x76, 0x65, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x53,
	0x69, 0x78, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x10, 0x0b, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x69, 0x67, 0x68, 0x74, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x69,
	0x6e, 0x65, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x65, 0x6e, 0x10, 0x11, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x6a,
	0x73, 0x6f, 0x6e, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x62, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_module_external_proto_rawDescOnce sync.Once
	file_tests_proto_module_external_proto_rawDescData = file_tests_proto_module_external_proto_rawDesc
)

func file_tests_proto_module_external_proto_rawDescGZIP() []byte {
	file_tests_proto_module_external_proto_rawDescOnce.Do(func() {
		file_tests_proto_module_external_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_module_external_proto_rawDescData)
	})
	return file_tests_proto_module_external_proto_rawDescData
}

var file_tests_proto_module_external_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_tests_proto_module_external_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tests_proto_module_external_proto_goTypes = []interface{}{
	(EnumNum1)(0),                  // 0: external.EnumNum1
	(Embed_EnumNum1)(0),            // 1: external.Embed.EnumNum1
	(Embed_Message_EnumNum1)(0),    // 2: external.Embed.Message.EnumNum1
	(*Embed)(nil),                  // 3: external.Embed
	(*Message1)(nil),               // 4: external.Message1
	(*Embed_Message)(nil),          // 5: external.Embed.Message
	(*Message1_Embed1)(nil),        // 6: external.Message1.Embed1
	(*Message1_Embed1_Embed2)(nil), // 7: external.Message1.Embed1.Embed2
}
var file_tests_proto_module_external_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_proto_module_external_proto_init() }
func file_tests_proto_module_external_proto_init() {
	if File_tests_proto_module_external_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_module_external_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embed); i {
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
		file_tests_proto_module_external_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message1); i {
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
		file_tests_proto_module_external_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embed_Message); i {
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
		file_tests_proto_module_external_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message1_Embed1); i {
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
		file_tests_proto_module_external_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message1_Embed1_Embed2); i {
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
			RawDescriptor: file_tests_proto_module_external_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_module_external_proto_goTypes,
		DependencyIndexes: file_tests_proto_module_external_proto_depIdxs,
		EnumInfos:         file_tests_proto_module_external_proto_enumTypes,
		MessageInfos:      file_tests_proto_module_external_proto_msgTypes,
	}.Build()
	File_tests_proto_module_external_proto = out.File
	file_tests_proto_module_external_proto_rawDesc = nil
	file_tests_proto_module_external_proto_goTypes = nil
	file_tests_proto_module_external_proto_depIdxs = nil
}

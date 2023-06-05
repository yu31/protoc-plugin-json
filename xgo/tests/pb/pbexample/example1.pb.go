// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: tests/proto/example/example1.proto

package pbexample

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
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

// EnumPlain1 used tests cases in this file.
type Enum1 int32

const (
	Enum1_Zero  Enum1 = 0
	Enum1_One   Enum1 = 2
	Enum1_Two   Enum1 = 3
	Enum1_Three Enum1 = 5
	Enum1_Four  Enum1 = 6
	Enum1_Five  Enum1 = 7
	Enum1_Six   Enum1 = 9
	Enum1_Seven Enum1 = 11
	Enum1_Eight Enum1 = 12
	Enum1_Nine  Enum1 = 15
	Enum1_Ten   Enum1 = 17
)

// Enum value maps for Enum1.
var (
	Enum1_name = map[int32]string{
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
	Enum1_value = map[string]int32{
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

func (x Enum1) Enum() *Enum1 {
	p := new(Enum1)
	*p = x
	return p
}

func (x Enum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum1) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_proto_example_example1_proto_enumTypes[0].Descriptor()
}

func (Enum1) Type() protoreflect.EnumType {
	return &file_tests_proto_example_example1_proto_enumTypes[0]
}

func (x Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum1.Descriptor instead.
func (Enum1) EnumDescriptor() ([]byte, []int) {
	return file_tests_proto_example_example1_proto_rawDescGZIP(), []int{0}
}

type Empty1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty1) Reset() {
	*x = Empty1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_example_example1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty1) ProtoMessage() {}

func (x *Empty1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_example_example1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty1.ProtoReflect.Descriptor instead.
func (*Empty1) Descriptor() ([]byte, []int) {
	return file_tests_proto_example_example1_proto_rawDescGZIP(), []int{0}
}

type Example1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	// Types that are assignable to OnetType1:
	//	*Example1_FString2
	//	*Example1_FString3
	OnetType1 isExample1_OnetType1 `protobuf_oneof:"OnetType1"`
	// Types that are assignable to OnetType2:
	//	*Example1_FString4
	//	*Example1_FString5
	OnetType2   isExample1_OnetType2    `protobuf_oneof:"OnetType2"`
	FInt32      int32                   `protobuf:"varint,11,opt,name=f_int32,json=fInt32,proto3" json:"f_int32,omitempty"`
	FInt64      int64                   `protobuf:"varint,12,opt,name=f_int64,json=fInt64,proto3" json:"f_int64,omitempty"`
	FUint32     uint32                  `protobuf:"varint,13,opt,name=f_uint32,json=fUint32,proto3" json:"f_uint32,omitempty"`
	FUint64     uint64                  `protobuf:"varint,14,opt,name=f_uint64,json=fUint64,proto3" json:"f_uint64,omitempty"`
	FSint32     int32                   `protobuf:"zigzag32,15,opt,name=f_sint32,json=fSint32,proto3" json:"f_sint32,omitempty"`
	FSint64     int64                   `protobuf:"zigzag64,16,opt,name=f_sint64,json=fSint64,proto3" json:"f_sint64,omitempty"`
	FSfixed32   int32                   `protobuf:"fixed32,17,opt,name=f_sfixed32,json=fSfixed32,proto3" json:"f_sfixed32,omitempty"`
	FSfixed64   int64                   `protobuf:"fixed64,18,opt,name=f_sfixed64,json=fSfixed64,proto3" json:"f_sfixed64,omitempty"`
	FFixed32    uint32                  `protobuf:"fixed32,19,opt,name=f_fixed32,json=fFixed32,proto3" json:"f_fixed32,omitempty"`
	FFixed64    uint64                  `protobuf:"fixed64,20,opt,name=f_fixed64,json=fFixed64,proto3" json:"f_fixed64,omitempty"`
	FFloat      float32                 `protobuf:"fixed32,21,opt,name=f_float,json=fFloat,proto3" json:"f_float,omitempty"`
	FDouble     float64                 `protobuf:"fixed64,22,opt,name=f_double,json=fDouble,proto3" json:"f_double,omitempty"`
	FBool1      bool                    `protobuf:"varint,31,opt,name=f_bool1,json=fBool1,proto3" json:"f_bool1,omitempty"`
	FBytes1     []byte                  `protobuf:"bytes,32,opt,name=f_bytes1,json=fBytes1,proto3" json:"f_bytes1,omitempty"`
	FEnum1      Enum1                   `protobuf:"varint,61,opt,name=f_enum1,json=fEnum1,proto3,enum=example1.Enum1" json:"f_enum1,omitempty"`
	FEnum2      Enum1                   `protobuf:"varint,62,opt,name=f_enum2,json=fEnum2,proto3,enum=example1.Enum1" json:"f_enum2,omitempty"`
	FEnum5      *Enum1                  `protobuf:"varint,65,opt,name=f_enum5,json=fEnum5,proto3,enum=example1.Enum1,oneof" json:"f_enum5,omitempty"`
	FEnum6      *Enum1                  `protobuf:"varint,66,opt,name=f_enum6,json=fEnum6,proto3,enum=example1.Enum1,oneof" json:"f_enum6,omitempty"`
	FDuration1  *durationpb.Duration    `protobuf:"bytes,41,opt,name=f_duration1,json=fDuration1,proto3" json:"f_duration1,omitempty"`
	FTimestamp1 *timestamppb.Timestamp  `protobuf:"bytes,42,opt,name=f_timestamp1,json=fTimestamp1,proto3" json:"f_timestamp1,omitempty"`
	FMessage11  *Message1               `protobuf:"bytes,51,opt,name=f_message11,json=fMessage11,proto3" json:"f_message11,omitempty"`
	FMessage12  *Message1_Embed1        `protobuf:"bytes,52,opt,name=f_message12,json=fMessage12,proto3" json:"f_message12,omitempty"`
	FMessage13  *Message1_Embed1_Embed2 `protobuf:"bytes,53,opt,name=f_message13,json=fMessage13,proto3" json:"f_message13,omitempty"`
	FAny1       *anypb.Any              `protobuf:"bytes,54,opt,name=f_any1,json=fAny1,proto3" json:"f_any1,omitempty"`
	RString1    []string                `protobuf:"bytes,71,rep,name=r_string1,json=rString1,proto3" json:"r_string1,omitempty"`
	RInt32      []int32                 `protobuf:"varint,72,rep,packed,name=r_int32,json=rInt32,proto3" json:"r_int32,omitempty"`
	RMessage    []*Message1             `protobuf:"bytes,73,rep,name=r_message,json=rMessage,proto3" json:"r_message,omitempty"`
	REnum       []Enum1                 `protobuf:"varint,74,rep,packed,name=r_enum,json=rEnum,proto3,enum=example1.Enum1" json:"r_enum,omitempty"`
	MString1    map[string]string       `protobuf:"bytes,81,rep,name=m_string1,json=mString1,proto3" json:"m_string1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MString2    map[int32]string        `protobuf:"bytes,84,rep,name=m_string2,json=mString2,proto3" json:"m_string2,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //  map<bool, string> m_string3 = 86; // TODO: unsupported now.
	MMessage1   map[string]*Message1    `protobuf:"bytes,82,rep,name=m_message1,json=mMessage1,proto3" json:"m_message1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MEnum       map[string]Enum1        `protobuf:"bytes,83,rep,name=m_enum,json=mEnum,proto3" json:"m_enum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=example1.Enum1"`
}

func (x *Example1) Reset() {
	*x = Example1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_example_example1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example1) ProtoMessage() {}

func (x *Example1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_example_example1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example1.ProtoReflect.Descriptor instead.
func (*Example1) Descriptor() ([]byte, []int) {
	return file_tests_proto_example_example1_proto_rawDescGZIP(), []int{1}
}

func (x *Example1) GetFString1() string {
	if x != nil {
		return x.FString1
	}
	return ""
}

func (m *Example1) GetOnetType1() isExample1_OnetType1 {
	if m != nil {
		return m.OnetType1
	}
	return nil
}

func (x *Example1) GetFString2() string {
	if x, ok := x.GetOnetType1().(*Example1_FString2); ok {
		return x.FString2
	}
	return ""
}

func (x *Example1) GetFString3() string {
	if x, ok := x.GetOnetType1().(*Example1_FString3); ok {
		return x.FString3
	}
	return ""
}

func (m *Example1) GetOnetType2() isExample1_OnetType2 {
	if m != nil {
		return m.OnetType2
	}
	return nil
}

func (x *Example1) GetFString4() string {
	if x, ok := x.GetOnetType2().(*Example1_FString4); ok {
		return x.FString4
	}
	return ""
}

func (x *Example1) GetFString5() string {
	if x, ok := x.GetOnetType2().(*Example1_FString5); ok {
		return x.FString5
	}
	return ""
}

func (x *Example1) GetFInt32() int32 {
	if x != nil {
		return x.FInt32
	}
	return 0
}

func (x *Example1) GetFInt64() int64 {
	if x != nil {
		return x.FInt64
	}
	return 0
}

func (x *Example1) GetFUint32() uint32 {
	if x != nil {
		return x.FUint32
	}
	return 0
}

func (x *Example1) GetFUint64() uint64 {
	if x != nil {
		return x.FUint64
	}
	return 0
}

func (x *Example1) GetFSint32() int32 {
	if x != nil {
		return x.FSint32
	}
	return 0
}

func (x *Example1) GetFSint64() int64 {
	if x != nil {
		return x.FSint64
	}
	return 0
}

func (x *Example1) GetFSfixed32() int32 {
	if x != nil {
		return x.FSfixed32
	}
	return 0
}

func (x *Example1) GetFSfixed64() int64 {
	if x != nil {
		return x.FSfixed64
	}
	return 0
}

func (x *Example1) GetFFixed32() uint32 {
	if x != nil {
		return x.FFixed32
	}
	return 0
}

func (x *Example1) GetFFixed64() uint64 {
	if x != nil {
		return x.FFixed64
	}
	return 0
}

func (x *Example1) GetFFloat() float32 {
	if x != nil {
		return x.FFloat
	}
	return 0
}

func (x *Example1) GetFDouble() float64 {
	if x != nil {
		return x.FDouble
	}
	return 0
}

func (x *Example1) GetFBool1() bool {
	if x != nil {
		return x.FBool1
	}
	return false
}

func (x *Example1) GetFBytes1() []byte {
	if x != nil {
		return x.FBytes1
	}
	return nil
}

func (x *Example1) GetFEnum1() Enum1 {
	if x != nil {
		return x.FEnum1
	}
	return Enum1_Zero
}

func (x *Example1) GetFEnum2() Enum1 {
	if x != nil {
		return x.FEnum2
	}
	return Enum1_Zero
}

func (x *Example1) GetFEnum5() Enum1 {
	if x != nil && x.FEnum5 != nil {
		return *x.FEnum5
	}
	return Enum1_Zero
}

func (x *Example1) GetFEnum6() Enum1 {
	if x != nil && x.FEnum6 != nil {
		return *x.FEnum6
	}
	return Enum1_Zero
}

func (x *Example1) GetFDuration1() *durationpb.Duration {
	if x != nil {
		return x.FDuration1
	}
	return nil
}

func (x *Example1) GetFTimestamp1() *timestamppb.Timestamp {
	if x != nil {
		return x.FTimestamp1
	}
	return nil
}

func (x *Example1) GetFMessage11() *Message1 {
	if x != nil {
		return x.FMessage11
	}
	return nil
}

func (x *Example1) GetFMessage12() *Message1_Embed1 {
	if x != nil {
		return x.FMessage12
	}
	return nil
}

func (x *Example1) GetFMessage13() *Message1_Embed1_Embed2 {
	if x != nil {
		return x.FMessage13
	}
	return nil
}

func (x *Example1) GetFAny1() *anypb.Any {
	if x != nil {
		return x.FAny1
	}
	return nil
}

func (x *Example1) GetRString1() []string {
	if x != nil {
		return x.RString1
	}
	return nil
}

func (x *Example1) GetRInt32() []int32 {
	if x != nil {
		return x.RInt32
	}
	return nil
}

func (x *Example1) GetRMessage() []*Message1 {
	if x != nil {
		return x.RMessage
	}
	return nil
}

func (x *Example1) GetREnum() []Enum1 {
	if x != nil {
		return x.REnum
	}
	return nil
}

func (x *Example1) GetMString1() map[string]string {
	if x != nil {
		return x.MString1
	}
	return nil
}

func (x *Example1) GetMString2() map[int32]string {
	if x != nil {
		return x.MString2
	}
	return nil
}

func (x *Example1) GetMMessage1() map[string]*Message1 {
	if x != nil {
		return x.MMessage1
	}
	return nil
}

func (x *Example1) GetMEnum() map[string]Enum1 {
	if x != nil {
		return x.MEnum
	}
	return nil
}

type isExample1_OnetType1 interface {
	isExample1_OnetType1()
}

type Example1_FString2 struct {
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3,oneof"`
}

type Example1_FString3 struct {
	FString3 string `protobuf:"bytes,3,opt,name=f_string3,json=fString3,proto3,oneof"`
}

func (*Example1_FString2) isExample1_OnetType1() {}

func (*Example1_FString3) isExample1_OnetType1() {}

type isExample1_OnetType2 interface {
	isExample1_OnetType2()
}

type Example1_FString4 struct {
	FString4 string `protobuf:"bytes,4,opt,name=f_string4,json=fString4,proto3,oneof"`
}

type Example1_FString5 struct {
	FString5 string `protobuf:"bytes,5,opt,name=f_string5,json=fString5,proto3,oneof"`
}

func (*Example1_FString4) isExample1_OnetType2() {}

func (*Example1_FString5) isExample1_OnetType2() {}

type Message1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
}

func (x *Message1) Reset() {
	*x = Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_example_example1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1) ProtoMessage() {}

func (x *Message1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_example_example1_proto_msgTypes[2]
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
	return file_tests_proto_example_example1_proto_rawDescGZIP(), []int{2}
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

type Message1_Embed1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
}

func (x *Message1_Embed1) Reset() {
	*x = Message1_Embed1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_example_example1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1_Embed1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1_Embed1) ProtoMessage() {}

func (x *Message1_Embed1) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_example_example1_proto_msgTypes[7]
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
	return file_tests_proto_example_example1_proto_rawDescGZIP(), []int{2, 0}
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

type Message1_Embed1_Embed2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FString1 string `protobuf:"bytes,1,opt,name=f_string1,json=fString1,proto3" json:"f_string1,omitempty"`
	FString2 string `protobuf:"bytes,2,opt,name=f_string2,json=fString2,proto3" json:"f_string2,omitempty"`
}

func (x *Message1_Embed1_Embed2) Reset() {
	*x = Message1_Embed1_Embed2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_example_example1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message1_Embed1_Embed2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1_Embed1_Embed2) ProtoMessage() {}

func (x *Message1_Embed1_Embed2) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_example_example1_proto_msgTypes[8]
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
	return file_tests_proto_example_example1_proto_rawDescGZIP(), []int{2, 0, 0}
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

var File_tests_proto_example_example1_proto protoreflect.FileDescriptor

var file_tests_proto_example_example1_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x31, 0x22, 0xa5, 0x0e, 0x0a, 0x08, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x31, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31,
	0x12, 0x1d, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x12,
	0x1d, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x33, 0x12, 0x1d,
	0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x34, 0x12, 0x1d, 0x0a,
	0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x35, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x66, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x66, 0x55, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x07, 0x66, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x5f,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x09,
	0x66, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x5f, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x12, 0x20, 0x01, 0x28, 0x10, 0x52, 0x09, 0x66,
	0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x13, 0x20, 0x01, 0x28, 0x07, 0x52, 0x08, 0x66, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x18, 0x14, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x66, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x66, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c,
	0x31, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x31, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x31, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x66, 0x42, 0x79, 0x74, 0x65, 0x73, 0x31, 0x12, 0x32, 0x0a, 0x07, 0x66, 0x5f,
	0x65, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x42, 0x08, 0x8a, 0xa2,
	0x1f, 0x04, 0x18, 0x01, 0x20, 0x01, 0x52, 0x06, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x2e,
	0x0a, 0x07, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31,
	0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x52, 0x06, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x32, 0x12, 0x37,
	0x0a, 0x07, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x35, 0x18, 0x41, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31,
	0x42, 0x08, 0x8a, 0xa2, 0x1f, 0x04, 0x18, 0x01, 0x20, 0x01, 0x48, 0x02, 0x52, 0x06, 0x66, 0x45,
	0x6e, 0x75, 0x6d, 0x35, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x07, 0x66, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x36, 0x18, 0x42, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x42, 0x04, 0x8a, 0xa2, 0x1f, 0x00, 0x48,
	0x03, 0x52, 0x06, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x36, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x0b,
	0x66, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x18, 0x29, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x12, 0x3d, 0x0a, 0x0c, 0x66, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x31, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x66, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x31, 0x12, 0x33, 0x0a, 0x0b, 0x66, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x31, 0x31, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31,
	0x52, 0x0a, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x31, 0x12, 0x3a, 0x0a, 0x0b,
	0x66, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x32, 0x18, 0x34, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x31, 0x52, 0x0a, 0x66, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x32, 0x12, 0x41, 0x0a, 0x0b, 0x66, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x33, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x32, 0x52,
	0x0a, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x33, 0x12, 0x2b, 0x0a, 0x06, 0x66,
	0x5f, 0x61, 0x6e, 0x79, 0x31, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x05, 0x66, 0x41, 0x6e, 0x79, 0x31, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x47, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x48, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x72, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x2f,
	0x0a, 0x09, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x49, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x31, 0x52, 0x08, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x4a, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31,
	0x52, 0x05, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x3d, 0x0a, 0x09, 0x6d, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x31, 0x18, 0x51, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x3d, 0x0a, 0x09, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x32, 0x18, 0x54, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x32, 0x12, 0x40, 0x0a, 0x0a, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x31, 0x18, 0x52, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x18, 0x53, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x45, 0x6e, 0x75,
	0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x0e, 0x4d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x49, 0x0a, 0x0a, 0x4d, 0x45, 0x6e,
	0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x4f, 0x6e, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x31, 0x42, 0x13, 0x0a, 0x09, 0x4f, 0x6e, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x32, 0x12, 0x06,
	0x92, 0xa2, 0x1f, 0x02, 0x20, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x35, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x36, 0x22, 0xcd,
	0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x32, 0x1a, 0x86, 0x01, 0x0a, 0x06, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x31,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x31, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x1a, 0x42, 0x0a, 0x06, 0x45, 0x6d,
	0x62, 0x65, 0x64, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x31, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x2a, 0x74,
	0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x65, 0x72, 0x6f, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77,
	0x6f, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x68, 0x72, 0x65, 0x65, 0x10, 0x05, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x6f, 0x75, 0x72, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x76, 0x65,
	0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x69, 0x78, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x65, 0x76, 0x65, 0x6e, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x69, 0x67, 0x68, 0x74, 0x10,
	0x0c, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x69, 0x6e, 0x65, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x54,
	0x65, 0x6e, 0x10, 0x11, 0x42, 0x18, 0x5a, 0x16, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_example_example1_proto_rawDescOnce sync.Once
	file_tests_proto_example_example1_proto_rawDescData = file_tests_proto_example_example1_proto_rawDesc
)

func file_tests_proto_example_example1_proto_rawDescGZIP() []byte {
	file_tests_proto_example_example1_proto_rawDescOnce.Do(func() {
		file_tests_proto_example_example1_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_example_example1_proto_rawDescData)
	})
	return file_tests_proto_example_example1_proto_rawDescData
}

var file_tests_proto_example_example1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_proto_example_example1_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tests_proto_example_example1_proto_goTypes = []interface{}{
	(Enum1)(0),                     // 0: example1.Enum1
	(*Empty1)(nil),                 // 1: example1.Empty1
	(*Example1)(nil),               // 2: example1.Example1
	(*Message1)(nil),               // 3: example1.Message1
	nil,                            // 4: example1.Example1.MString1Entry
	nil,                            // 5: example1.Example1.MString2Entry
	nil,                            // 6: example1.Example1.MMessage1Entry
	nil,                            // 7: example1.Example1.MEnumEntry
	(*Message1_Embed1)(nil),        // 8: example1.Message1.Embed1
	(*Message1_Embed1_Embed2)(nil), // 9: example1.Message1.Embed1.Embed2
	(*durationpb.Duration)(nil),    // 10: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),  // 11: google.protobuf.Timestamp
	(*anypb.Any)(nil),              // 12: google.protobuf.Any
}
var file_tests_proto_example_example1_proto_depIdxs = []int32{
	0,  // 0: example1.Example1.f_enum1:type_name -> example1.Enum1
	0,  // 1: example1.Example1.f_enum2:type_name -> example1.Enum1
	0,  // 2: example1.Example1.f_enum5:type_name -> example1.Enum1
	0,  // 3: example1.Example1.f_enum6:type_name -> example1.Enum1
	10, // 4: example1.Example1.f_duration1:type_name -> google.protobuf.Duration
	11, // 5: example1.Example1.f_timestamp1:type_name -> google.protobuf.Timestamp
	3,  // 6: example1.Example1.f_message11:type_name -> example1.Message1
	8,  // 7: example1.Example1.f_message12:type_name -> example1.Message1.Embed1
	9,  // 8: example1.Example1.f_message13:type_name -> example1.Message1.Embed1.Embed2
	12, // 9: example1.Example1.f_any1:type_name -> google.protobuf.Any
	3,  // 10: example1.Example1.r_message:type_name -> example1.Message1
	0,  // 11: example1.Example1.r_enum:type_name -> example1.Enum1
	4,  // 12: example1.Example1.m_string1:type_name -> example1.Example1.MString1Entry
	5,  // 13: example1.Example1.m_string2:type_name -> example1.Example1.MString2Entry
	6,  // 14: example1.Example1.m_message1:type_name -> example1.Example1.MMessage1Entry
	7,  // 15: example1.Example1.m_enum:type_name -> example1.Example1.MEnumEntry
	3,  // 16: example1.Example1.MMessage1Entry.value:type_name -> example1.Message1
	0,  // 17: example1.Example1.MEnumEntry.value:type_name -> example1.Enum1
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_tests_proto_example_example1_proto_init() }
func file_tests_proto_example_example1_proto_init() {
	if File_tests_proto_example_example1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_example_example1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty1); i {
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
		file_tests_proto_example_example1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example1); i {
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
		file_tests_proto_example_example1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_tests_proto_example_example1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_tests_proto_example_example1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	file_tests_proto_example_example1_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Example1_FString2)(nil),
		(*Example1_FString3)(nil),
		(*Example1_FString4)(nil),
		(*Example1_FString5)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_proto_example_example1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_example_example1_proto_goTypes,
		DependencyIndexes: file_tests_proto_example_example1_proto_depIdxs,
		EnumInfos:         file_tests_proto_example_example1_proto_enumTypes,
		MessageInfos:      file_tests_proto_example_example1_proto_msgTypes,
	}.Build()
	File_tests_proto_example_example1_proto = out.File
	file_tests_proto_example_example1_proto_rawDesc = nil
	file_tests_proto_example_example1_proto_goTypes = nil
	file_tests_proto_example_example1_proto_depIdxs = nil
}

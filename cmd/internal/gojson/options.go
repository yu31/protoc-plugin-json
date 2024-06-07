package gojson

import (
	"reflect"
	"strings"
	"time"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

func getTagName(ot interface{}) string {
	name := reflect.TypeOf(ot).Elem().Name()
	name = strings.TrimSuffix(name, "_") // string is String_
	parts := strings.Split(name, "_")
	return strings.ToLower(parts[1])
}

func loadMessageOptions(msg *protogen.Message) *pbjson.MessageOptions {
	i := proto.GetExtension(msg.Desc.Options(), pbjson.E_Message)
	options := i.(*pbjson.MessageOptions)
	if options == nil {
		options = &pbjson.MessageOptions{}
	}
	return options
}
func loadFieldOptions(field *protogen.Field) *pbjson.FieldOptions {
	i := proto.GetExtension(field.Desc.Options(), pbjson.E_Field)
	options := i.(*pbjson.FieldOptions)
	if options == nil {
		options = &pbjson.FieldOptions{}
	}
	if options.Inline == nil { // default to false.
		v := false
		options.Inline = &v
	}
	return options
}

func loadOneOfOptions(field *protogen.Field) *pbjson.OneofOptions {
	i := proto.GetExtension(field.Oneof.Desc.Options(), pbjson.E_Oneof)
	options := i.(*pbjson.OneofOptions)
	if options == nil {
		options = &pbjson.OneofOptions{}
	}
	if options.Inline == nil { // default to false.
		v := false
		options.Inline = &v
	}
	return options
}

func loadTypeFormatRepeated(TypeFormat *pbjson.TypeFormat) *pbjson.TypeRepeated {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeRepeated{}
	}

	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Repeated:
		return ot.Repeated
	default:
		err := pkerror.New(
			"type repeated only support kind of TypeFormat <repeated> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatMap(TypeFormat *pbjson.TypeFormat) *pbjson.TypeMap {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeMap{}
	}

	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Map:
		return ot.Map
	default:
		err := pkerror.New(
			"type map only support kind of TypeFormat <map> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

func loadTypeFormatInt32(TypeFormat *pbjson.TypeFormat) *pbjson.TypeInt32 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeInt32{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Int32:
		return ot.Int32
	default:
		err := pkerror.New(
			"type int32 only support kind of TypeFormat <int32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatInt64(TypeFormat *pbjson.TypeFormat) *pbjson.TypeInt64 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeInt64{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Int64:
		return ot.Int64
	default:
		err := pkerror.New(
			"type int64 only support kind of TypeFormat <int64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatSInt32(TypeFormat *pbjson.TypeFormat) *pbjson.TypeSInt32 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeSInt32{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Sint32:
		return ot.Sint32
	default:
		err := pkerror.New(
			"type sint32 only support kind of TypeFormat <sint32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatSInt64(TypeFormat *pbjson.TypeFormat) *pbjson.TypeSInt64 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeSInt64{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Sint64:
		return ot.Sint64
	default:
		err := pkerror.New(
			"type sint64 only support kind of TypeFormat <sint64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatSFInt32(TypeFormat *pbjson.TypeFormat) *pbjson.TypeSFixed32 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeSFixed32{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Sfixed32:
		return ot.Sfixed32
	default:
		err := pkerror.New(
			"type sfixed32 only support kind of TypeFormat <sfixed32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatSFInt64(TypeFormat *pbjson.TypeFormat) *pbjson.TypeSFixed64 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeSFixed64{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Sfixed64:
		return ot.Sfixed64
	default:
		err := pkerror.New(
			"type sfixed64 only support kind of TypeFormat <sfixed64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatUint32(TypeFormat *pbjson.TypeFormat) *pbjson.TypeUint32 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeUint32{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Uint32:
		return ot.Uint32
	default:
		err := pkerror.New(
			"type uint32 only support kind of TypeFormat <uint32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatUint64(TypeFormat *pbjson.TypeFormat) *pbjson.TypeUint64 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeUint64{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Uint64:
		return ot.Uint64
	default:
		err := pkerror.New(
			"type uint64 only support kind of TypeFormat <uint64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatFixed32(TypeFormat *pbjson.TypeFormat) *pbjson.TypeFixed32 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeFixed32{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Fixed32:
		return ot.Fixed32
	default:
		err := pkerror.New(
			"type fixed32 only support kind of TypeFormat <fixed32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatFixed64(TypeFormat *pbjson.TypeFormat) *pbjson.TypeFixed64 {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeFixed64{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Fixed64:
		return ot.Fixed64
	default:
		err := pkerror.New(
			"type fixed64 only support kind of TypeFormat <fixed64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatFloat(TypeFormat *pbjson.TypeFormat) *pbjson.TypeFloat {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeFloat{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Float:
		return ot.Float
	default:
		err := pkerror.New(
			"type float only support kind of TypeFormat <float> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatDouble(TypeFormat *pbjson.TypeFormat) *pbjson.TypeDouble {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeDouble{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Double:
		return ot.Double
	default:
		err := pkerror.New(
			"type double only support kind of TypeFormat <double> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

func loadTypeFormatBool(TypeFormat *pbjson.TypeFormat) *pbjson.TypeBool {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeBool{}
	}
	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Bool:
		return ot.Bool
	default:
		err := pkerror.New(
			"type bool only support kind of TypeFormat <bool> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

//func loadTypeFormatString(TypeFormat *pbjson.TypeFormat) *pbjson.TypeString {
//	if TypeFormat == nil || TypeFormat.Kind == nil {
//		return &pbjson.TypeString{}
//	}
//	switch ot := TypeFormat.Kind.(type) {
//	case *pbjson.TypeFormat_String_:
//		return ot.String_
//	default:
//		err := pkerror.New(
//			"type string only support kind of TypeFormat <string> and you provided: <%s>",
//			getTagName(ot),
//		)
//		panic(err)
//	}
//}
//func loadTypeFormatBytes(TypeFormat *pbjson.TypeFormat) *pbjson.TypeBytes {
//	if TypeFormat == nil || TypeFormat.Kind == nil {
//		return &pbjson.TypeBytes{}
//	}
//	switch ot := TypeFormat.Kind.(type) {
//	case *pbjson.TypeFormat_Bytes:
//		return ot.Bytes
//	default:
//		err := pkerror.New(
//			"type bytes only support kind of TypeFormat <bytes> and you provided: <%s>",
//			getTagName(ot),
//		)
//		panic(err)
//	}
//}
//func loadTypeFormatMessage(TypeFormat *pbjson.TypeFormat) *pbjson.TypeMessage {
//	if TypeFormat == nil || TypeFormat.Kind == nil {
//		return &pbjson.TypeMessage{}
//	}
//	switch ot := TypeFormat.Kind.(type) {
//	case *pbjson.TypeFormat_Message:
//		return ot.Message
//	default:
//		err := pkerror.New(
//			"type message only support kind of TypeFormat <message> and you provided: <%s>",
//			getTagName(ot),
//		)
//		panic(err)
//	}
//}

func loadTypeFormatEnum(TypeFormat *pbjson.TypeFormat) *pbjson.TypeEnum {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeEnum{}
	}

	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Enum:
		return ot.Enum
	default:
		err := pkerror.New(
			"type google.protobuf.Any only support kind of TypeFormat <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatAny(TypeFormat *pbjson.TypeFormat) *pbjson.TypeAny {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeAny{}
	}

	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Any:
		return ot.Any
	default:
		err := pkerror.New(
			"type google.protobuf.Any only support kind of TypeFormat <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatDuration(TypeFormat *pbjson.TypeFormat) *pbjson.TypeDuration {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeDuration{}
	}

	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Duration:
		return ot.Duration
	default:
		err := pkerror.New(
			"type google.protobuf.Any only support kind of TypeFormat <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeFormatTimestamp(TypeFormat *pbjson.TypeFormat) *pbjson.TypeTimestamp {
	if TypeFormat == nil || TypeFormat.Kind == nil {
		return &pbjson.TypeTimestamp{Layout: &pbjson.TypeTimestamp_Layout{}}
	}

	switch ot := TypeFormat.Kind.(type) {
	case *pbjson.TypeFormat_Timestamp:
		rt := ot.Timestamp
		if rt.Layout == nil {
			rt.Layout = &pbjson.TypeTimestamp_Layout{}
		}
		if rt.Layout.Golang == "" {
			rt.Layout.Golang = time.RFC3339Nano
		}
		return rt
	default:
		err := pkerror.New(
			"type google.protobuf.Any only support kind of TypeFormat <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

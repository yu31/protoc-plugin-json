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
	// Ignore field if json is "-"
	if options.Json != nil && *options.Json == "-" {
		options.Ignore = true
	}
	return options
}

func loadOneOfOptions(field *protogen.Field) *pbjson.OneofOptions {
	i := proto.GetExtension(field.Oneof.Desc.Options(), pbjson.E_Oneof)
	options := i.(*pbjson.OneofOptions)
	if options == nil {
		options = &pbjson.OneofOptions{}
	}
	// Ignore field if json is "-"
	if options.Json != nil && *options.Json == "-" {
		options.Ignore = true
	}
	return options
}

func loadPlainOptions(field *protogen.Field, options *pbjson.FieldOptions) *pbjson.PlainOptions {
	if options.Reference == nil || options.Reference.Kind == nil {
		return &pbjson.PlainOptions{}
	}

	switch ot := options.Reference.Kind.(type) {
	case *pbjson.TypeReference_Plain:
		return ot.Plain
	default:
		err := pkerror.New(
			"type %s only supports kind of TypeReference <plain> and you provided: <%s>",
			field.Desc.Kind().String(), getTagName(ot),
		)
		panic(err)
	}
}
func loadRepeatedOptions(field *protogen.Field, options *pbjson.FieldOptions) *pbjson.RepeatedOptions {
	if options.Reference == nil || options.Reference.Kind == nil {
		return &pbjson.RepeatedOptions{}
	}

	switch ot := options.Reference.Kind.(type) {
	case *pbjson.TypeReference_Repeated:
		return ot.Repeated
	default:
		err := pkerror.New(
			"type %s only supports kind of TypeReference <repeated> and you provided: <%s>",
			field.Desc.Kind().String(), getTagName(ot),
		)
		panic(err)
	}
}
func loadMapOptions(field *protogen.Field, options *pbjson.FieldOptions) *pbjson.MapOptions {
	if options.Reference == nil || options.Reference.Kind == nil {
		return &pbjson.MapOptions{}
	}

	switch ot := options.Reference.Kind.(type) {
	case *pbjson.TypeReference_Map:
		return ot.Map
	default:
		err := pkerror.New(
			"type %s only supports kind of TypeReference <map> and you provided: <%s>",
			field.Desc.Kind().String(), getTagName(ot),
		)
		panic(err)
	}
}

func loadTypeCodecInt32(typeCodec *pbjson.TypeCodec) *pbjson.TypeInt32 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeInt32{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Int32:
		return ot.Int32
	default:
		err := pkerror.New(
			"type int32 only supports kind of TypeCodec <int32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecInt64(typeCodec *pbjson.TypeCodec) *pbjson.TypeInt64 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeInt64{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Int64:
		return ot.Int64
	default:
		err := pkerror.New(
			"type int64 only supports kind of TypeCodec <int64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecSInt32(typeCodec *pbjson.TypeCodec) *pbjson.TypeSInt32 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeSInt32{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Sint32:
		return ot.Sint32
	default:
		err := pkerror.New(
			"type sint32 only supports kind of TypeCodec <sint32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecSInt64(typeCodec *pbjson.TypeCodec) *pbjson.TypeSInt64 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeSInt64{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Sint64:
		return ot.Sint64
	default:
		err := pkerror.New(
			"type sint64 only supports kind of TypeCodec <sint64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecSFInt32(typeCodec *pbjson.TypeCodec) *pbjson.TypeSFixed32 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeSFixed32{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Sfixed32:
		return ot.Sfixed32
	default:
		err := pkerror.New(
			"type sfixed32 only supports kind of TypeCodec <sfixed32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecSFInt64(typeCodec *pbjson.TypeCodec) *pbjson.TypeSFixed64 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeSFixed64{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Sfixed64:
		return ot.Sfixed64
	default:
		err := pkerror.New(
			"type sfixed64 only supports kind of TypeCodec <sfixed64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecUint32(typeCodec *pbjson.TypeCodec) *pbjson.TypeUint32 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeUint32{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Uint32:
		return ot.Uint32
	default:
		err := pkerror.New(
			"type uint32 only supports kind of TypeCodec <uint32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecUint64(typeCodec *pbjson.TypeCodec) *pbjson.TypeUint64 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeUint64{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Uint64:
		return ot.Uint64
	default:
		err := pkerror.New(
			"type uint64 only supports kind of TypeCodec <uint64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecFixed32(typeCodec *pbjson.TypeCodec) *pbjson.TypeFixed32 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeFixed32{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Fixed32:
		return ot.Fixed32
	default:
		err := pkerror.New(
			"type fixed32 only supports kind of TypeCodec <fixed32> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecFixed64(typeCodec *pbjson.TypeCodec) *pbjson.TypeFixed64 {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeFixed64{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Fixed64:
		return ot.Fixed64
	default:
		err := pkerror.New(
			"type fixed64 only supports kind of TypeCodec <fixed64> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecFloat(typeCodec *pbjson.TypeCodec) *pbjson.TypeFloat {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeFloat{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Float:
		return ot.Float
	default:
		err := pkerror.New(
			"type float only supports kind of TypeCodec <float> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecDouble(typeCodec *pbjson.TypeCodec) *pbjson.TypeDouble {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeDouble{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Double:
		return ot.Double
	default:
		err := pkerror.New(
			"type double only supports kind of TypeCodec <double> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

func loadTypeCodecBool(typeCodec *pbjson.TypeCodec) *pbjson.TypeBool {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeBool{}
	}
	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Bool:
		return ot.Bool
	default:
		err := pkerror.New(
			"type bool only supports kind of TypeCodec <bool> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

func loadTypeCodecEnum(typeCodec *pbjson.TypeCodec) *pbjson.TypeEnum {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeEnum{}
	}

	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Enum:
		return ot.Enum
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeCodec <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecAny(typeCodec *pbjson.TypeCodec) *pbjson.TypeAny {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeAny{}
	}

	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Any:
		return ot.Any
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeCodec <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecDuration(typeCodec *pbjson.TypeCodec) *pbjson.TypeDuration {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeDuration{}
	}

	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Duration:
		return ot.Duration
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeCodec <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeCodecTimestamp(typeCodec *pbjson.TypeCodec) *pbjson.TypeTimestamp {
	if typeCodec == nil || typeCodec.Kind == nil {
		return &pbjson.TypeTimestamp{}
	}

	switch ot := typeCodec.Kind.(type) {
	case *pbjson.TypeCodec_Timestamp:
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
			"type google.protobuf.Any only supports kind of TypeCodec <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

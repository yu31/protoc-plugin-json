package gojson

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"github.com/yu31/protoc-kit-go/helper/pkwkt"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

func genMessageID(file *protogen.File, msg *protogen.Message) string {
	filePath := file.Desc.Path()
	msgName := string(msg.Desc.FullName())
	if index := strings.Index(msgName, "."); index > 0 {
		msgName = msgName[index+1:]
	}
	id := fmt.Sprintf("[ERROR] - [plugin: gojson] - [ file: %s | message: %s ]", filePath, msgName)
	return id
}

func fieldGoType(protoGen *protogen.GeneratedFile, field *protogen.Field) (goType string) {
	if field.Desc.IsWeak() {
		//return "struct{}", false
		panic(fmt.Sprintf("unsupported case IsWeak; field: %s", field.Desc.FullName()))
	}

	switch field.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = protoGen.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = protoGen.QualifiedGoIdent(field.Message.GoIdent)
	}
	return goType
}

func loadOneOfFields(fieldSets []*FieldSet) (oneOfFields []*FieldSet) {
	for _, fieldSet := range fieldSets {
		if fieldSet.FieldIsInline() {
			oneOfFields = append(oneOfFields, loadOneOfFields(fieldSet.InlineSet.Childs)...)
		}
		if fieldSet.IsOneOfField() && len(fieldSet.OneOfSet.Parts) != 0 {
			oneOfFields = append(oneOfFields, loadOneOfFields(fieldSet.OneOfSet.Parts)...)
			oneOfFields = append(oneOfFields, fieldSet)
		}
	}
	return
}

func loadInlineFields(fieldSets []*FieldSet) (inlineFields []*FieldSet) {
	_inlineFields := _loadInlineFields(fieldSets)
	length := len(_inlineFields)
	if length == 0 {
		return
	}
	// Deduplication
	_exists := make(map[*FieldSet]struct{}, length)
	inlineFields = make([]*FieldSet, 0, length)
	for i := len(_inlineFields) - 1; i >= 0; i-- {
		fieldSet := _inlineFields[i]
		parent := fieldSet.ParentField()
		if _, ok := _exists[parent]; !ok {
			inlineFields = append(inlineFields, fieldSet)
			_exists[parent] = struct{}{}
		}
	}
	return
}
func _loadInlineFields(fieldSets []*FieldSet) (parentFields []*FieldSet) {
	for _, fieldSet := range fieldSets {
		if fieldSet.FieldIsInline() {
			parentFields = append(parentFields, _loadInlineFields(fieldSet.InlineSet.Childs)...)
		}
		if fieldSet.IsOneOfField() && len(fieldSet.OneOfSet.Parts) != 0 {
			parentFields = append(parentFields, _loadInlineFields(fieldSet.OneOfSet.Parts)...)
		}
		if parent := fieldSet.ParentField(); parent != nil {
			parentFields = append(parentFields, fieldSet)
		}
	}
	return
}

func loadMapKeyTypeInfo(field *protogen.Field, typeFormat *pbjson.TypeFormat) (typeName string, funcVars []string) {
	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.StringKind:
		typeName = "Str"
	case protoreflect.BoolKind:
		quote := loadTypeFormatBool(typeFormat).Codec != pbjson.TypeBool_Bool
		typeName = "Bool"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int32Kind:
		quote := loadTypeFormatInt32(typeFormat).Codec != pbjson.TypeInt32_Numeric
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int64Kind:
		quote := loadTypeFormatInt64(typeFormat).Codec != pbjson.TypeInt64_Numeric
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint32Kind:
		quote := loadTypeFormatSInt32(typeFormat).Codec != pbjson.TypeSInt32_Numeric
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint64Kind:
		quote := loadTypeFormatSInt64(typeFormat).Codec != pbjson.TypeSInt64_Numeric
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed32Kind:
		quote := loadTypeFormatSFInt32(typeFormat).Codec != pbjson.TypeSFixed32_Numeric
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed64Kind:
		quote := loadTypeFormatSFInt64(typeFormat).Codec != pbjson.TypeSFixed64_Numeric
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint32Kind:
		quote := loadTypeFormatUint32(typeFormat).Codec != pbjson.TypeUint32_Numeric
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint64Kind:
		quote := loadTypeFormatUint64(typeFormat).Codec != pbjson.TypeUint64_Numeric
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed32Kind:
		quote := loadTypeFormatFixed32(typeFormat).Codec != pbjson.TypeFixed32_Numeric
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed64Kind:
		quote := loadTypeFormatFixed64(typeFormat).Codec != pbjson.TypeFixed64_Numeric
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	default:
		err := pkerror.New("unsupported kind of %s as map key", kind.String())
		panic(err)
	}
	return
}
func loadValueTypeInfo(field *protogen.Field, typeFormat *pbjson.TypeFormat) (typeName string, funcVars []string) {
	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.Int32Kind:
		quote := loadTypeFormatInt32(typeFormat).Codec == pbjson.TypeInt32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int64Kind:
		quote := loadTypeFormatInt64(typeFormat).Codec == pbjson.TypeInt64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint32Kind:
		quote := loadTypeFormatSInt32(typeFormat).Codec == pbjson.TypeSInt32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint64Kind:
		quote := loadTypeFormatSInt64(typeFormat).Codec == pbjson.TypeSInt64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed32Kind:
		quote := loadTypeFormatSFInt32(typeFormat).Codec == pbjson.TypeSFixed32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed64Kind:
		quote := loadTypeFormatSFInt64(typeFormat).Codec == pbjson.TypeSFixed64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint32Kind:
		quote := loadTypeFormatUint32(typeFormat).Codec == pbjson.TypeUint32_String
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint64Kind:
		quote := loadTypeFormatUint64(typeFormat).Codec == pbjson.TypeUint64_String
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed32Kind:
		quote := loadTypeFormatFixed32(typeFormat).Codec == pbjson.TypeFixed32_String
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed64Kind:
		quote := loadTypeFormatFixed64(typeFormat).Codec == pbjson.TypeFixed64_String
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.FloatKind:
		quote := loadTypeFormatFloat(typeFormat).Codec == pbjson.TypeFloat_String
		typeName = "F32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.DoubleKind:
		quote := loadTypeFormatDouble(typeFormat).Codec == pbjson.TypeDouble_String
		typeName = "F64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.BoolKind:
		quote := loadTypeFormatBool(typeFormat).Codec == pbjson.TypeBool_String
		typeName = "Bool"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.StringKind:
		typeName = "Str"
	case protoreflect.BytesKind:
		typeName = "Bytes"
	case protoreflect.EnumKind:
		// FIXME: 需要补充指针型变量的测试.
		typeSet := loadTypeFormatEnum(typeFormat)
		switch typeSet.Codec {
		case pbjson.TypeEnum_Unset, pbjson.TypeEnum_Numeric:
			typeName = "EnumNum"
			funcVars = append(funcVars, strconv.FormatBool(false))
		case pbjson.TypeEnum_NumericString:
			typeName = "EnumNum"
			funcVars = append(funcVars, strconv.FormatBool(true))
		case pbjson.TypeEnum_EnumString:
			typeName = "EnumStr"
		}
	case protoreflect.MessageKind:
		switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
		case pkwkt.Any:
			typeAny := loadTypeFormatAny(typeFormat)
			switch typeAny.Codec {
			case pbjson.TypeAny_Object, pbjson.TypeAny_Unset:
				typeName = "WKTAnyObject"
			case pbjson.TypeAny_Proto:
				typeName = "WKTAnyProto"
			}
		case pkwkt.Duration:
			typeDuration := loadTypeFormatDuration(typeFormat)
			switch typeDuration.Codec {
			case pbjson.TypeDuration_Object, pbjson.TypeDuration_Unset:
				typeName = "WKTDurObject"
			case pbjson.TypeDuration_TimeString:
				typeName = "WKTDurTimeStr"
			case pbjson.TypeDuration_Nanosecond:
				typeName = "WKTDurNano"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeDuration_NanosecondString:
				typeName = "WKTDurNano"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeDuration_Microsecond:
				typeName = "WKTDurMicro"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeDuration_MicrosecondString:
				typeName = "WKTDurMicro"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeDuration_Millisecond:
				typeName = "WKTDurMilli"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeDuration_MillisecondString:
				typeName = "WKTDurMilli"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeDuration_Second:
				typeName = "WKTDurSecond"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeDuration_SecondString:
				typeName = "WKTDurSecond"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeDuration_Minute:
				typeName = "WKTDurMinute"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeDuration_MinuteString:
				typeName = "WKTDurMinute"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeDuration_Hour:
				typeName = "WKTDurHour"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeDuration_HourString:
				typeName = "WKTDurHour"
				funcVars = append(funcVars, strconv.FormatBool(true))
			}
		case pkwkt.Timestamp:
			typeTimestamp := loadTypeFormatTimestamp(typeFormat)
			switch typeTimestamp.Codec {
			case pbjson.TypeTimestamp_Object, pbjson.TypeTimestamp_Unset:
				typeName = "WKTTsObject"
			case pbjson.TypeTimestamp_TimeLayout:
				typeSet := loadTypeFormatTimestamp(typeFormat)
				// FIXME: valid the layout.
				layout := strconv.Quote(typeSet.Layout.Golang)
				typeName = "WKTTsLayout"
				funcVars = append(funcVars, layout)
			case pbjson.TypeTimestamp_UnixNano:
				typeName = "WKTTsUnixNano"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeTimestamp_UnixNanoString:
				typeName = "WKTTsUnixNano"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeTimestamp_UnixMicro:
				typeName = "WKTTsUnixMicro"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeTimestamp_UnixMicroString:
				typeName = "WKTTsUnixMicro"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeTimestamp_UnixMilli:
				typeName = "WKTTsUnixMilli"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeTimestamp_UnixMilliString:
				typeName = "WKTTsUnixMilli"
				funcVars = append(funcVars, strconv.FormatBool(true))
			case pbjson.TypeTimestamp_UnixSec:
				typeName = "WKTTsUnixSec"
				funcVars = append(funcVars, strconv.FormatBool(false))
			case pbjson.TypeTimestamp_UnixSecString:
				typeName = "WKTTsUnixSec"
				funcVars = append(funcVars, strconv.FormatBool(true))
			}
		default:
			typeName = "Message"
		}
	default:
		err := pkerror.New("marshal: unsupported kind of %s as value", kind.String())
		panic(err)
	}
	return
}

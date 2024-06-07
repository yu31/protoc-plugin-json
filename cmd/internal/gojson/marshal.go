package gojson

import (
	"os"
	"strconv"
	"strings"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"github.com/yu31/protoc-kit-go/helper/pkwkt"
	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/cmd/internal/gojson/pkg/importpkg"
	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

type Marshal struct {
	g    *protogen.GeneratedFile
	file *protogen.File

	// The message of currently being processed.
	message *protogen.Message

	// The message options of currently being processed.
	options *pbjson.MessageOptions
}

func (ml *Marshal) guessBufLen(fieldSets []*FieldSet) (bufLen int) {
	// Sum json beginning(`{`) and closing(`}`).
	bufLen += 2

	for _, fieldSet := range fieldSets {
		if fieldSet.IsOneOfField() {
			for _, oneSet := range fieldSet.OneOfSet.Parts {
				// Sum key length, colon separator(`:`) and comma separator(`,`).
				bufLen += len(oneSet.JSONKey) + 2
			}
		}
		// Sum key length, colon separator(`:`) and comma separator(`,`).
		bufLen += len(fieldSet.JSONKey) + 2
	}
	bufLen = truncateBufLen(bufLen * 2)
	return
}

func (ml *Marshal) GenerateCode(fieldSets []*FieldSet) {
	msg := ml.message
	bufLen := ml.guessBufLen(fieldSets)

	ml.g.P("// MarshalJSON implements interface json.Marshaler for proto message ", msg.Desc.Name(), " in file ", ml.file.Desc.Path())
	ml.g.P("func (x *", msg.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
	{
		ml.g.P("if x == nil {")
		ml.g.P(`	return []byte("null"), nil`)
		ml.g.P("}")
		if len(fieldSets) == 0 {
			ml.g.P(`return []byte("{}"), nil`)
		} else {
			ml.g.P("enc := ", importpkg.PJEncoder.Ident("New"), "(", bufLen, ")")
			ml.g.P("enc.AppendObjectBegin() // Add begin JSON identifier")
			ml.g.P("")
			for _, fieldSet := range fieldSets {
				catch := pkerror.Recover("gojson", ml.file, ml.message, fieldSet.Field, func() {
					ml.marshalField(fieldSet)
				})
				if catch {
					os.Exit(1)
				}
			}
			ml.g.P("enc.AppendObjectEnd() // Add end JSON identifier")
			ml.g.P("return enc.Bytes(), nil")
		}
	}
	// End of MarshalJSON.
	ml.g.P("}")
}

func (ml *Marshal) marshalField(fieldSet *FieldSet) {
	if fieldSet.FieldIsInline() && len(fieldSet.InlineSet.Childs) != 0 {
		// Confirm the parent field is not null
		fieldAlias := fieldSet.Alias()
		fieldValue := fieldSet.Value()
		ml.g.P("if ", fieldAlias, " := ", fieldValue, "; ", fieldAlias, " != nil {")
		for _, childField := range fieldSet.InlineSet.Childs {
			ml.marshalField(childField)
		}
		ml.g.P("}")
		return
	}

	switch {
	case fieldSet.IsOneOfField():
		ml.marshalOneOf(fieldSet)
	default:
		ml.encodeField(fieldSet)
	}
}

func (ml *Marshal) marshalOneOf(fieldSet *FieldSet) {
	if fieldSet.OneOfIsInline() {
		ml.marshalParts(fieldSet)
		return
	}

	fieldValue := fieldSet.Value()
	if options := fieldSet.OneOfSet.Options; options.Omitempty {
		ml.g.P("if ", fieldValue, " != nil {")
		{
			ml.g.P("enc.AppendObjectKey(", strconv.Quote(fieldSet.JSONKey), ")")
			ml.g.P("enc.AppendObjectBegin()")
			ml.marshalParts(fieldSet)
			ml.g.P("enc.AppendObjectEnd()")
		}
		ml.g.P("}")
	} else {
		ml.g.P("enc.AppendObjectKey(", strconv.Quote(fieldSet.JSONKey), ")")
		ml.g.P("if ", fieldValue, " != nil {") // // Confirm the oneof field is not null
		{
			ml.g.P("enc.AppendObjectBegin()")
			ml.marshalParts(fieldSet)
			ml.g.P("enc.AppendObjectEnd()")
		}
		ml.g.P("} else {")
		{
			ml.g.P("enc.AppendValNULL()")
		}
		ml.g.P("}")
	}
}
func (ml *Marshal) marshalParts(fieldSet *FieldSet) {
	partFields := fieldSet.OneOfSet.Parts
	if len(partFields) == 0 {
		return
	}

	fieldAlias := fieldSet.Alias()
	fieldValue := fieldSet.Value()
	ml.g.P("switch ", fieldAlias, " := ", fieldValue, ".(type) {")
	for _, partField := range partFields {
		partType := ml.g.QualifiedGoIdent(partField.Field.GoIdent)
		ml.g.P("case *", partType, ":")
		ml.marshalField(partField)
	}
	//ml.g.P("default:")
	//// To avoids unused panics. Cases occur in the all parts in oneof field are ignored.
	//ml.g.P("  _ = ", fieldAlias)
	ml.g.P("}")
}

func (ml *Marshal) encodeField(fieldSet *FieldSet) {
	var (
		funcName   string
		checkError bool
	)

	options := fieldSet.Options
	funcVars := []string{
		"enc",
		strconv.Quote(fieldSet.JSONKey),
		fieldSet.Value(),
		strconv.FormatBool(options.Omitempty),
	}

	switch {
	case fieldSet.Field.Desc.IsMap():
		formatMap := loadTypeFormatMap(options.Format)
		_mkName, _mkVars := ml.loadMapKeyTypeInfo(fieldSet.Field.Message.Fields[0], formatMap.Key)
		_mvName, _mvVars, _checkError := ml.loadValueTypeInfo(fieldSet.Field.Message.Fields[1], formatMap.Value)

		funcName = "AppendMap" + _mkName + _mvName
		checkError = _checkError
		funcVars = append(funcVars, _mkVars...)
		funcVars = append(funcVars, _mvVars...)
	case fieldSet.Field.Desc.IsList():
		formatRepeated := loadTypeFormatRepeated(options.Format)
		_typeName, _vars, _checkError := ml.loadValueTypeInfo(fieldSet.Field, formatRepeated.Elem)

		funcName = "AppendList" + _typeName
		checkError = _checkError
		funcVars = append(funcVars, _vars...)
	default:
		_typeName, _vars, _checkError := ml.loadValueTypeInfo(fieldSet.Field, options.Format)

		if pkfield.FieldIsOptional(fieldSet.Field) {
			funcName = "AppendPtr" + _typeName
		} else {
			funcName = "AppendVal" + _typeName
		}
		checkError = _checkError
		funcVars = append(funcVars, _vars...)
	}

	funcIdent := importpkg.PJEncoder.Ident(funcName)
	parameters := strings.Join(funcVars, ",")
	if checkError {
		ml.g.P("if err := ", funcIdent, "(", parameters, "); err != nil {")
		ml.g.P("    return nil, err")
		ml.g.P("}")
	} else {
		ml.g.P(funcIdent, "(", parameters, ")")
	}
}

func (ml *Marshal) loadMapKeyTypeInfo(field *protogen.Field, TypeFormat *pbjson.TypeFormat,
) (typeName string, funcVars []string) {
	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.StringKind:
		typeName = "Str"
	case protoreflect.BoolKind:
		quote := loadTypeFormatBool(TypeFormat).Codec != pbjson.TypeBool_Bool
		typeName = "Bool"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int32Kind:
		quote := loadTypeFormatInt32(TypeFormat).Codec != pbjson.TypeInt32_Numeric
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int64Kind:
		quote := loadTypeFormatInt64(TypeFormat).Codec != pbjson.TypeInt64_Numeric
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint32Kind:
		quote := loadTypeFormatSInt32(TypeFormat).Codec != pbjson.TypeSInt32_Numeric
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint64Kind:
		quote := loadTypeFormatSInt64(TypeFormat).Codec != pbjson.TypeSInt64_Numeric
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed32Kind:
		quote := loadTypeFormatSFInt32(TypeFormat).Codec != pbjson.TypeSFixed32_Numeric
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed64Kind:
		quote := loadTypeFormatSFInt64(TypeFormat).Codec != pbjson.TypeSFixed64_Numeric
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint32Kind:
		quote := loadTypeFormatUint32(TypeFormat).Codec != pbjson.TypeUint32_Numeric
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint64Kind:
		quote := loadTypeFormatUint64(TypeFormat).Codec != pbjson.TypeUint64_Numeric
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed32Kind:
		quote := loadTypeFormatFixed32(TypeFormat).Codec != pbjson.TypeFixed32_Numeric
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed64Kind:
		quote := loadTypeFormatFixed64(TypeFormat).Codec != pbjson.TypeFixed64_Numeric
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	default:
		err := pkerror.New("marshal: unsupported kind of %s as map key", kind.String())
		panic(err)
	}
	return
}

func (ml *Marshal) loadValueTypeInfo(field *protogen.Field, TypeFormat *pbjson.TypeFormat,
) (typeName string, funcVars []string, checkError bool) {
	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.Int32Kind:
		quote := loadTypeFormatInt32(TypeFormat).Codec == pbjson.TypeInt32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int64Kind:
		quote := loadTypeFormatInt64(TypeFormat).Codec == pbjson.TypeInt64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint32Kind:
		quote := loadTypeFormatSInt32(TypeFormat).Codec == pbjson.TypeSInt32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint64Kind:
		quote := loadTypeFormatSInt64(TypeFormat).Codec == pbjson.TypeSInt64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed32Kind:
		quote := loadTypeFormatSFInt32(TypeFormat).Codec == pbjson.TypeSFixed32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed64Kind:
		quote := loadTypeFormatSFInt64(TypeFormat).Codec == pbjson.TypeSFixed64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint32Kind:
		quote := loadTypeFormatUint32(TypeFormat).Codec == pbjson.TypeUint32_String
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint64Kind:
		quote := loadTypeFormatUint64(TypeFormat).Codec == pbjson.TypeUint64_String
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed32Kind:
		quote := loadTypeFormatFixed32(TypeFormat).Codec == pbjson.TypeFixed32_String
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed64Kind:
		quote := loadTypeFormatFixed64(TypeFormat).Codec == pbjson.TypeFixed64_String
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.FloatKind:
		quote := loadTypeFormatFloat(TypeFormat).Codec == pbjson.TypeFloat_String
		typeName = "F32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.DoubleKind:
		quote := loadTypeFormatDouble(TypeFormat).Codec == pbjson.TypeDouble_String
		typeName = "F64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.BoolKind:
		quote := loadTypeFormatBool(TypeFormat).Codec == pbjson.TypeBool_String
		typeName = "Bool"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.StringKind:
		typeName = "Str"
	case protoreflect.BytesKind:
		typeName = "Bytes"
		checkError = true
	case protoreflect.EnumKind:
		// FIXME: 需要补充指针型变量的测试.
		typeSet := loadTypeFormatEnum(TypeFormat)
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
			typeAny := loadTypeFormatAny(TypeFormat)
			switch typeAny.Codec {
			case pbjson.TypeAny_Object, pbjson.TypeAny_Unset:
				typeName = "WKTAnyObject"
				checkError = true
			case pbjson.TypeAny_Proto:
				typeName = "WKTAnyProto"
				checkError = true
			}
		case pkwkt.Duration:
			typeDuration := loadTypeFormatDuration(TypeFormat)
			switch typeDuration.Codec {
			case pbjson.TypeDuration_Object, pbjson.TypeDuration_Unset:
				typeName = "WKTDurObject"
				checkError = true
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
			typeTimestamp := loadTypeFormatTimestamp(TypeFormat)
			switch typeTimestamp.Codec {
			case pbjson.TypeTimestamp_Object, pbjson.TypeTimestamp_Unset:
				typeName = "WKTTsObject"
				checkError = true
			case pbjson.TypeTimestamp_TimeLayout:
				typeSet := loadTypeFormatTimestamp(TypeFormat)
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
			checkError = true
		}
	default:
		err := pkerror.New("marshal: unsupported kind of %s as value", kind.String())
		panic(err)
	}
	return
}

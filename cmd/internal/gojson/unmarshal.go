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

type Unmarshal struct {
	g    *protogen.GeneratedFile
	file *protogen.File

	// The message of currently being processed.
	message *protogen.Message

	// The message options of currently being processed.
	options *pbjson.MessageOptions
}

func (ul *Unmarshal) GenerateCode(fieldSets []*FieldSet) {
	msg := ul.message

	ul.g.P("// UnmarshalJSON implements json.Unmarshaler for proto message ", msg.Desc.Name(), " in file ", ul.file.Desc.Path())
	ul.g.P("func (x *", msg.GoIdent.GoName, ") UnmarshalJSON(b []byte) error {")
	{
		ul.g.P("if x == nil {")
		// FIXME: review the error message.
		ul.g.P("	return ", importpkg.Errors.Ident("New"), `("json: Unmarshal: `, string(msg.GoIdent.GoImportPath), ".(*", msg.GoIdent.GoName, `) is nil")`)
		ul.g.P("}")
		// Generated codes of loop scan.
		ul.unmarshalLoop(fieldSets)
		ul.g.P("return nil")
	}
	ul.g.P("}")
}

func (ul *Unmarshal) unmarshalLoop(fieldSets []*FieldSet) {
	if len(fieldSets) == 0 {
		return
	}

	// pre-check
	ul.g.P("var (")
	ul.g.P("    err error")
	ul.g.P("    isNULL bool")
	ul.g.P("    dec *", importpkg.PJDecoder.Ident("Decoder"))
	ul.g.P(")")

	ul.g.P("if dec, err = ", importpkg.PJDecoder.Ident("New"), "(b); err != nil {")
	ul.g.P("    return err")
	ul.g.P("}")

	ul.g.P("if isNULL, err = dec.BeforeScanJSON(); err != nil || isNULL {")
	ul.g.P("    return err")
	ul.g.P("}")

	ul.declareLocalVariables(fieldSets)

	ul.g.P("LOOP_SCAN:")
	ul.g.P("for { // Loop to read the JSON objects")
	ul.unmarshalRead(fieldSets)
	ul.g.P("}")
}
func (ul *Unmarshal) unmarshalRead(fieldSets []*FieldSet) {
	ul.g.P("var (")
	ul.g.P("    jsonKey string")
	ul.g.P("    isEnd bool")
	ul.g.P(")")
	ul.g.P("")
	ul.g.P("if isEnd, err = dec.BeforeScanNext(); err != nil {")
	ul.g.P("    return err")
	ul.g.P("}")
	ul.g.P("if isEnd {")
	ul.g.P("    break LOOP_SCAN")
	ul.g.P("}")
	ul.g.P("")
	// The `jsonKey` is unsafe after unmarshal end.
	ul.g.P("if jsonKey, err = dec.ReadJSONKey(); err != nil {")
	ul.g.P("    return err")
	ul.g.P("}")
	ul.g.P("switch jsonKey { // match the jsonKey")
	{
		for _, fieldSet := range fieldSets {
			catch := pkerror.Recover("gojson", ul.file, ul.message, fieldSet.Field, func() {
				ul.unmarshalField(fieldSet)
			})
			if catch {
				os.Exit(1)
			}
		}
		ul.g.P("default:")
		if ul.options.DisallowUnknownFields {
			// FIXME: optimized error.
			ul.g.P("return ", importpkg.PJDecoder.Ident("ErrUnknownField"), "(dec)")
		} else {
			ul.g.P("if err = dec.DiscardValue(); err != nil {")
			ul.g.P("    return err")
			ul.g.P("}")
		}
	}
	ul.g.P("} // end switch")
}

func (ul *Unmarshal) unmarshalField(fieldSet *FieldSet) {
	if fieldSet.FieldIsInline() {
		for _, childField := range fieldSet.InlineSet.Childs {
			ul.unmarshalField(childField)
		}
		return
	}

	if fieldSet.IsOneOfField() && fieldSet.OneOfIsInline() {
		for _, oneOfPart := range fieldSet.OneOfSet.Parts {
			ul.unmarshalField(oneOfPart)
		}
		return
	}

	ul.g.P("case ", strconv.Quote(fieldSet.JSONKey), ":")

	ul.prepareInitParentField(fieldSet)
	ul.prepareInitOneOfPart(fieldSet)

	switch {
	case fieldSet.IsOneOfField():
		ul.unmarshalOneOf(fieldSet)
	default:
		ul.decodeField(fieldSet)
	}
}

func (ul *Unmarshal) unmarshalOneOf(fieldSet *FieldSet) {
	// Read and parse oneof.
	ul.g.P("if isNULL, err = dec.BeforeReadObject(); err != nil {")
	ul.g.P("	return err")
	ul.g.P("}")

	ul.g.P("if isNULL {") // if started.
	{
		ul.g.P(fieldSet.Value(), " = nil")
	}
	ul.g.P("} else {")
	{
		oneOfKeyName := "oneOfKey" + fieldSet.ID()
		ul.g.P("for { // Loop to read oneof fields") // for started.
		{
			ul.g.P("if isEnd, err = dec.BeforeReadNext(); err != nil {")
			ul.g.P("	return err ")
			ul.g.P("}")
			ul.g.P("if isEnd {")
			ul.g.P("	break")
			ul.g.P("}")
			ul.g.P("var ", oneOfKeyName, " string")
			// The `oneOfKeyName` is unsafe after unmarshal end.
			ul.g.P("if ", oneOfKeyName, ", err = dec.ReadObjectKey(); err != nil {")
			ul.g.P("	return err")
			ul.g.P("}")
			ul.g.P("switch ", oneOfKeyName, " { // match the oneof key") // switch started.
			for _, oneOfPart := range fieldSet.OneOfSet.Parts {
				ul.unmarshalField(oneOfPart)
			}
			ul.g.P("default:")
			if ul.options.DisallowUnknownFields { // FIXME: control by oneof options?
				ul.g.P("return ", importpkg.PJDecoder.Ident("ErrUnknownOneOfField"), "(dec, ", oneOfKeyName, ")")
			} else {
				ul.g.P("if err = dec.DiscardValue(); err != nil {")
				ul.g.P("    return err")
				ul.g.P("}")
			}
			ul.g.P("}") // switch ended.
		}
		ul.g.P("}") // for ended.
	}
	ul.g.P("}") // if ended.
}

func (ul *Unmarshal) decodeField(fieldSet *FieldSet) {
	funcVars := []string{
		"dec",
	}

	var funcName string

	options := fieldSet.Options
	fieldValue := fieldSet.Value()

	switch {
	case fieldSet.Field.Desc.IsMap():
		mapOptions := loadMapOptions(fieldSet.Field, options)
		_mkName, _mkVars := ul.loadMapKeyTypeInfo(fieldSet.Field.Message.Fields[0], mapOptions.Key)
		_mvName, _mvVars := ul.loadValueTypeInfo(fieldSet.Field.Message.Fields[1], mapOptions.Value)

		funcName = "ReadMap" + _mkName + _mvName
		funcVars = append(funcVars, fieldValue)
		funcVars = append(funcVars, _mkVars...)
		funcVars = append(funcVars, _mvVars...)
	case fieldSet.Field.Desc.IsList():
		repeatedOptions := loadRepeatedOptions(fieldSet.Field, options)
		_typeName, _vars := ul.loadValueTypeInfo(fieldSet.Field, repeatedOptions.Elem)

		funcName = "ReadList" + _typeName
		funcVars = append(funcVars, fieldValue)
		funcVars = append(funcVars, _vars...)
	default:
		plainOptions := loadPlainOptions(fieldSet.Field, options)
		_typeName, _vars := ul.loadValueTypeInfo(fieldSet.Field, plainOptions.Value)

		if pkfield.FieldIsOptional(fieldSet.Field) {
			funcName = "ReadPtr" + _typeName
		} else {
			funcName = "ReadVal" + _typeName
		}
		switch fieldSet.Field.Desc.Kind() {
		case protoreflect.EnumKind, protoreflect.MessageKind:
			funcVars = append(funcVars, fieldValue)
		}
		funcVars = append(funcVars, _vars...)
	}

	funcGoIdent := importpkg.PJDecoder.Ident(funcName)
	parameters := strings.Join(funcVars, ",")

	ul.g.P("if ", fieldValue, ", err = ", funcGoIdent, "(", parameters, "); err != nil {")
	ul.g.P("    return err")
	ul.g.P("}")
}

func (ul *Unmarshal) loadMapKeyTypeInfo(field *protogen.Field, typeCodec *pbjson.TypeCodec,
) (typeName string, funcVars []string) {

	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.StringKind:
		typeName = "Str"
	case protoreflect.BoolKind:
		quote := loadTypeCodecBool(typeCodec).Codec != pbjson.TypeBool_Bool
		typeName = "Bool"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int32Kind:
		quote := loadTypeCodecInt32(typeCodec).Codec != pbjson.TypeInt32_Number
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int64Kind:
		quote := loadTypeCodecInt64(typeCodec).Codec != pbjson.TypeInt64_Number
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint32Kind:
		quote := loadTypeCodecSInt32(typeCodec).Codec != pbjson.TypeSInt32_Number
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint64Kind:
		quote := loadTypeCodecSInt64(typeCodec).Codec != pbjson.TypeSInt64_Number
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed32Kind:
		quote := loadTypeCodecSFInt32(typeCodec).Codec != pbjson.TypeSFixed32_Number
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed64Kind:
		quote := loadTypeCodecSFInt64(typeCodec).Codec != pbjson.TypeSFixed64_Number
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint32Kind:
		quote := loadTypeCodecUint32(typeCodec).Codec != pbjson.TypeUint32_Number
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint64Kind:
		quote := loadTypeCodecUint64(typeCodec).Codec != pbjson.TypeUint64_Number
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed32Kind:
		quote := loadTypeCodecFixed32(typeCodec).Codec != pbjson.TypeFixed32_Number
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed64Kind:
		quote := loadTypeCodecFixed64(typeCodec).Codec != pbjson.TypeFixed64_Number
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	default:
		err := pkerror.New("unmarshal: unsupported kind of %s as map key", kind.String())
		panic(err)
	}
	return
}
func (ul *Unmarshal) loadValueTypeInfo(field *protogen.Field, typeCodec *pbjson.TypeCodec,
) (typeName string, funcVars []string) {

	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.Int32Kind:
		quote := loadTypeCodecInt32(typeCodec).Codec == pbjson.TypeInt32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Int64Kind:
		quote := loadTypeCodecInt64(typeCodec).Codec == pbjson.TypeInt64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint32Kind:
		quote := loadTypeCodecSInt32(typeCodec).Codec == pbjson.TypeSInt32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sint64Kind:
		quote := loadTypeCodecSInt64(typeCodec).Codec == pbjson.TypeSInt64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed32Kind:
		quote := loadTypeCodecSFInt32(typeCodec).Codec == pbjson.TypeSFixed32_String
		typeName = "I32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Sfixed64Kind:
		quote := loadTypeCodecSFInt64(typeCodec).Codec == pbjson.TypeSFixed64_String
		typeName = "I64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint32Kind:
		quote := loadTypeCodecUint32(typeCodec).Codec == pbjson.TypeUint32_String
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Uint64Kind:
		quote := loadTypeCodecUint64(typeCodec).Codec == pbjson.TypeUint64_String
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed32Kind:
		quote := loadTypeCodecFixed32(typeCodec).Codec == pbjson.TypeFixed32_String
		typeName = "U32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.Fixed64Kind:
		quote := loadTypeCodecFixed64(typeCodec).Codec == pbjson.TypeFixed64_String
		typeName = "U64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.FloatKind:
		quote := loadTypeCodecFloat(typeCodec).Codec == pbjson.TypeFloat_String
		typeName = "F32"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.DoubleKind:
		quote := loadTypeCodecDouble(typeCodec).Codec == pbjson.TypeDouble_String
		typeName = "F64"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.BoolKind:
		quote := loadTypeCodecBool(typeCodec).Codec == pbjson.TypeBool_String
		typeName = "Bool"
		funcVars = append(funcVars, strconv.FormatBool(quote))
	case protoreflect.StringKind:
		typeName = "Str"
	case protoreflect.BytesKind:
		typeName = "Bytes"
	case protoreflect.EnumKind:
		// FIXME: 需要补充指针型变量的测试.
		typeSet := loadTypeCodecEnum(typeCodec)
		switch typeSet.Codec {
		case pbjson.TypeEnum_Unset, pbjson.TypeEnum_Number:
			typeName = "EnumNum"
			funcVars = append(funcVars, strconv.FormatBool(false))
		case pbjson.TypeEnum_NumberString:
			typeName = "EnumNum"
			funcVars = append(funcVars, strconv.FormatBool(true))
		case pbjson.TypeEnum_String:
			_fieldGoType := fieldGoType(ul.g, field)
			enumValue := _fieldGoType + "_value"
			typeName = "EnumStr"
			funcVars = append(funcVars, enumValue)
		}
	case protoreflect.MessageKind:
		switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
		case pkwkt.Any:
			typeAny := loadTypeCodecAny(typeCodec)
			switch typeAny.Codec {
			case pbjson.TypeAny_Object, pbjson.TypeAny_Unset:
				typeName = "WKTAnyObject"
			case pbjson.TypeAny_Proto:
				typeName = "WKTAnyProto"
			}
		case pkwkt.Duration:
			typeDuration := loadTypeCodecDuration(typeCodec)
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
			typeTimestamp := loadTypeCodecTimestamp(typeCodec)
			switch typeTimestamp.Codec {
			case pbjson.TypeTimestamp_Object, pbjson.TypeTimestamp_Unset:
				typeName = "WKTTsObject"
			case pbjson.TypeTimestamp_TimeLayout:
				typeSet := loadTypeCodecTimestamp(typeCodec)
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

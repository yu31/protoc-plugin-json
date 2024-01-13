package gojson

import (
	"os"
	"strconv"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"github.com/yu31/protoc-kit-go/helper/pkwkt"
	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/cmd/internal/gojson/pkg/importpkg"
	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

func (p *Plugin) generateCodeMarshal(fields []*Field, bufLen int) {
	msg := p.message

	p.g.P("// MarshalJSON implements interface json.Marshaler for proto message ", msg.Desc.Name(), " in file ", p.file.Desc.Path())
	p.g.P("func (x *", msg.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
	p.g.P("    if x == nil {")
	p.g.P(`        return []byte("null"), nil`)
	p.g.P("    }")
	if len(fields) == 0 {
		p.g.P(`return []byte("{}"), nil`)
	} else {
		p.g.P("var err error")
		p.g.P("encoder := ", importpkg.PJEncoder.Ident("New"), "(", bufLen, ")")
		p.g.P("")
		p.g.P("// Add begin JSON identifier")
		p.g.P("encoder.AppendObjectBegin()")
		p.g.P("")
		for _, field := range fields {
			catch := pkerror.Recover("gojson", p.file, p.message, field.Field, func() {
				p.marshalForField(field)
			})
			if catch {
				os.Exit(1)
			}
		}
		p.g.P("")
		p.g.P("// Add end JSON identifier")
		p.g.P("encoder.AppendObjectEnd()")
		p.g.P("return encoder.Bytes(), err")
	}
	// End of MarshalJSON.
	p.g.P("}")
}

func (p *Plugin) marshalForField(field *Field) {
	switch {
	case field.Field.Desc.IsMap():
		p.marshalMap(field)
	case field.Field.Desc.IsList():
		p.marshalRepeated(field)
	case pkfield.FieldIsOneOf(field.Field):
		p.marshalOneOf(field)
	default:
		p.marshalPlain(field)
	}
}

func (p *Plugin) marshalMap(field *Field) {
	options := field.Options
	jsonKey := field.JSONKey
	goName := field.Field.GoName
	if options.Omitempty {
		p.g.P("if len(x.", goName, ") != 0 {")
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendMap(field)
		p.g.P("}")
	} else {
		p.marshalAppendObjectKey(jsonKey)
		p.g.P("if x.", goName, "!= nil {")
		p.marshalAppendMap(field)
		p.g.P("} else {")
		p.marshalAppendNULL()
		p.g.P("}")
	}
}
func (p *Plugin) marshalRepeated(field *Field) {
	options := field.Options
	jsonKey := field.JSONKey
	goName := field.Field.GoName

	if options.Omitempty {
		p.g.P("if len(x.", goName, ") != 0 {")
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendRepeated(field)
		p.g.P("}")
	} else {
		p.marshalAppendObjectKey(jsonKey)
		p.g.P("if x.", goName, "!= nil {")
		p.marshalAppendRepeated(field)
		p.g.P("} else {")
		p.marshalAppendNULL()
		p.g.P("}")
	}
}
func (p *Plugin) marshalOneOf(field *Field) {
	options := field.OneOf.Options
	jsonKey := field.JSONKey
	goName := field.Field.Oneof.GoName

	p.g.P("switch ov := x.", goName, ".(type) {")
	for _, oneField := range field.OneOf.Parts {
		oneType := p.g.QualifiedGoIdent(oneField.Field.GoIdent)
		p.g.P("case *", oneType, ":")
		if !options.Inline {
			p.marshalAppendObjectKey(jsonKey)
			p.g.P("encoder.AppendObjectBegin()")
			p.marshalPlain(oneField)
			p.g.P("encoder.AppendObjectEnd()")
		} else {
			p.marshalPlain(oneField)
		}
	}
	if !options.Inline && !options.Omitempty {
		// If the oneof field are inline, It will be omitted for empty value.
		p.g.P("case nil:")
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendNULL()
	}
	p.g.P("default:")
	p.g.P("  _ = ov // to avoids unused panics")
	p.g.P("}")
}
func (p *Plugin) marshalPlain(field *Field) {
	options := field.Options
	jsonKey := field.JSONKey
	if options.Omitempty {
		notEmptyCond := p.marshalGetNotEmptyCond(field.Field)
		p.g.P("if ", notEmptyCond, " {")
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendLiteral(field.Field, options)
		p.g.P("}")
	} else {
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendLiteral(field.Field, options)
	}
}

func (p *Plugin) marshalGetNotEmptyCond(field *protogen.Field) (cond string) {
	isOptional := pkfield.FieldIsOptional(field)
	isOnoOf := pkfield.FieldIsOneOf(field)

	var receiver string
	if isOnoOf {
		receiver = "ov." + field.GoName
	} else {
		receiver = "x." + field.GoName
	}

	if isOptional {
		cond = receiver + " != nil "
		return
	}

	switch field.Desc.Kind() {
	case protoreflect.DoubleKind, protoreflect.FloatKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		cond = receiver + " != 0 "
	case protoreflect.BoolKind:
		cond = receiver
	case protoreflect.StringKind:
		cond = receiver + ` != "" `
	case protoreflect.BytesKind:
		cond = "len(" + receiver + ")" + " != 0 "
	case protoreflect.MessageKind:
		cond = receiver + " != nil "
	case protoreflect.EnumKind:
		cond = receiver + " != 0 "
	default:
		err := pkerror.New("marshal: unsupported kind of %s", field.Desc.Kind().String())
		panic(err)
	}
	return
}

func (p *Plugin) marshalAppendObjectKey(key string) {
	p.g.P("encoder.AppendObjectKey(", strconv.Quote(key), ")")
}
func (p *Plugin) marshalAppendNULL() {
	p.g.P("encoder.AppendLiteralNULL()")
}
func (p *Plugin) marshalAppendMap(field *Field) {
	p.g.P("encoder.AppendObjectBegin()")
	p.g.P("for mk, mv := range ", "x.", field.Field.GoName, " {") // start for loop
	p.marshalAppendMapKey(field.Field, field.Options)
	p.marshalAppendLiteral(field.Field, field.Options)
	p.g.P("}") // end for loop.
	p.g.P("encoder.AppendObjectEnd()")
}
func (p *Plugin) marshalAppendRepeated(field *Field) {
	p.g.P("encoder.AppendArrayBegin()")
	p.g.P("for _, ri := range ", "x.", field.Field.GoName, " {")
	p.marshalAppendLiteral(field.Field, field.Options)
	p.g.P("}")
	p.g.P("encoder.AppendArrayEnd()")
}
func (p *Plugin) marshalAppendMapKey(field *protogen.Field, options *pbjson.FieldOptions) {
	typeSet := loadMapOptions(field, options).Key

	kind := field.Desc.MapKey().Kind()
	switch kind {
	case protoreflect.StringKind:
		p.g.P("encoder.AppendMapKeyString(mk)")
	case protoreflect.Int32Kind:
		quote := "true"
		typeInt32 := loadTypeCodecInt32(typeSet)
		if typeInt32.Codec == pbjson.TypeInt32_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyInt32(mk, ", quote, ")")
	case protoreflect.Int64Kind:
		quote := "true"
		typeInt64 := loadTypeCodecInt64(typeSet)
		if typeInt64.Codec == pbjson.TypeInt64_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyInt64(mk, ", quote, ")")
	case protoreflect.Sint32Kind:
		quote := "true"
		typeSInt32 := loadTypeCodecSInt32(typeSet)
		if typeSInt32.Codec == pbjson.TypeSInt32_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyInt32(mk, ", quote, ")")
	case protoreflect.Sint64Kind:
		quote := "true"
		typeSInt64 := loadTypeCodecSInt64(typeSet)
		if typeSInt64.Codec == pbjson.TypeSInt64_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyInt64(mk, ", quote, ")")
	case protoreflect.Sfixed32Kind:
		quote := "true"
		typeSFixed32 := loadTypeCodecSFInt32(typeSet)
		if typeSFixed32.Codec == pbjson.TypeSFixed32_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyInt32(mk, ", quote, ")")
	case protoreflect.Sfixed64Kind:
		quote := "true"
		typeSFixed64 := loadTypeCodecSFInt64(typeSet)
		if typeSFixed64.Codec == pbjson.TypeSFixed64_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyInt64(mk, ", quote, ")")
	case protoreflect.Uint32Kind:
		quote := "true"
		typeUint32 := loadTypeCodecUint32(typeSet)
		if typeUint32.Codec == pbjson.TypeUint32_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyUInt32(mk, ", quote, ")")
	case protoreflect.Uint64Kind:
		quote := "true"
		typeUint64 := loadTypeCodecUint64(typeSet)
		if typeUint64.Codec == pbjson.TypeUint64_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyUInt64(mk, ", quote, ")")
	case protoreflect.Fixed32Kind:
		quote := "true"
		typeFixed32 := loadTypeCodecFixed32(typeSet)
		if typeFixed32.Codec == pbjson.TypeFixed32_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyUInt32(mk, ", quote, ")")
	case protoreflect.Fixed64Kind:
		quote := "true"
		typeFixed64 := loadTypeCodecFixed64(typeSet)
		if typeFixed64.Codec == pbjson.TypeFixed64_Number {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyUInt64(mk, ", quote, ")")
	case protoreflect.BoolKind:
		quote := "true"
		typeBool := loadTypeCodecBool(typeSet)
		if typeBool.Codec == pbjson.TypeBool_Bool {
			quote = "false"
		}
		p.g.P("encoder.AppendMapKeyBool(mk, ", quote, ")")
	default:
		err := pkerror.New("marshal: unsupported kind of %s as map key", kind.String())
		panic(err)
	}
}

func (p *Plugin) marshalAppendLiteral(field *protogen.Field, options *pbjson.FieldOptions) {
	var receiver string
	switch {
	case field.Desc.IsMap():
		receiver = "mv" // as map value.
	case field.Desc.IsList():
		receiver = "ri" // as repeated item.
	case pkfield.FieldIsOneOf(field):
		receiver = "ov." + field.GoName // filed in `oneof` parts
	default:
		receiver = "x." + field.GoName
	}

	var typeCodec *pbjson.TypeCodec
	switch {
	case field.Desc.IsMap():
		typeCodec = loadMapOptions(field, options).Value
	case field.Desc.IsList():
		typeCodec = loadRepeatedOptions(field, options).Elem
	default:
		typeCodec = loadPlainOptions(field, options).Value
	}

	isOptional := pkfield.FieldIsOptional(field)

	if field.Desc.IsMap() {
		// Only used by append value if the field is map.
		field = field.Message.Fields[1]
	}

	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.StringKind:
		if isOptional {
			p.g.P("encoder.AppendPointerString(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralString(", receiver, ")")
		}
	case protoreflect.BytesKind:
		p.g.P("encoder.AppendLiteralBytes(", receiver, ")")
	case protoreflect.Int32Kind:
		quote := "false"
		typeInt32 := loadTypeCodecInt32(typeCodec)
		if typeInt32.Codec == pbjson.TypeInt32_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerInt32(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt32(", receiver, ", ", quote, ")")
		}
	case protoreflect.Int64Kind:
		quote := "false"
		typeInt64 := loadTypeCodecInt64(typeCodec)
		if typeInt64.Codec == pbjson.TypeInt64_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerInt64(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt64(", receiver, ", ", quote, ")")
		}
	case protoreflect.Sint32Kind:
		quote := "false"
		typeSInt32 := loadTypeCodecSInt32(typeCodec)
		if typeSInt32.Codec == pbjson.TypeSInt32_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerInt32(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt32(", receiver, ", ", quote, ")")
		}
	case protoreflect.Sint64Kind:
		quote := "false"
		typeSInt64 := loadTypeCodecSInt64(typeCodec)
		if typeSInt64.Codec == pbjson.TypeSInt64_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerInt64(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt64(", receiver, ", ", quote, ")")
		}
	case protoreflect.Sfixed32Kind:
		quote := "false"
		typeSFixed32 := loadTypeCodecSFInt32(typeCodec)
		if typeSFixed32.Codec == pbjson.TypeSFixed32_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerInt32(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt32(", receiver, ", ", quote, ")")
		}
	case protoreflect.Sfixed64Kind:
		quote := "false"
		typeSFixed64 := loadTypeCodecSFInt64(typeCodec)
		if typeSFixed64.Codec == pbjson.TypeSFixed64_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerInt64(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt64(", receiver, ", ", quote, ")")
		}
	case protoreflect.Uint32Kind:
		quote := "false"
		typeUint32 := loadTypeCodecUint32(typeCodec)
		if typeUint32.Codec == pbjson.TypeUint32_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerUint32(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint32(", receiver, ", ", quote, ")")
		}
	case protoreflect.Uint64Kind:
		quote := "false"
		typeUint64 := loadTypeCodecUint64(typeCodec)
		if typeUint64.Codec == pbjson.TypeUint64_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerUint64(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint64(", receiver, ", ", quote, ")")
		}
	case protoreflect.Fixed32Kind:
		quote := "false"
		typeFixed32 := loadTypeCodecFixed32(typeCodec)
		if typeFixed32.Codec == pbjson.TypeFixed32_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerUint32(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint32(", receiver, ", ", quote, ")")
		}
	case protoreflect.Fixed64Kind:
		quote := "false"
		typeFixed64 := loadTypeCodecFixed64(typeCodec)
		if typeFixed64.Codec == pbjson.TypeFixed64_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerUint64(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint64(", receiver, ", ", quote, ")")
		}
	case protoreflect.FloatKind:
		quote := "false"
		typeFloat := loadTypeCodecFloat(typeCodec)
		if typeFloat.Codec == pbjson.TypeFloat_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerFloat32(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralFloat32(", receiver, ", ", quote, ")")
		}
	case protoreflect.DoubleKind:
		quote := "false"
		typeDouble := loadTypeCodecDouble(typeCodec)
		if typeDouble.Codec == pbjson.TypeDouble_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerFloat64(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ", ", quote, ")")
		}
	case protoreflect.BoolKind:
		quote := "false"
		typeBool := loadTypeCodecBool(typeCodec)
		if typeBool.Codec == pbjson.TypeBool_String {
			quote = "true"
		}
		if isOptional {
			p.g.P("encoder.AppendPointerBool(", receiver, ", ", quote, ")")
		} else {
			p.g.P("encoder.AppendLiteralBool(", receiver, ", ", quote, ")")
		}
	case protoreflect.MessageKind:
		p.marshalAppendMessage(field, typeCodec, receiver)
	case protoreflect.EnumKind:
		if isOptional {
			p.g.P("if ", receiver, "!= nil {")
		}

		typeEnum := loadTypeCodecEnum(typeCodec)
		switch typeEnum.Codec {
		case pbjson.TypeEnum_Unset, pbjson.TypeEnum_Number:
			p.g.P("encoder.AppendLiteralInt32(int32(", receiver, ".Number()", "), false)")
		case pbjson.TypeEnum_NumberString:
			p.g.P("encoder.AppendLiteralInt32(int32(", receiver, ".Number()", "), true)")
		case pbjson.TypeEnum_String:
			p.g.P("encoder.AppendLiteralString(", receiver, ".String()", ")")
		}

		if isOptional {
			p.g.P("} else {")
			p.marshalAppendNULL()
			p.g.P("}")
		}
	default:
		err := pkerror.New("marshal: unsupported kind of %s as value", kind.String())
		panic(err)
	}
}

func (p *Plugin) marshalAppendMessage(field *protogen.Field, typeCodec *pbjson.TypeCodec, receiver string) {
	// Supported Well Know Type.
	switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
	case pkwkt.Any:
		typeAny := loadTypeCodecAny(typeCodec)
		switch typeAny.Codec {
		case pbjson.TypeAny_Object:
		case pbjson.TypeAny_Proto:
			p.g.P("if err = encoder.AppendWKTAnyByProto(", receiver, "); err != nil {")
			p.g.P("    return nil, err")
			p.g.P("}")
			return
		}
	case pkwkt.Duration:
		typeDuration := loadTypeCodecDuration(typeCodec)
		switch typeDuration.Codec {
		case pbjson.TypeDuration_Object:
		case pbjson.TypeDuration_String:
			p.g.P("encoder.AppendLiteralString(", receiver, ".AsDuration().String())")
			return
		case pbjson.TypeDuration_Nanosecond:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Nanoseconds(), false)")
			return
		case pbjson.TypeDuration_NanosecondString:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Nanoseconds(), true)")
			return
		case pbjson.TypeDuration_Microsecond:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Microseconds(), false)")
			return
		case pbjson.TypeDuration_MicrosecondString:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Microseconds(), true)")
			return
		case pbjson.TypeDuration_Millisecond:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Milliseconds(), false)")
			return
		case pbjson.TypeDuration_MillisecondString:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Milliseconds(), true)")
			return
		case pbjson.TypeDuration_Second:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Seconds(), false)")
			return
		case pbjson.TypeDuration_SecondString:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Seconds(), true)")
			return
		case pbjson.TypeDuration_Minute:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Minutes(), false)")
			return
		case pbjson.TypeDuration_MinuteString:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Minutes(), true)")
			return
		case pbjson.TypeDuration_Hour:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Hours(), false)")
			return
		case pbjson.TypeDuration_HourString:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Hours(), true)")
			return
		}
	case pkwkt.Timestamp:
		typeTimestamp := loadTypeCodecTimestamp(typeCodec)
		switch typeTimestamp.Codec {
		case pbjson.TypeTimestamp_Object:
		case pbjson.TypeTimestamp_TimeLayout:
			layout := typeTimestamp.Layout.Golang
			p.g.P("encoder.AppendLiteralString(", receiver, ".AsTime().Format(", strconv.Quote(layout), "))")
			return
		case pbjson.TypeTimestamp_UnixNano:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixNano(), false)")
			return
		case pbjson.TypeTimestamp_UnixNanoString:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixNano(), true)")
			return
		case pbjson.TypeTimestamp_UnixMicro:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixMicro(), false)")
			return
		case pbjson.TypeTimestamp_UnixMicroString:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixMicro(), true)")
			return
		case pbjson.TypeTimestamp_UnixMilli:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixMilli(), false)")
			return
		case pbjson.TypeTimestamp_UnixMilliString:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixMilli(), true)")
			return
		case pbjson.TypeTimestamp_UnixSec:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().Unix(), false)")
			return
		case pbjson.TypeTimestamp_UnixSecString:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().Unix(), true)")
			return
		}
	default:
	}

	p.g.P("if err = encoder.AppendLiteralInterface(", receiver, "); err != nil {")
	p.g.P("    return nil, err")
	p.g.P("}")
}

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

func (p *Plugin) generateCodeUnmarshal(fields []*Field, variables []string) {
	msg := p.message

	p.g.P("// UnmarshalJSON implements json.Unmarshaler for proto message ", msg.Desc.Name(), " in file ", p.file.Desc.Path())
	p.g.P("func (x *", msg.GoIdent.GoName, ") UnmarshalJSON(b []byte) error {")
	p.g.P("    if x == nil {")
	p.g.P("        return ", importpkg.Errors.Ident("New"), `("json: Unmarshal: `, string(msg.GoIdent.GoImportPath), ".(*", msg.GoIdent.GoName, `) is nil")`)
	p.g.P("    }")
	if len(fields) != 0 {
		if len(variables) != 0 {
			// Generated variables to flag the oneof field are loaded.
			p.g.P("var (")
			for _, varIsFill := range variables {
				p.g.P(varIsFill, " bool")
			}
			p.g.P(")")
		}
		// Generated loop scan code.
		p.unmarshalLoopScan(fields)
	}
	p.g.P("    return nil")
	p.g.P("}")
}
func (p *Plugin) unmarshalLoopScan(fields []*Field) {
	// pre-check
	p.g.P("var (")
	p.g.P("    err error")
	p.g.P("    isNULL bool")
	p.g.P("    decoder *", importpkg.PJDecoder.Ident("Decoder"))
	p.g.P(")")
	p.g.P("if decoder, err = ", importpkg.PJDecoder.Ident("New"), "(b); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL, err = decoder.BeforeReadJSON(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL {")
	p.g.P("    return nil")
	p.g.P("}")

	p.g.P("LOOP_SCAN:")
	p.g.P("for { // Loop to scan object.")
	p.unmarshalLoopRead(fields)
	p.g.P("}")
}

func (p *Plugin) unmarshalLoopRead(fields []*Field) {
	p.g.P("var (")
	p.g.P("    jsonKey string")
	p.g.P("    isEnd bool")
	p.g.P(")")
	p.g.P("if isEnd, err = decoder.BeforeScanNext(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isEnd {")
	p.g.P("    break LOOP_SCAN")
	p.g.P("}")
	p.g.P("if jsonKey, err = decoder.ReadJSONKey(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("switch jsonKey { // match the jsonKey")
	for _, field := range fields {
		catch := pkerror.Recover("gojson", p.file, p.message, field.Field, func() {
			p.unmarshalForField(field)
		})
		if catch {
			os.Exit(1)
		}
	}
	p.g.P("default:")
	if p.msgOptions.DisallowUnknownFields {
		// FIXME: optimized error.
		p.g.P("return ", importpkg.FMT.Ident("Errorf"), `("json: unknown field %q", jsonKey)`)
	} else {
		p.g.P("if err = decoder.DiscardValue(jsonKey); err != nil {")
		p.g.P("    return err")
		p.g.P("}")
	}
	p.g.P("} // end switch")
}

func (p *Plugin) unmarshalForField(field *Field) {
	if pkfield.FieldIsOneOf(field.Field) {
		if options := field.OneOf.Options; options.Hide {
			p.unmarshalOneOfField(field)
			return
		}
	}
	p.g.P("case ", strconv.Quote(field.JSONKey), ":")
	switch {
	case field.Field.Desc.IsMap():
		p.unmarshalMap(field)
	case field.Field.Desc.IsList():
		p.unmarshalRepeated(field)
	case pkfield.FieldIsOneOf(field.Field):
		p.unmarshalOneOf(field)
	default:
		p.unmarshalPlain(field)
	}
}

func (p *Plugin) unmarshalMap(field *Field) {
	goName := field.Field.GoName
	goType := pkfield.FieldGoType(p.g, field.Field)

	// Read and parse map.
	p.g.P("if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL {")
	p.g.P("    x.", goName, " = nil")
	p.g.P("    continue LOOP_SCAN")
	p.g.P("}")
	p.g.P("if x.", goName, " == nil {")
	p.g.P("    x.", goName, " = ", "make(", goType, ")")
	p.g.P("}")
	p.g.P("for {") // start loop.
	p.g.P("	if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {")
	p.g.P("    	return err ")
	p.g.P("	}")
	p.g.P("	if isEnd {")
	p.g.P("    	break")
	p.g.P("	}")
	p.unmarshalReadMapKey(field.Field)
	p.unmarshalReadLiteral(field.Field, field.Options)
	p.g.P("}") // end loop
}
func (p *Plugin) unmarshalRepeated(field *Field) {
	goName := field.Field.GoName
	goType := pkfield.FieldGoType(p.g, field.Field)

	// Read and parse array. same as the standard json.
	p.g.P("if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL {")
	p.g.P("    x.", goName, " = nil")
	p.g.P("    continue LOOP_SCAN")
	p.g.P("}")
	p.g.P("if x.", goName, " == nil {")
	p.g.P("    x.", goName, " = ", "make(", goType, ", 0)")
	p.g.P("}")
	p.g.P("i := 0")
	p.g.P("for {") // start loop.
	p.g.P("	if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {")
	p.g.P("   		return err")
	p.g.P("	}")
	p.g.P("	if isEnd {")
	p.g.P("   		break")
	p.g.P("	}")
	p.unmarshalReadLiteral(field.Field, field.Options)
	p.g.P("	i++")
	p.g.P("}") // end loop.
	p.g.P("// truncate the slice to consistent with standard library json.")
	p.g.P("x.", goName, " = x.", goName, "[:i]")
}
func (p *Plugin) unmarshalOneOf(field *Field) {
	goName := field.Field.Oneof.GoName
	// Read and parse oneof.
	p.g.P("if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL {")
	p.g.P("    x.", goName, " = nil")
	p.g.P("    continue LOOP_SCAN")
	p.g.P("}")
	p.g.P("for {")
	p.g.P("    var oneofKey string")
	p.g.P("	if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {")
	p.g.P("    	return err ")
	p.g.P("	}")
	p.g.P("	if isEnd {")
	p.g.P("    	break")
	p.g.P("	}")
	p.g.P("	if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {")
	p.g.P("    	return err")
	p.g.P("	}")
	p.g.P("    switch oneofKey { // match oneof key")
	p.unmarshalOneOfField(field)
	p.g.P("    default:")
	if p.msgOptions.DisallowUnknownFields {
		p.g.P("    return ", importpkg.FMT.Ident("Errorf"), `("json: unknown field %q in oneof parts", oneofKey)`)
	} else {
		p.g.P("if err = decoder.DiscardValue(jsonKey); err != nil {")
		p.g.P("    return err")
		p.g.P("}")
	}
	p.g.P("    } // end switch")
	p.g.P("}")
}
func (p *Plugin) unmarshalPlain(field *Field) {
	options := field.Options
	p.unmarshalReadLiteral(field.Field, options)
}
func (p *Plugin) unmarshalOneOfField(field *Field) {
	oneOfName := field.Field.Oneof.GoName
	oneOfIsFill := genVariableOneOfIsFill(oneOfName)

	for _, oneField := range field.OneOf.Parts {
		oneOfType := p.g.QualifiedGoIdent(oneField.Field.GoIdent)
		p.g.P("case ", strconv.Quote(oneField.JSONKey), ":")
		p.g.P("	if ", oneOfIsFill, " {")
		p.g.P("    	return ", importpkg.FMT.Ident("Errorf"), `("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)`)
		p.g.P("	}")
		p.g.P(oneOfIsFill, " = true")
		p.g.P("")
		p.g.P("	var ok bool")
		p.g.P("	var ot *", oneOfType)
		p.g.P("	if ot, ok = x.", oneOfName, ".(*", oneOfType, "); !ok {")
		p.g.P("    	ot = new(", oneOfType, ")")
		p.g.P("	}")
		p.unmarshalReadLiteral(oneField.Field, oneField.Options)
		p.g.P("	x.", oneOfName, " = ot")
	}
}

func (p *Plugin) unmarshalReadMapKey(field *protogen.Field) {
	field = field.Message.Fields[0]
	goType := fieldGoType(p.g, field)

	p.g.P("var mk ", goType)
	switch field.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("if mk, err = decoder.ReadMapKeyInt32(jsonKey); err != nil {")
		p.g.P("    return err ")
		p.g.P("}")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("if mk, err = decoder.ReadMapKeyInt64(jsonKey); err != nil {")
		p.g.P("    return err ")
		p.g.P("}")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("if mk, err = decoder.ReadMapKeyUint32(jsonKey); err != nil {")
		p.g.P("    return err ")
		p.g.P("}")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("if mk, err = decoder.ReadMapKeyUint64(jsonKey); err != nil {")
		p.g.P("    return err ")
		p.g.P("}")
	case protoreflect.StringKind:
		p.g.P("if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {")
		p.g.P("    return err ")
		p.g.P("}")
	case protoreflect.BoolKind:
		p.g.P("if mk, err = decoder.ReadMapKeyBool(jsonKey); err != nil {")
		p.g.P("    return err ")
		p.g.P("}")
	default:
		err := pkerror.New("unmarshal: unsupported kind of %s as map key", field.Desc.Kind().String())
		panic(err)
	}
}

func (p *Plugin) unmarshalReadLiteral(field *protogen.Field, options *pbjson.FieldOptions) {
	goName := field.GoName
	isMap := field.Desc.IsMap()
	isList := field.Desc.IsList()
	isOneOf := pkfield.FieldIsOneOf(field)
	isOptional := pkfield.FieldIsOptional(field)

	if isMap {
		field = field.Message.Fields[1]
	}

	kind := field.Desc.Kind()
	goType := fieldGoType(p.g, field)
	if isOptional || kind == protoreflect.MessageKind {
		p.g.P("var vv *", goType)
	} else {
		p.g.P("var vv ", goType)
	}

	var receiver string
	switch {
	case isMap:
		receiver = "x." + goName + "[mk]"
	case isList:
		receiver = "x." + goName + "[i]"
		p.g.P("if i >= len(x.", goName, ") {")
		p.g.P("    x.", goName, " = append(", "x.", goName, ", vv", ")")
		p.g.P("}")
	case isOneOf:
		receiver = "ot." + goName
	default:
		receiver = "x." + goName
	}

	switch kind {
	case protoreflect.StringKind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerString", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralString", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerInt32", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralInt32", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerInt64", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralInt64", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerUint32", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralUint32", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerUint64", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralUint64", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.FloatKind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerFloat32", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralFloat32", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.DoubleKind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerFloat64", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralFloat64", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.BoolKind:
		if isOptional {
			p.g.P("if vv, err = decoder.ReadPointerBool", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		} else {
			p.g.P("if vv, err = decoder.ReadLiteralBool", "(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
		}
	case protoreflect.BytesKind:
		p.g.P("if vv, err = decoder.ReadLiteralBytes", "(jsonKey); err != nil {")
		p.g.P("    return err")
		p.g.P("}")
	case protoreflect.MessageKind:
		code, checkNULL := p.unmarshalReadMessage(field, options)
		if checkNULL {
			p.g.P("if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {")
			p.g.P("    return err")
			p.g.P("}")
			p.g.P("if !isNULL {")
		}

		p.g.P("if ", receiver, " != nil {")
		p.g.P("	vv = ", receiver)
		p.g.P("} else {")
		p.g.P("    vv = new(", goType, ")")
		p.g.P("}")
		p.g.P("if err = ", code, "; err != nil {")
		p.g.P("    return err")
		p.g.P("}")

		if checkNULL {
			p.g.P("}")
		}
	case protoreflect.EnumKind:
		enumValue := goType + "_value"
		enumName := goType + "_name"
		typeEnum := loadTypeSetEnum(options)

		if isOptional {
			p.g.P("var v1 *int32")
			if typeEnum.Format == pbjson.TypeEnum_String {
				p.g.P("if v1, err = decoder.ReadPointerEnumString", "(jsonKey,", enumValue, "); err != nil {")
				p.g.P("    return err")
				p.g.P("}")
			} else {
				p.g.P("if v1, err = decoder.ReadPointerEnumNumber", "(jsonKey,", enumName, "); err != nil {")
				p.g.P("    return err")
				p.g.P("}")
			}
			p.g.P("if v1 != nil {")
			p.g.P("	v2 := ", goType, "(*v1)")
			p.g.P("	vv = &v2")
			p.g.P("}")
		} else {
			p.g.P("var v1 int32")
			if typeEnum.Format == pbjson.TypeEnum_String {
				p.g.P("if v1, err = decoder.ReadLiteralEnumString", "(jsonKey,", enumValue, "); err != nil {")
				p.g.P("    return err")
				p.g.P("}")
			} else {
				p.g.P("if v1, err = decoder.ReadLiteralEnumNumber", "(jsonKey,", enumName, "); err != nil {")
				p.g.P("    return err")
				p.g.P("}")
			}
			p.g.P("vv = ", goType, "(v1)")
		}
	default:
		err := pkerror.New("unmarshal: unsupported kind of %s as value", kind.String())
		panic(err)
	}
	// Store the value.
	p.g.P(receiver, " = vv")
}
func (p *Plugin) unmarshalReadMessage(field *protogen.Field, options *pbjson.FieldOptions,
) (code string, checkNULL bool) {
	// Supported Well Know Type.
	switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
	case pkwkt.Any:
		typeAny := loadTypeSetAny(options)
		switch typeAny.Format {
		case pbjson.TypeAny_Native:
		case pbjson.TypeAny_Proto:
			code = "decoder.ReadWKTAnyByProto(jsonKey, vv)"
			return code, true
		}
	case pkwkt.Duration:
		typeDuration := loadTypeSetDuration(options)
		switch typeDuration.Format {
		case pbjson.TypeDuration_Native:
		case pbjson.TypeDuration_String:
			code = "decoder.ReadWKTDurationByString(jsonKey, vv)"
			return code, false
		case pbjson.TypeDuration_Nanoseconds:
			code = "decoder.ReadWKTDurationByNanoseconds(jsonKey, vv)"
			return code, false
		case pbjson.TypeDuration_Microseconds:
			code = "decoder.ReadWKTDurationByMicroseconds(jsonKey, vv)"
			return code, false
		case pbjson.TypeDuration_Milliseconds:
			code = "decoder.ReadWKTDurationByMilliseconds(jsonKey, vv)"
			return code, false
		case pbjson.TypeDuration_Seconds:
			code = "decoder.ReadWKTDurationBySeconds(jsonKey, vv)"
			return code, false
		case pbjson.TypeDuration_Minutes:
			code = "decoder.ReadWKTDurationByMinutes(jsonKey, vv)"
			return code, false
		case pbjson.TypeDuration_Hours:
			code = "decoder.ReadWKTDurationByHours(jsonKey, vv)"
			return code, false
		}
	case pkwkt.Timestamp:
		typeTimestamp := loadTypeSetTimestamp(options)
		switch typeTimestamp.Format {
		case pbjson.TypeTimestamp_Native:
		case pbjson.TypeTimestamp_TimeLayout:
			layout := typeTimestamp.Layout.Golang
			code = "decoder.ReadWKTTimestampByString(jsonKey, vv, " + strconv.Quote(layout) + ")"
			return code, false
		case pbjson.TypeTimestamp_UnixNano:
			code = "decoder.ReadWKTTimestampByUnixNano(jsonKey, vv)"
			return code, false
		case pbjson.TypeTimestamp_UnixMicro:
			code = "decoder.ReadWKTTimestampByUnixMicro(jsonKey, vv)"
			return code, false
		case pbjson.TypeTimestamp_UnixMilli:
			code = "decoder.ReadWKTTimestampByUnixMilli(jsonKey, vv)"
			return code, false
		case pbjson.TypeTimestamp_UnixSec:
			code = "decoder.ReadWKTTimestampByUnixSec(jsonKey, vv)"
			return code, false
		}
	default:
	}

	// The default code.
	code = "decoder.ReadLiteralInterface(jsonKey, vv)"
	return code, true
}

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
		if !options.Hide {
			p.marshalAppendObjectKey(jsonKey)
			p.g.P("encoder.AppendObjectBegin()")
			p.marshalPlain(oneField)
			p.g.P("encoder.AppendObjectEnd()")
		} else {
			p.marshalPlain(oneField)
		}
	}
	if !options.Hide && !options.Omitempty {
		// If the oneof field are hide, It will be omitted for empty value.
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
	//case protoreflect.BoolKind:
	//	cond = receiver
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
	p.marshalAppendMapKey(field.Field)
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
func (p *Plugin) marshalAppendMapKey(field *protogen.Field) {
	kind := field.Desc.MapKey().Kind()
	switch kind {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("encoder.AppendMapKeyInt32(mk)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("encoder.AppendMapKeyInt64(mk)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("encoder.AppendMapKeyUInt32(mk)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("encoder.AppendMapKeyUInt64(mk)")
	case protoreflect.StringKind:
		p.g.P("encoder.AppendMapKeyString(mk)")
	case protoreflect.BoolKind:
		p.g.P("encoder.AppendMapKeyBool(mk)")
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
	case protoreflect.BoolKind:
		if isOptional {
			p.g.P("encoder.AppendPointerBool(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralBool(", receiver, ")")
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerInt32(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt32(", receiver, ")")
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerInt64(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt64(", receiver, ")")
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerUint32(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint32(", receiver, ")")
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerUint64(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint64(", receiver, ")")
		}
	case protoreflect.FloatKind:
		if isOptional {
			p.g.P("encoder.AppendPointerFloat32(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralFloat32(", receiver, ")")
		}
	case protoreflect.DoubleKind:
		if isOptional {
			p.g.P("encoder.AppendPointerFloat64(", receiver, ")")
		} else {
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ")")
		}
	case protoreflect.MessageKind:
		p.marshalAppendMessage(field, options, receiver)
	case protoreflect.EnumKind:
		if isOptional {
			p.g.P("if ", receiver, "!= nil {")
		}

		typeEnum := loadTypeSetEnum(options)
		if typeEnum.Format == pbjson.TypeEnum_String {
			p.g.P("encoder.AppendLiteralString(", receiver, ".String()", ")")
		} else {
			p.g.P("encoder.AppendLiteralInt32(int32(", receiver, ".Number()", "))")
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

func (p *Plugin) marshalAppendMessage(field *protogen.Field, options *pbjson.FieldOptions, receiver string) {
	// Supported Well Know Type.
	switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
	case pkwkt.Any:
		typeAny := loadTypeSetAny(options)
		switch typeAny.Format {
		case pbjson.TypeAny_Native:
		case pbjson.TypeAny_Proto:
			p.g.P("if err = encoder.AppendWKTAnyByProto(", receiver, "); err != nil {")
			p.g.P("    return nil, err")
			p.g.P("}")
			return
		}
	case pkwkt.Duration:
		typeDuration := loadTypeSetDuration(options)
		switch typeDuration.Format {
		case pbjson.TypeDuration_Native:
		case pbjson.TypeDuration_String:
			p.g.P("encoder.AppendLiteralString(", receiver, ".AsDuration().String())")
			return
		case pbjson.TypeDuration_Nanoseconds:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Nanoseconds())")
			return
		case pbjson.TypeDuration_Microseconds:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Microseconds())")
			return
		case pbjson.TypeDuration_Milliseconds:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsDuration().Milliseconds())")
			return
		case pbjson.TypeDuration_Seconds:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Seconds())")
			return
		case pbjson.TypeDuration_Minutes:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Minutes())")
			return
		case pbjson.TypeDuration_Hours:
			p.g.P("encoder.AppendLiteralFloat64(", receiver, ".AsDuration().Hours())")
			return
		}
	case pkwkt.Timestamp:
		typeTimestamp := loadTypeSetTimestamp(options)
		switch typeTimestamp.Format {
		case pbjson.TypeTimestamp_Native:
		case pbjson.TypeTimestamp_TimeLayout:
			layout := typeTimestamp.Layout.Golang
			p.g.P("encoder.AppendLiteralString(", receiver, ".AsTime().Format(", strconv.Quote(layout), "))")
			return
		case pbjson.TypeTimestamp_UnixNano:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixNano())")
			return
		case pbjson.TypeTimestamp_UnixMicro:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixMicro())")
			return
		case pbjson.TypeTimestamp_UnixMilli:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().UnixMilli())")
			return
		case pbjson.TypeTimestamp_UnixSec:
			p.g.P("encoder.AppendLiteralInt64(", receiver, ".AsTime().Unix())")
			return
		}
	default:
	}

	p.g.P("if err = encoder.AppendLiteralInterface(", receiver, "); err != nil {")
	p.g.P("    return nil, err")
	p.g.P("}")
}

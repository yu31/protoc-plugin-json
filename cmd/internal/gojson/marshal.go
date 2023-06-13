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

func (p *Plugin) generateCodeForMarshal(fields []*protogen.Field) {
	msg := p.message

	p.g.P("// MarshalJSON implements interface json.Marshaler for proto message ", msg.Desc.Name(), " in file ", p.file.Desc.Path())
	p.g.P("func (x *", msg.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
	p.g.P("    if x == nil {")
	p.g.P(`        return []byte("null"), nil`)
	p.g.P("    }")
	if len(fields) == 0 {
		p.g.P(`return []byte("{}"), nil`)
	} else {
		bufLen := p.guessBufLength(fields)
		p.g.P("var err error")
		p.g.P("encoder := ", importpkg.PJEncoder.Ident("New"), "(", bufLen, ")")
		p.g.P("")
		p.g.P("// Add begin JSON identifier")
		p.g.P("encoder.AppendObjectBegin()")
		p.g.P("")
		for _, field := range fields {
			catch := pkerror.Recover("gojson", p.file, p.message, field, func() {
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

func (p *Plugin) marshalForField(field *protogen.Field) {
	switch {
	case field.Desc.IsMap():
		p.marshalForMap(field)
	case field.Desc.IsList():
		p.marshalForRepeated(field)
	case pkfield.FieldIsOneOf(field):
		p.marshalForOneOf(field.Oneof)
	default:
		p.marshalForPlain(field)
	}
}

func (p *Plugin) marshalForOneOf(oneof *protogen.Oneof) {
	options := p.loadOneOfOptions(oneof)
	if options.Ignore {
		return
	}
	jsonKey := p.getJSONKeyForOneof(options, oneof)

	p.g.P("switch ov := x.", oneof.GoName, ".(type) {")
	for _, oneField := range oneof.Fields {
		p.g.P("case *", p.g.QualifiedGoIdent(oneField.GoIdent), ":")
		if !options.Hide {
			p.marshalAppendObjectKey(jsonKey)
			p.g.P("encoder.AppendObjectBegin()")
			p.marshalForPlain(oneField)
			p.g.P("encoder.AppendObjectEnd()")
		} else {
			p.marshalForPlain(oneField)
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
func (p *Plugin) marshalForPlain(field *protogen.Field) {
	options := p.loadFieldOptions(field)
	if options.Ignore {
		return
	}
	jsonKey := p.getJSONKeyForField(options, field)
	if options.Omitempty {
		notEmptyCond := p.marshalGenNotEmptyCond(field)
		p.g.P("if ", notEmptyCond, " {")
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendValue(field)
		p.g.P("}")
	} else {
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendValue(field)
	}
}
func (p *Plugin) marshalForMap(field *protogen.Field) {
	options := p.loadFieldOptions(field)
	if options.Ignore {
		return
	}
	jsonKey := p.getJSONKeyForField(options, field)
	if options.Omitempty {
		p.g.P("if len(x.", field.GoName, ") != 0 {")
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendMap(field)
		p.g.P("}")
	} else {
		p.marshalAppendObjectKey(jsonKey)
		p.g.P("if x.", field.GoName, "!= nil {")
		p.marshalAppendMap(field)
		p.g.P("} else {")
		p.marshalAppendNULL()
		p.g.P("}")
	}
}
func (p *Plugin) marshalForRepeated(field *protogen.Field) {
	options := p.loadFieldOptions(field)
	if options.Ignore {
		return
	}
	jsonKey := p.getJSONKeyForField(options, field)
	if options.Omitempty {
		p.g.P("if len(x.", field.GoName, ") != 0 {")
		p.marshalAppendObjectKey(jsonKey)
		p.marshalAppendRepeated(field)
		p.g.P("}")
	} else {
		p.marshalAppendObjectKey(jsonKey)
		p.g.P("if x.", field.GoName, "!= nil {")
		p.marshalAppendRepeated(field)
		p.g.P("} else {")
		p.marshalAppendNULL()
		p.g.P("}")
	}
}

func (p *Plugin) marshalGenNotEmptyCond(field *protogen.Field) (cond string) {
	isOptional := pkfield.FieldIsOptional(field)
	isOnoOf := pkfield.FieldIsOneOf(field)

	var varName string
	if isOnoOf {
		varName = "ov." + field.GoName
	} else {
		varName = "x." + field.GoName
	}

	if isOptional {
		cond = varName + " != nil "
		return
	}

	switch field.Desc.Kind() {
	case protoreflect.DoubleKind, protoreflect.FloatKind,
		protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		cond = varName + " != 0 "
	//case protoreflect.BoolKind:
	//	cond = varName
	case protoreflect.StringKind:
		cond = varName + ` != "" `
	case protoreflect.BytesKind:
		cond = "len(" + varName + ")" + " != 0 "
	case protoreflect.MessageKind:
		cond = varName + " != nil "
	case protoreflect.EnumKind:
		cond = varName + " != 0 "
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
func (p *Plugin) marshalAppendMap(field *protogen.Field) {
	p.g.P("encoder.AppendObjectBegin()")
	p.g.P("for mk, mv := range ", "x.", field.GoName, " {") // start for loop
	switch field.Desc.MapKey().Kind() {
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
	//case protoreflect.BoolKind:
	//	p.g.P("encoder.AppendMapKeyBool(mk)")
	default:
		err := pkerror.New("marshal: unsupported type of map key: %s", field.Desc.MapKey().Kind().String())
		panic(err)
	}
	p.marshalAppendValue(field)
	p.g.P("}") // end for loop.
	p.g.P("encoder.AppendObjectEnd()")
}
func (p *Plugin) marshalAppendRepeated(field *protogen.Field) {
	p.g.P("encoder.AppendArrayBegin()")
	p.g.P("for _, ri := range ", "x.", field.GoName, " {")
	p.marshalAppendValue(field)
	p.g.P("}")
	p.g.P("encoder.AppendArrayEnd()")
}
func (p *Plugin) marshalAppendValue(field *protogen.Field) {
	options := p.loadFieldOptions(field)

	var varName string
	switch {
	case field.Desc.IsMap():
		varName = "mv" // as map value.
	case field.Desc.IsList():
		varName = "ri" // as repeated item.
	case pkfield.FieldIsOneOf(field):
		varName = "ov." + field.GoName // filed in `oneof` parts
	default:
		varName = "x." + field.GoName
	}

	isOptional := pkfield.FieldIsOptional(field)

	if field.Desc.IsMap() {
		// Only used by append value if the field is map.
		field = field.Message.Fields[1]
	}

	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		if isOptional {
			p.g.P("encoder.AppendPointerString(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralString(", varName, ")")
		}
	case protoreflect.BytesKind:
		p.g.P("encoder.AppendLiteralBytes(", varName, ")")
	case protoreflect.BoolKind:
		if isOptional {
			p.g.P("encoder.AppendPointerBool(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralBool(", varName, ")")
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerInt32(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt32(", varName, ")")
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerInt64(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralInt64(", varName, ")")
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerUint32(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint32(", varName, ")")
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerUint64(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralUint64(", varName, ")")
		}
	case protoreflect.FloatKind:
		if isOptional {
			p.g.P("encoder.AppendPointerFloat32(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralFloat32(", varName, ")")
		}
	case protoreflect.DoubleKind:
		if isOptional {
			p.g.P("encoder.AppendPointerFloat64(", varName, ")")
		} else {
			p.g.P("encoder.AppendLiteralFloat64(", varName, ")")
		}
	case protoreflect.MessageKind:
		p.marshalValueMessage(field, options, varName)
	case protoreflect.EnumKind:
		typeEnum := p.loadTypeSetEnum(options)

		if isOptional {
			p.g.P("if ", varName, "!= nil {")
		}

		if typeEnum.Format == pbjson.TypeEnum_String {
			p.g.P("encoder.AppendLiteralString(", varName, ".String()", ")")
		} else {
			p.g.P("encoder.AppendLiteralInt32(int32(", varName, ".Number()", "))")
		}

		if isOptional {
			p.g.P("} else {")
			p.marshalAppendNULL()
			p.g.P("}")
		}
	default:
		err := pkerror.New("marshal: unsupported kind of %s", field.Desc.Kind().String())
		panic(err)
	}
}

func (p *Plugin) marshalValueMessage(field *protogen.Field, options *pbjson.FieldOptions, varName string) {
	// Supported Well Know Type.
	switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
	case pkwkt.Any:
		typeAny := p.loadTypeSetAny(options)
		switch typeAny.Format {
		case pbjson.TypeAny_Native:
		case pbjson.TypeAny_Proto:
			p.g.P("if err = encoder.AppendWKTAnyByProto(", varName, "); err != nil {")
			p.g.P("    return nil, err")
			p.g.P("}")
			return
		}
	case pkwkt.Duration:
		typeDuration := p.loadTypeSetDuration(options)
		switch typeDuration.Format {
		case pbjson.TypeDuration_Native:
		case pbjson.TypeDuration_String:
			p.g.P("encoder.AppendLiteralString(", varName, ".AsDuration().String())")
			return
		case pbjson.TypeDuration_Nanoseconds:
			p.g.P("encoder.AppendLiteralInt64(", varName, ".AsDuration().Nanoseconds())")
			return
		case pbjson.TypeDuration_Microseconds:
			p.g.P("encoder.AppendLiteralInt64(", varName, ".AsDuration().Microseconds())")
			return
		case pbjson.TypeDuration_Milliseconds:
			p.g.P("encoder.AppendLiteralInt64(", varName, ".AsDuration().Milliseconds())")
			return
		case pbjson.TypeDuration_Seconds:
			p.g.P("encoder.AppendLiteralFloat64(", varName, ".AsDuration().Seconds())")
			return
		case pbjson.TypeDuration_Minutes:
			p.g.P("encoder.AppendLiteralFloat64(", varName, ".AsDuration().Minutes())")
			return
		case pbjson.TypeDuration_Hours:
			p.g.P("encoder.AppendLiteralFloat64(", varName, ".AsDuration().Hours())")
			return
		}
	case pkwkt.Timestamp:
		typeTimestamp := p.loadTypeSetTimestamp(options)
		switch typeTimestamp.Format {
		case pbjson.TypeTimestamp_Native:
		case pbjson.TypeTimestamp_TimeLayout:
			// TODO: check the layout.
			layout := typeTimestamp.Layout.Golang
			p.g.P("encoder.AppendLiteralString(", varName, ".AsTime().Format(", strconv.Quote(layout), "))")
			return
		case pbjson.TypeTimestamp_UnixNano:
			p.g.P("encoder.AppendLiteralInt64(", varName, ".AsTime().UnixNano())")
			return
		case pbjson.TypeTimestamp_UnixMicro:
			p.g.P("encoder.AppendLiteralInt64(", varName, ".AsTime().UnixMicro())")
			return
		case pbjson.TypeTimestamp_UnixMilli:
			p.g.P("encoder.AppendLiteralInt64(", varName, ".AsTime().UnixMilli())")
			return
		case pbjson.TypeTimestamp_UnixSec:
			p.g.P("encoder.AppendLiteralInt64(", varName, ".AsTime().Unix())")
			return
		}
	default:
	}

	p.g.P("if err = encoder.AppendLiteralInterface(", varName, "); err != nil {")
	p.g.P("    return nil, err")
	p.g.P("}")
}

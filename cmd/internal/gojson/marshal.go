package gojson

import (
	"os"
	"strconv"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
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
		p.marshalForOneof(field.Oneof)
	default:
		p.marshalForPlain(field)
	}
}

func (p *Plugin) marshalForOneof(oneof *protogen.Oneof) {
	options := p.loadOneofOptions(oneof)
	if options.Ignore {
		return
	}
	jsonKey := p.getJSONKeyForOneof(options, oneof)

	p.g.P("switch ov := x.", oneof.GoName, ".(type) {")
	for _, oneField := range oneof.Fields {
		p.g.P("case *", p.g.QualifiedGoIdent(oneField.GoIdent), ":")
		if !options.Hide {
			p.marshalAppendJSONKey(jsonKey)
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
		p.marshalAppendJSONKey(jsonKey)
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
		notEmptyCond := p.marshalGenNotEmptyCond(field, options)
		p.g.P("if ", notEmptyCond, " {")
		p.marshalAppendJSONKey(jsonKey)
		p.marshalAppendValue(field)
		p.g.P("}")
	} else {
		p.marshalAppendJSONKey(jsonKey)
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
		p.marshalAppendJSONKey(jsonKey)
		p.marshalAppendMap(field)
		p.g.P("}")
	} else {
		p.marshalAppendJSONKey(jsonKey)
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
		p.marshalAppendJSONKey(jsonKey)
		p.marshalAppendRepeated(field)
		p.g.P("}")
	} else {
		p.marshalAppendJSONKey(jsonKey)
		p.g.P("if x.", field.GoName, "!= nil {")
		p.marshalAppendRepeated(field)
		p.g.P("} else {")
		p.marshalAppendNULL()
		p.g.P("}")
	}
}

func (p *Plugin) marshalGenNotEmptyCond(field *protogen.Field, options *pbjson.FieldOptions) (cond string) {
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
		//// TODO: remove it.
		//if options.UseEnumString {
		//	// Can't omit empty value for field type is enum if you use enum string.
		//	options.Omitempty = false
		//	cond = ""
		//} else {
		//	cond = varName + " != 0 "
		//}
	default:
		err := pkerror.New("marshal: unsupported kind of %s", field.Desc.Kind().String())
		panic(err)
	}
	return
}

func (p *Plugin) marshalAppendJSONKey(key string) {
	p.g.P("encoder.AppendJSONKey(", strconv.Quote(key), ")")
}
func (p *Plugin) marshalAppendNULL() {
	p.g.P("encoder.AppendValueNULL()")
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
			p.g.P("encoder.AppendValueString(", varName, ")")
		}
	case protoreflect.BytesKind:
		p.g.P("encoder.AppendValueBytes(", varName, ")")
	case protoreflect.BoolKind:
		if isOptional {
			p.g.P("encoder.AppendPointerBool(", varName, ")")
		} else {
			p.g.P("encoder.AppendValueBool(", varName, ")")
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerInt32(", varName, ")")
		} else {
			p.g.P("encoder.AppendValueInt32(", varName, ")")
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerInt64(", varName, ")")
		} else {
			p.g.P("encoder.AppendValueInt64(", varName, ")")
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerUint32(", varName, ")")
		} else {
			p.g.P("encoder.AppendValueUint32(", varName, ")")
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if isOptional {
			p.g.P("encoder.AppendPointerUint64(", varName, ")")
		} else {
			p.g.P("encoder.AppendValueUint64(", varName, ")")
		}
	case protoreflect.FloatKind:
		if isOptional {
			p.g.P("encoder.AppendPointerFloat32(", varName, ")")
		} else {
			p.g.P("encoder.AppendValueFloat32(", varName, ")")
		}
	case protoreflect.DoubleKind:
		if isOptional {
			p.g.P("encoder.AppendPointerFloat64(", varName, ")")
		} else {
			p.g.P("encoder.AppendValueFloat64(", varName, ")")
		}
	case protoreflect.MessageKind:
		// TODO: supported Well Know Type.
		p.g.P("err = encoder.AppendValueInterface(", varName, ")")
		p.g.P("if err != nil {")
		p.g.P("    return nil, err")
		p.g.P("}")
	case protoreflect.EnumKind:
		if isOptional {
			p.g.P("if ", varName, "!= nil {")
		}
		if options.UseEnumString {
			p.g.P("encoder.AppendValueString(", varName, ".String()", ")")
		} else {
			p.g.P("encoder.AppendValueInt32(int32(", varName, ".Number()", "))")
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

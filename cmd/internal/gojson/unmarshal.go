package gojson

import (
	"fmt"
	"strconv"

	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/cmd/internal/gojson/pkg/importpkg"
)

func (p *Plugin) generateCodeUnmarshal(fields []*protogen.Field) {
	msg := p.message

	p.g.P("// UnmarshalJSON implements json.Unmarshaler for proto message ", msg.Desc.Name(), " in file ", p.file.Desc.Path())
	p.g.P("func (x *", msg.GoIdent.GoName, ") UnmarshalJSON(b []byte) error {")
	p.g.P("    if x == nil {")
	p.g.P("        return ", importpkg.Errors.Ident("New"), `("json: Unmarshal: `, string(msg.GoIdent.GoImportPath), ".(*", msg.GoIdent.GoName, `) is nil")`)
	p.g.P("    }")
	if len(fields) != 0 {
		// Generated flag variables to check oneof are loaded.
		p.unmarshalGenVariables(fields)

		// Generated scan code.
		p.unmarshalLoopScan(fields)
	}
	p.g.P("    return nil")
	p.g.P("}")
}
func (p *Plugin) unmarshalGenVariables(fields []*protogen.Field) {
	// Generated flag variables to check oneof are loaded.
	for _, field := range fields {
		if !pkfield.FieldIsOneOf(field) {
			continue
		}
		options := p.loadOneofOptions(field.Oneof)
		if options.Ignore {
			continue
		}
		varIsFill := p.genVariableOneofIsFill(field.Oneof.GoName)
		p.g.P("var ", varIsFill, " bool")
	}
	p.g.P("")
}
func (p *Plugin) unmarshalLoopScan(fields []*protogen.Field) {
	// prepare
	p.g.P("var (")
	p.g.P("    err error")
	p.g.P("    isNULL bool")
	p.g.P("    decoder *", importpkg.PJDecoder.Ident("Decoder"))
	p.g.P(")")
	p.g.P("if decoder, err = ", importpkg.PJDecoder.Ident("New"), "(b); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL, err = decoder.BeforeScanJSON(); err != nil {")
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

func (p *Plugin) unmarshalLoopRead(fields []*protogen.Field) {
	p.g.P("var (")
	p.g.P("    jsonKey string")
	p.g.P("    isEnd bool")
	p.g.P(")")
	p.g.P("if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isEnd {")
	p.g.P("    break LOOP_SCAN")
	p.g.P("}")
	p.g.P("if jsonKey, err = decoder.ReadJSONKey(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("switch jsonKey { // match the JSON key")
LOOP:
	for _, field := range fields {
		var jsonKey string
		if pkfield.FieldIsOneOf(field) {
			options := p.loadOneofOptions(field.Oneof)
			if options.Ignore {
				continue LOOP
			}
			if options.Hide {
				p.unmarshalOneOfField(field)
				continue LOOP
			}
			jsonKey = p.getJSONKeyForOneof(options, field.Oneof)
		} else {
			options := p.loadFieldOptions(field)
			if options.Ignore {
				continue LOOP
			}
			jsonKey = p.getJSONKeyForField(options, field)
		}
		p.g.P("case ", strconv.Quote(jsonKey), ":")
		switch {
		case pkfield.FieldIsOneOf(field):
			p.unmarshalOneOf(field)
		case field.Desc.IsMap():
			p.unmarshalMap(field)
		case field.Desc.IsList():
			p.unmarshalRepeated(field)
		default:
			p.unmarshalPlain(field)
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
	// enc switch
	p.g.P("}")
}

func (p *Plugin) unmarshalOneOf(field *protogen.Field) {
	// Read and parse oneof.
	p.g.P("if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL {")
	p.g.P("    x.", field.Oneof.GoName, " = nil")
	p.g.P("    continue LOOP_SCAN")
	p.g.P("}")
	p.g.P("for {")
	p.g.P("    var oneofKey string")
	p.g.P("	if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {")
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
	p.g.P("    }")
	p.g.P("}")
}

func (p *Plugin) unmarshalMap(field *protogen.Field) {
	goName := field.GoName
	goType := pkfield.FieldGoType(p.g, field)

	// Read and parse map.
	p.g.P("if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL {")
	p.g.P("    x.", field.GoName, " = nil")
	p.g.P("    continue LOOP_SCAN")
	p.g.P("}")
	p.g.P("if x.", goName, " == nil {")
	p.g.P("    x.", goName, " = ", "make(", goType, ")")
	p.g.P("}")
	p.g.P("for {") // start loop.
	p.g.P("	if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {")
	p.g.P("    	return err ")
	p.g.P("	}")
	p.g.P("	if isEnd {")
	p.g.P("    	break")
	p.g.P("	}")
	p.unmarshalReadMapKey(field)
	p.unmarshalReadMapValue(field)
	p.g.P("}") // end loop
}

func (p *Plugin) unmarshalRepeated(field *protogen.Field) {
	goName := field.GoName
	goType := pkfield.FieldGoType(p.g, field)

	// Read and parse array. same as the standard json.
	p.g.P("if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if isNULL {")
	p.g.P("    x.", field.GoName, " = nil")
	p.g.P("    continue LOOP_SCAN")
	p.g.P("}")
	p.g.P("if x.", goName, " == nil {")
	p.g.P("    x.", goName, " = ", "make(", goType, ", 0)")
	p.g.P("}")
	p.g.P("i := 0")
	p.g.P("length := len(x.", goName, ")")
	p.g.P("for {") // start loop.
	p.g.P("	if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {")
	p.g.P("   		return err")
	p.g.P("	}")
	p.g.P("	if isEnd {")
	p.g.P("   		break")
	p.g.P("	}")
	p.unmarshalReadArrayElem(field)
	p.g.P("}") // end loop.
	p.g.P("if i < length {")
	p.g.P("    // truncate the slice to Consistent with standard library json.")
	p.g.P("    x.", field.GoName, " = x.", field.GoName, "[:i]")
	p.g.P("}")
}

func (p *Plugin) unmarshalPlain(field *protogen.Field) {
	p.unmarshalReadLiteralValue(field)
}

func (p *Plugin) unmarshalOneOfField(field *protogen.Field) {
	for _, oneField := range field.Oneof.Fields {
		options := p.loadFieldOptions(field)
		if options.Ignore {
			continue
		}
		jsonKey := p.getJSONKeyForField(options, oneField)
		p.g.P("case ", strconv.Quote(jsonKey), ":")
		p.unmarshalReadLiteralValue(oneField)
	}
}

func (p *Plugin) unmarshalReadMapKey(field *protogen.Field) {
	_field := field.Message.Fields[0]
	goType := p.fieldGoType(_field)
	p.g.P("var mapKey ", goType)

	switch _field.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("mapKey, err = decoder.ReadMapKeyInt32(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("mapKey, err = decoder.ReadMapKeyInt64(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("mapKey, err = decoder.ReadMapKeyUint32(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("mapKey, err = decoder.ReadMapKeyUint64(jsonKey)")
	case protoreflect.StringKind:
		p.g.P("mapKey, err = decoder.ReadMapKeyString(jsonKey)")
	//case protoreflect.BoolKind:
	//	p.g.P("mapKey, err = decoder.ReadMapKeyBool(jsonKey)")
	//	checkKeyError()
	default:
		panic(fmt.Sprintf(
			"gojson: umarshal: unsupported type of map key, field: %s, kind: %s",
			field.Desc.FullName(), field.Desc.MapKey().Kind().String(),
		))
	}
	p.g.P("if err != nil {")
	p.g.P("    return err ")
	p.g.P("}")
}
func (p *Plugin) unmarshalReadMapValue(field *protogen.Field) {
	_field := field.Message.Fields[1]

	goType := p.fieldGoType(_field)
	if _field.Desc.Kind() == protoreflect.MessageKind {
		p.g.P("var mapVal *", goType)
	} else {
		p.g.P("var mapVal ", goType)
	}

	switch _field.Desc.Kind() {
	case protoreflect.StringKind:
		p.g.P("mapVal, err = decoder.ReadMapValueString", "(jsonKey)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("mapVal, err = decoder.ReadMapValueInt32", "(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("mapVal, err = decoder.ReadMapValueInt64", "(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("mapVal, err = decoder.ReadMapValueUint32", "(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("mapVal, err = decoder.ReadMapValueUint64", "(jsonKey)")
	case protoreflect.FloatKind:
		p.g.P("mapVal, err = decoder.ReadMapValueFloat32", "(jsonKey)")
	case protoreflect.DoubleKind:
		p.g.P("mapVal, err = decoder.ReadMapValueFloat64", "(jsonKey)")
	case protoreflect.BoolKind:
		p.g.P("mapVal, err = decoder.ReadMapValueBool", "(jsonKey)")
	case protoreflect.BytesKind:
		p.g.P("mapVal, err = decoder.ReadMapValueBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		p.g.P("initFN := func() interface{} {")
		p.g.P("	mapVal = x.", field.GoName, "[mapKey]")
		p.g.P("	if mapVal == nil {")
		p.g.P("		mapVal = new(", goType, ")")
		p.g.P("	}")
		p.g.P("     return mapVal")
		p.g.P("}")
		p.g.P("err = decoder.ReadMapValueInterface", "(jsonKey, initFN)")
	case protoreflect.EnumKind:
		options := p.loadFieldOptions(field)
		p.g.P("var v1 int32")
		if options.UseEnumString {
			p.g.P("v1, err = decoder.ReadMapValueEnumString", "(jsonKey,", goType, "_value)")
		} else {
			p.g.P("v1, err = decoder.ReadMapValueEnumNumber", "(jsonKey,", goType, "_name)")
		}
		p.g.P("mapVal = ", goType, "(v1)")
	default:
		panic(fmt.Sprintf(
			"gojson: unmarshal: unsupported kind of %s as value, field: %s", _field.Desc.Kind().String(), _field.Desc.FullName(),
		))
	}
	// Check error.
	p.g.P("if err != nil {")
	p.g.P("	return err")
	p.g.P("}")
	// Save the value.
	p.g.P("x.", field.GoName, "[mapKey] = mapVal")
}
func (p *Plugin) unmarshalReadArrayElem(field *protogen.Field) {
	goType := p.fieldGoType(field)
	if field.Desc.Kind() == protoreflect.MessageKind {
		p.g.P("var vv *", goType)
	} else {
		p.g.P("var vv ", goType)
	}

	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		p.g.P("vv, err = decoder.ReadArrayElemString", "(jsonKey)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("vv, err = decoder.ReadArrayElemInt32", "(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("vv, err = decoder.ReadArrayElemInt64", "(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("vv, err = decoder.ReadArrayElemUint32", "(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("vv, err = decoder.ReadArrayElemUint64", "(jsonKey)")
	case protoreflect.FloatKind:
		p.g.P("vv, err = decoder.ReadArrayElemFloat32", "(jsonKey)")
	case protoreflect.DoubleKind:
		p.g.P("vv, err = decoder.ReadArrayElemFloat64", "(jsonKey)")
	case protoreflect.BoolKind:
		p.g.P("vv, err = decoder.ReadArrayElemBool", "(jsonKey)")
	case protoreflect.BytesKind:
		p.g.P("vv, err = decoder.ReadArrayElemBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		p.g.P("initFN := func() interface{} {")
		p.g.P("	if i < length {")
		p.g.P("		vv = x.", field.GoName, "[i]")
		p.g.P("	}")
		p.g.P("	if vv == nil {")
		p.g.P("		vv = new(", goType, ")")
		p.g.P("	}")
		p.g.P("     return vv")
		p.g.P("}")
		p.g.P("err = decoder.ReadArrayElemInterface", "(jsonKey, initFN)")
	case protoreflect.EnumKind:
		options := p.loadFieldOptions(field)
		p.g.P("var v1 int32")
		if options.UseEnumString {
			p.g.P("v1, err = decoder.ReadArrayElemEnumString", "(jsonKey,", goType, "_value)")
		} else {
			p.g.P("v1, err = decoder.ReadArrayElemEnumNumber", "(jsonKey,", goType, "_name)")
		}
		p.g.P("vv = ", goType, "(v1)")
	default:
		panic(fmt.Sprintf(
			"gojson: unmarshal: unsupported kind of %s as value, field: %s", field.Desc.Kind().String(), field.Desc.FullName(),
		))
	}
	// Check error.
	p.g.P("if err != nil {")
	p.g.P("	return err")
	p.g.P("}")
	// Save the value.
	p.g.P("if i < length {")
	p.g.P("    x.", field.GoName, "[i] = vv")
	p.g.P("} else {")
	p.g.P("    x.", field.GoName, " = append(", "x.", field.GoName, ", vv", ")")
	p.g.P("}")
	p.g.P("i++")
}
func (p *Plugin) unmarshalReadLiteralValue(field *protogen.Field) {
	isOneOf := pkfield.FieldIsOneOf(field)
	isOptional := pkfield.FieldIsOptional(field)

	if isOneOf {
		oneOfName := field.Oneof.GoName
		oneOfType := p.g.QualifiedGoIdent(field.GoIdent)
		p.g.P("var ok bool")
		p.g.P("var ot *", oneOfType)
		p.g.P("if ot, ok = x.", oneOfName, ".(*", oneOfType, "); !ok {")
		p.g.P("    ot = new(", oneOfType, ")")
		p.g.P("}")
	}

	goType := p.fieldGoType(field)
	if isOptional || field.Desc.Kind() == protoreflect.MessageKind {
		p.g.P("var vv *", goType)
	} else {
		p.g.P("var vv ", goType)
	}

	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerString", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueString", "(jsonKey)")
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerInt32", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueInt32", "(jsonKey)")
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerInt64", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueInt64", "(jsonKey)")
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerUint32", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueUint32", "(jsonKey)")
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerUint64", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueUint64", "(jsonKey)")
		}
	case protoreflect.FloatKind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerFloat32", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueFloat32", "(jsonKey)")
		}
	case protoreflect.DoubleKind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerFloat64", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueFloat64", "(jsonKey)")
		}
	case protoreflect.BoolKind:
		if isOptional {
			p.g.P("vv, err = decoder.ReadPointerBool", "(jsonKey)")
		} else {
			p.g.P("vv, err = decoder.ReadValueBool", "(jsonKey)")
		}
	case protoreflect.BytesKind:
		p.g.P("vv, err = decoder.ReadValueBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		var receiver string
		switch {
		case isOneOf:
			receiver = "ot"
		default:
			receiver = "x"
		}
		p.g.P("initFN := func () interface{} {")
		p.g.P("    if ", receiver, ".", field.GoName, " != nil {")
		p.g.P("        vv = ", receiver, ".", field.GoName)
		p.g.P("    } else {")
		p.g.P("        vv = new(", goType, ")")
		p.g.P("    }")
		p.g.P("    return vv")
		p.g.P("}")
		p.g.P("err = decoder.ReadValueInterface", "(jsonKey, initFN)")
	case protoreflect.EnumKind:
		options := p.loadFieldOptions(field)
		if isOptional {
			p.g.P("var v1 *int32")
			if options.UseEnumString {
				p.g.P("v1, err = decoder.ReadPointerEnumString", "(jsonKey,", goType, "_value)")
			} else {
				p.g.P("v1, err = decoder.ReadPointerEnumNumber", "(jsonKey,", goType, "_name)")
			}
			p.g.P("if v1 != nil {")
			p.g.P("	v2 := ", goType, "(*v1)")
			p.g.P("	vv = &v2")
			p.g.P("}")
		} else {
			p.g.P("var v1 int32")
			if options.UseEnumString {
				p.g.P("v1, err = decoder.ReadValueEnumString", "(jsonKey,", goType, "_value)")
			} else {
				p.g.P("v1, err = decoder.ReadValueEnumNumber", "(jsonKey,", goType, "_name)")
			}
			p.g.P("vv = ", goType, "(v1)")
		}
	default:
		panic(fmt.Sprintf(
			"gojson: unmarshal: unsupported kind of %s as value, field: %s", field.Desc.Kind().String(), field.Desc.FullName(),
		))
	}
	// Check error.
	p.g.P("if err != nil {")
	p.g.P("	return err")
	p.g.P("}")
	// Save the value.
	switch {
	case isOneOf:
		oneOfName := field.Oneof.GoName
		varIsFill := p.genVariableOneofIsFill(oneOfName)
		p.g.P("if ", varIsFill, " {")
		p.g.P("    return ", importpkg.FMT.Ident("Errorf"), `("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)`)
		p.g.P("}")
		p.g.P(varIsFill, " = true")
		p.g.P("ot.", field.GoName, " = vv")
		p.g.P("x.", oneOfName, " = ot")
	default:
		p.g.P("x.", field.GoName, " = vv")
	}
}

func (p *Plugin) genVariableOneofIsFill(oneofName string) string {
	return "oneOfIsFill_" + oneofName
}

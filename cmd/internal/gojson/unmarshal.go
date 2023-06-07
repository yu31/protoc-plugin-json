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
	// End function.
	p.g.P("}")
}
func (p *Plugin) unmarshalGenVariables(fields []*protogen.Field) {
	// Generated flag variables to check oneof are loaded.
	for _, field := range fields {
		if !pkfield.FieldIsOneOf(field) {
			continue
		}
		options := p.loadOneofOptions(field.Oneof)
		if !(options.Ignore) {
			oneOfName := field.Oneof.GoName
			p.g.P("var ", p.genVariableOneofIsFill(oneOfName), " bool")
		}
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
	p.g.P("if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {")
	p.g.P("    return err")
	p.g.P("}")

	// loop scan object begin.
	p.g.P("// Loop to scan object.")
	p.g.P("LOOP_OBJECT:")
	p.g.P("for {")
	// Check error in every loop.
	p.g.P("	if err = decoder.ScanError(); err != nil {")
	p.g.P("    	return err")
	p.g.P("	}")

	p.unmarshalLoopRead(fields)

	// Loop scan object end
	p.g.P("}")
	p.g.P("// Check error in decoder")
	p.g.P("if err = decoder.ScanError(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
}

func (p *Plugin) unmarshalLoopRead(fields []*protogen.Field) {
	p.g.P("jsonKey, stop := decoder.ReadJSONKey()")
	p.g.P("if stop {")
	p.g.P("    break LOOP_OBJECT")
	p.g.P("}")
	p.g.P("switch { // match the JSON KEY")
LOOP:
	for _, field := range fields {
		var jsonKey string
		if pkfield.FieldIsOneOf(field) {
			options := p.loadOneofOptions(field.Oneof)
			if options.Ignore {
				continue LOOP
			}
			if options.Hide {
				p.unmarshalOneOfField(field, "jsonKey")
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
		p.g.P("case jsonKey == ", strconv.Quote(jsonKey), ":")
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
		p.g.P("return ", importpkg.FMT.Ident("Errorf"), `("json: unknown field %q", jsonKey)`)
	} else {
		p.g.P("if err = decoder.Discard(); err != nil { // discard unknown field")
		p.g.P("    return err")
		p.g.P("}")
	}
	// enc switch
	p.g.P("}")
	// After read value..
	p.unmarshalReadObjectValueAfter("LOOP_OBJECT")
}

func (p *Plugin) unmarshalOneOf(field *protogen.Field) {
	loopLabel := "LOOP_ONEOF_" + string(field.Oneof.Desc.Name())

	decodeOneOf := func() {
		p.g.P(loopLabel, ":")
		p.g.P("for {")
		p.g.P("	oneofKey, stop := decoder.ReadJSONKey()")
		p.g.P("	if stop {")
		p.g.P("    	break ", loopLabel)
		p.g.P("	}")

		// Read and process value.
		p.g.P("    switch {")
		p.unmarshalOneOfField(field, "oneofKey")
		p.g.P("    default:")
		if p.msgOptions.DisallowUnknownFields {
			p.g.P("    return ", importpkg.FMT.Ident("Errorf"), `("json: unknown field %q in oneof parts", oneofKey)`)
		} else {
			p.g.P("if err = decoder.Discard(); err != nil { // discard unknown field")
			p.g.P("    return err")
			p.g.P("}")
		}
		p.g.P("    }")

		p.unmarshalReadObjectValueAfter(loopLabel)
		// end loop.
		p.g.P("}")
		p.g.P("decoder.ScanNext()")
	}

	// read oneof
	p.g.P("if isNULL, err = decoder.CheckObjectBegin(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("switch {")
	p.g.P("case isNULL:")
	p.g.P("    x.", field.Oneof.GoName, " = nil")
	p.g.P("default:")
	decodeOneOf()
	p.g.P("}")
}

func (p *Plugin) unmarshalMap(field *protogen.Field) {
	goName := field.GoName
	goType := pkfield.FieldGoType(p.g, field)
	loopLabel := "LOOP_MAP_" + string(field.Desc.Name())

	decodeMap := func() {
		// create map if not initialized. Same as the standard json.Unmarshal
		p.g.P("if x.", goName, " == nil {")
		p.g.P("    x.", goName, " = ", "make(", goType, ")")
		p.g.P("}")
		p.g.P(loopLabel, ":")
		p.g.P("for {") // start loop.
		p.g.P("	if err = decoder.ScanError(); err != nil {")
		p.g.P("    	return err")
		p.g.P("	}")
		p.unmarshalReadMapKey(field)
		p.unmarshalReadMapValue(field)
		p.g.P("	if noMore {")
		p.g.P("    	break ", loopLabel)
		p.g.P("	}")
		p.g.P("}") // end loop
		p.g.P("decoder.ScanNext()")
	}

	// read map.
	p.g.P("var isEmpty bool")
	p.g.P("if isNULL, isEmpty, err = decoder.CheckBeforeReadMap(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("switch {")
	p.g.P("case isNULL:")
	p.g.P("    x.", field.GoName, " = nil")
	p.g.P("case isEmpty:")
	p.g.P("    // do nothing")
	p.g.P("default:")
	decodeMap()
	p.g.P("}")
}

func (p *Plugin) unmarshalRepeated(field *protogen.Field) {
	goName := field.GoName
	goType := pkfield.FieldGoType(p.g, field)

	loopLabel := "LOOP_LIST_" + string(field.Desc.Name())

	decodeList := func() {
		p.g.P("if x.", goName, " == nil {")
		p.g.P("    x.", goName, " = ", "make(", goType, ", 0)")
		p.g.P("}")

		p.g.P("i := 0")
		p.g.P("length := len(x.", goName, ")")
		p.g.P(loopLabel, ":")
		p.g.P("for ; ; {") // start loop.

		p.g.P("	if err = decoder.ScanError(); err != nil {")
		p.g.P("    	return err")
		p.g.P("	}")
		p.unmarshalReadArrayElem(field)
		p.g.P("	if noMore { // After read array value.")
		p.g.P("   		break ", loopLabel)
		p.g.P("	}")
		p.g.P("}") // end loop.
		p.g.P("if i < length {")
		p.g.P("    // truncate the slice to Consistent with standard library json.")
		p.g.P("    x.", field.GoName, " = x.", field.GoName, "[:i]")
		p.g.P("}")
		p.g.P("decoder.ScanNext()")
	}

	// read array. same as the standard json.
	p.g.P("var isEmpty bool")
	p.g.P("if isNULL, isEmpty, err = decoder.CheckBeforeReadArray(jsonKey); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("switch {")
	p.g.P("case isNULL:")
	p.g.P("    x.", field.GoName, " = nil")
	p.g.P("case isEmpty:")
	p.g.P("	if x.", goName, " == nil {")
	p.g.P("    	x.", goName, " = ", "make(", goType, ", 0)")
	p.g.P("	} else {")
	p.g.P("         x.", goName, " = ", "x.", goName, "[:0]")
	p.g.P("     }")
	p.g.P("default:")
	decodeList()
	p.g.P("}")
}

func (p *Plugin) unmarshalPlain(field *protogen.Field) {
	p.unmarshalReadLiteralValue(field)
}

func (p *Plugin) unmarshalOneOfField(field *protogen.Field, keyVariable string) {
	for _, oneField := range field.Oneof.Fields {
		options := p.loadFieldOptions(field)
		if options.Ignore {
			continue
		}
		p.g.P("case ", keyVariable, " == ", `"`, p.getJSONKeyForField(options, oneField), `"`, ":")
		p.unmarshalReadLiteralValue(oneField)
	}
}

func (p *Plugin) unmarshalReadMapKey(field *protogen.Field) {
	switch field.Desc.MapKey().Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("mapKey, _err := decoder.ReadMapKeyInt32(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("mapKey, _err := decoder.ReadMapKeyInt64(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("mapKey, _err := decoder.ReadMapKeyUint32(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("mapKey, _err := decoder.ReadMapKeyUint64(jsonKey)")
	case protoreflect.StringKind:
		p.g.P("mapKey, _err := decoder.ReadMapKeyString(jsonKey)")
	//case protoreflect.BoolKind:
	//	p.g.P("mapKey, _err := decoder.ReadMapKeyBool(jsonKey)")
	//	checkKeyError()
	default:
		panic(fmt.Sprintf(
			"gojson: umarshal: unsupported type of map key, field: %s, kind: %s",
			field.Desc.FullName(), field.Desc.MapKey().Kind().String(),
		))
	}
	p.g.P("if _err != nil {")
	p.g.P("    return _err ")
	p.g.P("}")
}
func (p *Plugin) unmarshalReadMapValue(field *protogen.Field) {
	_field := field.Message.Fields[1]
	switch _field.Desc.Kind() {
	case protoreflect.StringKind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueString", "(jsonKey)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueInt32", "(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueInt64", "(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueUint32", "(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueUint64", "(jsonKey)")
	case protoreflect.FloatKind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueFloat32", "(jsonKey)")
	case protoreflect.DoubleKind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueFloat64", "(jsonKey)")
	case protoreflect.BoolKind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueBool", "(jsonKey)")
	case protoreflect.BytesKind:
		p.g.P("vv, noMore, _err := decoder.ReadMapValueBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		valueType := p.g.QualifiedGoIdent(_field.Message.GoIdent)
		p.g.P("var vv *", valueType)
		p.g.P("initFN := func() interface{} {")
		p.g.P("	vv = x.", field.GoName, "[mapKey]")
		p.g.P("	if vv == nil {")
		p.g.P("		vv = new(", valueType, ")")
		p.g.P("	}")
		p.g.P("     return vv")
		p.g.P("}")
		p.g.P("noMore, _err := decoder.ReadMapValueInterface", "(jsonKey, initFN)")
	case protoreflect.EnumKind:
		options := p.loadFieldOptions(field)
		enumType := p.g.QualifiedGoIdent(_field.Enum.GoIdent)
		if options.UseEnumString {
			p.g.P("v1, noMore, _err := decoder.ReadMapValueEnumString", "(jsonKey,", enumType, "_value)")
		} else {
			p.g.P("v1, noMore, _err := decoder.ReadMapValueEnumNumber", "(jsonKey,", enumType, "_name)")
		}
		p.g.P("	vv := ", enumType, "(v1)")
	default:
		panic(fmt.Sprintf(
			"gojson: unmarshal: unsupported kind of %s as value, field: %s", _field.Desc.Kind().String(), _field.Desc.FullName(),
		))
	}
	// Check error.
	p.g.P("if _err != nil {")
	p.g.P("	return _err")
	p.g.P("}")
	// Save the value.
	p.g.P("x.", field.GoName, "[mapKey] = vv")
}
func (p *Plugin) unmarshalReadArrayElem(field *protogen.Field) {
	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemString", "(jsonKey)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemInt32", "(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemInt64", "(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemUint32", "(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemUint64", "(jsonKey)")
	case protoreflect.FloatKind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemFloat32", "(jsonKey)")
	case protoreflect.DoubleKind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemFloat64", "(jsonKey)")
	case protoreflect.BoolKind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemBool", "(jsonKey)")
	case protoreflect.BytesKind:
		p.g.P("vv, noMore, _err := decoder.ReadArrayElemBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		valueType := p.g.QualifiedGoIdent(field.Message.GoIdent)
		p.g.P("var vv *", valueType)
		p.g.P("initFN := func() interface{} {")
		p.g.P("	if i < length {")
		p.g.P("		vv = x.", field.GoName, "[i]")
		p.g.P("	}")
		p.g.P("	if vv == nil {")
		p.g.P("		vv = new(", valueType, ")")
		p.g.P("	}")
		p.g.P("     return vv")
		p.g.P("}")
		p.g.P("noMore, _err := decoder.ReadArrayElemInterface", "(jsonKey, initFN)")
	case protoreflect.EnumKind:
		options := p.loadFieldOptions(field)
		enumType := p.g.QualifiedGoIdent(field.Enum.GoIdent)
		if options.UseEnumString {
			p.g.P("v1, noMore, _err := decoder.ReadArrayElemEnumString", "(jsonKey,", enumType, "_value)")
		} else {
			p.g.P("v1, noMore, _err := decoder.ReadArrayElemEnumNumber", "(jsonKey,", enumType, "_name)")
		}
		p.g.P("	vv := ", enumType, "(v1)")
	default:
		panic(fmt.Sprintf(
			"gojson: unmarshal: unsupported kind of %s as value, field: %s", field.Desc.Kind().String(), field.Desc.FullName(),
		))
	}
	// Check error.
	p.g.P("if _err != nil {")
	p.g.P("	return _err")
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
		p.g.P("var ot *", oneOfType)
		p.g.P("if _ot, _ok := x.", oneOfName, ".(*", oneOfType, "); _ok {")
		p.g.P("    ot = _ot")
		p.g.P("} else {")
		p.g.P("    ot = new(", oneOfType, ")")
		p.g.P("}")
	}

	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerString", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueString", "(jsonKey)")
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerInt32", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueInt32", "(jsonKey)")
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerInt64", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueInt64", "(jsonKey)")
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerUint32", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueUint32", "(jsonKey)")
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerUint64", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueUint64", "(jsonKey)")
		}
	case protoreflect.FloatKind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerFloat32", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueFloat32", "(jsonKey)")
		}
	case protoreflect.DoubleKind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerFloat64", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueFloat64", "(jsonKey)")
		}
	case protoreflect.BoolKind:
		if isOptional {
			p.g.P("vv, _err := decoder.ReadPointerBool", "(jsonKey)")
		} else {
			p.g.P("vv, _err := decoder.ReadValueBool", "(jsonKey)")
		}
	case protoreflect.BytesKind:
		p.g.P("vv, _err := decoder.ReadValueBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		valueType := p.g.QualifiedGoIdent(field.Message.GoIdent)
		var receiver string
		switch {
		case isOneOf:
			receiver = "ot"
		default:
			receiver = "x"
		}
		p.g.P("var vv *", valueType)
		p.g.P("initFN := func () interface{} {")
		p.g.P("    if ", receiver, ".", field.GoName, " != nil {")
		p.g.P("        vv = ", receiver, ".", field.GoName)
		p.g.P("    } else {")
		p.g.P("        vv = new(", valueType, ")")
		p.g.P("    }")
		p.g.P("    return vv")
		p.g.P("}")
		p.g.P("_err := decoder.ReadValueInterface", "(jsonKey, initFN)")
	case protoreflect.EnumKind:
		options := p.loadFieldOptions(field)
		enumType := p.g.QualifiedGoIdent(field.Enum.GoIdent)
		if isOptional {
			p.g.P("var vv *", enumType)
			if options.UseEnumString {
				p.g.P("v1, _err := decoder.ReadPointerEnumString", "(jsonKey,", enumType, "_value)")
			} else {
				p.g.P("v1, _err := decoder.ReadPointerEnumNumber", "(jsonKey,", enumType, "_name)")
			}
			p.g.P("if v1 != nil {")
			p.g.P("	_vv := ", enumType, "(*v1)")
			p.g.P("	vv = &_vv")
			p.g.P("}")
		} else {
			if options.UseEnumString {
				p.g.P("v1, _err := decoder.ReadValueEnumString", "(jsonKey,", enumType, "_value)")
			} else {
				p.g.P("v1, _err := decoder.ReadValueEnumNumber", "(jsonKey,", enumType, "_name)")
			}
			p.g.P("	vv := ", enumType, "(v1)")
		}
	default:
		panic(fmt.Sprintf(
			"gojson: unmarshal: unsupported kind of %s as value, field: %s", field.Desc.Kind().String(), field.Desc.FullName(),
		))
	}
	// Check error.
	p.g.P("if _err != nil {")
	p.g.P("	return _err")
	p.g.P("}")
	// Save the value.
	switch {
	case isOneOf:
		oneOfName := field.Oneof.GoName
		//oneOfType := p.g.QualifiedGoIdent(field.GoIdent)
		varIsFill := p.genVariableOneofIsFill(oneOfName)
		p.g.P("if ", varIsFill, " {")
		p.g.P("    return ", importpkg.FMT.Ident("Errorf"), `("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)`)
		p.g.P("}")
		p.g.P(varIsFill, " = true")
		//p.g.P("ot := new(", oneOfType, ")")
		p.g.P("ot.", field.GoName, " = vv")
		p.g.P("x.", oneOfName, " = ot")
	default:
		p.g.P("x.", field.GoName, " = vv")
	}
}

func (p *Plugin) genVariableOneofIsFill(oneofName string) string {
	return "oneOfIsFill_" + oneofName
}

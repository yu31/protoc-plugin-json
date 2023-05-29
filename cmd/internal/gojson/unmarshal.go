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
			p.g.P("var ", p.genVariableOneofIsStore(oneOfName), " bool")
		}
	}
}
func (p *Plugin) unmarshalLoopScan(fields []*protogen.Field) {
	// prepare
	p.g.P("decoder, err := ", importpkg.PJDecoder.Ident("New"), "(b)")
	p.g.P("if err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("if ok, _err := decoder.CheckJSONBegin(); _err != nil || ok {")
	p.g.P("    return _err")
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
	// Before read key, Read opening charsets(`"`) of string key or closing charsets(`}`).
	p.unmarshalReadObjectKeyBefore("LOOP_OBJECT")
	p.g.P("// Read JSON key.")
	p.g.P("jsonKey := decoder.ReadObjectKey()")

	// Before read value
	p.unmarshalReadObjectValueBefore()
	p.g.P("// match field with JSON key.")
	p.g.P("switch {")
LOOP:
	for _, field := range fields {
		var jsonKey string
		if pkfield.FieldIsOneOf(field) {
			options := p.loadOneofOptions(field.Oneof)
			if options.Ignore {
				continue LOOP
			}
			if options.Hide {
				p.unmarshalDecodeOneOf(field, "jsonKey")
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
		p.g.P("_ = decoder.ReadItem() // discard unknown field")
	}
	// enc switch
	p.g.P("}")
	// After read value..
	p.unmarshalReadObjectValueAfter("LOOP_OBJECT")
}

func (p *Plugin) unmarshalOneOf(field *protogen.Field) {
	goType := pkfield.FieldGoType(p.g, field)
	loopLabel := "LOOP_ONEOF_" + string(field.Oneof.Desc.Name())

	decodeOneOf := func() {
		p.g.P(loopLabel, ":")
		p.g.P("for {")

		// Before read key, Read opening " of string key or closing }.
		p.unmarshalReadObjectKeyBefore(loopLabel)
		p.g.P("    oneofKey := decoder.ReadObjectKey() // Read key")
		// Before read value
		p.unmarshalReadObjectValueBefore()

		// Read and process value.
		p.g.P("    switch {")
		p.unmarshalDecodeOneOf(field, "oneofKey")
		p.g.P("    default:")
		if p.msgOptions.DisallowUnknownFields {
			p.g.P("    return ", importpkg.FMT.Ident("Errorf"), `("json: unknown oneof field %q", oneofKey)`)
		} else {
			p.g.P("_ = decoder.ReadItem() // discard unknown field")
		}
		p.g.P("    }")

		p.unmarshalReadObjectValueAfter(loopLabel)
		// end loop.
		p.g.P("}")
		p.g.P("decoder.ScanNext()")
	}

	// Check whether null.
	p.g.P("if decoder.OpCode == ", importpkg.PJDecoder.Ident("ScanLiteralBegin"), " {")
	p.g.P("    value := decoder.ReadItem()")
	p.g.P("    if value[0] != 'n' {")
	p.g.P("        return ", importpkg.FMT.Ident("Errorf"), `("json: cannot unmarshal %s as oneof into field %s of type `, goType, `", string(value), jsonKey)`)
	p.g.P("    }")
	p.g.P("} else {")
	// check is object.
	p.g.P("    if decoder.OpCode != ", importpkg.PJDecoder.Ident("ScanObjectBegin"), " {")
	p.g.P("        value := decoder.ReadItem()")
	p.g.P("        return ", importpkg.FMT.Ident("Errorf"), `("json: cannot unmarshal %s as oneof into field %s of type `, goType, `", string(value), jsonKey)`)
	p.g.P("    }")
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
		p.g.P("for {")

		p.g.P("if err = decoder.ScanError(); err != nil {")
		p.g.P("    return err")
		p.g.P("}")

		// Before read key, Read opening " of string key or closing }.
		p.unmarshalReadObjectKeyBefore(loopLabel)

		p.unmarshalReadMapKey(field)

		// Before read value
		p.unmarshalReadObjectValueBefore()

		// read and decode value.
		p.unmarshalReadMapValue(field)

		// After read value.
		p.unmarshalReadObjectValueAfter(loopLabel)

		// End loop.
		p.g.P("}")
		p.g.P("decoder.ScanNext()")
	}

	// Check whether null.
	p.g.P("if decoder.OpCode == ", importpkg.PJDecoder.Ident("ScanLiteralBegin"), " {")
	p.g.P("    value := decoder.ReadItem()")
	p.g.P("    if value[0] != 'n' {")
	p.g.P("        return ", importpkg.FMT.Ident("Errorf"), `("json: cannot unmarshal %s as map into field %s of type `, goType, `", string(value), jsonKey)`)
	p.g.P("    } else {")
	p.g.P("        x.", field.GoName, " = nil")
	p.g.P("    }")
	p.g.P("} else {")

	// check is map.
	p.g.P("    if decoder.OpCode != ", importpkg.PJDecoder.Ident("ScanObjectBegin"), " {")
	p.g.P("        value := decoder.ReadItem()")
	p.g.P("        return ", importpkg.FMT.Ident("Errorf"), `("json: cannot unmarshal %s as map into field %s of type `, goType, `", string(value), jsonKey)`)
	p.g.P("    }")

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
		p.g.P("for {")

		p.g.P("if err = decoder.ScanError(); err != nil {")
		p.g.P("    return err")
		p.g.P("}")

		// Before Read value.
		p.unmarshalReadArrayValueBefore(loopLabel)

		// Read list value.
		p.unmarshalReadArrayElem(field)

		// After read value.
		p.unmarshalReadArrayValueAfter(loopLabel)

		// end LOOP_LIST.
		p.g.P("}")

		// truncate the slice if necessary.
		p.g.P("if i < length {")
		p.g.P("    x.", field.GoName, " = x.", field.GoName, "[:i]")
		p.g.P("}")

		p.g.P("decoder.ScanNext()")
	}

	// Check whether null.
	p.g.P("if decoder.OpCode == ", importpkg.PJDecoder.Ident("ScanLiteralBegin"), " {")
	p.g.P("    value := decoder.ReadItem()")
	p.g.P("    if value[0] != 'n' {")
	p.g.P("        return ", importpkg.FMT.Ident("Errorf"), `("json: cannot unmarshal %s as array into field %s of type `, goType, `", string(value), jsonKey)`)
	p.g.P("    } else {")
	p.g.P("        x.", field.GoName, " = nil")
	p.g.P("    }")
	p.g.P("} else {")

	// check is list.
	p.g.P("    if decoder.OpCode != ", importpkg.PJDecoder.Ident("ScanArrayBegin"), " {")
	p.g.P("        value := decoder.ReadItem()")
	p.g.P("        return ", importpkg.FMT.Ident("Errorf"), `("json: cannot unmarshal %s as array into field %s of type `, goType, `", string(value), jsonKey)`)
	p.g.P("    }")

	decodeList()

	p.g.P("}")
}

func (p *Plugin) unmarshalPlain(field *protogen.Field) {
	if field.Desc.IsMap() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsMap; field: %s", field.Desc.FullName()))
	}
	if field.Desc.IsList() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsList; field: %s", field.Desc.FullName()))
	}
	if field.Desc.IsWeak() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsWeak; field: %s", field.Desc.FullName()))
	}
	if field.Desc.IsExtension() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsExtension; filed: %s", field.Desc.FullName()))
	}
	if field.Desc.IsPlaceholder() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsPlaceholder; filed: %s", field.Desc.FullName()))
	}
	if field.Desc.IsPacked() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsPacked; filed: %s", field.Desc.FullName()))
	}

	p.unmarshalReadLiteralValue(field)
}

func (p *Plugin) unmarshalDecodeOneOf(field *protogen.Field, keyVariable string) {
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
	options := p.loadFieldOptions(field)
	goName := field.GoName
	_field := field.Message.Fields[1]

	switch _field.Desc.Kind() {
	case protoreflect.StringKind:
		p.g.P("vv, _err := decoder.ReadMapValueString", "(jsonKey)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("vv, _err := decoder.ReadMapValueInt32", "(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("vv, _err := decoder.ReadMapValueInt64", "(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("vv, _err := decoder.ReadMapValueUint32", "(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("vv, _err := decoder.ReadMapValueUint64", "(jsonKey)")
	case protoreflect.FloatKind:
		p.g.P("vv, _err := decoder.ReadMapValueFloat32", "(jsonKey)")
	case protoreflect.DoubleKind:
		p.g.P("vv, _err := decoder.ReadMapValueFloat64", "(jsonKey)")
	case protoreflect.BoolKind:
		p.g.P("vv, _err := decoder.ReadMapValueBool", "(jsonKey)")
	case protoreflect.BytesKind:
		p.g.P("vv, _err := decoder.ReadMapValueBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		valueType := p.g.QualifiedGoIdent(_field.Message.GoIdent)
		p.g.P("vv := x.", goName, "[mapKey]")
		p.g.P("if vv == nil {")
		p.g.P("	vv = new(", valueType, ")")
		p.g.P("}")
		p.g.P("ok, _err := decoder.ReadMapValueInterface", "(jsonKey, vv)")
		p.g.P("if !ok { // The field is null")
		p.g.P("	vv = nil")
		p.g.P("}")
	case protoreflect.EnumKind:
		enumType := p.g.QualifiedGoIdent(_field.Enum.GoIdent)
		if options.UseEnumString {
			p.g.P("v1, _err := decoder.ReadMapValueEnumString", "(jsonKey,", enumType, "_value)")
		} else {
			p.g.P("v1, _err := decoder.ReadMapValueEnumNumber", "(jsonKey,", enumType, "_name)")
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
	options := p.loadFieldOptions(field)
	goName := field.GoName
	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		p.g.P("vv, _err := decoder.ReadArrayElemString", "(jsonKey)")
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("vv, _err := decoder.ReadArrayElemInt32", "(jsonKey)")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("vv, _err := decoder.ReadArrayElemInt64", "(jsonKey)")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("vv, _err := decoder.ReadArrayElemUint32", "(jsonKey)")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("vv, _err := decoder.ReadArrayElemUint64", "(jsonKey)")
	case protoreflect.FloatKind:
		p.g.P("vv, _err := decoder.ReadArrayElemFloat32", "(jsonKey)")
	case protoreflect.DoubleKind:
		p.g.P("vv, _err := decoder.ReadArrayElemFloat64", "(jsonKey)")
	case protoreflect.BoolKind:
		p.g.P("vv, _err := decoder.ReadArrayElemBool", "(jsonKey)")
	case protoreflect.BytesKind:
		p.g.P("vv, _err := decoder.ReadArrayElemBytes", "(jsonKey)")
	case protoreflect.MessageKind:
		valueType := p.g.QualifiedGoIdent(field.Message.GoIdent)
		p.g.P("var vv *", valueType)
		p.g.P("if i < length {")
		p.g.P("	vv = x.", goName, "[i]")
		p.g.P("}")
		p.g.P("if vv == nil {")
		p.g.P("	vv = new(", valueType, ")")
		p.g.P("}")
		p.g.P("ok, _err := decoder.ReadArrayElemInterface", "(jsonKey, vv)")
		p.g.P("if !ok { // The field is null")
		p.g.P("	vv = nil")
		p.g.P("}")
	case protoreflect.EnumKind:
		enumType := p.g.QualifiedGoIdent(field.Enum.GoIdent)
		if options.UseEnumString {
			p.g.P("v1, _err := decoder.ReadArrayElemEnumString", "(jsonKey,", enumType, "_value)")
		} else {
			p.g.P("v1, _err := decoder.ReadArrayElemEnumNumber", "(jsonKey,", enumType, "_name)")
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
	p.g.P("    x.", goName, "[i] = vv")
	p.g.P("} else {")
	p.g.P("    x.", goName, " = append(", "x.", goName, ", vv", ")")
	p.g.P("}")
	p.g.P("i++")
}
func (p *Plugin) unmarshalReadLiteralValue(field *protogen.Field) {
	options := p.loadFieldOptions(field)

	goName := field.GoName
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
		switch {
		case isOneOf:
			p.g.P("var vv *", valueType)
			p.g.P("if ot.", field.GoName, " != nil {")
			p.g.P("    vv = ot.", field.GoName)
			p.g.P("} else {")
			p.g.P("    vv = new(", valueType, ")")
			p.g.P("}")
		default:
			p.g.P("var vv *", valueType)
			p.g.P("if x.", goName, " == nil {")
			p.g.P("	vv = new(", valueType, ")")
			p.g.P("} else {")
			p.g.P("	vv = x.", goName)
			p.g.P("}")
		}
		p.g.P("ok, _err := decoder.ReadValueInterface", "(jsonKey, vv)")
		p.g.P("if !ok { // The field is null")
		p.g.P("	vv = nil")
		p.g.P("}")
	case protoreflect.EnumKind:
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
		varIsStore := p.genVariableOneofIsStore(oneOfName)
		p.g.P("if ", varIsStore, " {")
		p.g.P("    return ", importpkg.FMT.Ident("Errorf"), `("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)`)
		p.g.P("}")
		p.g.P(varIsStore, " = true")
		//p.g.P("ot := new(", oneOfType, ")")
		p.g.P("ot.", field.GoName, " = vv")
		p.g.P("x.", oneOfName, " = ot")
	default:
		p.g.P("x.", field.GoName, " = vv")
	}
}

//func (p *Plugin) unmarshalReadLiteralValue(field *protogen.Field) {
//	options := p.loadFieldOptions(field)
//
//	isMap := field.Desc.IsMap()
//	isList := field.Desc.IsList()
//	isOneOf := pkfield.FieldIsOneOf(field)
//	isOptional := pkfield.FieldIsOptional(field)
//
//	goName := field.GoName
//
//	var _field *protogen.Field
//
//	if isMap {
//		_field = field.Message.Fields[1]
//	} else {
//		_field = field
//	}
//
//	switch _field.Desc.Kind() {
//	case protoreflect.StringKind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerString", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueString", "(jsonKey)")
//		}
//	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerInt32", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueInt32", "(jsonKey)")
//		}
//	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerInt64", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueInt64", "(jsonKey)")
//		}
//	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerUint32", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueUint32", "(jsonKey)")
//		}
//	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerUint64", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueUint64", "(jsonKey)")
//		}
//	case protoreflect.FloatKind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerFloat32", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueFloat32", "(jsonKey)")
//		}
//	case protoreflect.DoubleKind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerFloat64", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueFloat64", "(jsonKey)")
//		}
//	case protoreflect.BoolKind:
//		if isOptional {
//			p.g.P("vv, _err := decoder.ReadPointerBool", "(jsonKey)")
//		} else {
//			p.g.P("vv, _err := decoder.ReadValueBool", "(jsonKey)")
//		}
//	case protoreflect.BytesKind:
//		p.g.P("vv, _err := decoder.ReadValueBytes", "(jsonKey)")
//	case protoreflect.MessageKind:
//		valueType := p.g.QualifiedGoIdent(_field.Message.GoIdent)
//		p.g.P("var vv *", valueType)
//		switch {
//		case isMap:
//			p.g.P("vv = x.", goName, "[mapKey]")
//			p.g.P("if vv == nil {")
//			p.g.P("vv = new(", valueType, ")")
//			p.g.P("}")
//		case isList:
//			p.g.P("if i < length {")
//			p.g.P("vv = x.", goName, "[i]")
//			p.g.P("}")
//			p.g.P("if vv == nil {")
//			p.g.P("vv = new(", valueType, ")")
//			p.g.P("}")
//		case isOneOf:
//			p.g.P("vv = new(", valueType, ")")
//		default:
//			p.g.P("if x.", goName, " == nil {")
//			p.g.P("	vv = new(", valueType, ")")
//			p.g.P("} else {")
//			p.g.P("	vv = x.", goName)
//			p.g.P("}")
//		}
//		p.g.P("ok, _err := decoder.ReadValueInterface", "(jsonKey, vv)")
//		p.g.P("if !ok { // The field is null")
//		p.g.P("	vv = nil")
//		p.g.P("}")
//	case protoreflect.EnumKind:
//		enumType := p.g.QualifiedGoIdent(_field.Enum.GoIdent)
//		if isOptional {
//			p.g.P("var vv *", enumType)
//			if options.UseEnumString {
//				p.g.P("v1, _err := decoder.ReadPointerEnumString", "(jsonKey,", enumType, "_value)")
//			} else {
//				p.g.P("v1, _err := decoder.ReadPointerEnumNumber", "(jsonKey,", enumType, "_name)")
//			}
//			p.g.P("if v1 != nil {")
//			p.g.P("	_vv := ", enumType, "(*v1)")
//			p.g.P("	vv = &_vv")
//			p.g.P("}")
//		} else {
//			if options.UseEnumString {
//				p.g.P("v1, _err := decoder.ReadValueEnumString", "(jsonKey,", enumType, "_value)")
//			} else {
//				p.g.P("v1, _err := decoder.ReadValueEnumNumber", "(jsonKey,", enumType, "_name)")
//			}
//			p.g.P("	vv := ", enumType, "(v1)")
//		}
//	default:
//		panic(fmt.Sprintf(
//			"gojson: unmarshal: unsupported kind of %s as value, field: %s", _field.Desc.Kind().String(), _field.Desc.FullName(),
//		))
//	}
//
//	// Check error.
//	p.g.P("if _err != nil {")
//	p.g.P("	return _err")
//	p.g.P("}")
//
//	// Save the value.
//	switch {
//	case isMap:
//		p.g.P("x.", field.GoName, "[mapKey] = vv")
//	case isList:
//		p.g.P("if i < length {")
//		p.g.P("    x.", goName, "[i] = vv")
//		p.g.P("} else {")
//		p.g.P("    x.", goName, " = append(", "x.", goName, ", vv", ")")
//		p.g.P("}")
//		p.g.P("i++")
//	case isOneOf:
//		oneOfName := field.Oneof.GoName
//		oneOfType := p.g.QualifiedGoIdent(field.GoIdent)
//		varIsStore := p.genVariableOneofIsStore(oneOfName)
//		p.g.P("if ", varIsStore, " {")
//		p.g.P("    return ", importpkg.FMT.Ident("Errorf"), `("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)`)
//		p.g.P("}")
//		p.g.P(varIsStore, " = true")
//		p.g.P("ot := new(", oneOfType, ")")
//		p.g.P("ot.", field.GoName, " = vv")
//		p.g.P("x.", oneOfName, " = ot")
//	default:
//		p.g.P("x.", field.GoName, " = vv")
//	}
//}

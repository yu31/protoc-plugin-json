package gojson

import (
	"os"
	"strconv"
	"strings"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
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
		formatMap := loadTypeFormatMap(options.Format)
		_mkName, _mkVars := loadMapKeyTypeInfo(fieldSet.Field.Message.Fields[0], formatMap.Key)
		_mvName, _mvVars := loadValueTypeInfo(fieldSet.Field.Message.Fields[1], formatMap.Value)

		funcName = "ReadMap" + _mkName + _mvName
		funcVars = append(funcVars, fieldValue)
		funcVars = append(funcVars, _mkVars...)
		funcVars = append(funcVars, _mvVars...)
	case fieldSet.Field.Desc.IsList():
		formatRepeated := loadTypeFormatRepeated(options.Format)
		_typeName, _vars := loadValueTypeInfo(fieldSet.Field, formatRepeated.Elem)

		funcName = "ReadList" + _typeName
		funcVars = append(funcVars, fieldValue)
		funcVars = append(funcVars, _vars...)
	default:
		_typeName, _vars := loadValueTypeInfo(fieldSet.Field, options.Format)

		if pkfield.FieldIsOptional(fieldSet.Field) {
			funcName = "ReadPtr" + _typeName
		} else {
			funcName = "ReadVal" + _typeName
		}
		switch fieldSet.Field.Desc.Kind() {
		case protoreflect.EnumKind, protoreflect.MessageKind:
			funcVars = append(funcVars, fieldValue)
		}
		// FIXME: we should also pass the field value as parameters if the field type if optional?
		funcVars = append(funcVars, _vars...)
	}

	funcGoIdent := importpkg.PJDecoder.Ident(funcName)
	parameters := strings.Join(funcVars, ",")
	ul.g.P("if ", fieldValue, ", err = ", funcGoIdent, "(", parameters, "); err != nil {")
	ul.g.P("    return err")
	ul.g.P("}")
}

func (ul *Unmarshal) declareLocalVariables(fieldSets []*FieldSet) {
	// local variables for check if the oneof field is loaded.
	if onfOfFields := loadOneOfFields(fieldSets); len(onfOfFields) != 0 {
		ul.g.P("// declares variables to report whether the oneof field is loaded.")
		ul.g.P("var (")
		for _, oneOfField := range onfOfFields {
			ul.g.P(ul.genVarNameOneOfIsLoad(oneOfField), " bool")
		}
		ul.g.P(")")
		ul.g.P("")
	}

	// local variables for reference the parent field.
	if inlineFields := loadInlineFields(fieldSets); len(inlineFields) != 0 {
		ul.g.P("// declares variables to simple to reference parent field")
		ul.g.P("var (")
		for _, fieldSet := range inlineFields {
			parent := fieldSet.ParentField()
			ul.g.P(parent.Alias() + " " + pkfield.FieldGoType(ul.g, parent.Field))
		}
		ul.g.P(")")
		ul.g.P("")
		ul.g.P("// declares anonymous to init the parent field.")
		for _, fieldSet := range inlineFields {
			ul.declareFuncInitParentField(fieldSet)
		}
		ul.g.P("")
	}
}

func (ul *Unmarshal) declareFuncInitParentField(fieldSet *FieldSet) {
	parentField := fieldSet.ParentField()
	parentAlias := parentField.Alias()
	parentValue := parentField.Value()

	ul.g.P(ul.genFuncNameInitParent(fieldSet), " := func() error {") // func started
	{
		ul.g.P("if ", parentAlias, " == nil {") // if started
		{
			ul.initParentFieldForAnonymousFunc(parentField)
			ul.prepareInitOneOfPart(parentField)
			ul.g.P("if ", parentValue, " == nil {")
			// Init the parent value.
			ul.g.P("    ", parentValue, " =  new(", fieldGoType(ul.g, parentField.Field), ")")
			ul.g.P("}")
			// Init the parentAlias
			ul.g.P(parentAlias, " = ", parentValue)
		}
		ul.g.P("}") // if ended.
		ul.g.P("return nil")
	}
	ul.g.P("}") // func ended
}

func (ul *Unmarshal) initParentFieldForAnonymousFunc(fieldSet *FieldSet) {
	if fieldSet.ParentField() == nil {
		return
	}
	// FIXME: skip if the parent is oneof part and the oneof field not inline?
	//if fieldSet.IsOneOfPart && !fieldSet.OneOfField().OneOfIsInline() {
	//	return
	//}
	ul.g.P("if _err := ", ul.genFuncNameInitParent(fieldSet), "(); _err != nil {")
	ul.g.P("    return _err")
	ul.g.P("}")
}
func (ul *Unmarshal) prepareInitParentField(fieldSet *FieldSet) {
	if fieldSet.ParentField() == nil {
		return
	}
	if fieldSet.IsOneOfPart && !fieldSet.OneOfField().OneOfIsInline() {
		return
	}
	ul.g.P("if err = ", ul.genFuncNameInitParent(fieldSet), "(); err != nil {")
	ul.g.P("    return err")
	ul.g.P("}")
}
func (ul *Unmarshal) prepareInitOneOfPart(fieldSet *FieldSet) {
	if !fieldSet.IsOneOfPart {
		return
	}

	oneOfField := fieldSet.OneOfField()
	oneOfIsLoad := ul.genVarNameOneOfIsLoad(oneOfField)
	ul.g.P("if ", oneOfIsLoad, " {")
	ul.g.P("	return ", importpkg.PJDecoder.Ident("ErrOneOfConflict"), `(dec)`)
	ul.g.P("}")
	ul.g.P(oneOfIsLoad, " = true")

	oneOfFieldAlias := oneOfField.Alias()
	oneOfFieldValue := oneOfField.Value()
	oneOfPartIdent := ul.g.QualifiedGoIdent(fieldSet.Field.GoIdent)

	ul.g.P(oneOfFieldAlias, ", ok := ", oneOfFieldValue, ".(*", oneOfPartIdent, ")")
	ul.g.P("if !ok {")
	{
		ul.g.P(oneOfFieldAlias, " = new(", oneOfPartIdent, ")")
		ul.g.P(oneOfFieldValue, " = ", oneOfFieldAlias)
	}
	ul.g.P("}")
}

func (ul *Unmarshal) genVarNameOneOfIsLoad(fieldSet *FieldSet) string {
	if !fieldSet.IsOneOfField() {
		panic("unknown logic error when generate variable name for check oneof field is loaded")
	}
	return "isLoad_" + fieldSet.Alias()
}
func (ul *Unmarshal) genFuncNameInitParent(fieldSet *FieldSet) string {
	parent := fieldSet.ParentField()
	if parent == nil {
		panic("unknown logic error when generate func name for init parent field")
	}
	return "init_" + parent.Alias()
}

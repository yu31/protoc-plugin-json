package gojson

import (
	"github.com/yu31/protoc-kit-go/utils/pkfield"

	"github.com/yu31/protoc-plugin-json/cmd/internal/gojson/pkg/importpkg"
)

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

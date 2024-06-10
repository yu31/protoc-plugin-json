package gojson

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/compiler/protogen"
)

func truncateBufLen(n int) int {
	const base = 8
	return (n/base + 1) * base
}

func fieldGoType(protoGen *protogen.GeneratedFile, field *protogen.Field) (goType string) {
	if field.Desc.IsWeak() {
		//return "struct{}", false
		panic(fmt.Sprintf("unsupported case IsWeak; field: %s", field.Desc.FullName()))
	}

	switch field.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = protoGen.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = protoGen.QualifiedGoIdent(field.Message.GoIdent)
	}
	return goType
}

func loadOneOfFields(fieldSets []*FieldSet) (oneOfFields []*FieldSet) {
	for _, fieldSet := range fieldSets {
		if fieldSet.FieldIsInline() {
			oneOfFields = append(oneOfFields, loadOneOfFields(fieldSet.InlineSet.Childs)...)
		}
		if fieldSet.IsOneOfField() && len(fieldSet.OneOfSet.Parts) != 0 {
			oneOfFields = append(oneOfFields, loadOneOfFields(fieldSet.OneOfSet.Parts)...)
			oneOfFields = append(oneOfFields, fieldSet)
		}
	}
	return
}

func loadInlineFields(fieldSets []*FieldSet) (inlineFields []*FieldSet) {
	_inlineFields := _loadInlineFields(fieldSets)
	length := len(_inlineFields)
	if length == 0 {
		return
	}
	// Deduplication
	_exists := make(map[*FieldSet]struct{}, length)
	inlineFields = make([]*FieldSet, 0, length)
	for i := len(_inlineFields) - 1; i >= 0; i-- {
		fieldSet := _inlineFields[i]
		parent := fieldSet.ParentField()
		if _, ok := _exists[parent]; !ok {
			inlineFields = append(inlineFields, fieldSet)
			_exists[parent] = struct{}{}
		}
	}
	return
}
func _loadInlineFields(fieldSets []*FieldSet) (parentFields []*FieldSet) {
	for _, fieldSet := range fieldSets {
		if fieldSet.FieldIsInline() {
			parentFields = append(parentFields, _loadInlineFields(fieldSet.InlineSet.Childs)...)
		}
		if fieldSet.IsOneOfField() && len(fieldSet.OneOfSet.Parts) != 0 {
			parentFields = append(parentFields, _loadInlineFields(fieldSet.OneOfSet.Parts)...)
		}
		if parent := fieldSet.ParentField(); parent != nil {
			parentFields = append(parentFields, fieldSet)
		}
	}
	return
}

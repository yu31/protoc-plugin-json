package gojson

import (
	"fmt"

	"github.com/yu31/protoc-kit-go/helper/pkwkt"
	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"

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

func getJSONKeyForField(field *protogen.Field, options *pbjson.FieldOptions) string {
	jsonKey := options.Json
	if options.Ignore {
		panic("the field should be ignore.")
	}
	if jsonKey != nil {
		return *jsonKey
	}
	return string(field.Desc.Name())
}

func getJSONKeyForOneOf(field *protogen.Field, options *pbjson.OneofOptions) string {
	jsonKey := options.Json
	if options.Ignore {
		panic("the oneof field should be ignore.")
	}
	if jsonKey != nil {
		return *jsonKey
	}
	return string(field.Oneof.Desc.Name())
}

// fieldIsInline return whether the field need to directly expanded in current areas.
// The field must be type message and not type well-known.
// And invalid for map and repeated.
func fieldIsInline(field *protogen.Field, options *pbjson.FieldOptions) bool {
	if field.Desc.IsMap() || field.Desc.IsList() {
		return false
	}
	if field.Desc.Kind() != protoreflect.MessageKind {
		return false
	}
	switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
	case pkwkt.Any, pkwkt.Duration, pkwkt.Timestamp:
		return false
	}
	return *options.Inline
}

type IdGenerator struct {
	id int64
}

func (x *IdGenerator) Take() int64 {
	x.id++
	return x.id
}
func (x *IdGenerator) Back() {
	x.id--
}

func loadFieldSets(idGen *IdGenerator, fields []*protogen.Field, parent *FieldSet, isOneOfPart bool) (fieldSets []*FieldSet) {
	fieldSets = make([]*FieldSet, 0, len(fields))

	var fieldLevel int
	switch {
	case isOneOfPart: // represents the fields is oneof parts.
		fieldLevel = parent.Level
	case parent != nil: // represents the fields is inline fields.
		fieldLevel = parent.Level + 1
	default:
		fieldLevel = 1 // The top fieldLevel.
	}

LOOP:
	for _, field := range fields {
		var (
			isOneOfField bool
			oneOfOptions *pbjson.OneofOptions
			fieldOptions *pbjson.FieldOptions
		)

		if pkfield.FieldIsOneOf(field) && !isOneOfPart { // Is the oneof field.
			if field.Oneof.Fields[0] != field {
				continue LOOP // only generate for first appearance
			}
			if oneOfOptions = loadOneOfOptions(field); oneOfOptions.Ignore {
				// The field and its subfields will all be ignored.
				continue LOOP
			}
			isOneOfField = true
		} else {
			if fieldOptions = loadFieldOptions(field); fieldOptions.Ignore {
				// The field is marked as ignored.
				continue LOOP
			}
		}

		fieldSet := &FieldSet{
			Field:       field,
			Options:     nil,
			JSONKey:     "",
			OneOfSet:    nil,
			IsOneOfPart: isOneOfPart,
			InlineSet:   nil,
			Parent:      parent,
			Level:       fieldLevel,
			Number:      idGen.Take(),
		}
		if isOneOfField {
			parts := loadFieldSets(idGen, field.Oneof.Fields, fieldSet, true)
			//if len(parts) == 0 {
			//	// Empty parts means no valid field in oneof areas(all the fields are marked as ignored).
			//	idGen.Back()
			//	continue LOOP
			//}
			fieldSet.JSONKey = getJSONKeyForOneOf(field, oneOfOptions)
			fieldSet.OneOfSet = &OneOfSet{Options: oneOfOptions, Parts: parts}
		} else {
			fieldSet.JSONKey = getJSONKeyForField(field, fieldOptions)
			fieldSet.Options = fieldOptions
			if fieldIsInline(field, fieldOptions) {
				childs := loadFieldSets(idGen, field.Message.Fields, fieldSet, false)
				//if len(childs) == 0 {
				//	// Empty childs means no valid field int the inline message
				//	// (all the fields are marked as ignored or the no fields int the fields).
				//	idGen.Back()
				//	continue LOOP
				//}
				fieldSet.InlineSet = &InlineSet{Childs: childs}
			}
		}
		fieldSets = append(fieldSets, fieldSet)
	}
	return
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

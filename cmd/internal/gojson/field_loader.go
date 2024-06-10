package gojson

import (
	"github.com/yu31/protoc-kit-go/helper/pkwkt"
	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

// IdGenerator for generate a unique id for every field.
type IdGenerator struct {
	id int64
}

func (x *IdGenerator) Take() int64 {
	x.id++
	return x.id
}

// FieldLoader for load the proto field into the fieldSet.
type FieldLoader struct{}

func (x *FieldLoader) Load(file *protogen.File, msg *protogen.Message) (fieldSets []*FieldSet) {
	idGen := &IdGenerator{}
	fieldSets = x.load(idGen, msg.Fields, nil, false)
	checkFields(file, msg, fieldSets)
	return
}

func (x *FieldLoader) load(idGen *IdGenerator, fields []*protogen.Field, parent *FieldSet, isOneOfPart bool) (fieldSets []*FieldSet) {
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
			parts := x.load(idGen, field.Oneof.Fields, fieldSet, true)
			fieldSet.JSONKey = x.getJSONKeyForOneOf(field, oneOfOptions)
			fieldSet.OneOfSet = &OneOfSet{Options: oneOfOptions, Parts: parts}
		} else {
			fieldSet.JSONKey = x.getJSONKeyForField(field, fieldOptions)
			fieldSet.Options = fieldOptions
			if x.fieldIsInline(field, fieldOptions) {
				childs := x.load(idGen, field.Message.Fields, fieldSet, false)
				fieldSet.InlineSet = &InlineSet{Childs: childs}
			}
		}
		fieldSets = append(fieldSets, fieldSet)
	}
	return
}

func (x *FieldLoader) getJSONKeyForField(field *protogen.Field, options *pbjson.FieldOptions) string {
	jsonKey := options.Json
	if options.Ignore {
		panic("the field should be ignore")
	}
	if jsonKey != nil {
		return *jsonKey
	}
	return string(field.Desc.Name())
}

func (x *FieldLoader) getJSONKeyForOneOf(field *protogen.Field, options *pbjson.OneofOptions) string {
	jsonKey := options.Json
	if options.Ignore {
		panic("the oneof field should be ignore")
	}
	if jsonKey != nil {
		return *jsonKey
	}
	return string(field.Oneof.Desc.Name())
}

// fieldIsInline return whether the field need to directly expanded in current areas.
// The field must be type message and not type well-known.
// And false for map and repeated.
func (x *FieldLoader) fieldIsInline(field *protogen.Field, options *pbjson.FieldOptions) bool {
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

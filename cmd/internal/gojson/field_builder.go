package gojson

import (
	"fmt"
	"os"

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

// FieldBuilder for load the proto field into the fieldSet.
type FieldBuilder struct {
	file    *protogen.File
	message *protogen.Message

	idGen *IdGenerator
}

func (x *FieldBuilder) Build(file *protogen.File, msg *protogen.Message) []*FieldSet {
	x.file = file
	x.message = msg
	x.idGen = &IdGenerator{}

	fieldSets := x.load(msg.Fields, nil, false)
	return fieldSets
}

func (x *FieldBuilder) load(fields []*protogen.Field, parent *FieldSet, isOneOfPart bool) (fieldSets []*FieldSet) {
	x.checkCycleReference(parent)

	var fieldLevel int
	switch {
	case isOneOfPart: // represents the fields is oneof parts.
		fieldLevel = parent.Level
	case parent != nil: // represents the fields is inline fields.
		fieldLevel = parent.Level + 1
	default:
		fieldLevel = 1 // The top fieldLevel.
	}

	fieldSets = make([]*FieldSet, 0, len(fields))
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
			Number:      x.idGen.Take(),
		}
		if isOneOfField {
			parts := x.load(field.Oneof.Fields, fieldSet, true)
			fieldSet.JSONKey = x.getJSONKeyForOneOf(field, oneOfOptions)
			fieldSet.OneOfSet = &OneOfSet{Options: oneOfOptions, Parts: parts}
		} else {
			fieldSet.JSONKey = x.getJSONKeyForField(field, fieldOptions)
			fieldSet.Options = fieldOptions
			if x.fieldIsInline(field, fieldOptions) {
				childs := x.load(field.Message.Fields, fieldSet, false)
				fieldSet.InlineSet = &InlineSet{Childs: childs}
			}
		}
		fieldSets = append(fieldSets, fieldSet)
	}
	return
}

func (x *FieldBuilder) checkCycleReference(parent *FieldSet) {
	if parent == nil {
		return
	}
	// Check cycle reference.
	for prev := parent.ParentField(); prev != nil; prev = prev.ParentField() {
		if parent.Field.Desc.FullName() == prev.Field.Desc.FullName() {
			id := genMessageID(x.file, x.message)
			println(fmt.Sprintf("%s - Found cycle reference in field %s", id, parent.Field.Desc.Name()))
			os.Exit(1)
		}
	}
}

func (x *FieldBuilder) getJSONKeyForField(field *protogen.Field, options *pbjson.FieldOptions) string {
	jsonKey := options.Json
	if options.Ignore {
		panic("the field should be ignore")
	}
	if jsonKey != nil {
		return *jsonKey
	}
	return string(field.Desc.Name())
}
func (x *FieldBuilder) getJSONKeyForOneOf(field *protogen.Field, options *pbjson.OneofOptions) string {
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
// And also not type map and repeated.
func (x *FieldBuilder) fieldIsInline(field *protogen.Field, options *pbjson.FieldOptions) bool {
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

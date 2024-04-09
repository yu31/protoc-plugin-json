package gojson

import (
	"strconv"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

type OneOfSet struct {
	// Options represents the Options for oneof.
	Options *pbjson.OneofOptions

	// Parts represents the list of fields which location in the oneof field.
	Parts []*FieldSet
}

type InlineSet struct {
	// Childs represents the list of fields which inline to the message.
	Childs []*FieldSet
}

type FieldSet struct {
	// Field represents the origin object in the protogen.
	Field *protogen.Field

	// Options represents the Options for field.
	Options *pbjson.FieldOptions

	// JSONKey represents the keyword in JSON format.
	JSONKey string

	// And the value is not nil means the field type if oneof.
	OneOfSet *OneOfSet

	// IsOneOfPart represents whether the field is part member in an oneof field.
	IsOneOfPart bool

	// InlineSet the value is not nil means the field is inline.
	InlineSet *InlineSet

	// Parent represents the field to which the inline message belongs or the oneof field for oneof part.
	Parent *FieldSet

	// Level is the level of the message to which the field belongs.
	Level int

	// Number is the unique number for this field.
	Number int64
}

// FieldIsInline indicates whether directly expanded in current ares.
func (x *FieldSet) FieldIsInline() bool {
	return x.InlineSet != nil
}

func (x *FieldSet) OneOfIsInline() bool {
	return *x.OneOfSet.Options.Inline
}

// IsOneOfField return the field whether is oneof.
func (x *FieldSet) IsOneOfField() bool {
	return x.OneOfSet != nil
}

func (x *FieldSet) OneOfField() *FieldSet {
	return x.Parent
}

func (x *FieldSet) ParentField() *FieldSet {
	if x.IsOneOfPart {
		return x.Parent.Parent
	}
	return x.Parent
}

func (x *FieldSet) GoName() string {
	if x.IsOneOfField() {
		return x.Field.Oneof.GoName
	}
	return x.Field.GoName
}

func (x *FieldSet) ID() string {
	return strconv.Itoa(x.Level) + "_" + strconv.FormatInt(x.Number, 10)
}

// Alias represents the variable name of the field alias.
func (x *FieldSet) Alias() string {
	switch {
	case x == nil:
		return "x"
	case x.IsOneOfPart:
		return "p" + x.ID()
	case x.IsOneOfField():
		return "o" + x.ID()
	default:
		return "a" + x.ID()
	}
}

// Value represents the variable name of the field value.
func (x *FieldSet) Value() string {
	return x.Parent.Alias() + "." + x.GoName()
}

package gojson

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// TODO: supported set by parameters?
const (
	assumeBufLenUnknown = 32
	assumeBufLenArray   = 8
	assumeBufLenMap     = 8
	alignmentBufLen     = 8
)

func guessEncodeBufLen(filedSets []*FieldSet) (bufLen int) {
	// Sum json beginning(`{`) and closing(`}`).
	bufLen = 2
	bufLen += sumBufLen(filedSets)
	// truncate to alignment.
	bufLen = (bufLen/alignmentBufLen + 1) * alignmentBufLen
	return
}

func sumBufLen(filedSets []*FieldSet) (bufLen int) {
	for _, fieldSet := range filedSets {
		bufLen += getFieldBufLen(fieldSet)
		switch {
		case fieldSet.FieldIsInline():
			bufLen += guessEncodeBufLen(fieldSet.InlineSet.Childs)
		case fieldSet.IsOneOfField():
			bufLen += guessEncodeBufLen(fieldSet.OneOfSet.Parts)
		}
	}
	return
}
func getFieldBufLen(fieldSet *FieldSet) (bufLen int) {
	if fieldSet.FieldIsInline() {
		return
	}
	if fieldSet.IsOneOfField() && fieldSet.OneOfIsInline() {
		return
	}
	// Sum length of JSON key and colon separator(`:`).
	bufLen = len(fieldSet.JSONKey) + 1
	if fieldSet.IsOneOfField() {
		// Sum oneof beginning(`{`) and closing(`}`).
		bufLen += 2
		return
	}
	// Sum length field value.
	switch {
	case fieldSet.Field.Desc.IsList():
		bufLen += getValueBufLen(fieldSet.Field) * assumeBufLenArray
	case fieldSet.Field.Desc.IsMap():
		keyLen := getValueBufLen(fieldSet.Field.Message.Fields[0])
		valLen := getValueBufLen(fieldSet.Field.Message.Fields[1])
		bufLen += (keyLen + valLen) * assumeBufLenMap
	default:
		bufLen += getValueBufLen(fieldSet.Field)
	}
	// Sum comma separator(`,`).
	bufLen += 1
	return
}
func getValueBufLen(field *protogen.Field) int {
	var bufLen int
	switch field.Desc.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.FloatKind, protoreflect.BoolKind, protoreflect.EnumKind:
		bufLen = 4
	case protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed64Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind,
		protoreflect.DoubleKind:
		bufLen = 8
	case protoreflect.StringKind, protoreflect.BytesKind:
		bufLen = assumeBufLenUnknown
	case protoreflect.MessageKind, protoreflect.GroupKind:
		bufLen = len(field.Message.Fields) * assumeBufLenUnknown
	}
	return bufLen
}

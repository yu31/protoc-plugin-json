package gojson

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"

	"google.golang.org/protobuf/compiler/protogen"
)

func genVariableOneOfIsFill(oneOfName string) string {
	return "oneOfIsFill_" + oneOfName
}

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

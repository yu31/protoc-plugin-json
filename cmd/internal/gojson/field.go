package gojson

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"

	"google.golang.org/protobuf/compiler/protogen"
)

func (p *Plugin) getJSONKeyForField(fieldOptions *pbjson.FieldOptions, field *protogen.Field) string {
	jsonKey := fieldOptions.Json
	if (fieldOptions.Ignore) || (jsonKey != nil && *jsonKey == "-") {
		panic("the field should be ignore.")
	}
	if jsonKey != nil {
		return *jsonKey
	}
	return string(field.Desc.Name())
}

func (p *Plugin) getJSONKeyForOneof(oneofOptions *pbjson.OneofOptions, oneof *protogen.Oneof) string {
	jsonKey := oneofOptions.Json
	if (oneofOptions.Ignore) || (jsonKey != nil && *jsonKey == "-") {
		panic("the oneof field should be ignore.")
	}
	if jsonKey != nil {
		return *jsonKey
	}
	return string(oneof.Desc.Name())
}

func (p *Plugin) guessBufLength(fields []*protogen.Field) int {
	n := 0

	for _, field := range fields {
		var jsonKey string
		if field.Oneof != nil {
			options := p.loadOneOfOptions(field.Oneof)
			if options.Ignore {
				continue
			}
			jsonKey = p.getJSONKeyForOneof(options, field.Oneof)
		} else {
			options := p.loadFieldOptions(field)
			if options.Ignore {
				continue
			}
			jsonKey = p.getJSONKeyForField(options, field)
		}

		// Sum key length.
		n += len(jsonKey)
		// Sum key/value separator(':')
		n += 1
		// Sum field separator(',')
		n += 1

		// Sum value of length.
	}

	n *= 2

	// Sum object begin and end marker('{', '}').
	n += 2
	return n
}

func (p *Plugin) fieldGoType(field *protogen.Field) (goType string) {
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
		goType = p.g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = p.g.QualifiedGoIdent(field.Message.GoIdent)
	}
	return goType
}

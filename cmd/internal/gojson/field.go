package gojson

import (
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
			options := p.loadOneofOptions(field.Oneof)
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

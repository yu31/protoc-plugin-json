package gojson

import (
	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

type Field struct {
	Field   *protogen.Field
	Options *pbjson.FieldOptions
	JSONKey string
	OneOf   *OneOf
}

type OneOf struct {
	Options *pbjson.OneofOptions
	Parts   []*Field
}

// codes is variables for flag oneof is fill.
func loadFields(msg *protogen.Message) (fields []*Field, bufLen int, variables []string) {
	if len(msg.Fields) == 0 {
		return
	}

	fields = make([]*Field, 0, len(msg.Fields))
	// Sum json beginning(`{`) and closing(`}`).
	bufLen += 2

	for _, field := range msg.Fields {
		isOneOf := pkfield.FieldIsOneOf(field)
		if isOneOf && field.Oneof.Fields[0] != field {
			continue // only generate for first appearance
		}

		var fieldOptions *pbjson.FieldOptions
		var jsonKey string
		var oneOf *OneOf

		if !isOneOf {
			fieldOptions = loadFieldOptions(field)
			if fieldOptions.Ignore {
				continue
			}
			jsonKey = getJSONKeyForField(field, fieldOptions)
		} else {
			oneOfOptions := loadOneOfOptions(field)
			if oneOfOptions.Ignore {
				continue
			}

			parts := make([]*Field, 0, len(field.Oneof.Fields))
			for _, oneField := range field.Oneof.Fields {
				oneOptions := loadFieldOptions(oneField)
				if oneOptions.Ignore {
					continue
				}
				oneKey := getJSONKeyForField(oneField, oneOptions)
				parts = append(parts, &Field{
					Field:   oneField,
					Options: oneOptions,
					JSONKey: oneKey,
					OneOf:   nil,
				})
				// Sum key length, colon separator(`:`) and comma separator(`,`).
				bufLen += len(oneKey) + 2
			}

			if len(parts) == 0 {
				// Empty parts means no valid field in oneof parts.
				continue
			}

			jsonKey = getJSONKeyForOneOf(field, oneOfOptions)
			variables = append(variables, genVariableOneOfIsFill(field.Oneof.GoName))

			oneOf = &OneOf{
				Options: oneOfOptions,
				Parts:   parts,
			}
		}

		fields = append(fields, &Field{
			Field:   field,
			Options: fieldOptions,
			JSONKey: jsonKey,
			OneOf:   oneOf,
		})
		// Sum key length, colon separator(`:`) and comma separator(`,`).
		bufLen += len(jsonKey) + 2
	}
	bufLen = truncateBufLen(bufLen * 2)
	return

}

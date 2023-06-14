package gojson

import (
	"fmt"
	"os"

	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *Plugin) checkJSONKey(fields []*Field) {
	msg := p.message

	cacheFields := make(map[string]*protogen.Field)

	dupOneOfs := make([]map[string][]string, 0)
	emptyOneOfs := make([][]string, 0)

	emptyFields := make([]string, 0)
	dupFields := make(map[string][]string)

	checkFieldDup := func(field *Field) {
		jsonKey := field.JSONKey
		if jsonKey == "" {
			emptyFields = append(emptyFields, string(field.Field.Desc.Name()))
			return
		}
		if _, ok := cacheFields[jsonKey]; ok {
			dupFields[jsonKey] = append(dupFields[jsonKey], string(field.Field.Desc.Name()))
			return
		}
		cacheFields[jsonKey] = field.Field
	}

LOOP:
	for _, field := range fields {
		// Check general field.
		if field.OneOf == nil {
			checkFieldDup(field)
			continue LOOP
		}

		if isOneOf := pkfield.FieldIsOneOf(field.Field); !isOneOf {
			checkFieldDup(field)
			continue
		}
		if field.OneOf.Options.Hide {
			for _, oneField := range field.OneOf.Parts {
				checkFieldDup(oneField)
			}
			continue LOOP
		}
		// oneOf key not hide in json. check it.
		checkFieldDup(field)

		// Check the fields in oneof part.
		cacheOneOf := make(map[string]*protogen.Field)
		dupOneOf := make(map[string][]string)
		emptyOneOf := make([]string, 0)

		for _, oneField := range field.OneOf.Parts {
			jsonKey := oneField.JSONKey
			if jsonKey == "" {
				emptyOneOf = append(emptyOneOf, string(oneField.Field.Desc.Name()))
			} else {
				if _, ok := cacheOneOf[jsonKey]; ok {
					dupOneOf[jsonKey] = append(dupOneOf[jsonKey], string(oneField.Field.Desc.Name()))
					continue
				}
				cacheOneOf[jsonKey] = oneField.Field
			}
		}

		for jsonKey := range dupOneOf {
			if x, ok := cacheOneOf[jsonKey]; ok {
				dupOneOf[jsonKey] = append(dupOneOf[jsonKey], string(x.Desc.Name()))
			}
		}
		if len(dupOneOf) != 0 {
			dupOneOfs = append(dupOneOfs, dupOneOf)
		}
		if len(emptyOneOf) != 0 {
			emptyOneOfs = append(emptyOneOfs, emptyOneOf)
		}
	}

	if len(dupFields) == 0 &&
		len(dupOneOfs) == 0 &&
		len(emptyFields) == 0 &&
		len(emptyOneOfs) == 0 {
		return
	}

	if len(emptyFields) != 0 {
		println(fmt.Sprintf(
			"gojson: <file(%s) message(%s)>: (general) the json key is empty in fields %v",
			string(p.file.GoImportPath), msg.GoIdent.GoName, emptyFields,
		))
	}

	for jsonKey, value := range dupFields {
		if x, ok := cacheFields[jsonKey]; ok {
			value = append(value, string(x.Desc.Name()))
		}
		println(fmt.Sprintf(
			"gojson: <file(%s) message(%s)>: (general) Found duplicate json key [%s] both in fields %v",
			string(p.file.GoImportPath), msg.GoIdent.GoName, jsonKey, value,
		))
	}

	for _, emptyOneOf := range emptyOneOfs {
		println(fmt.Sprintf(
			"gojson: <file(%s) message(%s)>: (oneof) the json key is empty in fields %v",
			string(p.file.GoImportPath), msg.GoIdent.GoName, emptyOneOf,
		))
	}

	for _, dupOneOf := range dupOneOfs {
		for jsonKey, value := range dupOneOf {
			println(fmt.Sprintf(
				"gojson: <file(%s) message(%s)>: (oneof) Found duplicate json key [%s] both in fields %v",
				string(p.file.GoImportPath), msg.GoIdent.GoName, jsonKey, value,
			))
		}
	}
	os.Exit(1)
}

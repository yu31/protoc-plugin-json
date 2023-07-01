package gojson

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func (p *Plugin) checkDuplicateKey(fields []*Field) {
	msg := p.message

	cacheFields := make(map[string]*Field)

	// [oneOfName]: [ [jsonKey]: [partLists] ]
	dupOneOf := make(map[string]map[string][]string)
	// [oneOfName]: [partLists]
	emptyOneOf := make(map[string][]string)

	// [fieldLists]
	emptyFields := make([]string, 0)
	// [jsonKey]: [fieldLists]
	dupFields := make(map[string][]string)

	checkFieldDup := func(field *Field) {
		jsonKey := field.JSONKey

		var fieldName string
		if field.OneOf != nil {
			fieldName = string(field.Field.Oneof.Desc.Name())
		} else {
			fieldName = string(field.Field.Desc.Name())
		}

		if jsonKey == "" {
			emptyFields = append(emptyFields, fieldName)
			return
		}

		if x, ok := cacheFields[jsonKey]; ok {
			if len(dupFields[jsonKey]) == 0 {
				var firstField string
				if x.OneOf != nil {
					firstField = string(x.Field.Oneof.Desc.Name())
				} else {
					firstField = string(x.Field.Desc.Name())
				}
				dupFields[jsonKey] = append(dupFields[jsonKey], firstField)
			}
			dupFields[jsonKey] = append(dupFields[jsonKey], fieldName)
			return
		}
		cacheFields[jsonKey] = field
	}

	checkOneOfDup := func(field *Field) {
		cacheOneOf := make(map[string]*protogen.Field)
		dupParts := make(map[string][]string)
		emptyParts := make([]string, 0)

		for _, oneField := range field.OneOf.Parts {
			jsonKey := oneField.JSONKey
			if jsonKey == "" {
				emptyParts = append(emptyParts, string(oneField.Field.Desc.Name()))
			} else {
				if x, ok := cacheOneOf[jsonKey]; ok {
					if len(dupParts[jsonKey]) == 0 {
						dupParts[jsonKey] = append(dupParts[jsonKey], string(x.Desc.Name()))
					}
					dupParts[jsonKey] = append(dupParts[jsonKey], string(oneField.Field.Desc.Name()))
				} else {
					cacheOneOf[jsonKey] = oneField.Field
				}
			}
		}

		oneOfName := string(field.Field.Oneof.Desc.Name())
		if len(dupParts) != 0 {
			dupOneOf[oneOfName] = dupParts
		}
		if len(emptyParts) != 0 {
			emptyOneOf[oneOfName] = emptyParts
		}
	}

LOOP:
	for _, field := range fields {
		// Check general field.
		if field.OneOf == nil {
			checkFieldDup(field)
			continue LOOP
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
		checkOneOfDup(field)
	}

	if len(dupFields) == 0 && len(dupOneOf) == 0 &&
		len(emptyFields) == 0 && len(emptyOneOf) == 0 {
		return
	}

	filePath := p.file.Desc.Path()
	msgName := string(msg.Desc.FullName())
	if index := strings.Index(msgName, "."); index > 0 {
		msgName = msgName[index+1:]
	}
	id := fmt.Sprintf("[ERROR] - [plugin: gojson] - [ file: %s | message: %s ]", filePath, msgName)

	if len(emptyFields) != 0 {
		println(fmt.Sprintf("%s - The json key is empty in fields %v", id, emptyFields))
	}

	for jsonKey, _fields := range dupFields {
		println(fmt.Sprintf("%s - The json key [%s] are duplicated both in fields %v", id, jsonKey, _fields))
	}

	for oneOfName, parts := range emptyOneOf {
		println(fmt.Sprintf("%s - The json key taht in oneof field [%s] is empty of parts %v", id, oneOfName, parts))
	}

	for oneOfName, dupParts := range dupOneOf {
		for jsonKey, parts := range dupParts {
			println(fmt.Sprintf("%s - The json key [%s] that in oneof field [%s] are duplicated both in parts %v", id, jsonKey, oneOfName, parts))
		}
	}

	os.Exit(1)
}

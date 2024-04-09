package gojson

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// checkFields for valid the fields before generate codes.
func checkFields(file *protogen.File, msg *protogen.Message, fieldSets []*FieldSet) {
	oneOfList := make([]*FieldSet, 0) // used to maintain order
	oneOfMap := make(map[*FieldSet][]*FieldSet)

	var reBuildFields func(_fieldSets []*FieldSet) (_all []*FieldSet)
	reBuildFields = func(_fieldSets []*FieldSet) (_all []*FieldSet) {
		for _, fieldSet := range _fieldSets {
			switch {
			case fieldSet.IsOneOfField() && !fieldSet.OneOfIsInline():
				_all = append(_all, fieldSet)
				oneOfList = append(oneOfList, fieldSet)
				oneOfMap[fieldSet] = append(oneOfMap[fieldSet], reBuildFields(fieldSet.OneOfSet.Parts)...)
			case fieldSet.IsOneOfField() && fieldSet.OneOfIsInline():
				_all = append(_all, reBuildFields(fieldSet.OneOfSet.Parts)...)
			case fieldSet.FieldIsInline():
				_all = append(_all, reBuildFields(fieldSet.InlineSet.Childs)...)
			default:
				_all = append(_all, fieldSet)
			}
		}
		return
	}

	fieldSetALL := reBuildFields(fieldSets)

	var exit bool
	if !checkJSONKey(file, msg, fieldSetALL) {
		exit = true
	}
	for _, oneOfField := range oneOfList {
		if !checkJSONKey(file, msg, oneOfMap[oneOfField]) {
			exit = true
		}
	}

	if exit {
		os.Exit(1)
	}
}

// checkJSONKey for valid the jsonKey.
// Check if is there any duplicate json key.
// Check if is there any emtpy json key.
func checkJSONKey(file *protogen.File, msg *protogen.Message, fieldSets []*FieldSet) (ok bool) {
	// Cache the first occurrence of the JSONKey of field.
	cacheFirst := make(map[string]*FieldSet)

	// The list of all duplicate JSONKey. For fixed output order.
	dupJSONKeys := make([]string, 0)
	// [jsonKey]: [fieldLists]
	dupFieldMap := make(map[string][]string)

	// The field lists of JSONKey is emtpy.
	emptyFields := make([]string, 0)

	checkFieldDup := func(fieldSet *FieldSet) {
		jsonKey := fieldSet.JSONKey
		protoName := fieldProtoName(fieldSet)
		if jsonKey == "" {
			emptyFields = append(emptyFields, protoName)
			return
		}
		if xx, ok := cacheFirst[jsonKey]; ok { // The jsonKey is duplication.
			if len(dupFieldMap[jsonKey]) == 0 {
				dupJSONKeys = append(dupJSONKeys, jsonKey)
				dupFieldMap[jsonKey] = append(dupFieldMap[jsonKey], fieldProtoName(xx))
			}
			dupFieldMap[jsonKey] = append(dupFieldMap[jsonKey], protoName)
		} else {
			cacheFirst[jsonKey] = fieldSet
		}
	}

	for _, fieldSet := range fieldSets {
		checkFieldDup(fieldSet)
	}

	if len(dupJSONKeys) == 0 && len(emptyFields) == 0 {
		return true
	}

	filePath := file.Desc.Path()
	msgName := string(msg.Desc.FullName())
	if index := strings.Index(msgName, "."); index > 0 {
		msgName = msgName[index+1:]
	}
	id := fmt.Sprintf("[ERROR] - [plugin: gojson] - [ file: %s | message: %s ]", filePath, msgName)

	if len(emptyFields) != 0 {
		println(fmt.Sprintf("%s - The json key is empty in fields %v", id, emptyFields))
	}

	for _, jsonKey := range dupJSONKeys {
		dupFields := dupFieldMap[jsonKey]
		println(fmt.Sprintf("%s - The json key <%s> are duplicated both in fields %v", id, jsonKey, dupFields))
	}
	return false
}

func fieldProtoName(fieldSet *FieldSet) string {
	var name string
	if fieldSet.IsOneOfField() {
		name = string(fieldSet.Field.Oneof.Desc.Name())
	} else {
		name = string(fieldSet.Field.Desc.Name())
	}
	if parent := fieldSet.ParentField(); parent != nil {
		name = fieldProtoName(parent) + "." + name
	}
	return name
}

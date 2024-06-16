package gojson

import (
	"os"
	"strconv"
	"strings"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"github.com/yu31/protoc-kit-go/helper/pkwkt"
	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-plugin-json/cmd/internal/gojson/pkg/importpkg"
	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

type Marshal struct {
	g    *protogen.GeneratedFile
	file *protogen.File

	// The message of currently being processed.
	message *protogen.Message

	// The message options of currently being processed.
	options *pbjson.MessageOptions
}

// FIXME: calculate the inline field.
func (ml *Marshal) guessBufLen(fieldSets []*FieldSet) (bufLen int) {
	// Sum json beginning(`{`) and closing(`}`).
	bufLen += 2

	for _, fieldSet := range fieldSets {
		if fieldSet.IsOneOfField() {
			for _, oneSet := range fieldSet.OneOfSet.Parts {
				// Sum key length, colon separator(`:`) and comma separator(`,`).
				bufLen += len(oneSet.JSONKey) + 2
			}
		}
		// Sum key length, colon separator(`:`) and comma separator(`,`).
		bufLen += len(fieldSet.JSONKey) + 2
	}
	bufLen = truncateBufLen(bufLen * 2)
	return
}

func (ml *Marshal) GenerateCode(fieldSets []*FieldSet) {
	msg := ml.message
	bufLen := ml.guessBufLen(fieldSets)

	ml.g.P("// MarshalJSON implements interface json.Marshaler for proto message ", msg.Desc.Name(), " in file ", ml.file.Desc.Path())
	ml.g.P("func (x *", msg.GoIdent.GoName, ") MarshalJSON() ([]byte, error) {")
	{
		ml.g.P("if x == nil {")
		ml.g.P(`	return []byte("null"), nil`)
		ml.g.P("}")
		if len(fieldSets) == 0 {
			ml.g.P(`return []byte("{}"), nil`)
		} else {
			ml.g.P("enc := ", importpkg.PJEncoder.Ident("New"), "(", bufLen, ")")
			ml.g.P("enc.AppendObjectBegin() // Add begin JSON identifier")
			ml.g.P("")
			for _, fieldSet := range fieldSets {
				catch := pkerror.Recover("gojson", ml.file, ml.message, fieldSet.Field, func() {
					ml.marshalField(fieldSet)
				})
				if catch {
					os.Exit(1)
				}
			}
			ml.g.P("enc.AppendObjectEnd() // Add end JSON identifier")
			ml.g.P("return enc.Bytes(), nil")
		}
	}
	// End of MarshalJSON.
	ml.g.P("}")
}

func (ml *Marshal) marshalField(fieldSet *FieldSet) {
	if fieldSet.FieldIsInline() && len(fieldSet.InlineSet.Childs) != 0 {
		// Confirm the parent field is not null
		fieldAlias := fieldSet.Alias()
		fieldValue := fieldSet.Value()
		ml.g.P("if ", fieldAlias, " := ", fieldValue, "; ", fieldAlias, " != nil {")
		for _, childField := range fieldSet.InlineSet.Childs {
			ml.marshalField(childField)
		}
		ml.g.P("}")
		return
	}

	switch {
	case fieldSet.IsOneOfField():
		ml.marshalOneOf(fieldSet)
	default:
		ml.encodeField(fieldSet)
	}
}

func (ml *Marshal) marshalOneOf(fieldSet *FieldSet) {
	if fieldSet.OneOfIsInline() {
		ml.marshalParts(fieldSet)
		return
	}

	fieldValue := fieldSet.Value()
	if options := fieldSet.OneOfSet.Options; options.Omitempty {
		ml.g.P("if ", fieldValue, " != nil {")
		{
			ml.g.P("enc.AppendObjectKey(", strconv.Quote(fieldSet.JSONKey), ")")
			ml.g.P("enc.AppendObjectBegin()")
			ml.marshalParts(fieldSet)
			ml.g.P("enc.AppendObjectEnd()")
		}
		ml.g.P("}")
	} else {
		ml.g.P("enc.AppendObjectKey(", strconv.Quote(fieldSet.JSONKey), ")")
		ml.g.P("if ", fieldValue, " != nil {") // // Confirm the oneof field is not null
		{
			ml.g.P("enc.AppendObjectBegin()")
			ml.marshalParts(fieldSet)
			ml.g.P("enc.AppendObjectEnd()")
		}
		ml.g.P("} else {")
		{
			ml.g.P("enc.AppendValNULL()")
		}
		ml.g.P("}")
	}
}
func (ml *Marshal) marshalParts(fieldSet *FieldSet) {
	partFields := fieldSet.OneOfSet.Parts
	if len(partFields) == 0 {
		return
	}

	fieldAlias := fieldSet.Alias()
	fieldValue := fieldSet.Value()
	ml.g.P("switch ", fieldAlias, " := ", fieldValue, ".(type) {")
	for _, partField := range partFields {
		partType := ml.g.QualifiedGoIdent(partField.Field.GoIdent)
		ml.g.P("case *", partType, ":")
		ml.marshalField(partField)
	}
	//ml.g.P("default:")
	//// To avoids unused panics. Cases occur in the all parts in oneof field are ignored.
	//ml.g.P("  _ = ", fieldAlias)
	ml.g.P("}")
}

func (ml *Marshal) encodeField(fieldSet *FieldSet) {
	var (
		funcName   string
		checkError bool
	)

	options := fieldSet.Options
	funcVars := []string{
		"enc",
		strconv.Quote(fieldSet.JSONKey),
		fieldSet.Value(),
		strconv.FormatBool(options.Omitempty),
	}

	switch {
	case fieldSet.Field.Desc.IsMap():
		formatMap := loadTypeFormatMap(options.Format)
		_mkName, _mkVars := loadMapKeyTypeInfo(fieldSet.Field.Message.Fields[0], formatMap.Key)
		_mvName, _mvVars := loadValueTypeInfo(fieldSet.Field.Message.Fields[1], formatMap.Value)
		checkError = ml.funcReturnWithError(fieldSet.Field.Message.Fields[1], formatMap.Value)

		funcName = "AppendMap" + _mkName + _mvName
		funcVars = append(funcVars, _mkVars...)
		funcVars = append(funcVars, _mvVars...)
	case fieldSet.Field.Desc.IsList():
		formatRepeated := loadTypeFormatRepeated(options.Format)
		_typeName, _vars := loadValueTypeInfo(fieldSet.Field, formatRepeated.Elem)
		checkError = ml.funcReturnWithError(fieldSet.Field, formatRepeated.Elem)

		funcName = "AppendList" + _typeName
		funcVars = append(funcVars, _vars...)
	default:
		_typeName, _vars := loadValueTypeInfo(fieldSet.Field, options.Format)
		checkError = ml.funcReturnWithError(fieldSet.Field, options.Format)

		if pkfield.FieldIsOptional(fieldSet.Field) {
			funcName = "AppendPtr" + _typeName
		} else {
			funcName = "AppendVal" + _typeName
		}
		funcVars = append(funcVars, _vars...)
	}

	funcIdent := importpkg.PJEncoder.Ident(funcName)
	parameters := strings.Join(funcVars, ",")
	if checkError {
		ml.g.P("if err := ", funcIdent, "(", parameters, "); err != nil {")
		ml.g.P("    return nil, err")
		ml.g.P("}")
	} else {
		ml.g.P(funcIdent, "(", parameters, ")")
	}
}
func (ml *Marshal) funcReturnWithError(field *protogen.Field, typeFormat *pbjson.TypeFormat) (yes bool) {
	kind := field.Desc.Kind()
	switch kind {
	case protoreflect.BytesKind:
		return true
	case protoreflect.MessageKind:
		switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
		case pkwkt.Any:
			typeAny := loadTypeFormatAny(typeFormat)
			switch typeAny.Codec {
			case pbjson.TypeAny_Unset, pbjson.TypeAny_Object, pbjson.TypeAny_Proto:
				return true
			}
		case pkwkt.Duration:
			typeDuration := loadTypeFormatDuration(typeFormat)
			switch typeDuration.Codec {
			case pbjson.TypeDuration_Unset, pbjson.TypeDuration_Object:
				return true
			}
		case pkwkt.Timestamp:
			typeTimestamp := loadTypeFormatTimestamp(typeFormat)
			switch typeTimestamp.Codec {
			case pbjson.TypeTimestamp_Unset, pbjson.TypeTimestamp_Object:
				return true
			}
		default:
			return true
		}
	}
	return false
}

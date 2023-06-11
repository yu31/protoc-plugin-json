package gojson

import (
	"reflect"
	"strings"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

func (p *Plugin) getTagName(ot interface{}) string {
	name := reflect.TypeOf(ot).Elem().Name()
	name = strings.TrimSuffix(name, "_") // string is String_
	parts := strings.Split(name, "_")
	return strings.ToLower(parts[1])
}

// The SerializeOptions priority from low to high is: file_options -> msg_options
func (p *Plugin) loadMessageOptions(msg *protogen.Message) *pbjson.MessageOptions {
	i := proto.GetExtension(msg.Desc.Options(), pbjson.E_Message)
	msgOptions := i.(*pbjson.MessageOptions)
	if msgOptions == nil {
		msgOptions = &pbjson.MessageOptions{}
	}
	return msgOptions
}
func (p *Plugin) loadFieldOptions(field *protogen.Field) *pbjson.FieldOptions {
	i := proto.GetExtension(field.Desc.Options(), pbjson.E_Field)
	fieldOptions := i.(*pbjson.FieldOptions)
	if fieldOptions == nil {
		fieldOptions = &pbjson.FieldOptions{}
	}
	// Ignore field if json == "-"
	if fieldOptions.Json != nil && *fieldOptions.Json == "-" {
		fieldOptions.Ignore = true
	}
	return fieldOptions
}

func (p *Plugin) loadOneofOptions(oneof *protogen.Oneof) *pbjson.OneofOptions {
	i := proto.GetExtension(oneof.Desc.Options(), pbjson.E_Oneof)
	oneofOptions := i.(*pbjson.OneofOptions)
	if oneofOptions == nil {
		oneofOptions = &pbjson.OneofOptions{}
	}
	// Ignore field if json == "-"
	if oneofOptions.Json != nil && *oneofOptions.Json == "-" {
		oneofOptions.Ignore = true
	}
	return oneofOptions
}

//func (p *Plugin) loadTypeSetEnum(options *pbjson.FieldOptions) *pbjson.TypeEnum {
//	if options.TypeSet == nil {
//		return &pbjson.TypeEnum{Format: nil}
//	}
//
//	switch ot := options.TypeSet.(type) {
//	case *pbjson.FieldOptions_Enum:
//		return ot.Enum
//	default:
//		err := pkerror.New(
//			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
//			p.getTagName(ot),
//		)
//		panic(err)
//	}
//}
func (p *Plugin) loadTypeSetAny(options *pbjson.FieldOptions) *pbjson.TypeAny {
	if options.TypeSet == nil {
		return &pbjson.TypeAny{Format: nil}
	}

	switch ot := options.TypeSet.(type) {
	case *pbjson.FieldOptions_Any:
		return ot.Any
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadTypeSetDuration(options *pbjson.FieldOptions) *pbjson.TypeDuration {
	if options.TypeSet == nil {
		return &pbjson.TypeDuration{Format: nil}
	}

	switch ot := options.TypeSet.(type) {
	case *pbjson.FieldOptions_Duration:
		return ot.Duration
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadTypeSetTimestamp(options *pbjson.FieldOptions) *pbjson.TypeTimestamp {
	if options.TypeSet == nil {
		return &pbjson.TypeTimestamp{Format: nil}
	}

	switch ot := options.TypeSet.(type) {
	case *pbjson.FieldOptions_Timestamp:
		return ot.Timestamp
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

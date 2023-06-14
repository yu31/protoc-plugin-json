package gojson

import (
	"reflect"
	"strings"
	"time"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

func getTagName(ot interface{}) string {
	name := reflect.TypeOf(ot).Elem().Name()
	name = strings.TrimSuffix(name, "_") // string is String_
	parts := strings.Split(name, "_")
	return strings.ToLower(parts[1])
}

func loadMessageOptions(msg *protogen.Message) *pbjson.MessageOptions {
	i := proto.GetExtension(msg.Desc.Options(), pbjson.E_Message)
	options := i.(*pbjson.MessageOptions)
	if options == nil {
		options = &pbjson.MessageOptions{}
	}
	return options
}
func loadFieldOptions(field *protogen.Field) *pbjson.FieldOptions {
	i := proto.GetExtension(field.Desc.Options(), pbjson.E_Field)
	options := i.(*pbjson.FieldOptions)
	if options == nil {
		options = &pbjson.FieldOptions{}
	}
	// Ignore field if json is "-"
	if options.Json != nil && *options.Json == "-" {
		options.Ignore = true
	}
	return options
}

func loadOneOfOptions(field *protogen.Field) *pbjson.OneofOptions {
	i := proto.GetExtension(field.Oneof.Desc.Options(), pbjson.E_Oneof)
	options := i.(*pbjson.OneofOptions)
	if options == nil {
		options = &pbjson.OneofOptions{}
	}
	// Ignore field if json is "-"
	if options.Json != nil && *options.Json == "-" {
		options.Ignore = true
	}
	return options
}

func loadTypeSetEnum(options *pbjson.FieldOptions) *pbjson.TypeEnum {
	if options.TypeSet == nil {
		return &pbjson.TypeEnum{Format: pbjson.TypeEnum_Number}
	}

	switch ot := options.TypeSet.(type) {
	case *pbjson.FieldOptions_Enum:
		return ot.Enum
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeSetAny(options *pbjson.FieldOptions) *pbjson.TypeAny {
	if options.TypeSet == nil {
		return &pbjson.TypeAny{Format: pbjson.TypeAny_Native}
	}

	switch ot := options.TypeSet.(type) {
	case *pbjson.FieldOptions_Any:
		return ot.Any
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeSetDuration(options *pbjson.FieldOptions) *pbjson.TypeDuration {
	if options.TypeSet == nil {
		return &pbjson.TypeDuration{Format: pbjson.TypeDuration_Native}
	}

	switch ot := options.TypeSet.(type) {
	case *pbjson.FieldOptions_Duration:
		return ot.Duration
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}
func loadTypeSetTimestamp(options *pbjson.FieldOptions) *pbjson.TypeTimestamp {
	if options.TypeSet == nil {
		return &pbjson.TypeTimestamp{Format: pbjson.TypeTimestamp_Native}
	}

	switch ot := options.TypeSet.(type) {
	case *pbjson.FieldOptions_Timestamp:
		rt := ot.Timestamp
		if rt.Layout == nil {
			rt.Layout = &pbjson.TypeTimestamp_Layout{}
		}
		if rt.Layout.Golang == "" {
			rt.Layout.Golang = time.RFC3339Nano
		}
		return rt
	default:
		err := pkerror.New(
			"type google.protobuf.Any only supports kind of TypeSet <any> and you provided: <%s>",
			getTagName(ot),
		)
		panic(err)
	}
}

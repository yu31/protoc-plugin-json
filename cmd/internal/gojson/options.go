package gojson

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
)

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

package utils

import "google.golang.org/protobuf/encoding/protojson"

var PMarshal = &protojson.MarshalOptions{
	Multiline:       false,
	Indent:          "",
	AllowPartial:    true,
	UseProtoNames:   true,
	UseEnumNumbers:  true,
	EmitUnpopulated: true,
}

var PUnmarshal = &protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: false,
	Resolver:       nil,
}

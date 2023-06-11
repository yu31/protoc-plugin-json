package jsonencoder

import "google.golang.org/protobuf/encoding/protojson"

var pMarshal = &protojson.MarshalOptions{
	Multiline:       false,
	Indent:          "",
	AllowPartial:    true,
	UseProtoNames:   true,
	UseEnumNumbers:  true,
	EmitUnpopulated: true,
}

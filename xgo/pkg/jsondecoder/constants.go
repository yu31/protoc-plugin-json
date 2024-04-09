package jsondecoder

import (
	"google.golang.org/protobuf/encoding/protojson"
)

var pUnmarshal = &protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: false,
	Resolver:       nil,
}

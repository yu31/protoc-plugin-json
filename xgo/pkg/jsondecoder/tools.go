package jsondecoder

import (
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/protobuf/encoding/protojson"
)

var pUnmarshal = &protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: false,
	Resolver:       nil,
}

func durationToPB(dd time.Duration, vv *duration.Duration) {
	nanos := dd.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9

	vv.Seconds = int64(secs)
	vv.Nanos = int32(nanos)
}

func assertInterface(vv interface{}) {
	if vv == nil {
		panic("jsondecoder: the interface is nil")
	}
}

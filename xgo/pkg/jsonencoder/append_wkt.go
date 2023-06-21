package jsonencoder

import (
	"github.com/golang/protobuf/ptypes/any"
)

func (enc *Encoder) AppendWKTAnyByProto(v *any.Any) error {
	b, err := pMarshal.Marshal(v)
	if err != nil {
		return err
	}
	enc.appendCommaSeparator()
	enc.writeBytes(b)
	return nil
}

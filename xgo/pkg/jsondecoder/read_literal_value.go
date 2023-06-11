package jsondecoder

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (dec *Decoder) ReadValueString(jsonKey string) (vv string, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueBool(jsonKey string) (vv bool, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToBool(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInt32(jsonKey string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToInt32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInt64(jsonKey string) (vv int64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueUint32(jsonKey string) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToUint32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueUint64(jsonKey string) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToUint64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(value), jsonKey)
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadValueFloat32(jsonKey string) (vv float32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueFloat64(jsonKey string) (vv float64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueEnumString(jsonKey string, em map[string]int32) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = parseEnumString(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueEnumNumber(jsonKey string, em map[int32]string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = parseEnumNumber(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as number into field %s of type Enum", string(value), jsonKey)
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadValueBytes(jsonKey string) (vv []byte, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	if vv, err = bytesToBytes(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type []byte", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInterface(jsonKey string, initFN func() interface{}) (err error) {
	_ = jsonKey

	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil
	}
	if err = parseInterface(value, initFN); err != nil {
		return
	}
	return
}

func (dec *Decoder) ReadValueWKTAnyExpand(jsonKey string, vv *any.Any) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return
	}
	if err = pUnmarshal.Unmarshal(value, vv); err != nil {
		return
	}
	return
}

func (dec *Decoder) ReadValueWKTDurationString(jsonKey string, vv *durationpb.Duration) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ss string
	if ss, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(value), jsonKey)
		return
	}
	if ss == "" { // Ignore empty string.
		return err
	}
	var dd time.Duration
	if dd, err = time.ParseDuration(ss); err != nil {
		return
	}
	durationToPB(dd, vv)
	return nil
}
func (dec *Decoder) ReadValueWKTDurationNanoseconds(jsonKey string, vv *durationpb.Duration) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	durationToPB(time.Duration(ii), vv)
	return nil
}
func (dec *Decoder) ReadValueWKTDurationMicroseconds(jsonKey string, vv *durationpb.Duration) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	durationToPB(time.Duration(ii)*time.Microsecond, vv)
	return nil
}
func (dec *Decoder) ReadValueWKTDurationMilliseconds(jsonKey string, vv *durationpb.Duration) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	durationToPB(time.Duration(ii)*time.Millisecond, vv)
	return nil
}

func (dec *Decoder) ReadValueWKTDurationSeconds(jsonKey string, vv *durationpb.Duration) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	durationToPB(time.Duration(ff*float64(time.Second)), vv)
	return nil
}
func (dec *Decoder) ReadValueWKTDurationMinutes(jsonKey string, vv *durationpb.Duration) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	durationToPB(time.Duration(ff*float64(time.Minute)), vv)
	return nil
}
func (dec *Decoder) ReadValueWKTDurationHours(jsonKey string, vv *durationpb.Duration) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	durationToPB(time.Duration(ff*float64(time.Hour)), vv)
	return nil
}

func (dec *Decoder) ReadValueWKTTimestampString(jsonKey string, vv *timestamppb.Timestamp, layout string) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ss string
	if ss, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(value), jsonKey)
		return
	}

	if ss == "" { // Ignore empty string.
		return err
	}

	var tt time.Time
	if tt, err = time.Parse(layout, ss); err != nil {
		return
	}
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}
func (dec *Decoder) ReadValueWKTTimestampUnixNano(jsonKey string, vv *timestamppb.Timestamp) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}

	tt := time.Unix(0, ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}
func (dec *Decoder) ReadValueWKTTimestampUnixMicro(jsonKey string, vv *timestamppb.Timestamp) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}

	tt := time.UnixMicro(ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}
func (dec *Decoder) ReadValueWKTTimestampUnixMilli(jsonKey string, vv *timestamppb.Timestamp) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}

	tt := time.UnixMilli(ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}
func (dec *Decoder) ReadValueWKTTimestampUnixSec(jsonKey string, vv *timestamppb.Timestamp) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}

	tt := time.Unix(ii, 0)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}

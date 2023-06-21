package jsondecoder

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (dec *Decoder) ReadWKTAnyByProto(jsonKey string, vv *any.Any) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
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

func (dec *Decoder) ReadWKTDurationByString(jsonKey string, vv *durationpb.Duration) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ss string
	if ss, err = dec.convertToString(jsonKey, value); err != nil {
		return
	}
	if ss == "" { // Empty string represents empty Duration.
		vv.Seconds = 0
		vv.Nanos = 0
		return err
	}
	var dd time.Duration
	if dd, err = time.ParseDuration(ss); err != nil {
		return
	}
	durationToPB(dd, vv)
	return nil
}
func (dec *Decoder) ReadWKTDurationByNanoseconds(jsonKey string, vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return err
	}
	durationToPB(time.Duration(ii), vv)
	return nil
}
func (dec *Decoder) ReadWKTDurationByMicroseconds(jsonKey string, vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return err
	}
	durationToPB(time.Duration(ii)*time.Microsecond, vv)
	return nil
}
func (dec *Decoder) ReadWKTDurationByMilliseconds(jsonKey string, vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return err
	}
	durationToPB(time.Duration(ii)*time.Millisecond, vv)
	return nil
}

func (dec *Decoder) ReadWKTDurationBySeconds(jsonKey string, vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = dec.convertToFloat64(jsonKey, value, unquote); err != nil {
		return err
	}
	durationToPB(time.Duration(ff*float64(time.Second)), vv)
	return nil
}
func (dec *Decoder) ReadWKTDurationByMinutes(jsonKey string, vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = dec.convertToFloat64(jsonKey, value, unquote); err != nil {
		return err
	}
	durationToPB(time.Duration(ff*float64(time.Minute)), vv)
	return nil
}
func (dec *Decoder) ReadWKTDurationByHours(jsonKey string, vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = dec.convertToFloat64(jsonKey, value, unquote); err != nil {
		return err
	}
	durationToPB(time.Duration(ff*float64(time.Hour)), vv)
	return nil
}

func (dec *Decoder) ReadWKTTimestampByString(jsonKey string, vv *timestamppb.Timestamp, layout string) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
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
func (dec *Decoder) ReadWKTTimestampByUnixNano(jsonKey string, vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return err
	}

	tt := time.Unix(0, ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}
func (dec *Decoder) ReadWKTTimestampByUnixMicro(jsonKey string, vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return err
	}

	tt := time.UnixMicro(ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}
func (dec *Decoder) ReadWKTTimestampByUnixMilli(jsonKey string, vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return err
	}

	tt := time.UnixMilli(ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}
func (dec *Decoder) ReadWKTTimestampByUnixSec(jsonKey string, vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return err
	}

	tt := time.Unix(ii, 0)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}

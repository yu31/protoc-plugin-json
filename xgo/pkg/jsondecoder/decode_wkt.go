package jsondecoder

import (
	"time"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func durationToPB(dd time.Duration, vv *duration.Duration) {
	nanos := dd.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9

	vv.Seconds = secs
	vv.Nanos = int32(nanos)
}

func (dec *Decoder) readValWKTAnyProto(vv *any.Any) (err error) {
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

// readValWKTDurTimeStr read the value of WKT duration by time string.
func (dec *Decoder) readValWKTDurTimeStr(vv *durationpb.Duration) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ss string
	if ss, err = convertToString(value); err != nil {
		return
	}
	if ss == "" { // Empty string represents empty Duration.
		vv.Seconds = 0
		vv.Nanos = 0
		return nil
	}
	var dd time.Duration
	if dd, err = time.ParseDuration(ss); err != nil {
		return
	}

	durationToPB(dd, vv)
	return nil
}

// readValWKTDurNano read the value of WKT duration with nanoseconds.
func (dec *Decoder) readValWKTDurNano(vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = convertToInt64(unquote, value); err != nil {
		return err
	}

	durationToPB(time.Duration(ii), vv)
	return nil
}

// readValWKTDurMicro read the value of WKT duration with microseconds.
func (dec *Decoder) readValWKTDurMicro(vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = convertToInt64(unquote, value); err != nil {
		return err
	}

	durationToPB(time.Duration(ii)*time.Microsecond, vv)
	return nil
}

// readValWKTDurMilli read the value of WKT duration with milliseconds.
func (dec *Decoder) readValWKTDurMilli(vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = convertToInt64(unquote, value); err != nil {
		return err
	}

	durationToPB(time.Duration(ii)*time.Millisecond, vv)
	return nil
}

// readValWKTDurSecond read the value of WKT duration with seconds.
func (dec *Decoder) readValWKTDurSecond(vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = convertToFloat64(unquote, value); err != nil {
		return err
	}

	durationToPB(time.Duration(ff*float64(time.Second)), vv)
	return nil
}

// readValWKTDurMinute read the value of WKT duration with minutes.
func (dec *Decoder) readValWKTDurMinute(vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = convertToFloat64(unquote, value); err != nil {
		return err
	}

	durationToPB(time.Duration(ff*float64(time.Minute)), vv)
	return nil
}

// readValWKTDurHour read the value of WKT duration with hours.
func (dec *Decoder) readValWKTDurHour(vv *durationpb.Duration, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ff float64
	if ff, err = convertToFloat64(unquote, value); err != nil {
		return err
	}

	durationToPB(time.Duration(ff*float64(time.Hour)), vv)
	return nil
}

// readValWKTTsLayout read the value of WKT timestamp with time string.
func (dec *Decoder) readValWKTTsLayout(vv *timestamppb.Timestamp, layout string) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ss string
	if ss, err = convertToString(value); err != nil {
		return
	}
	if ss == "" { // Ignore empty string.
		vv.Seconds = 0
		vv.Nanos = 0
		return nil
	}

	var tt time.Time
	if tt, err = time.Parse(layout, ss); err != nil {
		return
	}
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}

// readValWKTTsUnixNano read the value of WKT timestamp with unix nanoseconds.
func (dec *Decoder) readValWKTTsUnixNano(vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = convertToInt64(unquote, value); err != nil {
		return err
	}

	tt := time.Unix(0, ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}

// readValWKTTsUnixMicro read the value of WKT timestamp with unix microseconds.
func (dec *Decoder) readValWKTTsUnixMicro(vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = convertToInt64(unquote, value); err != nil {
		return err
	}

	tt := time.UnixMicro(ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}

// readValWKTTsUnixMilli read the value of WKT timestamp with unix milliseconds.
func (dec *Decoder) readValWKTTsUnixMilli(vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = convertToInt64(unquote, value); err != nil {
		return err
	}

	tt := time.UnixMilli(ii)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}

// readValWKTTsUnixSec read the value of WKT timestamp with unix seconds.
func (dec *Decoder) readValWKTTsUnixSec(vv *timestamppb.Timestamp, unquote bool) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var ii int64
	if ii, err = convertToInt64(unquote, value); err != nil {
		return err
	}

	tt := time.Unix(ii, 0)
	vv.Seconds = tt.Unix()
	vv.Nanos = int32(tt.Nanosecond())
	return nil
}

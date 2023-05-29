package benchmark

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbbenchmark"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

type CopySimple pbbenchmark.BenchModelSimple

var seedSimple = &pbbenchmark.BenchModelSimple{
	FString1:  "s11111",
	FString2:  "s11112",
	FString3:  "s11113",
	FString4:  "s11114",
	FString5:  "s11115",
	FString6:  "s11116",
	FString7:  "s11117",
	FString8:  "s11118",
	FString9:  "s11119",
	FInt32:    11111,
	FInt64:    11112,
	FUint32:   11113,
	FUint64:   11114,
	FSint32:   11115,
	FSint64:   11116,
	FSfixed32: 11117,
	FSfixed64: 11118,
	FFixed32:  11119,
	FFixed64:  11120,
	FFloat:    111.111,
	FDouble:   112.112,
	FBool1:    true,
	FBool2:    true,
	FBool3:    true,
}

var copySimple = &CopySimple{
	FString1:  "s11111",
	FString2:  "s11112",
	FString3:  "s11113",
	FString4:  "s11114",
	FString5:  "s11115",
	FString6:  "s11116",
	FString7:  "s11117",
	FString8:  "s11118",
	FString9:  "s11119",
	FInt32:    11111,
	FInt64:    11112,
	FUint32:   11113,
	FUint64:   11114,
	FSint32:   11115,
	FSint64:   11116,
	FSfixed32: 11117,
	FSfixed64: 11118,
	FFixed32:  11119,
	FFixed64:  11120,
	FFloat:    111.111,
	FDouble:   112.112,
	FBool1:    true,
	FBool2:    true,
	FBool3:    true,
}

func Benchmark_Simple_Marshal_Plugin(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := seedSimple.MarshalJSON()
			if err != nil {
				b.Fatal("simple marshal by plugin error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Simple_Marshal_Proto(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := utils.PMarshal.Marshal(seedSimple)
			if err != nil {
				b.Fatal("simple marshal by proto error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Simple_Marshal_StdJSON(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := json.Marshal(copySimple)
			if err != nil {
				b.Fatal("simple marshal by standard json error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Simple_Marshal_JSONIter(b *testing.B) {
	_json := jsoniter.ConfigCompatibleWithStandardLibrary

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := _json.Marshal(&copySimple)
			if err != nil {
				b.Fatal("simple marshal by json iterator error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Simple_Unmarshal_Plugin(b *testing.B) {
	bb, err := seedSimple.MarshalJSON()
	if err != nil {
		b.Fatal("simple marshal by plugin error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &pbbenchmark.BenchModelSimple{}
			err = dataNew.UnmarshalJSON(bb)
			if err != nil {
				b.Fatal("simple unmarshal by plugin error:", err)
			}
		}
	})
}

func Benchmark_Simple_Unmarshal_Proto(b *testing.B) {
	bb, err := utils.PMarshal.Marshal(seedSimple)
	if err != nil {
		b.Fatal("simple marshal by proto error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &pbbenchmark.BenchModelSimple{}
			err = utils.PUnmarshal.Unmarshal(bb, dataNew)
			if err != nil {
				b.Fatal("simple unmarshal by proto error:", err)
			}
		}
	})
}

func Benchmark_Simple_Unmarshal_StdJSON(b *testing.B) {
	bb, err := json.Marshal(copySimple)
	if err != nil {
		b.Fatal("simple marshal by standard json error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &CopySimple{}
			err = json.Unmarshal(bb, dataNew)
			if err != nil {
				b.Fatal("simple unmarshal by standard json error:", err)
			}
		}
	})
}

func Benchmark_Simple_Unmarshal_JSONIter(b *testing.B) {
	_json := jsoniter.ConfigCompatibleWithStandardLibrary

	bb, err := _json.Marshal(&copySimple)
	if err != nil {
		b.Fatal("simple marshal by json iterator error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &CopySimple{}
			err = _json.Unmarshal(bb, dataNew)
			if err != nil {
				b.Fatal("simple unmarshal by json iterator error:", err)
			}
		}
	})
}

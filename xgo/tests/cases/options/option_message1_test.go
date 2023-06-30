package options

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pboptions"
)

func Test_IgnoreMessage1(t *testing.T) {
	data := &pboptions.IgnoreMessage1{
		FString1: "ss101",
		FString2: "ss102",
		FString3: "ss103",
	}
	_, ok1 := interface{}(data).(json.Marshaler)
	require.False(t, ok1)
	_, ok2 := interface{}(data).(json.Unmarshaler)
	require.False(t, ok2)
}

func Test_DisallowUnknownFields(t *testing.T) {
	b1 := []byte(`{
		"f_string1": "ss101",
		"f_string2": "ss102",
		"f_string3": "ss103"
	}`)

	b2 := []byte(`{
		"f_string1": "ss101",
		"f_string2": "ss102",
		"f_string3": "ss103",
		"f_string4": "ss104"
	}`)

	var err error

	t.Run("DisallowUnknownMessage1", func(t *testing.T) {
		data1 := pboptions.DisallowUnknownMessage1{}
		err = data1.UnmarshalJSON(b1)
		require.Nil(t, err)

		data2 := pboptions.DisallowUnknownMessage1{}
		err = data2.UnmarshalJSON(b2)
		require.NotNil(t, err)
		fmt.Println("Test_DisallowUnknownFields:", err.Error())
	})

	t.Run("DisallowUnknownMessage2", func(t *testing.T) {
		data1 := pboptions.DisallowUnknownMessage2{}
		err = data1.UnmarshalJSON(b1)
		require.Nil(t, err)

		data2 := pboptions.DisallowUnknownMessage2{}
		err = data2.UnmarshalJSON(b2)
		require.Nil(t, err)
	})
}

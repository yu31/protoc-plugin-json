package jsondecoder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Object_IsNULL(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		bb := []byte(`null`)
		decoder, err := New(bb)
		require.Nil(t, err)

		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
	t.Run("case2", func(t *testing.T) {
		bb := []byte(`  null  `)
		decoder, err := New(bb)
		require.Nil(t, err)

		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
}
func Test_Object_IsEmpty(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		bb := []byte(`{}`)
		decoder, err := New(bb)
		require.Nil(t, err)

		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)
	})
	t.Run("case2", func(t *testing.T) {
		bb := []byte(`  {   }  `)
		decoder, err := New(bb)
		require.Nil(t, err)

		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)
	})
}

func Test_Object_Read1(t *testing.T) {
	bb := []byte(
		`
{
	"f_int1":101,"f_int2"  :102,  "f_int3"  :  103  ,"f_int4"  :  104  ,
	"f_int5"  :  105  ,
	"r_int1":[11,12,13],
	"m_int1":{"k11":11,"k12":12,"k13":13},

	"f_string1"   :   "s101"    ,
	"m_int2"  :  {  "k21"  :  21  ,  "k22"  :22  ,"k23":  23    }    ,
	"r_int2"  :  [  21  ,   22  ,  23  ]  ,

	"f_message1":  {
		"f_string1": "s11",
		"f_string2": "s12",
		"f_string3": "s13"
	},
	"f_message2"  :{"f_string1":"s21"  ,  "f_string1":  "s22","f_string1"  : "s23"}     ,
	"f_message3"  :  {   "f_string1": "s31","f_string1": "s32","f_string1": "s33"}     ,
	"m_int3":  {
		"k31": 31,
		"k32": 32  ,
		"k33": 33
	}    ,
	"r_int3":  [
		31  ,
		32,
		33
	],


	"r_message1": [
		211,
		212, 213,
		214,    215
	],
	"r_message2": [221,222,223,224,225],
	"r_message3": [  231  ,  232 , 233 , 234   , 235  ],
	"f_int6":106,"f_int7":107,
	"r_empty1":[],"r_empty2"   :[  ]   ,"r_empty3":  [  ]  ,"r_empty4"  :  [  ],
	"m_empty1":{},"m_empty2"   :{   }   ,"m_empty3":  {  }  ,"m_empty4"  :  {  },
	"f_uint1":301,"f_uint2":302,"f_uint3":303,"f_uint4":304,
	"r_null1":null,"r_null2"  :null  ,"r_null3":  null  ,"r_null4"  :  null,
	"m_null1":null,"m_null2"  :null  ,"m_null3":  null  ,"m_null4"  :  null
}
`,
	)

	decoder, err := New(bb)
	require.Nil(t, err)

	t.Run("prepare", func(t *testing.T) {
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)
	})
	t.Run("f_int1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_int1"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("101"), value)
	})
	t.Run("f_int2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_int2"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("102"), value)
	})
	t.Run("f_int3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_int3"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("103"), value)
	})
	t.Run("f_int4", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_int4"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("104"), value)
	})
	t.Run("f_int5", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_int5"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("105"), value)
	})
	t.Run("r_int1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_int1"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// Read element 1.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("11"), elem)

		// Read element 2.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("12"), elem)

		// Read element 3.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("13"), elem)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("m_int1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_int1"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// Read key/value 1.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k11"`), mapKey)
		mapVal, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`11`), mapVal)

		// Read element 2.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err = decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k12"`), mapKey)
		mapVal, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`12`), mapVal)

		// Read element 3.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err = decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k13"`), mapKey)
		mapVal, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`13`), mapVal)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("f_string1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_string1"`), objKey)

		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`"s101"`), value)
	})
	t.Run("m_int2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_int2"`), objKey)

		// prepare read map.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// Read key/value 1.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k21"`), mapKey)
		mapVal, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`21`), mapVal)

		// Read element 2.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err = decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k22"`), mapKey)
		mapVal, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`22`), mapVal)

		// Read element 3.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err = decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k23"`), mapKey)
		mapVal, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`23`), mapVal)

		// Is ended.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("r_int2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_int2"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// Read element 1.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("21"), elem)

		// Read element 2.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("22"), elem)

		// Read element 3.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("23"), elem)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("f_message1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_message1"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`{
		"f_string1": "s11",
		"f_string2": "s12",
		"f_string3": "s13"
	}`), value)
	})
	t.Run("f_message2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_message2"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`{"f_string1":"s21"  ,  "f_string1":  "s22","f_string1"  : "s23"}`), value)
	})
	t.Run("f_message3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_message3"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`{   "f_string1": "s31","f_string1": "s32","f_string1": "s33"}`), value)
	})
	t.Run("m_int3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_int3"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// Read key/value 1.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k31"`), mapKey)
		mapVal, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`31`), mapVal)

		// Read element 2.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err = decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k32"`), mapKey)
		mapVal, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`32`), mapVal)

		// Read element 3.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		mapKey, err = decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"k33"`), mapKey)
		mapVal, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`33`), mapVal)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("r_int3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_int3"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// Read element 1.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("31"), elem)

		// Read element 2.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("32"), elem)

		// Read element 3.
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)
		elem, err = decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("33"), elem)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("r_message1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_message1"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`[
		211,
		212, 213,
		214,    215
	]`), value)
	})
	t.Run("r_message2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_message2"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`[221,222,223,224,225]`), value)
	})
	t.Run("r_message3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_message3"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte(`[  231  ,  232 , 233 , 234   , 235  ]`), value)
	})
	t.Run("f_int6", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_int6"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("106"), value)
	})
	t.Run("f_int7", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_int7"`), objKey)
		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("107"), value)
	})

	t.Run("r_empty1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_empty1"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("r_empty2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_empty2"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("r_empty3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_empty3"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("r_empty4", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_empty4"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})

	t.Run("m_empty1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_empty1"`), objKey)

		// prepare read object.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("m_empty2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_empty2"`), objKey)

		// prepare read object.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("m_empty3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_empty3"`), objKey)

		// prepare read object.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})
	t.Run("m_empty4", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_empty4"`), objKey)

		// prepare read object.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		// Advance to next.
		err = decoder.scanNext()
		require.Nil(t, err)
	})

	t.Run("f_uint1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_uint1"`), objKey)

		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("301"), value)
	})
	t.Run("f_uint2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_uint2"`), objKey)

		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("302"), value)
	})
	t.Run("f_uint3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_uint3"`), objKey)

		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("303"), value)
	})
	t.Run("f_uint4", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"f_uint4"`), objKey)

		value, err := decoder.readLiteralValue()
		require.Nil(t, err)
		require.Equal(t, []byte("304"), value)
	})

	t.Run("r_null1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_null1"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
	t.Run("r_null2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_null2"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
	t.Run("r_null3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_null3"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
	t.Run("r_null4", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_null4"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})

	t.Run("m_null1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_null1"`), objKey)

		// prepare read object.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
	t.Run("m_null2", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_null2"`), objKey)

		// prepare read object.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
	t.Run("m_null3", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_null3"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})
	t.Run("m_null4", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"m_null4"`), objKey)

		// prepare read array.
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.True(t, isNULL)
	})

	t.Run("END", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)
	})
}

// Test only contains an empty array.
func Test_Object_Read2(t *testing.T) {
	bb := []byte(`       {    "r_string1"    : [      ]    }   `)

	decoder, err := New(bb)
	require.Nil(t, err)

	t.Run("prepare", func(t *testing.T) {
		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)
	})
	t.Run("r_string1", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.False(t, isEnd)

		objKey, err := decoder.readObjectKey()
		require.Nil(t, err)
		require.Equal(t, []byte(`"r_string1"`), objKey)

		isNULL, err := decoder.beforeReadObject()
		require.Nil(t, err)
		require.False(t, isNULL)

		// is ended
		isEnd, err = decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)

		err = decoder.scanNext()
		require.Nil(t, err)
	})

	t.Run("end", func(t *testing.T) {
		isEnd, err := decoder.beforeReadNext()
		require.Nil(t, err)
		require.True(t, isEnd)
	})
}

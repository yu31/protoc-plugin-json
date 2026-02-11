package inline

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbinline"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_InlineMessage01_EmptyValue(t *testing.T) {
	level01 := &pbinline.MessageLevel01{
		Level01FString1:                   "",
		Level01FString2:                   "",
		Level01PString1:                   nil,
		Level01PString2:                   nil,
		Level01RString1:                   nil,
		Level01RString2:                   nil,
		Level01MString1:                   nil,
		Level01MString2:                   nil,
		Level01FMessageExtern1:            nil,
		Level01FMessageExtern2:            nil,
		Level01FMessageInline02:           nil,
		Level01FMessageInline03:           nil,
		Level01OneOfExtern1:               nil,
		Level01OneOfInline1:               nil,
		Level01OneOfOmitempty1:            nil,
		Level01FMessageInlineEmtpy:        nil,
		Level01FMessageInlineIgnoreByRef:  nil,
		Level01FMessageInlineIgnoreFields: nil,
		Level01OneOfIgnoreSelf1:           nil,
		Level01OneOfIgnoreParts1:          nil,
	}

	var (
		err error
		bb  []byte
	)

	t.Run("Marshal", func(t *testing.T) {
		bb, err = level01.MarshalJSON()
		require.Nil(t, err)
		require.True(t, json.Valid(bb))
	})

	t.Run("Unmarshal", func(t *testing.T) {
		dataNew := &pbinline.MessageLevel01{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, level01, dataNew)
	})
}

/*
Test cases include:
 1. Include multi-level inline message.
 2. Include inline oneof and include multi-level inline message in oneof parts.
 3. Include options `omitempty` and `ignore`.
 4. Include the field type message(option inline) and has no fields.
 5. Include the field type message(option inline) and all fields are ignore.
 6. Include the field type message(option inline) and option ignore.
 7. Include the field type oneof(option inline) and option ignore.
 8. Include the field type oneof(option inline) and all parts are ignore.
*/
func Test_InlineMessage01_Complex(t *testing.T) {
	var ( // message in level01
		level02N01 *pbinline.MessageLevel02
		level02N02 *pbinline.MessageLevel02

		level03N01 *pbinline.MessageLevel03

		level04N01 *pbinline.MessageLevel04

		level06N01 *pbinline.MessageLevel06
		level06N02 *pbinline.MessageLevel06
		level06N03 *pbinline.MessageLevel06
		level06N04 *pbinline.MessageLevel06

		level07N01 *pbinline.MessageLevel07
		level07N02 *pbinline.MessageLevel07

		level08N01 *pbinline.MessageLevel08
		level08N02 *pbinline.MessageLevel08

		level10N01 *pbinline.MessageLevel10
		level10N02 *pbinline.MessageLevel10

		level11N01 *pbinline.MessageLevel11

		level12N01 *pbinline.MessageLevel12

		level14N01 *pbinline.MessageLevel14
		level14N02 *pbinline.MessageLevel14
		level14N03 *pbinline.MessageLevel14
		level14N04 *pbinline.MessageLevel14
		level14N05 *pbinline.MessageLevel14
		level14N06 *pbinline.MessageLevel14
		level14N07 *pbinline.MessageLevel14
		level14N08 *pbinline.MessageLevel14

		level15N01 *pbinline.MessageLevel15
		level15N02 *pbinline.MessageLevel15
		level15N03 *pbinline.MessageLevel15
		level15N04 *pbinline.MessageLevel15

		level16N01 *pbinline.MessageLevel16
		level16N02 *pbinline.MessageLevel16
		level16N03 *pbinline.MessageLevel16
		level16N04 *pbinline.MessageLevel16

		level18N01 *pbinline.MessageLevel18
		level18N02 *pbinline.MessageLevel18
		level18N03 *pbinline.MessageLevel18
		level18N04 *pbinline.MessageLevel18

		level19N01 *pbinline.MessageLevel19
		level19N02 *pbinline.MessageLevel19

		level20N01 *pbinline.MessageLevel20
		level20N02 *pbinline.MessageLevel20

		level22N01 *pbinline.MessageLevel22
		level22N02 *pbinline.MessageLevel22
		level22N03 *pbinline.MessageLevel22
		level22N04 *pbinline.MessageLevel22
		level22N05 *pbinline.MessageLevel22
		level22N06 *pbinline.MessageLevel22
		level22N07 *pbinline.MessageLevel22
		level22N08 *pbinline.MessageLevel22
		level22N09 *pbinline.MessageLevel22
		level22N10 *pbinline.MessageLevel22
		level22N11 *pbinline.MessageLevel22
		level22N12 *pbinline.MessageLevel22
		level22N13 *pbinline.MessageLevel22
		level22N14 *pbinline.MessageLevel22
		level22N15 *pbinline.MessageLevel22
		level22N16 *pbinline.MessageLevel22

		level23N01 *pbinline.MessageLevel23
		level23N02 *pbinline.MessageLevel23
		level23N03 *pbinline.MessageLevel23
		level23N04 *pbinline.MessageLevel23
		level23N05 *pbinline.MessageLevel23
		level23N06 *pbinline.MessageLevel23
		level23N07 *pbinline.MessageLevel23
		level23N08 *pbinline.MessageLevel23

		level24N01 *pbinline.MessageLevel24
		level24N02 *pbinline.MessageLevel24
		level24N03 *pbinline.MessageLevel24
		level24N04 *pbinline.MessageLevel24
		level24N05 *pbinline.MessageLevel24
		level24N06 *pbinline.MessageLevel24
		level24N07 *pbinline.MessageLevel24
		level24N08 *pbinline.MessageLevel24

		level26N01 *pbinline.MessageLevel26
		level26N02 *pbinline.MessageLevel26
		level26N03 *pbinline.MessageLevel26
		level26N04 *pbinline.MessageLevel26
		level26N05 *pbinline.MessageLevel26
		level26N06 *pbinline.MessageLevel26
		level26N07 *pbinline.MessageLevel26
		level26N08 *pbinline.MessageLevel26

		level27N01 *pbinline.MessageLevel27
		level27N02 *pbinline.MessageLevel27
		level27N03 *pbinline.MessageLevel27
		level27N04 *pbinline.MessageLevel27

		level28N01 *pbinline.MessageLevel28
		level28N02 *pbinline.MessageLevel28
		level28N03 *pbinline.MessageLevel28
		level28N04 *pbinline.MessageLevel28

		level30N01 *pbinline.MessageLevel30
		level30N02 *pbinline.MessageLevel30
		level30N03 *pbinline.MessageLevel30
		level30N04 *pbinline.MessageLevel30
		level30N05 *pbinline.MessageLevel30
		level30N06 *pbinline.MessageLevel30
		level30N07 *pbinline.MessageLevel30
		level30N08 *pbinline.MessageLevel30
		level30N09 *pbinline.MessageLevel30
		level30N10 *pbinline.MessageLevel30
		level30N11 *pbinline.MessageLevel30
		level30N12 *pbinline.MessageLevel30
		level30N13 *pbinline.MessageLevel30
		level30N14 *pbinline.MessageLevel30
		level30N15 *pbinline.MessageLevel30
		level30N16 *pbinline.MessageLevel30

		level31N01 *pbinline.MessageLevel31
		level31N02 *pbinline.MessageLevel31
		level31N03 *pbinline.MessageLevel31
		level31N04 *pbinline.MessageLevel31
		level31N05 *pbinline.MessageLevel31
		level31N06 *pbinline.MessageLevel31
		level31N07 *pbinline.MessageLevel31
		level31N08 *pbinline.MessageLevel31

		level32N01 *pbinline.MessageLevel32
		level32N02 *pbinline.MessageLevel32
		level32N03 *pbinline.MessageLevel32
		level32N04 *pbinline.MessageLevel32
		level32N05 *pbinline.MessageLevel32
		level32N06 *pbinline.MessageLevel32
		level32N07 *pbinline.MessageLevel32
		level32N08 *pbinline.MessageLevel32

		level34N01 *pbinline.MessageLevel34
		level34N02 *pbinline.MessageLevel34
		level34N03 *pbinline.MessageLevel34
		level34N04 *pbinline.MessageLevel34
		level34N05 *pbinline.MessageLevel34
		level34N06 *pbinline.MessageLevel34
		level34N07 *pbinline.MessageLevel34
		level34N08 *pbinline.MessageLevel34

		level35N01 *pbinline.MessageLevel35
		level35N02 *pbinline.MessageLevel35
		level35N03 *pbinline.MessageLevel35
		level35N04 *pbinline.MessageLevel35

		level36N01 *pbinline.MessageLevel36
		level36N02 *pbinline.MessageLevel36
		level36N03 *pbinline.MessageLevel36
		level36N04 *pbinline.MessageLevel36

		level38N01 *pbinline.MessageLevel38
		level38N02 *pbinline.MessageLevel38
		level38N03 *pbinline.MessageLevel38
		level38N04 *pbinline.MessageLevel38
		level38N05 *pbinline.MessageLevel38
		level38N06 *pbinline.MessageLevel38
		level38N07 *pbinline.MessageLevel38
		level38N08 *pbinline.MessageLevel38
		level38N09 *pbinline.MessageLevel38
		level38N10 *pbinline.MessageLevel38
		level38N11 *pbinline.MessageLevel38
		level38N12 *pbinline.MessageLevel38
		level38N13 *pbinline.MessageLevel38
		level38N14 *pbinline.MessageLevel38
		level38N15 *pbinline.MessageLevel38
		level38N16 *pbinline.MessageLevel38
		level38N17 *pbinline.MessageLevel38
		level38N18 *pbinline.MessageLevel38
		level38N19 *pbinline.MessageLevel38
		level38N20 *pbinline.MessageLevel38
		level38N21 *pbinline.MessageLevel38
		level38N22 *pbinline.MessageLevel38
		level38N23 *pbinline.MessageLevel38
		level38N24 *pbinline.MessageLevel38
		level38N25 *pbinline.MessageLevel38
		level38N26 *pbinline.MessageLevel38
		level38N27 *pbinline.MessageLevel38
		level38N28 *pbinline.MessageLevel38
		level38N29 *pbinline.MessageLevel38
		level38N30 *pbinline.MessageLevel38
		level38N31 *pbinline.MessageLevel38
		level38N32 *pbinline.MessageLevel38

		level39N01 *pbinline.MessageLevel39
		level39N02 *pbinline.MessageLevel39
		level39N03 *pbinline.MessageLevel39
		level39N04 *pbinline.MessageLevel39
		level39N05 *pbinline.MessageLevel39
		level39N06 *pbinline.MessageLevel39
		level39N07 *pbinline.MessageLevel39
		level39N08 *pbinline.MessageLevel39
		level39N09 *pbinline.MessageLevel39
		level39N10 *pbinline.MessageLevel39
		level39N11 *pbinline.MessageLevel39
		level39N12 *pbinline.MessageLevel39
		level39N13 *pbinline.MessageLevel39
		level39N14 *pbinline.MessageLevel39
		level39N15 *pbinline.MessageLevel39
		level39N16 *pbinline.MessageLevel39

		level40N01 *pbinline.MessageLevel40
		level40N02 *pbinline.MessageLevel40
		level40N03 *pbinline.MessageLevel40
		level40N04 *pbinline.MessageLevel40
		level40N05 *pbinline.MessageLevel40
		level40N06 *pbinline.MessageLevel40
		level40N07 *pbinline.MessageLevel40
		level40N08 *pbinline.MessageLevel40
		level40N09 *pbinline.MessageLevel40
		level40N10 *pbinline.MessageLevel40
		level40N11 *pbinline.MessageLevel40
		level40N12 *pbinline.MessageLevel40
		level40N13 *pbinline.MessageLevel40
		level40N14 *pbinline.MessageLevel40
		level40N15 *pbinline.MessageLevel40
		level40N16 *pbinline.MessageLevel40
	)

	{ // message in level02N01
		{ // message in level06N01
			{ // message for level14N01
				{ // message for level22N01
					level38N01 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N01-s101",
						Level38FString2:                   "L38-N01-s111",
						Level38PString1:                   utils.PointerString("L38-N01-p101"),
						Level38PString2:                   utils.PointerString("L38-N01-p111"),
						Level38RString1:                   []string{"L38-N01-r101", "L38-N01-r102", "L38-N01-r103"},
						Level38RString2:                   []string{"L38-N01-r111", "L38-N01-r112", "L38-N01-r113"},
						Level38MString1:                   map[string]string{"L38-N01-k101": "L38-N01-v101", "L38-N01-k102": "L38-N01-v102", "L38-N01-k103": "L38-N01-v103"},
						Level38MString2:                   map[string]string{"L38-N01-k111": "L38-N01-v111", "L38-N01-k112": "L38-N01-v112", "L38-N01-k113": "L38-N01-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N01-e101", FString2: "L38-N01-e102", FString3: "L38-N01-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N01-e111", FString2: "L38-N01-e112", FString3: "L38-N01-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N01-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N01-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N01-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N01-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N01-i111", IgnoreFieldsString2: "L38-N01-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N01-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N01-i131"},
					}
					level39N01 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N01-s101",
						Level39FString2:                   "L39-N01-s111",
						Level39PString1:                   utils.PointerString("L39-N01-p101"),
						Level39PString2:                   utils.PointerString("L39-N01-p111"),
						Level39RString1:                   []string{"L39-N01-r101", "L39-N01-r102", "L39-N01-r103"},
						Level39RString2:                   []string{"L39-N01-r111", "L39-N01-r112", "L39-N01-r113"},
						Level39MString1:                   map[string]string{"L39-N01-k101": "L39-N01-v101", "L39-N01-k102": "L39-N01-v102", "L39-N01-k103": "L39-N01-v103"},
						Level39MString2:                   map[string]string{"L39-N01-k111": "L39-N01-v111", "L39-N01-k112": "L39-N01-v112", "L39-N01-k113": "L39-N01-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N01-e101", FString2: "L39-N01-e102", FString3: "L39-N01-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N01-e111", FString2: "L39-N01-e112", FString3: "L39-N01-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N01-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N01-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N01-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N01-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N01-i111", IgnoreFieldsString2: "L39-N01-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N01-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N01-i131"},
					}
					level38N02 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N02-s101",
						Level38FString2:                   "L38-N02-s111",
						Level38PString1:                   utils.PointerString("L38-N02-p101"),
						Level38PString2:                   utils.PointerString("L38-N02-p111"),
						Level38RString1:                   []string{"L38-N02-r101", "L38-N02-r102", "L38-N02-r103"},
						Level38RString2:                   []string{"L38-N02-r111", "L38-N02-r112", "L38-N02-r113"},
						Level38MString1:                   map[string]string{"L38-N02-k101": "L38-N02-v101", "L38-N02-k102": "L38-N02-v102", "L38-N02-k103": "L38-N02-v103"},
						Level38MString2:                   map[string]string{"L38-N02-k111": "L38-N02-v111", "L38-N02-k112": "L38-N02-v112", "L38-N02-k113": "L38-N02-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N02-e101", FString2: "L38-N02-e102", FString3: "L38-N02-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N02-e111", FString2: "L38-N02-e112", FString3: "L38-N02-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N02-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N02-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N02-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N02-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N02-i111", IgnoreFieldsString2: "L38-N02-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N02-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N02-i131"},
					}
					level40N01 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N01-s101",
						Level40FString2:                   "L40-N01-s111",
						Level40PString1:                   utils.PointerString("L40-N01-p101"),
						Level40PString2:                   utils.PointerString("L40-N01-p111"),
						Level40RString1:                   []string{"L40-N01-r101", "L40-N01-r102", "L40-N01-r103"},
						Level40RString2:                   []string{"L40-N01-r111", "L40-N01-r112", "L40-N01-r113"},
						Level40MString1:                   map[string]string{"L40-N01-k101": "L40-N01-v101", "L40-N01-k102": "L40-N01-v102", "L40-N01-k103": "L40-N01-v103"},
						Level40MString2:                   map[string]string{"L40-N01-k111": "L40-N01-v111", "L40-N01-k112": "L40-N01-v112", "L40-N01-k113": "L40-N01-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N01-e101", FString2: "L40-N01-e102", FString3: "L40-N01-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N01-e111", FString2: "L40-N01-e112", FString3: "L40-N01-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N01-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N01-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N01-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N01-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N01-i111", IgnoreFieldsString2: "L40-N01-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N01-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N01-i131"},
					}
				}
				level22N01 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N01-s101",
					Level22FString2:                   "L22-N01-s111",
					Level22PString1:                   utils.PointerString("L22-N01-p101"),
					Level22PString2:                   utils.PointerString("L22-N01-p111"),
					Level22RString1:                   []string{"L22-N01-r101", "L22-N01-r102", "L22-N01-r103"},
					Level22RString2:                   []string{"L22-N01-r111", "L22-N01-r112", "L22-N01-r113"},
					Level22MString1:                   map[string]string{"L22-N01-k101": "L22-N01-v101", "L22-N01-k102": "L22-N01-v102", "L22-N01-k103": "L22-N01-v103"},
					Level22MString2:                   map[string]string{"L22-N01-k111": "L22-N01-v111", "L22-N01-k112": "L22-N01-v112", "L22-N01-k113": "L22-N01-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N01-e101", FString2: "L22-N01-e102", FString3: "L22-N01-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N01-e111", FString2: "L22-N01-e112", FString3: "L22-N01-e113"},
					Level22FMessageInline38:           level38N01,
					Level22FMessageInline39:           level39N01,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N02},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N01},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N01-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N01-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N01-i111", IgnoreFieldsString2: "L22-N01-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N01-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N01-i131"},
				}
				level23N01 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N01-s101",
					Level23FString2:                   "L23-N01-s111",
					Level23PString1:                   utils.PointerString("L23-N01-p101"),
					Level23PString2:                   utils.PointerString("L23-N01-p111"),
					Level23RString1:                   []string{"L23-N01-r101", "L23-N01-r102", "L23-N01-r103"},
					Level23RString2:                   []string{"L23-N01-r111", "L23-N01-r112", "L23-N01-r113"},
					Level23MString1:                   map[string]string{"L23-N01-k101": "L23-N01-v101", "L23-N01-k102": "L23-N01-v102", "L23-N01-k103": "L23-N01-v103"},
					Level23MString2:                   map[string]string{"L23-N01-k111": "L23-N01-v111", "L23-N01-k112": "L23-N01-v112", "L23-N01-k113": "L23-N01-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N01-e101", FString2: "L23-N01-e102", FString3: "L23-N01-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N01-e111", FString2: "L23-N01-e112", FString3: "L23-N01-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N01-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N01-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N01-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N01-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N01-i111", IgnoreFieldsString2: "L23-N01-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N01-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N01-i131"},
				}
				{ // message for level22N02
					level38N03 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N03-s101",
						Level38FString2:                   "L38-N03-s111",
						Level38PString1:                   utils.PointerString("L38-N03-p101"),
						Level38PString2:                   utils.PointerString("L38-N03-p111"),
						Level38RString1:                   []string{"L38-N03-r101", "L38-N03-r102", "L38-N03-r103"},
						Level38RString2:                   []string{"L38-N03-r111", "L38-N03-r112", "L38-N03-r113"},
						Level38MString1:                   map[string]string{"L38-N03-k101": "L38-N03-v101", "L38-N03-k102": "L38-N03-v102", "L38-N03-k103": "L38-N03-v103"},
						Level38MString2:                   map[string]string{"L38-N03-k111": "L38-N03-v111", "L38-N03-k112": "L38-N03-v112", "L38-N03-k113": "L38-N03-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N03-e101", FString2: "L38-N03-e102", FString3: "L38-N03-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N03-e111", FString2: "L38-N03-e112", FString3: "L38-N03-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N03-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N03-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N03-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N03-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N03-i111", IgnoreFieldsString2: "L38-N03-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N03-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N03-i131"},
					}
					level39N02 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N02-s101",
						Level39FString2:                   "L39-N02-s111",
						Level39PString1:                   utils.PointerString("L39-N02-p101"),
						Level39PString2:                   utils.PointerString("L39-N02-p111"),
						Level39RString1:                   []string{"L39-N02-r101", "L39-N02-r102", "L39-N02-r103"},
						Level39RString2:                   []string{"L39-N02-r111", "L39-N02-r112", "L39-N02-r113"},
						Level39MString1:                   map[string]string{"L39-N02-k101": "L39-N02-v101", "L39-N02-k102": "L39-N02-v102", "L39-N02-k103": "L39-N02-v103"},
						Level39MString2:                   map[string]string{"L39-N02-k111": "L39-N02-v111", "L39-N02-k112": "L39-N02-v112", "L39-N02-k113": "L39-N02-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N02-e101", FString2: "L39-N02-e102", FString3: "L39-N02-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N02-e111", FString2: "L39-N02-e112", FString3: "L39-N02-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N02-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N02-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N02-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N02-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N02-i111", IgnoreFieldsString2: "L39-N02-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N02-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N02-i131"},
					}
					level38N04 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N04-s101",
						Level38FString2:                   "L38-N04-s111",
						Level38PString1:                   utils.PointerString("L38-N04-p101"),
						Level38PString2:                   utils.PointerString("L38-N04-p111"),
						Level38RString1:                   []string{"L38-N04-r101", "L38-N04-r102", "L38-N04-r103"},
						Level38RString2:                   []string{"L38-N04-r111", "L38-N04-r112", "L38-N04-r113"},
						Level38MString1:                   map[string]string{"L38-N04-k101": "L38-N04-v101", "L38-N04-k102": "L38-N04-v102", "L38-N04-k103": "L38-N04-v103"},
						Level38MString2:                   map[string]string{"L38-N04-k111": "L38-N04-v111", "L38-N04-k112": "L38-N04-v112", "L38-N04-k113": "L38-N04-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N04-e101", FString2: "L38-N04-e102", FString3: "L38-N04-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N04-e111", FString2: "L38-N04-e112", FString3: "L38-N04-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N04-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N04-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N04-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N04-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N04-i111", IgnoreFieldsString2: "L38-N04-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N04-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N04-i131"},
					}
					level40N02 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N02-s101",
						Level40FString2:                   "L40-N02-s111",
						Level40PString1:                   utils.PointerString("L40-N02-p101"),
						Level40PString2:                   utils.PointerString("L40-N02-p111"),
						Level40RString1:                   []string{"L40-N02-r101", "L40-N02-r102", "L40-N02-r103"},
						Level40RString2:                   []string{"L40-N02-r111", "L40-N02-r112", "L40-N02-r113"},
						Level40MString1:                   map[string]string{"L40-N02-k101": "L40-N02-v101", "L40-N02-k102": "L40-N02-v102", "L40-N02-k103": "L40-N02-v103"},
						Level40MString2:                   map[string]string{"L40-N02-k111": "L40-N02-v111", "L40-N02-k112": "L40-N02-v112", "L40-N02-k113": "L40-N02-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N02-e101", FString2: "L40-N02-e102", FString3: "L40-N02-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N02-e111", FString2: "L40-N02-e112", FString3: "L40-N02-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N02-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N02-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N02-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N02-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N02-i111", IgnoreFieldsString2: "L40-N02-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N02-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N02-i131"},
					}
				}
				level22N02 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N02-s101",
					Level22FString2:                   "L22-N02-s111",
					Level22PString1:                   utils.PointerString("L22-N02-p101"),
					Level22PString2:                   utils.PointerString("L22-N02-p111"),
					Level22RString1:                   []string{"L22-N02-r101", "L22-N02-r102", "L22-N02-r103"},
					Level22RString2:                   []string{"L22-N02-r111", "L22-N02-r112", "L22-N02-r113"},
					Level22MString1:                   map[string]string{"L22-N02-k101": "L22-N02-v101", "L22-N02-k102": "L22-N02-v102", "L22-N02-k103": "L22-N02-v103"},
					Level22MString2:                   map[string]string{"L22-N02-k111": "L22-N02-v111", "L22-N02-k112": "L22-N02-v112", "L22-N02-k113": "L22-N02-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N02-e101", FString2: "L22-N02-e102", FString3: "L22-N02-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N02-e111", FString2: "L22-N02-e112", FString3: "L22-N02-e113"},
					Level22FMessageInline38:           level38N03,
					Level22FMessageInline39:           level39N02,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N04},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N02},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N02-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N02-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N02-i111", IgnoreFieldsString2: "L22-N02-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N02-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N02-i131"},
				}
				level24N01 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N01-s101",
					Level24FString2:                   "L24-N01-s111",
					Level24PString1:                   utils.PointerString("L24-N01-p101"),
					Level24PString2:                   utils.PointerString("L24-N01-p111"),
					Level24RString1:                   []string{"L24-N01-r101", "L24-N01-r102", "L24-N01-r103"},
					Level24RString2:                   []string{"L24-N01-r111", "L24-N01-r112", "L24-N01-r113"},
					Level24MString1:                   map[string]string{"L24-N01-k101": "L24-N01-v101", "L24-N01-k102": "L24-N01-v102", "L24-N01-k103": "L24-N01-v103"},
					Level24MString2:                   map[string]string{"L24-N01-k111": "L24-N01-v111", "L24-N01-k112": "L24-N01-v112", "L24-N01-k113": "L24-N01-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N01-e101", FString2: "L24-N01-e102", FString3: "L24-N01-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N01-e111", FString2: "L24-N01-e112", FString3: "L24-N01-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N01-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N01-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N01-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N01-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N01-i111", IgnoreFieldsString2: "L24-N01-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N01-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N01-i131"},
				}
			}
			level14N01 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N01-s101",
				Level14FString2:                   "L14-N01-s111",
				Level14PString1:                   utils.PointerString("L14-N01-p101"),
				Level14PString2:                   utils.PointerString("L14-N01-p111"),
				Level14RString1:                   []string{"L14-N01-r101", "L14-N01-r102", "L14-N01-r103"},
				Level14RString2:                   []string{"L14-N01-r111", "L14-N01-r112", "L14-N01-r113"},
				Level14MString1:                   map[string]string{"L14-N01-k101": "L14-N01-v101", "L14-N01-k102": "L14-N01-v102", "L14-N01-k103": "L14-N01-v103"},
				Level14MString2:                   map[string]string{"L14-N01-k111": "L14-N01-v111", "L14-N01-k112": "L14-N01-v112", "L14-N01-k113": "L14-N01-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N01-e101", FString2: "L14-N01-e102", FString3: "L14-N01-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N01-e111", FString2: "L14-N01-e112", FString3: "L14-N01-e113"},
				Level14FMessageInline22:           level22N01,
				Level14FMessageInline23:           level23N01,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N02},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N01},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N01-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N01-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N01-i111", IgnoreFieldsString2: "L14-N01-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N01-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N01-i131"},
			}
			level15N01 = &pbinline.MessageLevel15{
				Level15FString1:                   "L15-N01-s101",
				Level15FString2:                   "L15-N01-s111",
				Level15PString1:                   utils.PointerString("L15-N01-p101"),
				Level15PString2:                   utils.PointerString("L15-N01-p111"),
				Level15RString1:                   []string{"L15-N01-r101", "L15-N01-r102", "L15-N01-r103"},
				Level15RString2:                   []string{"L15-N01-r111", "L15-N01-r112", "L15-N01-r113"},
				Level15MString1:                   map[string]string{"L15-N01-k101": "L15-N01-v101", "L15-N01-k102": "L15-N01-v102", "L15-N01-k103": "L15-N01-v103"},
				Level15MString2:                   map[string]string{"L15-N01-k111": "L15-N01-v111", "L15-N01-k112": "L15-N01-v112", "L15-N01-k113": "L15-N01-v113"},
				Level15FMessageExtern1:            &pbexternal.Message1{FString1: "L15-N01-e101", FString2: "L15-N01-e102", FString3: "L15-N01-e103"},
				Level15FMessageExtern2:            &pbexternal.Message1{FString1: "L15-N01-e111", FString2: "L15-N01-e112", FString3: "L15-N01-e113"},
				Level15OneOfExtern1:               &pbinline.MessageLevel15_Level15One1String1{Level15One1String1: "L15-N01-es101"},
				Level15OneOfInline1:               &pbinline.MessageLevel15_Level15One2String1{Level15One2String1: "L15-N01-is101"},
				Level15OneOfOmitempty1:            &pbinline.MessageLevel15_Level15One3String1{Level15One3String1: "L15-N01-o101"},
				Level15FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level15FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L15-N01-i101"},
				Level15FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L15-N01-i111", IgnoreFieldsString2: "L15-N01-i112"},
				Level15OneOfIgnoreSelf1:           &pbinline.MessageLevel15_Level15One4String1{Level15One4String1: "L15-N01-i121"},
				Level15OneOfIgnoreParts1:          &pbinline.MessageLevel15_Level15One5String1{Level15One5String1: "L15-N01-i131"},
			}
			{ // message for level14N02
				{ // message for level22N03
					level38N05 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N05-s101",
						Level38FString2:                   "L38-N05-s111",
						Level38PString1:                   utils.PointerString("L38-N05-p101"),
						Level38PString2:                   utils.PointerString("L38-N05-p111"),
						Level38RString1:                   []string{"L38-N05-r101", "L38-N05-r102", "L38-N05-r103"},
						Level38RString2:                   []string{"L38-N05-r111", "L38-N05-r112", "L38-N05-r113"},
						Level38MString1:                   map[string]string{"L38-N05-k101": "L38-N05-v101", "L38-N05-k102": "L38-N05-v102", "L38-N05-k103": "L38-N05-v103"},
						Level38MString2:                   map[string]string{"L38-N05-k111": "L38-N05-v111", "L38-N05-k112": "L38-N05-v112", "L38-N05-k113": "L38-N05-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N05-e101", FString2: "L38-N05-e102", FString3: "L38-N05-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N05-e111", FString2: "L38-N05-e112", FString3: "L38-N05-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N05-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N05-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N05-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N05-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N05-i111", IgnoreFieldsString2: "L38-N05-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N05-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N05-i131"},
					}
					level39N03 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N03-s101",
						Level39FString2:                   "L39-N03-s111",
						Level39PString1:                   utils.PointerString("L39-N03-p101"),
						Level39PString2:                   utils.PointerString("L39-N03-p111"),
						Level39RString1:                   []string{"L39-N03-r101", "L39-N03-r102", "L39-N03-r103"},
						Level39RString2:                   []string{"L39-N03-r111", "L39-N03-r112", "L39-N03-r113"},
						Level39MString1:                   map[string]string{"L39-N03-k101": "L39-N03-v101", "L39-N03-k102": "L39-N03-v102", "L39-N03-k103": "L39-N03-v103"},
						Level39MString2:                   map[string]string{"L39-N03-k111": "L39-N03-v111", "L39-N03-k112": "L39-N03-v112", "L39-N03-k113": "L39-N03-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N03-e101", FString2: "L39-N03-e102", FString3: "L39-N03-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N03-e111", FString2: "L39-N03-e112", FString3: "L39-N03-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N03-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N03-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N03-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N03-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N03-i111", IgnoreFieldsString2: "L39-N03-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N03-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N03-i131"},
					}
					level38N06 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N06-s101",
						Level38FString2:                   "L38-N06-s111",
						Level38PString1:                   utils.PointerString("L38-N06-p101"),
						Level38PString2:                   utils.PointerString("L38-N06-p111"),
						Level38RString1:                   []string{"L38-N06-r101", "L38-N06-r102", "L38-N06-r103"},
						Level38RString2:                   []string{"L38-N06-r111", "L38-N06-r112", "L38-N06-r113"},
						Level38MString1:                   map[string]string{"L38-N06-k101": "L38-N06-v101", "L38-N06-k102": "L38-N06-v102", "L38-N06-k103": "L38-N06-v103"},
						Level38MString2:                   map[string]string{"L38-N06-k111": "L38-N06-v111", "L38-N06-k112": "L38-N06-v112", "L38-N06-k113": "L38-N06-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N06-e101", FString2: "L38-N06-e102", FString3: "L38-N06-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N06-e111", FString2: "L38-N06-e112", FString3: "L38-N06-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N06-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N06-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N06-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N06-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N06-i111", IgnoreFieldsString2: "L38-N06-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N06-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N06-i131"},
					}
					level40N03 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N03-s101",
						Level40FString2:                   "L40-N03-s111",
						Level40PString1:                   utils.PointerString("L40-N03-p101"),
						Level40PString2:                   utils.PointerString("L40-N03-p111"),
						Level40RString1:                   []string{"L40-N03-r101", "L40-N03-r102", "L40-N03-r103"},
						Level40RString2:                   []string{"L40-N03-r111", "L40-N03-r112", "L40-N03-r113"},
						Level40MString1:                   map[string]string{"L40-N03-k101": "L40-N03-v101", "L40-N03-k102": "L40-N03-v102", "L40-N03-k103": "L40-N03-v103"},
						Level40MString2:                   map[string]string{"L40-N03-k111": "L40-N03-v111", "L40-N03-k112": "L40-N03-v112", "L40-N03-k113": "L40-N03-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N03-e101", FString2: "L40-N03-e102", FString3: "L40-N03-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N03-e111", FString2: "L40-N03-e112", FString3: "L40-N03-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N03-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N03-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N03-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N03-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N03-i111", IgnoreFieldsString2: "L40-N03-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N03-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N03-i131"},
					}
				}
				level22N03 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N03-s101",
					Level22FString2:                   "L22-N03-s111",
					Level22PString1:                   utils.PointerString("L22-N03-p101"),
					Level22PString2:                   utils.PointerString("L22-N03-p111"),
					Level22RString1:                   []string{"L22-N03-r101", "L22-N03-r102", "L22-N03-r103"},
					Level22RString2:                   []string{"L22-N03-r111", "L22-N03-r112", "L22-N03-r113"},
					Level22MString1:                   map[string]string{"L22-N03-k101": "L22-N03-v101", "L22-N03-k102": "L22-N03-v102", "L22-N03-k103": "L22-N03-v103"},
					Level22MString2:                   map[string]string{"L22-N03-k111": "L22-N03-v111", "L22-N03-k112": "L22-N03-v112", "L22-N03-k113": "L22-N03-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N03-e101", FString2: "L22-N03-e102", FString3: "L22-N03-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N03-e111", FString2: "L22-N03-e112", FString3: "L22-N03-e113"},
					Level22FMessageInline38:           level38N05,
					Level22FMessageInline39:           level39N03,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N06},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N03},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N03-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N03-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N03-i111", IgnoreFieldsString2: "L22-N03-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N03-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N03-i131"},
				}
				level23N02 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N02-s101",
					Level23FString2:                   "L23-N02-s111",
					Level23PString1:                   utils.PointerString("L23-N02-p101"),
					Level23PString2:                   utils.PointerString("L23-N02-p111"),
					Level23RString1:                   []string{"L23-N02-r101", "L23-N02-r102", "L23-N02-r103"},
					Level23RString2:                   []string{"L23-N02-r111", "L23-N02-r112", "L23-N02-r113"},
					Level23MString1:                   map[string]string{"L23-N02-k101": "L23-N02-v101", "L23-N02-k102": "L23-N02-v102", "L23-N02-k103": "L23-N02-v103"},
					Level23MString2:                   map[string]string{"L23-N02-k111": "L23-N02-v111", "L23-N02-k112": "L23-N02-v112", "L23-N02-k113": "L23-N02-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N02-e101", FString2: "L23-N02-e102", FString3: "L23-N02-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N02-e111", FString2: "L23-N02-e112", FString3: "L23-N02-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N02-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N02-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N02-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N02-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N02-i111", IgnoreFieldsString2: "L23-N02-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N02-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N02-i131"},
				}
				{ // message for level22N04
					level38N07 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N07-s101",
						Level38FString2:                   "L38-N07-s111",
						Level38PString1:                   utils.PointerString("L38-N07-p101"),
						Level38PString2:                   utils.PointerString("L38-N07-p111"),
						Level38RString1:                   []string{"L38-N07-r101", "L38-N07-r102", "L38-N07-r103"},
						Level38RString2:                   []string{"L38-N07-r111", "L38-N07-r112", "L38-N07-r113"},
						Level38MString1:                   map[string]string{"L38-N07-k101": "L38-N07-v101", "L38-N07-k102": "L38-N07-v102", "L38-N07-k103": "L38-N07-v103"},
						Level38MString2:                   map[string]string{"L38-N07-k111": "L38-N07-v111", "L38-N07-k112": "L38-N07-v112", "L38-N07-k113": "L38-N07-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N07-e101", FString2: "L38-N07-e102", FString3: "L38-N07-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N07-e111", FString2: "L38-N07-e112", FString3: "L38-N07-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N07-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N07-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N07-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N07-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N07-i111", IgnoreFieldsString2: "L38-N07-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N07-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N07-i131"},
					}
					level39N04 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N04-s101",
						Level39FString2:                   "L39-N04-s111",
						Level39PString1:                   utils.PointerString("L39-N04-p101"),
						Level39PString2:                   utils.PointerString("L39-N04-p111"),
						Level39RString1:                   []string{"L39-N04-r101", "L39-N04-r102", "L39-N04-r103"},
						Level39RString2:                   []string{"L39-N04-r111", "L39-N04-r112", "L39-N04-r113"},
						Level39MString1:                   map[string]string{"L39-N04-k101": "L39-N04-v101", "L39-N04-k102": "L39-N04-v102", "L39-N04-k103": "L39-N04-v103"},
						Level39MString2:                   map[string]string{"L39-N04-k111": "L39-N04-v111", "L39-N04-k112": "L39-N04-v112", "L39-N04-k113": "L39-N04-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N04-e101", FString2: "L39-N04-e102", FString3: "L39-N04-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N04-e111", FString2: "L39-N04-e112", FString3: "L39-N04-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N04-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N04-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N04-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N04-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N04-i111", IgnoreFieldsString2: "L39-N04-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N04-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N04-i131"},
					}
					level38N08 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N08-s101",
						Level38FString2:                   "L38-N08-s111",
						Level38PString1:                   utils.PointerString("L38-N08-p101"),
						Level38PString2:                   utils.PointerString("L38-N08-p111"),
						Level38RString1:                   []string{"L38-N08-r101", "L38-N08-r102", "L38-N08-r103"},
						Level38RString2:                   []string{"L38-N08-r111", "L38-N08-r112", "L38-N08-r113"},
						Level38MString1:                   map[string]string{"L38-N08-k101": "L38-N08-v101", "L38-N08-k102": "L38-N08-v102", "L38-N08-k103": "L38-N08-v103"},
						Level38MString2:                   map[string]string{"L38-N08-k111": "L38-N08-v111", "L38-N08-k112": "L38-N08-v112", "L38-N08-k113": "L38-N08-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N08-e101", FString2: "L38-N08-e102", FString3: "L38-N08-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N08-e111", FString2: "L38-N08-e112", FString3: "L38-N08-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N08-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N08-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N08-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N08-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N08-i111", IgnoreFieldsString2: "L38-N08-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N08-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N08-i131"},
					}
					level40N04 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N04-s101",
						Level40FString2:                   "L40-N04-s111",
						Level40PString1:                   utils.PointerString("L40-N04-p101"),
						Level40PString2:                   utils.PointerString("L40-N04-p111"),
						Level40RString1:                   []string{"L40-N04-r101", "L40-N04-r102", "L40-N04-r103"},
						Level40RString2:                   []string{"L40-N04-r111", "L40-N04-r112", "L40-N04-r113"},
						Level40MString1:                   map[string]string{"L40-N04-k101": "L40-N04-v101", "L40-N04-k102": "L40-N04-v102", "L40-N04-k103": "L40-N04-v103"},
						Level40MString2:                   map[string]string{"L40-N04-k111": "L40-N04-v111", "L40-N04-k112": "L40-N04-v112", "L40-N04-k113": "L40-N04-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N04-e101", FString2: "L40-N04-e102", FString3: "L40-N04-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N04-e111", FString2: "L40-N04-e112", FString3: "L40-N04-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N04-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N04-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N04-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N04-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N04-i111", IgnoreFieldsString2: "L40-N04-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N04-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N04-i131"},
					}
				}
				level22N04 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N04-s101",
					Level22FString2:                   "L22-N04-s111",
					Level22PString1:                   utils.PointerString("L22-N04-p101"),
					Level22PString2:                   utils.PointerString("L22-N04-p111"),
					Level22RString1:                   []string{"L22-N04-r101", "L22-N04-r102", "L22-N04-r103"},
					Level22RString2:                   []string{"L22-N04-r111", "L22-N04-r112", "L22-N04-r113"},
					Level22MString1:                   map[string]string{"L22-N04-k101": "L22-N04-v101", "L22-N04-k102": "L22-N04-v102", "L22-N04-k103": "L22-N04-v103"},
					Level22MString2:                   map[string]string{"L22-N04-k111": "L22-N04-v111", "L22-N04-k112": "L22-N04-v112", "L22-N04-k113": "L22-N04-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N04-e101", FString2: "L22-N04-e102", FString3: "L22-N04-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N04-e111", FString2: "L22-N04-e112", FString3: "L22-N04-e113"},
					Level22FMessageInline38:           level38N07,
					Level22FMessageInline39:           level39N04,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N08},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N04},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N04-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N04-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N04-i111", IgnoreFieldsString2: "L22-N04-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N04-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N04-i131"},
				}
				level24N02 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N02-s101",
					Level24FString2:                   "L24-N02-s111",
					Level24PString1:                   utils.PointerString("L24-N02-p101"),
					Level24PString2:                   utils.PointerString("L24-N02-p111"),
					Level24RString1:                   []string{"L24-N02-r101", "L24-N02-r102", "L24-N02-r103"},
					Level24RString2:                   []string{"L24-N02-r111", "L24-N02-r112", "L24-N02-r113"},
					Level24MString1:                   map[string]string{"L24-N02-k101": "L24-N02-v101", "L24-N02-k102": "L24-N02-v102", "L24-N02-k103": "L24-N02-v103"},
					Level24MString2:                   map[string]string{"L24-N02-k111": "L24-N02-v111", "L24-N02-k112": "L24-N02-v112", "L24-N02-k113": "L24-N02-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N02-e101", FString2: "L24-N02-e102", FString3: "L24-N02-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N02-e111", FString2: "L24-N02-e112", FString3: "L24-N02-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N02-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N02-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N02-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N02-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N02-i111", IgnoreFieldsString2: "L24-N02-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N02-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N02-i131"},
				}
			}
			level14N02 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N02-s101",
				Level14FString2:                   "L14-N02-s111",
				Level14PString1:                   utils.PointerString("L14-N02-p101"),
				Level14PString2:                   utils.PointerString("L14-N02-p111"),
				Level14RString1:                   []string{"L14-N02-r101", "L14-N02-r102", "L14-N02-r103"},
				Level14RString2:                   []string{"L14-N02-r111", "L14-N02-r112", "L14-N02-r113"},
				Level14MString1:                   map[string]string{"L14-N02-k101": "L14-N02-v101", "L14-N02-k102": "L14-N02-v102", "L14-N02-k103": "L14-N02-v103"},
				Level14MString2:                   map[string]string{"L14-N02-k111": "L14-N02-v111", "L14-N02-k112": "L14-N02-v112", "L14-N02-k113": "L14-N02-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N02-e101", FString2: "L14-N02-e102", FString3: "L14-N02-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N02-e111", FString2: "L14-N02-e112", FString3: "L14-N02-e113"},
				Level14FMessageInline22:           level22N03,
				Level14FMessageInline23:           level23N02,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N04},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N02},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N02-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N02-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N02-i111", IgnoreFieldsString2: "L14-N02-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N02-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N02-i131"},
			}
			{ // message for level16N01
				{ // message for level26N01
					level30N01 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N01-s101",
						Level30FString2:                   "L30-N01-s111",
						Level30PString1:                   utils.PointerString("L30-N01-p101"),
						Level30PString2:                   utils.PointerString("L30-N01-p111"),
						Level30RString1:                   []string{"L30-N01-r101", "L30-N01-r102", "L30-N01-r103"},
						Level30RString2:                   []string{"L30-N01-r111", "L30-N01-r112", "L30-N01-r113"},
						Level30MString1:                   map[string]string{"L30-N01-k101": "L30-N01-v101", "L30-N01-k102": "L30-N01-v102", "L30-N01-k103": "L30-N01-v103"},
						Level30MString2:                   map[string]string{"L30-N01-k111": "L30-N01-v111", "L30-N01-k112": "L30-N01-v112", "L30-N01-k113": "L30-N01-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N01-e101", FString2: "L30-N01-e102", FString3: "L30-N01-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N01-e111", FString2: "L30-N01-e112", FString3: "L30-N01-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N01-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N01-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N01-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N01-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N01-i111", IgnoreFieldsString2: "L30-N01-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N01-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N01-i131"},
					}
					level31N01 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N01-s101",
						Level31FString2:                   "L31-N01-s111",
						Level31PString1:                   utils.PointerString("L31-N01-p101"),
						Level31PString2:                   utils.PointerString("L31-N01-p111"),
						Level31RString1:                   []string{"L31-N01-r101", "L31-N01-r102", "L31-N01-r103"},
						Level31RString2:                   []string{"L31-N01-r111", "L31-N01-r112", "L31-N01-r113"},
						Level31MString1:                   map[string]string{"L31-N01-k101": "L31-N01-v101", "L31-N01-k102": "L31-N01-v102", "L31-N01-k103": "L31-N01-v103"},
						Level31MString2:                   map[string]string{"L31-N01-k111": "L31-N01-v111", "L31-N01-k112": "L31-N01-v112", "L31-N01-k113": "L31-N01-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N01-e101", FString2: "L31-N01-e102", FString3: "L31-N01-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N01-e111", FString2: "L31-N01-e112", FString3: "L31-N01-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N01-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N01-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N01-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N01-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N01-i111", IgnoreFieldsString2: "L31-N01-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N01-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N01-i131"},
					}
					level30N02 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N02-s101",
						Level30FString2:                   "L30-N02-s111",
						Level30PString1:                   utils.PointerString("L30-N02-p101"),
						Level30PString2:                   utils.PointerString("L30-N02-p111"),
						Level30RString1:                   []string{"L30-N02-r101", "L30-N02-r102", "L30-N02-r103"},
						Level30RString2:                   []string{"L30-N02-r111", "L30-N02-r112", "L30-N02-r113"},
						Level30MString1:                   map[string]string{"L30-N02-k101": "L30-N02-v101", "L30-N02-k102": "L30-N02-v102", "L30-N02-k103": "L30-N02-v103"},
						Level30MString2:                   map[string]string{"L30-N02-k111": "L30-N02-v111", "L30-N02-k112": "L30-N02-v112", "L30-N02-k113": "L30-N02-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N02-e101", FString2: "L30-N02-e102", FString3: "L30-N02-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N02-e111", FString2: "L30-N02-e112", FString3: "L30-N02-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N02-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N02-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N02-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N02-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N02-i111", IgnoreFieldsString2: "L30-N02-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N02-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N02-i131"},
					}
					level32N01 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N01-s101",
						Level32FString2:                   "L32-N01-s111",
						Level32PString1:                   utils.PointerString("L32-N01-p101"),
						Level32PString2:                   utils.PointerString("L32-N01-p111"),
						Level32RString1:                   []string{"L32-N01-r101", "L32-N01-r102", "L32-N01-r103"},
						Level32RString2:                   []string{"L32-N01-r111", "L32-N01-r112", "L32-N01-r113"},
						Level32MString1:                   map[string]string{"L32-N01-k101": "L32-N01-v101", "L32-N01-k102": "L32-N01-v102", "L32-N01-k103": "L32-N01-v103"},
						Level32MString2:                   map[string]string{"L32-N01-k111": "L32-N01-v111", "L32-N01-k112": "L32-N01-v112", "L32-N01-k113": "L32-N01-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N01-e101", FString2: "L32-N01-e102", FString3: "L32-N01-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N01-e111", FString2: "L32-N01-e112", FString3: "L32-N01-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N01-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N01-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N01-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N01-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N01-i111", IgnoreFieldsString2: "L32-N01-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N01-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N01-i131"},
					}
				}
				level26N01 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N01-s101",
					Level26FString2:                   "L26-N01-s111",
					Level26PString1:                   utils.PointerString("L26-N01-p101"),
					Level26PString2:                   utils.PointerString("L26-N01-p111"),
					Level26RString1:                   []string{"L26-N01-r101", "L26-N01-r102", "L26-N01-r103"},
					Level26RString2:                   []string{"L26-N01-r111", "L26-N01-r112", "L26-N01-r113"},
					Level26MString1:                   map[string]string{"L26-N01-k101": "L26-N01-v101", "L26-N01-k102": "L26-N01-v102", "L26-N01-k103": "L26-N01-v103"},
					Level26MString2:                   map[string]string{"L26-N01-k111": "L26-N01-v111", "L26-N01-k112": "L26-N01-v112", "L26-N01-k113": "L26-N01-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N01-e101", FString2: "L26-N01-e102", FString3: "L26-N01-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N01-e111", FString2: "L26-N01-e112", FString3: "L26-N01-e113"},
					Level26FMessageInline30:           level30N01,
					Level26FMessageInline31:           level31N01,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N02},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N01},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N01-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N01-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N01-i111", IgnoreFieldsString2: "L26-N01-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N01-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N01-i131"},
				}
				level27N01 = &pbinline.MessageLevel27{
					Level27FString1:                   "L27-N01-s101",
					Level27FString2:                   "L27-N01-s111",
					Level27PString1:                   utils.PointerString("L27-N01-p101"),
					Level27PString2:                   utils.PointerString("L27-N01-p111"),
					Level27RString1:                   []string{"L27-N01-r101", "L27-N01-r102", "L27-N01-r103"},
					Level27RString2:                   []string{"L27-N01-r111", "L27-N01-r112", "L27-N01-r113"},
					Level27MString1:                   map[string]string{"L27-N01-k101": "L27-N01-v101", "L27-N01-k102": "L27-N01-v102", "L27-N01-k103": "L27-N01-v103"},
					Level27MString2:                   map[string]string{"L27-N01-k111": "L27-N01-v111", "L27-N01-k112": "L27-N01-v112", "L27-N01-k113": "L27-N01-v113"},
					Level27FMessageExtern1:            &pbexternal.Message1{FString1: "L27-N01-e101", FString2: "L27-N01-e102", FString3: "L27-N01-e103"},
					Level27FMessageExtern2:            &pbexternal.Message1{FString1: "L27-N01-e111", FString2: "L27-N01-e112", FString3: "L27-N01-e113"},
					Level27OneOfExtern1:               &pbinline.MessageLevel27_Level27One1String1{Level27One1String1: "L27-N01-es101"},
					Level27OneOfInline1:               &pbinline.MessageLevel27_Level27One2String1{Level27One2String1: "L27-N01-is101"},
					Level27OneOfOmitempty1:            &pbinline.MessageLevel27_Level27One3String1{Level27One3String1: "L27-N01-o101"},
					Level27FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level27FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L27-N01-i101"},
					Level27FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L27-N01-i111", IgnoreFieldsString2: "L27-N01-i112"},
					Level27OneOfIgnoreSelf1:           &pbinline.MessageLevel27_Level27One4String1{Level27One4String1: "L27-N01-i121"},
					Level27OneOfIgnoreParts1:          &pbinline.MessageLevel27_Level27One5String1{Level27One5String1: "L27-N01-i131"},
				}
				{ // message for level26N02
					level30N03 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N03-s101",
						Level30FString2:                   "L30-N03-s111",
						Level30PString1:                   utils.PointerString("L30-N03-p101"),
						Level30PString2:                   utils.PointerString("L30-N03-p111"),
						Level30RString1:                   []string{"L30-N03-r101", "L30-N03-r102", "L30-N03-r103"},
						Level30RString2:                   []string{"L30-N03-r111", "L30-N03-r112", "L30-N03-r113"},
						Level30MString1:                   map[string]string{"L30-N03-k101": "L30-N03-v101", "L30-N03-k102": "L30-N03-v102", "L30-N03-k103": "L30-N03-v103"},
						Level30MString2:                   map[string]string{"L30-N03-k111": "L30-N03-v111", "L30-N03-k112": "L30-N03-v112", "L30-N03-k113": "L30-N03-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N03-e101", FString2: "L30-N03-e102", FString3: "L30-N03-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N03-e111", FString2: "L30-N03-e112", FString3: "L30-N03-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N03-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N03-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N03-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N03-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N03-i111", IgnoreFieldsString2: "L30-N03-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N03-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N03-i131"},
					}
					level31N02 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N02-s101",
						Level31FString2:                   "L31-N02-s111",
						Level31PString1:                   utils.PointerString("L31-N02-p101"),
						Level31PString2:                   utils.PointerString("L31-N02-p111"),
						Level31RString1:                   []string{"L31-N02-r101", "L31-N02-r102", "L31-N02-r103"},
						Level31RString2:                   []string{"L31-N02-r111", "L31-N02-r112", "L31-N02-r113"},
						Level31MString1:                   map[string]string{"L31-N02-k101": "L31-N02-v101", "L31-N02-k102": "L31-N02-v102", "L31-N02-k103": "L31-N02-v103"},
						Level31MString2:                   map[string]string{"L31-N02-k111": "L31-N02-v111", "L31-N02-k112": "L31-N02-v112", "L31-N02-k113": "L31-N02-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N02-e101", FString2: "L31-N02-e102", FString3: "L31-N02-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N02-e111", FString2: "L31-N02-e112", FString3: "L31-N02-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N02-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N02-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N02-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N02-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N02-i111", IgnoreFieldsString2: "L31-N02-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N02-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N02-i131"},
					}
					level30N04 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N04-s101",
						Level30FString2:                   "L30-N04-s111",
						Level30PString1:                   utils.PointerString("L30-N04-p101"),
						Level30PString2:                   utils.PointerString("L30-N04-p111"),
						Level30RString1:                   []string{"L30-N04-r101", "L30-N04-r102", "L30-N04-r103"},
						Level30RString2:                   []string{"L30-N04-r111", "L30-N04-r112", "L30-N04-r113"},
						Level30MString1:                   map[string]string{"L30-N04-k101": "L30-N04-v101", "L30-N04-k102": "L30-N04-v102", "L30-N04-k103": "L30-N04-v103"},
						Level30MString2:                   map[string]string{"L30-N04-k111": "L30-N04-v111", "L30-N04-k112": "L30-N04-v112", "L30-N04-k113": "L30-N04-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N04-e101", FString2: "L30-N04-e102", FString3: "L30-N04-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N04-e111", FString2: "L30-N04-e112", FString3: "L30-N04-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N04-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N04-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N04-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N04-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N04-i111", IgnoreFieldsString2: "L30-N04-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N04-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N04-i131"},
					}
					level32N02 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N02-s101",
						Level32FString2:                   "L32-N02-s111",
						Level32PString1:                   utils.PointerString("L32-N02-p101"),
						Level32PString2:                   utils.PointerString("L32-N02-p111"),
						Level32RString1:                   []string{"L32-N02-r101", "L32-N02-r102", "L32-N02-r103"},
						Level32RString2:                   []string{"L32-N02-r111", "L32-N02-r112", "L32-N02-r113"},
						Level32MString1:                   map[string]string{"L32-N02-k101": "L32-N02-v101", "L32-N02-k102": "L32-N02-v102", "L32-N02-k103": "L32-N02-v103"},
						Level32MString2:                   map[string]string{"L32-N02-k111": "L32-N02-v111", "L32-N02-k112": "L32-N02-v112", "L32-N02-k113": "L32-N02-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N02-e101", FString2: "L32-N02-e102", FString3: "L32-N02-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N02-e111", FString2: "L32-N02-e112", FString3: "L32-N02-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N02-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N02-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N02-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N02-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N02-i111", IgnoreFieldsString2: "L32-N02-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N02-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N02-i131"},
					}
				}
				level26N02 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N02-s101",
					Level26FString2:                   "L26-N02-s111",
					Level26PString1:                   utils.PointerString("L26-N02-p101"),
					Level26PString2:                   utils.PointerString("L26-N02-p111"),
					Level26RString1:                   []string{"L26-N02-r101", "L26-N02-r102", "L26-N02-r103"},
					Level26RString2:                   []string{"L26-N02-r111", "L26-N02-r112", "L26-N02-r113"},
					Level26MString1:                   map[string]string{"L26-N02-k101": "L26-N02-v101", "L26-N02-k102": "L26-N02-v102", "L26-N02-k103": "L26-N02-v103"},
					Level26MString2:                   map[string]string{"L26-N02-k111": "L26-N02-v111", "L26-N02-k112": "L26-N02-v112", "L26-N02-k113": "L26-N02-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N02-e101", FString2: "L26-N02-e102", FString3: "L26-N02-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N02-e111", FString2: "L26-N02-e112", FString3: "L26-N02-e113"},
					Level26FMessageInline30:           level30N03,
					Level26FMessageInline31:           level31N02,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N04},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N02},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N02-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N02-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N02-i111", IgnoreFieldsString2: "L26-N02-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N02-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N02-i131"},
				}
				{ // message for level28N01
					level34N01 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N01-s101",
						Level34FString2:                   "L34-N01-s111",
						Level34PString1:                   utils.PointerString("L34-N01-p101"),
						Level34PString2:                   utils.PointerString("L34-N01-p111"),
						Level34RString1:                   []string{"L34-N01-r101", "L34-N01-r102", "L34-N01-r103"},
						Level34RString2:                   []string{"L34-N01-r111", "L34-N01-r112", "L34-N01-r113"},
						Level34MString1:                   map[string]string{"L34-N01-k101": "L34-N01-v101", "L34-N01-k102": "L34-N01-v102", "L34-N01-k103": "L34-N01-v103"},
						Level34MString2:                   map[string]string{"L34-N01-k111": "L34-N01-v111", "L34-N01-k112": "L34-N01-v112", "L34-N01-k113": "L34-N01-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N01-e101", FString2: "L34-N01-e102", FString3: "L34-N01-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N01-e111", FString2: "L34-N01-e112", FString3: "L34-N01-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N01-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N01-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N01-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N01-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N01-i111", IgnoreFieldsString2: "L34-N01-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N01-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N01-i131"},
					}
					level35N01 = &pbinline.MessageLevel35{
						Level35FString1:                   "L35-N01-s101",
						Level35FString2:                   "L35-N01-s111",
						Level35PString1:                   utils.PointerString("L35-N01-p101"),
						Level35PString2:                   utils.PointerString("L35-N01-p111"),
						Level35RString1:                   []string{"L35-N01-r101", "L35-N01-r102", "L35-N01-r103"},
						Level35RString2:                   []string{"L35-N01-r111", "L35-N01-r112", "L35-N01-r113"},
						Level35MString1:                   map[string]string{"L35-N01-k101": "L35-N01-v101", "L35-N01-k102": "L35-N01-v102", "L35-N01-k103": "L35-N01-v103"},
						Level35MString2:                   map[string]string{"L35-N01-k111": "L35-N01-v111", "L35-N01-k112": "L35-N01-v112", "L35-N01-k113": "L35-N01-v113"},
						Level35FMessageExtern1:            &pbexternal.Message1{FString1: "L35-N01-e101", FString2: "L35-N01-e102", FString3: "L35-N01-e103"},
						Level35FMessageExtern2:            &pbexternal.Message1{FString1: "L35-N01-e111", FString2: "L35-N01-e112", FString3: "L35-N01-e113"},
						Level35OneOfExtern1:               &pbinline.MessageLevel35_Level35One1String1{Level35One1String1: "L35-N01-es101"},
						Level35OneOfInline1:               &pbinline.MessageLevel35_Level35One2String1{Level35One2String1: "L35-N01-is101"},
						Level35OneOfOmitempty1:            &pbinline.MessageLevel35_Level35One3String1{Level35One3String1: "L35-N01-o101"},
						Level35FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level35FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L35-N01-i101"},
						Level35FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L35-N01-i111", IgnoreFieldsString2: "L35-N01-i112"},
						Level35OneOfIgnoreSelf1:           &pbinline.MessageLevel35_Level35One4String1{Level35One4String1: "L35-N01-i121"},
						Level35OneOfIgnoreParts1:          &pbinline.MessageLevel35_Level35One5String1{Level35One5String1: "L35-N01-i131"},
					}
					level34N02 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N02-s101",
						Level34FString2:                   "L34-N02-s111",
						Level34PString1:                   utils.PointerString("L34-N02-p101"),
						Level34PString2:                   utils.PointerString("L34-N02-p111"),
						Level34RString1:                   []string{"L34-N02-r101", "L34-N02-r102", "L34-N02-r103"},
						Level34RString2:                   []string{"L34-N02-r111", "L34-N02-r112", "L34-N02-r113"},
						Level34MString1:                   map[string]string{"L34-N02-k101": "L34-N02-v101", "L34-N02-k102": "L34-N02-v102", "L34-N02-k103": "L34-N02-v103"},
						Level34MString2:                   map[string]string{"L34-N02-k111": "L34-N02-v111", "L34-N02-k112": "L34-N02-v112", "L34-N02-k113": "L34-N02-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N02-e101", FString2: "L34-N02-e102", FString3: "L34-N02-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N02-e111", FString2: "L34-N02-e112", FString3: "L34-N02-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N02-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N02-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N02-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N02-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N02-i111", IgnoreFieldsString2: "L34-N02-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N02-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N02-i131"},
					}
					level36N01 = &pbinline.MessageLevel36{
						Level36FString1:                   "L36-N01-s101",
						Level36FString2:                   "L36-N01-s111",
						Level36PString1:                   utils.PointerString("L36-N01-p101"),
						Level36PString2:                   utils.PointerString("L36-N01-p111"),
						Level36RString1:                   []string{"L36-N01-r101", "L36-N01-r102", "L36-N01-r103"},
						Level36RString2:                   []string{"L36-N01-r111", "L36-N01-r112", "L36-N01-r113"},
						Level36MString1:                   map[string]string{"L36-N01-k101": "L36-N01-v101", "L36-N01-k102": "L36-N01-v102", "L36-N01-k103": "L36-N01-v103"},
						Level36MString2:                   map[string]string{"L36-N01-k111": "L36-N01-v111", "L36-N01-k112": "L36-N01-v112", "L36-N01-k113": "L36-N01-v113"},
						Level36FMessageExtern1:            &pbexternal.Message1{FString1: "L36-N01-e101", FString2: "L36-N01-e102", FString3: "L36-N01-e103"},
						Level36FMessageExtern2:            &pbexternal.Message1{FString1: "L36-N01-e111", FString2: "L36-N01-e112", FString3: "L36-N01-e113"},
						Level36OneOfExtern1:               &pbinline.MessageLevel36_Level36One1String1{Level36One1String1: "L36-N01-es101"},
						Level36OneOfInline1:               &pbinline.MessageLevel36_Level36One2String1{Level36One2String1: "L36-N01-is101"},
						Level36OneOfOmitempty1:            &pbinline.MessageLevel36_Level36One3String1{Level36One3String1: "L36-N01-o101"},
						Level36FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level36FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L36-N01-i101"},
						Level36FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L36-N01-i111", IgnoreFieldsString2: "L36-N01-i112"},
						Level36OneOfIgnoreSelf1:           &pbinline.MessageLevel36_Level36One4String1{Level36One4String1: "L36-N01-i121"},
						Level36OneOfIgnoreParts1:          &pbinline.MessageLevel36_Level36One5String1{Level36One5String1: "L36-N01-i131"},
					}
				}
				level28N01 = &pbinline.MessageLevel28{
					Level28FString1:                   "L28-N01-s101",
					Level28FString2:                   "L28-N01-s111",
					Level28PString1:                   utils.PointerString("L28-N01-p101"),
					Level28PString2:                   utils.PointerString("L28-N01-p111"),
					Level28RString1:                   []string{"L28-N01-r101", "L28-N01-r102", "L28-N01-r103"},
					Level28RString2:                   []string{"L28-N01-r111", "L28-N01-r112", "L28-N01-r113"},
					Level28MString1:                   map[string]string{"L28-N01-k101": "L28-N01-v101", "L28-N01-k102": "L28-N01-v102", "L28-N01-k103": "L28-N01-v103"},
					Level28MString2:                   map[string]string{"L28-N01-k111": "L28-N01-v111", "L28-N01-k112": "L28-N01-v112", "L28-N01-k113": "L28-N01-v113"},
					Level28FMessageExtern1:            &pbexternal.Message1{FString1: "L28-N01-e101", FString2: "L28-N01-e102", FString3: "L28-N01-e103"},
					Level28FMessageExtern2:            &pbexternal.Message1{FString1: "L28-N01-e111", FString2: "L28-N01-e112", FString3: "L28-N01-e113"},
					Level28FMessageInline34:           level34N01,
					Level28FMessageInline35:           level35N01,
					Level28OneOfExtern1:               &pbinline.MessageLevel28_Level28One1MessageInline34{Level28One1MessageInline34: level34N02},
					Level28OneOfInline1:               &pbinline.MessageLevel28_Level28One2MessageInline36{Level28One2MessageInline36: level36N01},
					Level28OneOfOmitempty1:            &pbinline.MessageLevel28_Level28One3String1{Level28One3String1: "L28-N01-o101"},
					Level28FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level28FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L28-N01-i101"},
					Level28FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L28-N01-i111", IgnoreFieldsString2: "L28-N01-i112"},
					Level28OneOfIgnoreSelf1:           &pbinline.MessageLevel28_Level28One4String1{Level28One4String1: "L28-N01-i121"},
					Level28OneOfIgnoreParts1:          &pbinline.MessageLevel28_Level28One5String1{Level28One5String1: "L28-N01-i131"},
				}
			}
			level16N01 = &pbinline.MessageLevel16{
				Level16FString1:                   "L16-N01-s101",
				Level16FString2:                   "L16-N01-s111",
				Level16PString1:                   utils.PointerString("L16-N01-p101"),
				Level16PString2:                   utils.PointerString("L16-N01-p111"),
				Level16RString1:                   []string{"L16-N01-r101", "L16-N01-r102", "L16-N01-r103"},
				Level16RString2:                   []string{"L16-N01-r111", "L16-N01-r112", "L16-N01-r113"},
				Level16MString1:                   map[string]string{"L16-N01-k101": "L16-N01-v101", "L16-N01-k102": "L16-N01-v102", "L16-N01-k103": "L16-N01-v103"},
				Level16MString2:                   map[string]string{"L16-N01-k111": "L16-N01-v111", "L16-N01-k112": "L16-N01-v112", "L16-N01-k113": "L16-N01-v113"},
				Level16FMessageExtern1:            &pbexternal.Message1{FString1: "L16-N01-e101", FString2: "L16-N01-e102", FString3: "L16-N01-e103"},
				Level16FMessageExtern2:            &pbexternal.Message1{FString1: "L16-N01-e111", FString2: "L16-N01-e112", FString3: "L16-N01-e113"},
				Level16FMessageInline26:           level26N01,
				Level16FMessageInline27:           level27N01,
				Level16OneOfExtern1:               &pbinline.MessageLevel16_Level16One1MessageInline26{Level16One1MessageInline26: level26N02},
				Level16OneOfInline1:               &pbinline.MessageLevel16_Level16One2MessageInline28{Level16One2MessageInline28: level28N01},
				Level16OneOfOmitempty1:            &pbinline.MessageLevel16_Level16One3String1{Level16One3String1: "L16-N01-o101"},
				Level16FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level16FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L16-N01-i101"},
				Level16FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L16-N01-i111", IgnoreFieldsString2: "L16-N01-i112"},
				Level16OneOfIgnoreSelf1:           &pbinline.MessageLevel16_Level16One4String1{Level16One4String1: "L16-N01-i121"},
				Level16OneOfIgnoreParts1:          &pbinline.MessageLevel16_Level16One5String1{Level16One5String1: "L16-N01-i131"},
			}
		}
		level06N01 = &pbinline.MessageLevel06{
			Level06FString1:                   "L06-N01-s101",
			Level06FString2:                   "L06-N01-s111",
			Level06PString1:                   utils.PointerString("L06-N01-p101"),
			Level06PString2:                   utils.PointerString("L06-N01-p111"),
			Level06RString1:                   []string{"L06-N01-r101", "L06-N01-r102", "L06-N01-r103"},
			Level06RString2:                   []string{"L06-N01-r111", "L06-N01-r112", "L06-N01-r113"},
			Level06MString1:                   map[string]string{"L06-N01-k101": "L06-N01-v101", "L06-N01-k102": "L06-N01-v102", "L06-N01-k103": "L06-N01-v103"},
			Level06MString2:                   map[string]string{"L06-N01-k111": "L06-N01-v111", "L06-N01-k112": "L06-N01-v112", "L06-N01-k113": "L06-N01-v113"},
			Level06FMessageExtern1:            &pbexternal.Message1{FString1: "L06-N01-e101", FString2: "L06-N01-e102", FString3: "L06-N01-e103"},
			Level06FMessageExtern2:            &pbexternal.Message1{FString1: "L06-N01-e111", FString2: "L06-N01-e112", FString3: "L06-N01-e113"},
			Level06FMessageInline14:           level14N01,
			Level06FMessageInline15:           level15N01,
			Level06OneOfExtern1:               &pbinline.MessageLevel06_Level06One1MessageInline14{Level06One1MessageInline14: level14N02},
			Level06OneOfInline1:               &pbinline.MessageLevel06_Level4One2MessageInline16{Level4One2MessageInline16: level16N01},
			Level06OneOfOmitempty1:            &pbinline.MessageLevel06_Level06One3String1{Level06One3String1: "L06-N01-o101"},
			Level06FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level06FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L06-N01-i101"},
			Level06FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L06-N01-i111", IgnoreFieldsString2: "L06-N01-i112"},
			Level06OneOfIgnoreSelf1:           &pbinline.MessageLevel06_Level06One4String1{Level06One4String1: "L06-N01-i121"},
			Level06OneOfIgnoreParts1:          &pbinline.MessageLevel06_Level06One5String1{Level06One5String1: "L06-N01-i131"},
		}
		level07N01 = &pbinline.MessageLevel07{
			Level07FString1:                   "L07-N01-s101",
			Level07FString2:                   "L07-N01-s111",
			Level07PString1:                   utils.PointerString("L07-N01-p101"),
			Level07PString2:                   utils.PointerString("L07-N01-p111"),
			Level07RString1:                   []string{"L07-N01-r101", "L07-N01-r102", "L07-N01-r103"},
			Level07RString2:                   []string{"L07-N01-r111", "L07-N01-r112", "L07-N01-r113"},
			Level07MString1:                   map[string]string{"L07-N01-k101": "L07-N01-v101", "L07-N01-k102": "L07-N01-v102", "L07-N01-k103": "L07-N01-v103"},
			Level07MString2:                   map[string]string{"L07-N01-k111": "L07-N01-v111", "L07-N01-k112": "L07-N01-v112", "L07-N01-k113": "L07-N01-v113"},
			Level07FMessageExtern1:            &pbexternal.Message1{FString1: "L07-N01-e101", FString2: "L07-N01-e102", FString3: "L07-N01-e103"},
			Level07FMessageExtern2:            &pbexternal.Message1{FString1: "L07-N01-e111", FString2: "L07-N01-e112", FString3: "L07-N01-e113"},
			Level07OneOfExtern1:               &pbinline.MessageLevel07_Level07One1String1{Level07One1String1: "L07-N01-es101"},
			Level07OneOfInline1:               &pbinline.MessageLevel07_Level07One2String1{Level07One2String1: "L07-N01-is101"},
			Level07OneOfOmitempty1:            &pbinline.MessageLevel07_Level07One3String1{Level07One3String1: "L07-N01-o101"},
			Level07FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level07FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L07-N01-i101"},
			Level07FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L07-N01-i111", IgnoreFieldsString2: "L07-N01-i112"},
			Level07OneOfIgnoreSelf1:           &pbinline.MessageLevel07_Level07One4String1{Level07One4String1: "L07-N01-i121"},
			Level07OneOfIgnoreParts1:          &pbinline.MessageLevel07_Level07One5String1{Level07One5String1: "L07-N01-i131"},
		}
		{ // message in level06N02
			{ // message for level14N03
				{ // message for level22N05
					level38N09 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N09-s101",
						Level38FString2:                   "L38-N09-s111",
						Level38PString1:                   utils.PointerString("L38-N09-p101"),
						Level38PString2:                   utils.PointerString("L38-N09-p111"),
						Level38RString1:                   []string{"L38-N09-r101", "L38-N09-r102", "L38-N09-r103"},
						Level38RString2:                   []string{"L38-N09-r111", "L38-N09-r112", "L38-N09-r113"},
						Level38MString1:                   map[string]string{"L38-N09-k101": "L38-N09-v101", "L38-N09-k102": "L38-N09-v102", "L38-N09-k103": "L38-N09-v103"},
						Level38MString2:                   map[string]string{"L38-N09-k111": "L38-N09-v111", "L38-N09-k112": "L38-N09-v112", "L38-N09-k113": "L38-N09-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N09-e101", FString2: "L38-N09-e102", FString3: "L38-N09-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N09-e111", FString2: "L38-N09-e112", FString3: "L38-N09-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N09-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N09-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N09-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N09-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N09-i111", IgnoreFieldsString2: "L38-N09-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N09-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N09-i131"},
					}
					level39N05 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N05-s101",
						Level39FString2:                   "L39-N05-s111",
						Level39PString1:                   utils.PointerString("L39-N05-p101"),
						Level39PString2:                   utils.PointerString("L39-N05-p111"),
						Level39RString1:                   []string{"L39-N05-r101", "L39-N05-r102", "L39-N05-r103"},
						Level39RString2:                   []string{"L39-N05-r111", "L39-N05-r112", "L39-N05-r113"},
						Level39MString1:                   map[string]string{"L39-N05-k101": "L39-N05-v101", "L39-N05-k102": "L39-N05-v102", "L39-N05-k103": "L39-N05-v103"},
						Level39MString2:                   map[string]string{"L39-N05-k111": "L39-N05-v111", "L39-N05-k112": "L39-N05-v112", "L39-N05-k113": "L39-N05-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N05-e101", FString2: "L39-N05-e102", FString3: "L39-N05-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N05-e111", FString2: "L39-N05-e112", FString3: "L39-N05-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N05-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N05-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N05-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N05-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N05-i111", IgnoreFieldsString2: "L39-N05-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N05-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N05-i131"},
					}
					level38N10 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N10-s101",
						Level38FString2:                   "L38-N10-s111",
						Level38PString1:                   utils.PointerString("L38-N10-p101"),
						Level38PString2:                   utils.PointerString("L38-N10-p111"),
						Level38RString1:                   []string{"L38-N10-r101", "L38-N10-r102", "L38-N10-r103"},
						Level38RString2:                   []string{"L38-N10-r111", "L38-N10-r112", "L38-N10-r113"},
						Level38MString1:                   map[string]string{"L38-N10-k101": "L38-N10-v101", "L38-N10-k102": "L38-N10-v102", "L38-N10-k103": "L38-N10-v103"},
						Level38MString2:                   map[string]string{"L38-N10-k111": "L38-N10-v111", "L38-N10-k112": "L38-N10-v112", "L38-N10-k113": "L38-N10-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N10-e101", FString2: "L38-N10-e102", FString3: "L38-N10-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N10-e111", FString2: "L38-N10-e112", FString3: "L38-N10-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N10-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N10-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N10-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N10-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N10-i111", IgnoreFieldsString2: "L38-N10-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N10-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N10-i131"},
					}
					level40N05 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N05-s101",
						Level40FString2:                   "L40-N05-s111",
						Level40PString1:                   utils.PointerString("L40-N05-p101"),
						Level40PString2:                   utils.PointerString("L40-N05-p111"),
						Level40RString1:                   []string{"L40-N05-r101", "L40-N05-r102", "L40-N05-r103"},
						Level40RString2:                   []string{"L40-N05-r111", "L40-N05-r112", "L40-N05-r113"},
						Level40MString1:                   map[string]string{"L40-N05-k101": "L40-N05-v101", "L40-N05-k102": "L40-N05-v102", "L40-N05-k103": "L40-N05-v103"},
						Level40MString2:                   map[string]string{"L40-N05-k111": "L40-N05-v111", "L40-N05-k112": "L40-N05-v112", "L40-N05-k113": "L40-N05-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N05-e101", FString2: "L40-N05-e102", FString3: "L40-N05-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N05-e111", FString2: "L40-N05-e112", FString3: "L40-N05-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N05-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N05-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N05-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N05-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N05-i111", IgnoreFieldsString2: "L40-N05-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N05-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N05-i131"},
					}
				}
				level22N05 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N05-s101",
					Level22FString2:                   "L22-N05-s111",
					Level22PString1:                   utils.PointerString("L22-N05-p101"),
					Level22PString2:                   utils.PointerString("L22-N05-p111"),
					Level22RString1:                   []string{"L22-N05-r101", "L22-N05-r102", "L22-N05-r103"},
					Level22RString2:                   []string{"L22-N05-r111", "L22-N05-r112", "L22-N05-r113"},
					Level22MString1:                   map[string]string{"L22-N05-k101": "L22-N05-v101", "L22-N05-k102": "L22-N05-v102", "L22-N05-k103": "L22-N05-v103"},
					Level22MString2:                   map[string]string{"L22-N05-k111": "L22-N05-v111", "L22-N05-k112": "L22-N05-v112", "L22-N05-k113": "L22-N05-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N05-e101", FString2: "L22-N05-e102", FString3: "L22-N05-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N05-e111", FString2: "L22-N05-e112", FString3: "L22-N05-e113"},
					Level22FMessageInline38:           level38N09,
					Level22FMessageInline39:           level39N05,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N10},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N05},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N05-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N05-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N05-i111", IgnoreFieldsString2: "L22-N05-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N05-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N05-i131"},
				}
				level23N03 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N03-s101",
					Level23FString2:                   "L23-N03-s111",
					Level23PString1:                   utils.PointerString("L23-N03-p101"),
					Level23PString2:                   utils.PointerString("L23-N03-p111"),
					Level23RString1:                   []string{"L23-N03-r101", "L23-N03-r102", "L23-N03-r103"},
					Level23RString2:                   []string{"L23-N03-r111", "L23-N03-r112", "L23-N03-r113"},
					Level23MString1:                   map[string]string{"L23-N03-k101": "L23-N03-v101", "L23-N03-k102": "L23-N03-v102", "L23-N03-k103": "L23-N03-v103"},
					Level23MString2:                   map[string]string{"L23-N03-k111": "L23-N03-v111", "L23-N03-k112": "L23-N03-v112", "L23-N03-k113": "L23-N03-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N03-e101", FString2: "L23-N03-e102", FString3: "L23-N03-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N03-e111", FString2: "L23-N03-e112", FString3: "L23-N03-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N03-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N03-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N03-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N03-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N03-i111", IgnoreFieldsString2: "L23-N03-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N03-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N03-i131"},
				}
				{ // message for level22N06
					level38N11 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N11-s101",
						Level38FString2:                   "L38-N11-s111",
						Level38PString1:                   utils.PointerString("L38-N11-p101"),
						Level38PString2:                   utils.PointerString("L38-N11-p111"),
						Level38RString1:                   []string{"L38-N11-r101", "L38-N11-r102", "L38-N11-r103"},
						Level38RString2:                   []string{"L38-N11-r111", "L38-N11-r112", "L38-N11-r113"},
						Level38MString1:                   map[string]string{"L38-N11-k101": "L38-N11-v101", "L38-N11-k102": "L38-N11-v102", "L38-N11-k103": "L38-N11-v103"},
						Level38MString2:                   map[string]string{"L38-N11-k111": "L38-N11-v111", "L38-N11-k112": "L38-N11-v112", "L38-N11-k113": "L38-N11-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N11-e101", FString2: "L38-N11-e102", FString3: "L38-N11-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N11-e111", FString2: "L38-N11-e112", FString3: "L38-N11-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N11-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N11-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N11-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N11-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N11-i111", IgnoreFieldsString2: "L38-N11-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N11-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N11-i131"},
					}
					level39N06 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N06-s101",
						Level39FString2:                   "L39-N06-s111",
						Level39PString1:                   utils.PointerString("L39-N06-p101"),
						Level39PString2:                   utils.PointerString("L39-N06-p111"),
						Level39RString1:                   []string{"L39-N06-r101", "L39-N06-r102", "L39-N06-r103"},
						Level39RString2:                   []string{"L39-N06-r111", "L39-N06-r112", "L39-N06-r113"},
						Level39MString1:                   map[string]string{"L39-N06-k101": "L39-N06-v101", "L39-N06-k102": "L39-N06-v102", "L39-N06-k103": "L39-N06-v103"},
						Level39MString2:                   map[string]string{"L39-N06-k111": "L39-N06-v111", "L39-N06-k112": "L39-N06-v112", "L39-N06-k113": "L39-N06-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N06-e101", FString2: "L39-N06-e102", FString3: "L39-N06-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N06-e111", FString2: "L39-N06-e112", FString3: "L39-N06-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N06-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N06-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N06-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N06-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N06-i111", IgnoreFieldsString2: "L39-N06-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N06-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N06-i131"},
					}
					level38N12 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N12-s101",
						Level38FString2:                   "L38-N12-s111",
						Level38PString1:                   utils.PointerString("L38-N12-p101"),
						Level38PString2:                   utils.PointerString("L38-N12-p111"),
						Level38RString1:                   []string{"L38-N12-r101", "L38-N12-r102", "L38-N12-r103"},
						Level38RString2:                   []string{"L38-N12-r111", "L38-N12-r112", "L38-N12-r113"},
						Level38MString1:                   map[string]string{"L38-N12-k101": "L38-N12-v101", "L38-N12-k102": "L38-N12-v102", "L38-N12-k103": "L38-N12-v103"},
						Level38MString2:                   map[string]string{"L38-N12-k111": "L38-N12-v111", "L38-N12-k112": "L38-N12-v112", "L38-N12-k113": "L38-N12-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N12-e101", FString2: "L38-N12-e102", FString3: "L38-N12-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N12-e111", FString2: "L38-N12-e112", FString3: "L38-N12-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N12-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N12-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N12-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N12-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N12-i111", IgnoreFieldsString2: "L38-N12-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N12-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N12-i131"},
					}
					level40N06 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N06-s101",
						Level40FString2:                   "L40-N06-s111",
						Level40PString1:                   utils.PointerString("L40-N06-p101"),
						Level40PString2:                   utils.PointerString("L40-N06-p111"),
						Level40RString1:                   []string{"L40-N06-r101", "L40-N06-r102", "L40-N06-r103"},
						Level40RString2:                   []string{"L40-N06-r111", "L40-N06-r112", "L40-N06-r113"},
						Level40MString1:                   map[string]string{"L40-N06-k101": "L40-N06-v101", "L40-N06-k102": "L40-N06-v102", "L40-N06-k103": "L40-N06-v103"},
						Level40MString2:                   map[string]string{"L40-N06-k111": "L40-N06-v111", "L40-N06-k112": "L40-N06-v112", "L40-N06-k113": "L40-N06-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N06-e101", FString2: "L40-N06-e102", FString3: "L40-N06-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N06-e111", FString2: "L40-N06-e112", FString3: "L40-N06-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N06-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N06-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N06-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N06-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N06-i111", IgnoreFieldsString2: "L40-N06-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N06-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N06-i131"},
					}
				}
				level22N06 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N06-s101",
					Level22FString2:                   "L22-N06-s111",
					Level22PString1:                   utils.PointerString("L22-N06-p101"),
					Level22PString2:                   utils.PointerString("L22-N06-p111"),
					Level22RString1:                   []string{"L22-N06-r101", "L22-N06-r102", "L22-N06-r103"},
					Level22RString2:                   []string{"L22-N06-r111", "L22-N06-r112", "L22-N06-r113"},
					Level22MString1:                   map[string]string{"L22-N06-k101": "L22-N06-v101", "L22-N06-k102": "L22-N06-v102", "L22-N06-k103": "L22-N06-v103"},
					Level22MString2:                   map[string]string{"L22-N06-k111": "L22-N06-v111", "L22-N06-k112": "L22-N06-v112", "L22-N06-k113": "L22-N06-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N06-e101", FString2: "L22-N06-e102", FString3: "L22-N06-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N06-e111", FString2: "L22-N06-e112", FString3: "L22-N06-e113"},
					Level22FMessageInline38:           level38N11,
					Level22FMessageInline39:           level39N06,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N12},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N06},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N06-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N06-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N06-i111", IgnoreFieldsString2: "L22-N06-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N06-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N06-i131"},
				}
				level24N03 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N03-s101",
					Level24FString2:                   "L24-N03-s111",
					Level24PString1:                   utils.PointerString("L24-N03-p101"),
					Level24PString2:                   utils.PointerString("L24-N03-p111"),
					Level24RString1:                   []string{"L24-N03-r101", "L24-N03-r102", "L24-N03-r103"},
					Level24RString2:                   []string{"L24-N03-r111", "L24-N03-r112", "L24-N03-r113"},
					Level24MString1:                   map[string]string{"L24-N03-k101": "L24-N03-v101", "L24-N03-k102": "L24-N03-v102", "L24-N03-k103": "L24-N03-v103"},
					Level24MString2:                   map[string]string{"L24-N03-k111": "L24-N03-v111", "L24-N03-k112": "L24-N03-v112", "L24-N03-k113": "L24-N03-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N03-e101", FString2: "L24-N03-e102", FString3: "L24-N03-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N03-e111", FString2: "L24-N03-e112", FString3: "L24-N03-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N03-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N03-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N03-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N03-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N03-i111", IgnoreFieldsString2: "L24-N03-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N03-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N03-i131"},
				}
			}
			level14N03 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N03-s101",
				Level14FString2:                   "L14-N03-s111",
				Level14PString1:                   utils.PointerString("L14-N03-p101"),
				Level14PString2:                   utils.PointerString("L14-N03-p111"),
				Level14RString1:                   []string{"L14-N03-r101", "L14-N03-r102", "L14-N03-r103"},
				Level14RString2:                   []string{"L14-N03-r111", "L14-N03-r112", "L14-N03-r113"},
				Level14MString1:                   map[string]string{"L14-N03-k101": "L14-N03-v101", "L14-N03-k102": "L14-N03-v102", "L14-N03-k103": "L14-N03-v103"},
				Level14MString2:                   map[string]string{"L14-N03-k111": "L14-N03-v111", "L14-N03-k112": "L14-N03-v112", "L14-N03-k113": "L14-N03-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N03-e101", FString2: "L14-N03-e102", FString3: "L14-N03-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N03-e111", FString2: "L14-N03-e112", FString3: "L14-N03-e113"},
				Level14FMessageInline22:           level22N05,
				Level14FMessageInline23:           level23N03,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N06},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N03},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N03-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N03-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N03-i111", IgnoreFieldsString2: "L14-N03-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N03-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N03-i131"},
			}
			level15N02 = &pbinline.MessageLevel15{
				Level15FString1:                   "L15-N02-s101",
				Level15FString2:                   "L15-N02-s111",
				Level15PString1:                   utils.PointerString("L15-N02-p101"),
				Level15PString2:                   utils.PointerString("L15-N02-p111"),
				Level15RString1:                   []string{"L15-N02-r101", "L15-N02-r102", "L15-N02-r103"},
				Level15RString2:                   []string{"L15-N02-r111", "L15-N02-r112", "L15-N02-r113"},
				Level15MString1:                   map[string]string{"L15-N02-k101": "L15-N02-v101", "L15-N02-k102": "L15-N02-v102", "L15-N02-k103": "L15-N02-v103"},
				Level15MString2:                   map[string]string{"L15-N02-k111": "L15-N02-v111", "L15-N02-k112": "L15-N02-v112", "L15-N02-k113": "L15-N02-v113"},
				Level15FMessageExtern1:            &pbexternal.Message1{FString1: "L15-N02-e101", FString2: "L15-N02-e102", FString3: "L15-N02-e103"},
				Level15FMessageExtern2:            &pbexternal.Message1{FString1: "L15-N02-e111", FString2: "L15-N02-e112", FString3: "L15-N02-e113"},
				Level15OneOfExtern1:               &pbinline.MessageLevel15_Level15One1String1{Level15One1String1: "L15-N02-es101"},
				Level15OneOfInline1:               &pbinline.MessageLevel15_Level15One2String1{Level15One2String1: "L15-N02-is101"},
				Level15OneOfOmitempty1:            &pbinline.MessageLevel15_Level15One3String1{Level15One3String1: "L15-N02-o101"},
				Level15FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level15FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L15-N02-i101"},
				Level15FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L15-N02-i111", IgnoreFieldsString2: "L15-N02-i112"},
				Level15OneOfIgnoreSelf1:           &pbinline.MessageLevel15_Level15One4String1{Level15One4String1: "L15-N02-i121"},
				Level15OneOfIgnoreParts1:          &pbinline.MessageLevel15_Level15One5String1{Level15One5String1: "L15-N02-i131"},
			}
			{ // message for level14N04
				{ // message for level22N07
					level38N13 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N13-s101",
						Level38FString2:                   "L38-N13-s111",
						Level38PString1:                   utils.PointerString("L38-N13-p101"),
						Level38PString2:                   utils.PointerString("L38-N13-p111"),
						Level38RString1:                   []string{"L38-N13-r101", "L38-N13-r102", "L38-N13-r103"},
						Level38RString2:                   []string{"L38-N13-r111", "L38-N13-r112", "L38-N13-r113"},
						Level38MString1:                   map[string]string{"L38-N13-k101": "L38-N13-v101", "L38-N13-k102": "L38-N13-v102", "L38-N13-k103": "L38-N13-v103"},
						Level38MString2:                   map[string]string{"L38-N13-k111": "L38-N13-v111", "L38-N13-k112": "L38-N13-v112", "L38-N13-k113": "L38-N13-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N13-e101", FString2: "L38-N13-e102", FString3: "L38-N13-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N13-e111", FString2: "L38-N13-e112", FString3: "L38-N13-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N13-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N13-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N13-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N13-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N13-i111", IgnoreFieldsString2: "L38-N13-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N13-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N13-i131"},
					}
					level39N07 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N07-s101",
						Level39FString2:                   "L39-N07-s111",
						Level39PString1:                   utils.PointerString("L39-N07-p101"),
						Level39PString2:                   utils.PointerString("L39-N07-p111"),
						Level39RString1:                   []string{"L39-N07-r101", "L39-N07-r102", "L39-N07-r103"},
						Level39RString2:                   []string{"L39-N07-r111", "L39-N07-r112", "L39-N07-r113"},
						Level39MString1:                   map[string]string{"L39-N07-k101": "L39-N07-v101", "L39-N07-k102": "L39-N07-v102", "L39-N07-k103": "L39-N07-v103"},
						Level39MString2:                   map[string]string{"L39-N07-k111": "L39-N07-v111", "L39-N07-k112": "L39-N07-v112", "L39-N07-k113": "L39-N07-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N07-e101", FString2: "L39-N07-e102", FString3: "L39-N07-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N07-e111", FString2: "L39-N07-e112", FString3: "L39-N07-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N07-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N07-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N07-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N07-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N07-i111", IgnoreFieldsString2: "L39-N07-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N07-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N07-i131"},
					}
					level38N14 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N14-s101",
						Level38FString2:                   "L38-N14-s111",
						Level38PString1:                   utils.PointerString("L38-N14-p101"),
						Level38PString2:                   utils.PointerString("L38-N14-p111"),
						Level38RString1:                   []string{"L38-N14-r101", "L38-N14-r102", "L38-N14-r103"},
						Level38RString2:                   []string{"L38-N14-r111", "L38-N14-r112", "L38-N14-r113"},
						Level38MString1:                   map[string]string{"L38-N14-k101": "L38-N14-v101", "L38-N14-k102": "L38-N14-v102", "L38-N14-k103": "L38-N14-v103"},
						Level38MString2:                   map[string]string{"L38-N14-k111": "L38-N14-v111", "L38-N14-k112": "L38-N14-v112", "L38-N14-k113": "L38-N14-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N14-e101", FString2: "L38-N14-e102", FString3: "L38-N14-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N14-e111", FString2: "L38-N14-e112", FString3: "L38-N14-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N14-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N14-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N14-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N14-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N14-i111", IgnoreFieldsString2: "L38-N14-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N14-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N14-i131"},
					}
					level40N07 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N07-s101",
						Level40FString2:                   "L40-N07-s111",
						Level40PString1:                   utils.PointerString("L40-N07-p101"),
						Level40PString2:                   utils.PointerString("L40-N07-p111"),
						Level40RString1:                   []string{"L40-N07-r101", "L40-N07-r102", "L40-N07-r103"},
						Level40RString2:                   []string{"L40-N07-r111", "L40-N07-r112", "L40-N07-r113"},
						Level40MString1:                   map[string]string{"L40-N07-k101": "L40-N07-v101", "L40-N07-k102": "L40-N07-v102", "L40-N07-k103": "L40-N07-v103"},
						Level40MString2:                   map[string]string{"L40-N07-k111": "L40-N07-v111", "L40-N07-k112": "L40-N07-v112", "L40-N07-k113": "L40-N07-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N07-e101", FString2: "L40-N07-e102", FString3: "L40-N07-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N07-e111", FString2: "L40-N07-e112", FString3: "L40-N07-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N07-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N07-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N07-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N07-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N07-i111", IgnoreFieldsString2: "L40-N07-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N07-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N07-i131"},
					}
				}
				level22N07 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N07-s101",
					Level22FString2:                   "L22-N07-s111",
					Level22PString1:                   utils.PointerString("L22-N07-p101"),
					Level22PString2:                   utils.PointerString("L22-N07-p111"),
					Level22RString1:                   []string{"L22-N07-r101", "L22-N07-r102", "L22-N07-r103"},
					Level22RString2:                   []string{"L22-N07-r111", "L22-N07-r112", "L22-N07-r113"},
					Level22MString1:                   map[string]string{"L22-N07-k101": "L22-N07-v101", "L22-N07-k102": "L22-N07-v102", "L22-N07-k103": "L22-N07-v103"},
					Level22MString2:                   map[string]string{"L22-N07-k111": "L22-N07-v111", "L22-N07-k112": "L22-N07-v112", "L22-N07-k113": "L22-N07-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N07-e101", FString2: "L22-N07-e102", FString3: "L22-N07-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N07-e111", FString2: "L22-N07-e112", FString3: "L22-N07-e113"},
					Level22FMessageInline38:           level38N13,
					Level22FMessageInline39:           level39N07,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N14},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N07},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N07-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N07-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N07-i111", IgnoreFieldsString2: "L22-N07-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N07-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N07-i131"},
				}
				level23N04 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N04-s101",
					Level23FString2:                   "L23-N04-s111",
					Level23PString1:                   utils.PointerString("L23-N04-p101"),
					Level23PString2:                   utils.PointerString("L23-N04-p111"),
					Level23RString1:                   []string{"L23-N04-r101", "L23-N04-r102", "L23-N04-r103"},
					Level23RString2:                   []string{"L23-N04-r111", "L23-N04-r112", "L23-N04-r113"},
					Level23MString1:                   map[string]string{"L23-N04-k101": "L23-N04-v101", "L23-N04-k102": "L23-N04-v102", "L23-N04-k103": "L23-N04-v103"},
					Level23MString2:                   map[string]string{"L23-N04-k111": "L23-N04-v111", "L23-N04-k112": "L23-N04-v112", "L23-N04-k113": "L23-N04-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N04-e101", FString2: "L23-N04-e102", FString3: "L23-N04-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N04-e111", FString2: "L23-N04-e112", FString3: "L23-N04-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N04-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N04-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N04-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N04-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N04-i111", IgnoreFieldsString2: "L23-N04-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N04-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N04-i131"},
				}
				{ // message for level22N08
					level38N15 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N15-s101",
						Level38FString2:                   "L38-N15-s111",
						Level38PString1:                   utils.PointerString("L38-N15-p101"),
						Level38PString2:                   utils.PointerString("L38-N15-p111"),
						Level38RString1:                   []string{"L38-N15-r101", "L38-N15-r102", "L38-N15-r103"},
						Level38RString2:                   []string{"L38-N15-r111", "L38-N15-r112", "L38-N15-r113"},
						Level38MString1:                   map[string]string{"L38-N15-k101": "L38-N15-v101", "L38-N15-k102": "L38-N15-v102", "L38-N15-k103": "L38-N15-v103"},
						Level38MString2:                   map[string]string{"L38-N15-k111": "L38-N15-v111", "L38-N15-k112": "L38-N15-v112", "L38-N15-k113": "L38-N15-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N15-e101", FString2: "L38-N15-e102", FString3: "L38-N15-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N15-e111", FString2: "L38-N15-e112", FString3: "L38-N15-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N15-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N15-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N15-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N15-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N15-i111", IgnoreFieldsString2: "L38-N15-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N15-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N15-i131"},
					}
					level39N08 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N08-s101",
						Level39FString2:                   "L39-N08-s111",
						Level39PString1:                   utils.PointerString("L39-N08-p101"),
						Level39PString2:                   utils.PointerString("L39-N08-p111"),
						Level39RString1:                   []string{"L39-N08-r101", "L39-N08-r102", "L39-N08-r103"},
						Level39RString2:                   []string{"L39-N08-r111", "L39-N08-r112", "L39-N08-r113"},
						Level39MString1:                   map[string]string{"L39-N08-k101": "L39-N08-v101", "L39-N08-k102": "L39-N08-v102", "L39-N08-k103": "L39-N08-v103"},
						Level39MString2:                   map[string]string{"L39-N08-k111": "L39-N08-v111", "L39-N08-k112": "L39-N08-v112", "L39-N08-k113": "L39-N08-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N08-e101", FString2: "L39-N08-e102", FString3: "L39-N08-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N08-e111", FString2: "L39-N08-e112", FString3: "L39-N08-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N08-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N08-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N08-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N08-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N08-i111", IgnoreFieldsString2: "L39-N08-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N08-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N08-i131"},
					}
					level38N16 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N16-s101",
						Level38FString2:                   "L38-N16-s111",
						Level38PString1:                   utils.PointerString("L38-N16-p101"),
						Level38PString2:                   utils.PointerString("L38-N16-p111"),
						Level38RString1:                   []string{"L38-N16-r101", "L38-N16-r102", "L38-N16-r103"},
						Level38RString2:                   []string{"L38-N16-r111", "L38-N16-r112", "L38-N16-r113"},
						Level38MString1:                   map[string]string{"L38-N16-k101": "L38-N16-v101", "L38-N16-k102": "L38-N16-v102", "L38-N16-k103": "L38-N16-v103"},
						Level38MString2:                   map[string]string{"L38-N16-k111": "L38-N16-v111", "L38-N16-k112": "L38-N16-v112", "L38-N16-k113": "L38-N16-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N16-e101", FString2: "L38-N16-e102", FString3: "L38-N16-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N16-e111", FString2: "L38-N16-e112", FString3: "L38-N16-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N16-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N16-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N16-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N16-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N16-i111", IgnoreFieldsString2: "L38-N16-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N16-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N16-i131"},
					}
					level40N08 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N08-s101",
						Level40FString2:                   "L40-N08-s111",
						Level40PString1:                   utils.PointerString("L40-N08-p101"),
						Level40PString2:                   utils.PointerString("L40-N08-p111"),
						Level40RString1:                   []string{"L40-N08-r101", "L40-N08-r102", "L40-N08-r103"},
						Level40RString2:                   []string{"L40-N08-r111", "L40-N08-r112", "L40-N08-r113"},
						Level40MString1:                   map[string]string{"L40-N08-k101": "L40-N08-v101", "L40-N08-k102": "L40-N08-v102", "L40-N08-k103": "L40-N08-v103"},
						Level40MString2:                   map[string]string{"L40-N08-k111": "L40-N08-v111", "L40-N08-k112": "L40-N08-v112", "L40-N08-k113": "L40-N08-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N08-e101", FString2: "L40-N08-e102", FString3: "L40-N08-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N08-e111", FString2: "L40-N08-e112", FString3: "L40-N08-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N08-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N08-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N08-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N08-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N08-i111", IgnoreFieldsString2: "L40-N08-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N08-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N08-i131"},
					}
				}
				level22N08 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N08-s101",
					Level22FString2:                   "L22-N08-s111",
					Level22PString1:                   utils.PointerString("L22-N08-p101"),
					Level22PString2:                   utils.PointerString("L22-N08-p111"),
					Level22RString1:                   []string{"L22-N08-r101", "L22-N08-r102", "L22-N08-r103"},
					Level22RString2:                   []string{"L22-N08-r111", "L22-N08-r112", "L22-N08-r113"},
					Level22MString1:                   map[string]string{"L22-N08-k101": "L22-N08-v101", "L22-N08-k102": "L22-N08-v102", "L22-N08-k103": "L22-N08-v103"},
					Level22MString2:                   map[string]string{"L22-N08-k111": "L22-N08-v111", "L22-N08-k112": "L22-N08-v112", "L22-N08-k113": "L22-N08-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N08-e101", FString2: "L22-N08-e102", FString3: "L22-N08-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N08-e111", FString2: "L22-N08-e112", FString3: "L22-N08-e113"},
					Level22FMessageInline38:           level38N15,
					Level22FMessageInline39:           level39N08,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N16},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N08},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N08-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N08-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N08-i111", IgnoreFieldsString2: "L22-N08-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N08-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N08-i131"},
				}
				level24N04 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N04-s101",
					Level24FString2:                   "L24-N04-s111",
					Level24PString1:                   utils.PointerString("L24-N04-p101"),
					Level24PString2:                   utils.PointerString("L24-N04-p111"),
					Level24RString1:                   []string{"L24-N04-r101", "L24-N04-r102", "L24-N04-r103"},
					Level24RString2:                   []string{"L24-N04-r111", "L24-N04-r112", "L24-N04-r113"},
					Level24MString1:                   map[string]string{"L24-N04-k101": "L24-N04-v101", "L24-N04-k102": "L24-N04-v102", "L24-N04-k103": "L24-N04-v103"},
					Level24MString2:                   map[string]string{"L24-N04-k111": "L24-N04-v111", "L24-N04-k112": "L24-N04-v112", "L24-N04-k113": "L24-N04-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N04-e101", FString2: "L24-N04-e102", FString3: "L24-N04-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N04-e111", FString2: "L24-N04-e112", FString3: "L24-N04-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N04-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N04-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N04-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N04-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N04-i111", IgnoreFieldsString2: "L24-N04-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N04-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N04-i131"},
				}
			}
			level14N04 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N04-s101",
				Level14FString2:                   "L14-N04-s111",
				Level14PString1:                   utils.PointerString("L14-N04-p101"),
				Level14PString2:                   utils.PointerString("L14-N04-p111"),
				Level14RString1:                   []string{"L14-N04-r101", "L14-N04-r102", "L14-N04-r103"},
				Level14RString2:                   []string{"L14-N04-r111", "L14-N04-r112", "L14-N04-r113"},
				Level14MString1:                   map[string]string{"L14-N04-k101": "L14-N04-v101", "L14-N04-k102": "L14-N04-v102", "L14-N04-k103": "L14-N04-v103"},
				Level14MString2:                   map[string]string{"L14-N04-k111": "L14-N04-v111", "L14-N04-k112": "L14-N04-v112", "L14-N04-k113": "L14-N04-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N04-e101", FString2: "L14-N04-e102", FString3: "L14-N04-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N04-e111", FString2: "L14-N04-e112", FString3: "L14-N04-e113"},
				Level14FMessageInline22:           level22N07,
				Level14FMessageInline23:           level23N04,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N08},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N04},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N04-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N04-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N04-i111", IgnoreFieldsString2: "L14-N04-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N04-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N04-i131"},
			}
			{ // message for level16N02
				{ // message for level26N03
					level30N05 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N05-s101",
						Level30FString2:                   "L30-N05-s111",
						Level30PString1:                   utils.PointerString("L30-N05-p101"),
						Level30PString2:                   utils.PointerString("L30-N05-p111"),
						Level30RString1:                   []string{"L30-N05-r101", "L30-N05-r102", "L30-N05-r103"},
						Level30RString2:                   []string{"L30-N05-r111", "L30-N05-r112", "L30-N05-r113"},
						Level30MString1:                   map[string]string{"L30-N05-k101": "L30-N05-v101", "L30-N05-k102": "L30-N05-v102", "L30-N05-k103": "L30-N05-v103"},
						Level30MString2:                   map[string]string{"L30-N05-k111": "L30-N05-v111", "L30-N05-k112": "L30-N05-v112", "L30-N05-k113": "L30-N05-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N05-e101", FString2: "L30-N05-e102", FString3: "L30-N05-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N05-e111", FString2: "L30-N05-e112", FString3: "L30-N05-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N05-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N05-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N05-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N05-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N05-i111", IgnoreFieldsString2: "L30-N05-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N05-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N05-i131"},
					}
					level31N03 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N03-s101",
						Level31FString2:                   "L31-N03-s111",
						Level31PString1:                   utils.PointerString("L31-N03-p101"),
						Level31PString2:                   utils.PointerString("L31-N03-p111"),
						Level31RString1:                   []string{"L31-N03-r101", "L31-N03-r102", "L31-N03-r103"},
						Level31RString2:                   []string{"L31-N03-r111", "L31-N03-r112", "L31-N03-r113"},
						Level31MString1:                   map[string]string{"L31-N03-k101": "L31-N03-v101", "L31-N03-k102": "L31-N03-v102", "L31-N03-k103": "L31-N03-v103"},
						Level31MString2:                   map[string]string{"L31-N03-k111": "L31-N03-v111", "L31-N03-k112": "L31-N03-v112", "L31-N03-k113": "L31-N03-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N03-e101", FString2: "L31-N03-e102", FString3: "L31-N03-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N03-e111", FString2: "L31-N03-e112", FString3: "L31-N03-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N03-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N03-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N03-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N03-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N03-i111", IgnoreFieldsString2: "L31-N03-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N03-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N03-i131"},
					}
					level30N06 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N06-s101",
						Level30FString2:                   "L30-N06-s111",
						Level30PString1:                   utils.PointerString("L30-N06-p101"),
						Level30PString2:                   utils.PointerString("L30-N06-p111"),
						Level30RString1:                   []string{"L30-N06-r101", "L30-N06-r102", "L30-N06-r103"},
						Level30RString2:                   []string{"L30-N06-r111", "L30-N06-r112", "L30-N06-r113"},
						Level30MString1:                   map[string]string{"L30-N06-k101": "L30-N06-v101", "L30-N06-k102": "L30-N06-v102", "L30-N06-k103": "L30-N06-v103"},
						Level30MString2:                   map[string]string{"L30-N06-k111": "L30-N06-v111", "L30-N06-k112": "L30-N06-v112", "L30-N06-k113": "L30-N06-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N06-e101", FString2: "L30-N06-e102", FString3: "L30-N06-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N06-e111", FString2: "L30-N06-e112", FString3: "L30-N06-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N06-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N06-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N06-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N06-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N06-i111", IgnoreFieldsString2: "L30-N06-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N06-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N06-i131"},
					}
					level32N03 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N03-s101",
						Level32FString2:                   "L32-N03-s111",
						Level32PString1:                   utils.PointerString("L32-N03-p101"),
						Level32PString2:                   utils.PointerString("L32-N03-p111"),
						Level32RString1:                   []string{"L32-N03-r101", "L32-N03-r102", "L32-N03-r103"},
						Level32RString2:                   []string{"L32-N03-r111", "L32-N03-r112", "L32-N03-r113"},
						Level32MString1:                   map[string]string{"L32-N03-k101": "L32-N03-v101", "L32-N03-k102": "L32-N03-v102", "L32-N03-k103": "L32-N03-v103"},
						Level32MString2:                   map[string]string{"L32-N03-k111": "L32-N03-v111", "L32-N03-k112": "L32-N03-v112", "L32-N03-k113": "L32-N03-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N03-e101", FString2: "L32-N03-e102", FString3: "L32-N03-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N03-e111", FString2: "L32-N03-e112", FString3: "L32-N03-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N03-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N03-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N03-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N03-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N03-i111", IgnoreFieldsString2: "L32-N03-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N03-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N03-i131"},
					}
				}
				level26N03 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N03-s101",
					Level26FString2:                   "L26-N03-s111",
					Level26PString1:                   utils.PointerString("L26-N03-p101"),
					Level26PString2:                   utils.PointerString("L26-N03-p111"),
					Level26RString1:                   []string{"L26-N03-r101", "L26-N03-r102", "L26-N03-r103"},
					Level26RString2:                   []string{"L26-N03-r111", "L26-N03-r112", "L26-N03-r113"},
					Level26MString1:                   map[string]string{"L26-N03-k101": "L26-N03-v101", "L26-N03-k102": "L26-N03-v102", "L26-N03-k103": "L26-N03-v103"},
					Level26MString2:                   map[string]string{"L26-N03-k111": "L26-N03-v111", "L26-N03-k112": "L26-N03-v112", "L26-N03-k113": "L26-N03-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N03-e101", FString2: "L26-N03-e102", FString3: "L26-N03-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N03-e111", FString2: "L26-N03-e112", FString3: "L26-N03-e113"},
					Level26FMessageInline30:           level30N05,
					Level26FMessageInline31:           level31N03,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N06},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N03},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N03-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N03-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N03-i111", IgnoreFieldsString2: "L26-N03-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N03-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N03-i131"},
				}
				level27N02 = &pbinline.MessageLevel27{
					Level27FString1:                   "L27-N02-s101",
					Level27FString2:                   "L27-N02-s111",
					Level27PString1:                   utils.PointerString("L27-N02-p101"),
					Level27PString2:                   utils.PointerString("L27-N02-p111"),
					Level27RString1:                   []string{"L27-N02-r101", "L27-N02-r102", "L27-N02-r103"},
					Level27RString2:                   []string{"L27-N02-r111", "L27-N02-r112", "L27-N02-r113"},
					Level27MString1:                   map[string]string{"L27-N02-k101": "L27-N02-v101", "L27-N02-k102": "L27-N02-v102", "L27-N02-k103": "L27-N02-v103"},
					Level27MString2:                   map[string]string{"L27-N02-k111": "L27-N02-v111", "L27-N02-k112": "L27-N02-v112", "L27-N02-k113": "L27-N02-v113"},
					Level27FMessageExtern1:            &pbexternal.Message1{FString1: "L27-N02-e101", FString2: "L27-N02-e102", FString3: "L27-N02-e103"},
					Level27FMessageExtern2:            &pbexternal.Message1{FString1: "L27-N02-e111", FString2: "L27-N02-e112", FString3: "L27-N02-e113"},
					Level27OneOfExtern1:               &pbinline.MessageLevel27_Level27One1String1{Level27One1String1: "L27-N02-es101"},
					Level27OneOfInline1:               &pbinline.MessageLevel27_Level27One2String1{Level27One2String1: "L27-N02-is101"},
					Level27OneOfOmitempty1:            &pbinline.MessageLevel27_Level27One3String1{Level27One3String1: "L27-N02-o101"},
					Level27FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level27FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L27-N02-i101"},
					Level27FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L27-N02-i111", IgnoreFieldsString2: "L27-N02-i112"},
					Level27OneOfIgnoreSelf1:           &pbinline.MessageLevel27_Level27One4String1{Level27One4String1: "L27-N02-i121"},
					Level27OneOfIgnoreParts1:          &pbinline.MessageLevel27_Level27One5String1{Level27One5String1: "L27-N02-i131"},
				}
				{ // message for level26N04
					level30N07 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N07-s101",
						Level30FString2:                   "L30-N07-s111",
						Level30PString1:                   utils.PointerString("L30-N07-p101"),
						Level30PString2:                   utils.PointerString("L30-N07-p111"),
						Level30RString1:                   []string{"L30-N07-r101", "L30-N07-r102", "L30-N07-r103"},
						Level30RString2:                   []string{"L30-N07-r111", "L30-N07-r112", "L30-N07-r113"},
						Level30MString1:                   map[string]string{"L30-N07-k101": "L30-N07-v101", "L30-N07-k102": "L30-N07-v102", "L30-N07-k103": "L30-N07-v103"},
						Level30MString2:                   map[string]string{"L30-N07-k111": "L30-N07-v111", "L30-N07-k112": "L30-N07-v112", "L30-N07-k113": "L30-N07-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N07-e101", FString2: "L30-N07-e102", FString3: "L30-N07-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N07-e111", FString2: "L30-N07-e112", FString3: "L30-N07-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N07-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N07-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N07-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N07-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N07-i111", IgnoreFieldsString2: "L30-N07-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N07-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N07-i131"},
					}
					level31N04 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N04-s101",
						Level31FString2:                   "L31-N04-s111",
						Level31PString1:                   utils.PointerString("L31-N04-p101"),
						Level31PString2:                   utils.PointerString("L31-N04-p111"),
						Level31RString1:                   []string{"L31-N04-r101", "L31-N04-r102", "L31-N04-r103"},
						Level31RString2:                   []string{"L31-N04-r111", "L31-N04-r112", "L31-N04-r113"},
						Level31MString1:                   map[string]string{"L31-N04-k101": "L31-N04-v101", "L31-N04-k102": "L31-N04-v102", "L31-N04-k103": "L31-N04-v103"},
						Level31MString2:                   map[string]string{"L31-N04-k111": "L31-N04-v111", "L31-N04-k112": "L31-N04-v112", "L31-N04-k113": "L31-N04-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N04-e101", FString2: "L31-N04-e102", FString3: "L31-N04-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N04-e111", FString2: "L31-N04-e112", FString3: "L31-N04-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N04-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N04-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N04-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N04-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N04-i111", IgnoreFieldsString2: "L31-N04-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N04-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N04-i131"},
					}
					level30N08 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N08-s101",
						Level30FString2:                   "L30-N08-s111",
						Level30PString1:                   utils.PointerString("L30-N08-p101"),
						Level30PString2:                   utils.PointerString("L30-N08-p111"),
						Level30RString1:                   []string{"L30-N08-r101", "L30-N08-r102", "L30-N08-r103"},
						Level30RString2:                   []string{"L30-N08-r111", "L30-N08-r112", "L30-N08-r113"},
						Level30MString1:                   map[string]string{"L30-N08-k101": "L30-N08-v101", "L30-N08-k102": "L30-N08-v102", "L30-N08-k103": "L30-N08-v103"},
						Level30MString2:                   map[string]string{"L30-N08-k111": "L30-N08-v111", "L30-N08-k112": "L30-N08-v112", "L30-N08-k113": "L30-N08-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N08-e101", FString2: "L30-N08-e102", FString3: "L30-N08-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N08-e111", FString2: "L30-N08-e112", FString3: "L30-N08-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N08-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N08-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N08-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N08-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N08-i111", IgnoreFieldsString2: "L30-N08-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N08-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N08-i131"},
					}
					level32N04 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N04-s101",
						Level32FString2:                   "L32-N04-s111",
						Level32PString1:                   utils.PointerString("L32-N04-p101"),
						Level32PString2:                   utils.PointerString("L32-N04-p111"),
						Level32RString1:                   []string{"L32-N04-r101", "L32-N04-r102", "L32-N04-r103"},
						Level32RString2:                   []string{"L32-N04-r111", "L32-N04-r112", "L32-N04-r113"},
						Level32MString1:                   map[string]string{"L32-N04-k101": "L32-N04-v101", "L32-N04-k102": "L32-N04-v102", "L32-N04-k103": "L32-N04-v103"},
						Level32MString2:                   map[string]string{"L32-N04-k111": "L32-N04-v111", "L32-N04-k112": "L32-N04-v112", "L32-N04-k113": "L32-N04-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N04-e101", FString2: "L32-N04-e102", FString3: "L32-N04-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N04-e111", FString2: "L32-N04-e112", FString3: "L32-N04-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N04-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N04-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N04-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N04-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N04-i111", IgnoreFieldsString2: "L32-N04-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N04-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N04-i131"},
					}
				}
				level26N04 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N04-s101",
					Level26FString2:                   "L26-N04-s111",
					Level26PString1:                   utils.PointerString("L26-N04-p101"),
					Level26PString2:                   utils.PointerString("L26-N04-p111"),
					Level26RString1:                   []string{"L26-N04-r101", "L26-N04-r102", "L26-N04-r103"},
					Level26RString2:                   []string{"L26-N04-r111", "L26-N04-r112", "L26-N04-r113"},
					Level26MString1:                   map[string]string{"L26-N04-k101": "L26-N04-v101", "L26-N04-k102": "L26-N04-v102", "L26-N04-k103": "L26-N04-v103"},
					Level26MString2:                   map[string]string{"L26-N04-k111": "L26-N04-v111", "L26-N04-k112": "L26-N04-v112", "L26-N04-k113": "L26-N04-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N04-e101", FString2: "L26-N04-e102", FString3: "L26-N04-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N04-e111", FString2: "L26-N04-e112", FString3: "L26-N04-e113"},
					Level26FMessageInline30:           level30N07,
					Level26FMessageInline31:           level31N04,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N08},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N04},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N04-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N04-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N04-i111", IgnoreFieldsString2: "L26-N04-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N04-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N04-i131"},
				}
				{ // message for level28N02
					level34N03 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N03-s101",
						Level34FString2:                   "L34-N03-s111",
						Level34PString1:                   utils.PointerString("L34-N03-p101"),
						Level34PString2:                   utils.PointerString("L34-N03-p111"),
						Level34RString1:                   []string{"L34-N03-r101", "L34-N03-r102", "L34-N03-r103"},
						Level34RString2:                   []string{"L34-N03-r111", "L34-N03-r112", "L34-N03-r113"},
						Level34MString1:                   map[string]string{"L34-N03-k101": "L34-N03-v101", "L34-N03-k102": "L34-N03-v102", "L34-N03-k103": "L34-N03-v103"},
						Level34MString2:                   map[string]string{"L34-N03-k111": "L34-N03-v111", "L34-N03-k112": "L34-N03-v112", "L34-N03-k113": "L34-N03-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N03-e101", FString2: "L34-N03-e102", FString3: "L34-N03-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N03-e111", FString2: "L34-N03-e112", FString3: "L34-N03-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N03-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N03-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N03-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N03-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N03-i111", IgnoreFieldsString2: "L34-N03-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N03-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N03-i131"},
					}
					level35N02 = &pbinline.MessageLevel35{
						Level35FString1:                   "L35-N02-s101",
						Level35FString2:                   "L35-N02-s111",
						Level35PString1:                   utils.PointerString("L35-N02-p101"),
						Level35PString2:                   utils.PointerString("L35-N02-p111"),
						Level35RString1:                   []string{"L35-N02-r101", "L35-N02-r102", "L35-N02-r103"},
						Level35RString2:                   []string{"L35-N02-r111", "L35-N02-r112", "L35-N02-r113"},
						Level35MString1:                   map[string]string{"L35-N02-k101": "L35-N02-v101", "L35-N02-k102": "L35-N02-v102", "L35-N02-k103": "L35-N02-v103"},
						Level35MString2:                   map[string]string{"L35-N02-k111": "L35-N02-v111", "L35-N02-k112": "L35-N02-v112", "L35-N02-k113": "L35-N02-v113"},
						Level35FMessageExtern1:            &pbexternal.Message1{FString1: "L35-N02-e101", FString2: "L35-N02-e102", FString3: "L35-N02-e103"},
						Level35FMessageExtern2:            &pbexternal.Message1{FString1: "L35-N02-e111", FString2: "L35-N02-e112", FString3: "L35-N02-e113"},
						Level35OneOfExtern1:               &pbinline.MessageLevel35_Level35One1String1{Level35One1String1: "L35-N02-es101"},
						Level35OneOfInline1:               &pbinline.MessageLevel35_Level35One2String1{Level35One2String1: "L35-N02-is101"},
						Level35OneOfOmitempty1:            &pbinline.MessageLevel35_Level35One3String1{Level35One3String1: "L35-N02-o101"},
						Level35FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level35FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L35-N02-i101"},
						Level35FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L35-N02-i111", IgnoreFieldsString2: "L35-N02-i112"},
						Level35OneOfIgnoreSelf1:           &pbinline.MessageLevel35_Level35One4String1{Level35One4String1: "L35-N02-i121"},
						Level35OneOfIgnoreParts1:          &pbinline.MessageLevel35_Level35One5String1{Level35One5String1: "L35-N02-i131"},
					}
					level34N04 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N04-s101",
						Level34FString2:                   "L34-N04-s111",
						Level34PString1:                   utils.PointerString("L34-N04-p101"),
						Level34PString2:                   utils.PointerString("L34-N04-p111"),
						Level34RString1:                   []string{"L34-N04-r101", "L34-N04-r102", "L34-N04-r103"},
						Level34RString2:                   []string{"L34-N04-r111", "L34-N04-r112", "L34-N04-r113"},
						Level34MString1:                   map[string]string{"L34-N04-k101": "L34-N04-v101", "L34-N04-k102": "L34-N04-v102", "L34-N04-k103": "L34-N04-v103"},
						Level34MString2:                   map[string]string{"L34-N04-k111": "L34-N04-v111", "L34-N04-k112": "L34-N04-v112", "L34-N04-k113": "L34-N04-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N04-e101", FString2: "L34-N04-e102", FString3: "L34-N04-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N04-e111", FString2: "L34-N04-e112", FString3: "L34-N04-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N04-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N04-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N04-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N04-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N04-i111", IgnoreFieldsString2: "L34-N04-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N04-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N04-i131"},
					}
					level36N02 = &pbinline.MessageLevel36{
						Level36FString1:                   "L36-N02-s101",
						Level36FString2:                   "L36-N02-s111",
						Level36PString1:                   utils.PointerString("L36-N02-p101"),
						Level36PString2:                   utils.PointerString("L36-N02-p111"),
						Level36RString1:                   []string{"L36-N02-r101", "L36-N02-r102", "L36-N02-r103"},
						Level36RString2:                   []string{"L36-N02-r111", "L36-N02-r112", "L36-N02-r113"},
						Level36MString1:                   map[string]string{"L36-N02-k101": "L36-N02-v101", "L36-N02-k102": "L36-N02-v102", "L36-N02-k103": "L36-N02-v103"},
						Level36MString2:                   map[string]string{"L36-N02-k111": "L36-N02-v111", "L36-N02-k112": "L36-N02-v112", "L36-N02-k113": "L36-N02-v113"},
						Level36FMessageExtern1:            &pbexternal.Message1{FString1: "L36-N02-e101", FString2: "L36-N02-e102", FString3: "L36-N02-e103"},
						Level36FMessageExtern2:            &pbexternal.Message1{FString1: "L36-N02-e111", FString2: "L36-N02-e112", FString3: "L36-N02-e113"},
						Level36OneOfExtern1:               &pbinline.MessageLevel36_Level36One1String1{Level36One1String1: "L36-N02-es101"},
						Level36OneOfInline1:               &pbinline.MessageLevel36_Level36One2String1{Level36One2String1: "L36-N02-is101"},
						Level36OneOfOmitempty1:            &pbinline.MessageLevel36_Level36One3String1{Level36One3String1: "L36-N02-o101"},
						Level36FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level36FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L36-N02-i101"},
						Level36FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L36-N02-i111", IgnoreFieldsString2: "L36-N02-i112"},
						Level36OneOfIgnoreSelf1:           &pbinline.MessageLevel36_Level36One4String1{Level36One4String1: "L36-N02-i121"},
						Level36OneOfIgnoreParts1:          &pbinline.MessageLevel36_Level36One5String1{Level36One5String1: "L36-N02-i131"},
					}
				}
				level28N02 = &pbinline.MessageLevel28{
					Level28FString1:                   "L28-N02-s101",
					Level28FString2:                   "L28-N02-s111",
					Level28PString1:                   utils.PointerString("L28-N02-p101"),
					Level28PString2:                   utils.PointerString("L28-N02-p111"),
					Level28RString1:                   []string{"L28-N02-r101", "L28-N02-r102", "L28-N02-r103"},
					Level28RString2:                   []string{"L28-N02-r111", "L28-N02-r112", "L28-N02-r113"},
					Level28MString1:                   map[string]string{"L28-N02-k101": "L28-N02-v101", "L28-N02-k102": "L28-N02-v102", "L28-N02-k103": "L28-N02-v103"},
					Level28MString2:                   map[string]string{"L28-N02-k111": "L28-N02-v111", "L28-N02-k112": "L28-N02-v112", "L28-N02-k113": "L28-N02-v113"},
					Level28FMessageExtern1:            &pbexternal.Message1{FString1: "L28-N02-e101", FString2: "L28-N02-e102", FString3: "L28-N02-e103"},
					Level28FMessageExtern2:            &pbexternal.Message1{FString1: "L28-N02-e111", FString2: "L28-N02-e112", FString3: "L28-N02-e113"},
					Level28FMessageInline34:           level34N03,
					Level28FMessageInline35:           level35N02,
					Level28OneOfExtern1:               &pbinline.MessageLevel28_Level28One1MessageInline34{Level28One1MessageInline34: level34N04},
					Level28OneOfInline1:               &pbinline.MessageLevel28_Level28One2MessageInline36{Level28One2MessageInline36: level36N02},
					Level28OneOfOmitempty1:            &pbinline.MessageLevel28_Level28One3String1{Level28One3String1: "L28-N02-o101"},
					Level28FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level28FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L28-N02-i101"},
					Level28FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L28-N02-i111", IgnoreFieldsString2: "L28-N02-i112"},
					Level28OneOfIgnoreSelf1:           &pbinline.MessageLevel28_Level28One4String1{Level28One4String1: "L28-N02-i121"},
					Level28OneOfIgnoreParts1:          &pbinline.MessageLevel28_Level28One5String1{Level28One5String1: "L28-N02-i131"},
				}
			}
			level16N02 = &pbinline.MessageLevel16{
				Level16FString1:                   "L16-N02-s101",
				Level16FString2:                   "L16-N02-s111",
				Level16PString1:                   utils.PointerString("L16-N02-p101"),
				Level16PString2:                   utils.PointerString("L16-N02-p111"),
				Level16RString1:                   []string{"L16-N02-r101", "L16-N02-r102", "L16-N02-r103"},
				Level16RString2:                   []string{"L16-N02-r111", "L16-N02-r112", "L16-N02-r113"},
				Level16MString1:                   map[string]string{"L16-N02-k101": "L16-N02-v101", "L16-N02-k102": "L16-N02-v102", "L16-N02-k103": "L16-N02-v103"},
				Level16MString2:                   map[string]string{"L16-N02-k111": "L16-N02-v111", "L16-N02-k112": "L16-N02-v112", "L16-N02-k113": "L16-N02-v113"},
				Level16FMessageExtern1:            &pbexternal.Message1{FString1: "L16-N02-e101", FString2: "L16-N02-e102", FString3: "L16-N02-e103"},
				Level16FMessageExtern2:            &pbexternal.Message1{FString1: "L16-N02-e111", FString2: "L16-N02-e112", FString3: "L16-N02-e113"},
				Level16FMessageInline26:           level26N03,
				Level16FMessageInline27:           level27N02,
				Level16OneOfExtern1:               &pbinline.MessageLevel16_Level16One1MessageInline26{Level16One1MessageInline26: level26N04},
				Level16OneOfInline1:               &pbinline.MessageLevel16_Level16One2MessageInline28{Level16One2MessageInline28: level28N02},
				Level16OneOfOmitempty1:            &pbinline.MessageLevel16_Level16One3String1{Level16One3String1: "L16-N02-o101"},
				Level16FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level16FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L16-N02-i101"},
				Level16FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L16-N02-i111", IgnoreFieldsString2: "L16-N02-i112"},
				Level16OneOfIgnoreSelf1:           &pbinline.MessageLevel16_Level16One4String1{Level16One4String1: "L16-N02-i121"},
				Level16OneOfIgnoreParts1:          &pbinline.MessageLevel16_Level16One5String1{Level16One5String1: "L16-N02-i131"},
			}
		}
		level06N02 = &pbinline.MessageLevel06{
			Level06FString1:                   "L06-N02-s101",
			Level06FString2:                   "L06-N02-s111",
			Level06PString1:                   utils.PointerString("L06-N02-p101"),
			Level06PString2:                   utils.PointerString("L06-N02-p111"),
			Level06RString1:                   []string{"L06-N02-r101", "L06-N02-r102", "L06-N02-r103"},
			Level06RString2:                   []string{"L06-N02-r111", "L06-N02-r112", "L06-N02-r113"},
			Level06MString1:                   map[string]string{"L06-N02-k101": "L06-N02-v101", "L06-N02-k102": "L06-N02-v102", "L06-N02-k103": "L06-N02-v103"},
			Level06MString2:                   map[string]string{"L06-N02-k111": "L06-N02-v111", "L06-N02-k112": "L06-N02-v112", "L06-N02-k113": "L06-N02-v113"},
			Level06FMessageExtern1:            &pbexternal.Message1{FString1: "L06-N02-e101", FString2: "L06-N02-e102", FString3: "L06-N02-e103"},
			Level06FMessageExtern2:            &pbexternal.Message1{FString1: "L06-N02-e111", FString2: "L06-N02-e112", FString3: "L06-N02-e113"},
			Level06FMessageInline14:           level14N03,
			Level06FMessageInline15:           level15N02,
			Level06OneOfExtern1:               &pbinline.MessageLevel06_Level06One1MessageInline14{Level06One1MessageInline14: level14N04},
			Level06OneOfInline1:               &pbinline.MessageLevel06_Level4One2MessageInline16{Level4One2MessageInline16: level16N02},
			Level06OneOfOmitempty1:            &pbinline.MessageLevel06_Level06One3String1{Level06One3String1: "L06-N02-o101"},
			Level06FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level06FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L06-N02-i101"},
			Level06FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L06-N02-i111", IgnoreFieldsString2: "L06-N02-i112"},
			Level06OneOfIgnoreSelf1:           &pbinline.MessageLevel06_Level06One4String1{Level06One4String1: "L06-N02-i121"},
			Level06OneOfIgnoreParts1:          &pbinline.MessageLevel06_Level06One5String1{Level06One5String1: "L06-N02-i131"},
		}
		{ // message in level08N01
			level18N01 = &pbinline.MessageLevel18{
				Level18FString1:                   "L18-N01-s101",
				Level18FString2:                   "L18-N01-s111",
				Level18PString1:                   utils.PointerString("L18-N01-p101"),
				Level18PString2:                   utils.PointerString("L18-N01-p111"),
				Level18RString1:                   []string{"L18-N01-r101", "L18-N01-r102", "L18-N01-r103"},
				Level18RString2:                   []string{"L18-N01-r111", "L18-N01-r112", "L18-N01-r113"},
				Level18MString1:                   map[string]string{"L18-N01-k101": "L18-N01-v101", "L18-N01-k102": "L18-N01-v102", "L18-N01-k103": "L18-N01-v103"},
				Level18MString2:                   map[string]string{"L18-N01-k111": "L18-N01-v111", "L18-N01-k112": "L18-N01-v112", "L18-N01-k113": "L18-N01-v113"},
				Level18FMessageExtern1:            &pbexternal.Message1{FString1: "L18-N01-e101", FString2: "L18-N01-e102", FString3: "L18-N01-e103"},
				Level18FMessageExtern2:            &pbexternal.Message1{FString1: "L18-N01-e111", FString2: "L18-N01-e112", FString3: "L18-N01-e113"},
				Level18OneOfExtern1:               &pbinline.MessageLevel18_Level18One1String1{Level18One1String1: "L18-N01-es101"},
				Level18OneOfInline1:               &pbinline.MessageLevel18_Level18One2String1{Level18One2String1: "L18-N01-is101"},
				Level18OneOfOmitempty1:            &pbinline.MessageLevel18_Level18One3String1{Level18One3String1: "L18-N01-o101"},
				Level18FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level18FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L18-N01-i101"},
				Level18FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L18-N01-i111", IgnoreFieldsString2: "L18-N01-i112"},
				Level18OneOfIgnoreSelf1:           &pbinline.MessageLevel18_Level18One4String1{Level18One4String1: "L18-N01-i121"},
				Level18OneOfIgnoreParts1:          &pbinline.MessageLevel18_Level18One5String1{Level18One5String1: "L18-N01-i131"},
			}
			level19N01 = &pbinline.MessageLevel19{
				Level19FString1:                   "L19-N01-s101",
				Level19FString2:                   "L19-N01-s111",
				Level19PString1:                   utils.PointerString("L19-N01-p101"),
				Level19PString2:                   utils.PointerString("L19-N01-p111"),
				Level19RString1:                   []string{"L19-N01-r101", "L19-N01-r102", "L19-N01-r103"},
				Level19RString2:                   []string{"L19-N01-r111", "L19-N01-r112", "L19-N01-r113"},
				Level19MString1:                   map[string]string{"L19-N01-k101": "L19-N01-v101", "L19-N01-k102": "L19-N01-v102", "L19-N01-k103": "L19-N01-v103"},
				Level19MString2:                   map[string]string{"L19-N01-k111": "L19-N01-v111", "L19-N01-k112": "L19-N01-v112", "L19-N01-k113": "L19-N01-v113"},
				Level19FMessageExtern1:            &pbexternal.Message1{FString1: "L19-N01-e101", FString2: "L19-N01-e102", FString3: "L19-N01-e103"},
				Level19FMessageExtern2:            &pbexternal.Message1{FString1: "L19-N01-e111", FString2: "L19-N01-e112", FString3: "L19-N01-e113"},
				Level19OneOfExtern1:               &pbinline.MessageLevel19_Level19One1String1{Level19One1String1: "L19-N01-es101"},
				Level19OneOfInline1:               &pbinline.MessageLevel19_Level19One2String1{Level19One2String1: "L19-N01-is101"},
				Level19OneOfOmitempty1:            &pbinline.MessageLevel19_Level19One3String1{Level19One3String1: "L19-N01-o101"},
				Level19FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level19FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L19-N01-i101"},
				Level19FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L19-N01-i111", IgnoreFieldsString2: "L19-N01-i112"},
				Level19OneOfIgnoreSelf1:           &pbinline.MessageLevel19_Level19One4String1{Level19One4String1: "L19-N01-i121"},
				Level19OneOfIgnoreParts1:          &pbinline.MessageLevel19_Level19One5String1{Level19One5String1: "L19-N01-i131"},
			}
			level18N02 = &pbinline.MessageLevel18{
				Level18FString1:                   "L18-N02-s101",
				Level18FString2:                   "L18-N02-s111",
				Level18PString1:                   utils.PointerString("L18-N02-p101"),
				Level18PString2:                   utils.PointerString("L18-N02-p111"),
				Level18RString1:                   []string{"L18-N02-r101", "L18-N02-r102", "L18-N02-r103"},
				Level18RString2:                   []string{"L18-N02-r111", "L18-N02-r112", "L18-N02-r113"},
				Level18MString1:                   map[string]string{"L18-N02-k101": "L18-N02-v101", "L18-N02-k102": "L18-N02-v102", "L18-N02-k103": "L18-N02-v103"},
				Level18MString2:                   map[string]string{"L18-N02-k111": "L18-N02-v111", "L18-N02-k112": "L18-N02-v112", "L18-N02-k113": "L18-N02-v113"},
				Level18FMessageExtern1:            &pbexternal.Message1{FString1: "L18-N02-e101", FString2: "L18-N02-e102", FString3: "L18-N02-e103"},
				Level18FMessageExtern2:            &pbexternal.Message1{FString1: "L18-N02-e111", FString2: "L18-N02-e112", FString3: "L18-N02-e113"},
				Level18OneOfExtern1:               &pbinline.MessageLevel18_Level18One1String1{Level18One1String1: "L18-N02-es101"},
				Level18OneOfInline1:               &pbinline.MessageLevel18_Level18One2String1{Level18One2String1: "L18-N02-is101"},
				Level18OneOfOmitempty1:            &pbinline.MessageLevel18_Level18One3String1{Level18One3String1: "L18-N02-o101"},
				Level18FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level18FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L18-N02-i101"},
				Level18FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L18-N02-i111", IgnoreFieldsString2: "L18-N02-i112"},
				Level18OneOfIgnoreSelf1:           &pbinline.MessageLevel18_Level18One4String1{Level18One4String1: "L18-N02-i121"},
				Level18OneOfIgnoreParts1:          &pbinline.MessageLevel18_Level18One5String1{Level18One5String1: "L18-N02-i131"},
			}
			level20N01 = &pbinline.MessageLevel20{
				Level20FString1:                   "L20-N01-s101",
				Level20FString2:                   "L20-N01-s111",
				Level20PString1:                   utils.PointerString("L20-N01-p101"),
				Level20PString2:                   utils.PointerString("L20-N01-p111"),
				Level20RString1:                   []string{"L20-N01-r101", "L20-N01-r102", "L20-N01-r103"},
				Level20RString2:                   []string{"L20-N01-r111", "L20-N01-r112", "L20-N01-r113"},
				Level20MString1:                   map[string]string{"L20-N01-k101": "L20-N01-v101", "L20-N01-k102": "L20-N01-v102", "L20-N01-k103": "L20-N01-v103"},
				Level20MString2:                   map[string]string{"L20-N01-k111": "L20-N01-v111", "L20-N01-k112": "L20-N01-v112", "L20-N01-k113": "L20-N01-v113"},
				Level20FMessageExtern1:            &pbexternal.Message1{FString1: "L20-N01-e101", FString2: "L20-N01-e102", FString3: "L20-N01-e103"},
				Level20FMessageExtern2:            &pbexternal.Message1{FString1: "L20-N01-e111", FString2: "L20-N01-e112", FString3: "L20-N01-e113"},
				Level20OneOfExtern1:               &pbinline.MessageLevel20_Level20One1String1{Level20One1String1: "L20-N01-es101"},
				Level20OneOfInline1:               &pbinline.MessageLevel20_Level20One2String1{Level20One2String1: "L20-N01-is101"},
				Level20OneOfOmitempty1:            &pbinline.MessageLevel20_Level20One3String1{Level20One3String1: "L20-N01-o101"},
				Level20FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level20FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L20-N01-i101"},
				Level20FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L20-N01-i111", IgnoreFieldsString2: "L20-N01-i112"},
				Level20OneOfIgnoreSelf1:           &pbinline.MessageLevel20_Level20One4String1{Level20One4String1: "L20-N01-i121"},
				Level20OneOfIgnoreParts1:          &pbinline.MessageLevel20_Level20One5String1{Level20One5String1: "L20-N01-i131"},
			}
		}
		level08N01 = &pbinline.MessageLevel08{
			Level08FString1:                   "L08-N01-s101",
			Level08FString2:                   "L08-N01-s111",
			Level08PString1:                   utils.PointerString("L08-N01-p101"),
			Level08PString2:                   utils.PointerString("L08-N01-p111"),
			Level08RString1:                   []string{"L08-N01-r101", "L08-N01-r102", "L08-N01-r103"},
			Level08RString2:                   []string{"L08-N01-r111", "L08-N01-r112", "L08-N01-r113"},
			Level08MString1:                   map[string]string{"L08-N01-k101": "L08-N01-v101", "L08-N01-k102": "L08-N01-v102", "L08-N01-k103": "L08-N01-v103"},
			Level08MString2:                   map[string]string{"L08-N01-k111": "L08-N01-v111", "L08-N01-k112": "L08-N01-v112", "L08-N01-k113": "L08-N01-v113"},
			Level08FMessageExtern1:            &pbexternal.Message1{FString1: "L08-N01-e101", FString2: "L08-N01-e102", FString3: "L08-N01-e103"},
			Level08FMessageExtern2:            &pbexternal.Message1{FString1: "L08-N01-e111", FString2: "L08-N01-e112", FString3: "L08-N01-e113"},
			Level08FMessageInline18:           level18N01,
			Level08FMessageInline19:           level19N01,
			Level08OneOfExtern1:               &pbinline.MessageLevel08_Level08One1MessageInline18{Level08One1MessageInline18: level18N02},
			Level08OneOfInline1:               &pbinline.MessageLevel08_Level08One2MessageInline20{Level08One2MessageInline20: level20N01},
			Level08OneOfOmitempty1:            &pbinline.MessageLevel08_Level08One3String1{Level08One3String1: "L08-N01-o101"},
			Level08FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level08FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L08-N01-i101"},
			Level08FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L08-N01-i111", IgnoreFieldsString2: "L08-N01-i112"},
			Level08OneOfIgnoreSelf1:           &pbinline.MessageLevel08_Level08One4String1{Level08One4String1: "L08-N01-i121"},
			Level08OneOfIgnoreParts1:          &pbinline.MessageLevel08_Level08One5String1{Level08One5String1: "L08-N01-i131"},
		}
	}
	level02N01 = &pbinline.MessageLevel02{
		Level02FString1:                   "L02-N01-s101",
		Level02FString2:                   "L02-N01-s111",
		Level02PString1:                   utils.PointerString("L02-N01-p101"),
		Level02PString2:                   utils.PointerString("L02-N01-p111"),
		Level02RString1:                   []string{"L02-N01-r101", "L02-N01-r102", "L02-N01-r103"},
		Level02RString2:                   []string{"L02-N01-r111", "L02-N01-r112", "L02-N01-r113"},
		Level02MString1:                   map[string]string{"L02-N01-k101": "L02-N01-v101", "L02-N01-k102": "L02-N01-v102", "L02-N01-k103": "L02-N01-v103"},
		Level02MString2:                   map[string]string{"L02-N01-k111": "L02-N01-v111", "L02-N01-k112": "L02-N01-v112", "L02-N01-k113": "L02-N01-v113"},
		Level02FMessageExtern1:            &pbexternal.Message1{FString1: "L02-N01-e101", FString2: "L02-N01-e102", FString3: "L02-N01-e103"},
		Level02FMessageExtern2:            &pbexternal.Message1{FString1: "L02-N01-e111", FString2: "L02-N01-e112", FString3: "L02-N01-e113"},
		Level02FMessageInline06:           level06N01,
		Level02FMessageInline07:           level07N01,
		Level02OneOfExtern1:               &pbinline.MessageLevel02_Level02One1MessageInline06{Level02One1MessageInline06: level06N02},
		Level02OneOfInline1:               &pbinline.MessageLevel02_Level02One2MessageInline08{Level02One2MessageInline08: level08N01},
		Level02OneOfOmitempty1:            &pbinline.MessageLevel02_Level02One3String1{Level02One3String1: "L02-N01-o101"},
		Level02FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
		Level02FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L02-N01-i101"},
		Level02FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L02-N01-i111", IgnoreFieldsString2: "L02-N01-i112"},
		Level02OneOfIgnoreSelf1:           &pbinline.MessageLevel02_Level02One4String1{Level02One4String1: "L02-N01-i121"},
		Level02OneOfIgnoreParts1:          &pbinline.MessageLevel02_Level02One5String1{Level02One5String1: "L02-N01-i131"},
	}

	level03N01 = &pbinline.MessageLevel03{
		Level03FString1:                   "L03-N01-s101",
		Level03FString2:                   "L03-N01-s111",
		Level03PString1:                   utils.PointerString("L03-N01-p101"),
		Level03PString2:                   utils.PointerString("L03-N01-p111"),
		Level03RString1:                   []string{"L03-N01-r101", "L03-N01-r102", "L03-N01-r103"},
		Level03RString2:                   []string{"L03-N01-r111", "L03-N01-r112", "L03-N01-r113"},
		Level03MString1:                   map[string]string{"L03-N01-k101": "L03-N01-v101", "L03-N01-k102": "L03-N01-v102", "L03-N01-k103": "L03-N01-v103"},
		Level03MString2:                   map[string]string{"L03-N01-k111": "L03-N01-v111", "L03-N01-k112": "L03-N01-v112", "L03-N01-k113": "L03-N01-v113"},
		Level03FMessageExtern1:            &pbexternal.Message1{FString1: "L03-N01-e101", FString2: "L03-N01-e102", FString3: "L03-N01-e103"},
		Level03FMessageExtern2:            &pbexternal.Message1{FString1: "L03-N01-e111", FString2: "L03-N01-e112", FString3: "L03-N01-e113"},
		Level03OneOfExtern1:               &pbinline.MessageLevel03_Level03One1String1{Level03One1String1: "L03-N01-es101"},
		Level03OneOfInline1:               &pbinline.MessageLevel03_Level03One2String1{Level03One2String1: "L03-N01-is101"},
		Level03OneOfOmitempty1:            &pbinline.MessageLevel03_Level03One3String1{Level03One3String1: "L03-N01-o101"},
		Level03FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
		Level03FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L03-N01-i101"},
		Level03FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L03-N01-i111", IgnoreFieldsString2: "L03-N01-i112"},
		Level03OneOfIgnoreSelf1:           &pbinline.MessageLevel03_Level03One4String1{Level03One4String1: "L03-N01-i121"},
		Level03OneOfIgnoreParts1:          &pbinline.MessageLevel03_Level03One5String1{Level03One5String1: "L03-N01-i131"},
	}
	// -------------------------------
	{ // message in level02N02
		{ // message in level06N03
			{ // message for level14N05
				{ // message for level22N09
					level38N17 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N17-s101",
						Level38FString2:                   "L38-N17-s111",
						Level38PString1:                   utils.PointerString("L38-N17-p101"),
						Level38PString2:                   utils.PointerString("L38-N17-p111"),
						Level38RString1:                   []string{"L38-N17-r101", "L38-N17-r102", "L38-N17-r103"},
						Level38RString2:                   []string{"L38-N17-r111", "L38-N17-r112", "L38-N17-r113"},
						Level38MString1:                   map[string]string{"L38-N17-k101": "L38-N17-v101", "L38-N17-k102": "L38-N17-v102", "L38-N17-k103": "L38-N17-v103"},
						Level38MString2:                   map[string]string{"L38-N17-k111": "L38-N17-v111", "L38-N17-k112": "L38-N17-v112", "L38-N17-k113": "L38-N17-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N17-e101", FString2: "L38-N17-e102", FString3: "L38-N17-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N17-e111", FString2: "L38-N17-e112", FString3: "L38-N17-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N17-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N17-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N17-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N17-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N17-i111", IgnoreFieldsString2: "L38-N17-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N17-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N17-i131"},
					}
					level39N09 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N09-s101",
						Level39FString2:                   "L39-N09-s111",
						Level39PString1:                   utils.PointerString("L39-N09-p101"),
						Level39PString2:                   utils.PointerString("L39-N09-p111"),
						Level39RString1:                   []string{"L39-N09-r101", "L39-N09-r102", "L39-N09-r103"},
						Level39RString2:                   []string{"L39-N09-r111", "L39-N09-r112", "L39-N09-r113"},
						Level39MString1:                   map[string]string{"L39-N09-k101": "L39-N09-v101", "L39-N09-k102": "L39-N09-v102", "L39-N09-k103": "L39-N09-v103"},
						Level39MString2:                   map[string]string{"L39-N09-k111": "L39-N09-v111", "L39-N09-k112": "L39-N09-v112", "L39-N09-k113": "L39-N09-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N09-e101", FString2: "L39-N09-e102", FString3: "L39-N09-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N09-e111", FString2: "L39-N09-e112", FString3: "L39-N09-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N09-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N09-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N09-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N09-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N09-i111", IgnoreFieldsString2: "L39-N09-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N09-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N09-i131"},
					}
					level38N18 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N18-s101",
						Level38FString2:                   "L38-N18-s111",
						Level38PString1:                   utils.PointerString("L38-N18-p101"),
						Level38PString2:                   utils.PointerString("L38-N18-p111"),
						Level38RString1:                   []string{"L38-N18-r101", "L38-N18-r102", "L38-N18-r103"},
						Level38RString2:                   []string{"L38-N18-r111", "L38-N18-r112", "L38-N18-r113"},
						Level38MString1:                   map[string]string{"L38-N18-k101": "L38-N18-v101", "L38-N18-k102": "L38-N18-v102", "L38-N18-k103": "L38-N18-v103"},
						Level38MString2:                   map[string]string{"L38-N18-k111": "L38-N18-v111", "L38-N18-k112": "L38-N18-v112", "L38-N18-k113": "L38-N18-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N18-e101", FString2: "L38-N18-e102", FString3: "L38-N18-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N18-e111", FString2: "L38-N18-e112", FString3: "L38-N18-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N18-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N18-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N18-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N18-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N18-i111", IgnoreFieldsString2: "L38-N18-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N18-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N18-i131"},
					}
					level40N09 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N09-s101",
						Level40FString2:                   "L40-N09-s111",
						Level40PString1:                   utils.PointerString("L40-N09-p101"),
						Level40PString2:                   utils.PointerString("L40-N09-p111"),
						Level40RString1:                   []string{"L40-N09-r101", "L40-N09-r102", "L40-N09-r103"},
						Level40RString2:                   []string{"L40-N09-r111", "L40-N09-r112", "L40-N09-r113"},
						Level40MString1:                   map[string]string{"L40-N09-k101": "L40-N09-v101", "L40-N09-k102": "L40-N09-v102", "L40-N09-k103": "L40-N09-v103"},
						Level40MString2:                   map[string]string{"L40-N09-k111": "L40-N09-v111", "L40-N09-k112": "L40-N09-v112", "L40-N09-k113": "L40-N09-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N09-e101", FString2: "L40-N09-e102", FString3: "L40-N09-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N09-e111", FString2: "L40-N09-e112", FString3: "L40-N09-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N09-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N09-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N09-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N09-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N09-i111", IgnoreFieldsString2: "L40-N09-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N09-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N09-i131"},
					}
				}
				level22N09 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N09-s101",
					Level22FString2:                   "L22-N09-s111",
					Level22PString1:                   utils.PointerString("L22-N09-p101"),
					Level22PString2:                   utils.PointerString("L22-N09-p111"),
					Level22RString1:                   []string{"L22-N09-r101", "L22-N09-r102", "L22-N09-r103"},
					Level22RString2:                   []string{"L22-N09-r111", "L22-N09-r112", "L22-N09-r113"},
					Level22MString1:                   map[string]string{"L22-N09-k101": "L22-N09-v101", "L22-N09-k102": "L22-N09-v102", "L22-N09-k103": "L22-N09-v103"},
					Level22MString2:                   map[string]string{"L22-N09-k111": "L22-N09-v111", "L22-N09-k112": "L22-N09-v112", "L22-N09-k113": "L22-N09-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N09-e101", FString2: "L22-N09-e102", FString3: "L22-N09-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N09-e111", FString2: "L22-N09-e112", FString3: "L22-N09-e113"},
					Level22FMessageInline38:           level38N17,
					Level22FMessageInline39:           level39N09,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N18},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N09},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N09-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N09-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N09-i111", IgnoreFieldsString2: "L22-N09-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N09-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N09-i131"},
				}
				level23N05 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N05-s101",
					Level23FString2:                   "L23-N05-s111",
					Level23PString1:                   utils.PointerString("L23-N05-p101"),
					Level23PString2:                   utils.PointerString("L23-N05-p111"),
					Level23RString1:                   []string{"L23-N05-r101", "L23-N05-r102", "L23-N05-r103"},
					Level23RString2:                   []string{"L23-N05-r111", "L23-N05-r112", "L23-N05-r113"},
					Level23MString1:                   map[string]string{"L23-N05-k101": "L23-N05-v101", "L23-N05-k102": "L23-N05-v102", "L23-N05-k103": "L23-N05-v103"},
					Level23MString2:                   map[string]string{"L23-N05-k111": "L23-N05-v111", "L23-N05-k112": "L23-N05-v112", "L23-N05-k113": "L23-N05-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N05-e101", FString2: "L23-N05-e102", FString3: "L23-N05-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N05-e111", FString2: "L23-N05-e112", FString3: "L23-N05-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N05-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N05-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N05-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N05-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N05-i111", IgnoreFieldsString2: "L23-N05-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N05-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N05-i131"},
				}
				{ // message for level22N10
					level38N19 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N19-s101",
						Level38FString2:                   "L38-N19-s111",
						Level38PString1:                   utils.PointerString("L38-N19-p101"),
						Level38PString2:                   utils.PointerString("L38-N19-p111"),
						Level38RString1:                   []string{"L38-N19-r101", "L38-N19-r102", "L38-N19-r103"},
						Level38RString2:                   []string{"L38-N19-r111", "L38-N19-r112", "L38-N19-r113"},
						Level38MString1:                   map[string]string{"L38-N19-k101": "L38-N19-v101", "L38-N19-k102": "L38-N19-v102", "L38-N19-k103": "L38-N19-v103"},
						Level38MString2:                   map[string]string{"L38-N19-k111": "L38-N19-v111", "L38-N19-k112": "L38-N19-v112", "L38-N19-k113": "L38-N19-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N19-e101", FString2: "L38-N19-e102", FString3: "L38-N19-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N19-e111", FString2: "L38-N19-e112", FString3: "L38-N19-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N19-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N19-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N19-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N19-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N19-i111", IgnoreFieldsString2: "L38-N19-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N19-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N19-i131"},
					}
					level39N10 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N10-s101",
						Level39FString2:                   "L39-N10-s111",
						Level39PString1:                   utils.PointerString("L39-N10-p101"),
						Level39PString2:                   utils.PointerString("L39-N10-p111"),
						Level39RString1:                   []string{"L39-N10-r101", "L39-N10-r102", "L39-N10-r103"},
						Level39RString2:                   []string{"L39-N10-r111", "L39-N10-r112", "L39-N10-r113"},
						Level39MString1:                   map[string]string{"L39-N10-k101": "L39-N10-v101", "L39-N10-k102": "L39-N10-v102", "L39-N10-k103": "L39-N10-v103"},
						Level39MString2:                   map[string]string{"L39-N10-k111": "L39-N10-v111", "L39-N10-k112": "L39-N10-v112", "L39-N10-k113": "L39-N10-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N10-e101", FString2: "L39-N10-e102", FString3: "L39-N10-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N10-e111", FString2: "L39-N10-e112", FString3: "L39-N10-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N10-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N10-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N10-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N10-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N10-i111", IgnoreFieldsString2: "L39-N10-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N10-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N10-i131"},
					}
					level38N20 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N20-s101",
						Level38FString2:                   "L38-N20-s111",
						Level38PString1:                   utils.PointerString("L38-N20-p101"),
						Level38PString2:                   utils.PointerString("L38-N20-p111"),
						Level38RString1:                   []string{"L38-N20-r101", "L38-N20-r102", "L38-N20-r103"},
						Level38RString2:                   []string{"L38-N20-r111", "L38-N20-r112", "L38-N20-r113"},
						Level38MString1:                   map[string]string{"L38-N20-k101": "L38-N20-v101", "L38-N20-k102": "L38-N20-v102", "L38-N20-k103": "L38-N20-v103"},
						Level38MString2:                   map[string]string{"L38-N20-k111": "L38-N20-v111", "L38-N20-k112": "L38-N20-v112", "L38-N20-k113": "L38-N20-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N20-e101", FString2: "L38-N20-e102", FString3: "L38-N20-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N20-e111", FString2: "L38-N20-e112", FString3: "L38-N20-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N20-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N20-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N20-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N20-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N20-i111", IgnoreFieldsString2: "L38-N20-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N20-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N20-i131"},
					}
					level40N10 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N10-s101",
						Level40FString2:                   "L40-N10-s111",
						Level40PString1:                   utils.PointerString("L40-N10-p101"),
						Level40PString2:                   utils.PointerString("L40-N10-p111"),
						Level40RString1:                   []string{"L40-N10-r101", "L40-N10-r102", "L40-N10-r103"},
						Level40RString2:                   []string{"L40-N10-r111", "L40-N10-r112", "L40-N10-r113"},
						Level40MString1:                   map[string]string{"L40-N10-k101": "L40-N10-v101", "L40-N10-k102": "L40-N10-v102", "L40-N10-k103": "L40-N10-v103"},
						Level40MString2:                   map[string]string{"L40-N10-k111": "L40-N10-v111", "L40-N10-k112": "L40-N10-v112", "L40-N10-k113": "L40-N10-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N10-e101", FString2: "L40-N10-e102", FString3: "L40-N10-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N10-e111", FString2: "L40-N10-e112", FString3: "L40-N10-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N10-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N10-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N10-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N10-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N10-i111", IgnoreFieldsString2: "L40-N10-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N10-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N10-i131"},
					}
				}
				level22N10 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N10-s101",
					Level22FString2:                   "L22-N10-s111",
					Level22PString1:                   utils.PointerString("L22-N10-p101"),
					Level22PString2:                   utils.PointerString("L22-N10-p111"),
					Level22RString1:                   []string{"L22-N10-r101", "L22-N10-r102", "L22-N10-r103"},
					Level22RString2:                   []string{"L22-N10-r111", "L22-N10-r112", "L22-N10-r113"},
					Level22MString1:                   map[string]string{"L22-N10-k101": "L22-N10-v101", "L22-N10-k102": "L22-N10-v102", "L22-N10-k103": "L22-N10-v103"},
					Level22MString2:                   map[string]string{"L22-N10-k111": "L22-N10-v111", "L22-N10-k112": "L22-N10-v112", "L22-N10-k113": "L22-N10-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N10-e101", FString2: "L22-N10-e102", FString3: "L22-N10-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N10-e111", FString2: "L22-N10-e112", FString3: "L22-N10-e113"},
					Level22FMessageInline38:           level38N19,
					Level22FMessageInline39:           level39N10,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N20},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N10},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N10-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N10-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N10-i111", IgnoreFieldsString2: "L22-N10-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N10-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N10-i131"},
				}
				level24N05 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N05-s101",
					Level24FString2:                   "L24-N05-s111",
					Level24PString1:                   utils.PointerString("L24-N05-p101"),
					Level24PString2:                   utils.PointerString("L24-N05-p111"),
					Level24RString1:                   []string{"L24-N05-r101", "L24-N05-r102", "L24-N05-r103"},
					Level24RString2:                   []string{"L24-N05-r111", "L24-N05-r112", "L24-N05-r113"},
					Level24MString1:                   map[string]string{"L24-N05-k101": "L24-N05-v101", "L24-N05-k102": "L24-N05-v102", "L24-N05-k103": "L24-N05-v103"},
					Level24MString2:                   map[string]string{"L24-N05-k111": "L24-N05-v111", "L24-N05-k112": "L24-N05-v112", "L24-N05-k113": "L24-N05-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N05-e101", FString2: "L24-N05-e102", FString3: "L24-N05-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N05-e111", FString2: "L24-N05-e112", FString3: "L24-N05-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N05-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N05-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N05-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N05-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N05-i111", IgnoreFieldsString2: "L24-N05-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N05-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N05-i131"},
				}
			}
			level14N05 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N05-s101",
				Level14FString2:                   "L14-N05-s111",
				Level14PString1:                   utils.PointerString("L14-N05-p101"),
				Level14PString2:                   utils.PointerString("L14-N05-p111"),
				Level14RString1:                   []string{"L14-N05-r101", "L14-N05-r102", "L14-N05-r103"},
				Level14RString2:                   []string{"L14-N05-r111", "L14-N05-r112", "L14-N05-r113"},
				Level14MString1:                   map[string]string{"L14-N05-k101": "L14-N05-v101", "L14-N05-k102": "L14-N05-v102", "L14-N05-k103": "L14-N05-v103"},
				Level14MString2:                   map[string]string{"L14-N05-k111": "L14-N05-v111", "L14-N05-k112": "L14-N05-v112", "L14-N05-k113": "L14-N05-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N05-e101", FString2: "L14-N05-e102", FString3: "L14-N05-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N05-e111", FString2: "L14-N05-e112", FString3: "L14-N05-e113"},
				Level14FMessageInline22:           level22N09,
				Level14FMessageInline23:           level23N05,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N10},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N05},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N05-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N05-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N05-i111", IgnoreFieldsString2: "L14-N05-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N05-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N05-i131"},
			}
			level15N03 = &pbinline.MessageLevel15{
				Level15FString1:                   "L15-N03-s101",
				Level15FString2:                   "L15-N03-s111",
				Level15PString1:                   utils.PointerString("L15-N03-p101"),
				Level15PString2:                   utils.PointerString("L15-N03-p111"),
				Level15RString1:                   []string{"L15-N03-r101", "L15-N03-r102", "L15-N03-r103"},
				Level15RString2:                   []string{"L15-N03-r111", "L15-N03-r112", "L15-N03-r113"},
				Level15MString1:                   map[string]string{"L15-N03-k101": "L15-N03-v101", "L15-N03-k102": "L15-N03-v102", "L15-N03-k103": "L15-N03-v103"},
				Level15MString2:                   map[string]string{"L15-N03-k111": "L15-N03-v111", "L15-N03-k112": "L15-N03-v112", "L15-N03-k113": "L15-N03-v113"},
				Level15FMessageExtern1:            &pbexternal.Message1{FString1: "L15-N03-e101", FString2: "L15-N03-e102", FString3: "L15-N03-e103"},
				Level15FMessageExtern2:            &pbexternal.Message1{FString1: "L15-N03-e111", FString2: "L15-N03-e112", FString3: "L15-N03-e113"},
				Level15OneOfExtern1:               &pbinline.MessageLevel15_Level15One1String1{Level15One1String1: "L15-N03-es101"},
				Level15OneOfInline1:               &pbinline.MessageLevel15_Level15One2String1{Level15One2String1: "L15-N03-is101"},
				Level15OneOfOmitempty1:            &pbinline.MessageLevel15_Level15One3String1{Level15One3String1: "L15-N03-o101"},
				Level15FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level15FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L15-N03-i101"},
				Level15FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L15-N03-i111", IgnoreFieldsString2: "L15-N03-i112"},
				Level15OneOfIgnoreSelf1:           &pbinline.MessageLevel15_Level15One4String1{Level15One4String1: "L15-N03-i121"},
				Level15OneOfIgnoreParts1:          &pbinline.MessageLevel15_Level15One5String1{Level15One5String1: "L15-N03-i131"},
			}
			{ // message for level14N06
				{ // message for level22N11
					level38N21 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N21-s101",
						Level38FString2:                   "L38-N21-s111",
						Level38PString1:                   utils.PointerString("L38-N21-p101"),
						Level38PString2:                   utils.PointerString("L38-N21-p111"),
						Level38RString1:                   []string{"L38-N21-r101", "L38-N21-r102", "L38-N21-r103"},
						Level38RString2:                   []string{"L38-N21-r111", "L38-N21-r112", "L38-N21-r113"},
						Level38MString1:                   map[string]string{"L38-N21-k101": "L38-N21-v101", "L38-N21-k102": "L38-N21-v102", "L38-N21-k103": "L38-N21-v103"},
						Level38MString2:                   map[string]string{"L38-N21-k111": "L38-N21-v111", "L38-N21-k112": "L38-N21-v112", "L38-N21-k113": "L38-N21-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N21-e101", FString2: "L38-N21-e102", FString3: "L38-N21-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N21-e111", FString2: "L38-N21-e112", FString3: "L38-N21-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N21-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N21-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N21-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N21-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N21-i111", IgnoreFieldsString2: "L38-N21-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N21-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N21-i131"},
					}
					level39N11 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N11-s101",
						Level39FString2:                   "L39-N11-s111",
						Level39PString1:                   utils.PointerString("L39-N11-p101"),
						Level39PString2:                   utils.PointerString("L39-N11-p111"),
						Level39RString1:                   []string{"L39-N11-r101", "L39-N11-r102", "L39-N11-r103"},
						Level39RString2:                   []string{"L39-N11-r111", "L39-N11-r112", "L39-N11-r113"},
						Level39MString1:                   map[string]string{"L39-N11-k101": "L39-N11-v101", "L39-N11-k102": "L39-N11-v102", "L39-N11-k103": "L39-N11-v103"},
						Level39MString2:                   map[string]string{"L39-N11-k111": "L39-N11-v111", "L39-N11-k112": "L39-N11-v112", "L39-N11-k113": "L39-N11-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N11-e101", FString2: "L39-N11-e102", FString3: "L39-N11-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N11-e111", FString2: "L39-N11-e112", FString3: "L39-N11-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N11-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N11-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N11-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N11-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N11-i111", IgnoreFieldsString2: "L39-N11-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N11-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N11-i131"},
					}
					level38N22 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N22-s101",
						Level38FString2:                   "L38-N22-s111",
						Level38PString1:                   utils.PointerString("L38-N22-p101"),
						Level38PString2:                   utils.PointerString("L38-N22-p111"),
						Level38RString1:                   []string{"L38-N22-r101", "L38-N22-r102", "L38-N22-r103"},
						Level38RString2:                   []string{"L38-N22-r111", "L38-N22-r112", "L38-N22-r113"},
						Level38MString1:                   map[string]string{"L38-N22-k101": "L38-N22-v101", "L38-N22-k102": "L38-N22-v102", "L38-N22-k103": "L38-N22-v103"},
						Level38MString2:                   map[string]string{"L38-N22-k111": "L38-N22-v111", "L38-N22-k112": "L38-N22-v112", "L38-N22-k113": "L38-N22-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N22-e101", FString2: "L38-N22-e102", FString3: "L38-N22-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N22-e111", FString2: "L38-N22-e112", FString3: "L38-N22-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N22-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N22-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N22-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N22-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N22-i111", IgnoreFieldsString2: "L38-N22-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N22-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N22-i131"},
					}
					level40N11 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N11-s101",
						Level40FString2:                   "L40-N11-s111",
						Level40PString1:                   utils.PointerString("L40-N11-p101"),
						Level40PString2:                   utils.PointerString("L40-N11-p111"),
						Level40RString1:                   []string{"L40-N11-r101", "L40-N11-r102", "L40-N11-r103"},
						Level40RString2:                   []string{"L40-N11-r111", "L40-N11-r112", "L40-N11-r113"},
						Level40MString1:                   map[string]string{"L40-N11-k101": "L40-N11-v101", "L40-N11-k102": "L40-N11-v102", "L40-N11-k103": "L40-N11-v103"},
						Level40MString2:                   map[string]string{"L40-N11-k111": "L40-N11-v111", "L40-N11-k112": "L40-N11-v112", "L40-N11-k113": "L40-N11-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N11-e101", FString2: "L40-N11-e102", FString3: "L40-N11-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N11-e111", FString2: "L40-N11-e112", FString3: "L40-N11-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N11-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N11-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N11-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N11-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N11-i111", IgnoreFieldsString2: "L40-N11-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N11-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N11-i131"},
					}
				}
				level22N11 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N11-s101",
					Level22FString2:                   "L22-N11-s111",
					Level22PString1:                   utils.PointerString("L22-N11-p101"),
					Level22PString2:                   utils.PointerString("L22-N11-p111"),
					Level22RString1:                   []string{"L22-N11-r101", "L22-N11-r102", "L22-N11-r103"},
					Level22RString2:                   []string{"L22-N11-r111", "L22-N11-r112", "L22-N11-r113"},
					Level22MString1:                   map[string]string{"L22-N11-k101": "L22-N11-v101", "L22-N11-k102": "L22-N11-v102", "L22-N11-k103": "L22-N11-v103"},
					Level22MString2:                   map[string]string{"L22-N11-k111": "L22-N11-v111", "L22-N11-k112": "L22-N11-v112", "L22-N11-k113": "L22-N11-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N11-e101", FString2: "L22-N11-e102", FString3: "L22-N11-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N11-e111", FString2: "L22-N11-e112", FString3: "L22-N11-e113"},
					Level22FMessageInline38:           level38N21,
					Level22FMessageInline39:           level39N11,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N22},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N11},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N11-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N11-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N11-i111", IgnoreFieldsString2: "L22-N11-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N11-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N11-i131"},
				}
				level23N06 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N06-s101",
					Level23FString2:                   "L23-N06-s111",
					Level23PString1:                   utils.PointerString("L23-N06-p101"),
					Level23PString2:                   utils.PointerString("L23-N06-p111"),
					Level23RString1:                   []string{"L23-N06-r101", "L23-N06-r102", "L23-N06-r103"},
					Level23RString2:                   []string{"L23-N06-r111", "L23-N06-r112", "L23-N06-r113"},
					Level23MString1:                   map[string]string{"L23-N06-k101": "L23-N06-v101", "L23-N06-k102": "L23-N06-v102", "L23-N06-k103": "L23-N06-v103"},
					Level23MString2:                   map[string]string{"L23-N06-k111": "L23-N06-v111", "L23-N06-k112": "L23-N06-v112", "L23-N06-k113": "L23-N06-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N06-e101", FString2: "L23-N06-e102", FString3: "L23-N06-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N06-e111", FString2: "L23-N06-e112", FString3: "L23-N06-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N06-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N06-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N06-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N06-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N06-i111", IgnoreFieldsString2: "L23-N06-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N06-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N06-i131"},
				}
				{ // message for level22N12
					level38N23 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N23-s101",
						Level38FString2:                   "L38-N23-s111",
						Level38PString1:                   utils.PointerString("L38-N23-p101"),
						Level38PString2:                   utils.PointerString("L38-N23-p111"),
						Level38RString1:                   []string{"L38-N23-r101", "L38-N23-r102", "L38-N23-r103"},
						Level38RString2:                   []string{"L38-N23-r111", "L38-N23-r112", "L38-N23-r113"},
						Level38MString1:                   map[string]string{"L38-N23-k101": "L38-N23-v101", "L38-N23-k102": "L38-N23-v102", "L38-N23-k103": "L38-N23-v103"},
						Level38MString2:                   map[string]string{"L38-N23-k111": "L38-N23-v111", "L38-N23-k112": "L38-N23-v112", "L38-N23-k113": "L38-N23-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N23-e101", FString2: "L38-N23-e102", FString3: "L38-N23-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N23-e111", FString2: "L38-N23-e112", FString3: "L38-N23-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N23-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N23-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N23-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N23-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N23-i111", IgnoreFieldsString2: "L38-N23-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N23-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N23-i131"},
					}
					level39N12 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N12-s101",
						Level39FString2:                   "L39-N12-s111",
						Level39PString1:                   utils.PointerString("L39-N12-p101"),
						Level39PString2:                   utils.PointerString("L39-N12-p111"),
						Level39RString1:                   []string{"L39-N12-r101", "L39-N12-r102", "L39-N12-r103"},
						Level39RString2:                   []string{"L39-N12-r111", "L39-N12-r112", "L39-N12-r113"},
						Level39MString1:                   map[string]string{"L39-N12-k101": "L39-N12-v101", "L39-N12-k102": "L39-N12-v102", "L39-N12-k103": "L39-N12-v103"},
						Level39MString2:                   map[string]string{"L39-N12-k111": "L39-N12-v111", "L39-N12-k112": "L39-N12-v112", "L39-N12-k113": "L39-N12-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N12-e101", FString2: "L39-N12-e102", FString3: "L39-N12-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N12-e111", FString2: "L39-N12-e112", FString3: "L39-N12-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N12-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N12-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N12-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N12-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N12-i111", IgnoreFieldsString2: "L39-N12-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N12-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N12-i131"},
					}
					level38N24 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N24-s101",
						Level38FString2:                   "L38-N24-s111",
						Level38PString1:                   utils.PointerString("L38-N24-p101"),
						Level38PString2:                   utils.PointerString("L38-N24-p111"),
						Level38RString1:                   []string{"L38-N24-r101", "L38-N24-r102", "L38-N24-r103"},
						Level38RString2:                   []string{"L38-N24-r111", "L38-N24-r112", "L38-N24-r113"},
						Level38MString1:                   map[string]string{"L38-N24-k101": "L38-N24-v101", "L38-N24-k102": "L38-N24-v102", "L38-N24-k103": "L38-N24-v103"},
						Level38MString2:                   map[string]string{"L38-N24-k111": "L38-N24-v111", "L38-N24-k112": "L38-N24-v112", "L38-N24-k113": "L38-N24-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N24-e101", FString2: "L38-N24-e102", FString3: "L38-N24-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N24-e111", FString2: "L38-N24-e112", FString3: "L38-N24-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N24-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N24-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N24-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N24-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N24-i111", IgnoreFieldsString2: "L38-N24-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N24-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N24-i131"},
					}
					level40N12 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N12-s101",
						Level40FString2:                   "L40-N12-s111",
						Level40PString1:                   utils.PointerString("L40-N12-p101"),
						Level40PString2:                   utils.PointerString("L40-N12-p111"),
						Level40RString1:                   []string{"L40-N12-r101", "L40-N12-r102", "L40-N12-r103"},
						Level40RString2:                   []string{"L40-N12-r111", "L40-N12-r112", "L40-N12-r113"},
						Level40MString1:                   map[string]string{"L40-N12-k101": "L40-N12-v101", "L40-N12-k102": "L40-N12-v102", "L40-N12-k103": "L40-N12-v103"},
						Level40MString2:                   map[string]string{"L40-N12-k111": "L40-N12-v111", "L40-N12-k112": "L40-N12-v112", "L40-N12-k113": "L40-N12-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N12-e101", FString2: "L40-N12-e102", FString3: "L40-N12-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N12-e111", FString2: "L40-N12-e112", FString3: "L40-N12-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N12-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N12-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N12-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N12-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N12-i111", IgnoreFieldsString2: "L40-N12-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N12-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N12-i131"},
					}
				}
				level22N12 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N12-s101",
					Level22FString2:                   "L22-N12-s111",
					Level22PString1:                   utils.PointerString("L22-N12-p101"),
					Level22PString2:                   utils.PointerString("L22-N12-p111"),
					Level22RString1:                   []string{"L22-N12-r101", "L22-N12-r102", "L22-N12-r103"},
					Level22RString2:                   []string{"L22-N12-r111", "L22-N12-r112", "L22-N12-r113"},
					Level22MString1:                   map[string]string{"L22-N12-k101": "L22-N12-v101", "L22-N12-k102": "L22-N12-v102", "L22-N12-k103": "L22-N12-v103"},
					Level22MString2:                   map[string]string{"L22-N12-k111": "L22-N12-v111", "L22-N12-k112": "L22-N12-v112", "L22-N12-k113": "L22-N12-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N12-e101", FString2: "L22-N12-e102", FString3: "L22-N12-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N12-e111", FString2: "L22-N12-e112", FString3: "L22-N12-e113"},
					Level22FMessageInline38:           level38N23,
					Level22FMessageInline39:           level39N12,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N24},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N12},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N12-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N12-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N12-i111", IgnoreFieldsString2: "L22-N12-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N12-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N12-i131"},
				}
				level24N06 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N06-s101",
					Level24FString2:                   "L24-N06-s111",
					Level24PString1:                   utils.PointerString("L24-N06-p101"),
					Level24PString2:                   utils.PointerString("L24-N06-p111"),
					Level24RString1:                   []string{"L24-N06-r101", "L24-N06-r102", "L24-N06-r103"},
					Level24RString2:                   []string{"L24-N06-r111", "L24-N06-r112", "L24-N06-r113"},
					Level24MString1:                   map[string]string{"L24-N06-k101": "L24-N06-v101", "L24-N06-k102": "L24-N06-v102", "L24-N06-k103": "L24-N06-v103"},
					Level24MString2:                   map[string]string{"L24-N06-k111": "L24-N06-v111", "L24-N06-k112": "L24-N06-v112", "L24-N06-k113": "L24-N06-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N06-e101", FString2: "L24-N06-e102", FString3: "L24-N06-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N06-e111", FString2: "L24-N06-e112", FString3: "L24-N06-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N06-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N06-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N06-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N06-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N06-i111", IgnoreFieldsString2: "L24-N06-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N06-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N06-i131"},
				}
			}
			level14N06 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N06-s101",
				Level14FString2:                   "L14-N06-s111",
				Level14PString1:                   utils.PointerString("L14-N06-p101"),
				Level14PString2:                   utils.PointerString("L14-N06-p111"),
				Level14RString1:                   []string{"L14-N06-r101", "L14-N06-r102", "L14-N06-r103"},
				Level14RString2:                   []string{"L14-N06-r111", "L14-N06-r112", "L14-N06-r113"},
				Level14MString1:                   map[string]string{"L14-N06-k101": "L14-N06-v101", "L14-N06-k102": "L14-N06-v102", "L14-N06-k103": "L14-N06-v103"},
				Level14MString2:                   map[string]string{"L14-N06-k111": "L14-N06-v111", "L14-N06-k112": "L14-N06-v112", "L14-N06-k113": "L14-N06-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N06-e101", FString2: "L14-N06-e102", FString3: "L14-N06-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N06-e111", FString2: "L14-N06-e112", FString3: "L14-N06-e113"},
				Level14FMessageInline22:           level22N11,
				Level14FMessageInline23:           level23N06,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N12},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N06},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N06-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N06-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N06-i111", IgnoreFieldsString2: "L14-N06-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N06-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N06-i131"},
			}
			{ // message for level16N03
				{ // message for level26N05
					level30N09 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N09-s101",
						Level30FString2:                   "L30-N09-s111",
						Level30PString1:                   utils.PointerString("L30-N09-p101"),
						Level30PString2:                   utils.PointerString("L30-N09-p111"),
						Level30RString1:                   []string{"L30-N09-r101", "L30-N09-r102", "L30-N09-r103"},
						Level30RString2:                   []string{"L30-N09-r111", "L30-N09-r112", "L30-N09-r113"},
						Level30MString1:                   map[string]string{"L30-N09-k101": "L30-N09-v101", "L30-N09-k102": "L30-N09-v102", "L30-N09-k103": "L30-N09-v103"},
						Level30MString2:                   map[string]string{"L30-N09-k111": "L30-N09-v111", "L30-N09-k112": "L30-N09-v112", "L30-N09-k113": "L30-N09-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N09-e101", FString2: "L30-N09-e102", FString3: "L30-N09-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N09-e111", FString2: "L30-N09-e112", FString3: "L30-N09-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N09-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N09-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N09-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N09-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N09-i111", IgnoreFieldsString2: "L30-N09-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N09-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N09-i131"},
					}
					level31N05 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N05-s101",
						Level31FString2:                   "L31-N05-s111",
						Level31PString1:                   utils.PointerString("L31-N05-p101"),
						Level31PString2:                   utils.PointerString("L31-N05-p111"),
						Level31RString1:                   []string{"L31-N05-r101", "L31-N05-r102", "L31-N05-r103"},
						Level31RString2:                   []string{"L31-N05-r111", "L31-N05-r112", "L31-N05-r113"},
						Level31MString1:                   map[string]string{"L31-N05-k101": "L31-N05-v101", "L31-N05-k102": "L31-N05-v102", "L31-N05-k103": "L31-N05-v103"},
						Level31MString2:                   map[string]string{"L31-N05-k111": "L31-N05-v111", "L31-N05-k112": "L31-N05-v112", "L31-N05-k113": "L31-N05-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N05-e101", FString2: "L31-N05-e102", FString3: "L31-N05-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N05-e111", FString2: "L31-N05-e112", FString3: "L31-N05-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N05-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N05-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N05-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N05-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N05-i111", IgnoreFieldsString2: "L31-N05-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N05-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N05-i131"},
					}
					level30N10 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N10-s101",
						Level30FString2:                   "L30-N10-s111",
						Level30PString1:                   utils.PointerString("L30-N10-p101"),
						Level30PString2:                   utils.PointerString("L30-N10-p111"),
						Level30RString1:                   []string{"L30-N10-r101", "L30-N10-r102", "L30-N10-r103"},
						Level30RString2:                   []string{"L30-N10-r111", "L30-N10-r112", "L30-N10-r113"},
						Level30MString1:                   map[string]string{"L30-N10-k101": "L30-N10-v101", "L30-N10-k102": "L30-N10-v102", "L30-N10-k103": "L30-N10-v103"},
						Level30MString2:                   map[string]string{"L30-N10-k111": "L30-N10-v111", "L30-N10-k112": "L30-N10-v112", "L30-N10-k113": "L30-N10-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N10-e101", FString2: "L30-N10-e102", FString3: "L30-N10-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N10-e111", FString2: "L30-N10-e112", FString3: "L30-N10-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N10-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N10-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N10-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N10-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N10-i111", IgnoreFieldsString2: "L30-N10-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N10-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N10-i131"},
					}
					level32N05 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N05-s101",
						Level32FString2:                   "L32-N05-s111",
						Level32PString1:                   utils.PointerString("L32-N05-p101"),
						Level32PString2:                   utils.PointerString("L32-N05-p111"),
						Level32RString1:                   []string{"L32-N05-r101", "L32-N05-r102", "L32-N05-r103"},
						Level32RString2:                   []string{"L32-N05-r111", "L32-N05-r112", "L32-N05-r113"},
						Level32MString1:                   map[string]string{"L32-N05-k101": "L32-N05-v101", "L32-N05-k102": "L32-N05-v102", "L32-N05-k103": "L32-N05-v103"},
						Level32MString2:                   map[string]string{"L32-N05-k111": "L32-N05-v111", "L32-N05-k112": "L32-N05-v112", "L32-N05-k113": "L32-N05-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N05-e101", FString2: "L32-N05-e102", FString3: "L32-N05-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N05-e111", FString2: "L32-N05-e112", FString3: "L32-N05-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N05-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N05-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N05-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N05-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N05-i111", IgnoreFieldsString2: "L32-N05-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N05-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N05-i131"},
					}
				}
				level26N05 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N05-s101",
					Level26FString2:                   "L26-N05-s111",
					Level26PString1:                   utils.PointerString("L26-N05-p101"),
					Level26PString2:                   utils.PointerString("L26-N05-p111"),
					Level26RString1:                   []string{"L26-N05-r101", "L26-N05-r102", "L26-N05-r103"},
					Level26RString2:                   []string{"L26-N05-r111", "L26-N05-r112", "L26-N05-r113"},
					Level26MString1:                   map[string]string{"L26-N05-k101": "L26-N05-v101", "L26-N05-k102": "L26-N05-v102", "L26-N05-k103": "L26-N05-v103"},
					Level26MString2:                   map[string]string{"L26-N05-k111": "L26-N05-v111", "L26-N05-k112": "L26-N05-v112", "L26-N05-k113": "L26-N05-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N05-e101", FString2: "L26-N05-e102", FString3: "L26-N05-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N05-e111", FString2: "L26-N05-e112", FString3: "L26-N05-e113"},
					Level26FMessageInline30:           level30N09,
					Level26FMessageInline31:           level31N05,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N10},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N05},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N05-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N05-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N05-i111", IgnoreFieldsString2: "L26-N05-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N05-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N05-i131"},
				}
				level27N03 = &pbinline.MessageLevel27{
					Level27FString1:                   "L27-N03-s101",
					Level27FString2:                   "L27-N03-s111",
					Level27PString1:                   utils.PointerString("L27-N03-p101"),
					Level27PString2:                   utils.PointerString("L27-N03-p111"),
					Level27RString1:                   []string{"L27-N03-r101", "L27-N03-r102", "L27-N03-r103"},
					Level27RString2:                   []string{"L27-N03-r111", "L27-N03-r112", "L27-N03-r113"},
					Level27MString1:                   map[string]string{"L27-N03-k101": "L27-N03-v101", "L27-N03-k102": "L27-N03-v102", "L27-N03-k103": "L27-N03-v103"},
					Level27MString2:                   map[string]string{"L27-N03-k111": "L27-N03-v111", "L27-N03-k112": "L27-N03-v112", "L27-N03-k113": "L27-N03-v113"},
					Level27FMessageExtern1:            &pbexternal.Message1{FString1: "L27-N03-e101", FString2: "L27-N03-e102", FString3: "L27-N03-e103"},
					Level27FMessageExtern2:            &pbexternal.Message1{FString1: "L27-N03-e111", FString2: "L27-N03-e112", FString3: "L27-N03-e113"},
					Level27OneOfExtern1:               &pbinline.MessageLevel27_Level27One1String1{Level27One1String1: "L27-N03-es101"},
					Level27OneOfInline1:               &pbinline.MessageLevel27_Level27One2String1{Level27One2String1: "L27-N03-is101"},
					Level27OneOfOmitempty1:            &pbinline.MessageLevel27_Level27One3String1{Level27One3String1: "L27-N03-o101"},
					Level27FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level27FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L27-N03-i101"},
					Level27FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L27-N03-i111", IgnoreFieldsString2: "L27-N03-i112"},
					Level27OneOfIgnoreSelf1:           &pbinline.MessageLevel27_Level27One4String1{Level27One4String1: "L27-N03-i121"},
					Level27OneOfIgnoreParts1:          &pbinline.MessageLevel27_Level27One5String1{Level27One5String1: "L27-N03-i131"},
				}
				{ // message for level26N06
					level30N11 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N11-s101",
						Level30FString2:                   "L30-N11-s111",
						Level30PString1:                   utils.PointerString("L30-N11-p101"),
						Level30PString2:                   utils.PointerString("L30-N11-p111"),
						Level30RString1:                   []string{"L30-N11-r101", "L30-N11-r102", "L30-N11-r103"},
						Level30RString2:                   []string{"L30-N11-r111", "L30-N11-r112", "L30-N11-r113"},
						Level30MString1:                   map[string]string{"L30-N11-k101": "L30-N11-v101", "L30-N11-k102": "L30-N11-v102", "L30-N11-k103": "L30-N11-v103"},
						Level30MString2:                   map[string]string{"L30-N11-k111": "L30-N11-v111", "L30-N11-k112": "L30-N11-v112", "L30-N11-k113": "L30-N11-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N11-e101", FString2: "L30-N11-e102", FString3: "L30-N11-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N11-e111", FString2: "L30-N11-e112", FString3: "L30-N11-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N11-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N11-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N11-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N11-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N11-i111", IgnoreFieldsString2: "L30-N11-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N11-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N11-i131"},
					}
					level31N06 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N06-s101",
						Level31FString2:                   "L31-N06-s111",
						Level31PString1:                   utils.PointerString("L31-N06-p101"),
						Level31PString2:                   utils.PointerString("L31-N06-p111"),
						Level31RString1:                   []string{"L31-N06-r101", "L31-N06-r102", "L31-N06-r103"},
						Level31RString2:                   []string{"L31-N06-r111", "L31-N06-r112", "L31-N06-r113"},
						Level31MString1:                   map[string]string{"L31-N06-k101": "L31-N06-v101", "L31-N06-k102": "L31-N06-v102", "L31-N06-k103": "L31-N06-v103"},
						Level31MString2:                   map[string]string{"L31-N06-k111": "L31-N06-v111", "L31-N06-k112": "L31-N06-v112", "L31-N06-k113": "L31-N06-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N06-e101", FString2: "L31-N06-e102", FString3: "L31-N06-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N06-e111", FString2: "L31-N06-e112", FString3: "L31-N06-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N06-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N06-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N06-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N06-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N06-i111", IgnoreFieldsString2: "L31-N06-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N06-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N06-i131"},
					}
					level30N12 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N12-s101",
						Level30FString2:                   "L30-N12-s111",
						Level30PString1:                   utils.PointerString("L30-N12-p101"),
						Level30PString2:                   utils.PointerString("L30-N12-p111"),
						Level30RString1:                   []string{"L30-N12-r101", "L30-N12-r102", "L30-N12-r103"},
						Level30RString2:                   []string{"L30-N12-r111", "L30-N12-r112", "L30-N12-r113"},
						Level30MString1:                   map[string]string{"L30-N12-k101": "L30-N12-v101", "L30-N12-k102": "L30-N12-v102", "L30-N12-k103": "L30-N12-v103"},
						Level30MString2:                   map[string]string{"L30-N12-k111": "L30-N12-v111", "L30-N12-k112": "L30-N12-v112", "L30-N12-k113": "L30-N12-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N12-e101", FString2: "L30-N12-e102", FString3: "L30-N12-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N12-e111", FString2: "L30-N12-e112", FString3: "L30-N12-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N12-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N12-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N12-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N12-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N12-i111", IgnoreFieldsString2: "L30-N12-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N12-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N12-i131"},
					}
					level32N06 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N06-s101",
						Level32FString2:                   "L32-N06-s111",
						Level32PString1:                   utils.PointerString("L32-N06-p101"),
						Level32PString2:                   utils.PointerString("L32-N06-p111"),
						Level32RString1:                   []string{"L32-N06-r101", "L32-N06-r102", "L32-N06-r103"},
						Level32RString2:                   []string{"L32-N06-r111", "L32-N06-r112", "L32-N06-r113"},
						Level32MString1:                   map[string]string{"L32-N06-k101": "L32-N06-v101", "L32-N06-k102": "L32-N06-v102", "L32-N06-k103": "L32-N06-v103"},
						Level32MString2:                   map[string]string{"L32-N06-k111": "L32-N06-v111", "L32-N06-k112": "L32-N06-v112", "L32-N06-k113": "L32-N06-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N06-e101", FString2: "L32-N06-e102", FString3: "L32-N06-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N06-e111", FString2: "L32-N06-e112", FString3: "L32-N06-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N06-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N06-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N06-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N06-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N06-i111", IgnoreFieldsString2: "L32-N06-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N06-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N06-i131"},
					}
				}
				level26N06 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N06-s101",
					Level26FString2:                   "L26-N06-s111",
					Level26PString1:                   utils.PointerString("L26-N06-p101"),
					Level26PString2:                   utils.PointerString("L26-N06-p111"),
					Level26RString1:                   []string{"L26-N06-r101", "L26-N06-r102", "L26-N06-r103"},
					Level26RString2:                   []string{"L26-N06-r111", "L26-N06-r112", "L26-N06-r113"},
					Level26MString1:                   map[string]string{"L26-N06-k101": "L26-N06-v101", "L26-N06-k102": "L26-N06-v102", "L26-N06-k103": "L26-N06-v103"},
					Level26MString2:                   map[string]string{"L26-N06-k111": "L26-N06-v111", "L26-N06-k112": "L26-N06-v112", "L26-N06-k113": "L26-N06-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N06-e101", FString2: "L26-N06-e102", FString3: "L26-N06-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N06-e111", FString2: "L26-N06-e112", FString3: "L26-N06-e113"},
					Level26FMessageInline30:           level30N11,
					Level26FMessageInline31:           level31N06,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N12},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N06},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N06-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N06-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N06-i111", IgnoreFieldsString2: "L26-N06-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N06-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N06-i131"},
				}
				{ // message for level28N03
					level34N05 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N05-s101",
						Level34FString2:                   "L34-N05-s111",
						Level34PString1:                   utils.PointerString("L34-N05-p101"),
						Level34PString2:                   utils.PointerString("L34-N05-p111"),
						Level34RString1:                   []string{"L34-N05-r101", "L34-N05-r102", "L34-N05-r103"},
						Level34RString2:                   []string{"L34-N05-r111", "L34-N05-r112", "L34-N05-r113"},
						Level34MString1:                   map[string]string{"L34-N05-k101": "L34-N05-v101", "L34-N05-k102": "L34-N05-v102", "L34-N05-k103": "L34-N05-v103"},
						Level34MString2:                   map[string]string{"L34-N05-k111": "L34-N05-v111", "L34-N05-k112": "L34-N05-v112", "L34-N05-k113": "L34-N05-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N05-e101", FString2: "L34-N05-e102", FString3: "L34-N05-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N05-e111", FString2: "L34-N05-e112", FString3: "L34-N05-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N05-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N05-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N05-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N05-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N05-i111", IgnoreFieldsString2: "L34-N05-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N05-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N05-i131"},
					}
					level35N03 = &pbinline.MessageLevel35{
						Level35FString1:                   "L35-N03-s101",
						Level35FString2:                   "L35-N03-s111",
						Level35PString1:                   utils.PointerString("L35-N03-p101"),
						Level35PString2:                   utils.PointerString("L35-N03-p111"),
						Level35RString1:                   []string{"L35-N03-r101", "L35-N03-r102", "L35-N03-r103"},
						Level35RString2:                   []string{"L35-N03-r111", "L35-N03-r112", "L35-N03-r113"},
						Level35MString1:                   map[string]string{"L35-N03-k101": "L35-N03-v101", "L35-N03-k102": "L35-N03-v102", "L35-N03-k103": "L35-N03-v103"},
						Level35MString2:                   map[string]string{"L35-N03-k111": "L35-N03-v111", "L35-N03-k112": "L35-N03-v112", "L35-N03-k113": "L35-N03-v113"},
						Level35FMessageExtern1:            &pbexternal.Message1{FString1: "L35-N03-e101", FString2: "L35-N03-e102", FString3: "L35-N03-e103"},
						Level35FMessageExtern2:            &pbexternal.Message1{FString1: "L35-N03-e111", FString2: "L35-N03-e112", FString3: "L35-N03-e113"},
						Level35OneOfExtern1:               &pbinline.MessageLevel35_Level35One1String1{Level35One1String1: "L35-N03-es101"},
						Level35OneOfInline1:               &pbinline.MessageLevel35_Level35One2String1{Level35One2String1: "L35-N03-is101"},
						Level35OneOfOmitempty1:            &pbinline.MessageLevel35_Level35One3String1{Level35One3String1: "L35-N03-o101"},
						Level35FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level35FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L35-N03-i101"},
						Level35FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L35-N03-i111", IgnoreFieldsString2: "L35-N03-i112"},
						Level35OneOfIgnoreSelf1:           &pbinline.MessageLevel35_Level35One4String1{Level35One4String1: "L35-N03-i121"},
						Level35OneOfIgnoreParts1:          &pbinline.MessageLevel35_Level35One5String1{Level35One5String1: "L35-N03-i131"},
					}
					level34N06 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N06-s101",
						Level34FString2:                   "L34-N06-s111",
						Level34PString1:                   utils.PointerString("L34-N06-p101"),
						Level34PString2:                   utils.PointerString("L34-N06-p111"),
						Level34RString1:                   []string{"L34-N06-r101", "L34-N06-r102", "L34-N06-r103"},
						Level34RString2:                   []string{"L34-N06-r111", "L34-N06-r112", "L34-N06-r113"},
						Level34MString1:                   map[string]string{"L34-N06-k101": "L34-N06-v101", "L34-N06-k102": "L34-N06-v102", "L34-N06-k103": "L34-N06-v103"},
						Level34MString2:                   map[string]string{"L34-N06-k111": "L34-N06-v111", "L34-N06-k112": "L34-N06-v112", "L34-N06-k113": "L34-N06-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N06-e101", FString2: "L34-N06-e102", FString3: "L34-N06-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N06-e111", FString2: "L34-N06-e112", FString3: "L34-N06-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N06-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N06-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N06-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N06-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N06-i111", IgnoreFieldsString2: "L34-N06-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N06-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N06-i131"},
					}
					level36N03 = &pbinline.MessageLevel36{
						Level36FString1:                   "L36-N03-s101",
						Level36FString2:                   "L36-N03-s111",
						Level36PString1:                   utils.PointerString("L36-N03-p101"),
						Level36PString2:                   utils.PointerString("L36-N03-p111"),
						Level36RString1:                   []string{"L36-N03-r101", "L36-N03-r102", "L36-N03-r103"},
						Level36RString2:                   []string{"L36-N03-r111", "L36-N03-r112", "L36-N03-r113"},
						Level36MString1:                   map[string]string{"L36-N03-k101": "L36-N03-v101", "L36-N03-k102": "L36-N03-v102", "L36-N03-k103": "L36-N03-v103"},
						Level36MString2:                   map[string]string{"L36-N03-k111": "L36-N03-v111", "L36-N03-k112": "L36-N03-v112", "L36-N03-k113": "L36-N03-v113"},
						Level36FMessageExtern1:            &pbexternal.Message1{FString1: "L36-N03-e101", FString2: "L36-N03-e102", FString3: "L36-N03-e103"},
						Level36FMessageExtern2:            &pbexternal.Message1{FString1: "L36-N03-e111", FString2: "L36-N03-e112", FString3: "L36-N03-e113"},
						Level36OneOfExtern1:               &pbinline.MessageLevel36_Level36One1String1{Level36One1String1: "L36-N03-es101"},
						Level36OneOfInline1:               &pbinline.MessageLevel36_Level36One2String1{Level36One2String1: "L36-N03-is101"},
						Level36OneOfOmitempty1:            &pbinline.MessageLevel36_Level36One3String1{Level36One3String1: "L36-N03-o101"},
						Level36FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level36FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L36-N03-i101"},
						Level36FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L36-N03-i111", IgnoreFieldsString2: "L36-N03-i112"},
						Level36OneOfIgnoreSelf1:           &pbinline.MessageLevel36_Level36One4String1{Level36One4String1: "L36-N03-i121"},
						Level36OneOfIgnoreParts1:          &pbinline.MessageLevel36_Level36One5String1{Level36One5String1: "L36-N03-i131"},
					}
				}
				level28N03 = &pbinline.MessageLevel28{
					Level28FString1:                   "L28-N03-s101",
					Level28FString2:                   "L28-N03-s111",
					Level28PString1:                   utils.PointerString("L28-N03-p101"),
					Level28PString2:                   utils.PointerString("L28-N03-p111"),
					Level28RString1:                   []string{"L28-N03-r101", "L28-N03-r102", "L28-N03-r103"},
					Level28RString2:                   []string{"L28-N03-r111", "L28-N03-r112", "L28-N03-r113"},
					Level28MString1:                   map[string]string{"L28-N03-k101": "L28-N03-v101", "L28-N03-k102": "L28-N03-v102", "L28-N03-k103": "L28-N03-v103"},
					Level28MString2:                   map[string]string{"L28-N03-k111": "L28-N03-v111", "L28-N03-k112": "L28-N03-v112", "L28-N03-k113": "L28-N03-v113"},
					Level28FMessageExtern1:            &pbexternal.Message1{FString1: "L28-N03-e101", FString2: "L28-N03-e102", FString3: "L28-N03-e103"},
					Level28FMessageExtern2:            &pbexternal.Message1{FString1: "L28-N03-e111", FString2: "L28-N03-e112", FString3: "L28-N03-e113"},
					Level28FMessageInline34:           level34N05,
					Level28FMessageInline35:           level35N03,
					Level28OneOfExtern1:               &pbinline.MessageLevel28_Level28One1MessageInline34{Level28One1MessageInline34: level34N06},
					Level28OneOfInline1:               &pbinline.MessageLevel28_Level28One2MessageInline36{Level28One2MessageInline36: level36N03},
					Level28OneOfOmitempty1:            &pbinline.MessageLevel28_Level28One3String1{Level28One3String1: "L28-N03-o101"},
					Level28FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level28FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L28-N03-i101"},
					Level28FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L28-N03-i111", IgnoreFieldsString2: "L28-N03-i112"},
					Level28OneOfIgnoreSelf1:           &pbinline.MessageLevel28_Level28One4String1{Level28One4String1: "L28-N03-i121"},
					Level28OneOfIgnoreParts1:          &pbinline.MessageLevel28_Level28One5String1{Level28One5String1: "L28-N03-i131"},
				}
			}
			level16N03 = &pbinline.MessageLevel16{
				Level16FString1:                   "L16-N03-s101",
				Level16FString2:                   "L16-N03-s111",
				Level16PString1:                   utils.PointerString("L16-N03-p101"),
				Level16PString2:                   utils.PointerString("L16-N03-p111"),
				Level16RString1:                   []string{"L16-N03-r101", "L16-N03-r102", "L16-N03-r103"},
				Level16RString2:                   []string{"L16-N03-r111", "L16-N03-r112", "L16-N03-r113"},
				Level16MString1:                   map[string]string{"L16-N03-k101": "L16-N03-v101", "L16-N03-k102": "L16-N03-v102", "L16-N03-k103": "L16-N03-v103"},
				Level16MString2:                   map[string]string{"L16-N03-k111": "L16-N03-v111", "L16-N03-k112": "L16-N03-v112", "L16-N03-k113": "L16-N03-v113"},
				Level16FMessageExtern1:            &pbexternal.Message1{FString1: "L16-N03-e101", FString2: "L16-N03-e102", FString3: "L16-N03-e103"},
				Level16FMessageExtern2:            &pbexternal.Message1{FString1: "L16-N03-e111", FString2: "L16-N03-e112", FString3: "L16-N03-e113"},
				Level16FMessageInline26:           level26N05,
				Level16FMessageInline27:           level27N03,
				Level16OneOfExtern1:               &pbinline.MessageLevel16_Level16One1MessageInline26{Level16One1MessageInline26: level26N06},
				Level16OneOfInline1:               &pbinline.MessageLevel16_Level16One2MessageInline28{Level16One2MessageInline28: level28N03},
				Level16OneOfOmitempty1:            &pbinline.MessageLevel16_Level16One3String1{Level16One3String1: "L16-N03-o101"},
				Level16FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level16FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L16-N03-i101"},
				Level16FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L16-N03-i111", IgnoreFieldsString2: "L16-N03-i112"},
				Level16OneOfIgnoreSelf1:           &pbinline.MessageLevel16_Level16One4String1{Level16One4String1: "L16-N03-i121"},
				Level16OneOfIgnoreParts1:          &pbinline.MessageLevel16_Level16One5String1{Level16One5String1: "L16-N03-i131"},
			}
		}
		level06N03 = &pbinline.MessageLevel06{
			Level06FString1:                   "L06-N03-s101",
			Level06FString2:                   "L06-N03-s111",
			Level06PString1:                   utils.PointerString("L06-N03-p101"),
			Level06PString2:                   utils.PointerString("L06-N03-p111"),
			Level06RString1:                   []string{"L06-N03-r101", "L06-N03-r102", "L06-N03-r103"},
			Level06RString2:                   []string{"L06-N03-r111", "L06-N03-r112", "L06-N03-r113"},
			Level06MString1:                   map[string]string{"L06-N03-k101": "L06-N03-v101", "L06-N03-k102": "L06-N03-v102", "L06-N03-k103": "L06-N03-v103"},
			Level06MString2:                   map[string]string{"L06-N03-k111": "L06-N03-v111", "L06-N03-k112": "L06-N03-v112", "L06-N03-k113": "L06-N03-v113"},
			Level06FMessageExtern1:            &pbexternal.Message1{FString1: "L06-N03-e101", FString2: "L06-N03-e102", FString3: "L06-N03-e103"},
			Level06FMessageExtern2:            &pbexternal.Message1{FString1: "L06-N03-e111", FString2: "L06-N03-e112", FString3: "L06-N03-e113"},
			Level06FMessageInline14:           level14N05,
			Level06FMessageInline15:           level15N03,
			Level06OneOfExtern1:               &pbinline.MessageLevel06_Level06One1MessageInline14{Level06One1MessageInline14: level14N06},
			Level06OneOfInline1:               &pbinline.MessageLevel06_Level4One2MessageInline16{Level4One2MessageInline16: level16N03},
			Level06OneOfOmitempty1:            &pbinline.MessageLevel06_Level06One3String1{Level06One3String1: "L06-N03-o101"},
			Level06FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level06FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L06-N03-i101"},
			Level06FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L06-N03-i111", IgnoreFieldsString2: "L06-N03-i112"},
			Level06OneOfIgnoreSelf1:           &pbinline.MessageLevel06_Level06One4String1{Level06One4String1: "L06-N03-i121"},
			Level06OneOfIgnoreParts1:          &pbinline.MessageLevel06_Level06One5String1{Level06One5String1: "L06-N03-i131"},
		}
		level07N02 = &pbinline.MessageLevel07{
			Level07FString1:                   "L07-N02-s101",
			Level07FString2:                   "L07-N02-s111",
			Level07PString1:                   utils.PointerString("L07-N02-p101"),
			Level07PString2:                   utils.PointerString("L07-N02-p111"),
			Level07RString1:                   []string{"L07-N02-r101", "L07-N02-r102", "L07-N02-r103"},
			Level07RString2:                   []string{"L07-N02-r111", "L07-N02-r112", "L07-N02-r113"},
			Level07MString1:                   map[string]string{"L07-N02-k101": "L07-N02-v101", "L07-N02-k102": "L07-N02-v102", "L07-N02-k103": "L07-N02-v103"},
			Level07MString2:                   map[string]string{"L07-N02-k111": "L07-N02-v111", "L07-N02-k112": "L07-N02-v112", "L07-N02-k113": "L07-N02-v113"},
			Level07FMessageExtern1:            &pbexternal.Message1{FString1: "L07-N02-e101", FString2: "L07-N02-e102", FString3: "L07-N02-e103"},
			Level07FMessageExtern2:            &pbexternal.Message1{FString1: "L07-N02-e111", FString2: "L07-N02-e112", FString3: "L07-N02-e113"},
			Level07OneOfExtern1:               &pbinline.MessageLevel07_Level07One1String1{Level07One1String1: "L07-N02-es101"},
			Level07OneOfInline1:               &pbinline.MessageLevel07_Level07One2String1{Level07One2String1: "L07-N02-is101"},
			Level07OneOfOmitempty1:            &pbinline.MessageLevel07_Level07One3String1{Level07One3String1: "L07-N02-o101"},
			Level07FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level07FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L07-N02-i101"},
			Level07FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L07-N02-i111", IgnoreFieldsString2: "L07-N02-i112"},
			Level07OneOfIgnoreSelf1:           &pbinline.MessageLevel07_Level07One4String1{Level07One4String1: "L07-N02-i121"},
			Level07OneOfIgnoreParts1:          &pbinline.MessageLevel07_Level07One5String1{Level07One5String1: "L07-N02-i131"},
		}
		{ // message in level06N04
			{ // message for level14N07
				{ // message for level22N13
					level38N25 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N25-s101",
						Level38FString2:                   "L38-N25-s111",
						Level38PString1:                   utils.PointerString("L38-N25-p101"),
						Level38PString2:                   utils.PointerString("L38-N25-p111"),
						Level38RString1:                   []string{"L38-N25-r101", "L38-N25-r102", "L38-N25-r103"},
						Level38RString2:                   []string{"L38-N25-r111", "L38-N25-r112", "L38-N25-r113"},
						Level38MString1:                   map[string]string{"L38-N25-k101": "L38-N25-v101", "L38-N25-k102": "L38-N25-v102", "L38-N25-k103": "L38-N25-v103"},
						Level38MString2:                   map[string]string{"L38-N25-k111": "L38-N25-v111", "L38-N25-k112": "L38-N25-v112", "L38-N25-k113": "L38-N25-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N25-e101", FString2: "L38-N25-e102", FString3: "L38-N25-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N25-e111", FString2: "L38-N25-e112", FString3: "L38-N25-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N25-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N25-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N25-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N25-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N25-i111", IgnoreFieldsString2: "L38-N25-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N25-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N25-i131"},
					}
					level39N13 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N13-s101",
						Level39FString2:                   "L39-N13-s111",
						Level39PString1:                   utils.PointerString("L39-N13-p101"),
						Level39PString2:                   utils.PointerString("L39-N13-p111"),
						Level39RString1:                   []string{"L39-N13-r101", "L39-N13-r102", "L39-N13-r103"},
						Level39RString2:                   []string{"L39-N13-r111", "L39-N13-r112", "L39-N13-r113"},
						Level39MString1:                   map[string]string{"L39-N13-k101": "L39-N13-v101", "L39-N13-k102": "L39-N13-v102", "L39-N13-k103": "L39-N13-v103"},
						Level39MString2:                   map[string]string{"L39-N13-k111": "L39-N13-v111", "L39-N13-k112": "L39-N13-v112", "L39-N13-k113": "L39-N13-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N13-e101", FString2: "L39-N13-e102", FString3: "L39-N13-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N13-e111", FString2: "L39-N13-e112", FString3: "L39-N13-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N13-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N13-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N13-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N13-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N13-i111", IgnoreFieldsString2: "L39-N13-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N13-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N13-i131"},
					}
					level38N26 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N26-s101",
						Level38FString2:                   "L38-N26-s111",
						Level38PString1:                   utils.PointerString("L38-N26-p101"),
						Level38PString2:                   utils.PointerString("L38-N26-p111"),
						Level38RString1:                   []string{"L38-N26-r101", "L38-N26-r102", "L38-N26-r103"},
						Level38RString2:                   []string{"L38-N26-r111", "L38-N26-r112", "L38-N26-r113"},
						Level38MString1:                   map[string]string{"L38-N26-k101": "L38-N26-v101", "L38-N26-k102": "L38-N26-v102", "L38-N26-k103": "L38-N26-v103"},
						Level38MString2:                   map[string]string{"L38-N26-k111": "L38-N26-v111", "L38-N26-k112": "L38-N26-v112", "L38-N26-k113": "L38-N26-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N26-e101", FString2: "L38-N26-e102", FString3: "L38-N26-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N26-e111", FString2: "L38-N26-e112", FString3: "L38-N26-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N26-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N26-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N26-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N26-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N26-i111", IgnoreFieldsString2: "L38-N26-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N26-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N26-i131"},
					}
					level40N13 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N13-s101",
						Level40FString2:                   "L40-N13-s111",
						Level40PString1:                   utils.PointerString("L40-N13-p101"),
						Level40PString2:                   utils.PointerString("L40-N13-p111"),
						Level40RString1:                   []string{"L40-N13-r101", "L40-N13-r102", "L40-N13-r103"},
						Level40RString2:                   []string{"L40-N13-r111", "L40-N13-r112", "L40-N13-r113"},
						Level40MString1:                   map[string]string{"L40-N13-k101": "L40-N13-v101", "L40-N13-k102": "L40-N13-v102", "L40-N13-k103": "L40-N13-v103"},
						Level40MString2:                   map[string]string{"L40-N13-k111": "L40-N13-v111", "L40-N13-k112": "L40-N13-v112", "L40-N13-k113": "L40-N13-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N13-e101", FString2: "L40-N13-e102", FString3: "L40-N13-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N13-e111", FString2: "L40-N13-e112", FString3: "L40-N13-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N13-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N13-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N13-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N13-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N13-i111", IgnoreFieldsString2: "L40-N13-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N13-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N13-i131"},
					}
				}
				level22N13 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N13-s101",
					Level22FString2:                   "L22-N13-s111",
					Level22PString1:                   utils.PointerString("L22-N13-p101"),
					Level22PString2:                   utils.PointerString("L22-N13-p111"),
					Level22RString1:                   []string{"L22-N13-r101", "L22-N13-r102", "L22-N13-r103"},
					Level22RString2:                   []string{"L22-N13-r111", "L22-N13-r112", "L22-N13-r113"},
					Level22MString1:                   map[string]string{"L22-N13-k101": "L22-N13-v101", "L22-N13-k102": "L22-N13-v102", "L22-N13-k103": "L22-N13-v103"},
					Level22MString2:                   map[string]string{"L22-N13-k111": "L22-N13-v111", "L22-N13-k112": "L22-N13-v112", "L22-N13-k113": "L22-N13-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N13-e101", FString2: "L22-N13-e102", FString3: "L22-N13-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N13-e111", FString2: "L22-N13-e112", FString3: "L22-N13-e113"},
					Level22FMessageInline38:           level38N25,
					Level22FMessageInline39:           level39N13,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N26},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N13},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N13-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N13-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N13-i111", IgnoreFieldsString2: "L22-N13-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N13-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N13-i131"},
				}
				level23N07 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N07-s101",
					Level23FString2:                   "L23-N07-s111",
					Level23PString1:                   utils.PointerString("L23-N07-p101"),
					Level23PString2:                   utils.PointerString("L23-N07-p111"),
					Level23RString1:                   []string{"L23-N07-r101", "L23-N07-r102", "L23-N07-r103"},
					Level23RString2:                   []string{"L23-N07-r111", "L23-N07-r112", "L23-N07-r113"},
					Level23MString1:                   map[string]string{"L23-N07-k101": "L23-N07-v101", "L23-N07-k102": "L23-N07-v102", "L23-N07-k103": "L23-N07-v103"},
					Level23MString2:                   map[string]string{"L23-N07-k111": "L23-N07-v111", "L23-N07-k112": "L23-N07-v112", "L23-N07-k113": "L23-N07-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N07-e101", FString2: "L23-N07-e102", FString3: "L23-N07-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N07-e111", FString2: "L23-N07-e112", FString3: "L23-N07-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N07-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N07-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N07-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N07-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N07-i111", IgnoreFieldsString2: "L23-N07-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N07-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N07-i131"},
				}
				{ // message for level22N14
					level38N27 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N27-s101",
						Level38FString2:                   "L38-N27-s111",
						Level38PString1:                   utils.PointerString("L38-N27-p101"),
						Level38PString2:                   utils.PointerString("L38-N27-p111"),
						Level38RString1:                   []string{"L38-N27-r101", "L38-N27-r102", "L38-N27-r103"},
						Level38RString2:                   []string{"L38-N27-r111", "L38-N27-r112", "L38-N27-r113"},
						Level38MString1:                   map[string]string{"L38-N27-k101": "L38-N27-v101", "L38-N27-k102": "L38-N27-v102", "L38-N27-k103": "L38-N27-v103"},
						Level38MString2:                   map[string]string{"L38-N27-k111": "L38-N27-v111", "L38-N27-k112": "L38-N27-v112", "L38-N27-k113": "L38-N27-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N27-e101", FString2: "L38-N27-e102", FString3: "L38-N27-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N27-e111", FString2: "L38-N27-e112", FString3: "L38-N27-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N27-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N27-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N27-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N27-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N27-i111", IgnoreFieldsString2: "L38-N27-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N27-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N27-i131"},
					}
					level39N14 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N14-s101",
						Level39FString2:                   "L39-N14-s111",
						Level39PString1:                   utils.PointerString("L39-N14-p101"),
						Level39PString2:                   utils.PointerString("L39-N14-p111"),
						Level39RString1:                   []string{"L39-N14-r101", "L39-N14-r102", "L39-N14-r103"},
						Level39RString2:                   []string{"L39-N14-r111", "L39-N14-r112", "L39-N14-r113"},
						Level39MString1:                   map[string]string{"L39-N14-k101": "L39-N14-v101", "L39-N14-k102": "L39-N14-v102", "L39-N14-k103": "L39-N14-v103"},
						Level39MString2:                   map[string]string{"L39-N14-k111": "L39-N14-v111", "L39-N14-k112": "L39-N14-v112", "L39-N14-k113": "L39-N14-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N14-e101", FString2: "L39-N14-e102", FString3: "L39-N14-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N14-e111", FString2: "L39-N14-e112", FString3: "L39-N14-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N14-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N14-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N14-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N14-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N14-i111", IgnoreFieldsString2: "L39-N14-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N14-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N14-i131"},
					}
					level38N28 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N28-s101",
						Level38FString2:                   "L38-N28-s111",
						Level38PString1:                   utils.PointerString("L38-N28-p101"),
						Level38PString2:                   utils.PointerString("L38-N28-p111"),
						Level38RString1:                   []string{"L38-N28-r101", "L38-N28-r102", "L38-N28-r103"},
						Level38RString2:                   []string{"L38-N28-r111", "L38-N28-r112", "L38-N28-r113"},
						Level38MString1:                   map[string]string{"L38-N28-k101": "L38-N28-v101", "L38-N28-k102": "L38-N28-v102", "L38-N28-k103": "L38-N28-v103"},
						Level38MString2:                   map[string]string{"L38-N28-k111": "L38-N28-v111", "L38-N28-k112": "L38-N28-v112", "L38-N28-k113": "L38-N28-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N28-e101", FString2: "L38-N28-e102", FString3: "L38-N28-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N28-e111", FString2: "L38-N28-e112", FString3: "L38-N28-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N28-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N28-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N28-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N28-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N28-i111", IgnoreFieldsString2: "L38-N28-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N28-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N28-i131"},
					}
					level40N14 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N14-s101",
						Level40FString2:                   "L40-N14-s111",
						Level40PString1:                   utils.PointerString("L40-N14-p101"),
						Level40PString2:                   utils.PointerString("L40-N14-p111"),
						Level40RString1:                   []string{"L40-N14-r101", "L40-N14-r102", "L40-N14-r103"},
						Level40RString2:                   []string{"L40-N14-r111", "L40-N14-r112", "L40-N14-r113"},
						Level40MString1:                   map[string]string{"L40-N14-k101": "L40-N14-v101", "L40-N14-k102": "L40-N14-v102", "L40-N14-k103": "L40-N14-v103"},
						Level40MString2:                   map[string]string{"L40-N14-k111": "L40-N14-v111", "L40-N14-k112": "L40-N14-v112", "L40-N14-k113": "L40-N14-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N14-e101", FString2: "L40-N14-e102", FString3: "L40-N14-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N14-e111", FString2: "L40-N14-e112", FString3: "L40-N14-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N14-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N14-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N14-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N14-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N14-i111", IgnoreFieldsString2: "L40-N14-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N14-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N14-i131"},
					}
				}
				level22N14 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N14-s101",
					Level22FString2:                   "L22-N14-s111",
					Level22PString1:                   utils.PointerString("L22-N14-p101"),
					Level22PString2:                   utils.PointerString("L22-N14-p111"),
					Level22RString1:                   []string{"L22-N14-r101", "L22-N14-r102", "L22-N14-r103"},
					Level22RString2:                   []string{"L22-N14-r111", "L22-N14-r112", "L22-N14-r113"},
					Level22MString1:                   map[string]string{"L22-N14-k101": "L22-N14-v101", "L22-N14-k102": "L22-N14-v102", "L22-N14-k103": "L22-N14-v103"},
					Level22MString2:                   map[string]string{"L22-N14-k111": "L22-N14-v111", "L22-N14-k112": "L22-N14-v112", "L22-N14-k113": "L22-N14-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N14-e101", FString2: "L22-N14-e102", FString3: "L22-N14-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N14-e111", FString2: "L22-N14-e112", FString3: "L22-N14-e113"},
					Level22FMessageInline38:           level38N27,
					Level22FMessageInline39:           level39N14,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N28},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N14},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N14-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N14-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N14-i111", IgnoreFieldsString2: "L22-N14-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N14-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N14-i131"},
				}
				level24N07 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N07-s101",
					Level24FString2:                   "L24-N07-s111",
					Level24PString1:                   utils.PointerString("L24-N07-p101"),
					Level24PString2:                   utils.PointerString("L24-N07-p111"),
					Level24RString1:                   []string{"L24-N07-r101", "L24-N07-r102", "L24-N07-r103"},
					Level24RString2:                   []string{"L24-N07-r111", "L24-N07-r112", "L24-N07-r113"},
					Level24MString1:                   map[string]string{"L24-N07-k101": "L24-N07-v101", "L24-N07-k102": "L24-N07-v102", "L24-N07-k103": "L24-N07-v103"},
					Level24MString2:                   map[string]string{"L24-N07-k111": "L24-N07-v111", "L24-N07-k112": "L24-N07-v112", "L24-N07-k113": "L24-N07-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N07-e101", FString2: "L24-N07-e102", FString3: "L24-N07-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N07-e111", FString2: "L24-N07-e112", FString3: "L24-N07-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N07-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N07-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N07-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N07-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N07-i111", IgnoreFieldsString2: "L24-N07-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N07-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N07-i131"},
				}
			}
			level14N07 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N07-s101",
				Level14FString2:                   "L14-N07-s111",
				Level14PString1:                   utils.PointerString("L14-N07-p101"),
				Level14PString2:                   utils.PointerString("L14-N07-p111"),
				Level14RString1:                   []string{"L14-N07-r101", "L14-N07-r102", "L14-N07-r103"},
				Level14RString2:                   []string{"L14-N07-r111", "L14-N07-r112", "L14-N07-r113"},
				Level14MString1:                   map[string]string{"L14-N07-k101": "L14-N07-v101", "L14-N07-k102": "L14-N07-v102", "L14-N07-k103": "L14-N07-v103"},
				Level14MString2:                   map[string]string{"L14-N07-k111": "L14-N07-v111", "L14-N07-k112": "L14-N07-v112", "L14-N07-k113": "L14-N07-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N07-e101", FString2: "L14-N07-e102", FString3: "L14-N07-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N07-e111", FString2: "L14-N07-e112", FString3: "L14-N07-e113"},
				Level14FMessageInline22:           level22N13,
				Level14FMessageInline23:           level23N07,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N14},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N07},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N07-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N07-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N07-i111", IgnoreFieldsString2: "L14-N07-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N07-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N07-i131"},
			}
			level15N04 = &pbinline.MessageLevel15{
				Level15FString1:                   "L15-N04-s101",
				Level15FString2:                   "L15-N04-s111",
				Level15PString1:                   utils.PointerString("L15-N04-p101"),
				Level15PString2:                   utils.PointerString("L15-N04-p111"),
				Level15RString1:                   []string{"L15-N04-r101", "L15-N04-r102", "L15-N04-r103"},
				Level15RString2:                   []string{"L15-N04-r111", "L15-N04-r112", "L15-N04-r113"},
				Level15MString1:                   map[string]string{"L15-N04-k101": "L15-N04-v101", "L15-N04-k102": "L15-N04-v102", "L15-N04-k103": "L15-N04-v103"},
				Level15MString2:                   map[string]string{"L15-N04-k111": "L15-N04-v111", "L15-N04-k112": "L15-N04-v112", "L15-N04-k113": "L15-N04-v113"},
				Level15FMessageExtern1:            &pbexternal.Message1{FString1: "L15-N04-e101", FString2: "L15-N04-e102", FString3: "L15-N04-e103"},
				Level15FMessageExtern2:            &pbexternal.Message1{FString1: "L15-N04-e111", FString2: "L15-N04-e112", FString3: "L15-N04-e113"},
				Level15OneOfExtern1:               &pbinline.MessageLevel15_Level15One1String1{Level15One1String1: "L15-N04-es101"},
				Level15OneOfInline1:               &pbinline.MessageLevel15_Level15One2String1{Level15One2String1: "L15-N04-is101"},
				Level15OneOfOmitempty1:            &pbinline.MessageLevel15_Level15One3String1{Level15One3String1: "L15-N04-o101"},
				Level15FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level15FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L15-N04-i101"},
				Level15FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L15-N04-i111", IgnoreFieldsString2: "L15-N04-i112"},
				Level15OneOfIgnoreSelf1:           &pbinline.MessageLevel15_Level15One4String1{Level15One4String1: "L15-N04-i121"},
				Level15OneOfIgnoreParts1:          &pbinline.MessageLevel15_Level15One5String1{Level15One5String1: "L15-N04-i131"},
			}
			{ // message for level14N08
				{ // message for level22N15
					level38N29 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N29-s101",
						Level38FString2:                   "L38-N29-s111",
						Level38PString1:                   utils.PointerString("L38-N29-p101"),
						Level38PString2:                   utils.PointerString("L38-N29-p111"),
						Level38RString1:                   []string{"L38-N29-r101", "L38-N29-r102", "L38-N29-r103"},
						Level38RString2:                   []string{"L38-N29-r111", "L38-N29-r112", "L38-N29-r113"},
						Level38MString1:                   map[string]string{"L38-N29-k101": "L38-N29-v101", "L38-N29-k102": "L38-N29-v102", "L38-N29-k103": "L38-N29-v103"},
						Level38MString2:                   map[string]string{"L38-N29-k111": "L38-N29-v111", "L38-N29-k112": "L38-N29-v112", "L38-N29-k113": "L38-N29-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N29-e101", FString2: "L38-N29-e102", FString3: "L38-N29-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N29-e111", FString2: "L38-N29-e112", FString3: "L38-N29-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N29-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N29-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N29-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N29-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N29-i111", IgnoreFieldsString2: "L38-N29-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N29-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N29-i131"},
					}
					level39N15 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N15-s101",
						Level39FString2:                   "L39-N15-s111",
						Level39PString1:                   utils.PointerString("L39-N15-p101"),
						Level39PString2:                   utils.PointerString("L39-N15-p111"),
						Level39RString1:                   []string{"L39-N15-r101", "L39-N15-r102", "L39-N15-r103"},
						Level39RString2:                   []string{"L39-N15-r111", "L39-N15-r112", "L39-N15-r113"},
						Level39MString1:                   map[string]string{"L39-N15-k101": "L39-N15-v101", "L39-N15-k102": "L39-N15-v102", "L39-N15-k103": "L39-N15-v103"},
						Level39MString2:                   map[string]string{"L39-N15-k111": "L39-N15-v111", "L39-N15-k112": "L39-N15-v112", "L39-N15-k113": "L39-N15-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N15-e101", FString2: "L39-N15-e102", FString3: "L39-N15-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N15-e111", FString2: "L39-N15-e112", FString3: "L39-N15-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N15-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N15-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N15-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N15-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N15-i111", IgnoreFieldsString2: "L39-N15-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N15-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N15-i131"},
					}
					level38N30 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N30-s101",
						Level38FString2:                   "L38-N30-s111",
						Level38PString1:                   utils.PointerString("L38-N30-p101"),
						Level38PString2:                   utils.PointerString("L38-N30-p111"),
						Level38RString1:                   []string{"L38-N30-r101", "L38-N30-r102", "L38-N30-r103"},
						Level38RString2:                   []string{"L38-N30-r111", "L38-N30-r112", "L38-N30-r113"},
						Level38MString1:                   map[string]string{"L38-N30-k101": "L38-N30-v101", "L38-N30-k102": "L38-N30-v102", "L38-N30-k103": "L38-N30-v103"},
						Level38MString2:                   map[string]string{"L38-N30-k111": "L38-N30-v111", "L38-N30-k112": "L38-N30-v112", "L38-N30-k113": "L38-N30-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N30-e101", FString2: "L38-N30-e102", FString3: "L38-N30-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N30-e111", FString2: "L38-N30-e112", FString3: "L38-N30-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N30-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N30-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N30-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N30-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N30-i111", IgnoreFieldsString2: "L38-N30-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N30-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N30-i131"},
					}
					level40N15 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N15-s101",
						Level40FString2:                   "L40-N15-s111",
						Level40PString1:                   utils.PointerString("L40-N15-p101"),
						Level40PString2:                   utils.PointerString("L40-N15-p111"),
						Level40RString1:                   []string{"L40-N15-r101", "L40-N15-r102", "L40-N15-r103"},
						Level40RString2:                   []string{"L40-N15-r111", "L40-N15-r112", "L40-N15-r113"},
						Level40MString1:                   map[string]string{"L40-N15-k101": "L40-N15-v101", "L40-N15-k102": "L40-N15-v102", "L40-N15-k103": "L40-N15-v103"},
						Level40MString2:                   map[string]string{"L40-N15-k111": "L40-N15-v111", "L40-N15-k112": "L40-N15-v112", "L40-N15-k113": "L40-N15-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N15-e101", FString2: "L40-N15-e102", FString3: "L40-N15-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N15-e111", FString2: "L40-N15-e112", FString3: "L40-N15-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N15-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N15-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N15-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N15-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N15-i111", IgnoreFieldsString2: "L40-N15-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N15-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N15-i131"},
					}
				}
				level22N15 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N15-s101",
					Level22FString2:                   "L22-N15-s111",
					Level22PString1:                   utils.PointerString("L22-N15-p101"),
					Level22PString2:                   utils.PointerString("L22-N15-p111"),
					Level22RString1:                   []string{"L22-N15-r101", "L22-N15-r102", "L22-N15-r103"},
					Level22RString2:                   []string{"L22-N15-r111", "L22-N15-r112", "L22-N15-r113"},
					Level22MString1:                   map[string]string{"L22-N15-k101": "L22-N15-v101", "L22-N15-k102": "L22-N15-v102", "L22-N15-k103": "L22-N15-v103"},
					Level22MString2:                   map[string]string{"L22-N15-k111": "L22-N15-v111", "L22-N15-k112": "L22-N15-v112", "L22-N15-k113": "L22-N15-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N15-e101", FString2: "L22-N15-e102", FString3: "L22-N15-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N15-e111", FString2: "L22-N15-e112", FString3: "L22-N15-e113"},
					Level22FMessageInline38:           level38N29,
					Level22FMessageInline39:           level39N15,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N30},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N15},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N15-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N15-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N15-i111", IgnoreFieldsString2: "L22-N15-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N15-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N15-i131"},
				}
				level23N08 = &pbinline.MessageLevel23{
					Level23FString1:                   "L23-N08-s101",
					Level23FString2:                   "L23-N08-s111",
					Level23PString1:                   utils.PointerString("L23-N08-p101"),
					Level23PString2:                   utils.PointerString("L23-N08-p111"),
					Level23RString1:                   []string{"L23-N08-r101", "L23-N08-r102", "L23-N08-r103"},
					Level23RString2:                   []string{"L23-N08-r111", "L23-N08-r112", "L23-N08-r113"},
					Level23MString1:                   map[string]string{"L23-N08-k101": "L23-N08-v101", "L23-N08-k102": "L23-N08-v102", "L23-N08-k103": "L23-N08-v103"},
					Level23MString2:                   map[string]string{"L23-N08-k111": "L23-N08-v111", "L23-N08-k112": "L23-N08-v112", "L23-N08-k113": "L23-N08-v113"},
					Level23FMessageExtern1:            &pbexternal.Message1{FString1: "L23-N08-e101", FString2: "L23-N08-e102", FString3: "L23-N08-e103"},
					Level23FMessageExtern2:            &pbexternal.Message1{FString1: "L23-N08-e111", FString2: "L23-N08-e112", FString3: "L23-N08-e113"},
					Level23OneOfExtern1:               &pbinline.MessageLevel23_Level23One1String1{Level23One1String1: "L23-N08-es101"},
					Level23OneOfInline1:               &pbinline.MessageLevel23_Level23One2String1{Level23One2String1: "L23-N08-is101"},
					Level23OneOfOmitempty1:            &pbinline.MessageLevel23_Level23One3String1{Level23One3String1: "L23-N08-o101"},
					Level23FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level23FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L23-N08-i101"},
					Level23FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L23-N08-i111", IgnoreFieldsString2: "L23-N08-i112"},
					Level23OneOfIgnoreSelf1:           &pbinline.MessageLevel23_Level23One4String1{Level23One4String1: "L23-N08-i121"},
					Level23OneOfIgnoreParts1:          &pbinline.MessageLevel23_Level23One5String1{Level23One5String1: "L23-N08-i131"},
				}
				{ // message for level22N16
					level38N31 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N31-s101",
						Level38FString2:                   "L38-N31-s111",
						Level38PString1:                   utils.PointerString("L38-N31-p101"),
						Level38PString2:                   utils.PointerString("L38-N31-p111"),
						Level38RString1:                   []string{"L38-N31-r101", "L38-N31-r102", "L38-N31-r103"},
						Level38RString2:                   []string{"L38-N31-r111", "L38-N31-r112", "L38-N31-r113"},
						Level38MString1:                   map[string]string{"L38-N31-k101": "L38-N31-v101", "L38-N31-k102": "L38-N31-v102", "L38-N31-k103": "L38-N31-v103"},
						Level38MString2:                   map[string]string{"L38-N31-k111": "L38-N31-v111", "L38-N31-k112": "L38-N31-v112", "L38-N31-k113": "L38-N31-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N31-e101", FString2: "L38-N31-e102", FString3: "L38-N31-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N31-e111", FString2: "L38-N31-e112", FString3: "L38-N31-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N31-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N31-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N31-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N31-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N31-i111", IgnoreFieldsString2: "L38-N31-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N31-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N31-i131"},
					}
					level39N16 = &pbinline.MessageLevel39{
						Level39FString1:                   "L39-N16-s101",
						Level39FString2:                   "L39-N16-s111",
						Level39PString1:                   utils.PointerString("L39-N16-p101"),
						Level39PString2:                   utils.PointerString("L39-N16-p111"),
						Level39RString1:                   []string{"L39-N16-r101", "L39-N16-r102", "L39-N16-r103"},
						Level39RString2:                   []string{"L39-N16-r111", "L39-N16-r112", "L39-N16-r113"},
						Level39MString1:                   map[string]string{"L39-N16-k101": "L39-N16-v101", "L39-N16-k102": "L39-N16-v102", "L39-N16-k103": "L39-N16-v103"},
						Level39MString2:                   map[string]string{"L39-N16-k111": "L39-N16-v111", "L39-N16-k112": "L39-N16-v112", "L39-N16-k113": "L39-N16-v113"},
						Level39FMessageExtern1:            &pbexternal.Message1{FString1: "L39-N16-e101", FString2: "L39-N16-e102", FString3: "L39-N16-e103"},
						Level39FMessageExtern2:            &pbexternal.Message1{FString1: "L39-N16-e111", FString2: "L39-N16-e112", FString3: "L39-N16-e113"},
						Level39OneOfExtern1:               &pbinline.MessageLevel39_Level39One1String1{Level39One1String1: "L39-N16-es101"},
						Level39OneOfInline1:               &pbinline.MessageLevel39_Level39One2String1{Level39One2String1: "L39-N16-is101"},
						Level39OneOfOmitempty1:            &pbinline.MessageLevel39_Level39One3String1{Level39One3String1: "L39-N16-o101"},
						Level39FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level39FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L39-N16-i101"},
						Level39FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L39-N16-i111", IgnoreFieldsString2: "L39-N16-i112"},
						Level39OneOfIgnoreSelf1:           &pbinline.MessageLevel39_Level39One4String1{Level39One4String1: "L39-N16-i121"},
						Level39OneOfIgnoreParts1:          &pbinline.MessageLevel39_Level39One5String1{Level39One5String1: "L39-N16-i131"},
					}
					level38N32 = &pbinline.MessageLevel38{
						Level38FString1:                   "L38-N32-s101",
						Level38FString2:                   "L38-N32-s111",
						Level38PString1:                   utils.PointerString("L38-N32-p101"),
						Level38PString2:                   utils.PointerString("L38-N32-p111"),
						Level38RString1:                   []string{"L38-N32-r101", "L38-N32-r102", "L38-N32-r103"},
						Level38RString2:                   []string{"L38-N32-r111", "L38-N32-r112", "L38-N32-r113"},
						Level38MString1:                   map[string]string{"L38-N32-k101": "L38-N32-v101", "L38-N32-k102": "L38-N32-v102", "L38-N32-k103": "L38-N32-v103"},
						Level38MString2:                   map[string]string{"L38-N32-k111": "L38-N32-v111", "L38-N32-k112": "L38-N32-v112", "L38-N32-k113": "L38-N32-v113"},
						Level38FMessageExtern1:            &pbexternal.Message1{FString1: "L38-N32-e101", FString2: "L38-N32-e102", FString3: "L38-N32-e103"},
						Level38FMessageExtern2:            &pbexternal.Message1{FString1: "L38-N32-e111", FString2: "L38-N32-e112", FString3: "L38-N32-e113"},
						Level38OneOfExtern1:               &pbinline.MessageLevel38_Level38One1String1{Level38One1String1: "L38-N32-es101"},
						Level38OneOfInline1:               &pbinline.MessageLevel38_Level38One2String1{Level38One2String1: "L38-N32-is101"},
						Level38OneOfOmitempty1:            &pbinline.MessageLevel38_Level38One3String1{Level38One3String1: "L38-N32-o101"},
						Level38FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level38FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L38-N32-i101"},
						Level38FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L38-N32-i111", IgnoreFieldsString2: "L38-N32-i112"},
						Level38OneOfIgnoreSelf1:           &pbinline.MessageLevel38_Level38One4String1{Level38One4String1: "L38-N32-i121"},
						Level38OneOfIgnoreParts1:          &pbinline.MessageLevel38_Level38One5String1{Level38One5String1: "L38-N32-i131"},
					}
					level40N16 = &pbinline.MessageLevel40{
						Level40FString1:                   "L40-N16-s101",
						Level40FString2:                   "L40-N16-s111",
						Level40PString1:                   utils.PointerString("L40-N16-p101"),
						Level40PString2:                   utils.PointerString("L40-N16-p111"),
						Level40RString1:                   []string{"L40-N16-r101", "L40-N16-r102", "L40-N16-r103"},
						Level40RString2:                   []string{"L40-N16-r111", "L40-N16-r112", "L40-N16-r113"},
						Level40MString1:                   map[string]string{"L40-N16-k101": "L40-N16-v101", "L40-N16-k102": "L40-N16-v102", "L40-N16-k103": "L40-N16-v103"},
						Level40MString2:                   map[string]string{"L40-N16-k111": "L40-N16-v111", "L40-N16-k112": "L40-N16-v112", "L40-N16-k113": "L40-N16-v113"},
						Level40FMessageExtern1:            &pbexternal.Message1{FString1: "L40-N16-e101", FString2: "L40-N16-e102", FString3: "L40-N16-e103"},
						Level40FMessageExtern2:            &pbexternal.Message1{FString1: "L40-N16-e111", FString2: "L40-N16-e112", FString3: "L40-N16-e113"},
						Level40OneOfExtern1:               &pbinline.MessageLevel40_Level40One1String1{Level40One1String1: "L40-N16-es101"},
						Level40OneOfInline1:               &pbinline.MessageLevel40_Level40One2String1{Level40One2String1: "L40-N16-is101"},
						Level40OneOfOmitempty1:            &pbinline.MessageLevel40_Level40One3String1{Level40One3String1: "L40-N16-o101"},
						Level40FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level40FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L40-N16-i101"},
						Level40FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L40-N16-i111", IgnoreFieldsString2: "L40-N16-i112"},
						Level40OneOfIgnoreSelf1:           &pbinline.MessageLevel40_Level40One4String1{Level40One4String1: "L40-N16-i121"},
						Level40OneOfIgnoreParts1:          &pbinline.MessageLevel40_Level40One5String1{Level40One5String1: "L40-N16-i131"},
					}
				}
				level22N16 = &pbinline.MessageLevel22{
					Level22FString1:                   "L22-N16-s101",
					Level22FString2:                   "L22-N16-s111",
					Level22PString1:                   utils.PointerString("L22-N16-p101"),
					Level22PString2:                   utils.PointerString("L22-N16-p111"),
					Level22RString1:                   []string{"L22-N16-r101", "L22-N16-r102", "L22-N16-r103"},
					Level22RString2:                   []string{"L22-N16-r111", "L22-N16-r112", "L22-N16-r113"},
					Level22MString1:                   map[string]string{"L22-N16-k101": "L22-N16-v101", "L22-N16-k102": "L22-N16-v102", "L22-N16-k103": "L22-N16-v103"},
					Level22MString2:                   map[string]string{"L22-N16-k111": "L22-N16-v111", "L22-N16-k112": "L22-N16-v112", "L22-N16-k113": "L22-N16-v113"},
					Level22FMessageExtern1:            &pbexternal.Message1{FString1: "L22-N16-e101", FString2: "L22-N16-e102", FString3: "L22-N16-e103"},
					Level22FMessageExtern2:            &pbexternal.Message1{FString1: "L22-N16-e111", FString2: "L22-N16-e112", FString3: "L22-N16-e113"},
					Level22FMessageInline38:           level38N31,
					Level22FMessageInline39:           level39N16,
					Level22OneOfExtern1:               &pbinline.MessageLevel22_Level22One1MessageInline38{Level22One1MessageInline38: level38N32},
					Level22OneOfInline1:               &pbinline.MessageLevel22_Level22One2MessageInline40{Level22One2MessageInline40: level40N16},
					Level22OneOfOmitempty1:            &pbinline.MessageLevel22_Level22One3String1{Level22One3String1: "L22-N16-o101"},
					Level22FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level22FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L22-N16-i101"},
					Level22FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L22-N16-i111", IgnoreFieldsString2: "L22-N16-i112"},
					Level22OneOfIgnoreSelf1:           &pbinline.MessageLevel22_Level22One4String1{Level22One4String1: "L22-N16-i121"},
					Level22OneOfIgnoreParts1:          &pbinline.MessageLevel22_Level22One5String1{Level22One5String1: "L22-N16-i131"},
				}
				level24N08 = &pbinline.MessageLevel24{
					Level24FString1:                   "L24-N08-s101",
					Level24FString2:                   "L24-N08-s111",
					Level24PString1:                   utils.PointerString("L24-N08-p101"),
					Level24PString2:                   utils.PointerString("L24-N08-p111"),
					Level24RString1:                   []string{"L24-N08-r101", "L24-N08-r102", "L24-N08-r103"},
					Level24RString2:                   []string{"L24-N08-r111", "L24-N08-r112", "L24-N08-r113"},
					Level24MString1:                   map[string]string{"L24-N08-k101": "L24-N08-v101", "L24-N08-k102": "L24-N08-v102", "L24-N08-k103": "L24-N08-v103"},
					Level24MString2:                   map[string]string{"L24-N08-k111": "L24-N08-v111", "L24-N08-k112": "L24-N08-v112", "L24-N08-k113": "L24-N08-v113"},
					Level24FMessageExtern1:            &pbexternal.Message1{FString1: "L24-N08-e101", FString2: "L24-N08-e102", FString3: "L24-N08-e103"},
					Level24FMessageExtern2:            &pbexternal.Message1{FString1: "L24-N08-e111", FString2: "L24-N08-e112", FString3: "L24-N08-e113"},
					Level24OneOfExtern1:               &pbinline.MessageLevel24_Level24One1String1{Level24One1String1: "L24-N08-es101"},
					Level24OneOfInline1:               &pbinline.MessageLevel24_Level24One2String1{Level24One2String1: "L24-N08-is101"},
					Level24OneOfOmitempty1:            &pbinline.MessageLevel24_Level24One3String1{Level24One3String1: "L24-N08-o101"},
					Level24FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level24FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L24-N08-i101"},
					Level24FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L24-N08-i111", IgnoreFieldsString2: "L24-N08-i112"},
					Level24OneOfIgnoreSelf1:           &pbinline.MessageLevel24_Level24One4String1{Level24One4String1: "L24-N08-i121"},
					Level24OneOfIgnoreParts1:          &pbinline.MessageLevel24_Level24One5String1{Level24One5String1: "L24-N08-i131"},
				}
			}
			level14N08 = &pbinline.MessageLevel14{
				Level14FString1:                   "L14-N08-s101",
				Level14FString2:                   "L14-N08-s111",
				Level14PString1:                   utils.PointerString("L14-N08-p101"),
				Level14PString2:                   utils.PointerString("L14-N08-p111"),
				Level14RString1:                   []string{"L14-N08-r101", "L14-N08-r102", "L14-N08-r103"},
				Level14RString2:                   []string{"L14-N08-r111", "L14-N08-r112", "L14-N08-r113"},
				Level14MString1:                   map[string]string{"L14-N08-k101": "L14-N08-v101", "L14-N08-k102": "L14-N08-v102", "L14-N08-k103": "L14-N08-v103"},
				Level14MString2:                   map[string]string{"L14-N08-k111": "L14-N08-v111", "L14-N08-k112": "L14-N08-v112", "L14-N08-k113": "L14-N08-v113"},
				Level14FMessageExtern1:            &pbexternal.Message1{FString1: "L14-N08-e101", FString2: "L14-N08-e102", FString3: "L14-N08-e103"},
				Level14FMessageExtern2:            &pbexternal.Message1{FString1: "L14-N08-e111", FString2: "L14-N08-e112", FString3: "L14-N08-e113"},
				Level14FMessageInline22:           level22N15,
				Level14FMessageInline23:           level23N08,
				Level14OneOfExtern1:               &pbinline.MessageLevel14_Level14One1MessageInline22{Level14One1MessageInline22: level22N16},
				Level14OneOfInline1:               &pbinline.MessageLevel14_Level14One2MessageInline24{Level14One2MessageInline24: level24N08},
				Level14OneOfOmitempty1:            &pbinline.MessageLevel14_Level14One3String1{Level14One3String1: "L14-N08-o101"},
				Level14FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level14FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L14-N08-i101"},
				Level14FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L14-N08-i111", IgnoreFieldsString2: "L14-N08-i112"},
				Level14OneOfIgnoreSelf1:           &pbinline.MessageLevel14_Level14One4String1{Level14One4String1: "L14-N08-i121"},
				Level14OneOfIgnoreParts1:          &pbinline.MessageLevel14_Level14One5String1{Level14One5String1: "L14-N08-i131"},
			}
			{ // message for level16N04
				{ // message for level26N07
					level30N13 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N13-s101",
						Level30FString2:                   "L30-N13-s111",
						Level30PString1:                   utils.PointerString("L30-N13-p101"),
						Level30PString2:                   utils.PointerString("L30-N13-p111"),
						Level30RString1:                   []string{"L30-N13-r101", "L30-N13-r102", "L30-N13-r103"},
						Level30RString2:                   []string{"L30-N13-r111", "L30-N13-r112", "L30-N13-r113"},
						Level30MString1:                   map[string]string{"L30-N13-k101": "L30-N13-v101", "L30-N13-k102": "L30-N13-v102", "L30-N13-k103": "L30-N13-v103"},
						Level30MString2:                   map[string]string{"L30-N13-k111": "L30-N13-v111", "L30-N13-k112": "L30-N13-v112", "L30-N13-k113": "L30-N13-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N13-e101", FString2: "L30-N13-e102", FString3: "L30-N13-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N13-e111", FString2: "L30-N13-e112", FString3: "L30-N13-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N13-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N13-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N13-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N13-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N13-i111", IgnoreFieldsString2: "L30-N13-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N13-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N13-i131"},
					}
					level31N07 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N07-s101",
						Level31FString2:                   "L31-N07-s111",
						Level31PString1:                   utils.PointerString("L31-N07-p101"),
						Level31PString2:                   utils.PointerString("L31-N07-p111"),
						Level31RString1:                   []string{"L31-N07-r101", "L31-N07-r102", "L31-N07-r103"},
						Level31RString2:                   []string{"L31-N07-r111", "L31-N07-r112", "L31-N07-r113"},
						Level31MString1:                   map[string]string{"L31-N07-k101": "L31-N07-v101", "L31-N07-k102": "L31-N07-v102", "L31-N07-k103": "L31-N07-v103"},
						Level31MString2:                   map[string]string{"L31-N07-k111": "L31-N07-v111", "L31-N07-k112": "L31-N07-v112", "L31-N07-k113": "L31-N07-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N07-e101", FString2: "L31-N07-e102", FString3: "L31-N07-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N07-e111", FString2: "L31-N07-e112", FString3: "L31-N07-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N07-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N07-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N07-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N07-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N07-i111", IgnoreFieldsString2: "L31-N07-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N07-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N07-i131"},
					}
					level30N14 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N14-s101",
						Level30FString2:                   "L30-N14-s111",
						Level30PString1:                   utils.PointerString("L30-N14-p101"),
						Level30PString2:                   utils.PointerString("L30-N14-p111"),
						Level30RString1:                   []string{"L30-N14-r101", "L30-N14-r102", "L30-N14-r103"},
						Level30RString2:                   []string{"L30-N14-r111", "L30-N14-r112", "L30-N14-r113"},
						Level30MString1:                   map[string]string{"L30-N14-k101": "L30-N14-v101", "L30-N14-k102": "L30-N14-v102", "L30-N14-k103": "L30-N14-v103"},
						Level30MString2:                   map[string]string{"L30-N14-k111": "L30-N14-v111", "L30-N14-k112": "L30-N14-v112", "L30-N14-k113": "L30-N14-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N14-e101", FString2: "L30-N14-e102", FString3: "L30-N14-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N14-e111", FString2: "L30-N14-e112", FString3: "L30-N14-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N14-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N14-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N14-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N14-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N14-i111", IgnoreFieldsString2: "L30-N14-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N14-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N14-i131"},
					}
					level32N07 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N07-s101",
						Level32FString2:                   "L32-N07-s111",
						Level32PString1:                   utils.PointerString("L32-N07-p101"),
						Level32PString2:                   utils.PointerString("L32-N07-p111"),
						Level32RString1:                   []string{"L32-N07-r101", "L32-N07-r102", "L32-N07-r103"},
						Level32RString2:                   []string{"L32-N07-r111", "L32-N07-r112", "L32-N07-r113"},
						Level32MString1:                   map[string]string{"L32-N07-k101": "L32-N07-v101", "L32-N07-k102": "L32-N07-v102", "L32-N07-k103": "L32-N07-v103"},
						Level32MString2:                   map[string]string{"L32-N07-k111": "L32-N07-v111", "L32-N07-k112": "L32-N07-v112", "L32-N07-k113": "L32-N07-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N07-e101", FString2: "L32-N07-e102", FString3: "L32-N07-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N07-e111", FString2: "L32-N07-e112", FString3: "L32-N07-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N07-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N07-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N07-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N07-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N07-i111", IgnoreFieldsString2: "L32-N07-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N07-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N07-i131"},
					}
				}
				level26N07 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N07-s101",
					Level26FString2:                   "L26-N07-s111",
					Level26PString1:                   utils.PointerString("L26-N07-p101"),
					Level26PString2:                   utils.PointerString("L26-N07-p111"),
					Level26RString1:                   []string{"L26-N07-r101", "L26-N07-r102", "L26-N07-r103"},
					Level26RString2:                   []string{"L26-N07-r111", "L26-N07-r112", "L26-N07-r113"},
					Level26MString1:                   map[string]string{"L26-N07-k101": "L26-N07-v101", "L26-N07-k102": "L26-N07-v102", "L26-N07-k103": "L26-N07-v103"},
					Level26MString2:                   map[string]string{"L26-N07-k111": "L26-N07-v111", "L26-N07-k112": "L26-N07-v112", "L26-N07-k113": "L26-N07-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N07-e101", FString2: "L26-N07-e102", FString3: "L26-N07-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N07-e111", FString2: "L26-N07-e112", FString3: "L26-N07-e113"},
					Level26FMessageInline30:           level30N13,
					Level26FMessageInline31:           level31N07,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N14},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N07},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N07-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N07-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N07-i111", IgnoreFieldsString2: "L26-N07-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N07-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N07-i131"},
				}
				level27N04 = &pbinline.MessageLevel27{
					Level27FString1:                   "L27-N04-s101",
					Level27FString2:                   "L27-N04-s111",
					Level27PString1:                   utils.PointerString("L27-N04-p101"),
					Level27PString2:                   utils.PointerString("L27-N04-p111"),
					Level27RString1:                   []string{"L27-N04-r101", "L27-N04-r102", "L27-N04-r103"},
					Level27RString2:                   []string{"L27-N04-r111", "L27-N04-r112", "L27-N04-r113"},
					Level27MString1:                   map[string]string{"L27-N04-k101": "L27-N04-v101", "L27-N04-k102": "L27-N04-v102", "L27-N04-k103": "L27-N04-v103"},
					Level27MString2:                   map[string]string{"L27-N04-k111": "L27-N04-v111", "L27-N04-k112": "L27-N04-v112", "L27-N04-k113": "L27-N04-v113"},
					Level27FMessageExtern1:            &pbexternal.Message1{FString1: "L27-N04-e101", FString2: "L27-N04-e102", FString3: "L27-N04-e103"},
					Level27FMessageExtern2:            &pbexternal.Message1{FString1: "L27-N04-e111", FString2: "L27-N04-e112", FString3: "L27-N04-e113"},
					Level27OneOfExtern1:               &pbinline.MessageLevel27_Level27One1String1{Level27One1String1: "L27-N04-es101"},
					Level27OneOfInline1:               &pbinline.MessageLevel27_Level27One2String1{Level27One2String1: "L27-N04-is101"},
					Level27OneOfOmitempty1:            &pbinline.MessageLevel27_Level27One3String1{Level27One3String1: "L27-N04-o101"},
					Level27FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level27FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L27-N04-i101"},
					Level27FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L27-N04-i111", IgnoreFieldsString2: "L27-N04-i112"},
					Level27OneOfIgnoreSelf1:           &pbinline.MessageLevel27_Level27One4String1{Level27One4String1: "L27-N04-i121"},
					Level27OneOfIgnoreParts1:          &pbinline.MessageLevel27_Level27One5String1{Level27One5String1: "L27-N04-i131"},
				}
				{ // message for level26N08
					level30N15 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N15-s101",
						Level30FString2:                   "L30-N15-s111",
						Level30PString1:                   utils.PointerString("L30-N15-p101"),
						Level30PString2:                   utils.PointerString("L30-N15-p111"),
						Level30RString1:                   []string{"L30-N15-r101", "L30-N15-r102", "L30-N15-r103"},
						Level30RString2:                   []string{"L30-N15-r111", "L30-N15-r112", "L30-N15-r113"},
						Level30MString1:                   map[string]string{"L30-N15-k101": "L30-N15-v101", "L30-N15-k102": "L30-N15-v102", "L30-N15-k103": "L30-N15-v103"},
						Level30MString2:                   map[string]string{"L30-N15-k111": "L30-N15-v111", "L30-N15-k112": "L30-N15-v112", "L30-N15-k113": "L30-N15-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N15-e101", FString2: "L30-N15-e102", FString3: "L30-N15-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N15-e111", FString2: "L30-N15-e112", FString3: "L30-N15-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N15-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N15-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N15-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N15-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N15-i111", IgnoreFieldsString2: "L30-N15-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N15-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N15-i131"},
					}
					level31N08 = &pbinline.MessageLevel31{
						Level31FString1:                   "L31-N08-s101",
						Level31FString2:                   "L31-N08-s111",
						Level31PString1:                   utils.PointerString("L31-N08-p101"),
						Level31PString2:                   utils.PointerString("L31-N08-p111"),
						Level31RString1:                   []string{"L31-N08-r101", "L31-N08-r102", "L31-N08-r103"},
						Level31RString2:                   []string{"L31-N08-r111", "L31-N08-r112", "L31-N08-r113"},
						Level31MString1:                   map[string]string{"L31-N08-k101": "L31-N08-v101", "L31-N08-k102": "L31-N08-v102", "L31-N08-k103": "L31-N08-v103"},
						Level31MString2:                   map[string]string{"L31-N08-k111": "L31-N08-v111", "L31-N08-k112": "L31-N08-v112", "L31-N08-k113": "L31-N08-v113"},
						Level31FMessageExtern1:            &pbexternal.Message1{FString1: "L31-N08-e101", FString2: "L31-N08-e102", FString3: "L31-N08-e103"},
						Level31FMessageExtern2:            &pbexternal.Message1{FString1: "L31-N08-e111", FString2: "L31-N08-e112", FString3: "L31-N08-e113"},
						Level31OneOfExtern1:               &pbinline.MessageLevel31_Level31One1String1{Level31One1String1: "L31-N08-es101"},
						Level31OneOfInline1:               &pbinline.MessageLevel31_Level31One2String1{Level31One2String1: "L31-N08-is101"},
						Level31OneOfOmitempty1:            &pbinline.MessageLevel31_Level31One3String1{Level31One3String1: "L31-N08-o101"},
						Level31FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level31FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L31-N08-i101"},
						Level31FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L31-N08-i111", IgnoreFieldsString2: "L31-N08-i112"},
						Level31OneOfIgnoreSelf1:           &pbinline.MessageLevel31_Level31One4String1{Level31One4String1: "L31-N08-i121"},
						Level31OneOfIgnoreParts1:          &pbinline.MessageLevel31_Level31One5String1{Level31One5String1: "L31-N08-i131"},
					}
					level30N16 = &pbinline.MessageLevel30{
						Level30FString1:                   "L30-N16-s101",
						Level30FString2:                   "L30-N16-s111",
						Level30PString1:                   utils.PointerString("L30-N16-p101"),
						Level30PString2:                   utils.PointerString("L30-N16-p111"),
						Level30RString1:                   []string{"L30-N16-r101", "L30-N16-r102", "L30-N16-r103"},
						Level30RString2:                   []string{"L30-N16-r111", "L30-N16-r112", "L30-N16-r113"},
						Level30MString1:                   map[string]string{"L30-N16-k101": "L30-N16-v101", "L30-N16-k102": "L30-N16-v102", "L30-N16-k103": "L30-N16-v103"},
						Level30MString2:                   map[string]string{"L30-N16-k111": "L30-N16-v111", "L30-N16-k112": "L30-N16-v112", "L30-N16-k113": "L30-N16-v113"},
						Level30FMessageExtern1:            &pbexternal.Message1{FString1: "L30-N16-e101", FString2: "L30-N16-e102", FString3: "L30-N16-e103"},
						Level30FMessageExtern2:            &pbexternal.Message1{FString1: "L30-N16-e111", FString2: "L30-N16-e112", FString3: "L30-N16-e113"},
						Level30OneOfExtern1:               &pbinline.MessageLevel30_Level30One1String1{Level30One1String1: "L30-N16-es101"},
						Level30OneOfInline1:               &pbinline.MessageLevel30_Level30One2String1{Level30One2String1: "L30-N16-is101"},
						Level30OneOfOmitempty1:            &pbinline.MessageLevel30_Level30One3String1{Level30One3String1: "L30-N16-o101"},
						Level30FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level30FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L30-N16-i101"},
						Level30FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L30-N16-i111", IgnoreFieldsString2: "L30-N16-i112"},
						Level30OneOfIgnoreSelf1:           &pbinline.MessageLevel30_Level30One4String1{Level30One4String1: "L30-N16-i121"},
						Level30OneOfIgnoreParts1:          &pbinline.MessageLevel30_Level30One5String1{Level30One5String1: "L30-N16-i131"},
					}
					level32N08 = &pbinline.MessageLevel32{
						Level32FString1:                   "L32-N08-s101",
						Level32FString2:                   "L32-N08-s111",
						Level32PString1:                   utils.PointerString("L32-N08-p101"),
						Level32PString2:                   utils.PointerString("L32-N08-p111"),
						Level32RString1:                   []string{"L32-N08-r101", "L32-N08-r102", "L32-N08-r103"},
						Level32RString2:                   []string{"L32-N08-r111", "L32-N08-r112", "L32-N08-r113"},
						Level32MString1:                   map[string]string{"L32-N08-k101": "L32-N08-v101", "L32-N08-k102": "L32-N08-v102", "L32-N08-k103": "L32-N08-v103"},
						Level32MString2:                   map[string]string{"L32-N08-k111": "L32-N08-v111", "L32-N08-k112": "L32-N08-v112", "L32-N08-k113": "L32-N08-v113"},
						Level32FMessageExtern1:            &pbexternal.Message1{FString1: "L32-N08-e101", FString2: "L32-N08-e102", FString3: "L32-N08-e103"},
						Level32FMessageExtern2:            &pbexternal.Message1{FString1: "L32-N08-e111", FString2: "L32-N08-e112", FString3: "L32-N08-e113"},
						Level32OneOfExtern1:               &pbinline.MessageLevel32_Level32One1String1{Level32One1String1: "L32-N08-es101"},
						Level32OneOfInline1:               &pbinline.MessageLevel32_Level32One2String1{Level32One2String1: "L32-N08-is101"},
						Level32OneOfOmitempty1:            &pbinline.MessageLevel32_Level32One3String1{Level32One3String1: "L32-N08-o101"},
						Level32FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level32FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L32-N08-i101"},
						Level32FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L32-N08-i111", IgnoreFieldsString2: "L32-N08-i112"},
						Level32OneOfIgnoreSelf1:           &pbinline.MessageLevel32_Level32One4String1{Level32One4String1: "L32-N08-i121"},
						Level32OneOfIgnoreParts1:          &pbinline.MessageLevel32_Level32One5String1{Level32One5String1: "L32-N08-i131"},
					}
				}
				level26N08 = &pbinline.MessageLevel26{
					Level26FString1:                   "L26-N08-s101",
					Level26FString2:                   "L26-N08-s111",
					Level26PString1:                   utils.PointerString("L26-N08-p101"),
					Level26PString2:                   utils.PointerString("L26-N08-p111"),
					Level26RString1:                   []string{"L26-N08-r101", "L26-N08-r102", "L26-N08-r103"},
					Level26RString2:                   []string{"L26-N08-r111", "L26-N08-r112", "L26-N08-r113"},
					Level26MString1:                   map[string]string{"L26-N08-k101": "L26-N08-v101", "L26-N08-k102": "L26-N08-v102", "L26-N08-k103": "L26-N08-v103"},
					Level26MString2:                   map[string]string{"L26-N08-k111": "L26-N08-v111", "L26-N08-k112": "L26-N08-v112", "L26-N08-k113": "L26-N08-v113"},
					Level26FMessageExtern1:            &pbexternal.Message1{FString1: "L26-N08-e101", FString2: "L26-N08-e102", FString3: "L26-N08-e103"},
					Level26FMessageExtern2:            &pbexternal.Message1{FString1: "L26-N08-e111", FString2: "L26-N08-e112", FString3: "L26-N08-e113"},
					Level26FMessageInline30:           level30N15,
					Level26FMessageInline31:           level31N08,
					Level26OneOfExtern1:               &pbinline.MessageLevel26_Level26One1MessageInline30{Level26One1MessageInline30: level30N16},
					Level26OneOfInline1:               &pbinline.MessageLevel26_Level26One2MessageInline32{Level26One2MessageInline32: level32N08},
					Level26OneOfOmitempty1:            &pbinline.MessageLevel26_Level26One3String1{Level26One3String1: "L26-N08-o101"},
					Level26FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level26FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L26-N08-i101"},
					Level26FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L26-N08-i111", IgnoreFieldsString2: "L26-N08-i112"},
					Level26OneOfIgnoreSelf1:           &pbinline.MessageLevel26_Level26One4String1{Level26One4String1: "L26-N08-i121"},
					Level26OneOfIgnoreParts1:          &pbinline.MessageLevel26_Level26One5String1{Level26One5String1: "L26-N08-i131"},
				}
				{ // message for level28N04
					level34N07 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N07-s101",
						Level34FString2:                   "L34-N07-s111",
						Level34PString1:                   utils.PointerString("L34-N07-p101"),
						Level34PString2:                   utils.PointerString("L34-N07-p111"),
						Level34RString1:                   []string{"L34-N07-r101", "L34-N07-r102", "L34-N07-r103"},
						Level34RString2:                   []string{"L34-N07-r111", "L34-N07-r112", "L34-N07-r113"},
						Level34MString1:                   map[string]string{"L34-N07-k101": "L34-N07-v101", "L34-N07-k102": "L34-N07-v102", "L34-N07-k103": "L34-N07-v103"},
						Level34MString2:                   map[string]string{"L34-N07-k111": "L34-N07-v111", "L34-N07-k112": "L34-N07-v112", "L34-N07-k113": "L34-N07-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N07-e101", FString2: "L34-N07-e102", FString3: "L34-N07-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N07-e111", FString2: "L34-N07-e112", FString3: "L34-N07-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N07-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N07-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N07-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N07-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N07-i111", IgnoreFieldsString2: "L34-N07-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N07-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N07-i131"},
					}
					level35N04 = &pbinline.MessageLevel35{
						Level35FString1:                   "L35-N04-s101",
						Level35FString2:                   "L35-N04-s111",
						Level35PString1:                   utils.PointerString("L35-N04-p101"),
						Level35PString2:                   utils.PointerString("L35-N04-p111"),
						Level35RString1:                   []string{"L35-N04-r101", "L35-N04-r102", "L35-N04-r103"},
						Level35RString2:                   []string{"L35-N04-r111", "L35-N04-r112", "L35-N04-r113"},
						Level35MString1:                   map[string]string{"L35-N04-k101": "L35-N04-v101", "L35-N04-k102": "L35-N04-v102", "L35-N04-k103": "L35-N04-v103"},
						Level35MString2:                   map[string]string{"L35-N04-k111": "L35-N04-v111", "L35-N04-k112": "L35-N04-v112", "L35-N04-k113": "L35-N04-v113"},
						Level35FMessageExtern1:            &pbexternal.Message1{FString1: "L35-N04-e101", FString2: "L35-N04-e102", FString3: "L35-N04-e103"},
						Level35FMessageExtern2:            &pbexternal.Message1{FString1: "L35-N04-e111", FString2: "L35-N04-e112", FString3: "L35-N04-e113"},
						Level35OneOfExtern1:               &pbinline.MessageLevel35_Level35One1String1{Level35One1String1: "L35-N04-es101"},
						Level35OneOfInline1:               &pbinline.MessageLevel35_Level35One2String1{Level35One2String1: "L35-N04-is101"},
						Level35OneOfOmitempty1:            &pbinline.MessageLevel35_Level35One3String1{Level35One3String1: "L35-N04-o101"},
						Level35FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level35FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L35-N04-i101"},
						Level35FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L35-N04-i111", IgnoreFieldsString2: "L35-N04-i112"},
						Level35OneOfIgnoreSelf1:           &pbinline.MessageLevel35_Level35One4String1{Level35One4String1: "L35-N04-i121"},
						Level35OneOfIgnoreParts1:          &pbinline.MessageLevel35_Level35One5String1{Level35One5String1: "L35-N04-i131"},
					}
					level34N08 = &pbinline.MessageLevel34{
						Level34FString1:                   "L34-N08-s101",
						Level34FString2:                   "L34-N08-s111",
						Level34PString1:                   utils.PointerString("L34-N08-p101"),
						Level34PString2:                   utils.PointerString("L34-N08-p111"),
						Level34RString1:                   []string{"L34-N08-r101", "L34-N08-r102", "L34-N08-r103"},
						Level34RString2:                   []string{"L34-N08-r111", "L34-N08-r112", "L34-N08-r113"},
						Level34MString1:                   map[string]string{"L34-N08-k101": "L34-N08-v101", "L34-N08-k102": "L34-N08-v102", "L34-N08-k103": "L34-N08-v103"},
						Level34MString2:                   map[string]string{"L34-N08-k111": "L34-N08-v111", "L34-N08-k112": "L34-N08-v112", "L34-N08-k113": "L34-N08-v113"},
						Level34FMessageExtern1:            &pbexternal.Message1{FString1: "L34-N08-e101", FString2: "L34-N08-e102", FString3: "L34-N08-e103"},
						Level34FMessageExtern2:            &pbexternal.Message1{FString1: "L34-N08-e111", FString2: "L34-N08-e112", FString3: "L34-N08-e113"},
						Level34OneOfExtern1:               &pbinline.MessageLevel34_Level34One1String1{Level34One1String1: "L34-N08-es101"},
						Level34OneOfInline1:               &pbinline.MessageLevel34_Level34One2String1{Level34One2String1: "L34-N08-is101"},
						Level34OneOfOmitempty1:            &pbinline.MessageLevel34_Level34One3String1{Level34One3String1: "L34-N08-o101"},
						Level34FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level34FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L34-N08-i101"},
						Level34FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L34-N08-i111", IgnoreFieldsString2: "L34-N08-i112"},
						Level34OneOfIgnoreSelf1:           &pbinline.MessageLevel34_Level34One4String1{Level34One4String1: "L34-N08-i121"},
						Level34OneOfIgnoreParts1:          &pbinline.MessageLevel34_Level34One5String1{Level34One5String1: "L34-N08-i131"},
					}
					level36N04 = &pbinline.MessageLevel36{
						Level36FString1:                   "L36-N04-s101",
						Level36FString2:                   "L36-N04-s111",
						Level36PString1:                   utils.PointerString("L36-N04-p101"),
						Level36PString2:                   utils.PointerString("L36-N04-p111"),
						Level36RString1:                   []string{"L36-N04-r101", "L36-N04-r102", "L36-N04-r103"},
						Level36RString2:                   []string{"L36-N04-r111", "L36-N04-r112", "L36-N04-r113"},
						Level36MString1:                   map[string]string{"L36-N04-k101": "L36-N04-v101", "L36-N04-k102": "L36-N04-v102", "L36-N04-k103": "L36-N04-v103"},
						Level36MString2:                   map[string]string{"L36-N04-k111": "L36-N04-v111", "L36-N04-k112": "L36-N04-v112", "L36-N04-k113": "L36-N04-v113"},
						Level36FMessageExtern1:            &pbexternal.Message1{FString1: "L36-N04-e101", FString2: "L36-N04-e102", FString3: "L36-N04-e103"},
						Level36FMessageExtern2:            &pbexternal.Message1{FString1: "L36-N04-e111", FString2: "L36-N04-e112", FString3: "L36-N04-e113"},
						Level36OneOfExtern1:               &pbinline.MessageLevel36_Level36One1String1{Level36One1String1: "L36-N04-es101"},
						Level36OneOfInline1:               &pbinline.MessageLevel36_Level36One2String1{Level36One2String1: "L36-N04-is101"},
						Level36OneOfOmitempty1:            &pbinline.MessageLevel36_Level36One3String1{Level36One3String1: "L36-N04-o101"},
						Level36FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
						Level36FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L36-N04-i101"},
						Level36FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L36-N04-i111", IgnoreFieldsString2: "L36-N04-i112"},
						Level36OneOfIgnoreSelf1:           &pbinline.MessageLevel36_Level36One4String1{Level36One4String1: "L36-N04-i121"},
						Level36OneOfIgnoreParts1:          &pbinline.MessageLevel36_Level36One5String1{Level36One5String1: "L36-N04-i131"},
					}
				}
				level28N04 = &pbinline.MessageLevel28{
					Level28FString1:                   "L28-N04-s101",
					Level28FString2:                   "L28-N04-s111",
					Level28PString1:                   utils.PointerString("L28-N04-p101"),
					Level28PString2:                   utils.PointerString("L28-N04-p111"),
					Level28RString1:                   []string{"L28-N04-r101", "L28-N04-r102", "L28-N04-r103"},
					Level28RString2:                   []string{"L28-N04-r111", "L28-N04-r112", "L28-N04-r113"},
					Level28MString1:                   map[string]string{"L28-N04-k101": "L28-N04-v101", "L28-N04-k102": "L28-N04-v102", "L28-N04-k103": "L28-N04-v103"},
					Level28MString2:                   map[string]string{"L28-N04-k111": "L28-N04-v111", "L28-N04-k112": "L28-N04-v112", "L28-N04-k113": "L28-N04-v113"},
					Level28FMessageExtern1:            &pbexternal.Message1{FString1: "L28-N04-e101", FString2: "L28-N04-e102", FString3: "L28-N04-e103"},
					Level28FMessageExtern2:            &pbexternal.Message1{FString1: "L28-N04-e111", FString2: "L28-N04-e112", FString3: "L28-N04-e113"},
					Level28FMessageInline34:           level34N07,
					Level28FMessageInline35:           level35N04,
					Level28OneOfExtern1:               &pbinline.MessageLevel28_Level28One1MessageInline34{Level28One1MessageInline34: level34N08},
					Level28OneOfInline1:               &pbinline.MessageLevel28_Level28One2MessageInline36{Level28One2MessageInline36: level36N04},
					Level28OneOfOmitempty1:            &pbinline.MessageLevel28_Level28One3String1{Level28One3String1: "L28-N04-o101"},
					Level28FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
					Level28FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L28-N04-i101"},
					Level28FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L28-N04-i111", IgnoreFieldsString2: "L28-N04-i112"},
					Level28OneOfIgnoreSelf1:           &pbinline.MessageLevel28_Level28One4String1{Level28One4String1: "L28-N04-i121"},
					Level28OneOfIgnoreParts1:          &pbinline.MessageLevel28_Level28One5String1{Level28One5String1: "L28-N04-i131"},
				}
			}
			level16N04 = &pbinline.MessageLevel16{
				Level16FString1:                   "L16-N04-s101",
				Level16FString2:                   "L16-N04-s111",
				Level16PString1:                   utils.PointerString("L16-N04-p101"),
				Level16PString2:                   utils.PointerString("L16-N04-p111"),
				Level16RString1:                   []string{"L16-N04-r101", "L16-N04-r102", "L16-N04-r103"},
				Level16RString2:                   []string{"L16-N04-r111", "L16-N04-r112", "L16-N04-r113"},
				Level16MString1:                   map[string]string{"L16-N04-k101": "L16-N04-v101", "L16-N04-k102": "L16-N04-v102", "L16-N04-k103": "L16-N04-v103"},
				Level16MString2:                   map[string]string{"L16-N04-k111": "L16-N04-v111", "L16-N04-k112": "L16-N04-v112", "L16-N04-k113": "L16-N04-v113"},
				Level16FMessageExtern1:            &pbexternal.Message1{FString1: "L16-N04-e101", FString2: "L16-N04-e102", FString3: "L16-N04-e103"},
				Level16FMessageExtern2:            &pbexternal.Message1{FString1: "L16-N04-e111", FString2: "L16-N04-e112", FString3: "L16-N04-e113"},
				Level16FMessageInline26:           level26N07,
				Level16FMessageInline27:           level27N04,
				Level16OneOfExtern1:               &pbinline.MessageLevel16_Level16One1MessageInline26{Level16One1MessageInline26: level26N08},
				Level16OneOfInline1:               &pbinline.MessageLevel16_Level16One2MessageInline28{Level16One2MessageInline28: level28N04},
				Level16OneOfOmitempty1:            &pbinline.MessageLevel16_Level16One3String1{Level16One3String1: "L16-N04-o101"},
				Level16FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level16FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L16-N04-i101"},
				Level16FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L16-N04-i111", IgnoreFieldsString2: "L16-N04-i112"},
				Level16OneOfIgnoreSelf1:           &pbinline.MessageLevel16_Level16One4String1{Level16One4String1: "L16-N04-i121"},
				Level16OneOfIgnoreParts1:          &pbinline.MessageLevel16_Level16One5String1{Level16One5String1: "L16-N04-i131"},
			}
		}
		level06N04 = &pbinline.MessageLevel06{
			Level06FString1:                   "L06-N04-s101",
			Level06FString2:                   "L06-N04-s111",
			Level06PString1:                   utils.PointerString("L06-N04-p101"),
			Level06PString2:                   utils.PointerString("L06-N04-p111"),
			Level06RString1:                   []string{"L06-N04-r101", "L06-N04-r102", "L06-N04-r103"},
			Level06RString2:                   []string{"L06-N04-r111", "L06-N04-r112", "L06-N04-r113"},
			Level06MString1:                   map[string]string{"L06-N04-k101": "L06-N04-v101", "L06-N04-k102": "L06-N04-v102", "L06-N04-k103": "L06-N04-v103"},
			Level06MString2:                   map[string]string{"L06-N04-k111": "L06-N04-v111", "L06-N04-k112": "L06-N04-v112", "L06-N04-k113": "L06-N04-v113"},
			Level06FMessageExtern1:            &pbexternal.Message1{FString1: "L06-N04-e101", FString2: "L06-N04-e102", FString3: "L06-N04-e103"},
			Level06FMessageExtern2:            &pbexternal.Message1{FString1: "L06-N04-e111", FString2: "L06-N04-e112", FString3: "L06-N04-e113"},
			Level06FMessageInline14:           level14N07,
			Level06FMessageInline15:           level15N04,
			Level06OneOfExtern1:               &pbinline.MessageLevel06_Level06One1MessageInline14{Level06One1MessageInline14: level14N08},
			Level06OneOfInline1:               &pbinline.MessageLevel06_Level4One2MessageInline16{Level4One2MessageInline16: level16N04},
			Level06OneOfOmitempty1:            &pbinline.MessageLevel06_Level06One3String1{Level06One3String1: "L06-N04-o101"},
			Level06FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level06FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L06-N04-i101"},
			Level06FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L06-N04-i111", IgnoreFieldsString2: "L06-N04-i112"},
			Level06OneOfIgnoreSelf1:           &pbinline.MessageLevel06_Level06One4String1{Level06One4String1: "L06-N04-i121"},
			Level06OneOfIgnoreParts1:          &pbinline.MessageLevel06_Level06One5String1{Level06One5String1: "L06-N04-i131"},
		}
		{ // message in level08N02
			level18N03 = &pbinline.MessageLevel18{
				Level18FString1:                   "L18-N03-s101",
				Level18FString2:                   "L18-N03-s111",
				Level18PString1:                   utils.PointerString("L18-N03-p101"),
				Level18PString2:                   utils.PointerString("L18-N03-p111"),
				Level18RString1:                   []string{"L18-N03-r101", "L18-N03-r102", "L18-N03-r103"},
				Level18RString2:                   []string{"L18-N03-r111", "L18-N03-r112", "L18-N03-r113"},
				Level18MString1:                   map[string]string{"L18-N03-k101": "L18-N03-v101", "L18-N03-k102": "L18-N03-v102", "L18-N03-k103": "L18-N03-v103"},
				Level18MString2:                   map[string]string{"L18-N03-k111": "L18-N03-v111", "L18-N03-k112": "L18-N03-v112", "L18-N03-k113": "L18-N03-v113"},
				Level18FMessageExtern1:            &pbexternal.Message1{FString1: "L18-N03-e101", FString2: "L18-N03-e102", FString3: "L18-N03-e103"},
				Level18FMessageExtern2:            &pbexternal.Message1{FString1: "L18-N03-e111", FString2: "L18-N03-e112", FString3: "L18-N03-e113"},
				Level18OneOfExtern1:               &pbinline.MessageLevel18_Level18One1String1{Level18One1String1: "L18-N03-es101"},
				Level18OneOfInline1:               &pbinline.MessageLevel18_Level18One2String1{Level18One2String1: "L18-N03-is101"},
				Level18OneOfOmitempty1:            &pbinline.MessageLevel18_Level18One3String1{Level18One3String1: "L18-N03-o101"},
				Level18FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level18FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L18-N03-i101"},
				Level18FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L18-N03-i111", IgnoreFieldsString2: "L18-N03-i112"},
				Level18OneOfIgnoreSelf1:           &pbinline.MessageLevel18_Level18One4String1{Level18One4String1: "L18-N03-i121"},
				Level18OneOfIgnoreParts1:          &pbinline.MessageLevel18_Level18One5String1{Level18One5String1: "L18-N03-i131"},
			}
			level19N02 = &pbinline.MessageLevel19{
				Level19FString1:                   "L19-N02-s101",
				Level19FString2:                   "L19-N02-s111",
				Level19PString1:                   utils.PointerString("L19-N02-p101"),
				Level19PString2:                   utils.PointerString("L19-N02-p111"),
				Level19RString1:                   []string{"L19-N02-r101", "L19-N02-r102", "L19-N02-r103"},
				Level19RString2:                   []string{"L19-N02-r111", "L19-N02-r112", "L19-N02-r113"},
				Level19MString1:                   map[string]string{"L19-N02-k101": "L19-N02-v101", "L19-N02-k102": "L19-N02-v102", "L19-N02-k103": "L19-N02-v103"},
				Level19MString2:                   map[string]string{"L19-N02-k111": "L19-N02-v111", "L19-N02-k112": "L19-N02-v112", "L19-N02-k113": "L19-N02-v113"},
				Level19FMessageExtern1:            &pbexternal.Message1{FString1: "L19-N02-e101", FString2: "L19-N02-e102", FString3: "L19-N02-e103"},
				Level19FMessageExtern2:            &pbexternal.Message1{FString1: "L19-N02-e111", FString2: "L19-N02-e112", FString3: "L19-N02-e113"},
				Level19OneOfExtern1:               &pbinline.MessageLevel19_Level19One1String1{Level19One1String1: "L19-N02-es101"},
				Level19OneOfInline1:               &pbinline.MessageLevel19_Level19One2String1{Level19One2String1: "L19-N02-is101"},
				Level19OneOfOmitempty1:            &pbinline.MessageLevel19_Level19One3String1{Level19One3String1: "L19-N02-o101"},
				Level19FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level19FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L19-N02-i101"},
				Level19FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L19-N02-i111", IgnoreFieldsString2: "L19-N02-i112"},
				Level19OneOfIgnoreSelf1:           &pbinline.MessageLevel19_Level19One4String1{Level19One4String1: "L19-N02-i121"},
				Level19OneOfIgnoreParts1:          &pbinline.MessageLevel19_Level19One5String1{Level19One5String1: "L19-N02-i131"},
			}
			level18N04 = &pbinline.MessageLevel18{
				Level18FString1:                   "L18-N04-s101",
				Level18FString2:                   "L18-N04-s111",
				Level18PString1:                   utils.PointerString("L18-N04-p101"),
				Level18PString2:                   utils.PointerString("L18-N04-p111"),
				Level18RString1:                   []string{"L18-N04-r101", "L18-N04-r102", "L18-N04-r103"},
				Level18RString2:                   []string{"L18-N04-r111", "L18-N04-r112", "L18-N04-r113"},
				Level18MString1:                   map[string]string{"L18-N04-k101": "L18-N04-v101", "L18-N04-k102": "L18-N04-v102", "L18-N04-k103": "L18-N04-v103"},
				Level18MString2:                   map[string]string{"L18-N04-k111": "L18-N04-v111", "L18-N04-k112": "L18-N04-v112", "L18-N04-k113": "L18-N04-v113"},
				Level18FMessageExtern1:            &pbexternal.Message1{FString1: "L18-N04-e101", FString2: "L18-N04-e102", FString3: "L18-N04-e103"},
				Level18FMessageExtern2:            &pbexternal.Message1{FString1: "L18-N04-e111", FString2: "L18-N04-e112", FString3: "L18-N04-e113"},
				Level18OneOfExtern1:               &pbinline.MessageLevel18_Level18One1String1{Level18One1String1: "L18-N04-es101"},
				Level18OneOfInline1:               &pbinline.MessageLevel18_Level18One2String1{Level18One2String1: "L18-N04-is101"},
				Level18OneOfOmitempty1:            &pbinline.MessageLevel18_Level18One3String1{Level18One3String1: "L18-N04-o101"},
				Level18FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level18FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L18-N04-i101"},
				Level18FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L18-N04-i111", IgnoreFieldsString2: "L18-N04-i112"},
				Level18OneOfIgnoreSelf1:           &pbinline.MessageLevel18_Level18One4String1{Level18One4String1: "L18-N04-i121"},
				Level18OneOfIgnoreParts1:          &pbinline.MessageLevel18_Level18One5String1{Level18One5String1: "L18-N04-i131"},
			}
			level20N02 = &pbinline.MessageLevel20{
				Level20FString1:                   "L20-N02-s101",
				Level20FString2:                   "L20-N02-s111",
				Level20PString1:                   utils.PointerString("L20-N02-p101"),
				Level20PString2:                   utils.PointerString("L20-N02-p111"),
				Level20RString1:                   []string{"L20-N02-r101", "L20-N02-r102", "L20-N02-r103"},
				Level20RString2:                   []string{"L20-N02-r111", "L20-N02-r112", "L20-N02-r113"},
				Level20MString1:                   map[string]string{"L20-N02-k101": "L20-N02-v101", "L20-N02-k102": "L20-N02-v102", "L20-N02-k103": "L20-N02-v103"},
				Level20MString2:                   map[string]string{"L20-N02-k111": "L20-N02-v111", "L20-N02-k112": "L20-N02-v112", "L20-N02-k113": "L20-N02-v113"},
				Level20FMessageExtern1:            &pbexternal.Message1{FString1: "L20-N02-e101", FString2: "L20-N02-e102", FString3: "L20-N02-e103"},
				Level20FMessageExtern2:            &pbexternal.Message1{FString1: "L20-N02-e111", FString2: "L20-N02-e112", FString3: "L20-N02-e113"},
				Level20OneOfExtern1:               &pbinline.MessageLevel20_Level20One1String1{Level20One1String1: "L20-N02-es101"},
				Level20OneOfInline1:               &pbinline.MessageLevel20_Level20One2String1{Level20One2String1: "L20-N02-is101"},
				Level20OneOfOmitempty1:            &pbinline.MessageLevel20_Level20One3String1{Level20One3String1: "L20-N02-o101"},
				Level20FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
				Level20FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L20-N02-i101"},
				Level20FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L20-N02-i111", IgnoreFieldsString2: "L20-N02-i112"},
				Level20OneOfIgnoreSelf1:           &pbinline.MessageLevel20_Level20One4String1{Level20One4String1: "L20-N02-i121"},
				Level20OneOfIgnoreParts1:          &pbinline.MessageLevel20_Level20One5String1{Level20One5String1: "L20-N02-i131"},
			}
		}
		level08N02 = &pbinline.MessageLevel08{
			Level08FString1:                   "L08-N02-s101",
			Level08FString2:                   "L08-N02-s111",
			Level08PString1:                   utils.PointerString("L08-N02-p101"),
			Level08PString2:                   utils.PointerString("L08-N02-p111"),
			Level08RString1:                   []string{"L08-N02-r101", "L08-N02-r102", "L08-N02-r103"},
			Level08RString2:                   []string{"L08-N02-r111", "L08-N02-r112", "L08-N02-r113"},
			Level08MString1:                   map[string]string{"L08-N02-k101": "L08-N02-v101", "L08-N02-k102": "L08-N02-v102", "L08-N02-k103": "L08-N02-v103"},
			Level08MString2:                   map[string]string{"L08-N02-k111": "L08-N02-v111", "L08-N02-k112": "L08-N02-v112", "L08-N02-k113": "L08-N02-v113"},
			Level08FMessageExtern1:            &pbexternal.Message1{FString1: "L08-N02-e101", FString2: "L08-N02-e102", FString3: "L08-N02-e103"},
			Level08FMessageExtern2:            &pbexternal.Message1{FString1: "L08-N02-e111", FString2: "L08-N02-e112", FString3: "L08-N02-e113"},
			Level08FMessageInline18:           level18N03,
			Level08FMessageInline19:           level19N02,
			Level08OneOfExtern1:               &pbinline.MessageLevel08_Level08One1MessageInline18{Level08One1MessageInline18: level18N04},
			Level08OneOfInline1:               &pbinline.MessageLevel08_Level08One2MessageInline20{Level08One2MessageInline20: level20N02},
			Level08OneOfOmitempty1:            &pbinline.MessageLevel08_Level08One3String1{Level08One3String1: "L08-N02-o101"},
			Level08FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level08FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L08-N02-i101"},
			Level08FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L08-N02-i111", IgnoreFieldsString2: "L08-N02-i112"},
			Level08OneOfIgnoreSelf1:           &pbinline.MessageLevel08_Level08One4String1{Level08One4String1: "L08-N02-i121"},
			Level08OneOfIgnoreParts1:          &pbinline.MessageLevel08_Level08One5String1{Level08One5String1: "L08-N02-i131"},
		}
	}
	level02N02 = &pbinline.MessageLevel02{
		Level02FString1:                   "L02-N02-s101",
		Level02FString2:                   "L02-N02-s111",
		Level02PString1:                   utils.PointerString("L02-N02-p101"),
		Level02PString2:                   utils.PointerString("L02-N02-p111"),
		Level02RString1:                   []string{"L02-N02-r101", "L02-N02-r102", "L02-N02-r103"},
		Level02RString2:                   []string{"L02-N02-r111", "L02-N02-r112", "L02-N02-r113"},
		Level02MString1:                   map[string]string{"L02-N02-k101": "L02-N02-v101", "L02-N02-k102": "L02-N02-v102", "L02-N02-k103": "L02-N02-v103"},
		Level02MString2:                   map[string]string{"L02-N02-k111": "L02-N02-v111", "L02-N02-k112": "L02-N02-v112", "L02-N02-k113": "L02-N02-v113"},
		Level02FMessageExtern1:            &pbexternal.Message1{FString1: "L02-N02-e101", FString2: "L02-N02-e102", FString3: "L02-N02-e103"},
		Level02FMessageExtern2:            &pbexternal.Message1{FString1: "L02-N02-e111", FString2: "L02-N02-e112", FString3: "L02-N02-e113"},
		Level02FMessageInline06:           level06N03,
		Level02FMessageInline07:           level07N02,
		Level02OneOfExtern1:               &pbinline.MessageLevel02_Level02One1MessageInline06{Level02One1MessageInline06: level06N04},
		Level02OneOfInline1:               &pbinline.MessageLevel02_Level02One2MessageInline08{Level02One2MessageInline08: level08N02},
		Level02OneOfOmitempty1:            &pbinline.MessageLevel02_Level02One3String1{Level02One3String1: "L02-N02-o101"},
		Level02FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
		Level02FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L02-N02-i101"},
		Level02FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L02-N02-i111", IgnoreFieldsString2: "L02-N02-i112"},
		Level02OneOfIgnoreSelf1:           &pbinline.MessageLevel02_Level02One4String1{Level02One4String1: "L02-N02-i121"},
		Level02OneOfIgnoreParts1:          &pbinline.MessageLevel02_Level02One5String1{Level02One5String1: "L02-N02-i131"},
	}

	{ // message for level04N01
		level10N01 = &pbinline.MessageLevel10{
			Level10FString1:                   "L10-N01-s101",
			Level10FString2:                   "L10-N01-s111",
			Level10PString1:                   utils.PointerString("L10-N01-p101"),
			Level10PString2:                   utils.PointerString("L10-N01-p111"),
			Level10RString1:                   []string{"L10-N01-r101", "L10-N01-r102", "L10-N01-r103"},
			Level10RString2:                   []string{"L10-N01-r111", "L10-N01-r112", "L10-N01-r113"},
			Level10MString1:                   map[string]string{"L10-N01-k101": "L10-N01-v101", "L10-N01-k102": "L10-N01-v102", "L10-N01-k103": "L10-N01-v103"},
			Level10MString2:                   map[string]string{"L10-N01-k111": "L10-N01-v111", "L10-N01-k112": "L10-N01-v112", "L10-N01-k113": "L10-N01-v113"},
			Level10FMessageExtern1:            &pbexternal.Message1{FString1: "L10-N01-e101", FString2: "L10-N01-e102", FString3: "L10-N01-e103"},
			Level10FMessageExtern2:            &pbexternal.Message1{FString1: "L10-N01-e111", FString2: "L10-N01-e112", FString3: "L10-N01-e113"},
			Level10OneOfExtern1:               &pbinline.MessageLevel10_Level10One1String1{Level10One1String1: "L10-N01-es101"},
			Level10OneOfInline1:               &pbinline.MessageLevel10_Level10One2String1{Level10One2String1: "L10-N01-is101"},
			Level10OneOfOmitempty1:            &pbinline.MessageLevel10_Level10One3String1{Level10One3String1: "L10-N01-o101"},
			Level10FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level10FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L10-N01-i101"},
			Level10FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L10-N01-i111", IgnoreFieldsString2: "L10-N01-i112"},
			Level10OneOfIgnoreSelf1:           &pbinline.MessageLevel10_Level10One4String1{Level10One4String1: "L10-N01-i121"},
			Level10OneOfIgnoreParts1:          &pbinline.MessageLevel10_Level10One5String1{Level10One5String1: "L10-N01-i131"},
		}
		level11N01 = &pbinline.MessageLevel11{
			Level11FString1:                   "L11-N01-s101",
			Level11FString2:                   "L11-N01-s111",
			Level11PString1:                   utils.PointerString("L11-N01-p101"),
			Level11PString2:                   utils.PointerString("L11-N01-p111"),
			Level11RString1:                   []string{"L11-N01-r101", "L11-N01-r102", "L11-N01-r103"},
			Level11RString2:                   []string{"L11-N01-r111", "L11-N01-r112", "L11-N01-r113"},
			Level11MString1:                   map[string]string{"L11-N01-k101": "L11-N01-v101", "L11-N01-k102": "L11-N01-v102", "L11-N01-k103": "L11-N01-v103"},
			Level11MString2:                   map[string]string{"L11-N01-k111": "L11-N01-v111", "L11-N01-k112": "L11-N01-v112", "L11-N01-k113": "L11-N01-v113"},
			Level11FMessageExtern1:            &pbexternal.Message1{FString1: "L11-N01-e101", FString2: "L11-N01-e102", FString3: "L11-N01-e103"},
			Level11FMessageExtern2:            &pbexternal.Message1{FString1: "L11-N01-e111", FString2: "L11-N01-e112", FString3: "L11-N01-e113"},
			Level11OneOfExtern1:               &pbinline.MessageLevel11_Level11One1String1{Level11One1String1: "L11-N01-es101"},
			Level11OneOfInline1:               &pbinline.MessageLevel11_Level11One2String1{Level11One2String1: "L11-N01-is101"},
			Level11OneOfOmitempty1:            &pbinline.MessageLevel11_Level11One3String1{Level11One3String1: "L11-N01-o101"},
			Level11FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level11FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L11-N01-i101"},
			Level11FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L11-N01-i111", IgnoreFieldsString2: "L11-N01-i112"},
			Level11OneOfIgnoreSelf1:           &pbinline.MessageLevel11_Level11One4String1{Level11One4String1: "L11-N01-i121"},
			Level11OneOfIgnoreParts1:          &pbinline.MessageLevel11_Level11One5String1{Level11One5String1: "L11-N01-i131"},
		}
		level10N02 = &pbinline.MessageLevel10{
			Level10FString1:                   "L10-N01-s101",
			Level10FString2:                   "L10-N01-s111",
			Level10PString1:                   utils.PointerString("L10-N01-p101"),
			Level10PString2:                   utils.PointerString("L10-N01-p111"),
			Level10RString1:                   []string{"L10-N01-r101", "L10-N01-r102", "L10-N01-r103"},
			Level10RString2:                   []string{"L10-N01-r111", "L10-N01-r112", "L10-N01-r113"},
			Level10MString1:                   map[string]string{"L10-N01-k101": "L10-N01-v101", "L10-N01-k102": "L10-N01-v102", "L10-N01-k103": "L10-N01-v103"},
			Level10MString2:                   map[string]string{"L10-N01-k111": "L10-N01-v111", "L10-N01-k112": "L10-N01-v112", "L10-N01-k113": "L10-N01-v113"},
			Level10FMessageExtern1:            &pbexternal.Message1{FString1: "L10-N01-e101", FString2: "L10-N01-e102", FString3: "L10-N01-e103"},
			Level10FMessageExtern2:            &pbexternal.Message1{FString1: "L10-N01-e111", FString2: "L10-N01-e112", FString3: "L10-N01-e113"},
			Level10OneOfExtern1:               &pbinline.MessageLevel10_Level10One1String1{Level10One1String1: "L10-N01-es101"},
			Level10OneOfInline1:               &pbinline.MessageLevel10_Level10One2String1{Level10One2String1: "L10-N01-is101"},
			Level10OneOfOmitempty1:            &pbinline.MessageLevel10_Level10One3String1{Level10One3String1: "L10-N01-o101"},
			Level10FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level10FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L10-N01-i101"},
			Level10FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L10-N01-i111", IgnoreFieldsString2: "L10-N01-i112"},
			Level10OneOfIgnoreSelf1:           &pbinline.MessageLevel10_Level10One4String1{Level10One4String1: "L10-N01-i121"},
			Level10OneOfIgnoreParts1:          &pbinline.MessageLevel10_Level10One5String1{Level10One5String1: "L10-N01-i131"},
		}
		level12N01 = &pbinline.MessageLevel12{
			Level12FString1:                   "L12-N01-s101",
			Level12FString2:                   "L12-N01-s111",
			Level12PString1:                   utils.PointerString("L12-N01-p101"),
			Level12PString2:                   utils.PointerString("L12-N01-p111"),
			Level12RString1:                   []string{"L12-N01-r101", "L12-N01-r102", "L12-N01-r103"},
			Level12RString2:                   []string{"L12-N01-r111", "L12-N01-r112", "L12-N01-r113"},
			Level12MString1:                   map[string]string{"L12-N01-k101": "L12-N01-v101", "L12-N01-k102": "L12-N01-v102", "L12-N01-k103": "L12-N01-v103"},
			Level12MString2:                   map[string]string{"L12-N01-k111": "L12-N01-v111", "L12-N01-k112": "L12-N01-v112", "L12-N01-k113": "L12-N01-v113"},
			Level12FMessageExtern1:            &pbexternal.Message1{FString1: "L12-N01-e101", FString2: "L12-N01-e102", FString3: "L12-N01-e103"},
			Level12FMessageExtern2:            &pbexternal.Message1{FString1: "L12-N01-e111", FString2: "L12-N01-e112", FString3: "L12-N01-e113"},
			Level12OneOfExtern1:               &pbinline.MessageLevel12_Level12One1String1{Level12One1String1: "L12-N01-es101"},
			Level12OneOfInline1:               &pbinline.MessageLevel12_Level12One2String1{Level12One2String1: "L12-N01-is101"},
			Level12OneOfOmitempty1:            &pbinline.MessageLevel12_Level12One3String1{Level12One3String1: "L12-N01-o101"},
			Level12FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
			Level12FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L12-N01-i101"},
			Level12FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L12-N01-i111", IgnoreFieldsString2: "L12-N01-i112"},
			Level12OneOfIgnoreSelf1:           &pbinline.MessageLevel12_Level12One4String1{Level12One4String1: "L12-N01-i121"},
			Level12OneOfIgnoreParts1:          &pbinline.MessageLevel12_Level12One5String1{Level12One5String1: "L12-N01-i131"},
		}
	}
	level04N01 = &pbinline.MessageLevel04{
		Level04FString1:                   "L04-N01-s101",
		Level04FString2:                   "L04-N01-s111",
		Level04PString1:                   utils.PointerString("L04-N01-p101"),
		Level04PString2:                   utils.PointerString("L04-N01-p111"),
		Level04RString1:                   []string{"L04-N01-r101", "L04-N01-r102", "L04-N01-r103"},
		Level04RString2:                   []string{"L04-N01-r111", "L04-N01-r112", "L04-N01-r113"},
		Level04MString1:                   map[string]string{"L04-N01-k101": "L04-N01-v101", "L04-N01-k102": "L04-N01-v102", "L04-N01-k103": "L04-N01-v103"},
		Level04MString2:                   map[string]string{"L04-N01-k111": "L04-N01-v111", "L04-N01-k112": "L04-N01-v112", "L04-N01-k113": "L04-N01-v113"},
		Level04FMessageExtern1:            &pbexternal.Message1{FString1: "L04-N01-e101", FString2: "L04-N01-e102", FString3: "L04-N01-e103"},
		Level04FMessageExtern2:            &pbexternal.Message1{FString1: "L04-N01-e111", FString2: "L04-N01-e112", FString3: "L04-N01-e113"},
		Level04FMessageInline10:           level10N01,
		Level04FMessageInline11:           level11N01,
		Level04OneOfExtern1:               &pbinline.MessageLevel04_Level04One1MessageInline10{Level04One1MessageInline10: level10N02},
		Level04OneOfInline1:               &pbinline.MessageLevel04_Level04One2MessageInline12{Level04One2MessageInline12: level12N01},
		Level04OneOfOmitempty1:            &pbinline.MessageLevel04_Level04One3String1{Level04One3String1: "L04-N01-o101"},
		Level04FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
		Level04FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L04-N01-i101"},
		Level04FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L04-N01-i111", IgnoreFieldsString2: "L04-N01-i112"},
		Level04OneOfIgnoreSelf1:           &pbinline.MessageLevel04_Level04One4String1{Level04One4String1: "L04-N01-i121"},
		Level04OneOfIgnoreParts1:          &pbinline.MessageLevel04_Level04One5String1{Level04One5String1: "L04-N01-i131"},
	}

	// level01 for test cases.
	level01 := &pbinline.MessageLevel01{
		Level01FString1:                   "L01-N01-s101",
		Level01FString2:                   "L01-N01-s111",
		Level01PString1:                   utils.PointerString("L01-N01-p101"),
		Level01PString2:                   utils.PointerString("L01-N01-p111"),
		Level01RString1:                   []string{"L01-N01-r101", "L01-N01-r102", "L01-N01-r103"},
		Level01RString2:                   []string{"L01-N01-r111", "L01-N01-r112", "L01-N01-r113"},
		Level01MString1:                   map[string]string{"L01-N01-k101": "L01-N01-v101", "L01-N01-k102": "L01-N01-v102", "L01-N01-k103": "L01-N01-v103"},
		Level01MString2:                   map[string]string{"L01-N01-k111": "L01-N01-v111", "L01-N01-k112": "L01-N01-v112", "L01-N01-k113": "L01-N01-v113"},
		Level01FMessageExtern1:            &pbexternal.Message1{FString1: "L01-N01-e101", FString2: "L01-N01-e102", FString3: "L01-N01-e103"},
		Level01FMessageExtern2:            &pbexternal.Message1{FString1: "L01-N01-e111", FString2: "L01-N01-e112", FString3: "L01-N01-e113"},
		Level01FMessageInline02:           level02N01,
		Level01FMessageInline03:           level03N01,
		Level01OneOfExtern1:               &pbinline.MessageLevel01_Level01One1MessageInline02{Level01One1MessageInline02: level02N02},
		Level01OneOfInline1:               &pbinline.MessageLevel01_Level01One2MessageInline04{Level01One2MessageInline04: level04N01},
		Level01OneOfOmitempty1:            &pbinline.MessageLevel01_Level01One3String1{Level01One3String1: "L01-N01-o101"},
		Level01FMessageInlineEmtpy:        &pbinline.MessageEmpty1{},
		Level01FMessageInlineIgnoreByRef:  &pbinline.MessageIgnoreByRef{IgnoreByRefString1: "L01-N01-i101"},
		Level01FMessageInlineIgnoreFields: &pbinline.MessageIgnoreFields{IgnoreFieldsString1: "L01-N01-i111", IgnoreFieldsString2: "L01-N01-i112"},
		Level01OneOfIgnoreSelf1:           &pbinline.MessageLevel01_Level01One4String1{Level01One4String1: "L01-N01-i121"},
		Level01OneOfIgnoreParts1:          &pbinline.MessageLevel01_Level01One5String1{Level01One5String1: "L01-N01-i131"},
	}

	var (
		err error
		bb  []byte
	)

	t.Run("Marshal", func(t *testing.T) {
		bb, err = level01.MarshalJSON()
		require.Nil(t, err)
		require.True(t, json.Valid(bb))
	})

	t.Run("Unmarshal", func(t *testing.T) {
		dataNew := &pbinline.MessageLevel01{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		{ // Reset ignore field.
			level01.Level01FMessageInlineEmtpy = nil
			level01.Level01FMessageInlineIgnoreByRef = nil
			level01.Level01FMessageInlineIgnoreFields = nil
			level01.Level01OneOfIgnoreSelf1 = nil
			level01.Level01OneOfIgnoreParts1 = nil

			level02N01.Level02FMessageInlineEmtpy = nil
			level02N01.Level02FMessageInlineIgnoreByRef = nil
			level02N01.Level02FMessageInlineIgnoreFields = nil
			level02N01.Level02OneOfIgnoreSelf1 = nil
			level02N01.Level02OneOfIgnoreParts1 = nil
			level02N02.Level02FMessageInlineEmtpy = nil
			level02N02.Level02FMessageInlineIgnoreByRef = nil
			level02N02.Level02FMessageInlineIgnoreFields = nil
			level02N02.Level02OneOfIgnoreSelf1 = nil
			level02N02.Level02OneOfIgnoreParts1 = nil

			level03N01.Level03FMessageInlineEmtpy = nil
			level03N01.Level03FMessageInlineIgnoreByRef = nil
			level03N01.Level03FMessageInlineIgnoreFields = nil
			level03N01.Level03OneOfIgnoreSelf1 = nil
			level03N01.Level03OneOfIgnoreParts1 = nil

			level04N01.Level04FMessageInlineEmtpy = nil
			level04N01.Level04FMessageInlineIgnoreByRef = nil
			level04N01.Level04FMessageInlineIgnoreFields = nil
			level04N01.Level04OneOfIgnoreSelf1 = nil
			level04N01.Level04OneOfIgnoreParts1 = nil

			level06N01.Level06FMessageInlineEmtpy = nil
			level06N01.Level06FMessageInlineIgnoreByRef = nil
			level06N01.Level06FMessageInlineIgnoreFields = nil
			level06N01.Level06OneOfIgnoreSelf1 = nil
			level06N01.Level06OneOfIgnoreParts1 = nil
			level06N02.Level06FMessageInlineEmtpy = nil
			level06N02.Level06FMessageInlineIgnoreByRef = nil
			level06N02.Level06FMessageInlineIgnoreFields = nil
			level06N02.Level06OneOfIgnoreSelf1 = nil
			level06N02.Level06OneOfIgnoreParts1 = nil
			level06N03.Level06FMessageInlineEmtpy = nil
			level06N03.Level06FMessageInlineIgnoreByRef = nil
			level06N03.Level06FMessageInlineIgnoreFields = nil
			level06N03.Level06OneOfIgnoreSelf1 = nil
			level06N03.Level06OneOfIgnoreParts1 = nil
			level06N04.Level06FMessageInlineEmtpy = nil
			level06N04.Level06FMessageInlineIgnoreByRef = nil
			level06N04.Level06FMessageInlineIgnoreFields = nil
			level06N04.Level06OneOfIgnoreSelf1 = nil
			level06N04.Level06OneOfIgnoreParts1 = nil

			level07N01.Level07FMessageInlineEmtpy = nil
			level07N01.Level07FMessageInlineIgnoreByRef = nil
			level07N01.Level07FMessageInlineIgnoreFields = nil
			level07N01.Level07OneOfIgnoreSelf1 = nil
			level07N01.Level07OneOfIgnoreParts1 = nil
			level07N02.Level07FMessageInlineEmtpy = nil
			level07N02.Level07FMessageInlineIgnoreByRef = nil
			level07N02.Level07FMessageInlineIgnoreFields = nil
			level07N02.Level07OneOfIgnoreSelf1 = nil
			level07N02.Level07OneOfIgnoreParts1 = nil

			level08N01.Level08FMessageInlineEmtpy = nil
			level08N01.Level08FMessageInlineIgnoreByRef = nil
			level08N01.Level08FMessageInlineIgnoreFields = nil
			level08N01.Level08OneOfIgnoreSelf1 = nil
			level08N01.Level08OneOfIgnoreParts1 = nil
			level08N02.Level08FMessageInlineEmtpy = nil
			level08N02.Level08FMessageInlineIgnoreByRef = nil
			level08N02.Level08FMessageInlineIgnoreFields = nil
			level08N02.Level08OneOfIgnoreSelf1 = nil
			level08N02.Level08OneOfIgnoreParts1 = nil

			level10N01.Level10FMessageInlineEmtpy = nil
			level10N01.Level10FMessageInlineIgnoreByRef = nil
			level10N01.Level10FMessageInlineIgnoreFields = nil
			level10N01.Level10OneOfIgnoreSelf1 = nil
			level10N01.Level10OneOfIgnoreParts1 = nil
			level10N02.Level10FMessageInlineEmtpy = nil
			level10N02.Level10FMessageInlineIgnoreByRef = nil
			level10N02.Level10FMessageInlineIgnoreFields = nil
			level10N02.Level10OneOfIgnoreSelf1 = nil
			level10N02.Level10OneOfIgnoreParts1 = nil

			level11N01.Level11FMessageInlineEmtpy = nil
			level11N01.Level11FMessageInlineIgnoreByRef = nil
			level11N01.Level11FMessageInlineIgnoreFields = nil
			level11N01.Level11OneOfIgnoreSelf1 = nil
			level11N01.Level11OneOfIgnoreParts1 = nil

			level12N01.Level12FMessageInlineEmtpy = nil
			level12N01.Level12FMessageInlineIgnoreByRef = nil
			level12N01.Level12FMessageInlineIgnoreFields = nil
			level12N01.Level12OneOfIgnoreSelf1 = nil
			level12N01.Level12OneOfIgnoreParts1 = nil

			level14N01.Level14FMessageInlineEmtpy = nil
			level14N01.Level14FMessageInlineIgnoreByRef = nil
			level14N01.Level14FMessageInlineIgnoreFields = nil
			level14N01.Level14OneOfIgnoreSelf1 = nil
			level14N01.Level14OneOfIgnoreParts1 = nil
			level14N02.Level14FMessageInlineEmtpy = nil
			level14N02.Level14FMessageInlineIgnoreByRef = nil
			level14N02.Level14FMessageInlineIgnoreFields = nil
			level14N02.Level14OneOfIgnoreSelf1 = nil
			level14N02.Level14OneOfIgnoreParts1 = nil
			level14N03.Level14FMessageInlineEmtpy = nil
			level14N03.Level14FMessageInlineIgnoreByRef = nil
			level14N03.Level14FMessageInlineIgnoreFields = nil
			level14N03.Level14OneOfIgnoreSelf1 = nil
			level14N03.Level14OneOfIgnoreParts1 = nil
			level14N04.Level14FMessageInlineEmtpy = nil
			level14N04.Level14FMessageInlineIgnoreByRef = nil
			level14N04.Level14FMessageInlineIgnoreFields = nil
			level14N04.Level14OneOfIgnoreSelf1 = nil
			level14N04.Level14OneOfIgnoreParts1 = nil
			level14N05.Level14FMessageInlineEmtpy = nil
			level14N05.Level14FMessageInlineIgnoreByRef = nil
			level14N05.Level14FMessageInlineIgnoreFields = nil
			level14N05.Level14OneOfIgnoreSelf1 = nil
			level14N05.Level14OneOfIgnoreParts1 = nil
			level14N06.Level14FMessageInlineEmtpy = nil
			level14N06.Level14FMessageInlineIgnoreByRef = nil
			level14N06.Level14FMessageInlineIgnoreFields = nil
			level14N06.Level14OneOfIgnoreSelf1 = nil
			level14N06.Level14OneOfIgnoreParts1 = nil
			level14N07.Level14FMessageInlineEmtpy = nil
			level14N07.Level14FMessageInlineIgnoreByRef = nil
			level14N07.Level14FMessageInlineIgnoreFields = nil
			level14N07.Level14OneOfIgnoreSelf1 = nil
			level14N07.Level14OneOfIgnoreParts1 = nil
			level14N08.Level14FMessageInlineEmtpy = nil
			level14N08.Level14FMessageInlineIgnoreByRef = nil
			level14N08.Level14FMessageInlineIgnoreFields = nil
			level14N08.Level14OneOfIgnoreSelf1 = nil
			level14N08.Level14OneOfIgnoreParts1 = nil

			level15N01.Level15FMessageInlineEmtpy = nil
			level15N01.Level15FMessageInlineIgnoreByRef = nil
			level15N01.Level15FMessageInlineIgnoreFields = nil
			level15N01.Level15OneOfIgnoreSelf1 = nil
			level15N01.Level15OneOfIgnoreParts1 = nil
			level15N02.Level15FMessageInlineEmtpy = nil
			level15N02.Level15FMessageInlineIgnoreByRef = nil
			level15N02.Level15FMessageInlineIgnoreFields = nil
			level15N02.Level15OneOfIgnoreSelf1 = nil
			level15N02.Level15OneOfIgnoreParts1 = nil
			level15N03.Level15FMessageInlineEmtpy = nil
			level15N03.Level15FMessageInlineIgnoreByRef = nil
			level15N03.Level15FMessageInlineIgnoreFields = nil
			level15N03.Level15OneOfIgnoreSelf1 = nil
			level15N03.Level15OneOfIgnoreParts1 = nil
			level15N04.Level15FMessageInlineEmtpy = nil
			level15N04.Level15FMessageInlineIgnoreByRef = nil
			level15N04.Level15FMessageInlineIgnoreFields = nil
			level15N04.Level15OneOfIgnoreSelf1 = nil
			level15N04.Level15OneOfIgnoreParts1 = nil

			level16N01.Level16FMessageInlineEmtpy = nil
			level16N01.Level16FMessageInlineIgnoreByRef = nil
			level16N01.Level16FMessageInlineIgnoreFields = nil
			level16N01.Level16OneOfIgnoreSelf1 = nil
			level16N01.Level16OneOfIgnoreParts1 = nil
			level16N02.Level16FMessageInlineEmtpy = nil
			level16N02.Level16FMessageInlineIgnoreByRef = nil
			level16N02.Level16FMessageInlineIgnoreFields = nil
			level16N02.Level16OneOfIgnoreSelf1 = nil
			level16N02.Level16OneOfIgnoreParts1 = nil
			level16N03.Level16FMessageInlineEmtpy = nil
			level16N03.Level16FMessageInlineIgnoreByRef = nil
			level16N03.Level16FMessageInlineIgnoreFields = nil
			level16N03.Level16OneOfIgnoreSelf1 = nil
			level16N03.Level16OneOfIgnoreParts1 = nil
			level16N04.Level16FMessageInlineEmtpy = nil
			level16N04.Level16FMessageInlineIgnoreByRef = nil
			level16N04.Level16FMessageInlineIgnoreFields = nil
			level16N04.Level16OneOfIgnoreSelf1 = nil
			level16N04.Level16OneOfIgnoreParts1 = nil

			level18N01.Level18FMessageInlineEmtpy = nil
			level18N01.Level18FMessageInlineIgnoreByRef = nil
			level18N01.Level18FMessageInlineIgnoreFields = nil
			level18N01.Level18OneOfIgnoreSelf1 = nil
			level18N01.Level18OneOfIgnoreParts1 = nil
			level18N02.Level18FMessageInlineEmtpy = nil
			level18N02.Level18FMessageInlineIgnoreByRef = nil
			level18N02.Level18FMessageInlineIgnoreFields = nil
			level18N02.Level18OneOfIgnoreSelf1 = nil
			level18N02.Level18OneOfIgnoreParts1 = nil
			level18N03.Level18FMessageInlineEmtpy = nil
			level18N03.Level18FMessageInlineIgnoreByRef = nil
			level18N03.Level18FMessageInlineIgnoreFields = nil
			level18N03.Level18OneOfIgnoreSelf1 = nil
			level18N03.Level18OneOfIgnoreParts1 = nil
			level18N04.Level18FMessageInlineEmtpy = nil
			level18N04.Level18FMessageInlineIgnoreByRef = nil
			level18N04.Level18FMessageInlineIgnoreFields = nil
			level18N04.Level18OneOfIgnoreSelf1 = nil
			level18N04.Level18OneOfIgnoreParts1 = nil

			level19N01.Level19FMessageInlineEmtpy = nil
			level19N01.Level19FMessageInlineIgnoreByRef = nil
			level19N01.Level19FMessageInlineIgnoreFields = nil
			level19N01.Level19OneOfIgnoreSelf1 = nil
			level19N01.Level19OneOfIgnoreParts1 = nil
			level19N02.Level19FMessageInlineEmtpy = nil
			level19N02.Level19FMessageInlineIgnoreByRef = nil
			level19N02.Level19FMessageInlineIgnoreFields = nil
			level19N02.Level19OneOfIgnoreSelf1 = nil
			level19N02.Level19OneOfIgnoreParts1 = nil

			level20N01.Level20FMessageInlineEmtpy = nil
			level20N01.Level20FMessageInlineIgnoreByRef = nil
			level20N01.Level20FMessageInlineIgnoreFields = nil
			level20N01.Level20OneOfIgnoreSelf1 = nil
			level20N01.Level20OneOfIgnoreParts1 = nil
			level20N02.Level20FMessageInlineEmtpy = nil
			level20N02.Level20FMessageInlineIgnoreByRef = nil
			level20N02.Level20FMessageInlineIgnoreFields = nil
			level20N02.Level20OneOfIgnoreSelf1 = nil
			level20N02.Level20OneOfIgnoreParts1 = nil

			level22N01.Level22FMessageInlineEmtpy = nil
			level22N01.Level22FMessageInlineIgnoreByRef = nil
			level22N01.Level22FMessageInlineIgnoreFields = nil
			level22N01.Level22OneOfIgnoreSelf1 = nil
			level22N01.Level22OneOfIgnoreParts1 = nil
			level22N02.Level22FMessageInlineEmtpy = nil
			level22N02.Level22FMessageInlineIgnoreByRef = nil
			level22N02.Level22FMessageInlineIgnoreFields = nil
			level22N02.Level22OneOfIgnoreSelf1 = nil
			level22N02.Level22OneOfIgnoreParts1 = nil
			level22N03.Level22FMessageInlineEmtpy = nil
			level22N03.Level22FMessageInlineIgnoreByRef = nil
			level22N03.Level22FMessageInlineIgnoreFields = nil
			level22N03.Level22OneOfIgnoreSelf1 = nil
			level22N03.Level22OneOfIgnoreParts1 = nil
			level22N04.Level22FMessageInlineEmtpy = nil
			level22N04.Level22FMessageInlineIgnoreByRef = nil
			level22N04.Level22FMessageInlineIgnoreFields = nil
			level22N04.Level22OneOfIgnoreSelf1 = nil
			level22N04.Level22OneOfIgnoreParts1 = nil
			level22N05.Level22FMessageInlineEmtpy = nil
			level22N05.Level22FMessageInlineIgnoreByRef = nil
			level22N05.Level22FMessageInlineIgnoreFields = nil
			level22N05.Level22OneOfIgnoreSelf1 = nil
			level22N05.Level22OneOfIgnoreParts1 = nil
			level22N06.Level22FMessageInlineEmtpy = nil
			level22N06.Level22FMessageInlineIgnoreByRef = nil
			level22N06.Level22FMessageInlineIgnoreFields = nil
			level22N06.Level22OneOfIgnoreSelf1 = nil
			level22N06.Level22OneOfIgnoreParts1 = nil
			level22N07.Level22FMessageInlineEmtpy = nil
			level22N07.Level22FMessageInlineIgnoreByRef = nil
			level22N07.Level22FMessageInlineIgnoreFields = nil
			level22N07.Level22OneOfIgnoreSelf1 = nil
			level22N07.Level22OneOfIgnoreParts1 = nil
			level22N08.Level22FMessageInlineEmtpy = nil
			level22N08.Level22FMessageInlineIgnoreByRef = nil
			level22N08.Level22FMessageInlineIgnoreFields = nil
			level22N08.Level22OneOfIgnoreSelf1 = nil
			level22N08.Level22OneOfIgnoreParts1 = nil
			level22N09.Level22FMessageInlineEmtpy = nil
			level22N09.Level22FMessageInlineIgnoreByRef = nil
			level22N09.Level22FMessageInlineIgnoreFields = nil
			level22N09.Level22OneOfIgnoreSelf1 = nil
			level22N09.Level22OneOfIgnoreParts1 = nil
			level22N10.Level22FMessageInlineEmtpy = nil
			level22N10.Level22FMessageInlineIgnoreByRef = nil
			level22N10.Level22FMessageInlineIgnoreFields = nil
			level22N10.Level22OneOfIgnoreSelf1 = nil
			level22N10.Level22OneOfIgnoreParts1 = nil
			level22N11.Level22FMessageInlineEmtpy = nil
			level22N11.Level22FMessageInlineIgnoreByRef = nil
			level22N11.Level22FMessageInlineIgnoreFields = nil
			level22N11.Level22OneOfIgnoreSelf1 = nil
			level22N11.Level22OneOfIgnoreParts1 = nil
			level22N12.Level22FMessageInlineEmtpy = nil
			level22N12.Level22FMessageInlineIgnoreByRef = nil
			level22N12.Level22FMessageInlineIgnoreFields = nil
			level22N12.Level22OneOfIgnoreSelf1 = nil
			level22N12.Level22OneOfIgnoreParts1 = nil
			level22N13.Level22FMessageInlineEmtpy = nil
			level22N13.Level22FMessageInlineIgnoreByRef = nil
			level22N13.Level22FMessageInlineIgnoreFields = nil
			level22N13.Level22OneOfIgnoreSelf1 = nil
			level22N13.Level22OneOfIgnoreParts1 = nil
			level22N14.Level22FMessageInlineEmtpy = nil
			level22N14.Level22FMessageInlineIgnoreByRef = nil
			level22N14.Level22FMessageInlineIgnoreFields = nil
			level22N14.Level22OneOfIgnoreSelf1 = nil
			level22N14.Level22OneOfIgnoreParts1 = nil
			level22N15.Level22FMessageInlineEmtpy = nil
			level22N15.Level22FMessageInlineIgnoreByRef = nil
			level22N15.Level22FMessageInlineIgnoreFields = nil
			level22N15.Level22OneOfIgnoreSelf1 = nil
			level22N15.Level22OneOfIgnoreParts1 = nil
			level22N16.Level22FMessageInlineEmtpy = nil
			level22N16.Level22FMessageInlineIgnoreByRef = nil
			level22N16.Level22FMessageInlineIgnoreFields = nil
			level22N16.Level22OneOfIgnoreSelf1 = nil
			level22N16.Level22OneOfIgnoreParts1 = nil

			level23N01.Level23FMessageInlineEmtpy = nil
			level23N01.Level23FMessageInlineIgnoreByRef = nil
			level23N01.Level23FMessageInlineIgnoreFields = nil
			level23N01.Level23OneOfIgnoreSelf1 = nil
			level23N01.Level23OneOfIgnoreParts1 = nil
			level23N02.Level23FMessageInlineEmtpy = nil
			level23N02.Level23FMessageInlineIgnoreByRef = nil
			level23N02.Level23FMessageInlineIgnoreFields = nil
			level23N02.Level23OneOfIgnoreSelf1 = nil
			level23N02.Level23OneOfIgnoreParts1 = nil
			level23N03.Level23FMessageInlineEmtpy = nil
			level23N03.Level23FMessageInlineIgnoreByRef = nil
			level23N03.Level23FMessageInlineIgnoreFields = nil
			level23N03.Level23OneOfIgnoreSelf1 = nil
			level23N03.Level23OneOfIgnoreParts1 = nil
			level23N04.Level23FMessageInlineEmtpy = nil
			level23N04.Level23FMessageInlineIgnoreByRef = nil
			level23N04.Level23FMessageInlineIgnoreFields = nil
			level23N04.Level23OneOfIgnoreSelf1 = nil
			level23N04.Level23OneOfIgnoreParts1 = nil
			level23N05.Level23FMessageInlineEmtpy = nil
			level23N05.Level23FMessageInlineIgnoreByRef = nil
			level23N05.Level23FMessageInlineIgnoreFields = nil
			level23N05.Level23OneOfIgnoreSelf1 = nil
			level23N05.Level23OneOfIgnoreParts1 = nil
			level23N06.Level23FMessageInlineEmtpy = nil
			level23N06.Level23FMessageInlineIgnoreByRef = nil
			level23N06.Level23FMessageInlineIgnoreFields = nil
			level23N06.Level23OneOfIgnoreSelf1 = nil
			level23N06.Level23OneOfIgnoreParts1 = nil
			level23N07.Level23FMessageInlineEmtpy = nil
			level23N07.Level23FMessageInlineIgnoreByRef = nil
			level23N07.Level23FMessageInlineIgnoreFields = nil
			level23N07.Level23OneOfIgnoreSelf1 = nil
			level23N07.Level23OneOfIgnoreParts1 = nil
			level23N08.Level23FMessageInlineEmtpy = nil
			level23N08.Level23FMessageInlineIgnoreByRef = nil
			level23N08.Level23FMessageInlineIgnoreFields = nil
			level23N08.Level23OneOfIgnoreSelf1 = nil
			level23N08.Level23OneOfIgnoreParts1 = nil

			level24N01.Level24FMessageInlineEmtpy = nil
			level24N01.Level24FMessageInlineIgnoreByRef = nil
			level24N01.Level24FMessageInlineIgnoreFields = nil
			level24N01.Level24OneOfIgnoreSelf1 = nil
			level24N01.Level24OneOfIgnoreParts1 = nil
			level24N02.Level24FMessageInlineEmtpy = nil
			level24N02.Level24FMessageInlineIgnoreByRef = nil
			level24N02.Level24FMessageInlineIgnoreFields = nil
			level24N02.Level24OneOfIgnoreSelf1 = nil
			level24N02.Level24OneOfIgnoreParts1 = nil
			level24N03.Level24FMessageInlineEmtpy = nil
			level24N03.Level24FMessageInlineIgnoreByRef = nil
			level24N03.Level24FMessageInlineIgnoreFields = nil
			level24N03.Level24OneOfIgnoreSelf1 = nil
			level24N03.Level24OneOfIgnoreParts1 = nil
			level24N04.Level24FMessageInlineEmtpy = nil
			level24N04.Level24FMessageInlineIgnoreByRef = nil
			level24N04.Level24FMessageInlineIgnoreFields = nil
			level24N04.Level24OneOfIgnoreSelf1 = nil
			level24N04.Level24OneOfIgnoreParts1 = nil
			level24N05.Level24FMessageInlineEmtpy = nil
			level24N05.Level24FMessageInlineIgnoreByRef = nil
			level24N05.Level24FMessageInlineIgnoreFields = nil
			level24N05.Level24OneOfIgnoreSelf1 = nil
			level24N05.Level24OneOfIgnoreParts1 = nil
			level24N06.Level24FMessageInlineEmtpy = nil
			level24N06.Level24FMessageInlineIgnoreByRef = nil
			level24N06.Level24FMessageInlineIgnoreFields = nil
			level24N06.Level24OneOfIgnoreSelf1 = nil
			level24N06.Level24OneOfIgnoreParts1 = nil
			level24N07.Level24FMessageInlineEmtpy = nil
			level24N07.Level24FMessageInlineIgnoreByRef = nil
			level24N07.Level24FMessageInlineIgnoreFields = nil
			level24N07.Level24OneOfIgnoreSelf1 = nil
			level24N07.Level24OneOfIgnoreParts1 = nil
			level24N08.Level24FMessageInlineEmtpy = nil
			level24N08.Level24FMessageInlineIgnoreByRef = nil
			level24N08.Level24FMessageInlineIgnoreFields = nil
			level24N08.Level24OneOfIgnoreSelf1 = nil
			level24N08.Level24OneOfIgnoreParts1 = nil

			level26N01.Level26FMessageInlineEmtpy = nil
			level26N01.Level26FMessageInlineIgnoreByRef = nil
			level26N01.Level26FMessageInlineIgnoreFields = nil
			level26N01.Level26OneOfIgnoreSelf1 = nil
			level26N01.Level26OneOfIgnoreParts1 = nil
			level26N02.Level26FMessageInlineEmtpy = nil
			level26N02.Level26FMessageInlineIgnoreByRef = nil
			level26N02.Level26FMessageInlineIgnoreFields = nil
			level26N02.Level26OneOfIgnoreSelf1 = nil
			level26N02.Level26OneOfIgnoreParts1 = nil
			level26N03.Level26FMessageInlineEmtpy = nil
			level26N03.Level26FMessageInlineIgnoreByRef = nil
			level26N03.Level26FMessageInlineIgnoreFields = nil
			level26N03.Level26OneOfIgnoreSelf1 = nil
			level26N03.Level26OneOfIgnoreParts1 = nil
			level26N04.Level26FMessageInlineEmtpy = nil
			level26N04.Level26FMessageInlineIgnoreByRef = nil
			level26N04.Level26FMessageInlineIgnoreFields = nil
			level26N04.Level26OneOfIgnoreSelf1 = nil
			level26N04.Level26OneOfIgnoreParts1 = nil
			level26N05.Level26FMessageInlineEmtpy = nil
			level26N05.Level26FMessageInlineIgnoreByRef = nil
			level26N05.Level26FMessageInlineIgnoreFields = nil
			level26N05.Level26OneOfIgnoreSelf1 = nil
			level26N05.Level26OneOfIgnoreParts1 = nil
			level26N06.Level26FMessageInlineEmtpy = nil
			level26N06.Level26FMessageInlineIgnoreByRef = nil
			level26N06.Level26FMessageInlineIgnoreFields = nil
			level26N06.Level26OneOfIgnoreSelf1 = nil
			level26N06.Level26OneOfIgnoreParts1 = nil
			level26N07.Level26FMessageInlineEmtpy = nil
			level26N07.Level26FMessageInlineIgnoreByRef = nil
			level26N07.Level26FMessageInlineIgnoreFields = nil
			level26N07.Level26OneOfIgnoreSelf1 = nil
			level26N07.Level26OneOfIgnoreParts1 = nil
			level26N08.Level26FMessageInlineEmtpy = nil
			level26N08.Level26FMessageInlineIgnoreByRef = nil
			level26N08.Level26FMessageInlineIgnoreFields = nil
			level26N08.Level26OneOfIgnoreSelf1 = nil
			level26N08.Level26OneOfIgnoreParts1 = nil

			level27N01.Level27FMessageInlineEmtpy = nil
			level27N01.Level27FMessageInlineIgnoreByRef = nil
			level27N01.Level27FMessageInlineIgnoreFields = nil
			level27N01.Level27OneOfIgnoreSelf1 = nil
			level27N01.Level27OneOfIgnoreParts1 = nil
			level27N02.Level27FMessageInlineEmtpy = nil
			level27N02.Level27FMessageInlineIgnoreByRef = nil
			level27N02.Level27FMessageInlineIgnoreFields = nil
			level27N02.Level27OneOfIgnoreSelf1 = nil
			level27N02.Level27OneOfIgnoreParts1 = nil
			level27N03.Level27FMessageInlineEmtpy = nil
			level27N03.Level27FMessageInlineIgnoreByRef = nil
			level27N03.Level27FMessageInlineIgnoreFields = nil
			level27N03.Level27OneOfIgnoreSelf1 = nil
			level27N03.Level27OneOfIgnoreParts1 = nil
			level27N04.Level27FMessageInlineEmtpy = nil
			level27N04.Level27FMessageInlineIgnoreByRef = nil
			level27N04.Level27FMessageInlineIgnoreFields = nil
			level27N04.Level27OneOfIgnoreSelf1 = nil
			level27N04.Level27OneOfIgnoreParts1 = nil

			level28N01.Level28FMessageInlineEmtpy = nil
			level28N01.Level28FMessageInlineIgnoreByRef = nil
			level28N01.Level28FMessageInlineIgnoreFields = nil
			level28N01.Level28OneOfIgnoreSelf1 = nil
			level28N01.Level28OneOfIgnoreParts1 = nil
			level28N02.Level28FMessageInlineEmtpy = nil
			level28N02.Level28FMessageInlineIgnoreByRef = nil
			level28N02.Level28FMessageInlineIgnoreFields = nil
			level28N02.Level28OneOfIgnoreSelf1 = nil
			level28N02.Level28OneOfIgnoreParts1 = nil
			level28N03.Level28FMessageInlineEmtpy = nil
			level28N03.Level28FMessageInlineIgnoreByRef = nil
			level28N03.Level28FMessageInlineIgnoreFields = nil
			level28N03.Level28OneOfIgnoreSelf1 = nil
			level28N03.Level28OneOfIgnoreParts1 = nil
			level28N04.Level28FMessageInlineEmtpy = nil
			level28N04.Level28FMessageInlineIgnoreByRef = nil
			level28N04.Level28FMessageInlineIgnoreFields = nil
			level28N04.Level28OneOfIgnoreSelf1 = nil
			level28N04.Level28OneOfIgnoreParts1 = nil

			level30N01.Level30FMessageInlineEmtpy = nil
			level30N01.Level30FMessageInlineIgnoreByRef = nil
			level30N01.Level30FMessageInlineIgnoreFields = nil
			level30N01.Level30OneOfIgnoreSelf1 = nil
			level30N01.Level30OneOfIgnoreParts1 = nil
			level30N02.Level30FMessageInlineEmtpy = nil
			level30N02.Level30FMessageInlineIgnoreByRef = nil
			level30N02.Level30FMessageInlineIgnoreFields = nil
			level30N02.Level30OneOfIgnoreSelf1 = nil
			level30N02.Level30OneOfIgnoreParts1 = nil
			level30N03.Level30FMessageInlineEmtpy = nil
			level30N03.Level30FMessageInlineIgnoreByRef = nil
			level30N03.Level30FMessageInlineIgnoreFields = nil
			level30N03.Level30OneOfIgnoreSelf1 = nil
			level30N03.Level30OneOfIgnoreParts1 = nil
			level30N04.Level30FMessageInlineEmtpy = nil
			level30N04.Level30FMessageInlineIgnoreByRef = nil
			level30N04.Level30FMessageInlineIgnoreFields = nil
			level30N04.Level30OneOfIgnoreSelf1 = nil
			level30N04.Level30OneOfIgnoreParts1 = nil
			level30N05.Level30FMessageInlineEmtpy = nil
			level30N05.Level30FMessageInlineIgnoreByRef = nil
			level30N05.Level30FMessageInlineIgnoreFields = nil
			level30N05.Level30OneOfIgnoreSelf1 = nil
			level30N05.Level30OneOfIgnoreParts1 = nil
			level30N06.Level30FMessageInlineEmtpy = nil
			level30N06.Level30FMessageInlineIgnoreByRef = nil
			level30N06.Level30FMessageInlineIgnoreFields = nil
			level30N06.Level30OneOfIgnoreSelf1 = nil
			level30N06.Level30OneOfIgnoreParts1 = nil
			level30N07.Level30FMessageInlineEmtpy = nil
			level30N07.Level30FMessageInlineIgnoreByRef = nil
			level30N07.Level30FMessageInlineIgnoreFields = nil
			level30N07.Level30OneOfIgnoreSelf1 = nil
			level30N07.Level30OneOfIgnoreParts1 = nil
			level30N08.Level30FMessageInlineEmtpy = nil
			level30N08.Level30FMessageInlineIgnoreByRef = nil
			level30N08.Level30FMessageInlineIgnoreFields = nil
			level30N08.Level30OneOfIgnoreSelf1 = nil
			level30N08.Level30OneOfIgnoreParts1 = nil
			level30N09.Level30FMessageInlineEmtpy = nil
			level30N09.Level30FMessageInlineIgnoreByRef = nil
			level30N09.Level30FMessageInlineIgnoreFields = nil
			level30N09.Level30OneOfIgnoreSelf1 = nil
			level30N09.Level30OneOfIgnoreParts1 = nil
			level30N10.Level30FMessageInlineEmtpy = nil
			level30N10.Level30FMessageInlineIgnoreByRef = nil
			level30N10.Level30FMessageInlineIgnoreFields = nil
			level30N10.Level30OneOfIgnoreSelf1 = nil
			level30N10.Level30OneOfIgnoreParts1 = nil
			level30N11.Level30FMessageInlineEmtpy = nil
			level30N11.Level30FMessageInlineIgnoreByRef = nil
			level30N11.Level30FMessageInlineIgnoreFields = nil
			level30N11.Level30OneOfIgnoreSelf1 = nil
			level30N11.Level30OneOfIgnoreParts1 = nil
			level30N12.Level30FMessageInlineEmtpy = nil
			level30N12.Level30FMessageInlineIgnoreByRef = nil
			level30N12.Level30FMessageInlineIgnoreFields = nil
			level30N12.Level30OneOfIgnoreSelf1 = nil
			level30N12.Level30OneOfIgnoreParts1 = nil
			level30N13.Level30FMessageInlineEmtpy = nil
			level30N13.Level30FMessageInlineIgnoreByRef = nil
			level30N13.Level30FMessageInlineIgnoreFields = nil
			level30N13.Level30OneOfIgnoreSelf1 = nil
			level30N13.Level30OneOfIgnoreParts1 = nil
			level30N14.Level30FMessageInlineEmtpy = nil
			level30N14.Level30FMessageInlineIgnoreByRef = nil
			level30N14.Level30FMessageInlineIgnoreFields = nil
			level30N14.Level30OneOfIgnoreSelf1 = nil
			level30N14.Level30OneOfIgnoreParts1 = nil
			level30N15.Level30FMessageInlineEmtpy = nil
			level30N15.Level30FMessageInlineIgnoreByRef = nil
			level30N15.Level30FMessageInlineIgnoreFields = nil
			level30N15.Level30OneOfIgnoreSelf1 = nil
			level30N15.Level30OneOfIgnoreParts1 = nil
			level30N16.Level30FMessageInlineEmtpy = nil
			level30N16.Level30FMessageInlineIgnoreByRef = nil
			level30N16.Level30FMessageInlineIgnoreFields = nil
			level30N16.Level30OneOfIgnoreSelf1 = nil
			level30N16.Level30OneOfIgnoreParts1 = nil

			level31N01.Level31FMessageInlineEmtpy = nil
			level31N01.Level31FMessageInlineIgnoreByRef = nil
			level31N01.Level31FMessageInlineIgnoreFields = nil
			level31N01.Level31OneOfIgnoreSelf1 = nil
			level31N01.Level31OneOfIgnoreParts1 = nil
			level31N02.Level31FMessageInlineEmtpy = nil
			level31N02.Level31FMessageInlineIgnoreByRef = nil
			level31N02.Level31FMessageInlineIgnoreFields = nil
			level31N02.Level31OneOfIgnoreSelf1 = nil
			level31N02.Level31OneOfIgnoreParts1 = nil
			level31N03.Level31FMessageInlineEmtpy = nil
			level31N03.Level31FMessageInlineIgnoreByRef = nil
			level31N03.Level31FMessageInlineIgnoreFields = nil
			level31N03.Level31OneOfIgnoreSelf1 = nil
			level31N03.Level31OneOfIgnoreParts1 = nil
			level31N04.Level31FMessageInlineEmtpy = nil
			level31N04.Level31FMessageInlineIgnoreByRef = nil
			level31N04.Level31FMessageInlineIgnoreFields = nil
			level31N04.Level31OneOfIgnoreSelf1 = nil
			level31N04.Level31OneOfIgnoreParts1 = nil
			level31N05.Level31FMessageInlineEmtpy = nil
			level31N05.Level31FMessageInlineIgnoreByRef = nil
			level31N05.Level31FMessageInlineIgnoreFields = nil
			level31N05.Level31OneOfIgnoreSelf1 = nil
			level31N05.Level31OneOfIgnoreParts1 = nil
			level31N06.Level31FMessageInlineEmtpy = nil
			level31N06.Level31FMessageInlineIgnoreByRef = nil
			level31N06.Level31FMessageInlineIgnoreFields = nil
			level31N06.Level31OneOfIgnoreSelf1 = nil
			level31N06.Level31OneOfIgnoreParts1 = nil
			level31N07.Level31FMessageInlineEmtpy = nil
			level31N07.Level31FMessageInlineIgnoreByRef = nil
			level31N07.Level31FMessageInlineIgnoreFields = nil
			level31N07.Level31OneOfIgnoreSelf1 = nil
			level31N07.Level31OneOfIgnoreParts1 = nil
			level31N08.Level31FMessageInlineEmtpy = nil
			level31N08.Level31FMessageInlineIgnoreByRef = nil
			level31N08.Level31FMessageInlineIgnoreFields = nil
			level31N08.Level31OneOfIgnoreSelf1 = nil
			level31N08.Level31OneOfIgnoreParts1 = nil

			level32N01.Level32FMessageInlineEmtpy = nil
			level32N01.Level32FMessageInlineIgnoreByRef = nil
			level32N01.Level32FMessageInlineIgnoreFields = nil
			level32N01.Level32OneOfIgnoreSelf1 = nil
			level32N01.Level32OneOfIgnoreParts1 = nil
			level32N02.Level32FMessageInlineEmtpy = nil
			level32N02.Level32FMessageInlineIgnoreByRef = nil
			level32N02.Level32FMessageInlineIgnoreFields = nil
			level32N02.Level32OneOfIgnoreSelf1 = nil
			level32N02.Level32OneOfIgnoreParts1 = nil
			level32N03.Level32FMessageInlineEmtpy = nil
			level32N03.Level32FMessageInlineIgnoreByRef = nil
			level32N03.Level32FMessageInlineIgnoreFields = nil
			level32N03.Level32OneOfIgnoreSelf1 = nil
			level32N03.Level32OneOfIgnoreParts1 = nil
			level32N04.Level32FMessageInlineEmtpy = nil
			level32N04.Level32FMessageInlineIgnoreByRef = nil
			level32N04.Level32FMessageInlineIgnoreFields = nil
			level32N04.Level32OneOfIgnoreSelf1 = nil
			level32N04.Level32OneOfIgnoreParts1 = nil
			level32N05.Level32FMessageInlineEmtpy = nil
			level32N05.Level32FMessageInlineIgnoreByRef = nil
			level32N05.Level32FMessageInlineIgnoreFields = nil
			level32N05.Level32OneOfIgnoreSelf1 = nil
			level32N05.Level32OneOfIgnoreParts1 = nil
			level32N06.Level32FMessageInlineEmtpy = nil
			level32N06.Level32FMessageInlineIgnoreByRef = nil
			level32N06.Level32FMessageInlineIgnoreFields = nil
			level32N06.Level32OneOfIgnoreSelf1 = nil
			level32N06.Level32OneOfIgnoreParts1 = nil
			level32N07.Level32FMessageInlineEmtpy = nil
			level32N07.Level32FMessageInlineIgnoreByRef = nil
			level32N07.Level32FMessageInlineIgnoreFields = nil
			level32N07.Level32OneOfIgnoreSelf1 = nil
			level32N07.Level32OneOfIgnoreParts1 = nil
			level32N08.Level32FMessageInlineEmtpy = nil
			level32N08.Level32FMessageInlineIgnoreByRef = nil
			level32N08.Level32FMessageInlineIgnoreFields = nil
			level32N08.Level32OneOfIgnoreSelf1 = nil
			level32N08.Level32OneOfIgnoreParts1 = nil

			level34N01.Level34FMessageInlineEmtpy = nil
			level34N01.Level34FMessageInlineIgnoreByRef = nil
			level34N01.Level34FMessageInlineIgnoreFields = nil
			level34N01.Level34OneOfIgnoreSelf1 = nil
			level34N01.Level34OneOfIgnoreParts1 = nil
			level34N02.Level34FMessageInlineEmtpy = nil
			level34N02.Level34FMessageInlineIgnoreByRef = nil
			level34N02.Level34FMessageInlineIgnoreFields = nil
			level34N02.Level34OneOfIgnoreSelf1 = nil
			level34N02.Level34OneOfIgnoreParts1 = nil
			level34N03.Level34FMessageInlineEmtpy = nil
			level34N03.Level34FMessageInlineIgnoreByRef = nil
			level34N03.Level34FMessageInlineIgnoreFields = nil
			level34N03.Level34OneOfIgnoreSelf1 = nil
			level34N03.Level34OneOfIgnoreParts1 = nil
			level34N04.Level34FMessageInlineEmtpy = nil
			level34N04.Level34FMessageInlineIgnoreByRef = nil
			level34N04.Level34FMessageInlineIgnoreFields = nil
			level34N04.Level34OneOfIgnoreSelf1 = nil
			level34N04.Level34OneOfIgnoreParts1 = nil
			level34N05.Level34FMessageInlineEmtpy = nil
			level34N05.Level34FMessageInlineIgnoreByRef = nil
			level34N05.Level34FMessageInlineIgnoreFields = nil
			level34N05.Level34OneOfIgnoreSelf1 = nil
			level34N05.Level34OneOfIgnoreParts1 = nil
			level34N06.Level34FMessageInlineEmtpy = nil
			level34N06.Level34FMessageInlineIgnoreByRef = nil
			level34N06.Level34FMessageInlineIgnoreFields = nil
			level34N06.Level34OneOfIgnoreSelf1 = nil
			level34N06.Level34OneOfIgnoreParts1 = nil
			level34N07.Level34FMessageInlineEmtpy = nil
			level34N07.Level34FMessageInlineIgnoreByRef = nil
			level34N07.Level34FMessageInlineIgnoreFields = nil
			level34N07.Level34OneOfIgnoreSelf1 = nil
			level34N07.Level34OneOfIgnoreParts1 = nil
			level34N08.Level34FMessageInlineEmtpy = nil
			level34N08.Level34FMessageInlineIgnoreByRef = nil
			level34N08.Level34FMessageInlineIgnoreFields = nil
			level34N08.Level34OneOfIgnoreSelf1 = nil
			level34N08.Level34OneOfIgnoreParts1 = nil

			level35N01.Level35FMessageInlineEmtpy = nil
			level35N01.Level35FMessageInlineIgnoreByRef = nil
			level35N01.Level35FMessageInlineIgnoreFields = nil
			level35N01.Level35OneOfIgnoreSelf1 = nil
			level35N01.Level35OneOfIgnoreParts1 = nil
			level35N02.Level35FMessageInlineEmtpy = nil
			level35N02.Level35FMessageInlineIgnoreByRef = nil
			level35N02.Level35FMessageInlineIgnoreFields = nil
			level35N02.Level35OneOfIgnoreSelf1 = nil
			level35N02.Level35OneOfIgnoreParts1 = nil
			level35N03.Level35FMessageInlineEmtpy = nil
			level35N03.Level35FMessageInlineIgnoreByRef = nil
			level35N03.Level35FMessageInlineIgnoreFields = nil
			level35N03.Level35OneOfIgnoreSelf1 = nil
			level35N03.Level35OneOfIgnoreParts1 = nil
			level35N04.Level35FMessageInlineEmtpy = nil
			level35N04.Level35FMessageInlineIgnoreByRef = nil
			level35N04.Level35FMessageInlineIgnoreFields = nil
			level35N04.Level35OneOfIgnoreSelf1 = nil
			level35N04.Level35OneOfIgnoreParts1 = nil

			level36N01.Level36FMessageInlineEmtpy = nil
			level36N01.Level36FMessageInlineIgnoreByRef = nil
			level36N01.Level36FMessageInlineIgnoreFields = nil
			level36N01.Level36OneOfIgnoreSelf1 = nil
			level36N01.Level36OneOfIgnoreParts1 = nil
			level36N02.Level36FMessageInlineEmtpy = nil
			level36N02.Level36FMessageInlineIgnoreByRef = nil
			level36N02.Level36FMessageInlineIgnoreFields = nil
			level36N02.Level36OneOfIgnoreSelf1 = nil
			level36N02.Level36OneOfIgnoreParts1 = nil
			level36N03.Level36FMessageInlineEmtpy = nil
			level36N03.Level36FMessageInlineIgnoreByRef = nil
			level36N03.Level36FMessageInlineIgnoreFields = nil
			level36N03.Level36OneOfIgnoreSelf1 = nil
			level36N03.Level36OneOfIgnoreParts1 = nil
			level36N04.Level36FMessageInlineEmtpy = nil
			level36N04.Level36FMessageInlineIgnoreByRef = nil
			level36N04.Level36FMessageInlineIgnoreFields = nil
			level36N04.Level36OneOfIgnoreSelf1 = nil
			level36N04.Level36OneOfIgnoreParts1 = nil

			level38N01.Level38FMessageInlineEmtpy = nil
			level38N01.Level38FMessageInlineIgnoreByRef = nil
			level38N01.Level38FMessageInlineIgnoreFields = nil
			level38N01.Level38OneOfIgnoreSelf1 = nil
			level38N01.Level38OneOfIgnoreParts1 = nil
			level38N02.Level38FMessageInlineEmtpy = nil
			level38N02.Level38FMessageInlineIgnoreByRef = nil
			level38N02.Level38FMessageInlineIgnoreFields = nil
			level38N02.Level38OneOfIgnoreSelf1 = nil
			level38N02.Level38OneOfIgnoreParts1 = nil
			level38N03.Level38FMessageInlineEmtpy = nil
			level38N03.Level38FMessageInlineIgnoreByRef = nil
			level38N03.Level38FMessageInlineIgnoreFields = nil
			level38N03.Level38OneOfIgnoreSelf1 = nil
			level38N03.Level38OneOfIgnoreParts1 = nil
			level38N04.Level38FMessageInlineEmtpy = nil
			level38N04.Level38FMessageInlineIgnoreByRef = nil
			level38N04.Level38FMessageInlineIgnoreFields = nil
			level38N04.Level38OneOfIgnoreSelf1 = nil
			level38N04.Level38OneOfIgnoreParts1 = nil
			level38N05.Level38FMessageInlineEmtpy = nil
			level38N05.Level38FMessageInlineIgnoreByRef = nil
			level38N05.Level38FMessageInlineIgnoreFields = nil
			level38N05.Level38OneOfIgnoreSelf1 = nil
			level38N05.Level38OneOfIgnoreParts1 = nil
			level38N06.Level38FMessageInlineEmtpy = nil
			level38N06.Level38FMessageInlineIgnoreByRef = nil
			level38N06.Level38FMessageInlineIgnoreFields = nil
			level38N06.Level38OneOfIgnoreSelf1 = nil
			level38N06.Level38OneOfIgnoreParts1 = nil
			level38N07.Level38FMessageInlineEmtpy = nil
			level38N07.Level38FMessageInlineIgnoreByRef = nil
			level38N07.Level38FMessageInlineIgnoreFields = nil
			level38N07.Level38OneOfIgnoreSelf1 = nil
			level38N07.Level38OneOfIgnoreParts1 = nil
			level38N08.Level38FMessageInlineEmtpy = nil
			level38N08.Level38FMessageInlineIgnoreByRef = nil
			level38N08.Level38FMessageInlineIgnoreFields = nil
			level38N08.Level38OneOfIgnoreSelf1 = nil
			level38N08.Level38OneOfIgnoreParts1 = nil
			level38N09.Level38FMessageInlineEmtpy = nil
			level38N09.Level38FMessageInlineIgnoreByRef = nil
			level38N09.Level38FMessageInlineIgnoreFields = nil
			level38N09.Level38OneOfIgnoreSelf1 = nil
			level38N09.Level38OneOfIgnoreParts1 = nil
			level38N10.Level38FMessageInlineEmtpy = nil
			level38N10.Level38FMessageInlineIgnoreByRef = nil
			level38N10.Level38FMessageInlineIgnoreFields = nil
			level38N10.Level38OneOfIgnoreSelf1 = nil
			level38N10.Level38OneOfIgnoreParts1 = nil
			level38N11.Level38FMessageInlineEmtpy = nil
			level38N11.Level38FMessageInlineIgnoreByRef = nil
			level38N11.Level38FMessageInlineIgnoreFields = nil
			level38N11.Level38OneOfIgnoreSelf1 = nil
			level38N11.Level38OneOfIgnoreParts1 = nil
			level38N12.Level38FMessageInlineEmtpy = nil
			level38N12.Level38FMessageInlineIgnoreByRef = nil
			level38N12.Level38FMessageInlineIgnoreFields = nil
			level38N12.Level38OneOfIgnoreSelf1 = nil
			level38N12.Level38OneOfIgnoreParts1 = nil
			level38N13.Level38FMessageInlineEmtpy = nil
			level38N13.Level38FMessageInlineIgnoreByRef = nil
			level38N13.Level38FMessageInlineIgnoreFields = nil
			level38N13.Level38OneOfIgnoreSelf1 = nil
			level38N13.Level38OneOfIgnoreParts1 = nil
			level38N14.Level38FMessageInlineEmtpy = nil
			level38N14.Level38FMessageInlineIgnoreByRef = nil
			level38N14.Level38FMessageInlineIgnoreFields = nil
			level38N14.Level38OneOfIgnoreSelf1 = nil
			level38N14.Level38OneOfIgnoreParts1 = nil
			level38N15.Level38FMessageInlineEmtpy = nil
			level38N15.Level38FMessageInlineIgnoreByRef = nil
			level38N15.Level38FMessageInlineIgnoreFields = nil
			level38N15.Level38OneOfIgnoreSelf1 = nil
			level38N15.Level38OneOfIgnoreParts1 = nil
			level38N16.Level38FMessageInlineEmtpy = nil
			level38N16.Level38FMessageInlineIgnoreByRef = nil
			level38N16.Level38FMessageInlineIgnoreFields = nil
			level38N16.Level38OneOfIgnoreSelf1 = nil
			level38N16.Level38OneOfIgnoreParts1 = nil
			level38N17.Level38FMessageInlineEmtpy = nil
			level38N17.Level38FMessageInlineIgnoreByRef = nil
			level38N17.Level38FMessageInlineIgnoreFields = nil
			level38N17.Level38OneOfIgnoreSelf1 = nil
			level38N17.Level38OneOfIgnoreParts1 = nil
			level38N18.Level38FMessageInlineEmtpy = nil
			level38N18.Level38FMessageInlineIgnoreByRef = nil
			level38N18.Level38FMessageInlineIgnoreFields = nil
			level38N18.Level38OneOfIgnoreSelf1 = nil
			level38N18.Level38OneOfIgnoreParts1 = nil
			level38N19.Level38FMessageInlineEmtpy = nil
			level38N19.Level38FMessageInlineIgnoreByRef = nil
			level38N19.Level38FMessageInlineIgnoreFields = nil
			level38N19.Level38OneOfIgnoreSelf1 = nil
			level38N19.Level38OneOfIgnoreParts1 = nil
			level38N20.Level38FMessageInlineEmtpy = nil
			level38N20.Level38FMessageInlineIgnoreByRef = nil
			level38N20.Level38FMessageInlineIgnoreFields = nil
			level38N20.Level38OneOfIgnoreSelf1 = nil
			level38N20.Level38OneOfIgnoreParts1 = nil
			level38N21.Level38FMessageInlineEmtpy = nil
			level38N21.Level38FMessageInlineIgnoreByRef = nil
			level38N21.Level38FMessageInlineIgnoreFields = nil
			level38N21.Level38OneOfIgnoreSelf1 = nil
			level38N21.Level38OneOfIgnoreParts1 = nil
			level38N22.Level38FMessageInlineEmtpy = nil
			level38N22.Level38FMessageInlineIgnoreByRef = nil
			level38N22.Level38FMessageInlineIgnoreFields = nil
			level38N22.Level38OneOfIgnoreSelf1 = nil
			level38N22.Level38OneOfIgnoreParts1 = nil
			level38N23.Level38FMessageInlineEmtpy = nil
			level38N23.Level38FMessageInlineIgnoreByRef = nil
			level38N23.Level38FMessageInlineIgnoreFields = nil
			level38N23.Level38OneOfIgnoreSelf1 = nil
			level38N23.Level38OneOfIgnoreParts1 = nil
			level38N24.Level38FMessageInlineEmtpy = nil
			level38N24.Level38FMessageInlineIgnoreByRef = nil
			level38N24.Level38FMessageInlineIgnoreFields = nil
			level38N24.Level38OneOfIgnoreSelf1 = nil
			level38N24.Level38OneOfIgnoreParts1 = nil
			level38N25.Level38FMessageInlineEmtpy = nil
			level38N25.Level38FMessageInlineIgnoreByRef = nil
			level38N25.Level38FMessageInlineIgnoreFields = nil
			level38N25.Level38OneOfIgnoreSelf1 = nil
			level38N25.Level38OneOfIgnoreParts1 = nil
			level38N26.Level38FMessageInlineEmtpy = nil
			level38N26.Level38FMessageInlineIgnoreByRef = nil
			level38N26.Level38FMessageInlineIgnoreFields = nil
			level38N26.Level38OneOfIgnoreSelf1 = nil
			level38N26.Level38OneOfIgnoreParts1 = nil
			level38N27.Level38FMessageInlineEmtpy = nil
			level38N27.Level38FMessageInlineIgnoreByRef = nil
			level38N27.Level38FMessageInlineIgnoreFields = nil
			level38N27.Level38OneOfIgnoreSelf1 = nil
			level38N27.Level38OneOfIgnoreParts1 = nil
			level38N28.Level38FMessageInlineEmtpy = nil
			level38N28.Level38FMessageInlineIgnoreByRef = nil
			level38N28.Level38FMessageInlineIgnoreFields = nil
			level38N28.Level38OneOfIgnoreSelf1 = nil
			level38N28.Level38OneOfIgnoreParts1 = nil
			level38N29.Level38FMessageInlineEmtpy = nil
			level38N29.Level38FMessageInlineIgnoreByRef = nil
			level38N29.Level38FMessageInlineIgnoreFields = nil
			level38N29.Level38OneOfIgnoreSelf1 = nil
			level38N29.Level38OneOfIgnoreParts1 = nil
			level38N30.Level38FMessageInlineEmtpy = nil
			level38N30.Level38FMessageInlineIgnoreByRef = nil
			level38N30.Level38FMessageInlineIgnoreFields = nil
			level38N30.Level38OneOfIgnoreSelf1 = nil
			level38N30.Level38OneOfIgnoreParts1 = nil
			level38N31.Level38FMessageInlineEmtpy = nil
			level38N31.Level38FMessageInlineIgnoreByRef = nil
			level38N31.Level38FMessageInlineIgnoreFields = nil
			level38N31.Level38OneOfIgnoreSelf1 = nil
			level38N31.Level38OneOfIgnoreParts1 = nil
			level38N32.Level38FMessageInlineEmtpy = nil
			level38N32.Level38FMessageInlineIgnoreByRef = nil
			level38N32.Level38FMessageInlineIgnoreFields = nil
			level38N32.Level38OneOfIgnoreSelf1 = nil
			level38N32.Level38OneOfIgnoreParts1 = nil

			level39N01.Level39FMessageInlineEmtpy = nil
			level39N01.Level39FMessageInlineIgnoreByRef = nil
			level39N01.Level39FMessageInlineIgnoreFields = nil
			level39N01.Level39OneOfIgnoreSelf1 = nil
			level39N01.Level39OneOfIgnoreParts1 = nil
			level39N02.Level39FMessageInlineEmtpy = nil
			level39N02.Level39FMessageInlineIgnoreByRef = nil
			level39N02.Level39FMessageInlineIgnoreFields = nil
			level39N02.Level39OneOfIgnoreSelf1 = nil
			level39N02.Level39OneOfIgnoreParts1 = nil
			level39N03.Level39FMessageInlineEmtpy = nil
			level39N03.Level39FMessageInlineIgnoreByRef = nil
			level39N03.Level39FMessageInlineIgnoreFields = nil
			level39N03.Level39OneOfIgnoreSelf1 = nil
			level39N03.Level39OneOfIgnoreParts1 = nil
			level39N04.Level39FMessageInlineEmtpy = nil
			level39N04.Level39FMessageInlineIgnoreByRef = nil
			level39N04.Level39FMessageInlineIgnoreFields = nil
			level39N04.Level39OneOfIgnoreSelf1 = nil
			level39N04.Level39OneOfIgnoreParts1 = nil
			level39N05.Level39FMessageInlineEmtpy = nil
			level39N05.Level39FMessageInlineIgnoreByRef = nil
			level39N05.Level39FMessageInlineIgnoreFields = nil
			level39N05.Level39OneOfIgnoreSelf1 = nil
			level39N05.Level39OneOfIgnoreParts1 = nil
			level39N06.Level39FMessageInlineEmtpy = nil
			level39N06.Level39FMessageInlineIgnoreByRef = nil
			level39N06.Level39FMessageInlineIgnoreFields = nil
			level39N06.Level39OneOfIgnoreSelf1 = nil
			level39N06.Level39OneOfIgnoreParts1 = nil
			level39N07.Level39FMessageInlineEmtpy = nil
			level39N07.Level39FMessageInlineIgnoreByRef = nil
			level39N07.Level39FMessageInlineIgnoreFields = nil
			level39N07.Level39OneOfIgnoreSelf1 = nil
			level39N07.Level39OneOfIgnoreParts1 = nil
			level39N08.Level39FMessageInlineEmtpy = nil
			level39N08.Level39FMessageInlineIgnoreByRef = nil
			level39N08.Level39FMessageInlineIgnoreFields = nil
			level39N08.Level39OneOfIgnoreSelf1 = nil
			level39N08.Level39OneOfIgnoreParts1 = nil
			level39N09.Level39FMessageInlineEmtpy = nil
			level39N09.Level39FMessageInlineIgnoreByRef = nil
			level39N09.Level39FMessageInlineIgnoreFields = nil
			level39N09.Level39OneOfIgnoreSelf1 = nil
			level39N09.Level39OneOfIgnoreParts1 = nil
			level39N10.Level39FMessageInlineEmtpy = nil
			level39N10.Level39FMessageInlineIgnoreByRef = nil
			level39N10.Level39FMessageInlineIgnoreFields = nil
			level39N10.Level39OneOfIgnoreSelf1 = nil
			level39N10.Level39OneOfIgnoreParts1 = nil
			level39N11.Level39FMessageInlineEmtpy = nil
			level39N11.Level39FMessageInlineIgnoreByRef = nil
			level39N11.Level39FMessageInlineIgnoreFields = nil
			level39N11.Level39OneOfIgnoreSelf1 = nil
			level39N11.Level39OneOfIgnoreParts1 = nil
			level39N12.Level39FMessageInlineEmtpy = nil
			level39N12.Level39FMessageInlineIgnoreByRef = nil
			level39N12.Level39FMessageInlineIgnoreFields = nil
			level39N12.Level39OneOfIgnoreSelf1 = nil
			level39N12.Level39OneOfIgnoreParts1 = nil
			level39N13.Level39FMessageInlineEmtpy = nil
			level39N13.Level39FMessageInlineIgnoreByRef = nil
			level39N13.Level39FMessageInlineIgnoreFields = nil
			level39N13.Level39OneOfIgnoreSelf1 = nil
			level39N13.Level39OneOfIgnoreParts1 = nil
			level39N14.Level39FMessageInlineEmtpy = nil
			level39N14.Level39FMessageInlineIgnoreByRef = nil
			level39N14.Level39FMessageInlineIgnoreFields = nil
			level39N14.Level39OneOfIgnoreSelf1 = nil
			level39N14.Level39OneOfIgnoreParts1 = nil
			level39N15.Level39FMessageInlineEmtpy = nil
			level39N15.Level39FMessageInlineIgnoreByRef = nil
			level39N15.Level39FMessageInlineIgnoreFields = nil
			level39N15.Level39OneOfIgnoreSelf1 = nil
			level39N15.Level39OneOfIgnoreParts1 = nil
			level39N16.Level39FMessageInlineEmtpy = nil
			level39N16.Level39FMessageInlineIgnoreByRef = nil
			level39N16.Level39FMessageInlineIgnoreFields = nil
			level39N16.Level39OneOfIgnoreSelf1 = nil
			level39N16.Level39OneOfIgnoreParts1 = nil

			level40N01.Level40FMessageInlineEmtpy = nil
			level40N01.Level40FMessageInlineIgnoreByRef = nil
			level40N01.Level40FMessageInlineIgnoreFields = nil
			level40N01.Level40OneOfIgnoreSelf1 = nil
			level40N01.Level40OneOfIgnoreParts1 = nil
			level40N02.Level40FMessageInlineEmtpy = nil
			level40N02.Level40FMessageInlineIgnoreByRef = nil
			level40N02.Level40FMessageInlineIgnoreFields = nil
			level40N02.Level40OneOfIgnoreSelf1 = nil
			level40N02.Level40OneOfIgnoreParts1 = nil
			level40N03.Level40FMessageInlineEmtpy = nil
			level40N03.Level40FMessageInlineIgnoreByRef = nil
			level40N03.Level40FMessageInlineIgnoreFields = nil
			level40N03.Level40OneOfIgnoreSelf1 = nil
			level40N03.Level40OneOfIgnoreParts1 = nil
			level40N04.Level40FMessageInlineEmtpy = nil
			level40N04.Level40FMessageInlineIgnoreByRef = nil
			level40N04.Level40FMessageInlineIgnoreFields = nil
			level40N04.Level40OneOfIgnoreSelf1 = nil
			level40N04.Level40OneOfIgnoreParts1 = nil
			level40N05.Level40FMessageInlineEmtpy = nil
			level40N05.Level40FMessageInlineIgnoreByRef = nil
			level40N05.Level40FMessageInlineIgnoreFields = nil
			level40N05.Level40OneOfIgnoreSelf1 = nil
			level40N05.Level40OneOfIgnoreParts1 = nil
			level40N06.Level40FMessageInlineEmtpy = nil
			level40N06.Level40FMessageInlineIgnoreByRef = nil
			level40N06.Level40FMessageInlineIgnoreFields = nil
			level40N06.Level40OneOfIgnoreSelf1 = nil
			level40N06.Level40OneOfIgnoreParts1 = nil
			level40N07.Level40FMessageInlineEmtpy = nil
			level40N07.Level40FMessageInlineIgnoreByRef = nil
			level40N07.Level40FMessageInlineIgnoreFields = nil
			level40N07.Level40OneOfIgnoreSelf1 = nil
			level40N07.Level40OneOfIgnoreParts1 = nil
			level40N08.Level40FMessageInlineEmtpy = nil
			level40N08.Level40FMessageInlineIgnoreByRef = nil
			level40N08.Level40FMessageInlineIgnoreFields = nil
			level40N08.Level40OneOfIgnoreSelf1 = nil
			level40N08.Level40OneOfIgnoreParts1 = nil
			level40N09.Level40FMessageInlineEmtpy = nil
			level40N09.Level40FMessageInlineIgnoreByRef = nil
			level40N09.Level40FMessageInlineIgnoreFields = nil
			level40N09.Level40OneOfIgnoreSelf1 = nil
			level40N09.Level40OneOfIgnoreParts1 = nil
			level40N10.Level40FMessageInlineEmtpy = nil
			level40N10.Level40FMessageInlineIgnoreByRef = nil
			level40N10.Level40FMessageInlineIgnoreFields = nil
			level40N10.Level40OneOfIgnoreSelf1 = nil
			level40N10.Level40OneOfIgnoreParts1 = nil
			level40N11.Level40FMessageInlineEmtpy = nil
			level40N11.Level40FMessageInlineIgnoreByRef = nil
			level40N11.Level40FMessageInlineIgnoreFields = nil
			level40N11.Level40OneOfIgnoreSelf1 = nil
			level40N11.Level40OneOfIgnoreParts1 = nil
			level40N12.Level40FMessageInlineEmtpy = nil
			level40N12.Level40FMessageInlineIgnoreByRef = nil
			level40N12.Level40FMessageInlineIgnoreFields = nil
			level40N12.Level40OneOfIgnoreSelf1 = nil
			level40N12.Level40OneOfIgnoreParts1 = nil
			level40N13.Level40FMessageInlineEmtpy = nil
			level40N13.Level40FMessageInlineIgnoreByRef = nil
			level40N13.Level40FMessageInlineIgnoreFields = nil
			level40N13.Level40OneOfIgnoreSelf1 = nil
			level40N13.Level40OneOfIgnoreParts1 = nil
			level40N14.Level40FMessageInlineEmtpy = nil
			level40N14.Level40FMessageInlineIgnoreByRef = nil
			level40N14.Level40FMessageInlineIgnoreFields = nil
			level40N14.Level40OneOfIgnoreSelf1 = nil
			level40N14.Level40OneOfIgnoreParts1 = nil
			level40N15.Level40FMessageInlineEmtpy = nil
			level40N15.Level40FMessageInlineIgnoreByRef = nil
			level40N15.Level40FMessageInlineIgnoreFields = nil
			level40N15.Level40OneOfIgnoreSelf1 = nil
			level40N15.Level40OneOfIgnoreParts1 = nil
			level40N16.Level40FMessageInlineEmtpy = nil
			level40N16.Level40FMessageInlineIgnoreByRef = nil
			level40N16.Level40FMessageInlineIgnoreFields = nil
			level40N16.Level40OneOfIgnoreSelf1 = nil
			level40N16.Level40OneOfIgnoreParts1 = nil

		}
		require.Equal(t, level01, dataNew)
	})

	// Cases for test check if the json key is at the correct level.
	t.Run("Check", func(t *testing.T) {

		data := map[string]interface{}{}
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		t.Run("Level1", func(t *testing.T) {
			// the keys in level1.
			keys := []string{
				"level01_c_f_string1",
				"level01_c_m_string1",
				"level01_c_p_string1",
				"level01_c_r_string1",
				"level01_f_message_extern1",
				"level01_f_message_extern2",
				"level01_f_message_inline_emtpy",
				"level01_f_message_inline_ignore_fields",
				"level01_f_string2",
				"level01_m_string2",
				"level01_one1_extern1",
				"level01_one3_string1",
				"level01_p_string2",
				"level01_r_string2",
				"level02_c_f_string1",
				"level02_c_m_string1",
				"level02_c_p_string1",
				"level02_c_r_string1",
				"level02_f_message_extern1",
				"level02_f_message_extern2",
				"level02_f_message_inline_emtpy",
				"level02_f_message_inline_ignore_fields",
				"level02_f_string2",
				"level02_m_string2",
				"level02_one1_extern1",
				"level02_one3_string1",
				"level02_p_string2",
				"level02_r_string2",
				"level03_c_f_string1",
				"level03_c_m_string1",
				"level03_c_one2_string1",
				"level03_c_p_string1",
				"level03_c_r_string1",
				"level03_f_message_extern1",
				"level03_f_message_extern2",
				"level03_f_message_inline_emtpy",
				"level03_f_message_inline_ignore_fields",
				"level03_f_string2",
				"level03_m_string2",
				"level03_one1_extern1",
				"level03_one3_string1",
				"level03_p_string2",
				"level03_r_string2",
				"level04_c_f_string1",
				"level04_c_m_string1",
				"level04_c_p_string1",
				"level04_c_r_string1",
				"level04_f_message_extern1",
				"level04_f_message_extern2",
				"level04_f_message_inline_emtpy",
				"level04_f_message_inline_ignore_fields",
				"level04_f_string2",
				"level04_m_string2",
				"level04_one1_extern1",
				"level04_one3_string1",
				"level04_p_string2",
				"level04_r_string2",
				"level06_c_f_string1",
				"level06_c_m_string1",
				"level06_c_p_string1",
				"level06_c_r_string1",
				"level06_f_message_extern1",
				"level06_f_message_extern2",
				"level06_f_message_inline_emtpy",
				"level06_f_message_inline_ignore_fields",
				"level06_f_string2",
				"level06_m_string2",
				"level06_one1_extern1",
				"level06_one3_string1",
				"level06_p_string2",
				"level06_r_string2",
				"level07_c_f_string1",
				"level07_c_m_string1",
				"level07_c_one2_string1",
				"level07_c_p_string1",
				"level07_c_r_string1",
				"level07_f_message_extern1",
				"level07_f_message_extern2",
				"level07_f_message_inline_emtpy",
				"level07_f_message_inline_ignore_fields",
				"level07_f_string2",
				"level07_m_string2",
				"level07_one1_extern1",
				"level07_one3_string1",
				"level07_p_string2",
				"level07_r_string2",
				"level08_c_f_string1",
				"level08_c_m_string1",
				"level08_c_p_string1",
				"level08_c_r_string1",
				"level08_f_message_extern1",
				"level08_f_message_extern2",
				"level08_f_message_inline_emtpy",
				"level08_f_message_inline_ignore_fields",
				"level08_f_string2",
				"level08_m_string2",
				"level08_one1_extern1",
				"level08_one3_string1",
				"level08_p_string2",
				"level08_r_string2",
				"level10_c_f_string1",
				"level10_c_m_string1",
				"level10_c_one2_string1",
				"level10_c_p_string1",
				"level10_c_r_string1",
				"level10_f_message_extern1",
				"level10_f_message_extern2",
				"level10_f_message_inline_emtpy",
				"level10_f_message_inline_ignore_fields",
				"level10_f_string2",
				"level10_m_string2",
				"level10_one1_extern1",
				"level10_one3_string1",
				"level10_p_string2",
				"level10_r_string2",
				"level11_c_f_string1",
				"level11_c_m_string1",
				"level11_c_one2_string1",
				"level11_c_p_string1",
				"level11_c_r_string1",
				"level11_f_message_extern1",
				"level11_f_message_extern2",
				"level11_f_message_inline_emtpy",
				"level11_f_message_inline_ignore_fields",
				"level11_f_string2",
				"level11_m_string2",
				"level11_one1_extern1",
				"level11_one3_string1",
				"level11_p_string2",
				"level11_r_string2",
				"level12_c_f_string1",
				"level12_c_m_string1",
				"level12_c_one2_string1",
				"level12_c_p_string1",
				"level12_c_r_string1",
				"level12_f_message_extern1",
				"level12_f_message_extern2",
				"level12_f_message_inline_emtpy",
				"level12_f_message_inline_ignore_fields",
				"level12_f_string2",
				"level12_m_string2",
				"level12_one1_extern1",
				"level12_one3_string1",
				"level12_p_string2",
				"level12_r_string2",
				"level14_c_f_string1",
				"level14_c_m_string1",
				"level14_c_p_string1",
				"level14_c_r_string1",
				"level14_f_message_extern1",
				"level14_f_message_extern2",
				"level14_f_message_inline_emtpy",
				"level14_f_message_inline_ignore_fields",
				"level14_f_string2",
				"level14_m_string2",
				"level14_one1_extern1",
				"level14_one3_string1",
				"level14_p_string2",
				"level14_r_string2",
				"level15_c_f_string1",
				"level15_c_m_string1",
				"level15_c_one2_string1",
				"level15_c_p_string1",
				"level15_c_r_string1",
				"level15_f_message_extern1",
				"level15_f_message_extern2",
				"level15_f_message_inline_emtpy",
				"level15_f_message_inline_ignore_fields",
				"level15_f_string2",
				"level15_m_string2",
				"level15_one1_extern1",
				"level15_one3_string1",
				"level15_p_string2",
				"level15_r_string2",
				"level16_c_f_string1",
				"level16_c_m_string1",
				"level16_c_p_string1",
				"level16_c_r_string1",
				"level16_f_message_extern1",
				"level16_f_message_extern2",
				"level16_f_message_inline_emtpy",
				"level16_f_message_inline_ignore_fields",
				"level16_f_string2",
				"level16_m_string2",
				"level16_one1_extern1",
				"level16_one3_string1",
				"level16_p_string2",
				"level16_r_string2",
				"level18_c_f_string1",
				"level18_c_m_string1",
				"level18_c_one2_string1",
				"level18_c_p_string1",
				"level18_c_r_string1",
				"level18_f_message_extern1",
				"level18_f_message_extern2",
				"level18_f_message_inline_emtpy",
				"level18_f_message_inline_ignore_fields",
				"level18_f_string2",
				"level18_m_string2",
				"level18_one1_extern1",
				"level18_one3_string1",
				"level18_p_string2",
				"level18_r_string2",
				"level19_c_f_string1",
				"level19_c_m_string1",
				"level19_c_one2_string1",
				"level19_c_p_string1",
				"level19_c_r_string1",
				"level19_f_message_extern1",
				"level19_f_message_extern2",
				"level19_f_message_inline_emtpy",
				"level19_f_message_inline_ignore_fields",
				"level19_f_string2",
				"level19_m_string2",
				"level19_one1_extern1",
				"level19_one3_string1",
				"level19_p_string2",
				"level19_r_string2",
				"level20_c_f_string1",
				"level20_c_m_string1",
				"level20_c_one2_string1",
				"level20_c_p_string1",
				"level20_c_r_string1",
				"level20_f_message_extern1",
				"level20_f_message_extern2",
				"level20_f_message_inline_emtpy",
				"level20_f_message_inline_ignore_fields",
				"level20_f_string2",
				"level20_m_string2",
				"level20_one1_extern1",
				"level20_one3_string1",
				"level20_p_string2",
				"level20_r_string2",
				"level22_c_f_string1",
				"level22_c_m_string1",
				"level22_c_p_string1",
				"level22_c_r_string1",
				"level22_f_message_extern1",
				"level22_f_message_extern2",
				"level22_f_message_inline_emtpy",
				"level22_f_message_inline_ignore_fields",
				"level22_f_string2",
				"level22_m_string2",
				"level22_one1_extern1",
				"level22_one3_string1",
				"level22_p_string2",
				"level22_r_string2",
				"level23_c_f_string1",
				"level23_c_m_string1",
				"level23_c_one2_string1",
				"level23_c_p_string1",
				"level23_c_r_string1",
				"level23_f_message_extern1",
				"level23_f_message_extern2",
				"level23_f_message_inline_emtpy",
				"level23_f_message_inline_ignore_fields",
				"level23_f_string2",
				"level23_m_string2",
				"level23_one1_extern1",
				"level23_one3_string1",
				"level23_p_string2",
				"level23_r_string2",
				"level24_c_f_string1",
				"level24_c_m_string1",
				"level24_c_one2_string1",
				"level24_c_p_string1",
				"level24_c_r_string1",
				"level24_f_message_extern1",
				"level24_f_message_extern2",
				"level24_f_message_inline_emtpy",
				"level24_f_message_inline_ignore_fields",
				"level24_f_string2",
				"level24_m_string2",
				"level24_one1_extern1",
				"level24_one3_string1",
				"level24_p_string2",
				"level24_r_string2",
				"level26_c_f_string1",
				"level26_c_m_string1",
				"level26_c_p_string1",
				"level26_c_r_string1",
				"level26_f_message_extern1",
				"level26_f_message_extern2",
				"level26_f_message_inline_emtpy",
				"level26_f_message_inline_ignore_fields",
				"level26_f_string2",
				"level26_m_string2",
				"level26_one1_extern1",
				"level26_one3_string1",
				"level26_p_string2",
				"level26_r_string2",
				"level27_c_f_string1",
				"level27_c_m_string1",
				"level27_c_one2_string1",
				"level27_c_p_string1",
				"level27_c_r_string1",
				"level27_f_message_extern1",
				"level27_f_message_extern2",
				"level27_f_message_inline_emtpy",
				"level27_f_message_inline_ignore_fields",
				"level27_f_string2",
				"level27_m_string2",
				"level27_one1_extern1",
				"level27_one3_string1",
				"level27_p_string2",
				"level27_r_string2",
				"level28_c_f_string1",
				"level28_c_m_string1",
				"level28_c_p_string1",
				"level28_c_r_string1",
				"level28_f_message_extern1",
				"level28_f_message_extern2",
				"level28_f_message_inline_emtpy",
				"level28_f_message_inline_ignore_fields",
				"level28_f_string2",
				"level28_m_string2",
				"level28_one1_extern1",
				"level28_one3_string1",
				"level28_p_string2",
				"level28_r_string2",
				"level30_c_f_string1",
				"level30_c_m_string1",
				"level30_c_one2_string1",
				"level30_c_p_string1",
				"level30_c_r_string1",
				"level30_f_message_extern1",
				"level30_f_message_extern2",
				"level30_f_message_inline_emtpy",
				"level30_f_message_inline_ignore_fields",
				"level30_f_string2",
				"level30_m_string2",
				"level30_one1_extern1",
				"level30_one3_string1",
				"level30_p_string2",
				"level30_r_string2",
				"level31_c_f_string1",
				"level31_c_m_string1",
				"level31_c_one2_string1",
				"level31_c_p_string1",
				"level31_c_r_string1",
				"level31_f_message_extern1",
				"level31_f_message_extern2",
				"level31_f_message_inline_emtpy",
				"level31_f_message_inline_ignore_fields",
				"level31_f_string2",
				"level31_m_string2",
				"level31_one1_extern1",
				"level31_one3_string1",
				"level31_p_string2",
				"level31_r_string2",
				"level32_c_f_string1",
				"level32_c_m_string1",
				"level32_c_one2_string1",
				"level32_c_p_string1",
				"level32_c_r_string1",
				"level32_f_message_extern1",
				"level32_f_message_extern2",
				"level32_f_message_inline_emtpy",
				"level32_f_message_inline_ignore_fields",
				"level32_f_string2",
				"level32_m_string2",
				"level32_one1_extern1",
				"level32_one3_string1",
				"level32_p_string2",
				"level32_r_string2",
				"level34_c_f_string1",
				"level34_c_m_string1",
				"level34_c_one2_string1",
				"level34_c_p_string1",
				"level34_c_r_string1",
				"level34_f_message_extern1",
				"level34_f_message_extern2",
				"level34_f_message_inline_emtpy",
				"level34_f_message_inline_ignore_fields",
				"level34_f_string2",
				"level34_m_string2",
				"level34_one1_extern1",
				"level34_one3_string1",
				"level34_p_string2",
				"level34_r_string2",
				"level35_c_f_string1",
				"level35_c_m_string1",
				"level35_c_one2_string1",
				"level35_c_p_string1",
				"level35_c_r_string1",
				"level35_f_message_extern1",
				"level35_f_message_extern2",
				"level35_f_message_inline_emtpy",
				"level35_f_message_inline_ignore_fields",
				"level35_f_string2",
				"level35_m_string2",
				"level35_one1_extern1",
				"level35_one3_string1",
				"level35_p_string2",
				"level35_r_string2",
				"level36_c_f_string1",
				"level36_c_m_string1",
				"level36_c_one2_string1",
				"level36_c_p_string1",
				"level36_c_r_string1",
				"level36_f_message_extern1",
				"level36_f_message_extern2",
				"level36_f_message_inline_emtpy",
				"level36_f_message_inline_ignore_fields",
				"level36_f_string2",
				"level36_m_string2",
				"level36_one1_extern1",
				"level36_one3_string1",
				"level36_p_string2",
				"level36_r_string2",
				"level38_c_f_string1",
				"level38_c_m_string1",
				"level38_c_one2_string1",
				"level38_c_p_string1",
				"level38_c_r_string1",
				"level38_f_message_extern1",
				"level38_f_message_extern2",
				"level38_f_message_inline_emtpy",
				"level38_f_message_inline_ignore_fields",
				"level38_f_string2",
				"level38_m_string2",
				"level38_one1_extern1",
				"level38_one3_string1",
				"level38_p_string2",
				"level38_r_string2",
				"level39_c_f_string1",
				"level39_c_m_string1",
				"level39_c_one2_string1",
				"level39_c_p_string1",
				"level39_c_r_string1",
				"level39_f_message_extern1",
				"level39_f_message_extern2",
				"level39_f_message_inline_emtpy",
				"level39_f_message_inline_ignore_fields",
				"level39_f_string2",
				"level39_m_string2",
				"level39_one1_extern1",
				"level39_one3_string1",
				"level39_p_string2",
				"level39_r_string2",
				"level40_c_f_string1",
				"level40_c_m_string1",
				"level40_c_one2_string1",
				"level40_c_p_string1",
				"level40_c_r_string1",
				"level40_f_message_extern1",
				"level40_f_message_extern2",
				"level40_f_message_inline_emtpy",
				"level40_f_message_inline_ignore_fields",
				"level40_f_string2",
				"level40_m_string2",
				"level40_one1_extern1",
				"level40_one3_string1",
				"level40_p_string2",
				"level40_r_string2",
			}
			for _, key := range keys {
				val, ok := data[key]
				require.True(t, ok, fmt.Sprintf("the json key [%s] not at in level 1", key))
				require.NotNil(t, val, fmt.Sprintf("the value of json key [%s] is nil in level 1", key))
			}
		})

		t.Run("Level2", func(t *testing.T) {
			// level22_one1_extern1
			t.Run("level22_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level38_c_f_string1",
					"level38_c_m_string1",
					"level38_c_one2_string1",
					"level38_c_p_string1",
					"level38_c_r_string1",
					"level38_f_message_extern1",
					"level38_f_message_extern2",
					"level38_f_message_inline_emtpy",
					"level38_f_message_inline_ignore_fields",
					"level38_f_string2",
					"level38_m_string2",
					"level38_one1_extern1",
					"level38_one3_string1",
					"level38_p_string2",
					"level38_r_string2",
				}
				values := data["level22_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level22_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level22_one1_extern1", key))
				}
			})

			// level14_one1_extern1
			t.Run("level14_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level22_c_f_string1",
					"level22_c_m_string1",
					"level22_c_p_string1",
					"level22_c_r_string1",
					"level22_f_message_extern1",
					"level22_f_message_extern2",
					"level22_f_message_inline_emtpy",
					"level22_f_message_inline_ignore_fields",
					"level22_f_string2",
					"level22_m_string2",
					"level22_one1_extern1",
					"level22_one3_string1",
					"level22_p_string2",
					"level22_r_string2",
					"level38_c_f_string1",
					"level38_c_m_string1",
					"level38_c_one2_string1",
					"level38_c_p_string1",
					"level38_c_r_string1",
					"level38_f_message_extern1",
					"level38_f_message_extern2",
					"level38_f_message_inline_emtpy",
					"level38_f_message_inline_ignore_fields",
					"level38_f_string2",
					"level38_m_string2",
					"level38_one1_extern1",
					"level38_one3_string1",
					"level38_p_string2",
					"level38_r_string2",
					"level39_c_f_string1",
					"level39_c_m_string1",
					"level39_c_one2_string1",
					"level39_c_p_string1",
					"level39_c_r_string1",
					"level39_f_message_extern1",
					"level39_f_message_extern2",
					"level39_f_message_inline_emtpy",
					"level39_f_message_inline_ignore_fields",
					"level39_f_string2",
					"level39_m_string2",
					"level39_one1_extern1",
					"level39_one3_string1",
					"level39_p_string2",
					"level39_r_string2",
					"level40_c_f_string1",
					"level40_c_m_string1",
					"level40_c_one2_string1",
					"level40_c_p_string1",
					"level40_c_r_string1",
					"level40_f_message_extern1",
					"level40_f_message_extern2",
					"level40_f_message_inline_emtpy",
					"level40_f_message_inline_ignore_fields",
					"level40_f_string2",
					"level40_m_string2",
					"level40_one1_extern1",
					"level40_one3_string1",
					"level40_p_string2",
					"level40_r_string2",
				}
				values := data["level14_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level14_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level14_one1_extern1", key))
				}
			})

			// level06_one1_extern1
			t.Run("level06_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level14_c_f_string1",
					"level14_c_m_string1",
					"level14_c_p_string1",
					"level14_c_r_string1",
					"level14_f_message_extern1",
					"level14_f_message_extern2",
					"level14_f_message_inline_emtpy",
					"level14_f_message_inline_ignore_fields",
					"level14_f_string2",
					"level14_m_string2",
					"level14_one1_extern1",
					"level14_one3_string1",
					"level14_p_string2",
					"level14_r_string2",
					"level22_c_f_string1",
					"level22_c_m_string1",
					"level22_c_p_string1",
					"level22_c_r_string1",
					"level22_f_message_extern1",
					"level22_f_message_extern2",
					"level22_f_message_inline_emtpy",
					"level22_f_message_inline_ignore_fields",
					"level22_f_string2",
					"level22_m_string2",
					"level22_one1_extern1",
					"level22_one3_string1",
					"level22_p_string2",
					"level22_r_string2",
					"level23_c_f_string1",
					"level23_c_m_string1",
					"level23_c_one2_string1",
					"level23_c_p_string1",
					"level23_c_r_string1",
					"level23_f_message_extern1",
					"level23_f_message_extern2",
					"level23_f_message_inline_emtpy",
					"level23_f_message_inline_ignore_fields",
					"level23_f_string2",
					"level23_m_string2",
					"level23_one1_extern1",
					"level23_one3_string1",
					"level23_p_string2",
					"level23_r_string2",
					"level24_c_f_string1",
					"level24_c_m_string1",
					"level24_c_one2_string1",
					"level24_c_p_string1",
					"level24_c_r_string1",
					"level24_f_message_extern1",
					"level24_f_message_extern2",
					"level24_f_message_inline_emtpy",
					"level24_f_message_inline_ignore_fields",
					"level24_f_string2",
					"level24_m_string2",
					"level24_one1_extern1",
					"level24_one3_string1",
					"level24_p_string2",
					"level24_r_string2",
					"level38_c_f_string1",
					"level38_c_m_string1",
					"level38_c_one2_string1",
					"level38_c_p_string1",
					"level38_c_r_string1",
					"level38_f_message_extern1",
					"level38_f_message_extern2",
					"level38_f_message_inline_emtpy",
					"level38_f_message_inline_ignore_fields",
					"level38_f_string2",
					"level38_m_string2",
					"level38_one1_extern1",
					"level38_one3_string1",
					"level38_p_string2",
					"level38_r_string2",
					"level39_c_f_string1",
					"level39_c_m_string1",
					"level39_c_one2_string1",
					"level39_c_p_string1",
					"level39_c_r_string1",
					"level39_f_message_extern1",
					"level39_f_message_extern2",
					"level39_f_message_inline_emtpy",
					"level39_f_message_inline_ignore_fields",
					"level39_f_string2",
					"level39_m_string2",
					"level39_one1_extern1",
					"level39_one3_string1",
					"level39_p_string2",
					"level39_r_string2",
					"level40_c_f_string1",
					"level40_c_m_string1",
					"level40_c_one2_string1",
					"level40_c_p_string1",
					"level40_c_r_string1",
					"level40_f_message_extern1",
					"level40_f_message_extern2",
					"level40_f_message_inline_emtpy",
					"level40_f_message_inline_ignore_fields",
					"level40_f_string2",
					"level40_m_string2",
					"level40_one1_extern1",
					"level40_one3_string1",
					"level40_p_string2",
					"level40_r_string2",
				}
				values := data["level06_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level06_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level06_one1_extern1", key))
				}
			})

			// level26_one1_extern1
			t.Run("level26_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level30_c_f_string1",
					"level30_c_m_string1",
					"level30_c_one2_string1",
					"level30_c_p_string1",
					"level30_c_r_string1",
					"level30_f_message_extern1",
					"level30_f_message_extern2",
					"level30_f_message_inline_emtpy",
					"level30_f_message_inline_ignore_fields",
					"level30_f_string2",
					"level30_m_string2",
					"level30_one1_extern1",
					"level30_one3_string1",
					"level30_p_string2",
					"level30_r_string2",
				}
				values := data["level26_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level26_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level26_one1_extern1", key))
				}
			})

			// level16_one1_extern1
			t.Run("level16_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level26_c_f_string1",
					"level26_c_m_string1",
					"level26_c_p_string1",
					"level26_c_r_string1",
					"level26_f_message_extern1",
					"level26_f_message_extern2",
					"level26_f_message_inline_emtpy",
					"level26_f_message_inline_ignore_fields",
					"level26_f_string2",
					"level26_m_string2",
					"level26_one1_extern1",
					"level26_one3_string1",
					"level26_p_string2",
					"level26_r_string2",
					"level30_c_f_string1",
					"level30_c_m_string1",
					"level30_c_one2_string1",
					"level30_c_p_string1",
					"level30_c_r_string1",
					"level30_f_message_extern1",
					"level30_f_message_extern2",
					"level30_f_message_inline_emtpy",
					"level30_f_message_inline_ignore_fields",
					"level30_f_string2",
					"level30_m_string2",
					"level30_one1_extern1",
					"level30_one3_string1",
					"level30_p_string2",
					"level30_r_string2",
					"level31_c_f_string1",
					"level31_c_m_string1",
					"level31_c_one2_string1",
					"level31_c_p_string1",
					"level31_c_r_string1",
					"level31_f_message_extern1",
					"level31_f_message_extern2",
					"level31_f_message_inline_emtpy",
					"level31_f_message_inline_ignore_fields",
					"level31_f_string2",
					"level31_m_string2",
					"level31_one1_extern1",
					"level31_one3_string1",
					"level31_p_string2",
					"level31_r_string2",
					"level32_c_f_string1",
					"level32_c_m_string1",
					"level32_c_one2_string1",
					"level32_c_p_string1",
					"level32_c_r_string1",
					"level32_f_message_extern1",
					"level32_f_message_extern2",
					"level32_f_message_inline_emtpy",
					"level32_f_message_inline_ignore_fields",
					"level32_f_string2",
					"level32_m_string2",
					"level32_one1_extern1",
					"level32_one3_string1",
					"level32_p_string2",
					"level32_r_string2",
				}
				values := data["level16_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level16_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level16_one1_extern1", key))
				}
			})

			// level28_one1_extern1
			t.Run("level28_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level34_c_f_string1",
					"level34_c_m_string1",
					"level34_c_one2_string1",
					"level34_c_p_string1",
					"level34_c_r_string1",
					"level34_f_message_extern1",
					"level34_f_message_extern2",
					"level34_f_message_inline_emtpy",
					"level34_f_message_inline_ignore_fields",
					"level34_f_string2",
					"level34_m_string2",
					"level34_one1_extern1",
					"level34_one3_string1",
					"level34_p_string2",
					"level34_r_string2",
				}
				values := data["level28_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level28_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level28_one1_extern1", key))
				}
			})

			// level02_one1_extern1
			t.Run("level02_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level06_c_f_string1",
					"level06_c_m_string1",
					"level06_c_p_string1",
					"level06_c_r_string1",
					"level06_f_message_extern1",
					"level06_f_message_extern2",
					"level06_f_message_inline_emtpy",
					"level06_f_message_inline_ignore_fields",
					"level06_f_string2",
					"level06_m_string2",
					"level06_one1_extern1",
					"level06_one3_string1",
					"level06_p_string2",
					"level06_r_string2",
					"level14_c_f_string1",
					"level14_c_m_string1",
					"level14_c_p_string1",
					"level14_c_r_string1",
					"level14_f_message_extern1",
					"level14_f_message_extern2",
					"level14_f_message_inline_emtpy",
					"level14_f_message_inline_ignore_fields",
					"level14_f_string2",
					"level14_m_string2",
					"level14_one1_extern1",
					"level14_one3_string1",
					"level14_p_string2",
					"level14_r_string2",
					"level15_c_f_string1",
					"level15_c_m_string1",
					"level15_c_one2_string1",
					"level15_c_p_string1",
					"level15_c_r_string1",
					"level15_f_message_extern1",
					"level15_f_message_extern2",
					"level15_f_message_inline_emtpy",
					"level15_f_message_inline_ignore_fields",
					"level15_f_string2",
					"level15_m_string2",
					"level15_one1_extern1",
					"level15_one3_string1",
					"level15_p_string2",
					"level15_r_string2",
					"level16_c_f_string1",
					"level16_c_m_string1",
					"level16_c_p_string1",
					"level16_c_r_string1",
					"level16_f_message_extern1",
					"level16_f_message_extern2",
					"level16_f_message_inline_emtpy",
					"level16_f_message_inline_ignore_fields",
					"level16_f_string2",
					"level16_m_string2",
					"level16_one1_extern1",
					"level16_one3_string1",
					"level16_p_string2",
					"level16_r_string2",
					"level22_c_f_string1",
					"level22_c_m_string1",
					"level22_c_p_string1",
					"level22_c_r_string1",
					"level22_f_message_extern1",
					"level22_f_message_extern2",
					"level22_f_message_inline_emtpy",
					"level22_f_message_inline_ignore_fields",
					"level22_f_string2",
					"level22_m_string2",
					"level22_one1_extern1",
					"level22_one3_string1",
					"level22_p_string2",
					"level22_r_string2",
					"level23_c_f_string1",
					"level23_c_m_string1",
					"level23_c_one2_string1",
					"level23_c_p_string1",
					"level23_c_r_string1",
					"level23_f_message_extern1",
					"level23_f_message_extern2",
					"level23_f_message_inline_emtpy",
					"level23_f_message_inline_ignore_fields",
					"level23_f_string2",
					"level23_m_string2",
					"level23_one1_extern1",
					"level23_one3_string1",
					"level23_p_string2",
					"level23_r_string2",
					"level24_c_f_string1",
					"level24_c_m_string1",
					"level24_c_one2_string1",
					"level24_c_p_string1",
					"level24_c_r_string1",
					"level24_f_message_extern1",
					"level24_f_message_extern2",
					"level24_f_message_inline_emtpy",
					"level24_f_message_inline_ignore_fields",
					"level24_f_string2",
					"level24_m_string2",
					"level24_one1_extern1",
					"level24_one3_string1",
					"level24_p_string2",
					"level24_r_string2",
					"level26_c_f_string1",
					"level26_c_m_string1",
					"level26_c_p_string1",
					"level26_c_r_string1",
					"level26_f_message_extern1",
					"level26_f_message_extern2",
					"level26_f_message_inline_emtpy",
					"level26_f_message_inline_ignore_fields",
					"level26_f_string2",
					"level26_m_string2",
					"level26_one1_extern1",
					"level26_one3_string1",
					"level26_p_string2",
					"level26_r_string2",
					"level27_c_f_string1",
					"level27_c_m_string1",
					"level27_c_one2_string1",
					"level27_c_p_string1",
					"level27_c_r_string1",
					"level27_f_message_extern1",
					"level27_f_message_extern2",
					"level27_f_message_inline_emtpy",
					"level27_f_message_inline_ignore_fields",
					"level27_f_string2",
					"level27_m_string2",
					"level27_one1_extern1",
					"level27_one3_string1",
					"level27_p_string2",
					"level27_r_string2",
					"level28_c_f_string1",
					"level28_c_m_string1",
					"level28_c_p_string1",
					"level28_c_r_string1",
					"level28_f_message_extern1",
					"level28_f_message_extern2",
					"level28_f_message_inline_emtpy",
					"level28_f_message_inline_ignore_fields",
					"level28_f_string2",
					"level28_m_string2",
					"level28_one1_extern1",
					"level28_one3_string1",
					"level28_p_string2",
					"level28_r_string2",
					"level30_c_f_string1",
					"level30_c_m_string1",
					"level30_c_one2_string1",
					"level30_c_p_string1",
					"level30_c_r_string1",
					"level30_f_message_extern1",
					"level30_f_message_extern2",
					"level30_f_message_inline_emtpy",
					"level30_f_message_inline_ignore_fields",
					"level30_f_string2",
					"level30_m_string2",
					"level30_one1_extern1",
					"level30_one3_string1",
					"level30_p_string2",
					"level30_r_string2",
					"level31_c_f_string1",
					"level31_c_m_string1",
					"level31_c_one2_string1",
					"level31_c_p_string1",
					"level31_c_r_string1",
					"level31_f_message_extern1",
					"level31_f_message_extern2",
					"level31_f_message_inline_emtpy",
					"level31_f_message_inline_ignore_fields",
					"level31_f_string2",
					"level31_m_string2",
					"level31_one1_extern1",
					"level31_one3_string1",
					"level31_p_string2",
					"level31_r_string2",
					"level32_c_f_string1",
					"level32_c_m_string1",
					"level32_c_one2_string1",
					"level32_c_p_string1",
					"level32_c_r_string1",
					"level32_f_message_extern1",
					"level32_f_message_extern2",
					"level32_f_message_inline_emtpy",
					"level32_f_message_inline_ignore_fields",
					"level32_f_string2",
					"level32_m_string2",
					"level32_one1_extern1",
					"level32_one3_string1",
					"level32_p_string2",
					"level32_r_string2",
					"level34_c_f_string1",
					"level34_c_m_string1",
					"level34_c_one2_string1",
					"level34_c_p_string1",
					"level34_c_r_string1",
					"level34_f_message_extern1",
					"level34_f_message_extern2",
					"level34_f_message_inline_emtpy",
					"level34_f_message_inline_ignore_fields",
					"level34_f_string2",
					"level34_m_string2",
					"level34_one1_extern1",
					"level34_one3_string1",
					"level34_p_string2",
					"level34_r_string2",
					"level35_c_f_string1",
					"level35_c_m_string1",
					"level35_c_one2_string1",
					"level35_c_p_string1",
					"level35_c_r_string1",
					"level35_f_message_extern1",
					"level35_f_message_extern2",
					"level35_f_message_inline_emtpy",
					"level35_f_message_inline_ignore_fields",
					"level35_f_string2",
					"level35_m_string2",
					"level35_one1_extern1",
					"level35_one3_string1",
					"level35_p_string2",
					"level35_r_string2",
					"level36_c_f_string1",
					"level36_c_m_string1",
					"level36_c_one2_string1",
					"level36_c_p_string1",
					"level36_c_r_string1",
					"level36_f_message_extern1",
					"level36_f_message_extern2",
					"level36_f_message_inline_emtpy",
					"level36_f_message_inline_ignore_fields",
					"level36_f_string2",
					"level36_m_string2",
					"level36_one1_extern1",
					"level36_one3_string1",
					"level36_p_string2",
					"level36_r_string2",
					"level38_c_f_string1",
					"level38_c_m_string1",
					"level38_c_one2_string1",
					"level38_c_p_string1",
					"level38_c_r_string1",
					"level38_f_message_extern1",
					"level38_f_message_extern2",
					"level38_f_message_inline_emtpy",
					"level38_f_message_inline_ignore_fields",
					"level38_f_string2",
					"level38_m_string2",
					"level38_one1_extern1",
					"level38_one3_string1",
					"level38_p_string2",
					"level38_r_string2",
					"level39_c_f_string1",
					"level39_c_m_string1",
					"level39_c_one2_string1",
					"level39_c_p_string1",
					"level39_c_r_string1",
					"level39_f_message_extern1",
					"level39_f_message_extern2",
					"level39_f_message_inline_emtpy",
					"level39_f_message_inline_ignore_fields",
					"level39_f_string2",
					"level39_m_string2",
					"level39_one1_extern1",
					"level39_one3_string1",
					"level39_p_string2",
					"level39_r_string2",
					"level40_c_f_string1",
					"level40_c_m_string1",
					"level40_c_one2_string1",
					"level40_c_p_string1",
					"level40_c_r_string1",
					"level40_f_message_extern1",
					"level40_f_message_extern2",
					"level40_f_message_inline_emtpy",
					"level40_f_message_inline_ignore_fields",
					"level40_f_string2",
					"level40_m_string2",
					"level40_one1_extern1",
					"level40_one3_string1",
					"level40_p_string2",
					"level40_r_string2",
				}
				values := data["level02_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level02_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level02_one1_extern1", key))
				}
			})

			// level08_one1_extern1
			t.Run("level08_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level18_c_f_string1",
					"level18_c_m_string1",
					"level18_c_one2_string1",
					"level18_c_p_string1",
					"level18_c_r_string1",
					"level18_f_message_extern1",
					"level18_f_message_extern2",
					"level18_f_message_inline_emtpy",
					"level18_f_message_inline_ignore_fields",
					"level18_f_string2",
					"level18_m_string2",
					"level18_one1_extern1",
					"level18_one3_string1",
					"level18_p_string2",
					"level18_r_string2",
				}
				values := data["level08_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level08_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level08_one1_extern1", key))
				}
			})

			// level01_one1_extern1
			t.Run("level01_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level02_c_f_string1",
					"level02_c_m_string1",
					"level02_c_p_string1",
					"level02_c_r_string1",
					"level02_f_message_extern1",
					"level02_f_message_extern2",
					"level02_f_message_inline_emtpy",
					"level02_f_message_inline_ignore_fields",
					"level02_f_string2",
					"level02_m_string2",
					"level02_one1_extern1",
					"level02_one3_string1",
					"level02_p_string2",
					"level02_r_string2",
					"level06_c_f_string1",
					"level06_c_m_string1",
					"level06_c_p_string1",
					"level06_c_r_string1",
					"level06_f_message_extern1",
					"level06_f_message_extern2",
					"level06_f_message_inline_emtpy",
					"level06_f_message_inline_ignore_fields",
					"level06_f_string2",
					"level06_m_string2",
					"level06_one1_extern1",
					"level06_one3_string1",
					"level06_p_string2",
					"level06_r_string2",
					"level07_c_f_string1",
					"level07_c_m_string1",
					"level07_c_one2_string1",
					"level07_c_p_string1",
					"level07_c_r_string1",
					"level07_f_message_extern1",
					"level07_f_message_extern2",
					"level07_f_message_inline_emtpy",
					"level07_f_message_inline_ignore_fields",
					"level07_f_string2",
					"level07_m_string2",
					"level07_one1_extern1",
					"level07_one3_string1",
					"level07_p_string2",
					"level07_r_string2",
					"level08_c_f_string1",
					"level08_c_m_string1",
					"level08_c_p_string1",
					"level08_c_r_string1",
					"level08_f_message_extern1",
					"level08_f_message_extern2",
					"level08_f_message_inline_emtpy",
					"level08_f_message_inline_ignore_fields",
					"level08_f_string2",
					"level08_m_string2",
					"level08_one1_extern1",
					"level08_one3_string1",
					"level08_p_string2",
					"level08_r_string2",
					"level14_c_f_string1",
					"level14_c_m_string1",
					"level14_c_p_string1",
					"level14_c_r_string1",
					"level14_f_message_extern1",
					"level14_f_message_extern2",
					"level14_f_message_inline_emtpy",
					"level14_f_message_inline_ignore_fields",
					"level14_f_string2",
					"level14_m_string2",
					"level14_one1_extern1",
					"level14_one3_string1",
					"level14_p_string2",
					"level14_r_string2",
					"level15_c_f_string1",
					"level15_c_m_string1",
					"level15_c_one2_string1",
					"level15_c_p_string1",
					"level15_c_r_string1",
					"level15_f_message_extern1",
					"level15_f_message_extern2",
					"level15_f_message_inline_emtpy",
					"level15_f_message_inline_ignore_fields",
					"level15_f_string2",
					"level15_m_string2",
					"level15_one1_extern1",
					"level15_one3_string1",
					"level15_p_string2",
					"level15_r_string2",
					"level16_c_f_string1",
					"level16_c_m_string1",
					"level16_c_p_string1",
					"level16_c_r_string1",
					"level16_f_message_extern1",
					"level16_f_message_extern2",
					"level16_f_message_inline_emtpy",
					"level16_f_message_inline_ignore_fields",
					"level16_f_string2",
					"level16_m_string2",
					"level16_one1_extern1",
					"level16_one3_string1",
					"level16_p_string2",
					"level16_r_string2",
					"level18_c_f_string1",
					"level18_c_m_string1",
					"level18_c_one2_string1",
					"level18_c_p_string1",
					"level18_c_r_string1",
					"level18_f_message_extern1",
					"level18_f_message_extern2",
					"level18_f_message_inline_emtpy",
					"level18_f_message_inline_ignore_fields",
					"level18_f_string2",
					"level18_m_string2",
					"level18_one1_extern1",
					"level18_one3_string1",
					"level18_p_string2",
					"level18_r_string2",
					"level19_c_f_string1",
					"level19_c_m_string1",
					"level19_c_one2_string1",
					"level19_c_p_string1",
					"level19_c_r_string1",
					"level19_f_message_extern1",
					"level19_f_message_extern2",
					"level19_f_message_inline_emtpy",
					"level19_f_message_inline_ignore_fields",
					"level19_f_string2",
					"level19_m_string2",
					"level19_one1_extern1",
					"level19_one3_string1",
					"level19_p_string2",
					"level19_r_string2",
					"level20_c_f_string1",
					"level20_c_m_string1",
					"level20_c_one2_string1",
					"level20_c_p_string1",
					"level20_c_r_string1",
					"level20_f_message_extern1",
					"level20_f_message_extern2",
					"level20_f_message_inline_emtpy",
					"level20_f_message_inline_ignore_fields",
					"level20_f_string2",
					"level20_m_string2",
					"level20_one1_extern1",
					"level20_one3_string1",
					"level20_p_string2",
					"level20_r_string2",
					"level22_c_f_string1",
					"level22_c_m_string1",
					"level22_c_p_string1",
					"level22_c_r_string1",
					"level22_f_message_extern1",
					"level22_f_message_extern2",
					"level22_f_message_inline_emtpy",
					"level22_f_message_inline_ignore_fields",
					"level22_f_string2",
					"level22_m_string2",
					"level22_one1_extern1",
					"level22_one3_string1",
					"level22_p_string2",
					"level22_r_string2",
					"level23_c_f_string1",
					"level23_c_m_string1",
					"level23_c_one2_string1",
					"level23_c_p_string1",
					"level23_c_r_string1",
					"level23_f_message_extern1",
					"level23_f_message_extern2",
					"level23_f_message_inline_emtpy",
					"level23_f_message_inline_ignore_fields",
					"level23_f_string2",
					"level23_m_string2",
					"level23_one1_extern1",
					"level23_one3_string1",
					"level23_p_string2",
					"level23_r_string2",
					"level24_c_f_string1",
					"level24_c_m_string1",
					"level24_c_one2_string1",
					"level24_c_p_string1",
					"level24_c_r_string1",
					"level24_f_message_extern1",
					"level24_f_message_extern2",
					"level24_f_message_inline_emtpy",
					"level24_f_message_inline_ignore_fields",
					"level24_f_string2",
					"level24_m_string2",
					"level24_one1_extern1",
					"level24_one3_string1",
					"level24_p_string2",
					"level24_r_string2",
					"level26_c_f_string1",
					"level26_c_m_string1",
					"level26_c_p_string1",
					"level26_c_r_string1",
					"level26_f_message_extern1",
					"level26_f_message_extern2",
					"level26_f_message_inline_emtpy",
					"level26_f_message_inline_ignore_fields",
					"level26_f_string2",
					"level26_m_string2",
					"level26_one1_extern1",
					"level26_one3_string1",
					"level26_p_string2",
					"level26_r_string2",
					"level27_c_f_string1",
					"level27_c_m_string1",
					"level27_c_one2_string1",
					"level27_c_p_string1",
					"level27_c_r_string1",
					"level27_f_message_extern1",
					"level27_f_message_extern2",
					"level27_f_message_inline_emtpy",
					"level27_f_message_inline_ignore_fields",
					"level27_f_string2",
					"level27_m_string2",
					"level27_one1_extern1",
					"level27_one3_string1",
					"level27_p_string2",
					"level27_r_string2",
					"level28_c_f_string1",
					"level28_c_m_string1",
					"level28_c_p_string1",
					"level28_c_r_string1",
					"level28_f_message_extern1",
					"level28_f_message_extern2",
					"level28_f_message_inline_emtpy",
					"level28_f_message_inline_ignore_fields",
					"level28_f_string2",
					"level28_m_string2",
					"level28_one1_extern1",
					"level28_one3_string1",
					"level28_p_string2",
					"level28_r_string2",
					"level30_c_f_string1",
					"level30_c_m_string1",
					"level30_c_one2_string1",
					"level30_c_p_string1",
					"level30_c_r_string1",
					"level30_f_message_extern1",
					"level30_f_message_extern2",
					"level30_f_message_inline_emtpy",
					"level30_f_message_inline_ignore_fields",
					"level30_f_string2",
					"level30_m_string2",
					"level30_one1_extern1",
					"level30_one3_string1",
					"level30_p_string2",
					"level30_r_string2",
					"level31_c_f_string1",
					"level31_c_m_string1",
					"level31_c_one2_string1",
					"level31_c_p_string1",
					"level31_c_r_string1",
					"level31_f_message_extern1",
					"level31_f_message_extern2",
					"level31_f_message_inline_emtpy",
					"level31_f_message_inline_ignore_fields",
					"level31_f_string2",
					"level31_m_string2",
					"level31_one1_extern1",
					"level31_one3_string1",
					"level31_p_string2",
					"level31_r_string2",
					"level32_c_f_string1",
					"level32_c_m_string1",
					"level32_c_one2_string1",
					"level32_c_p_string1",
					"level32_c_r_string1",
					"level32_f_message_extern1",
					"level32_f_message_extern2",
					"level32_f_message_inline_emtpy",
					"level32_f_message_inline_ignore_fields",
					"level32_f_string2",
					"level32_m_string2",
					"level32_one1_extern1",
					"level32_one3_string1",
					"level32_p_string2",
					"level32_r_string2",
					"level34_c_f_string1",
					"level34_c_m_string1",
					"level34_c_one2_string1",
					"level34_c_p_string1",
					"level34_c_r_string1",
					"level34_f_message_extern1",
					"level34_f_message_extern2",
					"level34_f_message_inline_emtpy",
					"level34_f_message_inline_ignore_fields",
					"level34_f_string2",
					"level34_m_string2",
					"level34_one1_extern1",
					"level34_one3_string1",
					"level34_p_string2",
					"level34_r_string2",
					"level35_c_f_string1",
					"level35_c_m_string1",
					"level35_c_one2_string1",
					"level35_c_p_string1",
					"level35_c_r_string1",
					"level35_f_message_extern1",
					"level35_f_message_extern2",
					"level35_f_message_inline_emtpy",
					"level35_f_message_inline_ignore_fields",
					"level35_f_string2",
					"level35_m_string2",
					"level35_one1_extern1",
					"level35_one3_string1",
					"level35_p_string2",
					"level35_r_string2",
					"level36_c_f_string1",
					"level36_c_m_string1",
					"level36_c_one2_string1",
					"level36_c_p_string1",
					"level36_c_r_string1",
					"level36_f_message_extern1",
					"level36_f_message_extern2",
					"level36_f_message_inline_emtpy",
					"level36_f_message_inline_ignore_fields",
					"level36_f_string2",
					"level36_m_string2",
					"level36_one1_extern1",
					"level36_one3_string1",
					"level36_p_string2",
					"level36_r_string2",
					"level38_c_f_string1",
					"level38_c_m_string1",
					"level38_c_one2_string1",
					"level38_c_p_string1",
					"level38_c_r_string1",
					"level38_f_message_extern1",
					"level38_f_message_extern2",
					"level38_f_message_inline_emtpy",
					"level38_f_message_inline_ignore_fields",
					"level38_f_string2",
					"level38_m_string2",
					"level38_one1_extern1",
					"level38_one3_string1",
					"level38_p_string2",
					"level38_r_string2",
					"level39_c_f_string1",
					"level39_c_m_string1",
					"level39_c_one2_string1",
					"level39_c_p_string1",
					"level39_c_r_string1",
					"level39_f_message_extern1",
					"level39_f_message_extern2",
					"level39_f_message_inline_emtpy",
					"level39_f_message_inline_ignore_fields",
					"level39_f_string2",
					"level39_m_string2",
					"level39_one1_extern1",
					"level39_one3_string1",
					"level39_p_string2",
					"level39_r_string2",
					"level40_c_f_string1",
					"level40_c_m_string1",
					"level40_c_one2_string1",
					"level40_c_p_string1",
					"level40_c_r_string1",
					"level40_f_message_extern1",
					"level40_f_message_extern2",
					"level40_f_message_inline_emtpy",
					"level40_f_message_inline_ignore_fields",
					"level40_f_string2",
					"level40_m_string2",
					"level40_one1_extern1",
					"level40_one3_string1",
					"level40_p_string2",
					"level40_r_string2",
				}
				values := data["level01_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level01_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level01_one1_extern1", key))
				}
			})

			// level04_one1_extern1
			t.Run("level04_one1_extern1", func(t *testing.T) {
				keys := []string{
					"level10_c_f_string1",
					"level10_c_m_string1",
					"level10_c_one2_string1",
					"level10_c_p_string1",
					"level10_c_r_string1",
					"level10_f_message_extern1",
					"level10_f_message_extern2",
					"level10_f_message_inline_emtpy",
					"level10_f_message_inline_ignore_fields",
					"level10_f_string2",
					"level10_m_string2",
					"level10_one1_extern1",
					"level10_one3_string1",
					"level10_p_string2",
					"level10_r_string2",
				}
				values := data["level04_one1_extern1"].(map[string]interface{})
				for _, key := range keys {
					val, ok := values[key]
					require.True(
						t, ok, fmt.Sprintf("the json key [%s] not at in level 2 with level04_one1_extern1", key))
					require.NotNil(
						t, val,
						fmt.Sprintf("the value of json key [%s] is nil in level 2 with level04_one1_extern1", key))
				}
			})
		})

		t.Run("Level3", func(t *testing.T) {
			// level14_one1_extern1
			t.Run("level14_one1_extern1", func(t *testing.T) {
				// level14_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})

			// level06_one1_extern1
			t.Run("level06_one1_extern1", func(t *testing.T) {
				// level06_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level06_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level06_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level06_one1_extern1.level22_one1_extern1", key))
					}
				})

				// level06_one1_extern1.level14_one1_extern1
				t.Run("level14_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level06_one1_extern1.level14_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level06_one1_extern1.level14_one1_extern1", key))
					}
				})
			})

			// level16_one1_extern1
			t.Run("level16_one1_extern1", func(t *testing.T) {
				// level16_one1_extern1.level26_one1_extern1
				t.Run("level26_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
					}
					values := data["level16_one1_extern1"].(map[string]interface{})["level26_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level16_one1_extern1.level26_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level16_one1_extern1.level26_one1_extern1", key))
					}
				})
			})

			// level02_one1_extern1
			t.Run("level02_one1_extern1", func(t *testing.T) {
				// level02_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level02_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level02_one1_extern1.level22_one1_extern1", key))
					}
				})

				// level02_one1_extern1.level14_one1_extern1
				t.Run("level14_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level02_one1_extern1.level14_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level02_one1_extern1.level14_one1_extern1", key))
					}
				})

				// level02_one1_extern1.level06_one1_extern1
				t.Run("level06_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level14_c_f_string1",
						"level14_c_m_string1",
						"level14_c_p_string1",
						"level14_c_r_string1",
						"level14_f_message_extern1",
						"level14_f_message_extern2",
						"level14_f_message_inline_emtpy",
						"level14_f_message_inline_ignore_fields",
						"level14_f_string2",
						"level14_m_string2",
						"level14_one1_extern1",
						"level14_one3_string1",
						"level14_p_string2",
						"level14_r_string2",
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level23_c_f_string1",
						"level23_c_m_string1",
						"level23_c_one2_string1",
						"level23_c_p_string1",
						"level23_c_r_string1",
						"level23_f_message_extern1",
						"level23_f_message_extern2",
						"level23_f_message_inline_emtpy",
						"level23_f_message_inline_ignore_fields",
						"level23_f_string2",
						"level23_m_string2",
						"level23_one1_extern1",
						"level23_one3_string1",
						"level23_p_string2",
						"level23_r_string2",
						"level24_c_f_string1",
						"level24_c_m_string1",
						"level24_c_one2_string1",
						"level24_c_p_string1",
						"level24_c_r_string1",
						"level24_f_message_extern1",
						"level24_f_message_extern2",
						"level24_f_message_inline_emtpy",
						"level24_f_message_inline_ignore_fields",
						"level24_f_string2",
						"level24_m_string2",
						"level24_one1_extern1",
						"level24_one3_string1",
						"level24_p_string2",
						"level24_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level02_one1_extern1.level06_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level02_one1_extern1.level06_one1_extern1", key))
					}
				})

				// level02_one1_extern1.level26_one1_extern1
				t.Run("level26_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level26_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level02_one1_extern1.level26_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level02_one1_extern1.level26_one1_extern1", key))
					}
				})

				// level02_one1_extern1.level16_one1_extern1
				t.Run("level16_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level26_c_f_string1",
						"level26_c_m_string1",
						"level26_c_p_string1",
						"level26_c_r_string1",
						"level26_f_message_extern1",
						"level26_f_message_extern2",
						"level26_f_message_inline_emtpy",
						"level26_f_message_inline_ignore_fields",
						"level26_f_string2",
						"level26_m_string2",
						"level26_one1_extern1",
						"level26_one3_string1",
						"level26_p_string2",
						"level26_r_string2",
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
						"level31_c_f_string1",
						"level31_c_m_string1",
						"level31_c_one2_string1",
						"level31_c_p_string1",
						"level31_c_r_string1",
						"level31_f_message_extern1",
						"level31_f_message_extern2",
						"level31_f_message_inline_emtpy",
						"level31_f_message_inline_ignore_fields",
						"level31_f_string2",
						"level31_m_string2",
						"level31_one1_extern1",
						"level31_one3_string1",
						"level31_p_string2",
						"level31_r_string2",
						"level32_c_f_string1",
						"level32_c_m_string1",
						"level32_c_one2_string1",
						"level32_c_p_string1",
						"level32_c_r_string1",
						"level32_f_message_extern1",
						"level32_f_message_extern2",
						"level32_f_message_inline_emtpy",
						"level32_f_message_inline_ignore_fields",
						"level32_f_string2",
						"level32_m_string2",
						"level32_one1_extern1",
						"level32_one3_string1",
						"level32_p_string2",
						"level32_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level16_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level02_one1_extern1.level16_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level02_one1_extern1.level16_one1_extern1", key))
					}
				})

				// level02_one1_extern1.level28_one1_extern1
				t.Run("level28_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level34_c_f_string1",
						"level34_c_m_string1",
						"level34_c_one2_string1",
						"level34_c_p_string1",
						"level34_c_r_string1",
						"level34_f_message_extern1",
						"level34_f_message_extern2",
						"level34_f_message_inline_emtpy",
						"level34_f_message_inline_ignore_fields",
						"level34_f_string2",
						"level34_m_string2",
						"level34_one1_extern1",
						"level34_one3_string1",
						"level34_p_string2",
						"level34_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level28_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level02_one1_extern1.level28_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level02_one1_extern1.level28_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1
			t.Run("level01_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level22_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level14_one1_extern1
				t.Run("level14_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level14_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level14_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level06_one1_extern1
				t.Run("level06_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level14_c_f_string1",
						"level14_c_m_string1",
						"level14_c_p_string1",
						"level14_c_r_string1",
						"level14_f_message_extern1",
						"level14_f_message_extern2",
						"level14_f_message_inline_emtpy",
						"level14_f_message_inline_ignore_fields",
						"level14_f_string2",
						"level14_m_string2",
						"level14_one1_extern1",
						"level14_one3_string1",
						"level14_p_string2",
						"level14_r_string2",
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level23_c_f_string1",
						"level23_c_m_string1",
						"level23_c_one2_string1",
						"level23_c_p_string1",
						"level23_c_r_string1",
						"level23_f_message_extern1",
						"level23_f_message_extern2",
						"level23_f_message_inline_emtpy",
						"level23_f_message_inline_ignore_fields",
						"level23_f_string2",
						"level23_m_string2",
						"level23_one1_extern1",
						"level23_one3_string1",
						"level23_p_string2",
						"level23_r_string2",
						"level24_c_f_string1",
						"level24_c_m_string1",
						"level24_c_one2_string1",
						"level24_c_p_string1",
						"level24_c_r_string1",
						"level24_f_message_extern1",
						"level24_f_message_extern2",
						"level24_f_message_inline_emtpy",
						"level24_f_message_inline_ignore_fields",
						"level24_f_string2",
						"level24_m_string2",
						"level24_one1_extern1",
						"level24_one3_string1",
						"level24_p_string2",
						"level24_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level06_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level06_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level26_one1_extern1
				t.Run("level26_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level26_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level26_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level26_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level16_one1_extern1
				t.Run("level16_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level26_c_f_string1",
						"level26_c_m_string1",
						"level26_c_p_string1",
						"level26_c_r_string1",
						"level26_f_message_extern1",
						"level26_f_message_extern2",
						"level26_f_message_inline_emtpy",
						"level26_f_message_inline_ignore_fields",
						"level26_f_string2",
						"level26_m_string2",
						"level26_one1_extern1",
						"level26_one3_string1",
						"level26_p_string2",
						"level26_r_string2",
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
						"level31_c_f_string1",
						"level31_c_m_string1",
						"level31_c_one2_string1",
						"level31_c_p_string1",
						"level31_c_r_string1",
						"level31_f_message_extern1",
						"level31_f_message_extern2",
						"level31_f_message_inline_emtpy",
						"level31_f_message_inline_ignore_fields",
						"level31_f_string2",
						"level31_m_string2",
						"level31_one1_extern1",
						"level31_one3_string1",
						"level31_p_string2",
						"level31_r_string2",
						"level32_c_f_string1",
						"level32_c_m_string1",
						"level32_c_one2_string1",
						"level32_c_p_string1",
						"level32_c_r_string1",
						"level32_f_message_extern1",
						"level32_f_message_extern2",
						"level32_f_message_inline_emtpy",
						"level32_f_message_inline_ignore_fields",
						"level32_f_string2",
						"level32_m_string2",
						"level32_one1_extern1",
						"level32_one3_string1",
						"level32_p_string2",
						"level32_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level16_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level16_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level16_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level28_one1_extern1
				t.Run("level28_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level34_c_f_string1",
						"level34_c_m_string1",
						"level34_c_one2_string1",
						"level34_c_p_string1",
						"level34_c_r_string1",
						"level34_f_message_extern1",
						"level34_f_message_extern2",
						"level34_f_message_inline_emtpy",
						"level34_f_message_inline_ignore_fields",
						"level34_f_string2",
						"level34_m_string2",
						"level34_one1_extern1",
						"level34_one3_string1",
						"level34_p_string2",
						"level34_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level28_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level28_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level28_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level02_one1_extern1
				t.Run("level02_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level06_c_f_string1",
						"level06_c_m_string1",
						"level06_c_p_string1",
						"level06_c_r_string1",
						"level06_f_message_extern1",
						"level06_f_message_extern2",
						"level06_f_message_inline_emtpy",
						"level06_f_message_inline_ignore_fields",
						"level06_f_string2",
						"level06_m_string2",
						"level06_one1_extern1",
						"level06_one3_string1",
						"level06_p_string2",
						"level06_r_string2",
						"level14_c_f_string1",
						"level14_c_m_string1",
						"level14_c_p_string1",
						"level14_c_r_string1",
						"level14_f_message_extern1",
						"level14_f_message_extern2",
						"level14_f_message_inline_emtpy",
						"level14_f_message_inline_ignore_fields",
						"level14_f_string2",
						"level14_m_string2",
						"level14_one1_extern1",
						"level14_one3_string1",
						"level14_p_string2",
						"level14_r_string2",
						"level15_c_f_string1",
						"level15_c_m_string1",
						"level15_c_one2_string1",
						"level15_c_p_string1",
						"level15_c_r_string1",
						"level15_f_message_extern1",
						"level15_f_message_extern2",
						"level15_f_message_inline_emtpy",
						"level15_f_message_inline_ignore_fields",
						"level15_f_string2",
						"level15_m_string2",
						"level15_one1_extern1",
						"level15_one3_string1",
						"level15_p_string2",
						"level15_r_string2",
						"level16_c_f_string1",
						"level16_c_m_string1",
						"level16_c_p_string1",
						"level16_c_r_string1",
						"level16_f_message_extern1",
						"level16_f_message_extern2",
						"level16_f_message_inline_emtpy",
						"level16_f_message_inline_ignore_fields",
						"level16_f_string2",
						"level16_m_string2",
						"level16_one1_extern1",
						"level16_one3_string1",
						"level16_p_string2",
						"level16_r_string2",
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level23_c_f_string1",
						"level23_c_m_string1",
						"level23_c_one2_string1",
						"level23_c_p_string1",
						"level23_c_r_string1",
						"level23_f_message_extern1",
						"level23_f_message_extern2",
						"level23_f_message_inline_emtpy",
						"level23_f_message_inline_ignore_fields",
						"level23_f_string2",
						"level23_m_string2",
						"level23_one1_extern1",
						"level23_one3_string1",
						"level23_p_string2",
						"level23_r_string2",
						"level24_c_f_string1",
						"level24_c_m_string1",
						"level24_c_one2_string1",
						"level24_c_p_string1",
						"level24_c_r_string1",
						"level24_f_message_extern1",
						"level24_f_message_extern2",
						"level24_f_message_inline_emtpy",
						"level24_f_message_inline_ignore_fields",
						"level24_f_string2",
						"level24_m_string2",
						"level24_one1_extern1",
						"level24_one3_string1",
						"level24_p_string2",
						"level24_r_string2",
						"level26_c_f_string1",
						"level26_c_m_string1",
						"level26_c_p_string1",
						"level26_c_r_string1",
						"level26_f_message_extern1",
						"level26_f_message_extern2",
						"level26_f_message_inline_emtpy",
						"level26_f_message_inline_ignore_fields",
						"level26_f_string2",
						"level26_m_string2",
						"level26_one1_extern1",
						"level26_one3_string1",
						"level26_p_string2",
						"level26_r_string2",
						"level27_c_f_string1",
						"level27_c_m_string1",
						"level27_c_one2_string1",
						"level27_c_p_string1",
						"level27_c_r_string1",
						"level27_f_message_extern1",
						"level27_f_message_extern2",
						"level27_f_message_inline_emtpy",
						"level27_f_message_inline_ignore_fields",
						"level27_f_string2",
						"level27_m_string2",
						"level27_one1_extern1",
						"level27_one3_string1",
						"level27_p_string2",
						"level27_r_string2",
						"level28_c_f_string1",
						"level28_c_m_string1",
						"level28_c_p_string1",
						"level28_c_r_string1",
						"level28_f_message_extern1",
						"level28_f_message_extern2",
						"level28_f_message_inline_emtpy",
						"level28_f_message_inline_ignore_fields",
						"level28_f_string2",
						"level28_m_string2",
						"level28_one1_extern1",
						"level28_one3_string1",
						"level28_p_string2",
						"level28_r_string2",
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
						"level31_c_f_string1",
						"level31_c_m_string1",
						"level31_c_one2_string1",
						"level31_c_p_string1",
						"level31_c_r_string1",
						"level31_f_message_extern1",
						"level31_f_message_extern2",
						"level31_f_message_inline_emtpy",
						"level31_f_message_inline_ignore_fields",
						"level31_f_string2",
						"level31_m_string2",
						"level31_one1_extern1",
						"level31_one3_string1",
						"level31_p_string2",
						"level31_r_string2",
						"level32_c_f_string1",
						"level32_c_m_string1",
						"level32_c_one2_string1",
						"level32_c_p_string1",
						"level32_c_r_string1",
						"level32_f_message_extern1",
						"level32_f_message_extern2",
						"level32_f_message_inline_emtpy",
						"level32_f_message_inline_ignore_fields",
						"level32_f_string2",
						"level32_m_string2",
						"level32_one1_extern1",
						"level32_one3_string1",
						"level32_p_string2",
						"level32_r_string2",
						"level34_c_f_string1",
						"level34_c_m_string1",
						"level34_c_one2_string1",
						"level34_c_p_string1",
						"level34_c_r_string1",
						"level34_f_message_extern1",
						"level34_f_message_extern2",
						"level34_f_message_inline_emtpy",
						"level34_f_message_inline_ignore_fields",
						"level34_f_string2",
						"level34_m_string2",
						"level34_one1_extern1",
						"level34_one3_string1",
						"level34_p_string2",
						"level34_r_string2",
						"level35_c_f_string1",
						"level35_c_m_string1",
						"level35_c_one2_string1",
						"level35_c_p_string1",
						"level35_c_r_string1",
						"level35_f_message_extern1",
						"level35_f_message_extern2",
						"level35_f_message_inline_emtpy",
						"level35_f_message_inline_ignore_fields",
						"level35_f_string2",
						"level35_m_string2",
						"level35_one1_extern1",
						"level35_one3_string1",
						"level35_p_string2",
						"level35_r_string2",
						"level36_c_f_string1",
						"level36_c_m_string1",
						"level36_c_one2_string1",
						"level36_c_p_string1",
						"level36_c_r_string1",
						"level36_f_message_extern1",
						"level36_f_message_extern2",
						"level36_f_message_inline_emtpy",
						"level36_f_message_inline_ignore_fields",
						"level36_f_string2",
						"level36_m_string2",
						"level36_one1_extern1",
						"level36_one3_string1",
						"level36_p_string2",
						"level36_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level02_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level02_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level08_one1_extern1
				t.Run("level08_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level18_c_f_string1",
						"level18_c_m_string1",
						"level18_c_one2_string1",
						"level18_c_p_string1",
						"level18_c_r_string1",
						"level18_f_message_extern1",
						"level18_f_message_extern2",
						"level18_f_message_inline_emtpy",
						"level18_f_message_inline_ignore_fields",
						"level18_f_string2",
						"level18_m_string2",
						"level18_one1_extern1",
						"level18_one3_string1",
						"level18_p_string2",
						"level18_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level08_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 3 with level01_one1_extern1.level08_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 3 with level01_one1_extern1.level08_one1_extern1", key))
					}
				})
			})
		})

		t.Run("Level4", func(t *testing.T) {
			// level06_one1_extern1.level14_one1_extern1
			t.Run("level06_one1_extern1.level14_one1_extern1", func(t *testing.T) {
				// level06_one1_extern1.level14_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})

			// level02_one1_extern1.level14_one1_extern1
			t.Run("level02_one1_extern1.level14_one1_extern1", func(t *testing.T) {
				// level02_one1_extern1.level14_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level02_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level02_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})

			// level02_one1_extern1.level06_one1_extern1
			t.Run("level02_one1_extern1.level06_one1_extern1", func(t *testing.T) {
				// level02_one1_extern1.level06_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level02_one1_extern1.level06_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level02_one1_extern1.level06_one1_extern1.level22_one1_extern1", key))
					}
				})

				// level02_one1_extern1.level06_one1_extern1.level14_one1_extern1
				t.Run("level14_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level02_one1_extern1.level06_one1_extern1.level14_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level02_one1_extern1.level06_one1_extern1.level14_one1_extern1", key))
					}
				})
			})

			// level02_one1_extern1.level16_one1_extern1
			t.Run("level02_one1_extern1.level16_one1_extern1", func(t *testing.T) {
				// level02_one1_extern1.level16_one1_extern1.level26_one1_extern1
				t.Run("level26_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level16_one1_extern1"].(map[string]interface{})["level26_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level02_one1_extern1.level16_one1_extern1.level26_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level02_one1_extern1.level16_one1_extern1.level26_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level14_one1_extern1
			t.Run("level01_one1_extern1.level14_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level14_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level06_one1_extern1
			t.Run("level01_one1_extern1.level06_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level06_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level06_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level06_one1_extern1.level22_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level06_one1_extern1.level14_one1_extern1
				t.Run("level14_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level06_one1_extern1.level14_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level06_one1_extern1.level14_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level16_one1_extern1
			t.Run("level01_one1_extern1.level16_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level16_one1_extern1.level26_one1_extern1
				t.Run("level26_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level16_one1_extern1"].(map[string]interface{})["level26_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level16_one1_extern1.level26_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level16_one1_extern1.level26_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level02_one1_extern1
			t.Run("level01_one1_extern1.level02_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level02_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level02_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level02_one1_extern1.level22_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level02_one1_extern1.level14_one1_extern1
				t.Run("level14_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level02_one1_extern1.level14_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level02_one1_extern1.level14_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level02_one1_extern1.level06_one1_extern1
				t.Run("level06_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level14_c_f_string1",
						"level14_c_m_string1",
						"level14_c_p_string1",
						"level14_c_r_string1",
						"level14_f_message_extern1",
						"level14_f_message_extern2",
						"level14_f_message_inline_emtpy",
						"level14_f_message_inline_ignore_fields",
						"level14_f_string2",
						"level14_m_string2",
						"level14_one1_extern1",
						"level14_one3_string1",
						"level14_p_string2",
						"level14_r_string2",
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level23_c_f_string1",
						"level23_c_m_string1",
						"level23_c_one2_string1",
						"level23_c_p_string1",
						"level23_c_r_string1",
						"level23_f_message_extern1",
						"level23_f_message_extern2",
						"level23_f_message_inline_emtpy",
						"level23_f_message_inline_ignore_fields",
						"level23_f_string2",
						"level23_m_string2",
						"level23_one1_extern1",
						"level23_one3_string1",
						"level23_p_string2",
						"level23_r_string2",
						"level24_c_f_string1",
						"level24_c_m_string1",
						"level24_c_one2_string1",
						"level24_c_p_string1",
						"level24_c_r_string1",
						"level24_f_message_extern1",
						"level24_f_message_extern2",
						"level24_f_message_inline_emtpy",
						"level24_f_message_inline_ignore_fields",
						"level24_f_string2",
						"level24_m_string2",
						"level24_one1_extern1",
						"level24_one3_string1",
						"level24_p_string2",
						"level24_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level02_one1_extern1.level26_one1_extern1
				t.Run("level26_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level26_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level02_one1_extern1.level26_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level02_one1_extern1.level26_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level02_one1_extern1.level16_one1_extern1
				t.Run("level16_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level26_c_f_string1",
						"level26_c_m_string1",
						"level26_c_p_string1",
						"level26_c_r_string1",
						"level26_f_message_extern1",
						"level26_f_message_extern2",
						"level26_f_message_inline_emtpy",
						"level26_f_message_inline_ignore_fields",
						"level26_f_string2",
						"level26_m_string2",
						"level26_one1_extern1",
						"level26_one3_string1",
						"level26_p_string2",
						"level26_r_string2",
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
						"level31_c_f_string1",
						"level31_c_m_string1",
						"level31_c_one2_string1",
						"level31_c_p_string1",
						"level31_c_r_string1",
						"level31_f_message_extern1",
						"level31_f_message_extern2",
						"level31_f_message_inline_emtpy",
						"level31_f_message_inline_ignore_fields",
						"level31_f_string2",
						"level31_m_string2",
						"level31_one1_extern1",
						"level31_one3_string1",
						"level31_p_string2",
						"level31_r_string2",
						"level32_c_f_string1",
						"level32_c_m_string1",
						"level32_c_one2_string1",
						"level32_c_p_string1",
						"level32_c_r_string1",
						"level32_f_message_extern1",
						"level32_f_message_extern2",
						"level32_f_message_inline_emtpy",
						"level32_f_message_inline_ignore_fields",
						"level32_f_string2",
						"level32_m_string2",
						"level32_one1_extern1",
						"level32_one3_string1",
						"level32_p_string2",
						"level32_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level16_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level02_one1_extern1.level16_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level02_one1_extern1.level16_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level02_one1_extern1.level28_one1_extern1
				t.Run("level28_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level34_c_f_string1",
						"level34_c_m_string1",
						"level34_c_one2_string1",
						"level34_c_p_string1",
						"level34_c_r_string1",
						"level34_f_message_extern1",
						"level34_f_message_extern2",
						"level34_f_message_inline_emtpy",
						"level34_f_message_inline_ignore_fields",
						"level34_f_string2",
						"level34_m_string2",
						"level34_one1_extern1",
						"level34_one3_string1",
						"level34_p_string2",
						"level34_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level28_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 4 with level01_one1_extern1.level02_one1_extern1.level28_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 4 with level01_one1_extern1.level02_one1_extern1.level28_one1_extern1", key))
					}
				})
			})
		})

		t.Run("Level5", func(t *testing.T) {
			// level02_one1_extern1.level06_one1_extern1.level14_one1_extern1
			t.Run("level02_one1_extern1.level06_one1_extern1.level14_one1_extern1", func(t *testing.T) {
				// level02_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 5 with level02_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 5 with level02_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level06_one1_extern1.level14_one1_extern1
			t.Run("level01_one1_extern1.level06_one1_extern1.level14_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 5 with level02_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 5 with level02_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level02_one1_extern1.level14_one1_extern1
			t.Run("level01_one1_extern1.level02_one1_extern1.level14_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level02_one1_extern1.level14_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 5 with level01_one1_extern1.level02_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 5 with level01_one1_extern1.level02_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level02_one1_extern1.level06_one1_extern1
			t.Run("level01_one1_extern1.level02_one1_extern1.level06_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level22_one1_extern1
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 5 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 5 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level22_one1_extern1", key))
					}
				})

				// level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level14_one1_extern1
				t.Run("level14_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level22_c_f_string1",
						"level22_c_m_string1",
						"level22_c_p_string1",
						"level22_c_r_string1",
						"level22_f_message_extern1",
						"level22_f_message_extern2",
						"level22_f_message_inline_emtpy",
						"level22_f_message_inline_ignore_fields",
						"level22_f_string2",
						"level22_m_string2",
						"level22_one1_extern1",
						"level22_one3_string1",
						"level22_p_string2",
						"level22_r_string2",
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
						"level39_c_f_string1",
						"level39_c_m_string1",
						"level39_c_one2_string1",
						"level39_c_p_string1",
						"level39_c_r_string1",
						"level39_f_message_extern1",
						"level39_f_message_extern2",
						"level39_f_message_inline_emtpy",
						"level39_f_message_inline_ignore_fields",
						"level39_f_string2",
						"level39_m_string2",
						"level39_one1_extern1",
						"level39_one3_string1",
						"level39_p_string2",
						"level39_r_string2",
						"level40_c_f_string1",
						"level40_c_m_string1",
						"level40_c_one2_string1",
						"level40_c_p_string1",
						"level40_c_r_string1",
						"level40_f_message_extern1",
						"level40_f_message_extern2",
						"level40_f_message_inline_emtpy",
						"level40_f_message_inline_ignore_fields",
						"level40_f_string2",
						"level40_m_string2",
						"level40_one1_extern1",
						"level40_one3_string1",
						"level40_p_string2",
						"level40_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 5 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level14_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 5 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level14_one1_extern1", key))
					}
				})
			})

			// level01_one1_extern1.level02_one1_extern1.level16_one1_extern1
			t.Run("level01_one1_extern1.level02_one1_extern1.level16_one1_extern1", func(t *testing.T) {
				// level01_one1_extern1.level02_one1_extern1.level16_one1_extern1.level26_one1_extern1
				t.Run("level26_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level30_c_f_string1",
						"level30_c_m_string1",
						"level30_c_one2_string1",
						"level30_c_p_string1",
						"level30_c_r_string1",
						"level30_f_message_extern1",
						"level30_f_message_extern2",
						"level30_f_message_inline_emtpy",
						"level30_f_message_inline_ignore_fields",
						"level30_f_string2",
						"level30_m_string2",
						"level30_one1_extern1",
						"level30_one3_string1",
						"level30_p_string2",
						"level30_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level16_one1_extern1"].(map[string]interface{})["level26_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 5 with level01_one1_extern1.level02_one1_extern1.level16_one1_extern1.level26_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 5 with level01_one1_extern1.level02_one1_extern1.level16_one1_extern1.level26_one1_extern1", key))
					}
				})
			})

		})

		t.Run("Level6", func(t *testing.T) {
			// level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level14_one1_extern1
			t.Run("level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level14_one1_extern1", func(t *testing.T) {
				t.Run("level22_one1_extern1", func(t *testing.T) {
					keys := []string{
						"level38_c_f_string1",
						"level38_c_m_string1",
						"level38_c_one2_string1",
						"level38_c_p_string1",
						"level38_c_r_string1",
						"level38_f_message_extern1",
						"level38_f_message_extern2",
						"level38_f_message_inline_emtpy",
						"level38_f_message_inline_ignore_fields",
						"level38_f_string2",
						"level38_m_string2",
						"level38_one1_extern1",
						"level38_one3_string1",
						"level38_p_string2",
						"level38_r_string2",
					}
					values := data["level01_one1_extern1"].(map[string]interface{})["level02_one1_extern1"].(map[string]interface{})["level06_one1_extern1"].(map[string]interface{})["level14_one1_extern1"].(map[string]interface{})["level22_one1_extern1"].(map[string]interface{})
					for _, key := range keys {
						val, ok := values[key]
						require.True(
							t, ok, fmt.Sprintf("the json key [%s] not at in level 5 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
						require.NotNil(
							t, val,
							fmt.Sprintf("the value of json key [%s] is nil in level 5 with level01_one1_extern1.level02_one1_extern1.level06_one1_extern1.level14_one1_extern1.level22_one1_extern1", key))
					}
				})
			})
		})
	})
}

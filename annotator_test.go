package stanfordcorenlp

import (
	"bytes"
	"testing"
)

func TestAnnotatorType_String(t *testing.T) {
	for _, testcase := range []struct {
		annotators AnnotatorType
		expected   string
	}{
		{AnnotatorTokenize, "tokenize"},
		{AnnotatorQuote, "quote"},
		{AnnotatorTokenize | AnnotatorSsplit | AnnotatorPos, "tokenize,ssplit,pos"},
		{0, ""},
		{1 << 30, ""},
	} {
		if testcase.expected != testcase.annotators.String() {
			t.Errorf("Not equal: \n"+
				"expected: %s\n"+
				"actual  : %s", testcase.expected, testcase.annotators.String())
		}
	}
}

func TestAnnotatorType_MarshalJSON(t *testing.T) {
	for _, testcase := range []struct {
		annotators AnnotatorType
		expected   []byte
	}{
		{AnnotatorTokenize, []byte(`"tokenize"`)},
		{AnnotatorQuote, []byte(`"quote"`)},
		{AnnotatorTokenize | AnnotatorSsplit | AnnotatorPos, []byte(`"tokenize,ssplit,pos"`)},
		{0, []byte(`""`)},
		{1 << 30, []byte(`""`)},
	} {
		b, err := testcase.annotators.MarshalJSON()
		if err != nil {
			t.Errorf("Received unexpected error: %+v", err)
			t.FailNow()
		}
		if !bytes.Equal(testcase.expected, b) {
			t.Errorf("Not equal: \n"+
				"expected: %s\n"+
				"actual  : %s", testcase.expected, b)
		}
	}
}

package stanfordcorenlp

import (
	"fmt"
	"strings"
)

type AnnotatorType int

const (
	AnnotatorTokenize AnnotatorType = 1 << iota
	AnnotatorCleanxml
	AnnotatorSsplit
	AnnotatorDocdate
	AnnotatorPos
	AnnotatorLemma
	AnnotatorNer
	AnnotatorRegexner
	AnnotatorSentiment
	AnnotatorParse
	AnnotatorDepparse
	AnnotatorDcoref
	AnnotatorRelation
	AnnotatorNatlog
	AnnotatorEntitylink
	AnnotatorKbp
	AnnotatorQuote
)

var (
	annotators = []AnnotatorType{
		AnnotatorTokenize,
		AnnotatorCleanxml,
		AnnotatorSsplit,
		AnnotatorDocdate,
		AnnotatorPos,
		AnnotatorLemma,
		AnnotatorNer,
		AnnotatorRegexner,
		AnnotatorSentiment,
		AnnotatorParse,
		AnnotatorDepparse,
		AnnotatorDcoref,
		AnnotatorRelation,
		AnnotatorNatlog,
		AnnotatorEntitylink,
		AnnotatorKbp,
		AnnotatorQuote,
	}

	annotatorStrings = []string{
		"tokenize",
		"cleanxml",
		"ssplit",
		"docdate",
		"pos",
		"lemma",
		"ner",
		"regexner",
		"sentiment",
		"parse",
		"depparse",
		"dcoref",
		"relation",
		"natlog",
		"entitylink",
		"kbp",
		"quote",
	}
)

func (a AnnotatorType) String() string {
	m := a
	if m == 0 {
		return ""
	}
	var out []string
	for i, x := range annotators {
		if x > m {
			break
		}
		if (m & x) != 0 {
			out = append(out, annotatorStrings[i])
			m = m ^ x
		}
	}
	if m != 0 {
		out = append(out, "")
	}
	return strings.Join(out, ",")
}

func (a AnnotatorType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", a.String())), nil
}

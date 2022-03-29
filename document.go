package stanfordcorenlp

import (
	"encoding/json"
	"strings"
)

type Document struct {
	Sentences []*Sentence `json:"sentences"`
}

func UnmarshalDocument(data []byte) (*Document, error) {
	var d Document
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (d *Document) String() string {
	strs := make([]string, len(d.Sentences))
	for i, s := range d.Sentences {
		strs[i] = s.String()
	}
	return strings.Join(strs, " ")
}

type Sentence struct {
	Index                        int               `json:"index"`
	Tokens                       []*Token          `json:"tokens"`
	Parse                        string            `json:"parse,omitempty"`
	BasicDependencies            []*DependencyNode `json:"basicDependencies,omitempty"`
	EnhancedDependencies         []*DependencyNode `json:"enhancedDependencies,omitempty"`
	EnhancedPlusPlusDependencies []*DependencyNode `json:"enhancedPlusPlusDependencies,omitempty"`
}

func UnmarshalSentence(data []byte) (*Sentence, error) {
	var s Sentence
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *Sentence) String() string {
	strs := make([]string, len(s.Tokens))
	for i, t := range s.Tokens {
		strs[i] = t.String()
	}
	return strings.Join(strs, " ")
}

type Token struct {
	Index                int    `json:"index"`
	Word                 string `json:"word"`
	OriginalText         string `json:"originalText,omitempty"`
	Lemma                string `json:"lemma,omitempty"`
	CharacterOffsetBegin int    `json:"characterOffsetBegin,omitempty"`
	CharacterOffsetEnd   int    `json:"characterOffsetEnd,omitempty"`
	Pos                  string `json:"pos,omitempty"`
	Before               string `json:"before,omitempty"`
	After                string `json:"after,omitempty"`
}

func (t *Token) String() string {
	return t.Word
}

type DependencyNode struct {
	Dep            string `json:"dep"`
	Governor       int    `json:"governor"`
	GovernorGloss  string `json:"governorGloss"`
	Dependent      int    `json:"dependent"`
	DependentGloss string `json:"dependentGloss"`
}

package stanfordcorenlp

import (
	"context"
	"os"
	"testing"
)

func TestClient_DoTokenize(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	client := NewClient(context.Background(), url)

	resp, err := client.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize,
	)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	sen, err := UnmarshalSentence(resp)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	if len(sen.Tokens) == 0 {
		t.Errorf("Should not be empty, but was %v", sen.Tokens)
		t.FailNow()
	}
	if sen.Tokens[0].Word == "" {
		t.Errorf("Should not be zero, but was %v", sen.Tokens[0].Word)
	}
}

func TestClinet_DoPos(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	client := NewClient(context.Background(), url)

	resp, err := client.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos,
	)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	doc, err := UnmarshalDocument(resp)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	if sentences := doc.Sentences; len(sentences) == 0 {
		t.Errorf("Should not be empty, but was %v", sentences)
		t.FailNow()
	}
	if tokens := doc.Sentences[0].Tokens; len(tokens) == 0 {
		t.Errorf("Should not be empty, but was %v", tokens)
		t.FailNow()
	}
	if word := doc.Sentences[0].Tokens[0].Word; word == "" {
		t.Errorf("Should not be zero, but was %v", word)
	}
}

func TestClient_DoLemma(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	client := NewClient(context.Background(), url)

	resp, err := client.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos|AnnotatorLemma,
	)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	doc, err := UnmarshalDocument(resp)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	if sentences := doc.Sentences; len(sentences) == 0 {
		t.Errorf("Should not be empty, but was %v", sentences)
		t.FailNow()
	}
	if tokens := doc.Sentences[0].Tokens; len(tokens) == 0 {
		t.Errorf("Should not be empty, but was %v", tokens)
		t.FailNow()
	}
	if lemma := doc.Sentences[0].Tokens[0].Lemma; lemma == "" {
		t.Errorf("Should not be zero, but was %v", lemma)
	}
}

func TestClient_DoParse(t *testing.T) {
	url := os.Getenv("STANFORD_CORENLP_URL")
	client := NewClient(context.Background(), url)

	resp, err := client.Do(
		context.Background(),
		"The quick brown fox jumped over the lazy dog.",
		AnnotatorTokenize|AnnotatorSsplit|AnnotatorPos|AnnotatorParse,
	)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	doc, err := UnmarshalDocument(resp)
	if err != nil {
		t.Errorf("Received unexpected error: %+v", err)
		t.FailNow()
	}

	if !(len(doc.Sentences) >= 1) {
		t.Errorf("\"%v\" is not greater than or equal to \"%v\"", len(doc.Sentences), 1)
		t.FailNow()
	}

	sentence := doc.Sentences[0]
	if parse := sentence.Parse; parse == "" {
		t.Errorf("Should not be zero, but was %v", parse)
		t.FailNow()
	}

	for _, dependencies := range [][]*DependencyNode{
		sentence.BasicDependencies,
		sentence.EnhancedDependencies,
		sentence.EnhancedPlusPlusDependencies,
	} {
		if len(dependencies) == 0 {
			t.Errorf("Should not be empty, but was %v", dependencies)
			t.FailNow()
		}
		if dep := dependencies[0].Dep; dep == "" {
			t.Errorf("Should not be zero, but was %v", dep)
		}
		if governorGloss := dependencies[0].GovernorGloss; governorGloss == "" {
			t.Errorf("Should not be zero, but was %v", governorGloss)
		}
		if dependentGloss := dependencies[0].DependentGloss; dependentGloss == "" {
			t.Errorf("Should not be zero, but was %v", dependentGloss)
		}
	}
}

# go-stanfordcorenlp
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

stanfordcorenlp is a simple Golang client for [Stanford CoreNLP server](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html).

# Install
```shell
go get -u github.com/shota3506/go-stanfordcorenlp
```

# Usage
## Run Stanford CoreNLP server
Please run the Stanford CoreNLP server following [the official documentation](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html).
Or you can run the server under [docker](https://stanfordnlp.github.io/CoreNLP/other-languages.html#docker).

Make sure you use version 4.0.0 or above.

## Tokenize
```go
package main

import (
	"context"
	"fmt"
	"log"

	corenlp "github.com/shota3506/go-stanfordcorenlp"
)

func main() {
	ctx := context.Background()

	// create client for Stanford CoreNLP
	client := corenlp.NewClient(ctx, "http://localhost:9000")

	// sample text
	text := "The quick brown fox jumped over the lazy dog."
	resp, err := client.Do(ctx, text, corenlp.AnnotatorTokenize)
	if err != nil {
		log.Fatal(err)
	}

	s, err := corenlp.UnmarshalSentence(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // The quick brown fox jumped over the lazy dog .

	fmt.Println(s.Tokens[1].Word) // quick
	fmt.Println(s.Tokens[3].Word) // fox
}
```

## Pos Tagging
```go
package main

import (
	"context"
	"fmt"
	"log"

	corenlp "github.com/shota3506/go-stanfordcorenlp"
)

func main() {
	ctx := context.Background()

	// create client for Stanford CoreNLP
	client := corenlp.NewClient(ctx, "http://localhost:9000")

	// sample text
	text := "The quick brown fox jumped over the lazy dog."
	resp, err := client.Do(ctx, text,
		corenlp.AnnotatorTokenize|corenlp.AnnotatorSsplit|corenlp.AnnotatorPos)
	if err != nil {
		log.Fatal(err)
	}

	d, err := corenlp.UnmarshalDocument(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d) // The quick brown fox jumped over the lazy dog .

	fmt.Println(d.Sentences[0].Tokens[1].Word) // quick
	fmt.Println(d.Sentences[0].Tokens[1].Pos)  // JJ

	fmt.Println(d.Sentences[0].Tokens[3].Word) // fox
	fmt.Println(d.Sentences[0].Tokens[3].Pos)  // NN
}
```

## Parsing
```go
package main

import (
	"context"
	"fmt"
	"log"

	corenlp "github.com/shota3506/go-stanfordcorenlp"
)

func main() {
	ctx := context.Background()

	// create client for Stanford CoreNLP
	client := corenlp.NewClient(ctx, "http://localhost:9000")

	// sample text
	text := "The quick brown fox jumped over the lazy dog."
	resp, err := client.Do(ctx, text,
		corenlp.AnnotatorTokenize|corenlp.AnnotatorSsplit|corenlp.AnnotatorPos|corenlp.AnnotatorParse)
	if err != nil {
		log.Fatal(err)
	}

	d, err := corenlp.UnmarshalDocument(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(d.Sentences[0].Parse)
	/*
	  (ROOT
	    (S
	      (NP (DT The) (JJ quick) (JJ brown) (NN fox))
	      (VP (VBD jumped)
	        (PP (IN over)
	          (NP (DT the) (JJ lazy) (NN dog))))
	      (. .)))
	*/
}
```

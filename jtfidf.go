// Copyright 2019 ramenjuniti.

/*
Package jtfidf provides calculations of TF(Term Frequency), IDF(Inverse Document Frequency) and TF-IDF values at Japanese documents.

Dependencies

This package uses [kagome](https://github.com/ikawaha/kagome) as Morphological Analyzer.

About how to calulate TF-IDF value

The calculation of the TF-IDF value in this package uses the IDF value plus 1.
This is to prevent the TF-IDF value from becoming 0.
*/
package jtfidf

import (
	"math"

	"github.com/ikawaha/kagome/tokenizer"
)

func splitTerm(d string) []string {
	t := tokenizer.New()
	tokens := t.Tokenize(d)
	tokens = tokens[1 : len(tokens)-1]
	terms := make([]string, len(tokens))

	for i, token := range tokens {
		terms[i] = token.Surface
	}

	return terms
}

// AllTf returns all TF values in a doucument
func AllTf(d string) map[string]float64 {
	terms := splitTerm(d)
	n := len(terms)
	tfs := map[string]float64{}

	for _, term := range terms {
		if _, ok := tfs[term]; ok {
			tfs[term]++
		} else {
			tfs[term] = 1
		}
	}

	for term := range tfs {
		tfs[term] /= float64(n)
	}

	return tfs
}

// Tf returns TF value in a document
func Tf(t, d string) float64 {
	terms := splitTerm(d)
	n := len(terms)
	var count int

	for _, term := range terms {
		if t == term {
			count++
		}
	}

	return float64(count) / float64(n)
}

// AllIdf returns all IDF values in documents
func AllIdf(ds []string) map[string]float64 {
	n := len(ds)
	terms := []string{}
	termsList := make([][]string, n)

	for _, d := range ds {
		terms = append(terms, splitTerm(d)...)
	}

	for i, d := range ds {
		termsList[i] = splitTerm(d)
	}

	idfs := map[string]float64{}

	for _, term := range terms {
		var df int
		for i := 0; i < len(termsList); i++ {
			for j := 0; j < len(termsList[i]); j++ {
				if termsList[i][j] == term {
					df++
					break
				}
			}
		}
		if _, ok := idfs[term]; !ok {
			idfs[term] = math.Log(float64(n) / float64(df))
		}
	}

	return idfs
}

// Idf retuns IDF value in documents
func Idf(t string, ds []string) float64 {
	n := len(ds)
	termsList := make([][]string, n)
	var df int

	for i, d := range ds {
		termsList[i] = splitTerm(d)
	}

	for i := 0; i < len(termsList); i++ {
		for j := 0; j < len(termsList[i]); j++ {
			if t == termsList[i][j] {
				df++
				break
			}
		}
	}

	if df == 0 {
		return 0
	}

	return math.Log(float64(n) / float64(df))
}

// AllTfidf retuns all TF-IDF values in documents
func AllTfidf(ds []string) []map[string]float64 {
	idfs := AllIdf(ds)
	tfidfs := []map[string]float64{}

	for _, d := range ds {
		tfidf := map[string]float64{}
		for term, tf := range AllTf(d) {
			tfidf[term] = tf * (idfs[term] + 1)
		}
		tfidfs = append(tfidfs, tfidf)
	}

	return tfidfs
}

// Tfidf returns TF-IDF value in documents
func Tfidf(t, d string, ds []string) float64 {
	return Tf(t, d) * (Idf(t, ds) + 1)
}

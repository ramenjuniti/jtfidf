// Copyright 2019 ramenjuniti.

/*
Package jtfidf provides calculations of TF(Term Frequency), IDF(Inverse Document Frequency) and TF-IDF values at Japanese documents.

Package jtfidf use https://github.com/ikawaha/kagome as Morphological Analyzer.
*/
package jtfidf

import (
	"math"

	"github.com/ikawaha/kagome/tokenizer"
)

// A TfValue is a map of TF values
type TfValue map[string]float64

// A IdfValue is a map of IDF values
type IdfValue map[string]float64

// A TfidfValue is a map of TF-IDF values
type TfidfValue map[string]float64

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

// NewTf returns a TfValue containing values
func NewTf(d string) TfValue {
	return AllTf(d)
}

// NewIdf returns a IdfValue containing values
func NewIdf(ds []string) IdfValue {
	return AllIdf(ds)
}

// NewTfidf returns a TfIdfValue containing values
func NewTfidf(ds []string) TfidfValue {
	return AllTfidf(ds)
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

// AllTfidf retuns all TF-IDF values in documents
func AllTfidf(ds []string) map[string]float64 {
	idfs := AllIdf(ds)
	tfidfs := map[string]float64{}

	for _, d := range ds {
		for term, tf := range AllTf(d) {
			tfidfs[term] = tf * (idfs[term] + 1)
		}
	}

	return tfidfs
}

// Tf returns TF value in a document
func (tfv *TfValue) Tf(t string) float64 {
	return (*tfv)[t]
}

// Idf retuns IDF value in documents
func (idfv *IdfValue) Idf(t string) float64 {
	return (*idfv)[t]
}

// Tfidf returns TF-IDF value in documents
func (tfidfv *TfidfValue) Tfidf(t string) float64 {
	return (*tfidfv)[t]
}

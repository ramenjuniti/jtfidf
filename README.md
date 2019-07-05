# jtfidf

[![GoDoc](http://godoc.org/github.com/fabioberger/chrome?status.svg)](https://godoc.org/github.com/ramenjuniti/jtfidf)
[![CircleCI](https://circleci.com/gh/ramenjuniti/jtfidf.svg?style=svg)](https://circleci.com/gh/ramenjuniti/jtfidf)
[![Go Report Card](https://goreportcard.com/badge/github.com/ramenjuniti/jtfidf)](https://goreportcard.com/report/github.com/ramenjuniti/jtfidf)
[![codecov](https://codecov.io/gh/ramenjuniti/jtfidf/branch/master/graph/badge.svg)](https://codecov.io/gh/ramenjuniti/jtfidf)

Package jtfidf provides calculations of TF(Term Frequency), IDF(Inverse Document Frequency) and TF-IDF values at Japanese documents.

Package jtfidf use [kagome](https://github.com/ikawaha/kagome) as Morphological Analyzer.

## Install

```
go get -u github.com/ramenjuniti/jtfidf
```

## Usage

All usage are described in [GoDoc](https://godoc.org/github.com/ramenjuniti/jtfidf).

### AllTf

```go
func ExampleAllTf() {
	fmt.Println(AllTf("寿司が食べたい。"))
	// Output: map[。:0.2 が:0.2 たい:0.2 寿司:0.2 食べ:0.2]
}
```

### Tf

```go
func ExampleTfValue_Tf() {
	tfs := NewTf("寿司を食べたい。")
	fmt.Println(tfs.Tf("寿司"))
	// Output: 0.2
}
```

### AllIdf

```go
func ExampleAllIdf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(AllIdf(ds))
	// Output: map[。:1 が:1 たい:1 寿司:1 食べ:1]
}
```

### Idf

```go
func ExampleIdfValue_Idf() {
	ds := []string{
		"寿司が食べたい。",
	}
	idfs := NewIdf(ds)
	fmt.Println(idfs.Idf("寿司"))
	// Output: 1
}
```

### AllTfidf

```go
func ExampleAllIdf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(AllIdf(ds))
	// Output: map[。:1 が:1 たい:1 寿司:1 食べ:1]
}
```

### Tfidf

```go
func ExampleTfidfValue_Tfidf() {
	ds := []string{
		"寿司が食べたい。",
	}
	tfidfs := NewTfidf(ds)
	fmt.Println(tfidfs.Tfidf("寿司"))
	// Output: 0.2
}
```

## License

This software is released under the MIT License, see LICENSE.

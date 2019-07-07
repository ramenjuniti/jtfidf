# jtfidf

[![GoDoc](http://godoc.org/github.com/fabioberger/chrome?status.svg)](https://godoc.org/github.com/ramenjuniti/jtfidf)
[![CircleCI](https://circleci.com/gh/ramenjuniti/jtfidf.svg?style=svg)](https://circleci.com/gh/ramenjuniti/jtfidf)
[![Go Report Card](https://goreportcard.com/badge/github.com/ramenjuniti/jtfidf)](https://goreportcard.com/report/github.com/ramenjuniti/jtfidf)
[![codecov](https://codecov.io/gh/ramenjuniti/jtfidf/branch/master/graph/badge.svg)](https://codecov.io/gh/ramenjuniti/jtfidf)

Package jtfidf provides calculations of TF(Term Frequency), IDF(Inverse Document Frequency) and TF-IDF values at Japanese documents.

## Dependencies

This package uses [kagome](https://github.com/ikawaha/kagome) as Morphological Analyzer.

## About how to calulate TF-IDF value

The calculation of the TF-IDF value in this package uses the IDF value plus 1.
This is to prevent the TF-IDF value from becoming 0.

## Install

```
go get -u github.com/ramenjuniti/jtfidf
```

## Usage

All usage are described in [GoDoc](https://godoc.org/github.com/ramenjuniti/jtfidf).

### AllTf

AllTf returns all TF values in a doucument.

```go
func ExampleAllTf() {
	fmt.Println(AllTf("寿司が食べたい。"))
	// Output: map[。:0.2 が:0.2 たい:0.2 寿司:0.2 食べ:0.2]
}
```

### Tf

Tf returns TF value in a document.

```go
func ExampleTf() {
	fmt.Println(Tf("寿司", "寿司が食べたい。"))
	// Output: 0.2
}
```

### AllIdf

AllIdf returns all IDF values in documents.

```go
func ExampleAllIdf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(AllIdf(ds))
	// Output: map[。:0 が:0 たい:0 寿司:0 食べ:0]
}
```

### Idf

Idf retuns IDF value in documents.

```go
func ExampleIdf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(Idf("寿司", ds))
	// Output: 0
}
```

### AllTfidf

AllTfidf retuns all TF-IDF values in documents.

```go
func ExampleAllTfidf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(AllTfidf(ds))
	// Output: [map[。:0.2 が:0.2 たい:0.2 寿司:0.2 食べ:0.2]]
}
```

### Tfidf

Tfidf returns TF-IDF value in documents.

```go
func ExampleTfidf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(Tfidf("寿司", ds[0], ds))
	// Output: 0.2
}
```

## License

This software is released under the MIT License, see LICENSE.

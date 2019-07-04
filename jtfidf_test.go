package jtfidf

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitTerm(t *testing.T) {
	cases := []struct {
		name string
		d    string
		want []string
	}{
		{
			name: "case1",
			d:    "寿司",
			want: []string{
				"寿司",
			},
		},
		{
			name: "case2",
			d:    "寿司が食べたい。",
			want: []string{
				"寿司",
				"が",
				"食べ",
				"たい",
				"。",
			},
		},
		{
			name: "case3",
			d:    "",
			want: []string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := splitTerm(c.d); !reflect.DeepEqual(c.want, got) {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

var tfCases = []struct {
	name string
	d    string
	want map[string]float64
}{
	{
		name: "case1",
		d:    "寿司",
		want: map[string]float64{
			"寿司": 1,
		},
	},
	{
		name: "case2",
		d:    "寿司が食べたい。",
		want: map[string]float64{
			"寿司": 0.2,
			"が":  0.2,
			"食べ": 0.2,
			"たい": 0.2,
			"。":  0.2,
		},
	},
	{
		name: "case2",
		d:    "寿司が食べたい。寿司は好きです。",
		want: map[string]float64{
			"寿司": 0.2,
			"が":  0.1,
			"食べ": 0.1,
			"たい": 0.1,
			"。":  0.2,
			"は":  0.1,
			"好き": 0.1,
			"です": 0.1,
		},
	},
	{
		name: "case3",
		d:    "",
		want: map[string]float64{},
	},
}

func TestNewTf(t *testing.T) {
	for _, c := range tfCases {
		t.Run(c.name, func(t *testing.T) {
			if got := NewTf(c.d); reflect.DeepEqual(c.want, got) {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

func TestAllTf(t *testing.T) {
	for _, c := range tfCases {
		t.Run(c.name, func(t *testing.T) {
			if got := AllTf(c.d); !reflect.DeepEqual(c.want, got) {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

func ExampleAllTf() {
	fmt.Println(AllTf("寿司が食べたい。"))
	// Output: map[。:0.2 が:0.2 たい:0.2 寿司:0.2 食べ:0.2]
}

func TestTf(t *testing.T) {
	for _, c := range tfCases {
		t.Run(c.name, func(t *testing.T) {
			tfs := NewTf(c.d)
			for term, tf := range c.want {
				if tfs.Tf(term) != tf {
					t.Errorf("want %v: %v, but got %v: %v", term, tf, term, tfs.Tf(term))
				}
			}
		})
	}
}

func ExampleTfValue_Tf() {
	tfs := NewTf("寿司を食べたい。")
	fmt.Println(tfs.Tf("寿司"))
	// Output: 0.2
}

var idfCases = []struct {
	name string
	ds   []string
	want map[string]float64
}{
	{
		name: "case1",
		ds: []string{
			"寿司が食べたい。",
		},
		want: map[string]float64{
			"寿司": 1,
			"が":  1,
			"食べ": 1,
			"たい": 1,
			"。":  1,
		},
	},
	{
		name: "case2",
		ds: []string{
			"寿司が食べたい。",
			"寿司が食べたい。",
		},
		want: map[string]float64{
			"寿司": 1,
			"が":  1,
			"食べ": 1,
			"たい": 1,
			"。":  1,
		},
	},
	{
		name: "case3",
		ds:   []string{},
		want: map[string]float64{},
	},
}

func TestNewIdf(t *testing.T) {
	for _, c := range idfCases {
		t.Run(c.name, func(t *testing.T) {
			if got := NewIdf(c.ds); reflect.DeepEqual(c.want, got) {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

func TestAllIdf(t *testing.T) {
	for _, c := range idfCases {
		t.Run(c.name, func(t *testing.T) {
			if got := AllIdf(c.ds); !reflect.DeepEqual(c.want, got) {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

func ExampleAllIdf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(AllIdf(ds))
	// Output: map[。:1 が:1 たい:1 寿司:1 食べ:1]
}

func TestIdf(t *testing.T) {
	for _, c := range idfCases {
		t.Run(c.name, func(t *testing.T) {
			idfs := NewIdf(c.ds)
			for term, tf := range c.want {
				if idfs.Idf(term) != tf {
					t.Errorf("want %v: %v, but got %v: %v", term, tf, term, idfs.Idf(term))
				}
			}
		})
	}
}

func ExampleIdfValue_Idf() {
	ds := []string{
		"寿司が食べたい。",
	}
	idfs := NewIdf(ds)
	fmt.Println(idfs.Idf("寿司"))
	// Output: 1
}

var tfidfCases = []struct {
	name string
	ds   []string
	want map[string]float64
}{
	{
		name: "case1",
		ds: []string{
			"寿司が食べたい。",
		},
		want: map[string]float64{
			"寿司": 0.2,
			"が":  0.2,
			"食べ": 0.2,
			"たい": 0.2,
			"。":  0.2,
		},
	},
	{
		name: "case2",
		ds: []string{
			"寿司が食べたい。",
			"寿司が食べたい。",
		},
		want: map[string]float64{
			"寿司": 0.2,
			"が":  0.2,
			"食べ": 0.2,
			"たい": 0.2,
			"。":  0.2,
		},
	},
	{
		name: "case3",
		ds:   []string{},
		want: map[string]float64{},
	},
}

func TestNewTfidf(t *testing.T) {
	for _, c := range tfidfCases {
		t.Run(c.name, func(t *testing.T) {
			if got := NewTfidf(c.ds); reflect.DeepEqual(c.want, got) {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

func TestAllTfidf(t *testing.T) {
	for _, c := range tfidfCases {
		t.Run(c.name, func(t *testing.T) {
			if got := AllTfidf(c.ds); !reflect.DeepEqual(c.want, got) {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

func ExampleAllTfidf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(AllTfidf(ds))
	// Output: map[。:0.2 が:0.2 たい:0.2 寿司:0.2 食べ:0.2]
}

func TestTfidf(t *testing.T) {
	for _, c := range tfidfCases {
		t.Run(c.name, func(t *testing.T) {
			tfidfs := NewTfidf(c.ds)
			for term, tf := range c.want {
				if tfidfs.Tfidf(term) != tf {
					t.Errorf("want %v: %v, but got %v: %v", term, tf, term, tfidfs.Tfidf(term))
				}
			}
		})
	}
}

func ExampleTfidfValue_Tfidf() {
	ds := []string{
		"寿司が食べたい。",
	}
	tfidfs := NewTfidf(ds)
	fmt.Println(tfidfs.Tfidf("寿司"))
	// Output: 0.2
}

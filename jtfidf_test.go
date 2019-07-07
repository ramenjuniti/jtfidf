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
			for term, tf := range c.want {
				if got := Tf(term, c.d); tf != got {
					t.Errorf("want %v, but got %v", tf, got)
				}
			}
		})
	}
}

func ExampleTf() {
	fmt.Println(Tf("寿司", "寿司が食べたい。"))
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
			"寿司": 0,
			"が":  0,
			"食べ": 0,
			"たい": 0,
			"。":  0,
		},
	},
	{
		name: "case2",
		ds: []string{
			"寿司が食べたい。",
			"寿司が食べたい。",
		},
		want: map[string]float64{
			"寿司": 0,
			"が":  0,
			"食べ": 0,
			"たい": 0,
			"。":  0,
		},
	},
	{
		name: "case3",
		ds:   []string{},
		want: map[string]float64{},
	},
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
	// Output: map[。:0 が:0 たい:0 寿司:0 食べ:0]
}

func TestIdf(t *testing.T) {
	for _, c := range idfCases {
		t.Run(c.name, func(t *testing.T) {
			for term, idf := range c.want {
				if got := Idf(term, c.ds); idf != got {
					t.Errorf("want %v, but got %v", idf, got)
				}
			}
		})
	}
}

func ExampleIdf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(Idf("寿司", ds))
	// Output: 0
}

var tfidfCases = []struct {
	name string
	ds   []string
	want []map[string]float64
}{
	{
		name: "case1",
		ds: []string{
			"寿司が食べたい。",
		},
		want: []map[string]float64{
			{
				"寿司": 0.2,
				"が":  0.2,
				"食べ": 0.2,
				"たい": 0.2,
				"。":  0.2,
			},
		},
	},
	{
		name: "case2",
		ds: []string{
			"寿司が食べたい。",
			"寿司が食べたい。",
		},
		want: []map[string]float64{
			{
				"寿司": 0.2,
				"が":  0.2,
				"食べ": 0.2,
				"たい": 0.2,
				"。":  0.2,
			},
			{
				"寿司": 0.2,
				"が":  0.2,
				"食べ": 0.2,
				"たい": 0.2,
				"。":  0.2,
			},
		},
	},
	{
		name: "case3",
		ds:   []string{},
		want: []map[string]float64{},
	},
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
	// Output: [map[。:0.2 が:0.2 たい:0.2 寿司:0.2 食べ:0.2]]
}

func TestTfidf(t *testing.T) {
	for _, c := range tfidfCases {
		t.Run(c.name, func(t *testing.T) {
			for i, d := range c.ds {
				t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
					for _, w := range c.want {
						for term, tfidf := range w {
							if got := Tfidf(term, d, c.ds); tfidf != got {
								t.Errorf("want %v, but got %v", tfidf, got)
							}
						}
					}
				})
			}
		})
	}
}

func ExampleTfidf() {
	ds := []string{
		"寿司が食べたい。",
	}
	fmt.Println(Tfidf("寿司", ds[0], ds))
	// Output: 0.2
}

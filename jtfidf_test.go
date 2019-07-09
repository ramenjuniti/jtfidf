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

func TestAllTf(t *testing.T) {
	cases := []struct {
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

	for _, c := range cases {
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
	cases := []struct {
		name string
		t    string
		d    string
		want float64
	}{
		{
			name: "case1",
			t:    "寿司",
			d:    "寿司",
			want: 1,
		},
		{
			name: "case2",
			t:    "寿司",
			d:    "寿司が食べたい。",
			want: 0.2,
		},
		{
			name: "case3",
			t:    "寿司",
			d:    "",
			want: 0,
		},
		{
			name: "case4",
			t:    "",
			d:    "",
			want: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Tf(c.t, c.d); c.want != got {
				t.Errorf("want %v, but got %v", c.want, got)
			}
		})
	}
}

func ExampleTf() {
	fmt.Println(Tf("寿司", "寿司が食べたい。"))
	// Output: 0.2
}

func TestAllIdf(t *testing.T) {
	cases := []struct {
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

	for _, c := range cases {
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
	cases := []struct {
		name string
		t    string
		ds   []string
		want float64
	}{
		{
			name: "case1",
			t:    "寿司",
			ds: []string{
				"寿司が食べたい",
			},
			want: 0,
		},
		{
			name: "case2",
			t:    "寿司",
			ds: []string{
				"寿司が食べたい",
				"寿司が食べたい",
			},
			want: 0,
		},
		{
			name: "case3",
			t:    "寿司",
			ds: []string{
				"",
			},
			want: 0,
		},
		{
			name: "case3",
			t:    "",
			ds: []string{
				"",
			},
			want: 0,
		},
		{
			name: "case3",
			t:    "",
			ds:   []string{},
			want: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Idf(c.t, c.ds); c.want != got {
				t.Errorf("want %v, but got %v", c.want, got)
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

func TestAllTfidf(t *testing.T) {
	cases := []struct {
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

	for _, c := range cases {
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
	cases := []struct {
		name string
		t    string
		d    string
		ds   []string
		want float64
	}{
		{
			name: "case1",
			t:    "寿司",
			d:    "寿司が食べたい。",
			ds: []string{
				"寿司が食べたい。",
			},
			want: 0.2,
		},
		{
			name: "case2",
			t:    "寿司",
			d:    "",
			ds: []string{
				"",
			},
			want: 0,
		},
		{
			name: "case3",
			t:    "",
			d:    "",
			ds: []string{
				"",
			},
			want: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Tfidf(c.t, c.d, c.ds); c.want != got {
				t.Errorf("want %v, but got %v", c.want, got)
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

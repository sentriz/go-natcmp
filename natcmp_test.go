package natcmp

import (
	"math/rand/v2"
	"slices"
	"testing"
)

var rnd = rand.New(rand.NewPCG(1, 2))

func TestCompare(t *testing.T) {
	eq(t, -1, Compare("abc10", "abc100"))
	eq(t, +1, Compare("abc100", "abc10"))
	eq(t, +1, Compare("abc10.20 final.zip", "abc10.10 final.zip"))
	eq(t, 0, Compare("", ""))
	eq(t, 0, Compare("abc100", "abc100"))

	resort := func(sorted []string, cmp func(a, b string) int) []string {
		vs := shuffled(sorted)
		slices.SortFunc(vs, cmp)
		return vs
	}

	for i := 0; i < 16; i++ {
		eqFunc(t, slices.Equal, sciValues, resort(sciValues, Compare))
	}
	for i := 0; i < 16; i++ {
		eqFunc(t, slices.Equal, docValues, resort(docValues, Compare))
	}
}

func FuzzCompareWithSlow(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b string) {
		exp := slowCompare(a, b)
		got := Compare(a, b)
		if got != exp {
			t.Errorf("compare %q %q, exp %d got %d", a, b, exp, got)
		}
	})
}

func BenchmarkCompare(b *testing.B) {
	benchCompare(b, Compare)
}

func BenchmarkCompareSlow(b *testing.B) {
	benchCompare(b, slowCompare)
}

func benchCompare(b *testing.B, cmp func(a, b string) int) {
	shuff := shuffled(sciValues)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := slices.Clone(shuff)
		b.StartTimer()
		slices.SortFunc(data, cmp)
	}
}

func eq[T comparable](t *testing.T, a, b T) {
	t.Helper()
	eqFunc(t, func(a, b T) bool { return a == b }, a, b)
}

func eqFunc[T any](t *testing.T, eq func(a, b T) bool, a, b T) {
	t.Helper()
	if !eq(a, b) {
		t.Errorf("%+v != %v", a, b)
	}
}

func shuffled[T any](sorted []T) []T {
	vs := slices.Clone(sorted)
	rnd.Shuffle(len(vs), func(i, j int) {
		vs[i], vs[j] = vs[j], vs[i]
	})
	return vs
}

var (
	sciValues = []string{
		"10X Radonius",
		"20X Radonius",
		"20X Radonius Prime",
		"30X Radonius",
		"40X Radonius",
		"200X Radonius",
		"1000X Radonius Maximus",
		"Allegia 6R Clasteron",
		"Allegia 50 Clasteron",
		"Allegia 50B Clasteron",
		"Allegia 51 Clasteron",
		"Allegia 500 Clasteron",
		"Alpha 2",
		"Alpha 2A",
		"Alpha 2A-900",
		"Alpha 2A-8000",
		"Alpha 100",
		"Alpha 200",
		"Callisto Morphamax",
		"Callisto Morphamax 500",
		"Callisto Morphamax 600",
		"Callisto Morphamax 700",
		"Callisto Morphamax 5000",
		"Callisto Morphamax 6000 SE",
		"Callisto Morphamax 6000 SE2",
		"Callisto Morphamax 7000",
		"Xiph Xlater 5",
		"Xiph Xlater 40",
		"Xiph Xlater 50",
		"Xiph Xlater 58",
		"Xiph Xlater 300",
		"Xiph Xlater 500",
		"Xiph Xlater 2000",
		"Xiph Xlater 5000",
		"Xiph Xlater 10000",
	}

	docValues = []string{
		"z1.doc",
		"z2.doc",
		"z3.doc",
		"z4.doc",
		"z5.doc",
		"z6.doc",
		"z7.doc",
		"z8.doc",
		"z9.doc",
		"z10.doc",
		"z11.doc",
		"z12.doc",
		"z13.doc",
		"z14.doc",
		"z15.doc",
		"z16.doc",
		"z17.doc",
		"z18.doc",
		"z19.doc",
		"z20.doc",
		"z100.doc",
		"z101.doc",
		"z102.doc",
	}
)

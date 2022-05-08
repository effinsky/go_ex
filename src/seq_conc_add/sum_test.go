package seqconcadd

import (
	"testing"
)

// test table
// with inline struct defs
var sumTests = []struct{ a, b, expected int }{
	{1, 2, 3},
	{2, 3, 5},
	{5, 3, 8},
}

var sumTestWithZeroArgs = []struct{ expected int }{
	{0},
}

func TestSum(t *testing.T) {
	for _, test := range sumTests {
		if got := Sum(test.a, test.b); got != test.expected {
			t.Errorf("Got: %v, wanted: %v", got, test.expected)
		}
	}
}

func TestSumZeroArgs(t *testing.T) {
	for _, test := range sumTestWithZeroArgs {
		if got := Sum(); got != test.expected {
			t.Errorf("Got: %v, wanted: %v", got, test.expected)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3)
	}
}

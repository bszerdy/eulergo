package internal

import (
	"reflect"
	"testing"
)

func TestFactors(t *testing.T) {
	data := []struct {
		n      int64
		expect []int64
	}{
		{8, []int64{2, 4}},
		{25, []int64{5}},
	}

	i := &Integer64{}
	for _, x := range data {
		i.n = x.n
		actual := i.Factors()

		if !reflect.DeepEqual(actual, x.expect) {
			t.Fatalf("TestFactors - expected: %v, got: %v", x.expect, actual)
		}
	}
}

func TestIsPrime(t *testing.T) {
	data := []struct {
		n      int
		expect bool
	}{
		{3, true},
		{4, false},
		{7, true},
	}

	i := &Integer{}
	for _, x := range data {
		i.n = x.n
		actual := i.IsPrime()

		if actual != x.expect {
			t.Fatalf("TestIsPrime - expected: %v, got: %v", x.expect, actual)
		}
	}
}

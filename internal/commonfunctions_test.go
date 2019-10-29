package internal

import "testing"

func TestIntegerContains(t *testing.T) {
	containsData := []struct {
		data []int

		expect bool
	}{
		{
			data:   []int{1, 2, 3, 4, 5},
			expect: true,
		}, {

			data:   []int{1, 3, 4, 5},
			expect: false,
		},
	}

	i := &Integer{n: 2}

	for _, x := range containsData {
		actual := i.Contains(x.data)
		if actual != x.expect {
			t.Fatalf("TestContains - expected: %v, got: %v", x.expect, actual)
		}
	}
}

func TestInteger64Contains(t *testing.T) {
	containsData := []struct {
		data   []int64
		expect bool
	}{
		{data: []int64{1, 2, 3, 4, 5},
			expect: true,
		}, {

			data:   []int64{1, 3, 4, 5},
			expect: false,
		},
	}

	i := &Integer64{n: 2}

	for _, x := range containsData {
		actual := i.Contains(x.data)
		if actual != x.expect {
			t.Fatalf("TestContains - expected: %v, got: %v", x.expect, actual)
		}
	}
}

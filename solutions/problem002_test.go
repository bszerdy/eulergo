package solutions

import "testing"

func TestProblem002(t *testing.T) {
	const expected = 4613732

	p := &Problem{}

	actual := p.Problem002()

	if actual != expected {
		t.Fatalf("Problem002 - expected: %d, got %d", expected, actual)
	}
}

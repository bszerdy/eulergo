package solutions

import (
	"testing"
)

func TestProblem001(t *testing.T) {
	const expected = 233168

	p := &Problem{}

	actual := p.Problem001()
	if actual != expected {
		t.Fatalf("Problem001 - expected: %d, got: %d", expected, actual)
	}
}

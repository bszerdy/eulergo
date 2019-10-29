package solutions

import "testing"

func TestProblem003(t *testing.T) {
	const expected = 6857
	p := &Problem{}

	actual := p.Problem003()
	if actual != expected {
		t.Fatalf("Problem003 - expected: %d, got: %d", expected, actual)
	}

}

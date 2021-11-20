package main

import "testing"

func TestActuate(t *testing.T) {
	dat := []struct {
		i int
		o int
	}{
		{i: 24, o: -1}, // case:0, input 24, expect output to be -1 (cooling)
		{i: 25, o: 0},
		{i: 30, o: 0},
		{i: 31, o: 1},
	}

	for n, d := range dat {
		if o := actuate(d.i); o != d.o {
			t.Errorf("case: %d, input: %d, expected: %d, got: %d\n", n, d.i, d.o, o)
		}
	}

}

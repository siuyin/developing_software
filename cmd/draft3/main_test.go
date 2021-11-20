package main

import "testing"

func TestActuate(t *testing.T) {
	dat := []struct {
		i int
		o int
	}{
		{i: 24, o: -1}, // case: 0, input 24, expect output to be -1 (too cold)
		{i: 25, o: 0},  // case: 1, input 25, expected output 0 (just right)
		{i: 30, o: 0},  // case: 2, input 30, exppected output 0 (just right)
		{i: 31, o: 1},  // case: 3, input 31, expected output 1 (too hot)
	}

	for n, d := range dat {
		if o := actuate(d.i); o != d.o {
			t.Errorf("case: %d, input: %d, expected: %d, got: %d\n", n, d.i, d.o, o)
		}
	}

}

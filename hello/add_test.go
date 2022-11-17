package main

import "testing"

func TestAbs(t *testing.T) {
	got := add(-1, 3)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

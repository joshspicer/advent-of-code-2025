package day04

import "testing"

func TestExample(t *testing.T) {
	t.Skip("TODO: add example expectations")

	part01, part02 := Run(true, "example")
	var expected01 int
	var expected02 int

	if part01 != expected01 {
		t.Fatalf("Part 01: got %d, expected %d", part01, expected01)
	}

	if part02 != expected02 {
		t.Fatalf("Part 02: got %d, expected %d", part02, expected02)
	}
}

func TestPuzzle(t *testing.T) {
	t.Skip("TODO: add puzzle expectations")

	part01, part02 := Run(true, "puzzle")
	var expected01 int
	var expected02 int

	if part01 != expected01 {
		t.Fatalf("Part 01: got %d, expected %d", part01, expected01)
	}

	if part02 != expected02 {
		t.Fatalf("Part 02: got %d, expected %d", part02, expected02)
	}
}

func TestSamples(t *testing.T) {
	t.Skip("TODO: add test cases in tests/")

	part01, part02 := Run(true, "tests/01")
	var expected01 int
	var expected02 int

	if part01 != expected01 {
		t.Fatalf("Part 01: got %d, expected %d", part01, expected01)
	}

	if part02 != expected02 {
		t.Fatalf("Part 02: got %d, expected %d", part02, expected02)
	}
}

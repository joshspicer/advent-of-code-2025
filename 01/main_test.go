package day01

import "testing"

func TestExample(t *testing.T) {
	var expected01 int64 = 3
	var expected02 int64 = 6
	part01, part02 := Run(true, "example")
	if part01 != expected01 {
		t.Errorf("Part 01: Got %d, expected %d", part01, expected01)
	}
	if part02 != expected02 {
		t.Errorf("Part 02: Got %d, expected %d", part02, expected02)
	}
}

func TestPuzzle(t *testing.T) {
	var expected01 int64 = 1102
	var expected02 int64 = 6175
	part01, part02 := Run(true, "puzzle")
	if part01 != expected01 {
		t.Errorf("Part 01: Got %d, expected %d", part01, expected01)
	}
	if part02 != expected02 {
		t.Errorf("Part 02: Got %d, expected %d", part02, expected02)
	}
}

func Test01(t *testing.T) {
	var expected int64 = 7
	_, part02 := Run(true, "tests/01")
	if part02 != expected {
		t.Errorf("Part 02: Got %d, expected %d", part02, expected)
	}
}

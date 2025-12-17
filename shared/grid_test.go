package shared

import (
	"fmt"
	"testing"
)

func TestToGrid(t *testing.T) {
	input := []string{
		"1234",
		"5678",
		"9ABC"}
	grid := ToGrid(input, func(r rune) string { return string(r) })

	t1 := grid.At(0, 0)
	if t1 != "1" {
		t.Fatalf("expected 1, got %s", t1)
	}

	t2 := grid.At(1, 0)
	if t2 != "5" {
		t.Fatalf("expected 5, got %s", t2)
	}

	tc := grid.At(2, 3)
	if tc != "C" {
		t.Fatalf("expected C, got %s", tc)
	}

	fmt.Print(grid.String())
}

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

	t1, err := grid.At(0, 0)
	if err != nil || t1 != "1" {
		t.Fatalf("expected 1, got %s", t1)
	}

	t2, err := grid.At(1, 0)
	if err != nil || t2 != "5" {
		t.Fatalf("expected 5, got %s", t2)
	}

	tc, err := grid.At(2, 3)
	if err != nil || tc != "C" {
		t.Fatalf("expected C, got %s", tc)
	}

	terr, err := grid.At(3, 3)
	if err == nil {
		t.Fatalf("expected error, got", terr)
	}

	fmt.Print(grid.String())
}

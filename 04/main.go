package day04

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
)

const Day = "04"

// ======================================================== //

var _debug bool

func DEBUG(a ...any) {
	if _debug {
		fmt.Fprint(os.Stderr, a...)
	}
}

func DEBUGF(format string, a ...any) {
	if _debug {
		fmt.Fprintf(os.Stderr, format, a...)
	}
}

// ======================================================== //

func Run(debug bool, targetFile string) (int, int) {
	_debug = debug

	lines := shared.ReadLines(targetFile)
	grid := shared.ToGrid(lines, func(r rune) string { return string(r) })

	var part01 int
	var part02 int

	part01 = run(grid, false)
	part02 = run(grid, true)

	return part01, part02
}

func run(grid shared.Grid[string], remove bool) int {
	result := 0
	keepGoing := true

	prevRolls := -1
	currentRolls := 0
	grid.ForEach(func(row, col int, v string) {
		if v == "@" {
			currentRolls += 1
		}
	})

	for keepGoing {
		prevRolls = currentRolls
		grid.ForEach(func(row, col int, v string) {
			if v == "@" {
				adjs := grid.CollectAdjacent(row, col, shared.EightAdjacentRelativeOffsets)
				numAdjacentRolls := 0
				for _, ch := range adjs {
					if ch == "@" {
						numAdjacentRolls += 1
					}
				}
				if numAdjacentRolls < 4 {
					DEBUGF("(%d, %d) has %d\n", row, col, numAdjacentRolls)
					if remove {
						currentRolls -= 1
						grid.Mutate(row, col, ".")
					}
					result += 1
				}
			}
		})
		keepGoing = remove && (prevRolls != currentRolls)
	}

	DEBUGF("result: %d", result)
	return result
}

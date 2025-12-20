package day07

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
)

const Day = "07"

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

// Run executes both parts and returns their results.
func Run(debug bool, targetFile string) (int, int) {
	_debug = debug

	lines := shared.ReadLines(targetFile)
	grid := shared.ToGrid(lines, func(r rune) string { return string(r) })

	DEBUG(grid.String() + "\n")

	var part01 int
	var part02 int

	// Counts how many times we observe this beam successfully split
	seen := shared.CreateSet()

	keepGoing := true
	for keepGoing {
		prevGrid := grid.Copy()
		grid.ForEach(func(row, col int, v string) {
			freeSpace := false
			switch v {
			case "S":
				grid.MutateIgnoringBounds(row+1, col, "|")
			case "^":
				// If a beam is above it
				above, err := grid.At(row-1, col)
				if err != nil {
					// ignore
				}
				if above == "|" {
					seen.Add(fmt.Sprintf("%d,%d", row, col))
					grid.MutateIgnoringBounds(row, col-1, "|")
					grid.MutateIgnoringBounds(row, col+1, "|")
				}
			case "|":
				below, err := grid.At(row+1, col)
				if err != nil {
					// ignore
				}
				if below != "^" {
					grid.MutateIgnoringBounds(row+1, col, "|")
				}
			case ".":
				freeSpace = true
			}

			// DEBUG
			if !freeSpace {
				DEBUG(grid.String() + "\n")
				DEBUG(seen.Size())
				DEBUG("\n\n\n\n")
			}

		})
		if grid.String() == prevGrid.String() {
			keepGoing = false
		}
	}

	part01 = seen.Size()

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

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

	keepGoing := true
	for keepGoing {
		prevGrid := grid.Copy()
		grid.ForEach(func(row, col int, v string) {
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
					if grid.MutateIgnoringBounds(row, col-1, "|") {
						// note: Arbitrarily choosing to count left successful mutations
						if prevLeft, err := prevGrid.At(row, col-1); err == nil {
							if prevLeft != "|" {
								part01++
							}
						}
					}
					if grid.MutateIgnoringBounds(row, col+1, "|") {
						// if prevRight, err := prevGrid.At(row, col+1); err == nil {
						// if prevRight != "|" {
						// part01++
						// }
						// }
					}
				}
			case "|":
				below, err := grid.At(row+1, col)
				if err != nil {
					// ignore
				}
				if below != "^" {
					grid.MutateIgnoringBounds(row+1, col, "|")
				}
			}

			// DEBUG(grid.String() + "\n")
		})
		if grid.String() == prevGrid.String() {
			keepGoing = false
		}
	}

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

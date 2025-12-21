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

func Run(debug bool, targetFile string) (int, int) {
	_debug = debug

	lines := shared.ReadLines(targetFile)
	grid := shared.ToGrid(lines, func(r rune) string { return string(r) })

	DEBUG(grid.String() + "\n")

	var part01 int
	var part02 int

	keepGoing := true
	seen := shared.CreateSet[string]()
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
				// Skip
			}

		})
		if grid.String() == prevGrid.String() {
			keepGoing = false
		}
	}

	part01 = seen.Size()
	fmt.Println(part01)

	graph := shared.ToDAG(grid, func(r, c int) ([]shared.Tuple[int], shared.DFlags) {
		flags := shared.None
		if val, err := grid.At(r, c); err == nil {
			edges := make([]shared.Tuple[int], 0)
			switch val {
			case "S":
				// Special case
				flags |= shared.Start
				edges = append(edges, shared.Tuple[int]{r + 1, c})
			case "|":
				edges = grid.CollectAdjacentIndices(r, c, []string{"|", "^"}, []shared.Tuple[int]{{1, 0}})
				if r == grid.Height()-1 {
					// Final Row
					flags |= shared.End
				}
			case "^":
				edges = grid.CollectAdjacentIndices(r, c, []string{"|"}, []shared.Tuple[int]{{0, -1}, {0, 1}})
			case ".":
				// Uninteresting
			default:
				panic("unexpected")
			}
			return edges, flags
		}
		return make([]shared.Tuple[int], 0), flags
	})
	start := graph.WithFlags(shared.Start)
	part02 = int(graph.CountPaths(start[0]))

	fmt.Println(part02)

	return part01, part02
}

package day05

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Day = "05"

// ======================================================== //

var _debug bool

func DEBUGLN(a ...any) {
	if _debug {
		fmt.Fprintln(os.Stderr, a...)
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

	ranges := make([]shared.Tuple[int], 0)
	targets := make([]int, 0)
	// Parse
	for _, entry := range lines {
		if entry == "" {
			continue
		}
		spl := strings.Split(entry, "-")
		if len(spl) > 1 {
			// Range
			first, err := strconv.Atoi(spl[0])
			if err != nil {
				panic("oops start")
			}
			second, err := strconv.Atoi(spl[1])
			if err != nil {
				panic("oops end")
			}
			ranges = append(ranges, shared.Tuple[int]{first, second})
		} else {
			// Target
			target, err := strconv.Atoi(entry)
			if err != nil {
				panic("oops target")
			}
			targets = append(targets, target)
		}
	}

	// DEBUGLN(ranges)
	// DEBUGLN(targets)

	// Order by first element
	// slices.SortFunc(ranges, func(a, b shared.Tuple[int]) int {
	// 	return a.First - b.First
	// })
	// Todo: Consolidate groups?

	DEBUGLN(ranges)

	var part01 int
	var part02 int

	for _, target := range targets {
		// Check all ranges (inclusive)
		// Where 'start' includes target
		for _, rng := range ranges {
			if target >= rng.First && target <= rng.Second {
				part01 += 1
				break
			}
		}
	}

	fmt.Println(part01)
	fmt.Println(part02) // TODO

	return part01, part02
}

package day05

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"slices"
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

	// Order by first element
	slices.SortFunc(ranges, func(a, b shared.Tuple[int]) int {
		return a.First - b.First
	})

	DEBUGLN(ranges)

	// Consolidate ranges by shortening so there is no overlap
	var merged = make([]shared.Tuple[int], 0)
	for i := 0; i < len(ranges); i++ {
		curr := &ranges[i]
		if len(merged) == 0 {
			merged = append(merged, *curr)
			continue
		}
		lastMerged := &merged[len(merged)-1]

		if lastMerged.Second >= curr.First {
			lastMerged.Second = shared.Max(lastMerged.Second, curr.Second)
		} else {
			merged = append(merged, *curr)
		}
	}

	// DEBUGLN(ranges)

	var part01 int
	var part02 int

	for _, target := range targets {
		for _, rng := range merged {
			if target >= rng.First && target <= rng.Second {
				DEBUGF("✅ %d is within the range %d-%d\n", target, rng.First, rng.Second)
				part01 += 1
				break
			}
		}
	}

	for _, rng := range merged {
		rangeSize := (rng.Second - rng.First) + 1
		DEBUGF("+ range %d-%d -> %d\n", rng.First, rng.Second, rangeSize)
		part02 += rangeSize
	}

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

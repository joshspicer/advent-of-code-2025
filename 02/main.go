package day02

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"strconv"
)

const Day = "02"

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
	items := shared.ReadCommaSeparatedLine(targetFile)
	pairs := shared.ParseDashedTuple(items)

	var part01 int
	var part02 int

	// DEBUG(pairs)
	for _, p := range pairs {
		part01 += examineRange(p, p1Filter)

	}
	DEBUGF("Part 01:   %d", part01)

	return part01, part02
}

// Find the invalid IDs by looking for any ID which is made only of
// some sequence of digits repeated twice.
// - 55 (5 twice)
// - 6464 (64 twice)
// - 123123 (123 twice)
func p1Filter(num int) bool {
	if num <= 10 {
		// Can't be silly if values can't repeat
		// 10 isn't silly either
		return false
	}
	s := fmt.Sprintf("%d", num)

	numBuckets := len(s)
	if numBuckets%2 != 0 {
		numBuckets--
	}

	for numBuckets > 1 {
		buckets := shared.BucketString(s, numBuckets)
		if shared.AllEqual(buckets...) {
			return true // Some bucketing of 'num' is self-similar
		}
		numBuckets = numBuckets - 2
	}
	// Assume failure
	return false
}

func examineRange(tuple shared.Tuple[string], filter func(num int) bool) int {
	sum := 0
	start, err := strconv.Atoi(tuple.First)
	if err != nil {
		panic("with first number")
	}
	end, err := strconv.Atoi(tuple.Second)
	if err != nil {
		panic("with second number")
	}
	for i := start; i <= end; i++ {
		if filter(i) {
			DEBUGF("** yes to %d\n", i)
			sum += i
		}
	}
	return sum
}

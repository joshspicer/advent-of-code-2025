package day03

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"strconv"
)

const Day = "03"

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

	var part01 int
	var part02 int

	for _, line := range lines {
		part01 += calculateBestJoltage(line)
	}

	DEBUGF("Part 1: %d", part01)

	return part01, part02
}

func calculateBestJoltage(line string) int {
	bestFirstIdx, bestFirstCh := 0, 0
	bestLastIdx, bestLastCh := len(line)-1, 0
	for idx, ch := range line {
		if idx == len(line)-1 {
			// Last number is off limits
			break
		}
		iCh := int(ch - '0')
		if iCh > bestFirstCh {
			bestFirstIdx, bestFirstCh = idx, iCh
		}
	}

	for idx := len(line) - 1; idx > bestFirstIdx; idx-- {
		revCh, err := strconv.Atoi(line[idx : idx+1])
		if err != nil {
			panic("atoi")
		}
		if revCh > bestLastCh {
			bestLastIdx, bestLastCh = idx, revCh
		}
	}

	concat := fmt.Sprintf("%d%d", bestFirstCh, bestLastCh)
	DEBUGF("%s  ([%d]-[%d])\n", concat, bestFirstIdx, bestLastIdx)
	result, err := strconv.Atoi(concat)
	if err != nil {
		panic("strconv of concat")
	}
	return result
}

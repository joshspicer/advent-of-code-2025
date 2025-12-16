package day03

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		part01 += calculateBestJoltage2(line, 2)
		part02 += calculateBestJoltage2(line, 12)
	}

	DEBUGF("Part 1: %d\n", part01)
	DEBUGF("Part 2: %d", part02)

	return part01, part02
}

func calculateBestJoltage2(line string, numBatteries int) int {
	length := len(line)
	lowerBound := 0
	batteriesAlreadyCollected := 0
	var result []string = make([]string, 0)

	for batteriesAlreadyCollected != numBatteries {
		tmpBestIdx, tmpBestVal, tmpBestString := 0, 0, "X"
		for idx := lowerBound; idx < (length - (numBatteries - batteriesAlreadyCollected - 1)); idx++ {
			currValString := line[idx : idx+1]
			curVal, err := strconv.Atoi(currValString)
			if err != nil {
				panic("oops")
			}
			if curVal > tmpBestVal {
				tmpBestIdx, tmpBestVal, tmpBestString = idx, curVal, currValString
			}
		}
		result = append(result, tmpBestString)
		batteriesAlreadyCollected++ // We captured a battery
		lowerBound = tmpBestIdx + 1 // We cannot use any numbers to the left of what we captured
	}
	num, err := strconv.Atoi(strings.Join(result, ""))
	if err != nil {
		panic("reconstruct failure")
	}
	return num
}

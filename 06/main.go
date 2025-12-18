package day06

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"strconv"
)

const Day = "06"

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

type MathProblem struct {
	Values    []int
	Operation string
}

func (mp *MathProblem) addValue(value string) {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic("yike")
	}
	mp.Values = append(mp.Values, num)
}

func (mp MathProblem) calculate() int {
	result := 1
	switch mp.Operation {
	case "*":
		for _, v := range mp.Values {
			result *= v
		}
	case "+":
		result = 0
		for _, v := range mp.Values {
			result += v
		}
	}
	return result
}

// Run executes both parts and returns their results.
func Run(debug bool, targetFile string) (int, int) {
	_debug = debug

	lines := shared.ReadLines(targetFile)

	numValues := len(lines) - 1
	mathProblems := make([]MathProblem, 0)

	for lineIdx, line := range lines {
		line += " " // Normalize

		currVal := ""
		currMathProblem := 0

		for _, rune := range line {
			if rune == ' ' {
				if currVal == "" {
					continue
				}
				if lineIdx == 0 {
					mathProblems = append(mathProblems, MathProblem{})
				}
				if lineIdx == numValues {
					mathProblems[currMathProblem].Operation = currVal
				} else {
					mathProblems[currMathProblem].addValue(currVal)
				}
				currVal = ""
				currMathProblem++
			} else {
				currVal += string(rune)
			}
		}
	}

	DEBUG(mathProblems)

	var part01 int
	var part02 int

	for _, mp := range mathProblems {
		part01 += mp.calculate()
	}

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

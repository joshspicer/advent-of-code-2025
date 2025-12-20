package day06

import (
	"advent-of-code-2025/shared"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
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
	_Values     []string
	Values      []int
	Operation   string
	LeftAligned bool
}

func (mp *MathProblem) addValue(value string) {
	indexWhitespace := strings.Index(value, " ")
	if indexWhitespace != -1 {
		mp.LeftAligned = indexWhitespace > 0
	}
	value = strings.TrimSpace(value)
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
	default:
		panic("aaaah")
	}
	return result
}

// Normalizes to a 'right aligned' MathProblem
// By reflecting all values across themselves
func (mp MathProblem) normalize() MathProblem {
	if !mp.LeftAligned {
		return mp
	}
	mp.Values = make([]int, 0)
	for _, _val := range mp._Values {
		_val := shared.ReverseString(_val)
		mp.addValue(_val)
	}
	mp.LeftAligned = false
	return mp
}

func (mp MathProblem) calculate2() int {
	maxNumLength := 0
	for _, v := range mp.Values {
		// log10(...)
		maxNumLength = shared.Max(len(fmt.Sprintf("%d", v)), maxNumLength)
	}

	mp = mp.normalize()

	// Now all MathProblem are 'right-handed'

	output := make([]string, maxNumLength)
	for slice := range maxNumLength {
		placeValue := int(math.Pow(10, float64((maxNumLength - slice - 1))))
		for _, val := range mp.Values {
			padding := placeValue
			v := strconv.Itoa(val / padding % 10)

			// This assumption breaks if the puzzle
			// had any zeros, but it seems like it doesn't...
			if v != "0" {
				output[slice] += v
			}
		}
	}

	values := make([]int, 0)
	for _, o := range output {
		converted, err := strconv.Atoi(o)
		if err != nil {
			converted = -1
		}
		values = append(values, converted)
	}

	mp2 := MathProblem{Values: values, Operation: mp.Operation, LeftAligned: mp.LeftAligned}
	DEBUGF("+ %v\n", mp2)
	return mp2.calculate()
}

// Run executes both parts and returns their results.
func Run(debug bool, targetFile string) (int, int) {
	_debug = debug

	lines := shared.ReadLines(targetFile)

	numValues := len(lines) - 1
	mathProblems := make([]MathProblem, 0)

	alignmentIdx := make([]int, 0)
	for idx, val := range lines[len(lines)-1] {
		if idx == 0 {
			continue
		}
		if val != ' ' {
			alignmentIdx = append(alignmentIdx, idx-1)
		}
	}
	alignmentIdx = append(alignmentIdx, len(lines[len(lines)-1]))

	for lineIdx, line := range lines {
		line += " " // Normalize

		currVal := ""
		currMathProblem := 0

		for idx, rune := range line {
			// Commit
			if slices.Contains(alignmentIdx, idx) {
				if lineIdx == 0 {
					mathProblems = append(mathProblems, MathProblem{})
				}
				if lineIdx == numValues {
					mathProblems[currMathProblem].Operation = strings.TrimSpace(currVal)
				} else {
					mathProblems[currMathProblem].addValue(currVal)
					mathProblems[currMathProblem]._Values = append(mathProblems[currMathProblem]._Values, currVal)
				}
				currVal = ""
				currMathProblem++
			} else {
				currVal += string(rune)
			}
		}
	}

	DEBUG(mathProblems)
	DEBUG("\n\n")

	var part01 int
	var part02 int

	for _, mp := range mathProblems {
		part01 += mp.calculate()
		part02 += mp.calculate2()
	}

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

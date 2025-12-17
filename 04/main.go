package day04

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
)

const Day = "04"

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

	var part01 int
	var part02 int

	fmt.Println(lines[0])

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

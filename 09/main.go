package day09

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const Day = "09"

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

type Rectangle struct {
	A       Point
	B       Point
	Product int
}

type Point shared.Tuple[int]

func Run(debug bool, targetFile string) (int, int) {
	_debug = debug

	lines := shared.ReadLines(targetFile)
	points := make([]Point, 0)
	for _, line := range lines {
		split := strings.Split(line, ",")
		column, err := strconv.Atoi(split[0])
		if err != nil {
			panic("column")
		}
		row, err := strconv.Atoi(split[1])
		if err != nil {
			panic("row")
		}
		points = append(points, Point{column, row})
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	product := func(a, b Point) int {
		return (abs(a.Second-b.Second) + 1) * (abs(a.First-b.First) + 1)
	}

	// n^2
	allPossiblePairs := make([]Rectangle, 0)
	for idxA := 0; idxA < len(points); idxA++ {
		a := points[idxA]
		for idxB := idxA + 1; idxB < len(points); idxB++ {
			b := points[idxB]
			rect := Rectangle{a, b, product(a, b)}
			// DEBUGF("%+v\n", rect)
			allPossiblePairs = append(allPossiblePairs, rect)
		}
	}

	slices.SortFunc(allPossiblePairs, func(a, b Rectangle) int {
		return b.Product - a.Product
	})

	var part01 int
	var part02 int

	part01 = allPossiblePairs[0].Product

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

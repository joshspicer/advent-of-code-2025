package day08

import (
	"advent-of-code-2025/shared"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const Day = "08"

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

type Point3D struct {
	x int
	y int
	z int
}

func Run(debug bool, targetFile string) (int, int) {
	_debug = debug

	lines := shared.ReadLines(targetFile)
	points := make([]Point3D, 0)
	for _, line := range lines {
		vals := strings.Split(line, ",")
		if len(vals) != 3 {
			panic("nooo")
		}
		x, err := strconv.Atoi(vals[0])
		if err != nil {
			panic("x no")
		}
		y, err := strconv.Atoi(vals[1])
		if err != nil {
			panic("y no")
		}
		z, err := strconv.Atoi(vals[2])
		if err != nil {
			panic("z no")
		}
		points = append(points, Point3D{x, y, z})
	}

	DEBUG(points)
	DEBUG("\n")

	var minimalPair shared.Tuple[Point3D]
	minimalDistance := math.MaxFloat64
	for idxA, a := range points {
		for idxB, b := range points {
			if idxA == idxB {
				continue
			}

			currDist := distance(a, b)
			// DEBUGF("%v and %v are %f apart\n", a, b, currDist)
			if currDist < minimalDistance {
				minimalPair = shared.Tuple[Point3D]{a, b}
				minimalDistance = currDist
			}
		}
	}
	DEBUG(minimalPair)
	DEBUGF("\n3d %f", distance(Point3D{0, 3, 0}, Point3D{4, 0, 0}))

	var part01 int
	var part02 int

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

func distance(a Point3D, b Point3D) float64 {
	X := math.Pow(float64(a.x-b.x), 2)
	Y := math.Pow(float64(a.y-b.y), 2)
	Z := math.Pow(float64(a.z-b.z), 2)
	return math.Pow(float64(X+Y+Z), 1.0/2.0)
}

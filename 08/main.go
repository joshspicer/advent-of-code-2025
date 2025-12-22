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
	isExample := strings.Contains(targetFile, "example")
	DEBUGF("isExample=%t\n", isExample)
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

	N := 1000
	if isExample {
		N = 10
	}

	less := func(p, q Point3D) bool {
		if p.x != q.x {
			return p.x < q.x
		}
		if p.y != q.y {
			return p.y < q.y
		}
		return p.z < q.z
	}

	normalizePair := func(a, b Point3D) shared.Tuple[Point3D] {
		if less(a, b) {
			return shared.Tuple[Point3D]{a, b}
		}
		return shared.Tuple[Point3D]{b, a}
	}

	graph := shared.MakeAdjacencyList[Point3D]()
	seen := make(map[shared.Tuple[Point3D]]bool, 0)
	// Build graph
	for range N {
		var minimalPair shared.Tuple[Point3D]
		minimalDistance := math.MaxFloat64
		found := false

		for idxA, a := range points {
			for idxB, b := range points {
				if idxA == idxB {
					continue
				}

				currPair := normalizePair(a, b)
				if seen[currPair] {
					continue
				}

				currDist := distance(a, b)
				if currDist < minimalDistance {
					minimalPair = currPair
					minimalDistance = currDist
					found = true
				}
			}
		}

		if !found {
			break // no unused edges left
		}

		seen[minimalPair] = true
		graph.AddEdge(minimalPair.First, minimalPair.Second, true)
	}

	part01 := 1
	var part02 int

	for idx, b := range graph.Components() {
		if idx > 2 { // Only three largest circuits
			break
		}
		DEBUG(len(b), b, "\n")
		part01 *= len(b)
	}

	fmt.Println(part01)
	fmt.Println(part02)

	return part01, part02
}

func distance(a Point3D, b Point3D) float64 {
	X := (a.x - b.x) * (a.x - b.x)
	Y := (a.y - b.y) * (a.y - b.y)
	Z := (a.z - b.z) * (a.z - b.z)
	return math.Pow(float64(X+Y+Z), 1.0/2.0)
}

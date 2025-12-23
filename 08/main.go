package day08

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"slices"
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

	graph := shared.MakeAdjacencyList[Point3D]()

	allPossiblePairs := make([]shared.Tuple[Point3D], 0)
	for idxA := 0; idxA < len(points); idxA++ {
		a := points[idxA]
		for idxB := idxA + 1; idxB < len(points); idxB++ {
			b := points[idxB]
			allPossiblePairs = append(allPossiblePairs, shared.Tuple[Point3D]{a, b})
		}
	}
	slices.SortFunc(allPossiblePairs, func(p, q shared.Tuple[Point3D]) int {
		Q := distance(q.First, q.Second)
		P := distance(p.First, p.Second)
		return P - Q
	})

	DEBUG(allPossiblePairs)

	part01 := 1
	var part02 int

	N := 1000
	if isExample {
		N = 10
	}

	for idx := range N {
		nextMinimalPair := allPossiblePairs[idx]
		graph.AddEdge(nextMinimalPair.First, nextMinimalPair.Second, true)
	}

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

func distance(a Point3D, b Point3D) int {
	X := (a.x - b.x) * (a.x - b.x)
	Y := (a.y - b.y) * (a.y - b.y)
	Z := (a.z - b.z) * (a.z - b.z)
	//return math.Pow(float64(X+Y+Z), 1.0/2.0)
	// Optimize by not taking the square root
	return X + Y + Z
}

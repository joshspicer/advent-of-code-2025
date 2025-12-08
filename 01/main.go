package day01

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"strconv"
)

const DAY = "01"

var _debug bool
var EXAMPLE = fmt.Sprintf("%s/%s", DAY, "example")
var PUZZLE = fmt.Sprintf("%s/%s", DAY, "puzzle")

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

func Run(debug bool, example bool) {
	_debug = debug
	targetFile := shared.If(example, EXAMPLE, PUZZLE)
	lines := shared.ReadLines(targetFile)
	state := state{50, 0}
	for _, l := range lines {
		dir, dist := parse(l)
		state.next(dir, dist)
		if state.position == 0 {
			state.stopsAtZero++
		}
	}
	part01(state)
	part02(state)
}

type Direction int

const (
	Left = iota
	Right
)

type state struct {
	position    int64
	stopsAtZero int64
}

func (s *state) next(direction Direction, distance int64) {
	DEBUG(s.position)
	DEBUG(" --> ")

	var val int64
	switch direction {
	case Left:
		val = (s.position + (100 - distance))

	case Right:
		val = (s.position + (100 + distance))
	}

	s.position = val % 100
	DEBUG(s.position, "\n")
}

func parse(line string) (Direction, int64) {
	var direction Direction
	rawDirection := line[0]
	switch rawDirection {
	case 'L':
		direction = Left
	case 'R':
		direction = Right
	default:
		panic("Invalid direction")
	}

	distance, err := strconv.ParseInt(line[1:], 10, 64)

	if err != nil {
		panic(err)
	}

	return direction, distance
}

func part01(state state) {
	fmt.Println(state.stopsAtZero)
}

func part02(state state) {
	fmt.Print("Part 2")
}

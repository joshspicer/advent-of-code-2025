package day01

import (
	"advent-of-code-2025/shared"
	"fmt"
	"os"
	"strconv"
)

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

func Run(debug bool, targetFile string) (int64, int64) {
	_debug = debug
	lines := shared.ReadLines(targetFile)
	state := state{50, 0, 0}
	for _, l := range lines {
		dir, dist := parse(l)
		state.next(dir, dist)
		if state.position == 0 {
			state.stopsAtZero++
		}
	}
	// Part 1
	fmt.Println(state.stopsAtZero)
	// Part 2
	fmt.Println(state.clicksPastZero)

	return state.stopsAtZero, state.clicksPastZero
}

type Direction string

const (
	Left  = "L"
	Right = "R"
)

type state struct {
	position       int64
	stopsAtZero    int64
	clicksPastZero int64
}

func (s *state) next(direction Direction, distance int64) {
	DEBUG(s.position)
	DEBUGF(" ---[%s%d]---> ", direction, distance)

	var val int64
	var oldPosition = s.position
	switch direction {
	case Left:
		val = (s.position + (100 - distance))
	case Right:
		val = (s.position + (100 + distance))
	}

	s.position = val % 100
	if s.position < 0 {
		s.position += 100
	}
	DEBUG(s.position)

	if oldPosition == 0 {
		if oldPosition+distance > 100 {
			count := (distance / 100)
			DEBUGF(" [*zero(%d)*] ", count)
			s.clicksPastZero += count
		}
	} else {
		switch direction {
		case Left:
			var count int64 = 0
			tmpDistance := distance
			for oldPosition-tmpDistance <= 0 {
				count += 1
				tmpDistance -= 100
			}
			DEBUGF(" [*left(%d)*] ", count)
			s.clicksPastZero += count
		case Right:
			var count int64 = 0
			tmpDistance := distance
			for oldPosition+tmpDistance >= 100 {
				count += 1
				tmpDistance -= 100
			}
			DEBUGF(" [*right(%d)*] ", count)
			s.clicksPastZero += count
		}
	}

	DEBUG("\n")
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

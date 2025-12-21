package shared

import (
	"fmt"
	"slices"
	"strings"
)

type Grid[T comparable] [][]T

func ToGrid[T comparable](lines []string, conv func(rune) T) Grid[T] {
	out := make(Grid[T], len(lines))
	for i, line := range lines {
		row := make([]T, 0, len(line))
		for _, ch := range line {
			row = append(row, conv(ch))
		}
		out[i] = row
	}
	return out
}

func (g Grid[T]) ForEach(fn func(row, col int, v T)) {
	for r := 0; r < g.Height(); r++ {
		for c := 0; c < g.Width(); c++ {
			fn(r, c, g[r][c])
		}
	}
}

func (g Grid[T]) Copy() Grid[T] {
	out := make(Grid[T], g.Height())
	for r := range g {
		out[r] = append([]T(nil), g[r]...)
	}
	return out
}

// Returns if the mutation occurred or not
func (g Grid[T]) MutateIgnoringBounds(row, col int, value T) bool {
	if !g.InBounds(row, col) {
		return false
	}
	g.Mutate(row, col, value)
	return true
}

func (g Grid[T]) Mutate(row, col int, value T) {
	g[row][col] = value
}

var EightAdjacentRelativeOffsets = []Tuple[int]{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1} /*{0, 0}*/, {0, 1},
	{1, -1}, {1, 0}, {1, 1}}

var CardinalOffsets = []Tuple[int]{
	{-1, 0},
	{0, -1} /*{0, 0}*/, {0, 1},
	{1, 0}}

// Returns the adjacent cells to the given row/col
// according to the provided 'relativeOffsets'
// If a cell doesn't exist (out of bounds), it is omitted
func (g Grid[T]) CollectAdjacent(row, col int, relativeOffsets []Tuple[int]) []T {
	var adjacents = make([]T, 0)
	for _, offset := range relativeOffsets {
		target, err := g.At(row+offset.First, col+offset.Second)
		if err != nil {
			// Skip
			continue
		}
		adjacents = append(adjacents, target)
	}
	return adjacents
}

// For a given (row,col) and a symbol
// Return a list of coordinates where expected == the value at that coordinate
func (g Grid[T]) CollectAdjacentIndices(row, col int, expected []T, relativeOffsets []Tuple[int]) []Tuple[int] {
	var indices = make([]Tuple[int], 0)
	for _, offset := range relativeOffsets {
		target, err := g.At(row+offset.First, col+offset.Second)
		if err != nil || !slices.Contains(expected, target) {
			// Skip
			continue
		}
		indices = append(indices, Tuple[int]{row + offset.First, col + offset.Second})
	}
	return indices
}

func (g Grid[T]) At(row, col int) (T, error) {
	if !g.InBounds(row, col) {
		return getGenericNull[T](), fmt.Errorf("BAD! (%d,%d) is out of bounds: maxX=%d maxY=%d", row, col, g.Height()-1, g.Width()-1)
	}
	return g[row][col], nil
}

func getGenericNull[T any]() T {
	var result T
	return result
}

func (g Grid[T]) Height() int {
	return len(g)
}

func (g Grid[T]) Width() int {
	return len(g[0])
}

func (g Grid[T]) InBounds(row, col int) bool {
	return row >= 0 && row < g.Height() &&
		col >= 0 && col < g.Width()
}

func (g Grid[T]) String() string {
	var b strings.Builder
	for r := 0; r < g.Height(); r++ {
		for c := 0; c < g.Width(); c++ {
			fmt.Fprint(&b, g[r][c])
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (g Grid[T]) Above(row, col int) (T, error) {
	return g.At(row-1, col)
}

func (g Grid[T]) Below(row, col int) (T, error) {
	return g.At(row+1, col)
}

func (g Grid[T]) Left(row, col int) (T, error) {
	return g.At(row, col-1)
}

func (g Grid[T]) Right(row, col int) (T, error) {
	return g.At(row, col+1)
}

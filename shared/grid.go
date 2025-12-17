package shared

import (
	"fmt"
	"strings"
)

type Grid[T comparable] [][]T

var eightAdjacent = []Tuple[int]{
	{-1, -1}, //...
}

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

func (g Grid[T]) collectEightAdjacent(row, col int) {

}

func (g Grid[T]) At(row, col int) T {
	if !g.InBounds(row, col) {
		panic(fmt.Sprintf("BAD! (%d,%d) is out of bounds: maxX=%d maxY=%d", row, col, g.Height()-1, g.Width()-1))
	}
	return g[row][col]
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

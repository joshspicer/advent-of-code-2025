package shared

import (
	"fmt"
	"slices"
	"strings"
)

type DFlags uint8

const (
	None  DFlags = 0
	Start DFlags = 1 << iota
	End
)

type Node = Tuple[int]

type DirectedAdjacencyList struct {
	Edges map[Node]*Set[Node]
	Flags map[Node]DFlags
}

// Processes a 2D Grid serially
// Applys fn() to each cell, returning its directed adjacency list.
func ToDAG[T comparable](input Grid[T], fn func(r, c int) ([]Node, DFlags)) DirectedAdjacencyList {
	adjList := DirectedAdjacencyList{Edges: make(map[Node]*Set[Node]), Flags: make(map[Node]DFlags)}
	input.ForEach(func(row, col int, v T) {
		from := Tuple[int]{row, col}
		toEdges, flags := fn(row, col)
		if len(toEdges) == 0 && flags == None {
			// Boring
			return
		}
		adjList.Flags[from] = flags
		if adjList.Edges[from] == nil {
			adjList.Edges[from] = CreateSet[Tuple[int]]()
		}
		for _, to := range toEdges {
			adjList.Edges[from].Add(to)
			if adjList.Edges[to] == nil {
				adjList.Edges[to] = CreateSet[Node]()
			}
		}
	})
	return adjList
}

func (adjList DirectedAdjacencyList) WithFlags(flags DFlags) []Node {
	result := make([]Node, 0)
	for node, v := range adjList.Flags {
		if v&flags != 0 {
			result = append(result, node)
		}
	}
	return result
}

func (adjList DirectedAdjacencyList) dfsImpl(curr, end Node, path []Node, visited *Set[Node], allPaths *[][]Node) {
	if visited.Contains(curr) {
		return
	}

	visited.Add(curr)
	defer visited.Remove(curr) // support backtracking

	path = append(path, curr)

	if curr == end {
		*allPaths = append(*allPaths, slices.Clone(path))
		return
	}

	// iterate my neighbors
	adjList.Edges[curr].ForEach(func(adj Node) {
		adjList.dfsImpl(adj, end, path, visited, allPaths)
	})
}

func (adjList DirectedAdjacencyList) AllPaths(start, end Node) [][]Node {
	var allPaths [][]Node
	adjList.dfsImpl(start, end, nil, CreateSet[Node](), &allPaths)
	return allPaths
}

func (adjList DirectedAdjacencyList) dfCount(curr Node, memo map[Node]uint64) uint64 {
	if memoized, ok := memo[curr]; ok {
		return memoized // reuse path we already deemed from curr to end
	}
	if adjList.Flags[curr]&DFlags(End) != 0 {
		memo[curr] = 1
		return 1 // base case, if we're at an end node there is one path to it
	}
	adjList.Edges[curr].ForEach(func(adj Node) {
		memo[curr] += adjList.dfCount(adj, memo) // calculate sum of paths over the neighbors
	})
	return memo[curr]
}

func (adjList DirectedAdjacencyList) CountPaths(start Node) uint64 {
	return adjList.dfCount(start, make(map[Node]uint64, 0))
}

func (adjList DirectedAdjacencyList) String() string {
	var builder strings.Builder
	for idx, x := range adjList.Edges {
		_, err := builder.WriteString(fmt.Sprintf("[%d, %d] => %v (flags=%d)\n", idx.First, idx.Second, x, adjList.Flags[idx]))
		if err != nil {
			panic("woooah")
		}
	}
	return builder.String()
}

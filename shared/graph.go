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

type AdjacencyList[N comparable] struct {
	Edges map[N]*Set[N]
	Flags map[N]DFlags
}

func (adjList *AdjacencyList[N]) ensureNode(n N) {
	if adjList.Edges[n] == nil {
		adjList.Edges[n] = CreateSet[N]()
	}
}

func (adjList *AdjacencyList[N]) AddEdge(from, to N, undirected bool) {
	adjList.ensureNode(from)
	adjList.Edges[from].Add(to)
	if undirected {
		adjList.AddEdge(to, from, false)
	}
}

func MakeAdjacencyList[N comparable]() AdjacencyList[N] {
	return AdjacencyList[N]{Edges: make(map[N]*Set[N]), Flags: make(map[N]DFlags)}
}

type point2D = Tuple[int]

// Processes a 2D Grid serially
// Applys fn() to each cell, returning its directed adjacency list.
// REMARKS:
//   - Not generalized over AdjacencyList
//   - Must give it an input grid that results in a DAG
func ToDAG[T comparable](input Grid[T], fn func(r, c int) ([]point2D, DFlags)) AdjacencyList[point2D] {
	adjList := MakeAdjacencyList[point2D]()
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
				adjList.Edges[to] = CreateSet[point2D]()
			}
		}
	})
	return adjList
}

func (adjList AdjacencyList[N]) WithFlags(flags DFlags) []N {
	result := make([]N, 0)
	for node, v := range adjList.Flags {
		if v&flags != 0 {
			result = append(result, node)
		}
	}
	return result
}

// Expects a *DAG*
func (adjList AdjacencyList[N]) dfsImpl(curr, end N, path []N, visited *Set[N], allPaths *[][]N) {
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
	adjList.Edges[curr].ForEach(func(adj N) {
		adjList.dfsImpl(adj, end, path, visited, allPaths)
	})
}

// Expects a *DAG*
func (adjList AdjacencyList[N]) AllPaths(start, end N) [][]N {
	var allPaths [][]N
	adjList.dfsImpl(start, end, nil, CreateSet[N](), &allPaths)
	return allPaths
}

// Expects a *DAG*
func (adjList AdjacencyList[N]) dfCount(curr N, memo map[N]uint64) uint64 {
	if memoized, ok := memo[curr]; ok {
		return memoized // reuse path we already deemed from curr to end
	}
	if adjList.Flags[curr]&DFlags(End) != 0 {
		memo[curr] = 1
		return 1 // base case, if we're at an end node there is one path to it
	}
	adjList.Edges[curr].ForEach(func(adj N) {
		memo[curr] += adjList.dfCount(adj, memo) // calculate sum of paths over the neighbors
	})
	return memo[curr]
}

// Expects a *DAG*
func (adjList AdjacencyList[N]) CountPaths(start N) uint64 {
	return adjList.dfCount(start, make(map[N]uint64, 0))
}

func (adjList AdjacencyList[N]) String() string {
	var builder strings.Builder
	for idx, x := range adjList.Edges {
		_, err := builder.WriteString(fmt.Sprintf("%v => %v (flags=%d)\n", idx, x, adjList.Flags[idx]))
		if err != nil {
			panic("woooah")
		}
	}
	return builder.String()
}

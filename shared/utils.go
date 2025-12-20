package shared

import (
	"os"
	"strings"
)

type Tuple[T any] struct {
	First  T
	Second T
}

func ReadLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func ReadCommaSeparatedLine(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), ",")
}

// "Buckets" a string (array of chars) into N buckets.
// 123456  --> [12], [34], [56]
// 1234567 --> [12], [34], [567]
//
// REMARKS:
//   - Last bucket might not equal the rest
func BucketString(input string, buckets int) []string {
	var output []string = make([]string, buckets)
	counter := 0
	targetBucket := 0
	itemsPerBucket := len(input) / buckets
	for counter != len(input) {
		output[targetBucket] = output[targetBucket] + string(input[counter])
		counter++
		if counter%itemsPerBucket == 0 && targetBucket < buckets-1 {
			targetBucket++
		}
	}
	return output
}

// True if all elements pass '==' check with each other
func AllEqual[T comparable](items ...T) bool {
	if len(items) < 2 {
		return true
	}
	for _, v := range items[1:] {
		if items[0] != v {
			return false
		}
	}
	return true
}

func ParseDashedTuple(items []string) []Tuple[string] {
	result := make([]Tuple[string], 0, len(items))
	for _, item := range items {
		raw := strings.Split(item, "-")
		if len(raw) != 2 {
			panic("Unexpected in ParseDashedTuple")
		}
		result = append(result, Tuple[string]{raw[0], raw[1]})
	}
	return result
}

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Map[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func Max(a, b int) int {
	// Ignore all the good special case handling
	if a > b {
		return a
	}
	return b
}

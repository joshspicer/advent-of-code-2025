package shared

import (
	"os"
	"strings"
)

func ReadLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func If[T any](cond bool, vtrue, vfalse T) T {
    if cond {
        return vtrue
    }
    return vfalse
}
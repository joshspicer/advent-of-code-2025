package main

import (
	day01 "advent-of-code-2025/01"
	"fmt"
	"os"
)

func main() {
	day := os.Getenv("DAY")
	debug := os.Getenv("DEBUG") != ""
	fileName := os.Getenv("QUALITY")
	if day == "" {
		panic("Missing DAY env variable")
	}
	switch day {
	case "01":
		day01.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	}
}

package main

import (
	day01 "advent-of-code-2025/01"
	"os"
)

func main() {
	day := os.Getenv("DAY")
	debug := os.Getenv("DEBUG") != ""
	example := os.Getenv("QUALITY") == "example"
	if day == "" {
		panic("Missing DAY env variable")
	}
	switch day {
	case "01":
		day01.Run(debug, example)
	}
}

package main

import (
	day01 "advent-of-code-2025/01"
	day02 "advent-of-code-2025/02"
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
	case "02":
		day02.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	default:
		err := fmt.Errorf("Missing day %s!", day)
		panic(err)
	}
}

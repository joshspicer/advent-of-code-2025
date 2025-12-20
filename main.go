package main

import (
	day01 "advent-of-code-2025/01"
	day02 "advent-of-code-2025/02"
	day03 "advent-of-code-2025/03"
	day04 "advent-of-code-2025/04"
	day05 "advent-of-code-2025/05"
	day06 "advent-of-code-2025/06"
	day07 "advent-of-code-2025/07"
	day08 "advent-of-code-2025/08"
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
	case "03":
		day03.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	case "04":
		day04.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	case "05":
		day05.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	case "06":
		day06.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	case "07":
		day07.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	case "08":
		day08.Run(debug, fmt.Sprintf("%s/%s", day, fileName))
	default:
		err := fmt.Errorf("missing day %s", day)
		panic(err)
	}
}

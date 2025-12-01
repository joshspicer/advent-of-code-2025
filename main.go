package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	day := time.Now().Day()
	if len(os.Args) > 1 {
		if d, err := strconv.Atoi(os.Args[1]); err == nil {
			day = d
		}
	}	
	dayDir := fmt.Sprintf("%02d", day)
	cmd := exec.Command("go", "run", dayDir+"/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Day %02d not yet implemented or failed: %v\n", day, err)
	}
}

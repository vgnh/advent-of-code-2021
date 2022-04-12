package main

import (
	"advent-of-code-2021/days"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	days.Day01()

	fmt.Printf("\nTime taken: %v\n", time.Since(start))
}

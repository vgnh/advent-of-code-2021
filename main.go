package main

import (
	"advent-of-code-2021/day01"
	"advent-of-code-2021/day02"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day01.Main()
	day02.Main()

	fmt.Printf("\nTime taken: %v\n", time.Since(start))
}

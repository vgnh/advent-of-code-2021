package main

import (
	"advent-of-code-2021/day01"
	"advent-of-code-2021/day02"
	"advent-of-code-2021/day03"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day01.Main()
	day02.Main()
	day03.Main()

	fmt.Printf("\nTime taken: %v\n", time.Since(start))
}

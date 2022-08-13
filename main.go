package main

import (
	"advent-of-code-2021/day01"
	"advent-of-code-2021/day02"
	"advent-of-code-2021/day03"
	"advent-of-code-2021/day04"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day01.Main()
	day02.Main()
	day03.Main()
	day04.Main()

	fmt.Printf("\nTime elapsed: %v\n", time.Since(start))
}

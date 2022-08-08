package day01

import (
	"advent-of-code-2021/utils"
	"fmt"
)

const filename = "./day01/input.txt"

var input = utils.MapToInt(utils.ReadLines(filename))

func part01() int {
	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] > input[i] {
			count++
		}
	}
	return count
}

func part02() int {
	count := 0
	for i := 0; i < len(input)-3; i++ {
		if utils.Sum(input[i+1:i+4]) > utils.Sum(input[i:i+3]) {
			count++
		}
	}
	return count
}

func Main() {
	fmt.Println("Advent of Code 2021, Day 01")
	fmt.Println(part01())
	fmt.Println(part02())
}

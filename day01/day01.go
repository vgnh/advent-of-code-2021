package day01

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
)

const FILENAME string = "./inputs/day01.txt"

func input() []int {
	strings := utils.ReadLines(FILENAME)

	var inputs []int
	for _, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		inputs = append(inputs, num)
	}

	return inputs
}

func part01() int {
	input := input()
	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] > input[i] {
			count++
		}
	}
	return count
}

func part02() int {
	input := input()
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

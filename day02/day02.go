package day02

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

const FILENAME string = "./inputs/day02.txt"

func instructions() []string {
	return utils.ReadLines(FILENAME)
}

func part01() int {
	horizontal, depth := 0, 0
	for _, ins := range instructions() {
		command := strings.Split(strings.TrimSpace(ins), " ")
		value, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	return horizontal * depth
}

func part02() int {
	horizontal, depth, aim := 0, 0, 0
	for _, ins := range instructions() {
		command := strings.Split(strings.TrimSpace(ins), " ")
		value, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	return horizontal * depth
}

func Main() {
	fmt.Println("Advent of Code 2021, Day 02")
	fmt.Println(part01())
	fmt.Println(part02())
}

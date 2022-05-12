package day04

import (
	"advent-of-code-2021/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const FILENAME string = "./inputs/day04.txt"

func input() []string {
	return strings.Split(utils.ReadString(FILENAME), "\n\n")
}

func bingoNumbers() []int {
	strs := strings.Split(strings.TrimSpace(input()[0]), ",")
	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}
	return nums
}

func allBoards() [][][]int {
	input := input()
	allBoards := make([][][]int, len(input[1:]))

	for i, inputString := range input[1:] {
		strs := strings.Split(inputString, "\n")
		board := make([][]int, len(strs))

		for j, line := range strs {
			reg := regexp.MustCompile(`\s+`)
			strArray := reg.Split(strings.TrimSpace(line), -1)

			intArray := make([]int, len(strArray))
			for k, s := range strArray {
				intArray[k], _ = strconv.Atoi(s)
			}

			board[j] = intArray
		}

		allBoards[i] = board
	}

	return allBoards
}

func checkWin(board [][]int) bool {
	// Check rows
	for _, row := range board {
		if utils.Sum(row) == -5 {
			return true
		}
	}

	// Check columns
	for j := range board[0] {
		allTrueInCol := true
		for _, ints := range board {
			if ints[j] != -1 {
				allTrueInCol = false
				break
			}
		}
		if allTrueInCol {
			return true
		}
	}

	return false
}

func markOnAllBoards(n int, boards [][][]int) {
	for _, board := range boards {
		for i := range board {
			for j := range board[0] {
				if board[i][j] == n {
					board[i][j] = -1
				}
			}
		}
	}
}

func sumOfUnmarked(board [][]int) int {
	sum := 0
	for _, ints := range board {
		for j := range board[0] {
			if ints[j] != -1 {
				sum += ints[j]
			}
		}
	}
	return sum
}

func part01() int {
	boards := allBoards()
	for _, num := range bingoNumbers() {
		markOnAllBoards(num, boards)
		for _, board := range boards {
			if checkWin(board) {
				return sumOfUnmarked(board) * num
			}
		}
	}
	return -1
}

func sliceRemoveAll(fromSlice [][][]int, indexes []int) [][][]int {
	var newSlice [][][]int
	for i, item := range fromSlice {
		if utils.SliceContains(indexes, i) {
			continue
		}
		newSlice = append(newSlice, item)
	}
	return newSlice
}

func part02() int {
	boards := allBoards()
	for _, num := range bingoNumbers() {
		markOnAllBoards(num, boards)
		if len(boards) != 1 {
			var boardsToRemove []int
			for i, board := range boards {
				if checkWin(board) {
					boardsToRemove = append(boardsToRemove, i)
				}
			}
			boards = sliceRemoveAll(boards, boardsToRemove)
		} else {
			if checkWin(boards[0]) {
				return sumOfUnmarked(boards[0]) * num
			}
		}
	}
	return -1
}

func Main() {
	fmt.Println("Advent of Code 2021, Day 04")
	fmt.Println(part01())
	fmt.Println(part02())
}

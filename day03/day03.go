package day03

import (
	"advent-of-code-2021/utils"
	"fmt"
	"strconv"
	"strings"
)

const filename string = "./day03/input.txt"

var binaryNumbers = utils.ReadLines(filename)

func part01() int64 {
	gamma, epsilon := "", ""

	newBinaries := make([]string, len(binaryNumbers[0]))
	for _, bin := range binaryNumbers {
		for i := 0; i < len(bin); i++ {
			newBinaries[i] += string(bin[i])
		}
	}

	for _, newBin := range newBinaries {
		if strings.Count(newBin, "0") > strings.Count(newBin, "1") {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 32)
	ep, _ := strconv.ParseInt(epsilon, 2, 32)
	return g * ep
}

func filterOnCharPosition(array []string, position int, character string) []string {
	var filteredArray []string
	for _, s := range array {
		if string(s[position]) == character {
			filteredArray = append(filteredArray, s)
		}
	}
	return filteredArray
}

func part02() int64 {
	generator, scrubber := int64(0), int64(0)
	genList, scrubList := binaryNumbers, binaryNumbers

	for i := 0; i < len(binaryNumbers[0]); i++ {
		binGenerator := make([]string, len(binaryNumbers[0]))
		for _, bin := range genList {
			for j := 0; j < len(bin); j++ {
				binGenerator[j] += string(bin[j])
			}
		}

		binScrubber := make([]string, len(binaryNumbers[0]))
		for _, bin := range scrubList {
			for j := 0; j < len(bin); j++ {
				binScrubber[j] += string(bin[j])
			}
		}

		if strings.Count(binGenerator[i], "0") > strings.Count(binGenerator[i], "1") {
			genList = filterOnCharPosition(genList, i, "0")
		} else {
			genList = filterOnCharPosition(genList, i, "1")
		}

		if strings.Count(binScrubber[i], "0") > strings.Count(binScrubber[i], "1") {
			scrubList = filterOnCharPosition(scrubList, i, "1")
		} else {
			scrubList = filterOnCharPosition(scrubList, i, "0")
		}

		if len(genList) == 1 {
			generator, _ = strconv.ParseInt(genList[0], 2, 32)
		}
		if len(scrubList) == 1 {
			scrubber, _ = strconv.ParseInt(scrubList[0], 2, 32)
		}
	}
	return generator * scrubber
}

func Main() {
	fmt.Println("Advent of Code 2021, Day 03")
	fmt.Println(part01())
	fmt.Println(part02())
}

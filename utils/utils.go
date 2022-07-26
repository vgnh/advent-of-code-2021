package utils

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	// "os"
	"strings"
)

func ReadLines(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
	/* file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines */
}

func ReadString(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sum[T Number](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}

func SliceContains[T comparable](slice []T, valueToSearch T) bool {
	for _, value := range slice {
		if value == valueToSearch {
			return true
		}
	}
	return false
}

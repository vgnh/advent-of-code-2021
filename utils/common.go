package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	strs := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i := range strs {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strs
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
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(data))
}

func MapToInt(strs []string) []int {
	ints := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		ints[i] = num
	}
	return ints
}

func MapToStr(ints []int) []string {
	strs := make([]string, len(ints))
	for i, num := range ints {
		strs[i] = strconv.Itoa(num)
	}
	return strs
}

func Min(slice []int) int {
	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func InsertAt[T any](slice []T, index int, value T) []T {
	if len(slice) == index {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func Sum[T int | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func Count[T comparable](slice []T, item T) int {
	count := 0
	for _, v := range slice {
		if v == item {
			count++
		}
	}
	return count
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

func Copy2d[T any](matrix [][]T) [][]T {
	dup := make([][]T, len(matrix))
	for i := range matrix {
		dup[i] = make([]T, len(matrix[i]))
		copy(dup[i], matrix[i])
	}
	return dup
}

func IndexOf[T comparable](slice []T, item T) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

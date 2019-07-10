package day2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// GetResult returns the result for Advent of Code Day 2
func GetResult(part string) int {

	f, _ := os.Open("day2/day2.input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		stringNumbers := strings.Split(line, "\t")
		numbers := []int{}

		for _, v := range stringNumbers {
			number, _ := strconv.Atoi(v)
			numbers = append(numbers, number)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		fmt.Println(numbers)

	OuterLoop:
		for i := 0; i < len(numbers); i++ {
			dividend := numbers[i]
			for j := i + 1; j < len(numbers); j++ {
				divisor := numbers[j]
				if dividend%divisor == 0 {
					result += dividend / divisor
					break OuterLoop
				}
			}
		}

	}

	return result
}

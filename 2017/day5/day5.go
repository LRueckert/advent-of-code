package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// GetResult returns the result for Advent of Code Day 5 b
func GetResult(part string) int {
	f, _ := os.Open("day5/day5.input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var instructions []int

	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		instructions = append(instructions, number)

	}
	fmt.Println(instructions)

	walker := 0
	result := 0

	for 0 <= walker && walker < len(instructions) {
		tmp := instructions[walker]
		if tmp > 2 {
			instructions[walker]--
		} else {
			instructions[walker]++
		}
		walker += tmp
		result++
	}

	return result
}

// // GetResult returns the result for Advent of Code Day 5 a
// func GetResult() int {
// 	f, _ := os.Open("day5/day5.input")
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	scanner.Split(bufio.ScanLines)

// 	var instructions []int

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		number, _ := strconv.Atoi(line)
// 		instructions = append(instructions, number)

// 	}
// 	fmt.Println(instructions)

// 	walker := 0
// 	result := 0

// 	for 0 <= walker && walker < len(instructions) {
// 		tmp := walker
// 		walker += instructions[walker]
// 		instructions[tmp]++
// 		result++
// 	}

// 	return result
// }

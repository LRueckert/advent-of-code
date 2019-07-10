package day1

import (
	"bufio"
	"os"
	"strconv"
)

var file string

// GetResult returns the result for Advent of Code Day 1
func GetResult(part string) int {

	input := []string{}

	if file == "" {
		file = "day1/day1.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return calculateResult(part, input)
}

func calculateResult(part string, input []string) int {
	firstPart := part == "A"
	secondPart := part == "B"
	result := 0
	looping := true
	var resultMap map[int]bool

	if secondPart {
		resultMap = make(map[int]bool)
		resultMap[0] = true
	}

	for looping {
		if firstPart {
			looping = false
		}
		for _, line := range input {
			value, _ := strconv.Atoi(line)
			result += value
			if secondPart {
				if resultMap[result] {
					looping = false
					return result
				}
				resultMap[result] = true
			}
		}
	}

	return result
}

package day1

import (
	"io/ioutil"
	"strconv"
)

const test = "12131415"

// GetResult returns the result for Advent of Code Day 1
func GetResult(part string) int {

	dat, _ := ioutil.ReadFile("day1/day1.input")

	use := dat
	result := 0
	offset := len(use) / 2
	for i := 0; i < len(use); i++ {
		compare := i + offset
		if compare >= len(use) {
			compare = i - offset
		}
		if use[i] == use[compare] {
			number, _ := strconv.Atoi(string(use[i]))
			result += number
		}
	}
	return result
}

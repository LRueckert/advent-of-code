package day5

import (
	"io/ioutil"
	"math"
	"strings"
	"unicode"
)

var file string

// GetResult returns the result for Advent of Code Day 5
func GetResult(part string) int {

	firstPart := part == "A"

	if file == "" {
		file = "day5/day5.input"
	}

	bytes, _ := ioutil.ReadFile(file)
	input := strings.TrimSpace(string(bytes))

	if firstPart {
		return calculateResultA(part, input)
	}

	return calculateResultB(part, input)
}

func calculateResultA(part string, input string) int {

	return react(input)
}

func calculateResultB(part string, input string) int {
	lowest := math.MaxInt64
	for i := 'A'; i <= 'Z'; i++ {
		removedUpper := strings.ReplaceAll(input, string(i), "")
		removed := strings.ReplaceAll(removedUpper, string(unicode.ToLower(i)), "")
		result := react(removed)
		if result < lowest {
			lowest = result
		}
	}
	return lowest

}

func react(input string) int {
	for i := 1; i < len(input); i++ {
		runeA := rune(input[i-1])
		runeB := rune(input[i])

		if runeA != runeB && (unicode.ToLower(runeA) == runeB || runeA == unicode.ToLower(runeB)) {
			return react(input[:i-1] + input[i+1:])
		}
	}
	return len(input)
}

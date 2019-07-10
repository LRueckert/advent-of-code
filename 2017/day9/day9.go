package day9

import "io/ioutil"

var file string

// GetResult returns the result for Advent of Code Day 9
func GetResult(part string) int {
	bytes, _ := ioutil.ReadFile("day9/day9.input")
	input := string(bytes)

	return calculateResult(part, input)
}

func calculateResult(part string, input string) int {
	firstPart := part == "A"
	secondPart := part == "B"
	result := 0
	currentValue := 0

	ignore := false

	for i := 0; i < len(input); i++ {
		char := string(input[i])

		if ignore {
			switch char {
			case ">":
				ignore = false
			case "!":
				i++
			default:
				if secondPart {
					result++
				}
			}
		} else {
			switch char {
			case "{":
				currentValue++
			case "}":
				if firstPart {
					result += currentValue
				}
				currentValue--
			case "<":
				ignore = true
			case "!":
				i++
			}
		}
	}

	return result
}

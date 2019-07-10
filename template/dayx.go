package dayx

import (
	"bufio"
	"os"
)

var file string

// GetResult returns the result for Advent of Code Day x
func GetResult(part string) int {

	input := []string{}

	firstPart := part == "A"

	if file == "" {
		file = "dayx/dayx.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if firstPart {
		return calculateResultA(part, input)
	}

	return calculateResultB(part, input)
}

func calculateResultA(part string, input []string) int {

	result := 0

	return result
}

func calculateResultB(part string, input []string) int {

	result := 0

	return result

}

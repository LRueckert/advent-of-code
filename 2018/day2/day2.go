package day2

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var file string

// GetResult returns the result for Advent of Code Day 2
func GetResult(part string) string {

	input := []string{}

	firstPart := part == "A"

	if file == "" {
		file = "day2/day2.input"
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
		return strconv.Itoa(calculateResultA(part, input))
	}

	return calculateResultB(part, input)
}

func calculateResultA(part string, input []string) int {

	doubles := 0
	triples := 0

	for _, line := range input {
		letterMap := make(map[rune]int)
		for _, char := range line {
			letterMap[char]++
		}
		hasDouble := false
		hasTriple := false
		for _, count := range letterMap {
			if count == 2 {
				hasDouble = true
			}
			if count == 3 {
				hasTriple = true
			}
		}
		if hasDouble {
			doubles++
		}
		if hasTriple {
			triples++
		}
	}

	return doubles * triples
}

func calculateResultB(part string, input []string) string {
	sort.Strings(input)
	var result string

	for index, one := range input {
		difference := differentCharacter(one, input[index+1])
		if difference != -1 {
			result = one[:difference] + one[difference+1:]
			return result
		}
	}

	return result
}

func differentCharacter(one string, two string) int {
	difference := -1
	for index, runeA := range one {
		runeB := []rune(two)[index]
		if runeA != runeB {
			if difference != -1 {
				return -1
			}
			difference = index
		}
	}
	return difference
}

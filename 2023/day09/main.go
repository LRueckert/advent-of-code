package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file string

type Sequence struct {
	Numbers []int
}

func (s *Sequence) Predict(forwards bool) int {
	var relevantNum int
	if forwards {
		relevantNum = s.Numbers[len(s.Numbers)-1]
	} else {
		relevantNum = s.Numbers[0]
	}
	sub := Sequence{make([]int, len(s.Numbers)-1)}
	done := true
	for i := 1; i < len(s.Numbers); i++ {
		diff := s.Numbers[i] - s.Numbers[i-1]
		sub.Numbers[i-1] = diff
		if diff != 0 {
			done = false
		}
	}
	if done {
		return relevantNum
	}
	previousDiff := sub.Predict(forwards)
	if forwards {
		return previousDiff + relevantNum
	} else {
		return relevantNum - previousDiff
	}
}

func calculateResultA(input []string) int {
	result := 0
	sequences := ParseSequences(input)
	for _, s := range sequences {
		result += s.Predict(true)
	}

	return result
}

func calculateResultB(input []string) int {
	result := 0
	sequences := ParseSequences(input)
	for _, s := range sequences {
		result += s.Predict(false)
	}

	return result
}

func ParseSequences(input []string) []Sequence {
	sequences := make([]Sequence, len(input))
	for i, line := range input {
		numbers := strings.Split(line, " ")
		sequences[i] = Sequence{make([]int, len(numbers))}
		for j := range numbers {
			number, _ := strconv.Atoi(numbers[j])
			sequences[i].Numbers[j] = number
		}
	}
	return sequences
}

func getResult(part string) int {
	input := getInput()
	firstPart := part == "A"

	if firstPart {
		return calculateResultA(input)
	}

	return calculateResultB(input)
}

func getInput() []string {
	input := []string{}

	if file == "" {
		file = "input.txt"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}

func main() {
	argsWithProg := os.Args

	var part string
	if len(argsWithProg) < 2 {
		part = "A"
	} else {
		part = argsWithProg[1]
	}

	fmt.Println(getResult(part))
}

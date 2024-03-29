package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var file string

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

func getResult(part string) int {

	input := []string{}

	if file == "" {
		file = "input"
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

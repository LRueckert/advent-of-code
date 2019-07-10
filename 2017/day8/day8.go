package day8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var file string

// GetResult returns the result for Advent of Code Day 8
func GetResult(part string) int {
	firstPart := part == "A"
	secondPart := part == "B"
	if file == "" {
		file = "day8/day8.input"
	}
	f, _ := os.Open(file)
	defer f.Close()
	var result int
	registers := make(map[string]int)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if _, ok := registers[parts[0]]; !ok {
			registers[parts[0]] = 0
		}
		if _, ok := registers[parts[4]]; !ok {
			registers[parts[4]] = 0
		}

		if evaluate(registers[parts[4]], parts[5], parts[6]) {
			value, _ := strconv.Atoi(parts[2])
			switch parts[1] {
			case "inc":
				registers[parts[0]] += value
			case "dec":
				registers[parts[0]] -= value
			}
			if secondPart && registers[parts[0]] > result {
				result = registers[parts[0]]
			}
		}
	}

	if firstPart {
		for _, register := range registers {
			if register > result {
				result = register
			}
		}
	}

	return result
}

func evaluate(register int, operand string, valueString string) bool {
	value, _ := strconv.Atoi(valueString)
	switch operand {
	case ">":
		return register > value
	case "<":
		return register < value
	case "==":
		return register == value
	case ">=":
		return register >= value
	case "<=":
		return register <= value
	case "!=":
		return register != value
	default:
		return false
	}
}

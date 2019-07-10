package day8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var file string

// GetResult returns the result for Advent of Code Day x
func GetResult(part string) int {

	input := []int{}

	firstPart := part == "A"

	if file == "" {
		file = "day8/day8.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, value := range parts {
			value, _ := strconv.Atoi(value)
			input = append(input, value)
		}
	}

	if firstPart {
		return calculateResultA(part, input)
	}

	return calculateResultB(part, input)
}

func calculateResultA(part string, input []int) int {

	result, _ := getNodeInfoA(input)

	return result
}

func calculateResultB(part string, input []int) int {

	result, _ := getNodeInfoB(input)

	return result

}

func getNodeInfoA(input []int) (sum, length int) {
	length = 2
	childCount := input[0]
	metaDataCount := input[1]
	for i := 0; i < childCount; i++ {
		nodeSum, nodeLength := getNodeInfoA(input[length:])
		sum += nodeSum
		length += nodeLength
	}
	for i := 0; i < metaDataCount; i++ {
		sum += input[length]
		length++
	}

	return sum, length
}

func getNodeInfoB(input []int) (sum, length int) {
	length = 2
	childCount := input[0]
	children := make([]int, childCount)
	metaDataCount := input[1]
	for i := 0; i < childCount; i++ {
		nodeSum, nodeLength := getNodeInfoB(input[length:])
		children[i] = nodeSum
		length += nodeLength
	}
	for i := 0; i < metaDataCount; i++ {
		childIndex := input[length] - 1
		if childCount > 0 {
			if childIndex < len(children) {
				sum += children[childIndex]
			}
		} else {
			sum += input[length]
		}
		length++
	}

	return sum, length
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	firstPart := part == "A"

	if file == "" {
		file = "day3.input"
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

	clothMap := make(map[int]map[int]int)
	result := 0

	for _, line := range input {
		idpart := strings.Split(line, "@")
		contentSplit := strings.Split(idpart[1], ":")
		coordinates := strings.Split(contentSplit[0], ",")
		xCoordinate, _ := strconv.Atoi(strings.TrimSpace(coordinates[0]))
		yCoordinate, _ := strconv.Atoi(strings.TrimSpace(coordinates[1]))
		size := strings.Split(contentSplit[1], "x")
		xSize, _ := strconv.Atoi(strings.TrimSpace(size[0]))
		ySize, _ := strconv.Atoi(strings.TrimSpace(size[1]))

		for x := 0; x < xSize; x++ {
			for y := 0; y < ySize; y++ {
				if clothMap[x+xCoordinate] == nil {
					clothMap[x+xCoordinate] = make(map[int]int)
				}
				clothMap[x+xCoordinate][y+yCoordinate]++
				if clothMap[x+xCoordinate][y+yCoordinate] == 2 {
					result++
				}
			}
		}
	}

	return result
}

func calculateResultB(part string, input []string) int {

	clothMap := make(map[int]map[int]int)
	result := 0

	for _, line := range input {
		idpart := strings.Split(line, "@")
		contentSplit := strings.Split(idpart[1], ":")
		coordinates := strings.Split(contentSplit[0], ",")
		xCoordinate, _ := strconv.Atoi(strings.TrimSpace(coordinates[0]))
		yCoordinate, _ := strconv.Atoi(strings.TrimSpace(coordinates[1]))
		size := strings.Split(contentSplit[1], "x")
		xSize, _ := strconv.Atoi(strings.TrimSpace(size[0]))
		ySize, _ := strconv.Atoi(strings.TrimSpace(size[1]))

		for x := 0; x < xSize; x++ {
			for y := 0; y < ySize; y++ {
				if clothMap[x+xCoordinate] == nil {
					clothMap[x+xCoordinate] = make(map[int]int)
				}
				clothMap[x+xCoordinate][y+yCoordinate]++
			}
		}
	}

outerLoop:
	for _, line := range input {

		idpart := strings.Split(line, "@")
		id, _ := strconv.Atoi(strings.TrimSpace(strings.Split(idpart[0], "#")[1]))
		contentSplit := strings.Split(idpart[1], ":")
		coordinates := strings.Split(contentSplit[0], ",")
		xCoordinate, _ := strconv.Atoi(strings.TrimSpace(coordinates[0]))
		yCoordinate, _ := strconv.Atoi(strings.TrimSpace(coordinates[1]))
		size := strings.Split(contentSplit[1], "x")
		xSize, _ := strconv.Atoi(strings.TrimSpace(size[0]))
		ySize, _ := strconv.Atoi(strings.TrimSpace(size[1]))

		for x := 0; x < xSize; x++ {
			for y := 0; y < ySize; y++ {
				if clothMap[x+xCoordinate][y+yCoordinate] != 1 {
					continue outerLoop
				}
			}
		}
		return id
	}

	return result

}

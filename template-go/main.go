package main

import (
	"bufio"
	"fmt"
	"os"
)

var file string

func getResult(part string) int {
	input := []string{}

	firstPart := part == "A"

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

	if firstPart {
		return calculateResultA(input)
	}

	return calculateResultB(input)
}

func calculateResultA(input []string) int {

	result := 0

	return result
}

func calculateResultB(input []string) int {

	result := 0

	return result

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

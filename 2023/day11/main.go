package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var file string

type Coord struct {
	X int
	Y int
}

func Distance(a, b Coord) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func calculateResultA(input []string) int {
	result := 0
	galaxies := ParseGalaxies(input, 2)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := Distance(galaxies[i], galaxies[j])
			result += distance
		}
	}
	return result
}

func calculateResultB(input []string) int {
	result := 0
	galaxies := ParseGalaxies(input, 1000000)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := Distance(galaxies[i], galaxies[j])
			result += distance
		}
	}
	return result

}

func ParseGalaxies(input []string, expansionFactor int) []Coord {
	width := len(input[0])
	columnFilled := make([]bool, width)
	emptyColumns := []int{}
	emptyRows := []int{}
	galaxies := []Coord{}
	// Find all empty rows and initial galaxy coordinates
	for y, line := range input {
		rowEmpty := true
		for x, el := range line {
			if el == '#' {
				galaxies = append(galaxies, Coord{x, y})
				rowEmpty = false
				columnFilled[x] = true
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, y)
		}
	}
	// use helper array to determine empty columns
	for i, hasGalaxy := range columnFilled {
		if !hasGalaxy {
			emptyColumns = append(emptyColumns, i)
		}
	}
	// use empty columns and rows to adjust galaxy coordinates
	for i := range galaxies {
		galaxies[i].X += smallerElements(emptyColumns, galaxies[i].X) * (expansionFactor - 1)
		galaxies[i].Y += smallerElements(emptyRows, galaxies[i].Y) * (expansionFactor - 1)
	}
	return galaxies
}

// return the number of elements in slice s which are smaller than x
func smallerElements(s []int, x int) int {
	result := 0
	for _, el := range s {
		if el < x {
			result++
		}
	}
	return result
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

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var file string
var distanceLimit int

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

type coordinates struct {
	X int
	Y int
}

func getResult(part string) int {

	input := []coordinates{}

	firstPart := part == "A"

	if file == "" {
		file = "day6.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		input = append(input, coordinates{x, y})
	}

	if firstPart {
		return calculateResultA(part, input)
	}

	return calculateResultB(part, input)
}

func calculateResultA(part string, input []coordinates) int {
	var xMax, yMax int
	for _, value := range input {
		if value.X > xMax {
			xMax = value.X
		}
		if value.Y > yMax {
			yMax = value.Y
		}
	}
	// Make and fill the grid with the closest points
	grid := make([][]int, xMax+1)
	for index := range grid {
		grid[index] = make([]int, yMax+1)
	}

	for x, column := range grid {
		for y := range column {
			grid[x][y] = lowestDistance(x, y, input)
		}
	}

	// Create an array to store the sizes of the area for each point
	sizeMap := make([]int, len(input))
	for x, column := range grid {
		for y := range column {
			if grid[x][y] > -1 {
				sizeMap[grid[x][y]]++
			}
		}
	}

	// Remove all points that have area on the edges of the grid (they are infinite)
	for x, column := range grid {
		for y := range column {
			if x == 0 || y == 0 {
				if grid[x][y] > -1 {
					sizeMap[grid[x][y]] = 0
				}
			}
		}
	}

	var maxSize int
	for _, value := range sizeMap {
		if value > maxSize {
			maxSize = value
		}
	}

	return maxSize
}

func calculateResultB(part string, input []coordinates) int {
	if distanceLimit == 0 {
		distanceLimit = 10000
	}

	var xMax, yMax int
	for _, value := range input {
		if value.X > xMax {
			xMax = value.X
		}
		if value.Y > yMax {
			yMax = value.Y
		}
	}
	// Make and fill the grid with the area that is bewlow the limit
	grid := make([][]bool, xMax+1)
	for index := range grid {
		grid[index] = make([]bool, yMax+1)
	}

	for x, column := range grid {
		for y := range column {
			grid[x][y] = belowThreshold(x, y, input)
		}
	}

	// Remove all points that have area on the edges of the grid (they are infinite)
	areaSize := 0
	for x, column := range grid {
		for y := range column {
			if grid[x][y] {
				areaSize++
			}
		}
	}

	return areaSize
}

func lowestDistance(x, y int, input []coordinates) int {
	min := math.MaxInt64
	result := 0
	for index, point := range input {
		distance := manhattanDistance(x, y, point)
		if distance < min {
			result = index
			min = distance
		} else if distance == min {
			result = -1
		}
	}
	return result
}

func manhattanDistance(x, y int, point coordinates) int {
	return abs(x-point.X) + abs(y-point.Y)
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func belowThreshold(x, y int, input []coordinates) bool {
	sum := 0
	for _, point := range input {
		sum += manhattanDistance(x, y, point)
	}
	if sum < distanceLimit {
		return true
	}
	return false
}

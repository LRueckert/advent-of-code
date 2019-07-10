package day10

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var file string

type position struct {
	X int
	Y int
}

type velocity struct {
	X int
	Y int
}

type point struct {
	Position position
	Velocity velocity
}

func (p *point) Move() {
	(*p).Position.X += (*p).Velocity.X
	(*p).Position.Y += (*p).Velocity.Y
}

type points []point

func (p *points) Move() {
	for index := range *p {
		(*p)[index].Move()
	}
}

// func (p *points) NextSum() int {
// 	sum := 0
// 	for _, point := range *p {
// 		sum += abs(point.Position.X + point.Velocity.X)
// 		sum += abs(point.Position.Y + point.Velocity.Y)
// 	}
// 	return sum
// }

func (p *points) GetBounds() (xMax, xMin, yMax, yMin int) {
	xMin = math.MaxInt64
	yMin = math.MaxInt64
	for _, point := range *p {
		newX := point.Position.X
		newY := point.Position.Y
		if newX > xMax {
			xMax = newX
		}
		if newX < xMin {
			xMin = newX
		}
		if newY > yMax {
			yMax = newY
		}
		if newY < yMin {
			yMin = newY
		}
	}
	return xMax + 1, xMin, yMax + 1, yMin
}

func (p *points) NextSum() (sum int) {
	xMin := math.MaxInt64
	yMin := math.MaxInt64
	var xMax, yMax int
	for _, point := range *p {
		newX := point.Position.X + point.Velocity.X
		newY := point.Position.Y + point.Velocity.Y
		if newX > xMax {
			xMax = newX
		}
		if newX < xMin {
			xMin = newX
		}
		if newY > yMax {
			yMax = newY
		}
		if newY < yMin {
			yMin = newY
		}
	}
	xDiff := xMax - xMin
	yDiff := yMax - yMin
	sum = xDiff * yDiff
	return
}

func (p *points) AllWithin(x int, y int) bool {
	for _, point := range *p {
		if abs(point.Position.X) > x/2 || abs(point.Position.Y) > y/2 {
			return false
		}
	}
	return true
}

// GetResult returns the result for Advent of Code Day x
func GetResult(part string) int {

	input := []point{}

	firstPart := part == "A"

	if file == "" {
		file = "day10/day10.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parsePoint(line))
	}

	if firstPart {
		return calculateResult(part, input)
	}

	return calculateResult(part, input)
}

func calculateResult(part string, input []point) int {

	var nextSum, xMax, xMin, yMax, yMin int
	points := points(input)
	currentSum := math.MaxInt64
	nextSum = points.NextSum()
	steps := 0
	for currentSum > nextSum {
		currentSum = nextSum
		points.Move()
		nextSum = points.NextSum()
		steps++
	}

	xMax, xMin, yMax, yMin = points.GetBounds()

	grid, xSize, ySize := initializeGrid(xMax, xMin, yMax, yMin)
	for _, point := range points {
		yAdjusted := point.Position.Y - yMin
		xAdjusted := point.Position.X - xMin
		if yAdjusted >= 0 && yAdjusted < ySize && xAdjusted >= 0 && xAdjusted < xSize {
			grid[yAdjusted][xAdjusted] = "#"
		}
	}
	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println("Steps: ", steps)

	return 0
}

func parsePoint(line string) point {
	re := regexp.MustCompile(`position=<(?P<position>.*)> velocity=<(?P<velocity>.*)>`)
	parts := re.FindStringSubmatch(line)
	pos := parts[1]
	positionParts := strings.Split(pos, ", ")
	xPos, _ := strconv.Atoi(strings.TrimSpace(positionParts[0]))
	yPos, _ := strconv.Atoi(strings.TrimSpace(positionParts[1]))

	vel := parts[2]
	velocityParts := strings.Split(vel, ", ")
	xVel, _ := strconv.Atoi(strings.TrimSpace(velocityParts[0]))
	yVel, _ := strconv.Atoi(strings.TrimSpace(velocityParts[1]))

	return point{position{xPos, yPos}, velocity{xVel, yVel}}
}

func initializeGrid(xMax, xMin, yMax, yMin int) (grid [][]string, xSize, ySize int) {
	xSize = xMax - xMin
	ySize = yMax - yMin
	grid = make([][]string, ySize)
	for y := range grid {
		grid[y] = make([]string, xSize)
	}
	for y, row := range grid {
		for x := range row {
			grid[y][x] = "."
		}
	}
	return
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

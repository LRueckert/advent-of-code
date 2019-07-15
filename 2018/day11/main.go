package main

import (
	"fmt"
	"os"
	"time"
)

var serialNumber = 6392
var cellMap [301][301]int

type grid struct {
	size  int
	x     int
	y     int
	value int
}

func (g *grid) moveY() {

	value := g.value

	for i := 0; i < g.size; i++ {
		value -= cellMap[g.x+i][g.y]
		value += cellMap[g.x+i][g.y+g.size]
	}
	(*g).value = value
	(*g).y = g.y + 1
}

func (g *grid) moveX() {

	(*g).x++
	(*g).y = 1

	(*g).value = calculateGridPower(cellMap, g.x, g.y, g.size)
}

func getResult(part string) string {

	start := time.Now()

	var result string

	for x, row := range cellMap {
		for y := range row {
			cellMap[x][y] = calculateCellPower(x, y)
		}
	}

	firstPart := part == "A"

	if firstPart {
		result = calculateResultA(part)
	} else {
		result = calculateResultB(part)
	}

	fmt.Println(time.Since(start))
	return result
}

func calculateResultA(part string) string {

	var cellGrid grid
	s := 3
	cellGrid = grid{s, 0, 0, 0}
	var max, xMax, yMax int
	for x := 0; x < 300-s; x++ {
		cellGrid.moveX()
		if cellGrid.value > max {
			max = cellGrid.value
			xMax = cellGrid.x
			yMax = cellGrid.y
		}
		for y := 0; y < 300-s; y++ {
			cellGrid.moveY()
			if cellGrid.value > max {
				max = cellGrid.value
				xMax = cellGrid.x
				yMax = cellGrid.y
			}
		}
	}

	return fmt.Sprintf("%v,%v", xMax, yMax)
}

func calculateResultB(part string) string {

	var cellGrid grid

	var max, xMax, yMax, sizeMax int
	for s := 1; s < 301; s++ {
		cellGrid = grid{s, 0, 0, 0}
		for x := 0; x < 300-s; x++ {
			cellGrid.moveX()
			if cellGrid.value > max {
				max = cellGrid.value
				xMax = cellGrid.x
				yMax = cellGrid.y
				sizeMax = cellGrid.size
			}
			for y := 0; y < 300-s; y++ {
				cellGrid.moveY()
				if cellGrid.value > max {
					max = cellGrid.value
					xMax = cellGrid.x
					yMax = cellGrid.y
					sizeMax = cellGrid.size
				}
			}
		}
	}

	return fmt.Sprintf("%v,%v,%v", xMax, yMax, sizeMax)

}

func calculateCellPower(x, y int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serialNumber
	powerLevel *= rackID
	powerLevel = powerLevel % 1000 / 100
	powerLevel -= 5

	return powerLevel
}

func calculateGridPower(cellMap [301][301]int, x, y, size int) int {

	var result int
	for dx := 0; dx < size; dx++ {
		for dy := 0; dy < size; dy++ {
			result += cellMap[x+dx][y+dy]
		}
	}

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

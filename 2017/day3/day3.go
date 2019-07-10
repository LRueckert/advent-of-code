package day3

import (
	"fmt"
)

const right = 0
const up = 1
const left = 2
const down = 3

type point struct {
	X int
	Y int
}

type grid map[point]int

// GetResult returns the result for Advent of Code Day 3
func GetResult(part string) int {

	target := 368078
	result := 0

	grid := grid{point{0, 0}: 1}

	var x int
	var y int
	directionBorders := []int{0, 0, 0, 0}
	direction := 0

	for i := 2; result <= target; i++ {
		switch direction {
		case right:
			x++
			if x > directionBorders[right] {
				directionBorders[right] = x
				direction = up
			}
		case up:
			y++
			if y > directionBorders[up] {
				directionBorders[up] = y
				direction = left
			}
		case left:
			x--
			if x < directionBorders[left] {
				directionBorders[left] = x
				direction = down
			}
		case down:
			y--
			if y < directionBorders[down] {
				directionBorders[down] = y
				direction = right
			}
		}
		result = addNumberToGrid(grid, point{x, y})
	}
	fmt.Println(grid)

	return result
}

func addNumberToGrid(g grid, p point) int {
	x := p.X
	y := p.Y
	result := 0

	result += g[point{x + 1, y}]
	result += g[point{x - 1, y}]
	result += g[point{x, y + 1}]
	result += g[point{x, y - 1}]
	result += g[point{x + 1, y + 1}]
	result += g[point{x + 1, y - 1}]
	result += g[point{x - 1, y + 1}]
	result += g[point{x - 1, y - 1}]

	g[point{x, y}] = result

	return result
}

// GetResult returns the result for Advent of Code Day 3
// func GetResult() int {

// 		target := 368078

// 		var x float64
// 		var y float64
// 		directionBorders := []float64{0, 0, 0, 0}
// 		direction := 0

// 		for i := 2; i <= target; i++ {
// 			switch direction {
// 			case right:
// 				x++
// 				if x > directionBorders[right] {
// 					directionBorders[right] = x
// 					direction = up
// 				}
// 			case up:
// 				y++
// 				if y > directionBorders[up] {
// 					directionBorders[up] = y
// 					direction = left
// 				}
// 			case left:
// 				x--
// 				if x < directionBorders[left] {
// 					directionBorders[left] = x
// 					direction = down
// 				}
// 			case down:
// 				y--
// 				if y < directionBorders[down] {
// 					directionBorders[down] = y
// 					direction = right
// 				}
// 			}

// 		}
// 		fmt.Printf("(%v/%v)", x, y)

// 		return int(math.Abs(x) + math.Abs(y))
// 	}

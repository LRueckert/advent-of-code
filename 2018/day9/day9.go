package day9

import (
	"fmt"
	"sort"
)

var numPlayers = 438
var lastMarble = 71626

type marble struct {
	Value int
	Left  *marble
	Right *marble
}

func (m *marble) placeNext(value int) *marble {
	first := m.Right
	second := first.Right
	new := &marble{value, first, second}
	first.Right = new
	second.Left = new
	return new
}

func (m *marble) remove(value int) (int, *marble) {
	walker := m
	for i := 0; i < 7; i++ {
		walker = walker.Left
	}
	right := walker.Right
	left := walker.Left
	walker.Left.Right = right
	walker.Right.Left = left
	return walker.Value + value, right
}

func (m *marble) String() string {
	result := fmt.Sprintf("%v", m.Value)
	for walker := m.Right; walker != m; walker = walker.Right {
		result += fmt.Sprintf(" - %v", walker.Value)
	}
	return result
}

// GetResult returns the result for Advent of Code Day x
func GetResult(part string) int {

	firstPart := part == "A"

	if firstPart {
		return calculateResult()
	}
	lastMarble = 100 * lastMarble
	return calculateResult()
}

func calculateResult() int {

	elves := make([]int, numPlayers)

	currentElf := 0
	firstMarble := &marble{0, nil, nil}
	firstMarble.Left = firstMarble
	firstMarble.Right = firstMarble

	currentMarble := firstMarble

	for nextMarble := 1; nextMarble <= lastMarble; nextMarble++ {
		if nextMarble%23 == 0 {
			var points int
			points, currentMarble = currentMarble.remove(nextMarble)
			elves[currentElf] += points
		} else {
			currentMarble = currentMarble.placeNext(nextMarble)
		}
		currentElf++
		if currentElf >= numPlayers {
			currentElf = 0
		}
	}

	sort.IntSlice(elves).Sort()

	return elves[numPlayers-1]
}

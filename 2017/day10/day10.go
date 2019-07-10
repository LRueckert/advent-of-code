package day10

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var file string
var listLength int
var list []int

// GetResult returns the result for Advent of Code Day 9
func GetResult(part string) int {
	if file == "" {
		file = "day10/day10.input"
		listLength = 256
	}
	bytes, _ := ioutil.ReadFile(file)
	input := string(bytes)
	lengthStrings := strings.Split(input, ",")
	lengths := make([]int, len(lengthStrings))
	for i, v := range lengthStrings {
		lengths[i], _ = strconv.Atoi(v)
	}
	position := 0
	skip := 0

	list = make([]int, listLength)
	for i := range list {
		list[i] = i
	}

	for _, length := range lengths {
		reverse(position, length)
		position = (position + length + skip) % len(list)
		skip++
	}

	return list[0] * list[1]
}

func reverse(position int, length int) {
	for i := 0; i < length/2; i++ {
		temp := list[(position+i)%len(list)]
		list[(position+i)%len(list)] = list[(position+length-i-1)%len(list)]
		list[(position+length-i-1)%len(list)] = temp
	}
}

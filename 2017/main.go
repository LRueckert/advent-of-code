package main

import (
	"advent-of-code/2017/day1"
	"advent-of-code/2017/day10"
	"advent-of-code/2017/day2"
	"advent-of-code/2017/day3"
	"advent-of-code/2017/day4"
	"advent-of-code/2017/day5"
	"advent-of-code/2017/day6"
	"advent-of-code/2017/day7"
	"advent-of-code/2017/day8"
	"advent-of-code/2017/day9"
	"fmt"
	"os"
	"strconv"
)

func main() {

	argsWithProg := os.Args
	day, _ := strconv.Atoi(argsWithProg[1])

	var part string
	if len(argsWithProg) < 3 {
		part = "B"
	} else {
		part = argsWithProg[2]
	}
	result := ""
	switch day {
	case 1:
		result = strconv.Itoa(day1.GetResult(part))
	case 2:
		result = strconv.Itoa(day2.GetResult(part))
	case 3:
		result = strconv.Itoa(day3.GetResult(part))
	case 4:
		result = strconv.Itoa(day4.GetResult(part))
	case 5:
		result = strconv.Itoa(day5.GetResult(part))
	case 6:
		result = strconv.Itoa(day6.GetResult(part))
	case 7:
		result = strconv.Itoa(day7.GetResult(part))
	case 8:
		result = strconv.Itoa(day8.GetResult(part))
	case 9:
		result = strconv.Itoa(day9.GetResult(part))
	case 10:
		result = strconv.Itoa(day10.GetResult(part))
	}
	fmt.Println(result)

}

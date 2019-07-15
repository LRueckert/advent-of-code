package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

type chronologically []string

func (s chronologically) Len() int {
	return len(s)
}
func (s chronologically) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s chronologically) Less(i, j int) bool {
	iParts := strings.Split(s[i], "]")
	iDate, _ := time.Parse("2006-01-02 15:04", iParts[0][1:])
	jParts := strings.Split(s[j], "]")
	jDate, _ := time.Parse("2006-01-02 15:04", jParts[0][1:])
	return iDate.Before(jDate)
}

func getResult(part string) int {

	input := []string{}

	firstPart := part == "A"

	if file == "" {
		file = "day4.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return calculateResult(part, input, firstPart)
}

func calculateResult(part string, input []string, firstPart bool) int {

	sleepMap := make(map[int]map[string][]bool)
	sleepSumMap := make(map[int]int)
	var id int

	// Sort Input chronologically
	sort.Sort(chronologically(input))

	// Create and FIll Sleepmap
	for _, line := range input {
		parts := strings.Split(line, "]")
		dateParts := strings.Split(parts[0][1:], " ")
		dateTime, _ := time.Parse("2006-01-02", dateParts[0])
		hour, _ := strconv.Atoi(strings.Split(dateParts[1], ":")[0])
		if hour != 0 {
			dateTime = dateTime.AddDate(0, 0, 1)
		}
		date := strings.Split(dateTime.String(), " ")[0]
		minute, _ := strconv.Atoi(strings.Split(dateParts[1], ":")[1])

		contentParts := strings.Split(parts[1], " ")

		switch contentParts[1] {
		case "Guard":
			id, _ = strconv.Atoi(contentParts[2][1:])
			if sleepMap[id] == nil {
				sleepMap[id] = make(map[string][]bool)
			}
			//If shift starts before midnight, add one day
			if sleepMap[id][date] == nil {
				sleepMap[id][date] = make([]bool, 61, 61)
			}
		case "falls":
			for index := minute; index <= 60; index++ {
				sleepMap[id][date][index] = true
				sleepSumMap[id]++
			}
		case "wakes":
			for index := minute; index <= 60; index++ {
				sleepMap[id][date][index] = false
				sleepSumMap[id]--
			}
		}
	}

	if firstPart {
		// Use SleepSumMap to find guard that sleeps the most
		var sleepiest int
		mostSleep := 0
		for guard, sum := range sleepSumMap {
			if sum > mostSleep {
				sleepiest = guard
				mostSleep = sum
			}
		}

		// Aggregate the sleeping minutes over the days
		var sleptMinutes [61]int
		for _, minutes := range sleepMap[sleepiest] {
			for minute, sleeping := range minutes {
				if sleeping {
					sleptMinutes[minute]++
				}
			}
		}

		// Find the minute he sleeps the most
		var sleepiestMinute int
		sleepCount := 0
		for minute, count := range sleptMinutes {
			if count > sleepCount {
				sleepiestMinute = minute
				sleepCount = count
			}
		}

		return sleepiest * sleepiestMinute

	}

	// Aggregate the sleeping minutes over the days
	resultMinute := 0
	resultGuard := 0
	highestCount := 0
	guardMap := make(map[int][]int)
	for guard := range sleepMap {
		guardMap[guard] = make([]int, 61, 61)
		for _, minutes := range sleepMap[guard] {
			for minute, sleeping := range minutes {
				if sleeping {
					guardMap[guard][minute]++
					if guardMap[guard][minute] > highestCount {
						resultGuard = guard
						resultMinute = minute
						highestCount = guardMap[guard][minute]
					}
				}
			}
		}
	}

	return resultMinute * resultGuard

}

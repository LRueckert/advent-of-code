package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var file string

var windowTime, configTime, propagateTime time.Duration

type flowerPot struct {
	number int
	plant  bool
	left   *flowerPot
	right  *flowerPot
}

func (p flowerPot) String() string {
	if p.plant {
		return "#"
	}
	return "."
}

func (p flowerPot) AdaptPreviousConfig(config [5]bool) (output [5]bool) {
	copy(output[:], config[1:])
	if r := p.right; r != nil {
		if rr := r.right; rr != nil {
			output[4] = rr.plant
		} else {
			output[4] = false
		}
	} else {
		output[4] = false
	}
	return
}

func (p *flowerPot) Propagate(produceMap map[[5]bool]bool, previousConfig [5]bool) [5]bool {
	config := p.AdaptPreviousConfig(previousConfig)
	_, exists := produceMap[config]
	p.plant = exists
	return config
}

type flowerPots struct {
	first *flowerPot
	last  *flowerPot
}

func (p *flowerPots) prepend(value bool) {
	newPot := flowerPot{p.first.number - 1, value, nil, p.first}
	p.first.left = &newPot
	p.first = &newPot
}

func (p *flowerPots) append(value bool) {
	newPot := flowerPot{p.last.number + 1, value, p.last, nil}
	p.last.right = &newPot
	p.last = &newPot
}

func (p flowerPots) String() string {
	output := ""
	for i := 0; i > p.first.number; i-- {
		output += " "
	}
	output += "0\n"

	for iter := p.first; iter != p.last; iter = iter.right {
		output += iter.String()
	}
	output += p.last.String()
	return output
}

func (p flowerPots) Value() (result int) {
	for iter := p.first; iter != p.last; iter = iter.right {
		if iter.plant {
			result += iter.number
		}
	}
	if p.last.plant {
		result += p.last.number
	}
	return result
}

func (p *flowerPots) Propagate(produceMap map[[5]bool]bool) {
	var config [5]bool
	prependConfig := [5]bool{false, false, false, p.first.plant, p.first.right.plant}
	config = prependConfig
	for iter := p.first; iter != p.last; iter = iter.right {
		config = iter.Propagate(produceMap, config)
	}
	config = p.last.Propagate(produceMap, config)
	var appendConfig [5]bool
	copy(appendConfig[:], config[1:])
	appendConfig[4] = false
	if produceMap[appendConfig] {
		p.append(true)
	}
	if produceMap[prependConfig] {
		p.prepend(true)
	}
}

// GetResult returns the result for Advent of Code Day x
func GetResult(part string) int {

	start := time.Now()

	var pots flowerPots
	produceMap := make(map[[5]bool]bool)

	firstPart := part == "A"

	if file == "" {
		file = "day12/day12.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ": ")
	initialState := strings.Split(parts[1], "")

	var pot flowerPot

	for index, value := range initialState {
		if index == 0 {
			pot = flowerPot{0, value == "#", nil, nil}
			pots.first = &pot
			pots.last = &pot
		} else {
			pots.append(value == "#")
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		if len(parts) > 1 {
			if parts[1] == "#" {
				first := parts[0][0] == '#'
				second := parts[0][1] == '#'
				third := parts[0][2] == '#'
				fourth := parts[0][3] == '#'
				fifth := parts[0][4] == '#'
				key := [5]bool{first, second, third, fourth, fifth}
				produceMap[key] = true
			}
		}
	}

	var result int
	if firstPart {
		result = calculateResultA(part, produceMap, pots)
	} else {
		result = calculateResultB(part, produceMap, pots)
	}

	// fmt.Printf("Total Time: %v - windowTime: %v - propagateTime: %v\n", time.Since(start), windowTime, propagateTime)
	fmt.Printf("Adapt Config: %v\n", time.Since(start))
	return result
}

func calculateResultA(part string, produceMap map[[5]bool]bool, pots flowerPots) int {

	// fmt.Println(pots)
	for i := 0; i < 20; i++ {
		pots.Propagate(produceMap)
		// fmt.Println(pots)
	}

	return pots.Value()
}

func calculateResultB(part string, produceMap map[[5]bool]bool, pots flowerPots) int {

	for i := 0; i < 1e3; i++ {
		pots.Propagate(produceMap)
	}
	value := pots.Value()
	value += (5e10 - 1e3) * 51

	return value

}

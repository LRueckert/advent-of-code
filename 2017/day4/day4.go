package day4

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

// GetResult returns the result for Advent of Code Day 4 b
func GetResult(part string) int {

	f, _ := os.Open("day4/day4.input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		passwords := strings.Split(line, " ")

		for i, v := range passwords {
			chars := strings.Split(v, "")
			sort.Strings(chars)
			passwords[i] = strings.Join(chars, "")
		}

		sort.Strings(passwords)
		valid := true
		for i := 0; i < len(passwords)-1; i++ {
			if passwords[i] == passwords[i+1] {
				valid = false
			}
		}
		if valid {
			result++
		}

	}

	return result
}

// // GetResult returns the result for Advent of Code Day 4 a
// func GetResult() int {

// 	f, _ := os.Open("day4/day4.input")
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	scanner.Split(bufio.ScanLines)
// 	result := 0

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		passwords := strings.Split(line, " ")

// 		sort.Strings(passwords)
// 		valid := true
// 		for i := 0; i < len(passwords)-1; i++ {
// 			if passwords[i] == passwords[i+1] {
// 				valid = false
// 			}
// 		}
// 		if valid {
// 			result++
// 		}

// 	}

// 	return result
// }

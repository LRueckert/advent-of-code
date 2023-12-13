package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// var file string

// var hashRe = regexp.MustCompile(`#+`)

// type Row struct {
// 	Pattern string
// 	Groups  []int
// }

// func (r *Row) Expand() {
// 	slice := []string{r.Pattern, r.Pattern, r.Pattern, r.Pattern, r.Pattern}
// 	r.Pattern = strings.Join(slice, "?")
// 	var newGroups []int
// 	for i := 0; i < 5; i++ {
// 		newGroups = append(newGroups, r.Groups...)
// 	}
// 	r.Groups = newGroups
// }

// func indentedLog(depth int, pattern string, a ...any) {
// 	if depth < 10 {
// 		fmt.Printf(strings.Repeat("  ", depth)+pattern, a...)
// 	}
// }

// func filterEmpty(s []string) (filtered []string) {
// 	for _, el := range s {
// 		if el != "" {
// 			filtered = append(filtered, el)
// 		}
// 	}
// 	return
// }

// func calculateResultA(input []string) int {
// 	result := 0
// 	rows := ParseRows(input)
// 	// figure out groups without separators
// 	for _, row := range rows {
// 		dotSplit := filterEmpty(strings.Split(row.Pattern, "."))
// 		arrangements := solveForDotsplit(dotSplit, row.Groups, 1)
// 		fmt.Printf("== %v ==", row)
// 		fmt.Printf("==> %d\n", len(arrangements))
// 		result += len(arrangements)
// 	}
// 	return result
// }

// func solveForDotsplit(dotSplit []string, groups []int, depth int) map[string]bool {
// 	indentedLog(depth, "Solving for dotsplit %v with groups %v\n", dotSplit, groups)
// 	// try to match dotSplit with groups
// 	// 2 cases:
// 	// - dotSplit has the number of groups that we need -> figure out permutations for groups
// 	permutations := make(map[string]bool)
// 	if len(dotSplit) == len(groups) {
// 		permutations = solveForMatchingGroups(dotSplit, groups, depth)
// 	} else if len(dotSplit) < len(groups) {
// 		// - dotsplit has less groups than we need -> figure out places where we can place a dot to increase the number of groups
// 		// then for each possible dot placement recurse to find solution
// 		possibleSplits := possibleSplits(dotSplit)
// 		for _, spot := range possibleSplits {
// 			newSplit := splitToGenerateNewDotsplit(dotSplit, spot)
// 			indentedLog(depth, "New separated split %v\n", newSplit)
// 			subPerms := solveForDotsplit(newSplit, groups, depth+1)
// 			for perm := range subPerms {
// 				permutations[perm] = true
// 			}
// 		}
// 	} else {
// 		// dotsplit has more groups than we need -> pairwise merge two groups to decrease number of groups
// 		for i := 0; i < len(dotSplit)-1; i++ {
// 			newSplit := mergeToGenerateNewDotsplit(dotSplit, i)
// 			indentedLog(depth, "New merged split %v\n", newSplit)
// 			subPerms := solveForDotsplit(newSplit, groups, depth+1)
// 			for perm := range subPerms {
// 				permutations[perm] = true
// 			}
// 		}
// 	}
// 	indentedLog(depth, "==> %d\n", len(permutations))
// 	for perm := range permutations {
// 		indentedLog(depth, "==> %v\n", perm)
// 	}
// 	return permutations
// }

// func splitToGenerateNewDotsplit(dotsplit []string, split []int) []string {
// 	newSplit := make([]string, 0)
// 	for i, el := range dotsplit {
// 		if i != split[0] {
// 			newSplit = append(newSplit, el)
// 		} else {
// 			newSplit = append(newSplit, el[:split[1]], el[split[1]+1:])
// 		}
// 	}
// 	return newSplit
// }

// func mergeToGenerateNewDotsplit(dotsplit []string, pos int) []string {
// 	newSplit := make([]string, 0)
// 	for i := 0; i < len(dotsplit); i++ {
// 		if i != pos {
// 			newSplit = append(newSplit, dotsplit[i])
// 		} else {
// 			newSplit = append(newSplit, dotsplit[i]+dotsplit[i+1])
// 			i++
// 		}
// 	}
// 	return newSplit
// }

// // returns the number of spots in dotsplits where you can put a dot to increase the number of groups
// func possibleSplits(dotSplit []string) (possible [][]int) {
// 	for i, group := range dotSplit {
// 		for j := 1; j < len(group)-1; j++ {
// 			if string(group[j]) == "?" {
// 				possible = append(possible, []int{i, j})
// 			}
// 		}
// 	}
// 	return
// }

// func solveForMatchingGroups(dotSplit []string, groups []int, depth int) map[string]bool {
// 	indentedLog(depth, "Solving for %v - %v\n", dotSplit, groups)
// 	expectedLength := len(dotSplit) - 1
// 	for _, split := range dotSplit {
// 		expectedLength += len(split)
// 	}
// 	permutations := make(map[string]bool)
// 	parts := make([][]string, len(groups))
// 	if len(dotSplit) == len(groups) {
// 		for i := 0; i < len(dotSplit); i++ {
// 			parts[i] = calculateGroupPermutations(dotSplit[i], groups[i], depth)
// 		}
// 	}
// 	groupPermutations := GenerateStringPermutations(parts)
// 	for _, perm := range groupPermutations {
// 		if len(perm) != expectedLength {
// 			indentedLog(depth, "!!!! Permutation %s is not expected length %d\n", perm, expectedLength)
// 			indentedLog(depth, "parts: %v\n", parts)
// 		}
// 		permutations[perm] = true
// 	}
// 	indentedLog(depth, "Solving for %v - %v => %d\n", dotSplit, groups, len(permutations))
// 	return permutations
// }

// func GenerateStringPermutations(in [][]string) []string {
// 	out := []string{}
// 	if len(in) == 1 {
// 		return in[0]
// 	}
// 	restPerm := GenerateStringPermutations(in[1:])
// 	for _, el := range in[0] {
// 		for _, rest := range restPerm {
// 			out = append(out, el+"."+rest)
// 		}
// 	}
// 	return out
// }

// func calculateGroupPermutations(pattern string, length int, depth int) (permutations []string) {
// 	hashBlock := strings.Repeat("#", length)
// 	indentedLog(depth, "Permutations for length of %d over %s\n", length, pattern)
// 	if len(pattern) < length {
// 		return
// 	}
// 	if len(pattern) == length {
// 		permutations = append(permutations, hashBlock)
// 		return
// 	}
// 	// cannot solve permutation if continuous block is always going to be too long
// 	firstHash := strings.Index(pattern, "#")
// 	lastHash := strings.LastIndex(pattern, "#")
// 	if length < (lastHash - (firstHash - 1)) {
// 		return
// 	}
// 	// has exactly one solution if one of th ends is a #
// 	if firstHash == 0 {
// 		permutations = append(permutations, fmt.Sprintf("%s%s", hashBlock, strings.Repeat(".", len(pattern)-length)))
// 		return
// 	}
// 	if firstHash == len(pattern)-1 {
// 		permutations = append(permutations, fmt.Sprintf("%s%s", strings.Repeat(".", len(pattern)-length), hashBlock))
// 		return
// 	}
// 	hashSplit := strings.Split(pattern, "#")
// 	indentedLog(depth, " hashtag split %v (%d)\n", hashSplit, len(hashSplit))
// 	// if after splitting on hashtags we have exactly one group, then its all questionmarks and we make the permutations
// 	if len(hashSplit) == 1 {
// 		for i := 0; i <= len(pattern)-length; i++ {
// 			value := fmt.Sprintf("%s%s%s", strings.Repeat(".", i), hashBlock, strings.Repeat(".", len(pattern)-length-i))
// 			permutations = append(permutations, value)
// 		}
// 		return
// 	}
// 	// we had some # in our pattern and need to account for those (we had one between each element)
// 	remaining := length - (len(hashSplit) - 1)
// 	// since we have to create one continuous block we can remove all groups that are not first or last and further reduce the amount to distribute
// 	indentedLog(depth, " hashtag split %v (%d) with %d remaining\n", hashSplit, len(hashSplit), remaining)
// 	for i := 1; i < len(hashSplit)-1; i++ {
// 		remaining -= len(hashSplit[i])
// 	}
// 	if remaining == 0 {
// 		permutations = append(permutations, strings.Replace(pattern, "?", ".", -1))
// 		return
// 	}
// 	if remaining < 0 {
// 		return
// 	}
// 	left := hashSplit[0]
// 	right := hashSplit[len(hashSplit)-1]
// 	leftPerms := calculateGroupPermutations(left, int(math.Min(float64(len(left)), float64(remaining))), depth+1)
// 	leftPerms = append(leftPerms, strings.Repeat(".", len(left)))
// 	rightPerms := calculateGroupPermutations(right, int(math.Min(float64(len(right)), float64(remaining))), depth+1)
// 	rightPerms = append(rightPerms, strings.Repeat(".", len(right)))
// 	indentedLog(depth, " leftPerms: %v\n", leftPerms)
// 	indentedLog(depth, " rightPerms: %v\n", rightPerms)
// 	for _, leftPerm := range leftPerms {
// 		for _, rightPerm := range rightPerms {
// 			if hashCount(leftPerm+rightPerm) == remaining {
// 				value := fmt.Sprintf("%s%s%s", leftPerm, strings.Repeat("#", length-remaining), rightPerm)
// 				match := hashRe.FindString(value)
// 				if len(match) == length {
// 					permutations = append(permutations, value)
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// func hashCount(s string) int {
// 	count := 0
// 	for _, el := range s {
// 		if el == '#' {
// 			count++
// 		}
// 	}
// 	return count
// }

// func calculateResultB(input []string) int {
// 	result := 0
// 	rows := ParseRows(input)
// 	// figure out groups without separators
// 	for _, row := range rows {
// 		row.Expand()
// 		dotSplit := filterEmpty(strings.Split(row.Pattern, "."))
// 		arrangements := solveForDotsplit(dotSplit, row.Groups, 1)
// 		fmt.Printf("== %v ==", row)
// 		fmt.Printf("==> %d\n", len(arrangements))
// 		result += len(arrangements)
// 	}
// 	return result
// }

// func ParseRows(input []string) []Row {
// 	rows := make([]Row, len(input))
// 	for i, line := range input {
// 		split := strings.Split(line, " ")
// 		sGroups := strings.Split(split[1], ",")
// 		groups := make([]int, len(sGroups))
// 		for j, el := range sGroups {
// 			groups[j], _ = strconv.Atoi(el)
// 		}
// 		rows[i] = Row{Pattern: split[0], Groups: groups}
// 	}
// 	return rows
// }

// func getResult(part string) int {
// 	input := getInput()
// 	firstPart := part == "A"

// 	if firstPart {
// 		return calculateResultA(input)
// 	}

// 	return calculateResultB(input)
// }

// func getInput() []string {
// 	input := []string{}

// 	if file == "" {
// 		file = "input.txt"
// 	}
// 	f, _ := os.Open(file)
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	scanner.Split(bufio.ScanLines)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		input = append(input, line)
// 	}

// 	return input
// }

// func main() {
// 	argsWithProg := os.Args

// 	var part string
// 	if len(argsWithProg) < 2 {
// 		part = "A"
// 	} else {
// 		part = argsWithProg[1]
// 	}

// 	fmt.Println(getResult(part))
// }

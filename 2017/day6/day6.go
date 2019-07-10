package day6

// GetResult returns the result for Advent of Code Day 6 a
func GetResult(part string) int {

	// input := []int{0, 2, 7, 0}
	input := []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}

	var memory [][]int
	memory = append(memory, createCopy(input))

	result := 0

	for true {
		result++
		index, highest := getHighest(input)
		input[index] = 0
		for i := 1; i <= highest; i++ {
			input[(index+i)%len(input)]++
		}
		for i, v := range memory {
			if areEqual(v, input) {
				return result - i
			}
		}
		// fmt.Println(memory)
		memory = append(memory, createCopy(input))
	}

	return 0
}

func createCopy(array []int) []int {
	tmp := make([]int, len(array))
	copy(tmp, array)
	return tmp
}

func areEqual(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getHighest(array []int) (int, int) {
	highest := 0
	index := 0
	for i, v := range array {
		if v > highest {
			highest = v
			index = i
		}
	}
	return index, highest
}

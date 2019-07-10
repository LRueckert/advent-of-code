package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	Root node
}

func (t tree) Balance() int {
	return t.Root.Balance(0)
}

type node struct {
	Name     string
	Parent   *node
	Weight   int
	Children []*node
}

func (n node) GetWeight() int {
	weight := n.Weight
	for _, v := range n.Children {
		weight += v.GetWeight()
	}
	return weight
}

func (n node) Balance(difference int) int {
	weights := make([]int, len(n.Children))
	for i, v := range n.Children {
		weights[i] = v.GetWeight()
	}
	for i := 0; i < len(weights); i++ {
		if weights[i] != weights[(i+1)%len(weights)] && weights[i] != weights[(i+2)%len(weights)] {
			return n.Children[i].Balance(weights[i] - weights[(i+1)%len(weights)])
		}
	}

	return n.Weight - difference
}

// GetResult returns the result for Advent of Code Day 7 a
func GetResult(part string) int {
	f, _ := os.Open("day7/day7.input")
	defer f.Close()
	var result tree

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	nodes := make(map[string]*node)
	childMap := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Split(line, " ")[0]
		parts := strings.Split(line, "->")
		for i, v := range parts {
			// fmt.Println(i, v)

			switch i {
			case 0:
				weightString := strings.Split(strings.TrimSpace(v), " ")[1]
				weight, _ := strconv.Atoi(weightString[1 : len(weightString)-1])
				nodes[name] = &node{name, nil, weight, []*node{}}
			case 1:
				childMap[name] = strings.Split(strings.TrimSpace(v), ", ")
			}
		}
	}
	for parent, v := range childMap {
		for _, child := range v {
			(*nodes[parent]).Children = append((*nodes[parent]).Children, nodes[child])
			(*nodes[child]).Parent = nodes[parent]
		}
	}
	for _, v := range nodes {
		if v.Parent == nil {
			result.Root = *v
			break
		}
	}

	return result.Balance()
}

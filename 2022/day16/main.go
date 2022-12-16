package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var file string

var (
	fullRE = regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)
)

type Valve struct {
	Name     string
	FlowRate int
	Opened   int
	Tunnels  map[string]int
}

func (v Valve) Value(minutes int) int {
	return v.FlowRate * minutes
}

func (v Valve) String() string {
	text := fmt.Sprintf("[%s] - %d :", v.Name, v.FlowRate)
	valves := []string{}
	for name := range v.Tunnels {
		valves = append(valves, name)
	}
	slices.Sort(valves)
	for _, to := range valves {
		text += fmt.Sprintf("%d -> %s | ", v.Tunnels[to], to)
	}
	return text
}

type Volcano struct {
	Position string
	Valves   map[string]*Valve
}

func (v Volcano) String() string {
	text := ""
	valves := []string{}
	for name := range v.Valves {
		valves = append(valves, name)
	}
	slices.Sort(valves)
	for _, name := range valves {
		text += v.Valves[name].String() + "\n"
	}
	return text
}

// Simplify the graph by removing valves with flowRate 0 as they are pointless
// except for the start node
func (v *Volcano) Simplify() {
	for name, valve := range v.Valves {
		if name != v.Position && valve.FlowRate == 0 {
			// To remove a tunnel instead add new tunnels connection to all other tunnels
			// GG -> FF & FF -> EE => GG -> EE (with sum of distances)
			for from, fromDistance := range valve.Tunnels {
				for to, toDistance := range valve.Tunnels {
					if from != to {
						v.Valves[from].Tunnels[to] = fromDistance + toDistance
						v.Valves[to].Tunnels[from] = fromDistance + toDistance
						delete(v.Valves[from].Tunnels, name)
						delete(v.Valves[to].Tunnels, name)
					}
				}
			}
			delete(v.Valves, name)
		}
	}
}

// CompleteDistanceMap uses dijsktra to find the shortest connection from each valve to each other valve
func (v *Volcano) CompleteDistanceMap() {
	for name := range v.Valves {
		v.Dijkstra(name)
	}
}

// Dijkstra uses dikstra algorithm to find the shortest distance from this node to all other nodes
func (v *Volcano) Dijkstra(name string) {
	visited := []string{}
	unvisited := []string{}
	for valve := range v.Valves {
		if valve != name {
			unvisited = append(unvisited, valve)
			if _, ok := v.Valves[name].Tunnels[valve]; !ok {
				v.Valves[name].Tunnels[valve] = math.MaxInt
			}
		}
	}
	current := name
	for len(unvisited) > 0 {
		// Get next value to check
		slices.SortFunc(unvisited, func(a, b string) bool {
			return v.Valves[name].Tunnels[a] <= v.Valves[name].Tunnels[b]
		})
		current, unvisited = unvisited[0], unvisited[1:]
		visited = append(visited, current)
		// Process current node
		for to, distance := range v.Valves[current].Tunnels {
			if name != to {
				indirectDistance := v.Valves[name].Tunnels[current] + distance
				if indirectDistance < v.Valves[name].Tunnels[to] {
					v.Valves[name].Tunnels[to] = indirectDistance
				}
			}
		}
	}
}

// OpenBestValve - Best Valve is the one that when opened results in the most flow
func (v *Volcano) OpenBestValve(timer int) int {
	max := 0
	valve := ""
	remaining := 0
	for to, distance := range v.Valves[v.Position].Tunnels {
		// Find Valve that is not opened and can still be reached
		if v.Valves[to].Opened <= 0 && distance+1 <= timer && v.Valves[to].FlowRate > 0 {
			// Calculate how much flow it would produce
			intermediate := timer - distance - 1
			value := v.Valves[to].Value(intermediate)
			// Calculate how much potential flow it leaves over
			potential := 0
			for nextTo, nextDistance := range v.Valves[to].Tunnels {
				if v.Valves[nextTo].Opened <= 0 && nextTo != to && nextDistance+1 <= intermediate {
					potential += v.Valves[nextTo].Value(intermediate - 1 - nextDistance)
				}
			}
			fmt.Printf("[%s] Value: %d - Remaining: %d - Potential: %d => Sum: %d\n", to, value, intermediate, potential, value+potential)
			// Find the valve with the highest sum of value and potential
			if value+potential > max {
				fmt.Printf("Max = %d + %d = %d from %s\n", value, potential, value+potential, to)
				max = value + potential
				valve = to
				remaining = timer - distance - 1
			}
		}
	}
	if valve != "" {
		fmt.Printf("\n===> Open valve [%s] - %d remaining\n\n", valve, remaining)
		v.Valves[valve].Opened = remaining
		v.Position = valve
	}
	return remaining
}

func (v Volcano) TryPath(valves []string) int {
	value := 0
	timer := 30
	i := 0
	current := "AA"
	for timer > 0 && i < len(valves) {
		next := valves[i]
		distance := v.Valves[current].Tunnels[next]
		remaining := timer - 1 - distance
		if remaining > 0 {
			value += v.Valves[next].Value(remaining)
			timer = remaining
			i++
			current = next
		} else {
			return value
		}
	}
	// Explorative code with single length path
	potential := 0
	for to, distance := range v.Valves[current].Tunnels {
		potential += v.Valves[to].Value(timer - 1 - distance)
	}
	fmt.Printf("Value: %d - Remaining: %d - Potential: %d - Sum: %d\n", value, timer, potential, value+potential)

	return value
}

func GenerateVolcano(input []string) Volcano {
	valves := make(map[string]*Valve)
	// Initialize all valves
	for _, line := range input {
		matches := fullRE.FindStringSubmatch(line)
		name := matches[1]
		flowRate, _ := strconv.Atoi(matches[2])
		valves[name] = &Valve{Name: name, FlowRate: flowRate, Tunnels: make(map[string]int)}
		for _, to := range strings.Split(matches[3], ", ") {
			valves[name].Tunnels[to] = 1
		}
	}

	return Volcano{Position: "AA", Valves: valves}
}

func RemoveFromSlice(slice []string, remove string) []string {
	newSlice := []string{}
	for _, element := range slice {
		if element != remove {
			newSlice = append(newSlice, element)
		}
	}
	return newSlice
}

func FindBestPath(v Volcano, timer int, position string, closed []string) int {
	// fmt.Printf("Going to %s - remaining time: %d - still closed: %v\n", position, timer, closed)
	max := 0
	// Exit condition
	if len(closed) == 1 {
		distance := v.Valves[position].Tunnels[closed[0]]
		return v.Valves[closed[0]].Value(timer - distance - 1)
	}
	// Recursion body
	slices.SortFunc(closed, func(a, b string) bool {
		aDistance := v.Valves[position].Tunnels[a]
		bDistance := v.Valves[position].Tunnels[b]
		return !(v.Valves[a].Value(timer-1-aDistance) <= v.Valves[b].Value(timer-1-bDistance))
	})
	for _, to := range closed {
		distance := v.Valves[position].Tunnels[to]
		// Find Valve that is not opened and can still be reached
		if v.Valves[to].Opened <= 0 && distance+1 <= timer && v.Valves[to].FlowRate > 0 {
			// Calculate how much flow it would produce
			intermediate := timer - distance - 1
			value := v.Valves[to].Value(intermediate)
			// Calculate how much leftOver flow remains
			leftOver := FindBestPath(v, intermediate, to, RemoveFromSlice(closed, to))
			// Find the valve with the highest sum of value and potential
			if value+leftOver > max {
				max = value + leftOver
			}
		}
	}
	return max
}

func calculateResultA(input []string) int {

	volcano := GenerateVolcano(input)
	volcano.Simplify()
	volcano.CompleteDistanceMap()
	fmt.Println(volcano)

	// path := []string{"DD", "BB", "JJ", "HH", "EE", "CC"}
	// for _, test := range path {
	// 	fmt.Println(test)
	// 	volcano.TryPath([]string{test})
	// }

	// optimal := []string{"DD"}
	// fmt.Println("Optimal\n======")
	// volcano.TryPath(optimal)

	// mine := []string{"JJ"}
	// fmt.Println("Mine\n======")
	// volcano.TryPath(mine)
	closed := []string{}
	for valve := range volcano.Valves {
		if valve != "AA" {
			closed = append(closed, valve)
		}
	}

	return FindBestPath(volcano, 30, "AA", closed)
}

func calculateResultB(input []string) int {

	result := 0

	return result

}

func getResult(part string) int {
	input := getInput()
	firstPart := part == "A"

	if firstPart {
		return calculateResultA(input)
	}

	return calculateResultB(input)
}

func getInput() []string {
	input := []string{}

	if file == "" {
		file = "input.txt"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}

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

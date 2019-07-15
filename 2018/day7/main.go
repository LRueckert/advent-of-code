package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var file string
var baseDuration = 60
var numWorkers = 5

type requirement struct {
	Required string
	Step     string
}

type worker struct {
	Time int
}

func (worker *worker) doWork(start int, work string) workResult {
	worker.Time = start + getDuration(work)
	return workResult{work, worker.Time}
}

type workerByTime []worker

func (s workerByTime) Len() int {
	return len(s)
}
func (s workerByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s workerByTime) Less(i, j int) bool {
	return s[i].Time < s[j].Time
}

type workResult struct {
	Work string
	Time int
}

type resultByTime []workResult

func (s resultByTime) Len() int {
	return len(s)
}
func (s resultByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s resultByTime) Less(i, j int) bool {
	return s[i].Time < s[j].Time
}

func getResult(part string) string {

	input := []requirement{}

	firstPart := part == "A"

	if file == "" {
		file = "day7.input"
	}
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		required := parts[1]
		step := parts[7]
		input = append(input, requirement{required, step})
	}

	if firstPart {
		return calculateResultA(part, input)
	}

	return strconv.Itoa(calculateResultB(part, input))
}

func calculateResultA(part string, input []requirement) string {

	result := ""

	graph := createGraph(input)

	for len(graph) > 0 {
		ready := findReadyNodes(graph)
		next := ready[0]
		result += next
		delete(graph, next)
		graph = updateGraph(graph, next)
	}

	return result
}

func calculateResultB(part string, input []requirement) int {

	graph := createGraph(input)
	time := 0
	workers := make([]worker, numWorkers)
	finished := []workResult{}

	fmt.Printf("workers: %v, baseDuration: %v \n", numWorkers, baseDuration)

	for len(graph) > 0 || len(finished) > 0 { // continue as long as there is work to be processed
		sort.Sort(workerByTime(workers))               // sort workers so that the earliest available worker is first
		ready := findReadyNodes(graph)                 // find the workpackages that are ready (sorted by workorder)
		if len(ready) > 0 && workers[0].Time <= time { // do work as long as there is ready work and available workers
			work := ready[0]
			finished = append(finished, workers[0].doWork(time, work))
			delete(graph, work)
			continue
		}
		if len(finished) > 0 { // if there is no more ready work skip ahead to when the next work is done
			sort.Sort(resultByTime(finished))
			time = finished[0].Time
			graph = updateGraph(graph, finished[0].Work)
			finished = finished[1:]
		}
	}

	return time
}

func findReadyNodes(graph map[string][]string) (result []string) {

	for node, requirements := range graph {
		if len(requirements) == 0 {
			result = append(result, node)
		}
	}
	sort.Sort(sort.StringSlice(result))
	return
}

func removeNode(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func createGraph(input []requirement) map[string][]string {
	graph := make(map[string][]string)
	for _, requirement := range input {
		if graph[requirement.Required] == nil {
			graph[requirement.Required] = []string{}
		}
		if graph[requirement.Step] == nil {
			graph[requirement.Step] = []string{requirement.Required}
		} else {
			graph[requirement.Step] = append(graph[requirement.Step], requirement.Required)
		}
	}

	return graph
}

func updateGraph(graph map[string][]string, next string) map[string][]string {

	for key, requirements := range graph {
		graph[key] = removeNode(requirements, next)
	}

	return graph
}

func getDuration(s string) int {
	return baseDuration + int(s[0]) - 64
}

func progressTime() {

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

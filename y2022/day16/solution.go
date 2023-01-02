package day16

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PartOne() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	valves, err := parseValveConnectionsAndFlows(file)
	if err != nil {
		fmt.Errorf("failed to parse valve connections and flows: %w", err)
	}

	maxPressureRelief := findMaxPressureReliefInTime(valves, 30)

	fmt.Println(fmt.Sprintf("max pressure relief: %d", maxPressureRelief))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	valves, err := parseValveConnectionsAndFlows(file)
	if err != nil {
		fmt.Errorf("failed to parse valve connections and flows: %w", err)
	}

	pressureReliefResults := findPressureReliefsInTime(valves, 26)
	sort.Slice(pressureReliefResults, func(i, j int) bool {
		return pressureReliefResults[i].relief > pressureReliefResults[j].relief
	})

	maxPressureRelief := 0
	for i, resultI := range pressureReliefResults {
		for j, resultJ := range pressureReliefResults {
			if resultI.relief+resultJ.relief < maxPressureRelief {
				break
			}
			if i == j {
				continue
			}

			maxPressureRelief = max(maxPressureRelief, max(resultI.relief, resultJ.relief))

			distinct := true
			for valve, open := range resultI.released {
				if open && resultJ.released[valve] {
					distinct = false
					break
				}
			}
			if distinct {
				maxPressureRelief = max(maxPressureRelief, resultJ.relief+resultI.relief)
			}
		}
	}

	fmt.Println(fmt.Sprintf("max pressure relief with help: %d", maxPressureRelief))

	return nil
}

type valve struct {
	name              string
	flow              int
	connections       map[string]int
	directConnections []string
}

func parseValveConnectionsAndFlows(r io.Reader) (map[string]valve, error) {
	valves := make(map[string]valve)

	s := bufio.NewScanner(r)
	for s.Scan() {
		vSplits := strings.Split(s.Text(), ";")
		vNameSplits := strings.Split(vSplits[0], " ")

		vName := vNameSplits[1]

		vFlowStr := strings.TrimPrefix(vNameSplits[4], "rate=")
		vFlow, err := strconv.Atoi(vFlowStr)
		if err != nil {
			return nil, fmt.Errorf("failed to convert valve flow rate to int: %w", err)
		}

		connections := make(map[string]int)
		vConnectionsStrs := strings.Split(vSplits[1], " ")
		for i := 5; i < len(vConnectionsStrs); i++ {
			connections[strings.TrimRight(vConnectionsStrs[i], ",")] = 1
			connections[vName] = 0
		}

		valves[vName] = valve{
			name:        vName,
			flow:        vFlow,
			connections: connections,
		}
	}

	// use floyd-warshall algorithm to pre-compute distances between valves
	for _, v := range valves {
		for _, c := range valves {
			if _, ok := v.connections[c.name]; !ok {
				v.connections[c.name] = math.MaxInt
			}
		}
	}

	for _, k := range valves {
		for _, i := range valves {
			for _, j := range valves {
				if i.connections[k.name] != math.MaxInt && k.connections[j.name] != math.MaxInt && i.connections[j.name] > i.connections[k.name]+k.connections[j.name] {
					i.connections[j.name] = i.connections[k.name] + k.connections[j.name]
				}
			}
		}
	}

	return valves, nil
}

func findMaxPressureReliefInTime(valves map[string]valve, minutes int) int {
	stack := list.New()
	stack.PushFront(node{
		curr:           "AA",
		released:       make(map[string]bool, len(valves)),
		minutes:        minutes,
		pressureRelief: 0,
	})

	maxPressureRelief := 0
	for stack.Len() > 0 {
		n := stack.Remove(stack.Front()).(node)
		if n.minutes <= 0 {
			continue
		}

		if valves[n.curr].flow > 0 {
			n.minutes--
			n.pressureRelief += valves[n.curr].flow * n.minutes
			n.released[n.curr] = true
		}
		maxPressureRelief = max(maxPressureRelief, n.pressureRelief)

		for connection, dist := range valves[n.curr].connections {
			if dist >= n.minutes {
				continue
			}
			if n.released[connection] {
				continue
			}
			if valves[connection].flow == 0 {
				continue
			}
			rc := make(map[string]bool, len(n.released))
			for key, value := range n.released {
				rc[key] = value
			}
			stack.PushBack(node{
				curr:           connection,
				released:       rc,
				minutes:        n.minutes - dist,
				pressureRelief: n.pressureRelief,
			})
		}
	}

	return maxPressureRelief
}

type pressureReliefResult struct {
	relief   int
	released map[string]bool
}

func findPressureReliefsInTime(valves map[string]valve, minutes int) []pressureReliefResult {
	stack := list.New()
	stack.PushFront(node{
		curr:           "AA",
		released:       make(map[string]bool, len(valves)),
		minutes:        minutes,
		pressureRelief: 0,
	})

	var pressureReliefs []pressureReliefResult
	for stack.Len() > 0 {
		n := stack.Remove(stack.Front()).(node)
		if n.minutes <= 0 {
			continue
		}

		if valves[n.curr].flow > 0 {
			n.minutes--
			n.pressureRelief += valves[n.curr].flow * n.minutes
			n.released[n.curr] = true
			rc := make(map[string]bool, len(n.released))
			for key, value := range n.released {
				rc[key] = value
			}
			pressureReliefs = append(pressureReliefs, pressureReliefResult{
				relief:   n.pressureRelief,
				released: rc,
			})
		}

		for connection, dist := range valves[n.curr].connections {
			if dist >= n.minutes {
				continue
			}
			if n.released[connection] {
				continue
			}
			if valves[connection].flow == 0 {
				continue
			}
			rc := make(map[string]bool, len(n.released))
			for key, value := range n.released {
				rc[key] = value
			}
			stack.PushBack(node{
				curr:           connection,
				released:       rc,
				minutes:        n.minutes - dist,
				pressureRelief: n.pressureRelief,
			})
		}
	}

	return pressureReliefs
}

type node struct {
	curr           string
	released       map[string]bool
	minutes        int
	pressureRelief int
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

package day12

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"math"
	"os"
)

func PartOne() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	heightmap := buildHeightmap(file)

	minSearchPathLength := findShortestRouteInHeightmapFromRune(heightmap, 'S')

	fmt.Println(fmt.Sprintf("min search path length: %d", minSearchPathLength))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	heightmap := buildHeightmap(file)

	minSearchPathLength := findShortestRouteInHeightmapFromRune(heightmap, 'a')

	fmt.Println(fmt.Sprintf("min search path length: %d", minSearchPathLength))

	return nil
}

type monkey struct {
	num         int
	items       []int
	op          func(int) int
	divisor     int
	trueMonkey  int
	falseMonkey int
}

func buildHeightmap(r io.Reader) [][]rune {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var heightmap [][]rune
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		heightmap = append(heightmap, row)
	}
	return heightmap
}

func findShortestRouteInHeightmapFromRune(heightmap [][]rune, startRune rune) int {
	minDist := math.MaxInt
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			if heightmap[i][j] == startRune {
				minDist = min(minDist, searchHeightmap(heightmap, i, j))
			}
		}
	}
	return minDist
}

type node struct {
	row  int
	col  int
	dist int
}

func searchHeightmap(heightmap [][]rune, startRow, startCol int) int {
	queue := list.List{}
	visited := make(map[[2]int]bool, len(heightmap)*len(heightmap[0]))
	queue.PushBack(node{row: startRow, col: startCol, dist: 0})

	minDist := math.MaxInt

	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(node)

		if heightmap[n.row][n.col] == 'E' {
			minDist = min(minDist, n.dist)
		}

		for _, dir := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			nextPoint := [2]int{n.row + dir[0], n.col + dir[1]}
			if visited[nextPoint] || nextPoint[0] < 0 || nextPoint[0] >= len(heightmap) || nextPoint[1] < 0 || nextPoint[1] >= len(heightmap[nextPoint[0]]) {
				continue
			}
			if heightmap[n.row][n.col] != 'S' && heightmap[nextPoint[0]][nextPoint[1]]-heightmap[n.row][n.col] > 1 {
				continue
			}
			if heightmap[nextPoint[0]][nextPoint[1]] == 'E' && heightmap[n.row][n.col] != 'z' {
				// cannot end until we get to 'z'
				continue
			}
			visited[nextPoint] = true
			queue.PushBack(node{row: nextPoint[0], col: nextPoint[1], dist: n.dist + 1})
		}

	}

	return minDist
}

func min(i ...int) int {
	m := i[0]
	for j := 1; j < len(i); j++ {
		if i[j] < m {
			m = i[j]
		}
	}
	return m
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

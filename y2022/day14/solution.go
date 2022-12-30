package day14

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PartOne() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	cave, _, err := buildCave(file)
	if err != nil {
		return fmt.Errorf("unable to build cave: %w", err)
	}

	numSand := 0
	infinite := false
	for !infinite {
		numSand++
		cave, infinite = simSandFallInfiniteCave(cave, 0, 500)
	}

	fmt.Println(fmt.Sprintf("num sand before infinite: %d", numSand-1))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	cave, maxRow, err := buildCave(file)
	if err != nil {
		return fmt.Errorf("unable to build cave: %w", err)
	}

	numSand := findNumSandFallBackToStart(cave, maxRow, 0, 500)

	fmt.Println(fmt.Sprintf("num sand to go back to start: %d", numSand))

	return nil
}

// returns cave, maxRow, and any error encountered
func buildCave(r io.Reader) ([][]bool, int, error) {
	cave := make([][]bool, 200)
	for i, _ := range cave {
		cave[i] = make([]bool, 1000)
	}
	maxRow := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		splits := strings.Split(s.Text(), " -> ")
		for start, finish := 0, 1; finish < len(splits); start, finish = start+1, finish+1 {
			firstSplit := strings.Split(splits[start], ",")
			secondSplit := strings.Split(splits[finish], ",")
			firstCol, err := strconv.Atoi(firstSplit[0])
			if err != nil {
				return nil, -1, fmt.Errorf("failed to parse firstCol: %w", err)
			}
			secondCol, err := strconv.Atoi(secondSplit[0])
			if err != nil {
				return nil, -1, fmt.Errorf("failed to parse secondCol: %w", err)
			}
			firstRow, err := strconv.Atoi(firstSplit[1])
			if err != nil {
				return nil, -1, fmt.Errorf("failed to parse firstRow: %w", err)
			}
			secondRow, err := strconv.Atoi(secondSplit[1])
			if err != nil {
				return nil, -1, fmt.Errorf("failed to parse secondRow: %w", err)
			}

			maxRow = max(maxRow, max(firstRow, secondRow))
			colMod := 0
			rowMod := 0
			diff := 0
			if firstCol < secondCol {
				colMod = 1
				diff = secondCol - firstCol
			} else if firstCol > secondCol {
				colMod = -1
				diff = firstCol - secondCol
			} else if firstRow < secondRow {
				rowMod = 1
				diff = secondRow - firstRow
			} else if firstRow > secondRow {
				rowMod = -1
				diff = firstRow - secondRow
			}
			for i := 0; i <= diff; i++ {
				cave[firstRow+(i*rowMod)][firstCol+(i*colMod)] = true
			}
		}
	}

	return cave, maxRow, nil
}

// returns new cave after piece of sand has fallen and whether the sand fell through
func simSandFallInfiniteCave(cave [][]bool, startRow, startCol int) ([][]bool, bool) {
	row, col := startRow, startCol
	infiniteFall := false
	for {
		if row >= len(cave) || col < 0 || col >= len(cave[row]) {
			infiniteFall = true
			break
		}
		if !cave[row][col] {
			row++
			continue
		}
		if !cave[row][col-1] {
			col--
			row++
			continue
		} else if !cave[row][col+1] {
			col++
			row++
			continue
		} else {
			cave[row-1][col] = true
			break
		}
	}
	return cave, infiniteFall
}

// returns new cave after piece of sand has fallen and whether the sand fell through
func findNumSandFallBackToStart(cave [][]bool, maxRow, startRow, startCol int) int {
	numSand := 0

	for !cave[startRow][startCol] {
		numSand++
		row, col := startRow, startCol
		for {
			if row == maxRow+2 {
				cave[row-1][col] = true
				break
			}
			if !cave[row][col] {
				row++
				continue
			}
			if !cave[row][col-1] {
				col--
				row++
				continue
			} else if !cave[row][col+1] {
				col++
				row++
				continue
			} else {
				cave[row-1][col] = true
				break
			}
		}
	}
	return numSand
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

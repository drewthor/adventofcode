package day09

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

	spotsTailVisited, err := findNumSpotsTailOfRopeTouches(file, 2)
	if err != nil {
		return fmt.Errorf("failed to find num spots rope tail visited: %w", err)
	}

	fmt.Println(fmt.Sprintf("num spots rope tail visited: %d", spotsTailVisited))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	spotsTailVisited, err := findNumSpotsTailOfRopeTouches(file, 10)
	if err != nil {
		return fmt.Errorf("failed to find num spots rope tail visited: %w", err)
	}

	fmt.Println(fmt.Sprintf("num spots rope tail visited: %d", spotsTailVisited))

	return nil
}

func findNumSpotsTailOfRopeTouches(r io.Reader, numKnots int) (int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	visited := make(map[[2]int]bool)
	tails := make([][2]int, numKnots)

	visited[tails[len(tails)-1]] = true
	spotsVisited := 1

	for i := 0; i < len(lines); i++ {
		splits := strings.Split(lines[i], " ")

		nSpots, err := strconv.Atoi(splits[1])
		if err != nil {
			return -1, fmt.Errorf("failed to get num spots to move as int: %d", nSpots)
		}

		for nSpots > 0 {
			switch splits[0] {
			case "L":
				tails[0][0] -= 1
				break
			case "R":
				tails[0][0] += 1
				break
			case "U":
				tails[0][1] += 1
				break
			case "D":
				tails[0][1] -= 1
				break
			}

			for j := 1; j < len(tails); j++ {
				if abs(tails[j][1]-tails[j-1][1]) > 1 || abs(tails[j][0]-tails[j-1][0]) > 1 {
					if tails[j-1][1]-tails[j][1] > 1 && tails[j-1][0]-tails[j][0] > 1 {
						tails[j][0] = tails[j-1][0] - 1
						tails[j][1] = tails[j-1][1] - 1
					} else if tails[j-1][1]-tails[j][1] < -1 && tails[j-1][0]-tails[j][0] > 1 {
						tails[j][0] = tails[j-1][0] - 1
						tails[j][1] = tails[j-1][1] + 1
					} else if tails[j-1][1]-tails[j][1] < -1 && tails[j-1][0]-tails[j][0] < -1 {
						tails[j][0] = tails[j-1][0] + 1
						tails[j][1] = tails[j-1][1] + 1
					} else if tails[j-1][1]-tails[j][1] > 1 && tails[j-1][0]-tails[j][0] < -1 {
						tails[j][0] = tails[j-1][0] + 1
						tails[j][1] = tails[j-1][1] - 1
					} else if tails[j-1][1]-tails[j][1] > 1 {
						tails[j][0] = tails[j-1][0]
						tails[j][1] = tails[j-1][1] - 1
					} else if tails[j-1][1]-tails[j][1] < -1 {
						tails[j][0] = tails[j-1][0]
						tails[j][1] = tails[j-1][1] + 1
					} else if tails[j-1][0]-tails[j][0] > 1 {
						tails[j][0] = tails[j-1][0] - 1
						tails[j][1] = tails[j-1][1]
					} else if tails[j-1][0]-tails[j][0] < -1 {
						tails[j][0] = tails[j-1][0] + 1
						tails[j][1] = tails[j-1][1]
					}
				}
			}

			if !visited[tails[len(tails)-1]] {
				spotsVisited++
				visited[tails[len(tails)-1]] = true
			}
			nSpots--
		}
	}
	return spotsVisited, nil
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

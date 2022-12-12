package day08

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PartOne() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	treeMatrix, err := buildTreeMatrix(file)
	if err != nil {
		return fmt.Errorf("failed to build tree matrix: %w", err)
	}

	visibilityFromLT := findTreesVisibilityFromLeftAndTop(treeMatrix)
	visibilityFromRB := findTreeVisibilityFromRightAndBottom(treeMatrix)

	numTreesVisible := 0
	for i := 0; i < len(treeMatrix); i++ {
		for j := 0; j < len(treeMatrix[i]); j++ {
			if visibilityFromLT[i][j] || visibilityFromRB[i][j] {
				numTreesVisible++
			}
		}
	}

	fmt.Println(fmt.Sprintf("num trees visible from outside grid: %d", numTreesVisible))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	treeMatrix, err := buildTreeMatrix(file)
	if err != nil {
		return fmt.Errorf("failed to build tree matrix: %w", err)
	}

	scoreLT := findScenicScoreOfTreesFromLeftAndTop(treeMatrix)
	scoreRB := findScenicScoreOfTreesFromRightAndBottom(treeMatrix)

	highestScore := 0
	for i := 0; i < len(treeMatrix); i++ {
		for j := 0; j < len(treeMatrix[i]); j++ {
			if scoreLT[i][j]*scoreRB[i][j] > highestScore {
				highestScore = scoreLT[i][j] * scoreRB[i][j]
			}
		}
	}

	fmt.Println(fmt.Sprintf("highest scenic score of tree: %d", highestScore))

	return nil
}

func buildTreeMatrix(r io.Reader) ([][]int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var matrix [][]int
	for i := 0; i < len(lines); i++ {
		var curr []int
		for j := 0; j < len(lines[i]); j++ {
			i, err := strconv.Atoi(string(lines[i][j]))
			if err != nil {
				return nil, fmt.Errorf("failed to parse input; expected integer: %w", err)
			}
			curr = append(curr, i)
		}
		matrix = append(matrix, curr)
	}
	return matrix, nil
}

func findTreesVisibilityFromLeftAndTop(treeMatrix [][]int) [][]bool {
	if len(treeMatrix) == 0 {
		return nil
	}

	visibility := make([][]bool, len(treeMatrix))

	highestTreeColFromTop := make([]int, len(treeMatrix[0]))
	for i := 0; i < len(highestTreeColFromTop); i++ {
		highestTreeColFromTop[i] = -1
	}
	for i := 0; i < len(treeMatrix); i++ {
		visibility[i] = make([]bool, len(treeMatrix[i]))
		highestTreeFromLeft := -1
		for j := 0; j < len(treeMatrix[i]); j++ {
			visibility[i][j] = treeMatrix[i][j] > highestTreeFromLeft || treeMatrix[i][j] > highestTreeColFromTop[j]
			highestTreeFromLeft = max(highestTreeFromLeft, treeMatrix[i][j])
			highestTreeColFromTop[j] = max(highestTreeColFromTop[j], treeMatrix[i][j])
		}
	}

	return visibility
}

func findTreeVisibilityFromRightAndBottom(treeMatrix [][]int) [][]bool {
	if len(treeMatrix) == 0 {
		return nil
	}
	visibility := make([][]bool, len(treeMatrix))

	highestTreeColFromBot := make([]int, len(treeMatrix[0]))
	for i := 0; i < len(highestTreeColFromBot); i++ {
		highestTreeColFromBot[i] = -1
	}
	for i := len(treeMatrix) - 1; i >= 0; i-- {
		visibility[i] = make([]bool, len(treeMatrix[i]))
		highestTreeFromRight := -1
		for j := len(treeMatrix[i]) - 1; j >= 0; j-- {
			visibility[i][j] = treeMatrix[i][j] > highestTreeFromRight || treeMatrix[i][j] > highestTreeColFromBot[j]
			highestTreeFromRight = max(highestTreeFromRight, treeMatrix[i][j])
			highestTreeColFromBot[j] = max(highestTreeColFromBot[j], treeMatrix[i][j])
		}
	}

	return visibility
}

func findScenicScoreOfTreesFromLeftAndTop(treeMatrix [][]int) [][]int {
	if len(treeMatrix) == 0 {
		return nil
	}

	scenicScore := make([][]int, len(treeMatrix))

	for i := 0; i < len(treeMatrix); i++ {
		scenicScore[i] = make([]int, len(treeMatrix[i]))
		for j := 0; j < len(treeMatrix[i]); j++ {
			leftScore := 0
			for k := j - 1; k >= 0; k-- {
				leftScore++
				if treeMatrix[i][j] <= treeMatrix[i][k] {
					break
				}
			}
			rightScore := 0
			for k := i - 1; k >= 0; k-- {
				rightScore++
				if treeMatrix[i][j] <= treeMatrix[k][j] {
					break
				}
			}
			scenicScore[i][j] = leftScore * rightScore
		}
	}

	return scenicScore
}

func findScenicScoreOfTreesFromRightAndBottom(treeMatrix [][]int) [][]int {
	if len(treeMatrix) == 0 {
		return nil
	}

	scenicScore := make([][]int, len(treeMatrix))

	for i := len(treeMatrix) - 1; i >= 0; i-- {
		scenicScore[i] = make([]int, len(treeMatrix[i]))
		for j := len(treeMatrix[i]) - 1; j >= 0; j-- {
			rightScore := 0
			for k := j + 1; k < len(treeMatrix[i]); k++ {
				rightScore++
				if treeMatrix[i][j] <= treeMatrix[i][k] {
					break
				}
			}
			botScore := 0
			for k := i + 1; k < len(treeMatrix); k++ {
				botScore++
				if treeMatrix[i][j] <= treeMatrix[k][j] {
					break
				}
			}
			scenicScore[i][j] = rightScore * botScore
		}
	}

	return scenicScore
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

package day07

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"math"
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

	fileTree, err := buildFileTree(file)
	if err != nil {
		return fmt.Errorf("failed to build file tree: %w", err)
	}

	sum, _ := findSumDirectoriesSizeBelowSize(fileTree, 100000)

	fmt.Println(fmt.Sprintf("sum direcctories below size: %d", sum))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	fileTree, err := buildFileTree(file)
	if err != nil {
		return fmt.Errorf("failed to build file tree: %w", err)
	}

	calcDirSizes(fileTree)

	smallestDirSizeAboveTarget := findSmallestDirectorySizeAboveTarget(fileTree, fileTree.size-(70000000-30000000))

	fmt.Println(fmt.Sprintf("smallest directory size above target: %d", smallestDirSizeAboveTarget))

	return nil
}

type dir struct {
	name    string
	size    int
	subDirs map[string]*dir
}

func buildFileTree(r io.Reader) (*dir, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	rootDir := &dir{subDirs: make(map[string]*dir)}
	workingDir := rootDir
	stack := list.List{}
	stack.PushFront(workingDir)
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			workingDir.name = strings.Split(lines[i], " ")[2]
			continue
		}
		if strings.HasPrefix(lines[i], "$ ls") {
			for i = i + 1; i < len(lines); i++ {
				if strings.HasPrefix(lines[i], "dir") {
					dirName := strings.Split(lines[i], " ")[1]
					d := dir{name: dirName, subDirs: make(map[string]*dir)}
					workingDir.subDirs[dirName] = &d
				} else if strings.HasPrefix(lines[i], "$") {
					i--
					break
				} else {
					fileSize, err := strconv.Atoi(strings.Split(lines[i], " ")[0])
					if err != nil {
						return rootDir, fmt.Errorf("parsed non-integer: %w", err)
					}
					workingDir.size += fileSize
				}
			}
		} else if lines[i] == "$ cd .." {
			stack.Remove(stack.Front())
			workingDir = stack.Front().Value.(*dir)
		} else if strings.HasPrefix(lines[i], "$ cd") {
			dirName := strings.Split(lines[i], " ")[2]
			workingDir = workingDir.subDirs[dirName]
			stack.PushFront(workingDir)
		}
	}
	return rootDir, nil
}

func calcDirSizes(d *dir) int {
	for _, subDir := range d.subDirs {
		subDirSize := calcDirSizes(subDir)
		d.size += subDirSize
	}

	return d.size
}

func findSumDirectoriesSizeBelowSize(d *dir, targetSize int) (int, int) {
	sum := 0
	for _, subDir := range d.subDirs {
		subDirSum, subDirSize := findSumDirectoriesSizeBelowSize(subDir, targetSize)
		d.size += subDirSize
		sum += subDirSum
	}

	if d.size <= targetSize {
		sum += d.size
	}

	return sum, d.size
}

func findSmallestDirectorySizeAboveTarget(d *dir, targetSize int) int {
	smallestDirSizeAboveTarget := math.MaxInt
	for _, subDir := range d.subDirs {
		smallestSubDirSizeAboveTarget := findSmallestDirectorySizeAboveTarget(subDir, targetSize)
		smallestDirSizeAboveTarget = min(smallestDirSizeAboveTarget, smallestSubDirSizeAboveTarget)
	}

	if d.size >= targetSize {
		smallestDirSizeAboveTarget = min(smallestDirSizeAboveTarget, d.size)
	}

	return smallestDirSizeAboveTarget
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

package day01

import (
	"bufio"
	"fmt"
	"io"
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

	elfCals, err := findSortedElfCalories(file)
	if err != nil {
		return fmt.Errorf("failed to elf calories: %w", err)
	}

	maxElfCal := 0
	if len(elfCals) > 0 {
		maxElfCal = elfCals[0]
	}

	fmt.Println(fmt.Sprintf("max elf calories: %d", maxElfCal))

	//if _, err := w.Write([]byte(fmt.Sprintf("max elf calories: %d", maxElfCal))); err != nil {
	//	return fmt.Errorf("failed to write output to writer: %w", err)
	//}

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	elfCals, err := findSortedElfCalories(file)
	if err != nil {
		return fmt.Errorf("failed to elf calories: %w", err)
	}

	sumTop3MaxElfCals := 0
	if len(elfCals) >= 3 {
		sumTop3MaxElfCals += elfCals[0] + elfCals[1] + elfCals[2]
	}

	fmt.Println(fmt.Sprintf("max elf calories: %d", sumTop3MaxElfCals))

	//if _, err := w.Write([]byte(fmt.Sprintf("max elf calories: %d", sumTop3MaxElfCals))); err != nil {
	//	return fmt.Errorf("failed to write output to writer: %w", err)
	//}

	return nil
}

// findSortedElfCalories returns a descending sorted list of elf calories that they are carrying
func findSortedElfCalories(r io.Reader) ([]int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var elfCals []int
	currElfCal := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			elfCals = append(elfCals, currElfCal)
			currElfCal = 0
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse elf calorie line: %s to int: %w", line, err)
		}

		currElfCal += cal
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCals)))

	return elfCals, nil

}

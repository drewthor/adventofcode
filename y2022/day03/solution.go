package day03

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PartOne() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	incorrectItemPrioritySum := calculateIncorrectItemPrioritySum(file)

	fmt.Println(fmt.Sprintf("incorrect item priority sum: %d", incorrectItemPrioritySum))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	groupBadgeSum := calculateGroupBadgeSum(file)

	fmt.Println(fmt.Sprintf("group badge sum: %d", groupBadgeSum))

	return nil
}

func calculateIncorrectItemPrioritySum(r io.Reader) int {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	incorrectItemPrioritySum := 0
	for _, line := range lines {
		m := make(map[rune]bool, len(line)/2)

		for i := 0; i < len(line)/2; i++ {
			m[rune(line[i])] = true
		}
		for i := len(line) / 2; i < len(line); i++ {
			if m[rune(line[i])] == true {
				incorrectItemPrioritySum += calculatePriorityOfItem(rune(line[i]))
				break
			}
		}
	}

	return incorrectItemPrioritySum
}

func calculateGroupBadgeSum(r io.Reader) int {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	groupBadgeSum := 0
	for i := 0; i < len(lines); i += 3 {
		var groupRucksacks []map[rune]bool
		for j := 0; j < 3; j++ {
			m := make(map[rune]bool)
			for k := 0; k < len(lines[i+j]); k++ {
				m[rune(lines[i+j][k])] = true
			}
			groupRucksacks = append(groupRucksacks, m)
		}
		for item, _ := range groupRucksacks[0] {
			if groupRucksacks[1][item] == true && groupRucksacks[2][item] == true {
				groupBadgeSum += calculatePriorityOfItem(item)
				break
			}
		}
	}

	return groupBadgeSum
}

func calculatePriorityOfItem(r rune) int {
	if r < 'a' {
		return int(r - 'A' + 27)
	} else {
		return int(r - 'a' + 1)
	}
}

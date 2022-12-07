package day06

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

	startOfPacketMarker := findStartOfPacketMarker(file)

	fmt.Println(fmt.Sprintf("start of packet marker: %d", startOfPacketMarker))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	startOfMessageMarker := findStartOfMessageMarker(file)

	fmt.Println(fmt.Sprintf("start of message marker: %d", startOfMessageMarker))

	return nil
}

func findStartOfPacketMarker(r io.Reader) int {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if len(lines) > 1 {
		return -1
	}

	return findIndexOfUniqueNLastChars(lines[0], 4)
}

func findStartOfMessageMarker(r io.Reader) int {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if len(lines) > 1 {
		return -1
	}

	return findIndexOfUniqueNLastChars(lines[0], 14)
}

func findIndexOfUniqueNLastChars(s string, n int) int {
	lastNUnique := make(map[rune]int, n)
	for i, char := range s {
		if len(lastNUnique) == n {
			return i
		}
		lastNUnique[char]++
		if i >= n {
			if lastNUnique[rune(s[i-n])] == 1 {
				delete(lastNUnique, rune(s[i-n]))
			} else {
				lastNUnique[rune(s[i-n])]--
			}
		}
	}

	return -1

}

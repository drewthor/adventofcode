package day13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
)

func PartOne() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	lines, err := parseLines(file)
	if err != nil {
		return fmt.Errorf("unable to parse lines: %w", err)
	}

	sum := findSumCorrectlyOrderedPairs(lines)

	fmt.Println(fmt.Sprintf("sum incorrect order pair indexes: %d", sum))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	lines, err := parseLines(file)
	if err != nil {
		return fmt.Errorf("unable to parse lines: %w", err)
	}

	decoderKey, err := findDecoderKey(lines)
	if err != nil {
		return fmt.Errorf("unable to find decoder key: %w", err)
	}

	fmt.Println(fmt.Sprintf("decoder key: %d", decoderKey))

	return nil
}

func parseLines(r io.Reader) ([]any, error) {
	var lines []any

	s := bufio.NewScanner(r)
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		var t any
		if err := json.Unmarshal([]byte(s.Text()), &t); err != nil {
			return nil, fmt.Errorf("failed to parse line as json: %w", err)
		}
		lines = append(lines, t)
	}

	return lines, nil
}

func findSumCorrectlyOrderedPairs(lines []any) int {
	sum := 0

	for i, pairNum := 0, 1; i < len(lines); i, pairNum = i+2, pairNum+1 {
		if isOrdered(lines[i], lines[i+1]) >= 0 {
			sum += pairNum
		}
	}

	return sum
}

func findDecoderKey(lines []any) (int, error) {
	sort.Slice(lines, func(i, j int) bool {
		return isOrdered(lines[i], lines[j]) == 1
	})

	firstKeyLoc := -1
	var firstDecoderKey any
	if err := json.Unmarshal([]byte(`[[2]]`), &firstDecoderKey); err != nil {
		return -1, fmt.Errorf("unable to unmarshal first decoder key: %w", err)
	}
	secondKeyLoc := -1
	var secondDecoderKey any
	if err := json.Unmarshal([]byte(`[[6]]`), &secondDecoderKey); err != nil {
		return -1, fmt.Errorf("unable to unmarshal first decoder key: %w", err)
	}
	for i, line := range lines {
		if firstKeyLoc == -1 && isOrdered(firstDecoderKey, line) >= 0 {
			firstKeyLoc = i + 1
		}
		if secondKeyLoc == -1 && isOrdered(secondDecoderKey, line) >= 0 {
			secondKeyLoc = i + 2
		}
	}

	return firstKeyLoc * secondKeyLoc, nil
}

// 0: continue checking order 1: ordered -1: not ordered
func isOrdered(first, second any) int {
	switch {
	case reflect.TypeOf(first).Kind() == reflect.Slice && reflect.TypeOf(second).Kind() == reflect.Slice:
		for i, j := range first.([]any) {
			if i < len(second.([]any)) {
				switch isOrdered(j, second.([]any)[i]) {
				case -1:
					return -1
				case 1:
					return 1
				}
			}
		}
		return isOrdered(float64(len(first.([]any))), float64(len(second.([]any))))
	case reflect.TypeOf(first).Kind() == reflect.Float64 && reflect.TypeOf(second).Kind() == reflect.Float64:
		if first.(float64) < second.(float64) {
			return 1
		} else if first.(float64) == second.(float64) {
			return 0
		} else {
			return -1
		}
	case reflect.TypeOf(first).Kind() == reflect.Slice && reflect.TypeOf(second).Kind() == reflect.Float64:
		return isOrdered(first, []any{second.(float64)})
	case reflect.TypeOf(first).Kind() == reflect.Float64 && reflect.TypeOf(second).Kind() == reflect.Slice:
		return isOrdered([]any{first.(float64)}, second)
	default:
		return 1
	}
}

package day04

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

	numFullyOverlappingShiftPairs, err := calcNumFullyOverlappingShiftPairs(file)
	if err != nil {
		return fmt.Errorf("failed to calculate num fully overlapping shift pairs: %w", err)
	}

	fmt.Println(fmt.Sprintf("num fully overlapping shift pairs: %d", numFullyOverlappingShiftPairs))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	numOverlappingShiftPairs, err := calcNumOverlappingShiftPairs(file)
	if err != nil {
		return fmt.Errorf("failed to calculate num overlapping shift pairs: %w", err)
	}

	fmt.Println(fmt.Sprintf("num overlapping shift pairs: %d", numOverlappingShiftPairs))

	return nil
}

func calcNumFullyOverlappingShiftPairs(r io.Reader) (int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	numFullyOverlappingShiftPairs := 0
	for _, line := range lines {
		shiftPairs := strings.Split(line, ",")

		firstShiftPair := strings.Split(shiftPairs[0], "-")
		secondShiftPair := strings.Split(shiftPairs[1], "-")

		firstShiftStart, err := strconv.Atoi(firstShiftPair[0])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}
		firstShiftEnd, err := strconv.Atoi(firstShiftPair[1])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}

		secondShiftStart, err := strconv.Atoi(secondShiftPair[0])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}
		secondShiftEnd, err := strconv.Atoi(secondShiftPair[1])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}

		if firstShiftStart <= secondShiftStart && firstShiftEnd >= secondShiftEnd {
			numFullyOverlappingShiftPairs++
		} else if secondShiftStart <= firstShiftStart && secondShiftEnd >= firstShiftEnd {
			numFullyOverlappingShiftPairs++
		}
	}

	return numFullyOverlappingShiftPairs, nil
}

func calcNumOverlappingShiftPairs(r io.Reader) (int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	numOverlappingShiftPairs := 0
	for _, line := range lines {
		shiftPairs := strings.Split(line, ",")

		firstShiftPair := strings.Split(shiftPairs[0], "-")
		secondShiftPair := strings.Split(shiftPairs[1], "-")

		firstShiftStart, err := strconv.Atoi(firstShiftPair[0])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}
		firstShiftEnd, err := strconv.Atoi(firstShiftPair[1])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}

		secondShiftStart, err := strconv.Atoi(secondShiftPair[0])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}
		secondShiftEnd, err := strconv.Atoi(secondShiftPair[1])
		if err != nil {
			return -1, fmt.Errorf("encountered non-integer when expecting integer: %w", err)
		}

		if secondShiftStart >= firstShiftStart && secondShiftStart <= firstShiftEnd {
			numOverlappingShiftPairs++
		} else if firstShiftStart >= secondShiftStart && firstShiftStart <= secondShiftEnd {
			numOverlappingShiftPairs++
		} else if firstShiftStart <= secondShiftStart && firstShiftEnd >= secondShiftEnd {
			numOverlappingShiftPairs++
		} else if secondShiftStart <= firstShiftStart && secondShiftEnd >= firstShiftEnd {
			numOverlappingShiftPairs++
		}
	}

	return numOverlappingShiftPairs, nil
}

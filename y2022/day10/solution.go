package day10

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

	sumCycles, err := sumSignalStrengthsDuringCycles(file, []int{20, 60, 100, 140, 180, 220})
	if err != nil {
		return fmt.Errorf("failed to sum signal strengths during cycles: %w", err)
	}

	fmt.Println(fmt.Sprintf("sum signal strengths during cycles: %d", sumCycles))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	image, err := calcImageOfSignalTransmission(file)
	if err != nil {
		return fmt.Errorf("failed to calc image of signal transmission: %w", err)
	}

	fmt.Println("image of signal transmission")
	for i := 0; i < len(image); i++ {
		fmt.Println(image[i])
	}

	return nil
}

func sumSignalStrengthsDuringCycles(r io.Reader, cycles []int) (int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	cycle := 0
	signal := 1

	sum := 0
	for i := 0; i < len(lines); i++ {
		splits := strings.Split(lines[i], " ")

		cycleIncr := 1
		signalIncr := 0
		switch splits[0] {
		case "addx":
			cycleIncr = 2
			add, err := strconv.Atoi(splits[1])
			if err != nil {
				return -1, fmt.Errorf("failed to parse signal add to int: %w", err)
			}
			signalIncr = add
			break
		}

		for j := 0; j < cycleIncr; j++ {
			cycle++
			for k := 0; k < len(cycles); k++ {
				if cycles[k] == cycle {
					sum += cycle * signal
					break
				}
			}
		}
		signal += signalIncr
	}
	return sum, nil
}

func calcImageOfSignalTransmission(r io.Reader) ([6][40]string, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var image [6][40]string

	cycle := 0
	signal := 1

	for i := 0; i < len(lines); i++ {
		splits := strings.Split(lines[i], " ")

		cycleIncr := 1
		signalIncr := 0
		switch splits[0] {
		case "addx":
			cycleIncr = 2
			add, err := strconv.Atoi(splits[1])
			if err != nil {
				return [6][40]string{}, fmt.Errorf("failed to parse signal add to int: %w", err)
			}
			signalIncr = add
			break
		}

		for j := 0; j < cycleIncr; j++ {
			cycle++
			xpos := (cycle - 1) % 40
			fill := "."
			if abs(xpos-signal) <= 1 {
				fill = "#"
			}
			image[(cycle-1)/40][xpos] = fill
		}
		signal += signalIncr
	}
	return image, nil
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

package day01

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func PartOne() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	score, err := findScoreOfRPSFirstStrategy(file)
	if err != nil {
		return fmt.Errorf("failed to elf calories: %w", err)
	}

	fmt.Println(fmt.Sprintf("score after play: %d", score))

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

	score, err := findScoreOfRPSSecondStrategy(file)
	if err != nil {
		return fmt.Errorf("failed to elf calories: %w", err)
	}

	fmt.Println(fmt.Sprintf("score after play: %d", score))

	//if _, err := w.Write([]byte(fmt.Sprintf("max elf calories: %d", sumTop3MaxElfCals))); err != nil {
	//	return fmt.Errorf("failed to write output to writer: %w", err)
	//}

	return nil
}

// A,B,C opponent play rock,paper,scissors X,Y,Z what you play for rock, paper scissors
func findScoreOfRPSFirstStrategy(r io.Reader) (int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	score := 0
	for _, line := range lines {
		line = strings.TrimRight(line, "\n")
		if len(line) == 0 {
			continue
		}

		plays := strings.Split(line, " ")
		if len(plays) < 2 {
			return -1, fmt.Errorf("invalid play in input; expecting 2 plays on line")
		}

		score += getScoreOfPlay(plays[0], plays[1])
	}

	return score, nil
}

// A,B,C opponent play rock,paper,scissors X,Y,Z what you need to accomplish lose, draw, win
func findScoreOfRPSSecondStrategy(r io.Reader) (int, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	score := 0
	for _, line := range lines {
		line = strings.TrimRight(line, "\n")
		if len(line) == 0 {
			continue
		}

		plays := strings.Split(line, " ")
		if len(plays) < 2 {
			return -1, fmt.Errorf("invalid play in input; expecting 2 plays on line")
		}

		score += getScoreOfIntendedOutcome(plays[0], plays[1])
	}

	return score, nil
}

// A,B,C opponent play rock,paper,scissors X,Y,Z what you play for rock, paper scissors
func getScoreOfPlay(opp string, play string) int {
	score := 0

	switch play {
	case "X":
		// rock
		score += 1
		if opp == "A" {
			score += 3
		} else if opp == "C" {
			score += 6
		}
		break
	case "Y":
		// paper
		score += 2
		if opp == "A" {
			score += 6
		} else if opp == "B" {
			score += 3
		}
		break
	case "Z":
		// scissors
		score += 3
		if opp == "B" {
			score += 6
		} else if opp == "C" {
			score += 3
		}
		break
	}
	return score
}

// A,B,C opponent play rock,paper,scissors X,Y,Z what outcome you need lose, draw, win
func getScoreOfIntendedOutcome(opp string, play string) int {
	score := 0

	switch play {
	case "X":
		// lose
		if opp == "A" {
			score += 3
		} else if opp == "B" {
			score += 1
		} else if opp == "C" {
			score += 2
		}
		break
	case "Y":
		// draw
		score += 3
		if opp == "A" {
			score += 1
		} else if opp == "B" {
			score += 2
		} else if opp == "C" {
			score += 3
		}
		break
	case "Z":
		// win
		score += 6
		if opp == "A" {
			score += 2
		} else if opp == "B" {
			score += 3
		} else if opp == "C" {
			score += 1
		}
		break
	}
	return score
}

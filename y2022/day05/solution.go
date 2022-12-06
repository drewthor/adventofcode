package day05

import (
	"bufio"
	"container/list"
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

	crateStacks, crateMoves, err := parseCrateStacksAndMovements(file)
	if err != nil {
		return fmt.Errorf("failed to parse crate stacks and movements: %w", err)
	}

	crateStacks = handleCrateStacksMoves9000(crateStacks, crateMoves)

	var topOfStacks []string
	for _, crateStack := range crateStacks {
		if crateStack.Len() > 0 {
			topOfStacks = append(topOfStacks, string(crateStack.Front().Value.(rune)))
		}
	}

	fmt.Println(fmt.Sprintf("top of each stack: %v", topOfStacks))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	crateStacks, crateMoves, err := parseCrateStacksAndMovements(file)
	if err != nil {
		return fmt.Errorf("failed to parse crate stacks and movements: %w", err)
	}

	crateStacks = handleCrateStacksMoves9001(crateStacks, crateMoves)

	var topOfStacks []string
	for _, crateStack := range crateStacks {
		if crateStack.Len() > 0 {
			topOfStacks = append(topOfStacks, string(crateStack.Front().Value.(rune)))
		}
	}

	fmt.Println(fmt.Sprintf("top of each stack: %v", topOfStacks))

	return nil
}

type crateMove struct {
	source    int
	dest      int
	numCrates int
}

func parseCrateStacksAndMovements(r io.Reader) ([]*list.List, []crateMove, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var crateStacks []*list.List
	var crateMoves []crateMove
	for i, line := range lines {
		if i == 0 {
			// 3 chars per crate is 1 space in-between
			numCrateStacks := (len(line) + 1) / 4
			for j := 0; j < numCrateStacks; j++ {
				crateStacks = append(crateStacks, list.New())
			}
		}

		if len(line) == 0 {
			continue
		}

		if (rune(line[0]) == ' ' && rune(line[1]) != '1') || rune(line[0]) == '[' {
			for j, crateNum := 0, 0; j < len(line); j, crateNum = j+4, crateNum+1 {
				if rune(line[j]) == '[' {
					crateStacks[crateNum].PushBack(rune(line[j+1]))
				}
			}
		} else if rune(line[1]) == '1' && rune(line[0]) == ' ' {
			continue
		} else if rune(line[0]) == 'm' {
			words := strings.Split(line, " ")
			numCrates, err := strconv.Atoi(words[1])
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert num crates to integer: %w", err)
			}
			source, err := strconv.Atoi(words[3])
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert crate source to integer: %w", err)
			}
			dest, err := strconv.Atoi(words[5])
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert crate dest to integer: %w", err)
			}
			crateMoves = append(crateMoves, crateMove{
				source:    source,
				dest:      dest,
				numCrates: numCrates,
			})
		}
	}

	return crateStacks, crateMoves, nil
}

func handleCrateStacksMoves9000(crateStacks []*list.List, crateMoves []crateMove) []*list.List {
	for _, crateMove := range crateMoves {
		for i := 0; i < crateMove.numCrates; i++ {
			el := crateStacks[crateMove.source-1].Front()
			crateStacks[crateMove.dest-1].PushFront(el.Value.(rune))
			crateStacks[crateMove.source-1].Remove(el)
		}
	}

	return crateStacks
}

func handleCrateStacksMoves9001(crateStacks []*list.List, crateMoves []crateMove) []*list.List {
	for _, crateMove := range crateMoves {
		curr := crateStacks[crateMove.source-1].Front()
		destPtr := crateStacks[crateMove.dest-1].PushFront(curr.Value.(rune))
		crateStacks[crateMove.source-1].Remove(curr)
		for i := 1; i < crateMove.numCrates; i++ {
			curr = crateStacks[crateMove.source-1].Front()
			destPtr = crateStacks[crateMove.dest-1].InsertAfter(curr.Value.(rune), destPtr)
			crateStacks[crateMove.source-1].Remove(curr)
		}
	}

	return crateStacks
}

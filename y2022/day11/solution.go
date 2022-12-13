package day11

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	monkeyStarts, err := buildMonkeyStarts(file)
	if err != nil {
		return fmt.Errorf("failed to build monkey starts: %w", err)
	}

	monkeyBusinessLevel := findMonkeyBusinessLevel(monkeyStarts, 20, 3, math.MaxInt)

	fmt.Println(fmt.Sprintf("monkey business level: %d", monkeyBusinessLevel))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	monkeyStarts, err := buildMonkeyStarts(file)
	if err != nil {
		return fmt.Errorf("failed to build monkey starts: %w", err)
	}

	lcm := 1
	for _, monkey := range monkeyStarts {
		lcm *= monkey.divisor
	}

	monkeyBusinessLevel := findMonkeyBusinessLevel(monkeyStarts, 10000, 1, lcm)

	fmt.Println(fmt.Sprintf("monkey business level: %d", monkeyBusinessLevel))

	return nil
}

type monkey struct {
	num         int
	items       []int
	op          func(int) int
	divisor     int
	trueMonkey  int
	falseMonkey int
}

func buildMonkeyStarts(r io.Reader) ([]monkey, error) {
	var lines []string

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var monkeys []monkey
	for i := 0; i < len(lines); i += 7 {
		mNumSplits := strings.Split(lines[i], " ")
		mNum := strings.TrimRight(mNumSplits[1], ":")
		num, err := strconv.Atoi(mNum)
		if err != nil {
			return nil, fmt.Errorf("failed to convert monkey number to int: %w", err)
		}

		var items []int
		itemNums := strings.Split(strings.TrimLeft(lines[i+1], "Starting items: "), ", ")
		for _, itemNum := range itemNums {
			iNum, err := strconv.Atoi(itemNum)
			if err != nil {
				return nil, fmt.Errorf("failed to convert item number to int: %w", err)
			}
			items = append(items, iNum)
		}

		var op func(int) int
		ops := strings.Split(strings.TrimSpace(lines[i+2]), " ")
		opSymbol := ops[4]
		opStr := ops[5]
		switch opSymbol {
		case "*":
			if opStr == "old" {
				op = func(x int) int {
					return x * x
				}
			} else {
				opNum, err := strconv.Atoi(opStr)
				if err != nil {
					return nil, fmt.Errorf("failed to convert monkey op num to int: %w", err)
				}
				op = func(x int) int {
					return x * opNum
				}
			}
			break
		case "+":
			if opStr == "old" {
				op = func(x int) int {
					return x + x
				}
			} else {
				opNum, err := strconv.Atoi(opStr)
				if err != nil {
					return nil, fmt.Errorf("failed to convert monkey op num to int: %w", err)
				}
				op = func(x int) int {
					return x + opNum
				}
			}
			break
		}

		div, err := strconv.Atoi(strings.Split(strings.TrimSpace(lines[i+3]), " ")[3])
		if err != nil {
			return nil, fmt.Errorf("failed to convert monkey divisor to int: %w", err)
		}

		trueMonkey, err := strconv.Atoi(strings.Split(strings.TrimSpace(lines[i+4]), " ")[5])
		if err != nil {
			return nil, fmt.Errorf("failed to convert true monkey to int: %w", err)
		}

		falseMonkey, err := strconv.Atoi(strings.Split(strings.TrimSpace(lines[i+5]), " ")[5])
		if err != nil {
			return nil, fmt.Errorf("failed to convert false monkey to int: %w", err)
		}

		monkeys = append(monkeys, monkey{
			num:         num,
			items:       items,
			op:          op,
			divisor:     div,
			trueMonkey:  trueMonkey,
			falseMonkey: falseMonkey,
		})
	}

	return monkeys, nil
}

func findMonkeyBusinessLevel(monkeys []monkey, rounds int, worryLevel int, mod int) int {
	monkeyActivity := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			for k := 0; k < len(monkeys[j].items); k++ {
				monkeyActivity[j]++
				newWorryLevel := monkeys[j].op(monkeys[j].items[k]) / worryLevel % mod
				if newWorryLevel%monkeys[j].divisor == 0 {
					monkeys[monkeys[j].trueMonkey].items = append(monkeys[monkeys[j].trueMonkey].items, newWorryLevel)
				} else {
					monkeys[monkeys[j].falseMonkey].items = append(monkeys[monkeys[j].falseMonkey].items, newWorryLevel)
				}
			}
			monkeys[j].items = nil
		}
	}
	sort.Ints(monkeyActivity)
	return monkeyActivity[len(monkeyActivity)-1] * monkeyActivity[len(monkeyActivity)-2]
}

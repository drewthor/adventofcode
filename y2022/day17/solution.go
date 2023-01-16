package day17

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PartOne(r io.Reader) (int, error) {
	jetStream := parseJetStream(r)

	rockHeight := rockHeightAfterRocksFall(2022, jetStream)

	return rockHeight, nil
}

func PartTwo(r io.Reader) (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return -1, fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	jetStream := parseJetStream(r)
	if err != nil {
		return -1, fmt.Errorf("failed to parse jet stream: %w", err)
	}

	cycle, caveTop, endJet, caveHeightBeginCycle, cycleHeight, startRock, endRock := detectCycle(jetStream)
	if !cycle {
		return -1, fmt.Errorf("failed to find cycle in rock falling")
	}

	totalRocks := 1_000_000_000_000
	cycleLength := endRock - startRock
	numCycles := (totalRocks - startRock - 1) / cycleLength
	finalStart := startRock + (numCycles * cycleLength)
	cyclesHeight := numCycles * cycleHeight

	rockHeightLastRocks := maxHeightAfterRockFall(caveTop, finalStart, totalRocks, endJet, jetStream)

	return caveHeightBeginCycle + cyclesHeight + rockHeightLastRocks - len(caveTop) + 1, nil
}

func parseJetStream(r io.Reader) string {
	s := bufio.NewScanner(r)
	for s.Scan() {
		return s.Text()
	}
	return ""
}

func rockHeightAfterRocksFall(numRocks int, jetStream string) int {
	return maxHeightAfterRockFall([][7]bool{{true, true, true, true, true, true, true}}, 0, numRocks, 0, jetStream)
}

// returns the cave top, jetStream index, cave start height, cycle height, rock start and rock index end for the first cycle found when rocks fall
func detectCycle(jetStream string) (bool, [][7]bool, int, int, int, int, int) {
	numRocks := 10000
	var cave [][7]bool

	for i := 0; i < 50; i++ {
		cave = append(cave, [7]bool{})
	}

	for i := 0; i < len(cave[0]); i++ {
		cave[0][i] = true
	}

	rockItersSeen := make(map[rockIter]caveState)

	maxHeight := 0
	for i, j := 0, 0; i < numRocks; i++ {
		startx := 0
		var rock [][2]int

		rockIndex := i % 5

		switch rockIndex {
		case 0:
			startx = 2
			rock = [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
			break
		case 1:
			startx = 3
			rock = [][2]int{{0, 0}, {1, 1}, {-1, 1}, {0, 2}}
			break
		case 2:
			startx = 2
			rock = [][2]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
			break
		case 3:
			startx = 2
			rock = [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
			break
		case 4:
			startx = 2
			rock = [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
			break
		}

		var snapshot [16][7]bool
		s, e := max(maxHeight-len(snapshot)+1, 0), max(maxHeight+1, len(snapshot))
		copy(snapshot[:], cave[s:e])
		r := rockIter{
			top:       snapshot,
			jetIndex:  j % len(jetStream),
			rockIndex: rockIndex,
		}
		if prevCave, ok := rockItersSeen[r]; ok {
			caveTop := make([][7]bool, len(snapshot))
			copy(caveTop, r.top[:])
			return true, caveTop, r.jetIndex, prevCave.caveHeight, maxHeight - prevCave.caveHeight, prevCave.numRock, i
		}

		rockItersSeen[r] = caveState{
			numRock:    i,
			caveHeight: maxHeight,
		}

		x, y := startx, maxHeight+4
		falling := true
		for ; falling; j++ {
			// first move by jet stream and then down
			jetMove := 0
			switch jetStream[j%len(jetStream)] {
			case '>':
				jetMove = 1
				break
			case '<':
				jetMove = -1
				break
			}
			jetCollide := false
			for _, piece := range rock {
				newx, newy := jetMove+piece[0]+x, piece[1]+y
				if newx < 0 || newx >= len(cave[0]) || cave[newy][newx] {
					jetCollide = true
					break
				}
			}
			if !jetCollide {
				x += jetMove
			}

			downCollide := false
			for _, piece := range rock {
				if y > 0 && cave[piece[1]+y-1][piece[0]+x] {
					downCollide = true
					break
				}
			}

			if y == 0 || downCollide {
				falling = false
				for _, piece := range rock {
					newx, newy := piece[0]+x, piece[1]+y
					cave[newy][newx] = true
					maxHeight = max(maxHeight, newy)
				}
			} else {
				y--
			}
		}
		if len(cave)-maxHeight < 8 {
			for k := 0; k < 8; k++ {
				cave = append(cave, [7]bool{})
			}
		}
	}
	return false, nil, -1, -1, -1, -1, -1
}

func maxHeightAfterRockFall(cave [][7]bool, startRock, endRock int, startJet int, jetStream string) int {
	maxHeight := 0
	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			if cave[i][j] {
				maxHeight = i
			}
		}
	}
	for i, j := startRock, startJet; i < endRock; i++ {
		if len(cave)-maxHeight < 8 {
			for k := 0; k < 8; k++ {
				cave = append(cave, [7]bool{})
			}
		}
		startx := 0
		var rock [][2]int

		switch i % 5 {
		case 0:
			startx = 2
			rock = [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
			break
		case 1:
			startx = 3
			rock = [][2]int{{0, 0}, {1, 1}, {-1, 1}, {0, 2}}
			break
		case 2:
			startx = 2
			rock = [][2]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
			break
		case 3:
			startx = 2
			rock = [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
			break
		case 4:
			startx = 2
			rock = [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
			break
		}

		x, y := startx, maxHeight+4
		falling := true
		for ; falling; j++ {
			// first move by jet stream and then down
			jetMove := 0
			switch jetStream[j%len(jetStream)] {
			case '>':
				jetMove = 1
				break
			case '<':
				jetMove = -1
				break
			}
			jetCollide := false
			for _, piece := range rock {
				newx, newy := jetMove+piece[0]+x, piece[1]+y
				if newx < 0 || newx >= len(cave[0]) || cave[newy][newx] {
					jetCollide = true
					break
				}
			}
			if !jetCollide {
				x += jetMove
			}

			downCollide := false
			for _, piece := range rock {
				if y > 0 && cave[piece[1]+y-1][piece[0]+x] {
					downCollide = true
					break
				}
			}

			if y == 0 || downCollide {
				falling = false
				for _, piece := range rock {
					newx, newy := piece[0]+x, piece[1]+y
					cave[newy][newx] = true
					maxHeight = max(maxHeight, newy)
				}
			} else {
				y--
			}
		}
	}
	return maxHeight
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type rockIter struct {
	top       [16][7]bool
	jetIndex  int
	rockIndex int
}

type caveState struct {
	numRock    int
	caveHeight int
}

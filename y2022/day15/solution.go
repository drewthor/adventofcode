package day15

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

	sensors, beacons, err := parseSensorsAndBeacons(file)
	if err != nil {
		return fmt.Errorf("unable to parse sensors and beacons: %w", err)
	}

	numInvalidLocsOnRow := findNumInvalidLocationsOnRow(sensors, beacons, 2000000)

	fmt.Println(fmt.Sprintf("num invalid locations on row: %d", numInvalidLocsOnRow))

	return nil
}

func PartTwo() error {
	file, err := os.Open("input.txt")
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer file.Close()

	sensors, beacons, err := parseSensorsAndBeacons(file)
	if err != nil {
		return fmt.Errorf("unable to parse sensors and beacons: %w", err)
	}

	tuningFrequency := findTuningFrequencyMissingBeaconLocation(sensors, beacons)

	fmt.Println(fmt.Sprintf("tuning frequency missing beacon: %d", tuningFrequency))

	return nil
}

type loc struct {
	x, y int
}

type interval struct {
	start, finish int
}

func parseSensorsAndBeacons(r io.Reader) ([]loc, []loc, error) {
	var sensors []loc
	var beacons []loc

	s := bufio.NewScanner(r)
	for s.Scan() {
		lines := s.Text()
		sbSplits := strings.Split(strings.ReplaceAll(lines, ",", ""), ":")

		sSplits := strings.Split(sbSplits[0], " ")
		sxStr := strings.TrimPrefix(sSplits[2], "x=")
		sx, err := strconv.Atoi(sxStr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert sensor x to int: %w", err)
		}
		syStr := strings.TrimPrefix(sSplits[3], "y=")
		sy, err := strconv.Atoi(syStr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert sensor y to int: %w", err)
		}

		sensors = append(sensors, loc{
			x: sx,
			y: sy,
		})

		bSplits := strings.Split(sbSplits[1], " ")
		bxStr := strings.TrimPrefix(bSplits[5], "x=")
		bx, err := strconv.Atoi(bxStr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert beacon x to int: %w", err)
		}
		byStr := strings.TrimPrefix(bSplits[6], "y=")
		by, err := strconv.Atoi(byStr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert beacon y to int: %w", err)
		}

		beacons = append(beacons, loc{
			x: bx,
			y: by,
		})
	}

	return sensors, beacons, nil
}

func findInvalidLocationsOnRow(sensors, beacons []loc, row int) []interval {
	var invalidRanges []interval

	for i := 0; i < len(sensors); i++ {
		manDist := calcManDist(sensors[i], beacons[i])

		distToRow := abs(row - sensors[i].y)
		if distToRow > manDist {
			continue
		}

		invalidRanges = append(invalidRanges, interval{
			start:  sensors[i].x - (manDist - distToRow),
			finish: sensors[i].x + (manDist - distToRow),
		})
	}

	return invalidRanges
}

func findInvalidLocationsOnCol(sensors, beacons []loc, col int) []interval {
	var invalidRanges []interval

	for i := 0; i < len(sensors); i++ {
		manDist := calcManDist(sensors[i], beacons[i])

		distToRow := abs(col - sensors[i].x)
		if distToRow > manDist {
			continue
		}

		invalidRanges = append(invalidRanges, interval{
			start:  sensors[i].y - (manDist - distToRow),
			finish: sensors[i].y + (manDist - distToRow),
		})
	}

	return invalidRanges
}

func findNumInvalidLocationsOnRow(sensors, beacons []loc, row int) int {
	condensed := condenseRanges(findInvalidLocationsOnRow(sensors, beacons, row))

	numInvalidLocations := 0
	for _, invalidRange := range condensed {
		numInvalidLocations += invalidRange.finish - invalidRange.start
	}

	return numInvalidLocations
}

func findNumInvalidConstrainedLocationsOnRow(sensors, beacons []loc, row int, low, high int) int {
	constrainedAndCondensed := condenseRanges(constrainRanges(findInvalidLocationsOnRow(sensors, beacons, row), low, high))

	numInvalidLocations := 0
	for _, invalidRange := range constrainedAndCondensed {
		numInvalidLocations += invalidRange.finish - invalidRange.start
	}

	return numInvalidLocations
}

func findNumInvalidConstrainedLocationsOnCol(sensors, beacons []loc, col int, low, high int) int {
	constrainedAndCondensed := condenseRanges(constrainRanges(findInvalidLocationsOnCol(sensors, beacons, col), low, high))

	numInvalidLocations := 0
	for _, invalidRange := range constrainedAndCondensed {
		numInvalidLocations += invalidRange.finish - invalidRange.start
	}

	return numInvalidLocations
}

func findTuningFrequencyMissingBeaconLocation(sensors, beacons []loc) int {
	for i := 0; i < 4000000; i++ {
		numInvalid := findNumInvalidConstrainedLocationsOnRow(sensors, beacons, i, 0, 4000000)
		if numInvalid < 4000000 {
			for j := 0; j <= 4000000; j++ {
				if findNumInvalidConstrainedLocationsOnCol(sensors, beacons, j, 0, 4000000) < 4000000 {
					return j*4000000 + i
				}
			}
		}
	}
	return -1
}

func calcManDist(l1, l2 loc) int {
	return abs(l2.y-l1.y) + abs(l2.x-l1.x)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func condenseRanges(ranges []interval) []interval {
	var condensed []interval

	overlap := true
	for overlap {
		overlap = false

		condensed = nil
		for _, r := range ranges {
			foundOverlap := false
			for i, c := range condensed {
				foundOverlap = r.start <= c.finish && r.finish >= c.start
				if foundOverlap {
					condensed[i].start = min(condensed[i].start, r.start)
					condensed[i].finish = max(condensed[i].finish, r.finish)
					break
				}
			}
			if !foundOverlap {
				condensed = append(condensed, r)
			}
			overlap = overlap || foundOverlap
		}

		ranges = condensed
	}

	return condensed
}

func constrainRanges(ranges []interval, low, high int) []interval {
	for i, _ := range ranges {
		ranges[i].start = max(ranges[i].start, low)
		ranges[i].finish = min(ranges[i].finish, high)
	}

	return ranges
}

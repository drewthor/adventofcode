package day17

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	sampleFile, err := os.Open("sample-input.txt")
	if err != nil {
		t.Errorf("could not open sample-input file: %v", err)
	}
	defer sampleFile.Close()

	sampleWant := 3068
	sampleGot, err := PartOne(sampleFile)
	if err != nil {
		t.Error(err)
	}
	if sampleGot != sampleWant {
		t.Errorf("sample result wanted %v but got %v", sampleWant, sampleGot)
	}

	realFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("could not open input file: %v", err)
	}
	defer realFile.Close()

	realWant := 3179
	realGot, err := PartOne(realFile)
	if err != nil {
		t.Error(err)
	}
	if realGot != realWant {
		t.Errorf("result wanted %v but got %v", realWant, realGot)
	}

	t.Logf("part one: %v", realGot)
}

func TestPartTwo(t *testing.T) {
	sampleFile, err := os.Open("sample-input.txt")
	if err != nil {
		t.Errorf("could not open sample-input file: %v", err)
	}
	defer sampleFile.Close()

	sampleWant := 1514285714288
	sampleGot, err := PartTwo(sampleFile)
	if err != nil {
		t.Error(err)
	}
	if sampleGot != sampleWant {
		t.Errorf("sample result wanted %v but got %v", sampleWant, sampleGot)
	}

	realFile, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("could not open input file: %v", err)
	}
	defer realFile.Close()

	realWant := 1567723342929
	realGot, err := PartTwo(realFile)
	if err != nil {
		t.Error(err)
	}
	if realGot != realWant {
		t.Errorf("result wanted %v but got %v", realWant, realGot)
	}

	t.Logf("part two: %v", realGot)
}

package day02

import (
	"testing"
)

func TestMin(t *testing.T) {
	expect := uint(100)
	result, err := min(300, 122, 100, 440, 231)
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

func TestCalcArea(t *testing.T) {
	expect := uint(58)
	result, err := calcWrappingArea(2, 3, 4)
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

func TestSplitLine(t *testing.T) {
	expect := [3]uint64{23, 72, 12}
	result, err := splitLine("23x72x12")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %v but got %v", expect, result)
	}
}

func TestPart1(t *testing.T) {
	var expect uint = 1_606_483
	result, err := part1("./input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

func TestPart2(t *testing.T) {
	var expect uint64 = 3_842_356
	result, err := part2("./input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

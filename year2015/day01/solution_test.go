package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expect := 74
	result, err := part1("./input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

func TestPart2(t *testing.T) {
	expect := 1795
	result, err := part2("./input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

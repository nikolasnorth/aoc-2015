package day04

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expect := 117946
	result := part1("ckczppom")
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

func TestPart2(t *testing.T) {
	expect := 3938038
	result := part2("ckczppom")
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

package day06

import (
	"reflect"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	type test struct {
		s         string
		expectCmd Command
		expectLo  Coordinate
		expectHi  Coordinate
	}

	tests := []test{
		{
			s:         "turn off 660,55 through 986,197",
			expectCmd: turnOff,
			expectLo:  Coordinate{x: 660, y: 55},
			expectHi:  Coordinate{x: 986, y: 197},
		},
		{
			s:         "toggle 322,558 through 977,958",
			expectCmd: toggle,
			expectLo:  Coordinate{x: 322, y: 558},
			expectHi:  Coordinate{x: 977, y: 958},
		},
		{
			s:         "turn on 226,196 through 599,390",
			expectCmd: turnOn,
			expectLo:  Coordinate{x: 226, y: 196},
			expectHi:  Coordinate{x: 599, y: 390},
		},
	}

	for _, test := range tests {
		resultCmd, resultLo, resultHi, err := parseInstruction(test.s)
		if err != nil {
			t.Errorf("got unexpected error: %q", err)
		}
		if resultCmd != test.expectCmd {
			t.Errorf("expected instruction %d but got %d", test.expectCmd, resultCmd)
		}
		if !reflect.DeepEqual(resultLo, test.expectLo) {
			t.Errorf("expected %v but got %v", test.expectLo, resultLo)
		}
		if !reflect.DeepEqual(resultHi, test.expectHi) {
			t.Errorf("expected %v but got %v", test.expectHi, resultHi)
		}
	}
}

func TestPart1(t *testing.T) {
	expect := 400410
	result, err := part1("input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

func TestPart2(t *testing.T) {
	expect := 15343601
	result, err := part2("input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

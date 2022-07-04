package day05

import (
	"testing"
)

func TestContainsConsecutiveChars(t *testing.T) {
	type test struct {
		s      string
		expect bool
	}

	tests := []test{
		{s: "abcdeef", expect: true},
		{s: "abcdef", expect: false},
	}

	for _, test := range tests {
		result := containsConsecutiveChars(test.s)
		if result != test.expect {
			t.Errorf("expected %t but got %t", test.expect, result)
		}
	}
}

func TestContainsCharSandwich(t *testing.T) {
	type test struct {
		s      string
		expect bool
	}

	tests := []test{
		{s: "xyx", expect: true},
		{s: "abcdefeghi", expect: true},
		{s: "aaa", expect: true},
		{s: "abc", expect: false},
	}

	for _, test := range tests {
		result := containsCharSandwich(test.s)
		if result != test.expect {
			t.Errorf("expected %t but got %t", test.expect, result)
		}
	}
}

func TestContainsCharPair(t *testing.T) {
	type test struct {
		s      string
		expect bool
	}

	tests := []test{
		{s: "aabcdefgaa", expect: true},
		{s: "xyxy", expect: true},
		{s: "aaa", expect: false},
	}

	for _, test := range tests {
		result := containsCharPair(test.s)
		if result != test.expect {
			t.Errorf("expected %t but got %t", test.expect, result)
		}
	}
}

func TestPart1(t *testing.T) {
	expect := 258
	result, err := part1("./input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if expect != result {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

func TestPart2(t *testing.T) {
	expect := 53
	result, err := part2("./input.txt")
	if err != nil {
		t.Errorf("got unexpected error: %q", err)
	}
	if result != expect {
		t.Errorf("expected %d but got %d", expect, result)
	}
}

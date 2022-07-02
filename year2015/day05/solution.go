package day05

import (
	"bufio"
	"os"
	"strings"
)

// containsConsecutiveChars checks whether s contains a pair of consecutive matching characters.
func containsConsecutiveChars(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
	}
	return false
}

// containsCharSandwich checks whether s contains some pair of identical characters with exactly one character in between.
func containsCharSandwich(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}

// containsPair checks whether s contains a pair of some two characters without any overlap.
func containsCharPair(s string) bool {
	for i := 0; i < len(s)-1; i += 2 {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	for i := 1; i < len(s)-1; i += 2 {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

func part1(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	vowels := [...]rune{'a', 'e', 'i', 'o', 'u'}
	antiPatterns := [...]string{"ab", "cd", "pq", "xy"}
	niceCount := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		numVowels := 0
		for _, vowel := range vowels {
			numVowels += strings.Count(s.Text(), string(vowel))
		}
		containsAntiPattern := false
		for _, pattern := range antiPatterns {
			if strings.Contains(s.Text(), pattern) {
				containsAntiPattern = true
				break
			}
		}
		if numVowels >= 3 && !containsAntiPattern && containsConsecutiveChars(s.Text()) {
			niceCount += 1
		}
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	return niceCount, nil
}

func part2(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	niceCount := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		if containsCharPair(s.Text()) && containsCharSandwich(s.Text()) {
			niceCount += 1
		}
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	return niceCount, nil
}

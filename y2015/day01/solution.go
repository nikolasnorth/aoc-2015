package day01

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func part1(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	floor := 0
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "(":
			floor++
		case ")":
			floor--
		default:
			msg := fmt.Sprintf("received unexpected token: %q", s.Text())
			return 0, errors.New(msg)
		}
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	return floor, nil
}

func part2(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	floor, pos := 0, 1
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "(":
			floor++
		case ")":
			floor--
		default:
			msg := fmt.Sprintf("received unexpected token: %q", s.Text())
			return 0, errors.New(msg)
		}
		if floor == -1 {
			return pos, nil
		}
		pos++
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	return 0, errors.New("basement is never reached")
}

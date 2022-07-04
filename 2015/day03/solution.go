package day03

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Coordinate struct {
	x int
	y int
}

func part1(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	visited := make(map[Coordinate]bool)
	c := Coordinate{x: 0, y: 0}
	visited[c] = true

	x, y := 0, 0

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "^":
			y += 1
		case ">":
			x += 1
		case "v":
			y -= 1
		case "<":
			x -= 1
		default:
			msg := fmt.Sprintf("invalid character %q", s.Text())
			return 0, errors.New(msg)
		}
		c = Coordinate{x: x, y: y}
		visited[c] = true
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	numKeys := 0
	for range visited {
		numKeys += 1
	}

	return numKeys, nil
}

func part2(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	visited := make(map[Coordinate]bool)
	c := Coordinate{x: 0, y: 0}
	visited[c] = true

	xSanta, ySanta := 0, 0
	xRobo, yRobo := 0, 0
	isSanta := true

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	for s.Scan() {
		switch s.Text() {
		case "^":
			if isSanta {
				ySanta += 1
			} else {
				yRobo += 1
			}
		case ">":
			if isSanta {
				xSanta += 1
			} else {
				xRobo += 1
			}
		case "v":
			if isSanta {
				ySanta -= 1
			} else {
				yRobo -= 1
			}
		case "<":
			if isSanta {
				xSanta -= 1
			} else {
				xRobo -= 1
			}
		default:
			msg := fmt.Sprintf("invalid character %q", s.Text())
			return 0, errors.New(msg)
		}

		if isSanta {
			c = Coordinate{x: xSanta, y: ySanta}
		} else {
			c = Coordinate{x: xRobo, y: yRobo}
		}
		visited[c] = true
		isSanta = !isSanta
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	numKeys := 0
	for range visited {
		numKeys += 1
	}

	return numKeys, nil
}

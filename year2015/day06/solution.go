package day06

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command uint8

type Coordinate struct {
	x int
	y int
}

const (
	turnOff Command = 0
	turnOn  Command = 1
	toggle  Command = 2
)

func parseInstruction(s string) (Command, Coordinate, Coordinate, error) {
	words := strings.Split(s, " ")
	loStr, hiStr := words[len(words)-3], words[len(words)-1]

	loSplit := strings.Split(loStr, ",")
	hiSplit := strings.Split(hiStr, ",")

	loX, err := strconv.Atoi(loSplit[0])
	if err != nil {
		return 0, Coordinate{}, Coordinate{}, nil
	}

	loY, err := strconv.Atoi(loSplit[1])
	if err != nil {
		return 0, Coordinate{}, Coordinate{}, err
	}

	hiX, err := strconv.Atoi(hiSplit[0])
	if err != nil {
		return 0, Coordinate{}, Coordinate{}, err
	}

	hiY, err := strconv.Atoi(hiSplit[1])
	if err != nil {
		return 0, Coordinate{}, Coordinate{}, err
	}

	lo := Coordinate{x: loX, y: loY}
	hi := Coordinate{x: hiX, y: hiY}

	if words[0] == "toggle" {
		return toggle, lo, hi, nil
	} else if words[0] == "turn" && words[1] == "on" {
		return turnOn, lo, hi, nil
	} else if words[0] == "turn" && words[1] == "off" {
		return turnOff, lo, hi, nil
	}

	msg := fmt.Sprintf("could not parse invalid instruction: %q", s)
	return 0, Coordinate{}, Coordinate{}, errors.New(msg)
}

func part1(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	grid := [1000][1000]bool{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		cmd, lo, hi, err := parseInstruction(s.Text())
		if err != nil {
			return 0, err
		}

		switch cmd {
		case turnOff:
			for i := lo.x; i <= hi.x; i++ {
				for j := lo.y; j <= hi.y; j++ {
					grid[i][j] = false
				}
			}
		case turnOn:
			for i := lo.x; i <= hi.x; i++ {
				for j := lo.y; j <= hi.y; j++ {
					grid[i][j] = true
				}
			}
		case toggle:
			for i := lo.x; i <= hi.x; i++ {
				for j := lo.y; j <= hi.y; j++ {
					grid[i][j] = !grid[i][j]
				}
			}
		}
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	n := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] {
				n++
			}
		}
	}

	return n, nil
}

func part2(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	grid := [1000][1000]uint8{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		cmd, lo, hi, err := parseInstruction(s.Text())
		if err != nil {
			return 0, err
		}

		switch cmd {
		case turnOff:
			for i := lo.x; i <= hi.x; i++ {
				for j := lo.y; j <= hi.y; j++ {
					if grid[i][j] > 0 {
						grid[i][j] -= 1
					}
				}
			}
		case turnOn:
			for i := lo.x; i <= hi.x; i++ {
				for j := lo.y; j <= hi.y; j++ {
					grid[i][j] += 1
				}
			}
		case toggle:
			for i := lo.x; i <= hi.x; i++ {
				for j := lo.y; j <= hi.y; j++ {
					grid[i][j] += 2
				}
			}
		}
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	n := 0
	for i, row := range grid {
		for j := range row {
			n += int(grid[i][j])
		}
	}

	return n, nil
}

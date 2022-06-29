package day02

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(nums ...uint) (uint, error) {
	if len(nums) == 0 {
		return 0, errors.New("min() was given no arguments")
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min, nil
}

func calcWrappingArea(l, w, h uint) (uint, error) {
	x := l * w
	y := w * h
	z := h * l

	minSide, err := min(x, y, z)
	if err != nil {
		return 0, err
	}

	return 2*x + 2*y + 2*z + minSide, nil
}

func splitLine(s string) ([3]uint64, error) {
	var args [3]uint64

	split := strings.Split(s, "x")
	if len(split) != 3 {
		msg := fmt.Sprintf("invalid input: %q", s)
		return args, errors.New(msg)
	}

	l, err := strconv.ParseUint(split[0], 10, 64)
	if err != nil {
		return args, err
	}

	w, err := strconv.ParseUint(split[1], 10, 64)
	if err != nil {
		return args, err
	}

	h, err := strconv.ParseUint(split[2], 10, 64)
	if err != nil {
		return args, err
	}

	args[0], args[1], args[2] = l, w, h
	return args, nil
}

func part1(filename string) (uint, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var totalArea uint
	s := bufio.NewScanner(f)
	for s.Scan() {
		args, err := splitLine(s.Text())
		if err != nil {
			return 0, err
		}

		l, w, h := uint(args[0]), uint(args[1]), uint(args[2])
		area, err := calcWrappingArea(l, w, h)
		if err != nil {
			return 0, err
		}

		totalArea += area
	}

	err = s.Err()
	if err != nil {
		return 0, err
	}

	return totalArea, nil
}

func part2(filename string) (uint64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var ribbonLen uint64 = 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		args, err := splitLine(s.Text())
		if err != nil {
			return 0, err
		}

		sort.Slice(args[:], func(i, j int) bool { return args[i] < args[j] })
		volume := args[0] * args[1] * args[2]
		ribbonLen += (2*args[0] + 2*args[1]) + volume
	}

	return ribbonLen, nil
}

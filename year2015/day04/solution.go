package day04

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(input string) int {
	for i := 1; i < math.MaxInt; i++ {
		data := []byte(input + strconv.Itoa(i))
		hash := md5.Sum(data)
		hashStr := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hashStr, "00000") {
			return i
		}
	}
	return -1
}

func part2(input string) int {
	for i := 1; i < math.MaxInt; i++ {
		data := []byte(input + strconv.Itoa(i))
		hash := md5.Sum(data)
		hashStr := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hashStr, "000000") {
			return i
		}
	}
	return -1
}

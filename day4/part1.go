package day4

import (
	"math"
	"strconv"
	"strings"
)

func toNums(s string) []int {
	s = strings.TrimSpace(s)

	result := []int{}

	splitted := strings.Split(s, " ")

	for _, i := range splitted {
		trimmed := strings.TrimSpace(i)

		if len(trimmed) > 0 {
			num, err := strconv.Atoi(trimmed)

			if err != nil {
				panic(err)
			}

			result = append(result, num)
		}
	}

	return result
}

func intersectionCount(a []int, b []int) int {
	if len(a) > len(b) {
		tmp := b
		b = a
		a = tmp
	}

	dict := make(map[int]bool)
	count := 0

	for _, item := range a {
		dict[item] = true
	}

	for _, item := range b {
		_, ok := dict[item]

		if ok {
			count += 1
		}
	}

	return count
}

func Part1(content string) int {
	lines := strings.Split(content, "\n")

	result := 0

	for i := range lines {
		line := strings.TrimSpace(lines[i])
		nums := strings.Split(line, ": ")[1]
		pair := strings.Split(nums, " | ")

		winning := toNums(pair[0])
		ihave := toNums(pair[1])

		count := intersectionCount(winning, ihave)

		result += int(math.Pow(2, float64(count-1)))
	}

	return result
}

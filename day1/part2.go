package day1

import (
	"aoc/utils"
	"sort"
	"strings"
)

func Part2(content string) int {
	elves := strings.Split(content, "\r\n\r\n")

	acc := []int{}

	for _, elf := range elves {
		calories := strings.Split(elf, "\r\n")

		c := 0

		for _, k := range calories {
			c += utils.LineToNums(k)[0]
		}

		acc = append(acc, c)
	}

	sort.Ints(acc)

	k := len(acc)

	return acc[k-1] + acc[k-2] + acc[k-3]
}

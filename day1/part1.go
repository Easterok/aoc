package day1

import (
	"aoc/utils"
	"strings"
)

func Part1(content string) int {
	elves := strings.Split(content, "\r\n\r\n")

	result := 0

	for _, elf := range elves {
		calories := strings.Split(elf, "\r\n")

		c := 0

		for _, k := range calories {
			c += utils.LineToNums(k)[0]
		}

		result = max(c, result)
	}

	return result
}

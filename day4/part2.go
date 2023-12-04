package day4

import (
	"strings"
)

func Part2(content string) int {
	lines := strings.Split(content, "\n")

	copies := make(map[int]int)

	result := 0

	for i := range lines {
		line := strings.TrimSpace(lines[i])
		nums := strings.Split(line, ": ")[1]
		pair := strings.Split(nums, " | ")

		winning := toNums(pair[0])
		ihave := toNums(pair[1])

		count := intersectionCount(winning, ihave)

		result += 1

		if count > 0 {
			repeat := copies[i]

			for k := 0; k < repeat+1; k++ {
				for times := 0; times < count; times++ {
					prev, ok := copies[i+times+1]

					if ok {
						copies[i+times+1] = prev + 1
					} else {
						copies[i+times+1] = 1
					}
				}
			}
		}
	}

	for _, k := range copies {
		result += k
	}

	return result
}

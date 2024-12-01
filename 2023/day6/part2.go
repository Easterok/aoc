package day6

import (
	"aoc/utils"
	"strings"
)

func Part2(content string) uint64 {
	lines := strings.Split(content, "\r\n")

	time := utils.LineToUint64(strings.ReplaceAll(strings.ReplaceAll(lines[0], "Time:", ""), " ", ""))
	distance := utils.LineToUint64(strings.ReplaceAll(strings.ReplaceAll(lines[1], "Distance: ", ""), " ", ""))

	acc := make([]uint64, len(time))

	for index, duration := range time {
		winDistance := distance[index]

		delay := uint64(1)

		for delay < duration {
			travelTime := duration - delay
			speed := delay // mm/ms

			winnable := speed*travelTime > winDistance

			if winnable {
				acc[index] = acc[index] + 1
			}

			delay += 1
		}
	}

	result := uint64(1)

	for _, i := range acc {
		result = result * i
	}

	return result
}

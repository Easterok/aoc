package day6

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Part1(content string) int {
	lines := strings.Split(content, "\r\n")

	time := utils.LineToNums(strings.ReplaceAll(lines[0], "Time:", ""))
	distance := utils.LineToNums(strings.ReplaceAll(lines[1], "Distance: ", ""))

	acc := make([]int, len(time))

	for index, duration := range time {
		winDistance := distance[index]

		delay := 1

		fmt.Println(winDistance, duration)

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

	fmt.Println(acc)

	result := 1

	for _, i := range acc {
		result = result * i
	}

	return result
}

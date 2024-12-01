package day9

import (
	"aoc/utils"
	"strings"
)

func extrapolateBackwards(nums []int) int {
	first := nums[0]

	acc := []int{}

	done := true

	for i := 1; i < len(nums); i++ {
		a := nums[i-1]
		b := nums[i]

		next := b - a

		acc = append(acc, next)

		done = next == 0 && done
	}

	if done {
		return first
	}

	return first - extrapolateBackwards(acc)
}

func Part2(content string) int {
	history := strings.Split(content, "\r\n")

	result := 0

	for _, k := range history {
		nums := utils.LineToNums(k)

		result = result + extrapolateBackwards(nums)
	}

	return result
}

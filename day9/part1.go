package day9

import (
	"aoc/utils"
	"strings"
)

func extrapolate(nums []int) int {
	last := nums[len(nums)-1]

	acc := []int{}

	for i := 1; i < len(nums); i++ {
		a := nums[i-1]
		b := nums[i]

		acc = append(acc, b-a)
	}

	done := true

	index := 0

	for index < len(acc) && done {
		done = acc[index] == 0 && done

		index += 1
	}

	if done {
		return last
	}

	return last + extrapolate(acc)
}

func Part1(content string) int {
	history := strings.Split(content, "\r\n")

	result := 0

	for _, k := range history {
		nums := utils.LineToNums(k)

		result = result + extrapolate(nums)
	}

	return result
}

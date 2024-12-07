package aoc_2024

import (
	"fmt"
	"strconv"
	"strings"
)

func toInts(s ...string) []int {
	acc := make([]int, len(s))

	for index, str := range s {
		num, _ := strconv.Atoi(str)

		acc[index] = num
	}

	return acc
}

type Day7Operation = string

const (
	ADD           Day7Operation = "+"
	MULTPLY       Day7Operation = "*"
	CONCATENATION Day7Operation = "||"
)

func day7evaluate(nums []int, total int, operation Day7Operation) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return nums[0] == total
	}

	var res int

	if operation == ADD {
		res = nums[0] + nums[1]
	} else if operation == MULTPLY {
		res = nums[0] * nums[1]
	}

	nextNums := []int{res}
	nextNums = append(nextNums, nums[2:]...)

	return day7evaluate(nextNums, total, ADD) || day7evaluate(nextNums, total, MULTPLY)
}

func Day7P1(s string) string {
	count := 0

	content := strings.Split(s, "\n")

	for _, line := range content {
		a := strings.Split(line, ": ")
		total := toInts(a[0])[0]
		nums := toInts(strings.Split(a[1], " ")...)

		if day7evaluate(nums, total, ADD) || day7evaluate(nums, total, MULTPLY) {
			count += total
		}
	}

	return fmt.Sprintf("%d", count)
}

func day7evaluate2(nums []int, total int, operation Day7Operation) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return nums[0] == total
	}

	var res int

	if operation == ADD {
		res = nums[0] + nums[1]
	} else if operation == MULTPLY {
		res = nums[0] * nums[1]
	} else if operation == CONCATENATION {
		res = toNums(fmt.Sprintf("%d%d", nums[0], nums[1]))[0]
	}

	nextNums := []int{res}
	nextNums = append(nextNums, nums[2:]...)

	return day7evaluate2(nextNums, total, ADD) || day7evaluate2(nextNums, total, MULTPLY) || day7evaluate2(nextNums, total, CONCATENATION)
}

func Day7P2(s string) string {
	count := 0

	content := strings.Split(s, "\n")

	for _, line := range content {
		a := strings.Split(line, ": ")
		total := toInts(a[0])[0]
		nums := toInts(strings.Split(a[1], " ")...)

		if day7evaluate2(nums, total, ADD) || day7evaluate2(nums, total, MULTPLY) || day7evaluate2(nums, total, CONCATENATION) {
			count += total
		}
	}

	return fmt.Sprintf("%d", count)
}

package aoc_2024

import (
	"strconv"
	"strings"
)

func Day2P1(content string) int {
	count := 0

	for _, line := range strings.Split(content, "\n") {
		var prev int
		inc := true

		shouldAdd := true

		for index, c := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(c)

			if index == 0 {
				prev = num
				continue
			}

			if prev == num {
				shouldAdd = false
				break
			}

			if index == 1 {
				inc = prev < num

				if inc && prev+3 < num {
					shouldAdd = false
					break
				}

				if !inc && prev-3 > num {
					shouldAdd = false
					break
				}

				prev = num

				continue
			}

			if inc && (prev > num || prev < num-3) {
				shouldAdd = false
				break
			}

			if !inc && (prev < num || prev-3 > num) {
				shouldAdd = false
				break
			}

			prev = num
		}

		if shouldAdd {
			count += 1
		}
	}

	return count
}

func toNums(l string) []int {
	s := strings.Split(l, " ")

	acc := []int{}

	for _, num := range s {
		n, _ := strconv.Atoi(num)

		acc = append(acc, n)
	}

	return acc
}

func evaluate(nums []int) bool {
	inc := true
	var prev int

	for index, num := range nums {
		if index == 0 {
			prev = num
			continue
		}

		if prev == num {
			return false
		}

		if index == 1 {
			inc = prev < num

			if inc && prev+3 < num {
				return false
			}

			if !inc && prev-3 > num {
				return false
			}

			prev = num

			continue
		}

		if inc && (prev > num || prev < num-3) {
			return false
		}

		if !inc && (prev < num || prev-3 > num) {
			return false
		}

		prev = num
	}

	return true
}

func Day2P2(content string) int {
	count := 0

	for _, line := range strings.Split(content, "\n") {
		nums := toNums(line)

		if evaluate(nums) {
			count += 1
		} else {
			ev := false
			// :(
			for i := range nums {
				s := []int{}
				cp := nums[:i]
				a := nums[i+1:]
				s = append(s, cp...)
				s = append(s, a...)

				if evaluate(s) {
					ev = true
					break
				}
			}

			if ev {
				count += 1
			}
		}
	}

	return count
}

package aoc_2024

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1P1(s string) string {
	left := []int{}
	right := []int{}

	for _, line := range strings.Split(s, "\n") {
		splitted := strings.Split(line, "   ")

		l, _ := strconv.Atoi(splitted[0])
		r, _ := strconv.Atoi(splitted[1])

		left = append(left, l)
		right = append(right, r)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	result := float64(0)

	for i := 0; i < len(left); i++ {
		a := left[i]
		b := right[i]

		result += math.Abs(float64(a - b))
	}

	return fmt.Sprintf("%d", int(result))
}

func Day1P2(s string) string {
	left := []int{}
	right := map[int]int{}

	for _, line := range strings.Split(s, "\n") {
		splitted := strings.Split(line, "   ")

		l, _ := strconv.Atoi(splitted[0])
		r, _ := strconv.Atoi(splitted[1])

		left = append(left, l)
		right[r] += 1
	}

	result := 0

	for _, i := range left {
		a := right[i]

		result += a * i
	}

	return fmt.Sprintf("%d", int(result))
}

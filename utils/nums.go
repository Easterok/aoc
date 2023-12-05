package utils

import (
	"strconv"
	"strings"
)

func Clamp(index int, _min int, _max int) int {
	return min(_max, max(index, _min))
}

func LineToNums(s string) []int {
	s = strings.TrimSpace(s)

	result := []int{}

	splitted := strings.Split(s, " ")

	for _, i := range splitted {
		trimmed := strings.TrimSpace(i)

		if len(trimmed) > 0 {
			num, err := strconv.Atoi(trimmed)

			if err != nil {
				panic(err)
			}

			result = append(result, num)
		}
	}

	return result
}

func LineToUint64(s string) []uint64 {
	s = strings.TrimSpace(s)

	result := []uint64{}

	splitted := strings.Split(s, " ")

	for _, i := range splitted {
		trimmed := strings.TrimSpace(i)

		if len(trimmed) > 0 {
			num, err := strconv.ParseUint(trimmed, 10, 64)

			if err != nil {
				panic(err)
			}

			result = append(result, num)
		}
	}

	return result
}

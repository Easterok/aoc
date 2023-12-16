package day3

import (
	"strings"
)

func isUppercase(s rune) bool {
	return strings.ToUpper(string(s)) == string(s)
}

func appearance(s string) rune {
	mid := len(s) / 2
	left := s[:mid]
	right := s[mid:]

	i := 0

	for i < len(left) {
		if strings.Contains(right, string(left[i])) {
			return rune(left[i])
		}

		i += 1
	}

	panic("error")
}

func score(s string) int {
	ch := appearance(s)

	if isUppercase(ch) {
		return int(27 + ch - rune('A'))
	}

	return int(ch - rune('a') + 1)
}

func Part1(content string) int {
	lines := strings.Split(content, "\r\n")

	total := 0

	for _, line := range lines {
		total += score(line)
	}

	return total
}

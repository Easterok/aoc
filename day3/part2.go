package day3

import (
	"sort"
	"strings"
)

func appearance2(s string) rune {
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

type SortByLength []string

func (a SortByLength) Len() int           { return len(a) }
func (a SortByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func commonRune(s []string) rune {
	sort.Sort(SortByLength(s))

	a, b, c := s[0], s[1], s[2]
	m := len(a)

	i := 0

	for i < m {
		ch := string(a[i])

		if strings.Contains(a, ch) && strings.Contains(b, ch) && strings.Contains(c, ch) {
			return rune(a[i])
		}

		i += 1
	}

	panic("invalid")
}

func score2(ch rune) int {
	if isUppercase(ch) {
		return int(27 + ch - rune('A'))
	}

	return int(ch - rune('a') + 1)
}

func Part2(content string) int {
	lines := strings.Split(content, "\r\n")

	total := 0

	i := 0

	for i < len(lines) {
		r := commonRune(lines[i : i+3])

		total += score2(r)

		i += 3
	}

	return total
}

package day13

import (
	"strings"
)

func isMirror(top, bottom []string) bool {
	m := min(len(top), len(bottom))

	result := true

	for i := 0; i < m && result; i++ {
		for j := 0; j < len(top[0]) && result; j++ {
			result = result && top[len(top)-1-i] == bottom[i]
		}
	}

	return result
}

func findMirror(group string) int {
	c := strings.Split(group, "\r\n")

	for i := 1; i < len(c); i++ {
		above := c[:i]
		below := c[i:]

		if isMirror(above, below) {
			return i
		}
	}

	return 0
}

func turn90degString(group string) string {
	splitted := strings.Split(group, "\r\n")

	acc := []string{}

	col := len(splitted[0])
	row := len(splitted)

	for c := 0; c < col; c++ {
		a := make([]string, row)

		for r := 0; r < row; r++ {
			a = append(a, string(splitted[r][c]))
		}

		acc = append(acc, strings.Join(a, ""))
	}

	return strings.Join(acc, "\r\n")
}

func Part1(content string) int {
	groups := strings.Split(content, "\r\n\r\n")

	result := 0

	for _, k := range groups {
		result += findMirror(k) * 100
		result += findMirror(turn90degString(k))
	}

	return result
}

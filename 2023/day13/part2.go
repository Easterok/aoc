package day13

import (
	"strings"
)

func revert(arr []string) []string {
	acc := []string{}

	for i := len(arr) - 1; i >= 0; i-- {
		acc = append(acc, arr[i])
	}

	return acc
}

func isMirror2(top, bottom []string) bool {
	reverted := revert(top)

	m := min(len(reverted), len(bottom))

	result := 0

	for i := 0; i < m; i++ {
		for j := 0; j < len(reverted[0]); j++ {
			if bottom[i][j] != reverted[i][j] {
				result += 1
			}
		}
	}

	return result == 1
}

func findMirror2(group string) int {
	c := strings.Split(group, "\r\n")

	for i := 1; i < len(c); i++ {
		above := c[:i]
		below := c[i:]

		if isMirror2(above, below) {
			return i
		}
	}

	return 0
}

func Part2(content string) int {
	groups := strings.Split(content, "\r\n\r\n")

	result := 0

	for _, k := range groups {
		result += findMirror2(k) * 100
		result += findMirror2(turn90degString(k))
	}

	return result
}

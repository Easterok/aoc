package day12

import (
	"aoc/utils"
	"strings"
)

func findVariants(springs string, missed []int) int {
	if springs == "" {
		if len(missed) == 0 {
			return 1
		}

		return 0
	}

	if len(missed) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		}

		return 1
	}

	ch := springs[0]
	result := 0

	if ch == '.' || ch == '?' {
		result += findVariants(springs[1:], missed)
	}

	if ch == '#' || ch == '?' {
		f := missed[0]

		if f <= len(springs) && !strings.Contains(springs[:f], ".") && (f == len(springs) || springs[f] != '#') {
			if f == len(springs) {
				result += findVariants("", missed[1:])
			} else {
				result += findVariants(springs[f+1:], missed[1:])
			}
		}
	}

	return result
}

func Part1(content string) int {
	lines := strings.Split(content, "\r\n")

	result := 0

	for _, line := range lines {
		splitted := strings.Split(line, " ")

		springs := splitted[0]

		missed := utils.LineToNums(strings.ReplaceAll(splitted[1], ",", " "))

		result += findVariants(springs, missed)
	}

	return result
}

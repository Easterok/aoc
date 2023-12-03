package day3

import (
	"aoc/utils"
	"strconv"
	"strings"
	"unicode"
)

func isSymbol(s byte) bool {
	return !unicode.IsDigit(rune(s)) && s != '.'
}

func isAdjacent(num []int, inline string, step int) bool {
	min := 0
	max := len(inline) - 1

	start := utils.Clamp(num[0]-1, min, max)
	end := utils.Clamp(num[len(num)-1]+1, min, max)

	c := isSymbol(inline[start]) || isSymbol(inline[utils.Clamp(start-step, min, max)]) || isSymbol(inline[utils.Clamp(start+step, min, max)]) || isSymbol(inline[utils.Clamp(end+step, min, max)]) || isSymbol(inline[utils.Clamp(end-step, min, max)]) || isSymbol(inline[utils.Clamp(end, min, max)])

	if c {
		return true
	}

	for _, n := range num {
		top := utils.Clamp(n-step, min, max)
		bottom := utils.Clamp(n+step, min, max)

		if isSymbol(inline[top]) || isSymbol(inline[bottom]) {
			return true
		}
	}

	return false
}

func toNum(num []int, inline string) int {
	start := num[0]
	end := num[len(num)-1]

	s := inline[start : end+1]

	res, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return res
}

func Part1(content string) int {
	lines := strings.Split(content, "\n")

	result := 0

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	cols := len(lines[0])

	inline := strings.Join(lines, "")

	num := []int{}

	for index, ch := range inline {
		if unicode.IsDigit(ch) {
			num = append(num, index)
		} else {
			if len(num) > 0 {
				res := isAdjacent(num, inline, cols)

				if res {
					m := toNum(num, inline)

					result += m
				}

				num = []int{}
			}
		}
	}

	return result
}

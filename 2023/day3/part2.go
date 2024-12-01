package day3

import (
	"sort"
	"strings"
	"unicode"
)

func ExtractNum(index []int, inline string) int {
	sort.Ints(index)

	visited := make(map[int]bool)
	res := []int{}

	for _, k := range index {
		_, ok := visited[k]

		if ok {
			continue
		}

		visited[k] = true

		lp := k - 1
		rp := k + 1

		for lp >= 0 && unicode.IsDigit(rune(inline[lp])) {
			visited[lp] = true

			lp = lp - 1
		}

		for rp < len(inline) && unicode.IsDigit(rune(inline[rp])) {
			visited[rp] = true

			rp = rp + 1
		}

		num := toNum([]int{lp + 1, rp - 1}, inline)

		res = append(res, num)
	}

	if len(res) != 2 {
		return 0
	}

	return res[0] * res[1]
}

func Part2(content string) int {
	lines := strings.Split(content, "\n")

	result := 0

	ratio := []int{}

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	cols := len(lines[0])

	inline := strings.Join(lines, "")

	for index, ch := range inline {
		if ch == '*' {
			ratio = append(ratio, index)
		}
	}

	for _, index := range ratio {
		l := index - 1
		c := index
		r := index + 1

		test := []int{}
		toLookUp := []int{}

		if l >= 0 {
			test = append(test, l)

			if unicode.IsDigit(rune(inline[l])) {
				toLookUp = append(toLookUp, l)
			}
		}

		test = append(test, c)

		if r <= len(inline)-1 {
			test = append(test, r)

			if unicode.IsDigit(rune(inline[r])) {
				toLookUp = append(toLookUp, r)
			}
		}

		for _, i := range test {
			top_hit := i-cols > 0 && unicode.IsDigit(rune(inline[i-cols]))
			bottom_hit := i+cols < len(inline)-1 && unicode.IsDigit(rune(inline[i+cols]))

			if top_hit {
				toLookUp = append(toLookUp, i-cols)
			}

			if bottom_hit {
				toLookUp = append(toLookUp, i+cols)
			}
		}

		result += ExtractNum(toLookUp, inline)
	}

	return result
}

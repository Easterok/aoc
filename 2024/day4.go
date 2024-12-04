package aoc_2024

import (
	"fmt"
	"strings"
)

const XMAS_LEN = len("XMAS")

func xmas(s string) int {
	count := 0
	index := 0

	for index <= len(s)-XMAS_LEN {
		c := s[index : index+XMAS_LEN]
		if c == "XMAS" || c == "XSAM" {
			count += 1
		}

		index += 1
	}

	return count
}

func verticalTransform(s []string) []string {
	l := len(strings.Split(s[0], ""))
	res := make([]string, l)

	index := 0

	for index < l {
		acc := make([]string, len(s))
		i := 0

		for i < len(s) {
			acc[i] = string(s[i][index])
			i += 1
		}

		res[index] = strings.Join(acc, "")

		index += 1
	}

	return res
}

func diagTransform(s []string) []string {
	l := len(strings.Split(s[0], ""))
	res := make([]string, l+len(s)-1)

	setIndex := 0

	for row := 0; row < len(s); row++ {
		r := row
		c := 0

		add := ""

		for r >= 0 && c <= l {
			add += string(s[r][c])

			r -= 1
			c += 1
		}

		res[setIndex] = add
		setIndex += 1
	}

	for col := 1; col < l; col++ {
		r := len(s) - 1
		c := col

		add := ""

		for r >= 0 && c < l {
			add += string(s[r][c])

			r -= 1
			c += 1
		}

		res[setIndex] = add
		setIndex += 1
	}

	return res
}

func diagTransform2(s []string) []string {
	l := len(strings.Split(s[0], ""))
	res := make([]string, l+len(s)-1)

	setIndex := 0

	for row := 0; row < len(s); row++ {
		r := row
		c := l - 1

		add := ""

		for r >= 0 && c >= 0 {
			add += string(s[r][c])

			r -= 1
			c -= 1
		}

		res[setIndex] = add
		setIndex += 1
	}

	for col := l - 2; col >= 0; col-- {
		r := len(s) - 1
		c := col

		add := ""

		for r >= 0 && c >= 0 {
			add += string(s[r][c])

			r -= 1
			c -= 1
		}

		res[setIndex] = add
		setIndex += 1
	}

	return res
}

func Day4P1(content string) string {
	count := 0
	spl := strings.Split(content, "\n")
	vert := verticalTransform(spl)
	diag := diagTransform(spl)
	diag2 := diagTransform2(spl)

	for _, line := range spl {
		count += xmas(line)
	}

	for _, line := range vert {
		count += xmas(line)
	}

	for _, line := range diag {
		count += xmas(line)
	}

	for _, line := range diag2 {
		count += xmas(line)
	}

	return fmt.Sprintf("%d", count)
}

func xmas2(s []string) bool {
	a := string(s[0][0]) + string(s[1][1]) + string(s[2][2])
	b := string(s[2][0]) + string(s[1][1]) + string(s[0][2])

	return (a == "MAS" || a == "SAM") && (b == "MAS" || b == "SAM")
}

func Day4P2(content string) string {
	count := 0
	spl := strings.Split(content, "\n")
	l := strings.Split(spl[0], "")

	for row := 0; row < len(spl)-2; row++ {
		for col := 0; col < len(l)-2; col++ {
			box := make([]string, 3)

			for r := 0; r < 3; r++ {
				a := ""

				for c := 0; c < 3; c++ {
					a += string(spl[row+r][col+c])
				}

				box[r] = a
			}

			if xmas2(box) {
				count += 1
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

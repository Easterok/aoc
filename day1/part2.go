package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var wordToNumer = map[string]byte{
	"one":   byte(49),
	"two":   byte(50),
	"three": byte(51),
	"four":  byte(52),
	"five":  byte(53),
	"six":   byte(54),
	"seven": byte(55),
	"eight": byte(56),
	"nine":  byte(57),
}

func clamp(index int, _min int, _max int) int {
	return min(_max, max(index, _min))
}

func IsDigit(index int, line string, reverse bool) (bool, byte) {
	ch := line[index]

	if unicode.IsDigit(rune(ch)) {
		return true, ch
	}

	a, b, c := func() (int, int, int) {
		if reverse {
			return -3, -4, -5
		}

		return 3, 4, 5
	}()

	index_3 := clamp(index+a, 0, len(line)-1)
	index_4 := clamp(index+b, 0, len(line)-1)
	index_5 := clamp(index+c, 0, len(line)-1)

	a1, a2, b1, b2, c1, c2 := func() (int, int, int, int, int, int) {
		if reverse {
			return index_3, index, index_4, index, index_5, index
		}

		return index, index_3, index, index_4, index, index_5
	}()

	v3, ex3 := wordToNumer[line[a1:a2]]
	v4, ex4 := wordToNumer[line[b1:b2]]
	v5, ex5 := wordToNumer[line[c1:c2]]

	if ex3 {
		return true, v3
	}

	if ex4 {
		return true, v4
	}

	if ex5 {
		return true, v5
	}

	return false, byte(0)
}

func AocDay1Part2(content string) int {
	lines := strings.Split(content, "\n")

	result := 0

	for _, line := range lines {
		lp, rp := 0, len(line)-1
		var first, second byte

		for lp < len(line)-1 && first == 0 {
			check, el := IsDigit(lp, line, false)

			if check {
				first = el
			}

			lp += 1
		}

		for rp > -1 && second == 0 {
			check, el := IsDigit(rp, line, true)

			if check {
				second = el
			}

			rp -= 1
		}

		conc := fmt.Sprintf("%c%c", first, second)

		r, err := strconv.Atoi(conc)

		if err != nil {
			panic(r)
		}

		result += r
	}

	return result
}

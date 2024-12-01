package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func AocDay1Part1(content string) {
	lines := strings.Split(content, "\n")

	i := 0

	result := 0

	for i < len(lines) {
		line := lines[i]

		lp, rp := 0, len(line)-1

		var first, second byte

		lineResult := 0

		for lp < len(line) && first == 0 {
			char := line[lp]

			if unicode.IsDigit(rune(char)) {
				first = char
			}

			lp++
		}

		for rp > -1 && second == 0 {
			char := line[rp]

			if unicode.IsDigit(rune(char)) {
				second = char
			}

			rp = rp - 1
		}

		concatted := fmt.Sprintf("%c%c", first, second)

		lineResult, err := strconv.Atoi(concatted)

		if err != nil {
			panic(err)
		}

		result = result + lineResult

		i++
	}

	fmt.Println(result)
}

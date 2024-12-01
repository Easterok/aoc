package day15

import (
	"strings"
)

func Part1(content string) int {
	seq := strings.Split(content, ",")

	total := 0

	for _, k := range seq {
		cv := rune(0)

		for _, ch := range k {
			cv += ch

			cv = cv * 17
			cv = cv % 256
		}

		total += int(cv)
	}

	return total
}

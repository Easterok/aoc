package day2

import (
	"strings"
)

var opMap = map[string]map[string]string{
	"A": {
		"win":  "Y",
		"lose": "Z",
	},
	"B": {
		"win":  "Z",
		"lose": "A",
	},
	"C": {
		"win":  "A",
		"lose": "B",
	},
}

func score2(line string) int {
	s := strings.Split(line, " ")

	opponent := s[0]
	me := s[1]

	if me == "Y" {
		return 3 + scoreMap[opponent]
	}

	if me == "Z" {
		return 6 + scoreMap[opMap[opponent]["win"]]
	}

	return scoreMap[opMap[opponent]["lose"]]
}

func Part2(content string) int {
	rounds := strings.Split(content, "\r\n")

	total := 0

	for _, k := range rounds {
		total += score2(k)
	}

	return total
}

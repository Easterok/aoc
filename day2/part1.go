package day2

import (
	"strings"
)

var scoreMap = map[string]int{
	// rock
	"X": 1,
	"A": 1,

	// paper
	"Y": 2,
	"B": 2,

	// scissors
	"Z": 3,
	"C": 3,
}

var winMap = map[string]map[string]bool{
	"X": {
		"B": false,
		"C": true,
	},
	"Y": {
		"A": true,
		"C": false,
	},
	"Z": {
		"A": false,
		"B": true,
	},
}

func score(line string) int {
	s := strings.Split(line, " ")

	opponent := s[0]
	me := s[1]

	if scoreMap[me] == scoreMap[opponent] {
		return scoreMap[me] + 3
	}

	win := winMap[me][opponent]

	if win {
		return scoreMap[me] + 6
	}

	return scoreMap[me]
}

func Part1(content string) int {
	rounds := strings.Split(content, "\r\n")

	total := 0

	for _, k := range rounds {
		total += score(k)
	}

	return total
}

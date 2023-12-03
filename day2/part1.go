package day2

import (
	"strconv"
	"strings"
)

var colorMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func GetRedGreenBlue(game string, gameId int) bool {
	toses := strings.Split(game, ", ")

	reached := false

	for _, toss := range toses {
		res := strings.Split(strings.TrimSpace(toss), " ")

		value, err := strconv.Atoi(res[0])

		if err != nil {
			panic(err)
		}

		color := res[1]
		compareTo := colorMap[color]

		if value > compareTo {
			reached = true
		}
	}

	return reached
}

func Part1(content string) uint32 {
	lines := strings.Split(content, "\n")

	var result uint32

	for _, line := range lines {
		splitted := strings.Split(line, ": ")

		game := splitted[0]
		other := splitted[1]

		gameId, err := strconv.Atoi(strings.Replace(game, "Game ", "", -1))

		if err != nil {
			panic(err)
		}

		probably := true

		games := strings.Split(other, "; ")

		for _, game := range games {
			reached := GetRedGreenBlue(game, gameId)

			if reached {
				probably = false
			}
		}

		if probably {
			// fmt.Println(line)

			result = result + uint32(gameId)
		}
	}

	return result
}

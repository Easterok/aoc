package day2

import (
	"strconv"
	"strings"
)

type RGB struct {
	R int
	B int
	G int
}

func GetRedGreenBlueCount(game string, gameId int) (uint16, uint16, uint16) {
	toses := strings.Split(game, ", ")

	acc := [3]uint16{0, 0, 0}

	for _, toss := range toses {
		res := strings.Split(strings.TrimSpace(toss), " ")

		value, err := strconv.Atoi(res[0])

		if err != nil {
			panic(err)
		}

		color := res[1]

		if color == "red" {
			acc[0] = max(acc[0], uint16(value))
		} else if color == "green" {
			acc[1] = max(acc[1], uint16(value))
		} else {
			acc[2] = max(acc[2], uint16(value))
		}
	}

	return acc[0], acc[1], acc[2]
}

func Part2(content string) uint64 {
	lines := strings.Split(content, "\n")

	var result uint64

	for _, line := range lines {
		splitted := strings.Split(line, ": ")

		game := splitted[0]
		other := splitted[1]

		gameId, err := strconv.Atoi(strings.Replace(game, "Game ", "", -1))

		if err != nil {
			panic(err)
		}

		games := strings.Split(other, "; ")

		m_r, m_g, m_b := uint16(0), uint16(0), uint16(0)

		for _, game := range games {

			r, g, b := GetRedGreenBlueCount(game, gameId)

			m_r = max(r, m_r)
			m_g = max(g, m_g)
			m_b = max(b, m_b)
		}

		result += uint64(m_r * m_g * m_b)
	}

	return result
}

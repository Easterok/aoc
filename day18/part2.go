package day18

import (
	"aoc/utils"
	"math"
	"strings"
)

func parseHex(hex string) (string, uint64) {
	prep := strings.ReplaceAll(strings.ReplaceAll(hex, ")", ""), "(", "")

	c := prep[1 : len(prep)-1]
	direction := prep[len(prep)-1]

	dir := ""

	if direction == '0' {
		dir = "R"
	} else if direction == '1' {
		dir = "D"
	} else if direction == '2' {
		dir = "L"
	} else if direction == '3' {
		dir = "U"
	} else {
		panic("Invalid direction: " + string(direction))
	}

	return dir, uint64(utils.HexToInt64(c))
}

func Part2(content string) uint64 {
	lines := strings.Split(content, "\r\n")

	acc := [][]uint64{{0, 0}}
	boundary := uint64(0)

	for _, line := range lines {
		s := strings.Split(line, " ")

		dir, count := parseHex(s[2])

		d := directionMap[dir]
		dr, dc := uint64(d[0]), uint64(d[1])

		last := acc[len(acc)-1]
		r, c := uint64(last[0]), uint64(last[1])

		boundary += count

		acc = append(acc, []uint64{r + dr*count, c + dc*count})
	}

	area := uint64(0)

	for i := 0; i < len(acc); i++ {
		left := i - 1

		if i == 0 {
			left = len(acc) - 1
		}

		// https://en.wikipedia.org/wiki/Shoelace_formula
		area += acc[i][0] * (acc[left][1] - acc[(i+1)%len(acc)][1])
	}

	area = uint64(math.Abs(float64(area)) / 2)

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	i := area - boundary/2 + 1

	return i + boundary
}

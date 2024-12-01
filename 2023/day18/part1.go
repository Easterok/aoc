package day18

import (
	"aoc/utils"
	"math"
	"strings"
)

var directionMap = map[string][]int{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

func Part1(content string) int {
	lines := strings.Split(content, "\r\n")

	acc := [][]int{{0, 0}}
	boundary := 0

	for _, line := range lines {
		s := strings.Split(line, " ")

		dir, count, _ := s[0], utils.LineToNums(s[1])[0], s[2]

		d := directionMap[dir]
		dr, dc := d[0], d[1]

		last := acc[len(acc)-1]
		r, c := last[0], last[1]

		boundary += count

		acc = append(acc, []int{r + dr*count, c + dc*count})
	}

	area := 0

	for i := 0; i < len(acc); i++ {
		left := i - 1

		if i == 0 {
			left = len(acc) - 1
		}

		// https://en.wikipedia.org/wiki/Shoelace_formula
		area += acc[i][0] * (acc[left][1] - acc[(i+1)%len(acc)][1])
	}

	area = int(math.Abs(float64(area)) / 2)

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	i := area - boundary/2 + 1

	return i + boundary
}

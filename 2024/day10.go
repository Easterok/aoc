package aoc_2024

import (
	"fmt"
	"strings"
)

var d = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type GridPosition struct {
	row int
	col int
}

func day10eval(row, col int, value int, acc *[][]string, visited *map[GridPosition]bool) int {
	if row < 0 || row >= len(*acc) || col < 0 || col >= len((*acc)[0]) {
		return 0
	}

	curr := toInts((*acc)[row][col])[0]

	pos := GridPosition{
		row: row,
		col: col,
	}

	if (*visited)[pos] || curr != value {
		return 0
	}

	if curr == 9 {
		(*visited)[pos] = true

		return 1
	}

	next := value + 1
	res := 0

	(*visited)[pos] = true

	for _, dd := range d {
		r := row + dd[0]
		c := col + dd[1]

		res += day10eval(r, c, next, acc, visited)
	}

	return res
}

func Day10P1(content string) string {
	count := 0

	lines := strings.Split(content, "\n")

	acc := make([][]string, len(lines))

	for row, line := range lines {
		spl := strings.Split(line, "")
		acc[row] = spl
	}

	for row, line := range acc {
		for col, ch := range line {
			if ch == "0" {
				count += day10eval(row, col, 0, &acc, &map[GridPosition]bool{})
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

func day10eval2(row, col int, value int, acc *[][]string) int {
	if row < 0 || row >= len(*acc) || col < 0 || col >= len((*acc)[0]) {
		return 0
	}

	curr := toInts((*acc)[row][col])[0]

	if curr != value {
		return 0
	}

	if curr == 9 {
		return 1
	}

	next := value + 1
	res := 0

	for _, dd := range d {
		r := row + dd[0]
		c := col + dd[1]

		res += day10eval2(r, c, next, acc)
	}

	return res
}

func Day10P2(content string) string {
	count := 0

	lines := strings.Split(content, "\n")

	acc := make([][]string, len(lines))

	for row, line := range lines {
		spl := strings.Split(line, "")
		acc[row] = spl
	}

	for row, line := range acc {
		for col, ch := range line {
			if ch == "0" {
				count += day10eval2(row, col, 0, &acc)
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

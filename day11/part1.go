package day11

import (
	"strings"
)

func expandedRowsAndCols(galaxy []string) Expanded {
	rows := map[int]bool{}

	for r, row := range galaxy {
		shouldExpandRow := true

		for _, col := range row {
			shouldExpandRow = shouldExpandRow && col == '.'
		}

		if shouldExpandRow {
			rows[r] = true
		}
	}

	cols := map[int]bool{}

	for c := range galaxy[0] {
		shouldExpandCol := true

		for r := range galaxy {
			shouldExpandCol = shouldExpandCol && galaxy[r][c] == '.'
		}

		if shouldExpandCol {
			cols[c] = true
		}
	}

	return Expanded{rows: rows, cols: cols}
}

type Node struct {
	row int
	col int
}

func makeMap(entry []string) []Node {
	res := []Node{}

	for r, row := range entry {
		for c, item := range row {
			if item == '#' {
				res = append(res, Node{col: c, row: r})
			}
		}
	}

	return res
}

func Part1(content string) int {
	lines := strings.Split(content, "\r\n")

	expanded := expandedRowsAndCols(lines)
	entry := makeMap(lines)

	result := 0
	scale := 2

	for index, start := range entry {
		for _, dist := range entry[index+1:] {
			r := min(start.row, dist.row)
			r_max := max(start.row, dist.row)

			c := min(start.col, dist.col)
			c_max := max(start.col, dist.col)

			for r < r_max {
				if expanded.rows[r] {
					result += scale
				} else {
					result += 1
				}

				r += 1
			}

			for c < c_max {
				if expanded.cols[c] {
					result += scale
				} else {
					result += 1
				}

				c += 1
			}
		}
	}

	return result
}

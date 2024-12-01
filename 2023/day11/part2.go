package day11

import (
	"strings"
)

type Expanded struct {
	rows map[int]bool
	cols map[int]bool
}

func Part2(content string) uint64 {
	lines := strings.Split(content, "\r\n")
	expanded := expandedRowsAndCols(lines)
	entry := makeMap(lines)

	scale := uint64(1000000)

	result := uint64(0)

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

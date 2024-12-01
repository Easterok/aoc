package day16

import (
	"strings"
)

func walk2(row, col int, lines []string, direction string, v *[][]string) {
	ROWS := len(lines)
	COLS := len(lines[0])

	if row < 0 || row > ROWS-1 || col < 0 || col > COLS-1 {
		return
	}

	c := lines[row][col]
	visit := (*v)[row][col]

	if strings.Contains(visit, direction) {
		return
	}

	(*v)[row][col] = visit + direction

	if c == '-' {
		if direction == "l" {
			walk2(row, col-1, lines, direction, v)
		} else if direction == "r" {
			walk2(row, col+1, lines, direction, v)
		} else if direction == "t" || direction == "b" {
			walk2(row, col-1, lines, "l", v)
			walk2(row, col+1, lines, "r", v)
		}
	} else if c == '|' {
		if direction == "l" || direction == "r" {
			walk2(row-1, col, lines, "t", v)
			walk2(row+1, col, lines, "b", v)
		} else if direction == "t" {
			walk2(row-1, col, lines, direction, v)
		} else if direction == "b" {
			walk2(row+1, col, lines, direction, v)
		}
	} else if c == '/' {
		if direction == "l" {
			walk2(row+1, col, lines, "b", v)
		} else if direction == "r" {
			walk2(row-1, col, lines, "t", v)
		} else if direction == "t" {
			walk2(row, col+1, lines, "r", v)
		} else if direction == "b" {
			walk2(row, col-1, lines, "l", v)
		}
	} else if c == '\\' {
		if direction == "l" {
			walk2(row-1, col, lines, "t", v)
		} else if direction == "r" {
			walk2(row+1, col, lines, "b", v)
		} else if direction == "t" {
			walk2(row, col-1, lines, "l", v)
		} else if direction == "b" {
			walk2(row, col+1, lines, "r", v)
		}
	} else {
		if direction == "r" {
			walk2(row, col+1, lines, direction, v)
		} else if direction == "l" {
			walk2(row, col-1, lines, direction, v)
		} else if direction == "b" {
			walk2(row+1, col, lines, direction, v)
		} else if direction == "t" {
			walk2(row-1, col, lines, direction, v)
		}
	}
}

func findMax(lines []string, r int, c int, d string) int {
	var visited = [][]string{}

	for _, l := range lines {
		acc := []string{}

		for range l {
			acc = append(acc, "")
		}

		visited = append(visited, acc)
	}

	walk2(r, c, lines, d, &visited)

	total := 0

	for i := range visited {
		for k := range visited[i] {
			if visited[i][k] != "" {
				total += 1
			}
		}
	}

	return total
}

func Part2(content string) int {
	lines := strings.Split(content, "\r\n")

	total := 0

	for r := range lines {
		total = max(total, findMax(lines, r, 0, "r"))
		total = max(total, findMax(lines, r, len(lines[0])-1, "l"))
	}

	for c := range lines[0] {
		total = max(total, findMax(lines, 0, c, "b"))
		total = max(total, findMax(lines, len(lines)-1, c, "t"))
	}

	return total
}

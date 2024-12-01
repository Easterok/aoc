package day16

import (
	"strings"
)

var visited = [][]string{}

func walk(row, col int, lines []string, direction string) {
	ROWS := len(lines)
	COLS := len(lines[0])

	if row < 0 || row > ROWS-1 || col < 0 || col > COLS-1 {
		return
	}

	c := lines[row][col]
	visit := visited[row][col]

	if strings.Contains(visit, direction) {
		return
	}

	visited[row][col] = visit + direction

	if c == '-' {
		if direction == "l" {
			walk(row, col-1, lines, direction)
		} else if direction == "r" {
			walk(row, col+1, lines, direction)
		} else if direction == "t" || direction == "b" {
			walk(row, col-1, lines, "l")
			walk(row, col+1, lines, "r")
		}
	} else if c == '|' {
		if direction == "l" || direction == "r" {
			walk(row-1, col, lines, "t")
			walk(row+1, col, lines, "b")
		} else if direction == "t" {
			walk(row-1, col, lines, direction)
		} else if direction == "b" {
			walk(row+1, col, lines, direction)
		}
	} else if c == '/' {
		if direction == "l" {
			walk(row+1, col, lines, "b")
		} else if direction == "r" {
			walk(row-1, col, lines, "t")
		} else if direction == "t" {
			walk(row, col+1, lines, "r")
		} else if direction == "b" {
			walk(row, col-1, lines, "l")
		}
	} else if c == '\\' {
		if direction == "l" {
			walk(row-1, col, lines, "t")
		} else if direction == "r" {
			walk(row+1, col, lines, "b")
		} else if direction == "t" {
			walk(row, col-1, lines, "l")
		} else if direction == "b" {
			walk(row, col+1, lines, "r")
		}
	} else {
		if direction == "r" {
			walk(row, col+1, lines, direction)
		} else if direction == "l" {
			walk(row, col-1, lines, direction)
		} else if direction == "b" {
			walk(row+1, col, lines, direction)
		} else if direction == "t" {
			walk(row-1, col, lines, direction)
		}
	}
}

func Part1(content string) int {
	lines := strings.Split(content, "\r\n")

	for _, l := range lines {
		acc := []string{}

		for range l {
			acc = append(acc, "")
		}

		visited = append(visited, acc)
	}

	walk(0, 0, lines, "r")

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

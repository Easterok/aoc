package day11

import (
	"fmt"
	"math"
	"strings"
)

func expand(galaxy [][]string) [][]string {
	rows := map[int]bool{}

	for r, row := range galaxy {
		shouldExpandRow := true

		for _, col := range row {
			shouldExpandRow = shouldExpandRow && col == "."
		}

		if shouldExpandRow {
			rows[r] = true
		}
	}

	cols := map[int]bool{}

	for c := range galaxy[0] {
		shouldExpandCol := true

		for r := range galaxy {
			shouldExpandCol = shouldExpandCol && galaxy[r][c] == "."
		}

		if shouldExpandCol {
			cols[c] = true
		}
	}

	expanded := [][]string{}
	emptyRow := []string{}

	for i := 0; i < len(galaxy)+len(cols); i++ {
		emptyRow = append(emptyRow, ".")
	}

	for r, row := range galaxy {
		res := []string{}

		for c, item := range row {
			res = append(res, item)

			if cols[c] {
				res = append(res, ".")
			}
		}

		expanded = append(expanded, res)

		if rows[r] {
			expanded = append(expanded, emptyRow)
		}
	}

	return expanded
}

type Node struct {
	row int
	col int
}

func makeMap(entry [][]string) []Node {
	res := []Node{}

	for r, row := range entry {
		for c, item := range row {
			if item == "#" {
				res = append(res, Node{col: c, row: r})
			}
		}
	}

	return res
}

// . . # .
// . . . .
// # . . .

func shortestLength(node Node, dist Node) int {
	col_length := int(math.Abs(float64(dist.col - node.col)))
	row_length := int(math.Abs(float64(dist.row - node.row)))

	if col_length == 0 {
		return row_length
	}

	if row_length == 0 {
		return col_length
	}

	col_step := 1

	if dist.col < node.col {
		col_step = -1
	}

	steps, r, c := 0, 0, 0

	t := "col"

	for node.row+r != dist.row || node.col+c != dist.col {
		if t == "col" {
			c += col_step

			if node.row+r != dist.row {
				t = "row"
			}
		} else {
			r += 1

			if node.col+c != dist.col {
				t = "col"
			}
		}

		steps += 1
	}

	return steps
}

func Part1(content string) int {
	galaxy := [][]string{}

	for _, row := range strings.Split(content, "\r\n") {
		galaxy = append(galaxy, strings.Split(row, ""))
	}

	expanded := expand(galaxy)
	entry := makeMap(expanded)

	for _, e := range expanded {
		fmt.Println(e)
	}

	result := 0

	for index, start := range entry {
		for _, dist := range entry[index+1:] {
			result += shortestLength(start, dist)
		}
	}

	return result
}

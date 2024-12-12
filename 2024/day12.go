package aoc_2024

import (
	"fmt"
	"strings"
)

var day12dd = []GridPosition{{row: -1, col: 0}, {row: 1, col: 0}, {row: 0, col: -1}, {row: 0, col: 1}}

func day12group(row, col int, ch string, content [][]string, visited *map[GridPosition]bool) []GridPosition {
	if row < 0 || row >= len(content) || col < 0 || col >= len(content[0]) {
		return []GridPosition{}
	}

	pos := GridPosition{row: row, col: col}

	if (*visited)[pos] {
		return []GridPosition{}
	}

	newch := content[row][col]

	if newch != ch {
		return []GridPosition{}
	}

	(*visited)[pos] = true
	acc := []GridPosition{pos}

	for _, d := range day12dd {
		r := row + d.row
		c := col + d.col

		acc = append(acc, day12group(r, c, ch, content, visited)...)
	}

	return acc
}

func day12eval(group []GridPosition) int {
	m := map[GridPosition]bool{}

	for _, g := range group {
		m[g] = true
	}

	count := 0

	for _, g := range group {
		base := 4

		l := m[GridPosition{row: g.row, col: g.col - 1}]
		r := m[GridPosition{row: g.row, col: g.col + 1}]
		t := m[GridPosition{row: g.row - 1, col: g.col}]
		b := m[GridPosition{row: g.row + 1, col: g.col}]

		if l {
			base -= 1
		}

		if t {
			base -= 1
		}

		if r {
			base -= 1
		}

		if b {
			base -= 1
		}

		count += base
	}

	return len(group) * count
}

func day12eval2(group []GridPosition) int {
	m := map[GridPosition]bool{}

	for _, g := range group {
		m[g] = true
	}

	left := map[GridPosition]bool{}
	right := map[GridPosition]bool{}
	top := map[GridPosition]bool{}
	bottom := map[GridPosition]bool{}

	for _, g := range group {
		if !m[GridPosition{row: g.row - 1, col: g.col}] {
			top[g] = true
		}

		if !m[GridPosition{row: g.row + 1, col: g.col}] {
			bottom[g] = true
		}

		if !m[GridPosition{row: g.row, col: g.col - 1}] {
			left[g] = true
		}

		if !m[GridPosition{row: g.row, col: g.col + 1}] {
			right[g] = true
		}
	}

	count := 0

	for t := range top {
		if left[t] {
			count += 1
		}

		if right[t] {
			count += 1
		}

		if right[GridPosition{row: t.row - 1, col: t.col - 1}] && !left[t] {
			count += 1
		}

		if left[GridPosition{row: t.row - 1, col: t.col + 1}] && !right[t] {
			count += 1
		}
	}

	for t := range bottom {
		if left[t] {
			count += 1
		}

		if right[t] {
			count += 1
		}

		if right[GridPosition{row: t.row + 1, col: t.col - 1}] && !left[t] {
			count += 1
		}

		if left[GridPosition{row: t.row + 1, col: t.col + 1}] && !right[t] {
			count += 1
		}
	}

	return len(group) * count
}

func Day12P1(content string) string {
	lines := strings.Split(content, "\n")
	acc := make([][]string, len(lines))

	for index, line := range lines {
		acc[index] = strings.Split(line, "")
	}

	count := 0

	visited := map[GridPosition]bool{}
	groups := [][]GridPosition{}

	for row, line := range acc {
		for col, ch := range line {
			pos := GridPosition{row: row, col: col}

			if visited[pos] {
				continue
			}

			groups = append(groups, day12group(row, col, ch, acc, &visited))
		}
	}

	for _, group := range groups {
		count += day12eval(group)
	}

	return fmt.Sprintf("%d", count)
}

func Day12P2(content string) string {
	lines := strings.Split(content, "\n")
	acc := make([][]string, len(lines))

	for index, line := range lines {
		acc[index] = strings.Split(line, "")
	}

	count := 0

	visited := map[GridPosition]bool{}
	groups := [][]GridPosition{}

	for row, line := range acc {
		for col, ch := range line {
			pos := GridPosition{row: row, col: col}

			if visited[pos] {
				continue
			}

			groups = append(groups, day12group(row, col, ch, acc, &visited))
		}
	}

	for _, group := range groups {
		count += day12eval2(group)
	}

	return fmt.Sprintf("%d", count)
}

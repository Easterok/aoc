package aoc_2024

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

var day18dirs = []GridPosition{
	{row: 0, col: 1},
	{row: 0, col: -1},
	{row: 1, col: 0},
	{row: -1, col: 0},
}

func day18eval(start, end GridPosition, grid [][]int) int {
	queue := []GridPosition{start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dd := range day18dirs {
			next := GridPosition{row: pos.row + dd.row, col: pos.col + dd.col}

			if next.row < 0 || next.col < 0 || next.row > 70 || next.col > 70 {
				continue
			}

			val := grid[next.row][next.col]

			if val == math.MaxInt {
				grid[next.row][next.col] = grid[pos.row][pos.col] + 1
				queue = append(queue, next)
			}
		}
	}

	return grid[end.row][end.col]
}

func Day18P1(content string) string {
	size := 71
	fallen := 1024

	grid := make([][]int, size)

	for index := range grid {
		grid[index] = slices.Repeat([]int{math.MaxInt}, size)
	}

	for _, a := range strings.Split(content, "\n")[:fallen] {
		spl := toInts(strings.Split(a, ",")...)
		grid[spl[1]][spl[0]] = -1
	}

	start := GridPosition{row: 0, col: 0}
	end := GridPosition{row: size - 1, col: size - 1}

	grid[start.row][start.col] = 0

	return fmt.Sprintf("%d", day18eval(start, end, grid))
}

func Day18P2(content string) string {
	size := 71
	fallen := 1024

	grid := make([][]int, size)

	for index := range grid {
		grid[index] = slices.Repeat([]int{math.MaxInt}, size)
	}

	lines := strings.Split(content, "\n")

	for _, a := range lines[:fallen] {
		spl := toInts(strings.Split(a, ",")...)
		grid[spl[1]][spl[0]] = -1
	}

	start := GridPosition{row: 0, col: 0}
	end := GridPosition{row: size - 1, col: size - 1}

	result := ""
	grid[start.row][start.col] = 0

	for _, br := range lines[fallen:] {
		spl := toInts(strings.Split(br, ",")...)
		grid[spl[1]][spl[0]] = -1

		g := make([][]int, len(grid))

		for index := range g {
			cp := make([]int, len(grid[index]))
			copy(cp, grid[index])
			g[index] = cp
		}

		res := day18eval(start, end, g)

		if res == math.MaxInt {
			result = br
			break
		}
	}

	return result
}

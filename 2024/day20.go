package aoc_2024

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

var gridStepCross = []Day20Node{
	{row: 0, col: 1},
	{row: 0, col: -1},
	{row: 1, col: 0},
	{row: -1, col: 0},
}

type Day20Node struct {
	row  int
	col  int
	cost int
}

func (n *Day20Node) toPos() GridPosition {
	return GridPosition{row: n.row, col: n.col}
}

func day20deq(q *[]Day20Node) Day20Node {
	index := 0
	minValue := (*q)[index].cost
	item := (*q)[index]

	for i := range *q {
		if (*q)[i].cost < minValue {
			index = i
			minValue = (*q)[i].cost
			item = (*q)[i]
		}
	}

	(*q) = append((*q)[:index], (*q)[index+1:]...)

	return item
}

func day20visited(grid [][]int, start Day20Node) map[GridPosition]int {
	visited := map[GridPosition]int{
		start.toPos(): 0,
	}

	rowMax := len(grid)
	colMax := len(grid[0])

	queue := []Day20Node{start}

	for len(queue) > 0 {
		node := day20deq(&queue)
		pos := node.toPos()

		if val, ok := visited[pos]; ok && val < node.cost {
			continue
		}

		for _, dd := range gridStepCross {
			next := GridPosition{row: pos.row + dd.row, col: pos.col + dd.col}

			if next.row < 0 || next.row >= rowMax || next.col < 0 || next.col >= colMax || grid[next.row][next.col] == -1 {
				continue
			}

			if val, ok := visited[next]; !ok || val > node.cost+1 {
				visited[next] = node.cost + 1
				queue = append(queue, Day20Node{row: next.row, col: next.col, cost: node.cost + 1})
			}
		}
	}

	return visited
}

func Day20P1(content string) string {
	lines := strings.Split(content, "\n")
	acc := make([][]int, len(lines))

	var startPos Day20Node

	for row, line := range lines {
		acc[row] = slices.Repeat([]int{math.MaxInt}, len(line))

		for col, ch := range line {
			node := Day20Node{row: row, col: col, cost: 0}

			if ch == '#' {
				acc[row][col] = -1
			}

			if ch == 'S' {
				startPos = node
			}
		}
	}

	target := day20visited(acc, startPos)
	asd := []GridPosition{}

	for k := range target {
		asd = append(asd, k)
	}

	count := 0

	for index, i := range asd {
		for _, j := range asd[index+1:] {
			cheatCost := math.Abs(float64(i.row-j.row)) + math.Abs(float64(i.col-j.col))
			cost := math.Abs(float64(target[i] - target[j]))

			if cheatCost <= 2 && (cost-cheatCost >= 100) {
				count += 1
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

func Day20P2(content string) string {
	lines := strings.Split(content, "\n")
	acc := make([][]int, len(lines))

	var startPos Day20Node

	for row, line := range lines {
		acc[row] = slices.Repeat([]int{math.MaxInt}, len(line))

		for col, ch := range line {
			node := Day20Node{row: row, col: col, cost: 0}

			if ch == '#' {
				acc[row][col] = -1
			}

			if ch == 'S' {
				startPos = node
			}
		}
	}

	target := day20visited(acc, startPos)
	asd := []GridPosition{}

	for k := range target {
		asd = append(asd, k)
	}

	count := 0

	for index, i := range asd {
		for _, j := range asd[index+1:] {
			cheatCost := math.Abs(float64(i.row-j.row)) + math.Abs(float64(i.col-j.col))
			cost := math.Abs(float64(target[i] - target[j]))

			if cheatCost <= 20 && (cost-cheatCost >= 100) {
				count += 1
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

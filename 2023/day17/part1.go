package day17

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func toVisitedKey(node []int) string {
	return fmt.Sprintf("%d_%d_%d_%d_%d", node[1], node[2], node[3], node[4], node[5])
}

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// just implement PriorityQueue lol
func deq(q *[][]int) []int {
	index := 0
	minValue := (*q)[index][0]
	item := (*q)[index]

	for i := range *q {
		if (*q)[i][0] < minValue {
			index = i
			minValue = (*q)[i][0]
			item = (*q)[i]
		}
	}

	(*q) = append((*q)[:index], (*q)[index+1:]...)

	return item
}

func shouldAdd(ndr, ndc, dr, dc int) bool {
	return (ndr != dr || ndc != dc) && (ndr != -dr || ndc != -dc)
}

func dijkstra(node []int, nums *[][]int, visited *map[string]bool) int {
	queue := [][]int{node}

	for len(queue) > 0 {
		first := deq(&queue)

		heatLoss, r, c, dr, dc, n := first[0], first[1], first[2], first[3], first[4], first[5]

		if r == len(*nums)-1 && c == len((*nums)[0])-1 {
			return heatLoss
		}

		key := toVisitedKey(first)

		if (*visited)[key] {
			continue
		}

		(*visited)[key] = true

		if n < 3 && shouldAdd(0, 0, dr, dc) {
			nr := r + dr
			nc := c + dc

			if nr >= 0 && nr < len(*nums) && nc >= 0 && nc < len((*nums)[0]) {
				queue = append(queue, []int{heatLoss + (*nums)[nr][nc], nr, nc, dr, dc, n + 1})
			}
		}

		for _, dir := range directions {
			ndr := dir[0]
			ndc := dir[1]
			nr := r + ndr
			nc := c + ndc

			if shouldAdd(ndr, ndc, dr, dc) && nr >= 0 && nr < len(*nums) && nc >= 0 && nc < len((*nums)[0]) {
				queue = append(queue, []int{heatLoss + (*nums)[nr][nc], nr, nc, ndr, ndc, 1})
			}
		}
	}

	return 0
}

func Part1(content string) int {
	c := strings.Split(content, "\r\n")
	nums := [][]int{}
	visited := map[string]bool{}

	for _, row := range c {
		acc := []int{}

		for _, col := range row {
			acc = append(acc, utils.LineToNums(string(col))...)
		}

		nums = append(nums, acc)
	}

	heatLoss := 0
	row := 0
	col := 0
	dr := 0
	dc := 0
	stepsCount := 0

	result := dijkstra([]int{heatLoss, row, col, dr, dc, stepsCount}, &nums, &visited)

	return result
}

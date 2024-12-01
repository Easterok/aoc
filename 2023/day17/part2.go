package day17

import (
	"aoc/utils"
	"strings"
)

func dijkstra2(node []int, nums *[][]int, visited *map[string]bool) int {
	queue := [][]int{node}

	for len(queue) > 0 {
		first := deq(&queue)

		heatLoss, r, c, dr, dc, n := first[0], first[1], first[2], first[3], first[4], first[5]

		if r == len(*nums)-1 && c == len((*nums)[0])-1 && n >= 4 {
			return heatLoss
		}

		key := toVisitedKey(first)

		if (*visited)[key] {
			continue
		}

		(*visited)[key] = true

		if n < 10 && shouldAdd(0, 0, dr, dc) {
			nr := r + dr
			nc := c + dc

			if nr >= 0 && nr < len(*nums) && nc >= 0 && nc < len((*nums)[0]) {
				queue = append(queue, []int{heatLoss + (*nums)[nr][nc], nr, nc, dr, dc, n + 1})
			}
		}

		if n >= 4 || !shouldAdd(0, 0, dr, dc) {
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
	}

	return 0
}

func Part2(content string) int {
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

	result := dijkstra2([]int{heatLoss, row, col, dr, dc, stepsCount}, &nums, &visited)

	return result
}

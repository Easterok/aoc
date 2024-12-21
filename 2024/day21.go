package aoc_2024

import (
	"fmt"
	"strings"
)

var numpad = map[string]GridPosition{
	"7": {row: 0, col: 0},
	"8": {row: 0, col: 1},
	"9": {row: 0, col: 2},

	"4": {row: 1, col: 0},
	"5": {row: 1, col: 1},
	"6": {row: 1, col: 2},

	"1": {row: 2, col: 0},
	"2": {row: 2, col: 1},
	"3": {row: 2, col: 2},

	"0": {row: 3, col: 1},
	"A": {row: 3, col: 2},
}

var numpadGrid = [][]int{
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
	{-1, 0, 0},
}

var dirGrid = [][]int{
	{-1, 0, 0},
	{0, 0, 0},
}

var dirpad = map[string]GridPosition{
	"^": {row: 0, col: 1},
	"A": {row: 0, col: 2},

	"<": {row: 1, col: 0},
	"v": {row: 1, col: 1},
	">": {row: 1, col: 2},
}

func day21pad(from, to GridPosition, keys bool) string {
	grid := numpadGrid
	if !keys {
		grid = dirGrid
	}

	down := to.row > from.row
	right := to.col > from.col

	vertical := ""

	if down {
		vertical = strings.Repeat("v", to.row-from.row)
	} else {
		vertical = strings.Repeat("^", from.row-to.row)
	}

	horizontal := ""

	if right {
		horizontal = strings.Repeat(">", to.col-from.col)
	} else {
		horizontal = strings.Repeat("<", from.col-to.col)
	}

	aCorner := grid[to.row][from.col]
	bCorner := grid[from.row][to.col]

	if right && aCorner != -1 {
		return vertical + horizontal + "A"
	}

	if bCorner != -1 {
		return horizontal + vertical + "A"
	}

	return vertical + horizontal + "A"
}

func day21walk(code string, keys bool) string {
	pad := numpad
	if !keys {
		pad = dirpad
	}

	start := pad["A"]
	acc := []string{}

	for _, ch := range code {
		to := pad[string(ch)]

		dirs := day21pad(start, to, keys)
		acc = append(acc, dirs)
		start = to
	}

	return strings.Join(acc, "")
}

func Day21P1(content string) string {
	lines := strings.Split(content, "\n")

	count := 0

	for _, line := range lines {
		code := line
		val := toInts(line[:len(line)-1])[0]

		nums := day21walk(code, true)
		pad1 := day21walk(nums, false)
		pad2 := day21walk(pad1, false)

		count += val * len(pad2)
	}

	return fmt.Sprintf("%d", count)
}

type Day21Key struct {
	s GridPosition
	t GridPosition
	d int
}

var day21cache = map[Day21Key]int{}

func day21s(curr, target GridPosition, depth int) int {
	k := Day21Key{
		s: curr,
		t: target,
		d: depth,
	}

	if val, ok := day21cache[k]; ok {
		return val
	}

	p := day21pad(curr, target, false)

	if depth == 24 {
		return len(p)
	}

	l := 0
	s := dirpad["A"]

	for _, ch := range p {
		t := dirpad[string(ch)]
		l += day21s(s, t, depth+1)
		s = t
	}

	day21cache[k] = l

	return l
}

func day21rec(code string) int {
	curr := dirpad["A"]
	l := 0

	for _, ch := range code {
		t := dirpad[string(ch)]
		l += day21s(curr, t, 0)
		curr = t
	}

	return l
}

func Day21P2(content string) string {
	lines := strings.Split(content, "\n")

	count := 0

	for _, line := range lines {
		code := line
		val := toInts(line[:len(line)-1])[0]

		nums := day21walk(code, true)

		count += val * day21rec(nums)
	}

	return fmt.Sprintf("%d", count)
}

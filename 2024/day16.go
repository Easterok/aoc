package aoc_2024

import (
	"fmt"
	"math"
	"strings"
)

var EAST = GridPosition{row: 0, col: 1}
var WEST = GridPosition{row: 0, col: -1}
var SOUTH = GridPosition{row: 1, col: 0}
var NORTH = GridPosition{row: -1, col: 0}

var day16dirs = map[GridPosition][]GridPosition{
	EAST: []GridPosition{
		NORTH,
		SOUTH,
	},
	WEST: []GridPosition{
		NORTH,
		SOUTH,
	},
	SOUTH: []GridPosition{
		EAST,
		WEST,
	},
	NORTH: []GridPosition{
		EAST,
		WEST,
	},
}

var day16dirs2 = map[GridPosition][]GridPosition{
	EAST: []GridPosition{
		NORTH,
		EAST,
		SOUTH,
	},
	WEST: []GridPosition{
		NORTH,
		WEST,
		SOUTH,
	},
	SOUTH: []GridPosition{
		EAST,
		SOUTH,
		WEST,
	},
	NORTH: []GridPosition{
		EAST,
		NORTH,
		WEST,
	},
}

func day16fill(pos, end, dir GridPosition, total int, mp *map[GridPosition]int) {
	if (*mp)[pos] == -1 {
		return
	}

	(*mp)[pos] = min(total, (*mp)[pos])

	if (*mp)[pos] < total {
		return
	}

	if pos == end {
		return
	}

	nextDirs := day16dirs[dir]

	for _, _dir := range nextDirs {
		day16fill(
			GridPosition{row: pos.row + _dir.row, col: pos.col + _dir.col},
			end,
			_dir,
			total+1001,
			mp,
		)
	}

	next := GridPosition{row: pos.row + dir.row, col: pos.col + dir.col}

	day16fill(next, end, dir, total+1, mp)
}

func Day16P1(content string) string {
	mp := map[GridPosition]int{}
	lines := strings.Split(content, "\n")

	var startPos GridPosition
	var endPos GridPosition

	for row, line := range lines {
		for col, ch := range line {
			pos := GridPosition{row: row, col: col}

			if ch == 'S' {
				startPos = pos
			} else if ch == 'E' {
				endPos = pos
				mp[pos] = math.MaxInt
			} else if ch == '#' {
				mp[pos] = -1
			} else {
				mp[pos] = math.MaxInt
			}
		}
	}

	day16fill(startPos, endPos, EAST, 0, &mp)

	return fmt.Sprintf("%d", mp[endPos])
}

func deq(q *[]Day16Node) Day16Node {
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

type Day16Node struct {
	cost int

	row int
	col int
	dr  int
	dc  int
}

type Day16Lowest struct {
	row int
	col int
	dr  int
	dc  int
}

func (n *Day16Node) Lowest() Day16Lowest {
	return Day16Lowest{
		row: n.row,
		col: n.col,
		dr:  n.dr,
		dc:  n.dc,
	}
}

func (n *Day16Node) Pos() GridPosition {
	return GridPosition{
		row: n.row,
		col: n.col,
	}
}

func Day16P2(content string) string {
	mp := map[GridPosition]int{}
	lines := strings.Split(content, "\n")

	var startPos GridPosition
	var endPos GridPosition

	for row, line := range lines {
		for col, ch := range line {
			pos := GridPosition{row: row, col: col}

			if ch == 'S' {
				startPos = pos
			} else if ch == 'E' {
				endPos = pos
				mp[pos] = math.MaxInt
			} else if ch == '#' {
				mp[pos] = -1
			} else {
				mp[pos] = math.MaxInt
			}
		}
	}

	sNode := Day16Node{cost: 0, row: startPos.row, col: startPos.col, dr: 0, dc: 1}
	queue := []Day16Node{
		sNode,
	}

	lowerCost := map[Day16Lowest]int{
		sNode.Lowest(): 0,
	}
	back := map[Day16Lowest]map[Day16Lowest]bool{}
	best := math.MaxInt
	end := map[Day16Node]bool{}

	for len(queue) > 0 {
		item := deq(&queue)

		il := item.Lowest()
		lower, ok := lowerCost[il]
		if !ok {
			lower = math.MaxInt
		}

		if item.cost > lower {
			continue
		}

		pos := GridPosition{row: item.row, col: item.col}

		if pos == endPos {
			if item.cost > best {
				break
			}
			best = item.cost
			end[item] = true
		}

		acc := []Day16Node{
			{cost: item.cost + 1, row: item.row + item.dr, col: item.col + item.dc, dr: item.dr, dc: item.dc},
			{cost: item.cost + 1000, row: item.row, col: item.col, dr: item.dc, dc: -item.dr},
			{cost: item.cost + 1000, row: item.row, col: item.col, dr: -item.dc, dc: item.dr},
		}

		for _, n := range acc {
			if mp[n.Pos()] == -1 {
				continue
			}

			l := n.Lowest()
			lowest, ok := lowerCost[l]
			if !ok {
				lowest = math.MaxInt
			}

			if n.cost > lowest {
				continue
			}
			if n.cost < lowest {
				back[l] = map[Day16Lowest]bool{}
				lowerCost[l] = n.cost
			}

			back[l][il] = true
			queue = append(queue, n)
		}
	}

	acc := []Day16Lowest{}
	for k := range end {
		acc = append(acc, k.Lowest())
	}
	seen := map[Day16Lowest]bool{}

	for len(acc) > 0 {
		a := acc[0]
		acc = acc[1:]

		b := back[a]

		for last := range b {
			if seen[last] {
				continue
			}
			seen[last] = true
			acc = append(acc, last)
		}
	}

	res := map[GridPosition]bool{}

	for k := range seen {
		res[GridPosition{row: k.row, col: k.col}] = true
	}

	return fmt.Sprintf("%d", len(res)+1)
}

package aoc_2024

import (
	"fmt"
	"strings"
)

func day15TryToShift(pos GridPosition, dd GridPosition, mp *map[GridPosition]int) bool {
	if (*mp)[pos] == 0 {
		return true
	}
	if (*mp)[pos] == -1 {
		return false
	}

	next := GridPosition{row: pos.row + dd.row, col: pos.col + dd.col}
	prev := (*mp)[pos]
	res := day15TryToShift(next, dd, mp)

	if res {
		(*mp)[pos] = 0
		(*mp)[next] = prev
	}

	return res
}

func Day15P1(content string) string {
	s := strings.Split(content, "\n\n")

	grid := strings.Split(s[0], "\n")
	moves := strings.Join(strings.Split(s[1], "\n"), "")

	m := map[GridPosition]int{}

	var robotPos GridPosition

	for row, g := range grid {
		for col, ch := range g {
			pos := GridPosition{row: row, col: col}

			m[pos] = 0

			if ch == '#' {
				m[pos] = -1
			}

			if ch == 'O' {
				m[pos] = 1
			}

			if ch == '@' {
				robotPos = pos
			}
		}
	}

	ddmove := map[rune]GridPosition{
		'^': GridPosition{row: -1, col: 0},
		'v': GridPosition{row: 1, col: 0},
		'<': GridPosition{row: 0, col: -1},
		'>': GridPosition{row: 0, col: 1},
	}

	for _, move := range moves {
		dd := ddmove[move]

		next := GridPosition{row: robotPos.row + dd.row, col: robotPos.col + dd.col}

		s := day15TryToShift(next, dd, &m)

		if s {
			robotPos = next
		}
	}

	count := 0

	for k, v := range m {
		if v == 1 {
			count += 100*k.row + k.col
		}
	}

	return fmt.Sprintf("%d", count)
}

func day15resize(grid []string) []string {
	acc := make([]string, len(grid))

	for index, line := range grid {
		a := ""

		for _, ch := range line {
			if ch == '#' {
				a += "##"
			}

			if ch == 'O' {
				a += "[]"
			}

			if ch == '.' {
				a += ".."
			}

			if ch == '@' {
				a += "@."
			}
		}

		acc[index] = a
	}

	return acc
}

func day15Canbemoved(pos GridPosition, dd GridPosition, mp *map[GridPosition]int) bool {
	if (*mp)[pos] == 0 {
		return true
	}

	if (*mp)[pos] == -1 {
		return false
	}

	if dd.col != 0 {
		return day15Canbemoved(GridPosition{row: pos.row, col: pos.col + dd.col}, dd, mp)
	}

	a := pos
	b := pos

	if (*mp)[a] == -2 {
		b = GridPosition{row: pos.row, col: pos.col + 1}
	} else if (*mp)[a] == 2 {
		a = GridPosition{row: pos.row, col: pos.col - 1}
	}

	aNext := GridPosition{row: a.row + dd.row, col: a.col}
	bNext := GridPosition{row: b.row + dd.row, col: b.col}

	return day15Canbemoved(aNext, dd, mp) && day15Canbemoved(bNext, dd, mp)
}

func day15move(pos GridPosition, dd GridPosition, mp *map[GridPosition]int) {
	if dd.col != 0 {
		if (*mp)[pos] == -2 || (*mp)[pos] == 2 {
			prev := (*mp)[pos]
			next := GridPosition{row: pos.row, col: pos.col + dd.col}

			day15move(next, dd, mp)

			(*mp)[pos] = 0
			(*mp)[next] = prev
		}
	} else {
		if (*mp)[pos] == -2 || (*mp)[pos] == 2 {
			a := pos
			b := pos

			if (*mp)[a] == -2 {
				b = GridPosition{row: pos.row, col: pos.col + 1}
			} else if (*mp)[a] == 2 {
				a = GridPosition{row: pos.row, col: pos.col - 1}
			}

			prevA := (*mp)[a]
			prevB := (*mp)[b]

			aNext := GridPosition{row: a.row + dd.row, col: a.col}
			bNext := GridPosition{row: b.row + dd.row, col: b.col}

			day15move(aNext, dd, mp)
			day15move(bNext, dd, mp)

			(*mp)[a] = 0
			(*mp)[aNext] = prevA

			(*mp)[b] = 0
			(*mp)[bNext] = prevB
		}
	}
}

func Day15P2(content string, wasd bool) string {
	s := strings.Split(content, "\n\n")

	grid := day15resize(strings.Split(s[0], "\n"))
	moves := strings.Join(strings.Split(s[1], "\n"), "")

	m := map[GridPosition]int{}

	var robotPos GridPosition

	for row, g := range grid {
		for col, ch := range g {
			pos := GridPosition{row: row, col: col}

			m[pos] = 0

			if ch == '#' {
				m[pos] = -1
			}

			if ch == '[' {
				m[pos] = -2
			}
			if ch == ']' {
				m[pos] = 2
			}

			if ch == '@' {
				robotPos = pos
			}
		}
	}

	ddmove := map[rune]GridPosition{
		'^': GridPosition{row: -1, col: 0},
		'v': GridPosition{row: 1, col: 0},
		'<': GridPosition{row: 0, col: -1},
		'>': GridPosition{row: 0, col: 1},
	}

	if wasd {
		asd := make([][]string, len(grid))

		for ind := range grid {
			asd[ind] = make([]string, len(grid[ind]))
		}

		for k, v := range m {
			if v == 0 {
				asd[k.row][k.col] = "."
			} else if v == -1 {
				asd[k.row][k.col] = "#"
			} else if v == 2 {
				asd[k.row][k.col] = "]"
			} else if v == -2 {
				asd[k.row][k.col] = "["
			}
		}

		asd[robotPos.row][robotPos.col] = "@"

		for _, a := range asd {
			fmt.Println(strings.Join(a, ""))
		}

		for {

			mmove := ""
			fmt.Scanln(&mmove)

			var dd GridPosition

			if mmove == "w" {
				dd = ddmove['^']
			} else if mmove == "a" {
				dd = ddmove['<']
			} else if mmove == "d" {
				dd = ddmove['>']
			} else if mmove == "s" {
				dd = ddmove['v']
			} else {
				continue
			}

			next := GridPosition{row: robotPos.row + dd.row, col: robotPos.col + dd.col}

			res := day15Canbemoved(next, dd, &m)

			if res {
				day15move(next, dd, &m)

				robotPos = next
			}

			for k, v := range m {
				if v == 0 {
					asd[k.row][k.col] = "."
				} else if v == -1 {
					asd[k.row][k.col] = "#"
				} else if v == 2 {
					asd[k.row][k.col] = "]"
				} else if v == -2 {
					asd[k.row][k.col] = "["
				}
			}

			asd[robotPos.row][robotPos.col] = "@"

			for _, a := range asd {
				fmt.Println(strings.Join(a, ""))
			}
		}
	} else {
		for _, move := range moves {
			dd := ddmove[move]

			next := GridPosition{row: robotPos.row + dd.row, col: robotPos.col + dd.col}

			res := day15Canbemoved(next, dd, &m)

			if res {
				day15move(next, dd, &m)

				robotPos = next
			}
		}
	}

	count := 0

	for k, v := range m {
		if v == -2 {
			count += k.row*100 + k.col
		}
	}

	return fmt.Sprintf("%d", count)
}

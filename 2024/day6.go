package aoc_2024

import (
	"fmt"
	"strings"
)

type StartPos struct {
	row int
	col int
	dir Direction
}

type Direction string

const (
	UP    Direction = "up"
	LEFT            = "left"
	RIGHT           = "right"
	DOWN            = "down"
)

var Moves = map[Direction]StartPos{
	UP: StartPos{
		row: -1,
		col: 0,
	},
	RIGHT: StartPos{
		row: 0,
		col: 1,
	},
	DOWN: StartPos{
		row: 1,
		col: 0,
	},
	LEFT: StartPos{
		row: 0,
		col: -1,
	},
}

func exit(pos *StartPos, dir *Direction, rowMax, colMax int) bool {
	switch *dir {
	case UP:
		return pos.row == 0
	case RIGHT:
		return pos.col == colMax-1
	case DOWN:
		return pos.row == rowMax-1
	case LEFT:
		return pos.col == 0
	default:
		return pos.col == 0
	}
}

func nextDir(dir Direction) Direction {
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		return UP
	}
}

func day6move(start *StartPos, content []string) int {
	rowMax := len(content)
	colMax := len(content[0])

	visited := map[StartPos]bool{}
	dir := UP

	for !exit(start, &dir, rowMax, colMax) {
		visited[*start] = true

		m := Moves[dir]
		start.row += m.row
		start.col += m.col

		if content[start.row][start.col] == '#' {
			m := Moves[dir]
			start.row -= m.row
			start.col -= m.col
			dir = nextDir(dir)
			m = Moves[dir]
			start.row += m.row
			start.col += m.col
		}
	}

	return len(visited) + 1
}

func Day6P1(content string) string {
	lines := strings.Split(content, "\n")

	var startPos *StartPos

outer:
	for row, line := range lines {
		for col, ch := range line {
			if ch == '^' {
				startPos = &StartPos{
					row: row,
					col: col,
				}
				break outer
			}
		}
	}

	count := day6move(startPos, lines)

	return fmt.Sprintf("%d", count)
}

func hasInPositions(position *StartPos, visited []StartPos) bool {
	for _, visit := range visited {
		if visit.col == position.col && position.row == visit.row {
			return true
		}
	}

	return false
}

func day6isLoop(pos *StartPos, visited *map[Direction][]StartPos, content []string) bool {
	turn := nextDir(pos.dir)

	poss, ok := (*visited)[turn]
	move := Moves[turn]

	if ok && len(poss) > 0 {
		s := StartPos{row: pos.row, col: pos.col}

		for s.row+move.row >= 0 &&
			s.col+move.col >= 0 &&
			s.col+move.col < len(content[0]) &&
			s.row+move.row < len(content) &&
			content[s.row+move.row][s.col+move.col] != '#' {
			if hasInPositions(&s, poss) {
				return true
			}

			s.row = s.row + move.row
			s.col = s.col + move.col
		}
	}

	return false
}

func day6move2(start *StartPos, content []string) int {
	rowMax := len(content)
	colMax := len(content[0])

	visited := map[StartPos]Direction{}

	s := &StartPos{col: start.col, row: start.row, dir: start.dir}

	for !exit(s, &s.dir, rowMax, colMax) {
		visited[StartPos{col: s.col, row: s.row}] = s.dir

		m := Moves[s.dir]
		s.row += m.row
		s.col += m.col

		if content[s.row][s.col] == '#' {
			m := Moves[s.dir]
			s.row -= m.row
			s.col -= m.col
			s.dir = nextDir(s.dir)
		}
	}

	visited[StartPos{col: s.col, row: s.row}] = s.dir

	c := make([][]string, len(content))

	for index, ctn := range content {
		c[index] = strings.Split(ctn, "")
	}

	count := 0

	for vis := range visited {
		if vis.col == start.col && vis.row == start.row {
			continue
		}

		visited2 := map[StartPos]bool{}
		visited2[StartPos{col: vis.col, row: vis.row, dir: vis.dir}] = true
		include := false

		c[vis.row][vis.col] = "#"

		s = &StartPos{col: start.col, row: start.row, dir: start.dir}

		for !exit(s, &s.dir, rowMax, colMax) {
			if _, ok := visited2[StartPos{col: s.col, row: s.row, dir: s.dir}]; ok {
				include = true

				break
			}

			visited2[StartPos{col: s.col, row: s.row, dir: s.dir}] = true

			m := Moves[s.dir]
			s.row += m.row
			s.col += m.col

			if c[s.row][s.col] == "#" {
				m := Moves[s.dir]
				s.row -= m.row
				s.col -= m.col
				s.dir = nextDir(s.dir)
			}
		}

		if include {
			count += 1
		}

		c[vis.row][vis.col] = "."
	}

	return count
}

func Day6P2(content string) string {
	lines := strings.Split(content, "\n")

	var startPos *StartPos

outer:
	for row, line := range lines {
		for col, ch := range line {
			if ch == '^' {
				startPos = &StartPos{
					row: row,
					col: col,
					dir: UP,
				}
				break outer
			}
		}
	}

	count := day6move2(startPos, lines)

	return fmt.Sprintf("%d", count)
}

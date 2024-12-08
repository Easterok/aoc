package aoc_2024

import (
	"fmt"
	"strings"
	"unicode"
)

type Day8Pos struct {
	row int
	col int
}

func day8inbound(pos Day8Pos, rowBound, colBound int) bool {
	return pos.row >= 0 && pos.row < rowBound && pos.col >= 0 && pos.col < colBound
}

func day8evaluate(acc []Day8Pos, rowBound, colBound int, uniq *map[Day8Pos]bool) int {
	if len(acc) == 0 || len(acc) == 1 {
		return 0
	}

	res := 0
	first := acc[0]

	for _, pos := range acc[1:] {
		rowDiff := first.row - pos.row
		colDiff := first.col - pos.col

		firstAntinode := Day8Pos{
			row: first.row + rowDiff,
			col: first.col + colDiff,
		}

		secondAntinode := Day8Pos{
			row: pos.row - rowDiff,
			col: pos.col - colDiff,
		}

		_, ok := (*uniq)[firstAntinode]

		if !ok && day8inbound(firstAntinode, rowBound, colBound) {
			res += 1
			(*uniq)[firstAntinode] = true
		}

		_, ok = (*uniq)[secondAntinode]

		if !ok && day8inbound(secondAntinode, rowBound, colBound) {
			res += 1
			(*uniq)[secondAntinode] = true
		}
	}

	return res + day8evaluate(acc[1:], rowBound, colBound, uniq)
}

func Day8P1(content string) string {
	c := strings.Split(content, "\n")

	count := 0
	acc := map[rune][]Day8Pos{}
	rowBound := len(c)
	colBound := len(strings.Split(c[0], ""))

	for row, line := range c {
		for col, ch := range line {
			if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
				_, ok := acc[ch]

				if ok {
					acc[ch] = append(acc[ch], Day8Pos{row: row, col: col})
				} else {
					acc[ch] = []Day8Pos{{row: row, col: col}}
				}
			}
		}
	}

	uniq := map[Day8Pos]bool{}

	for _, acc := range acc {
		count += day8evaluate(acc, rowBound, colBound, &uniq)
	}

	return fmt.Sprintf("%d", count)
}

func day8evaluate2(acc []Day8Pos, rowBound, colBound int, uniq *map[Day8Pos]bool) int {
	if len(acc) == 0 || len(acc) == 1 {
		return 0
	}

	res := 0
	first := acc[0]

	for _, pos := range acc[1:] {
		rowDiff := first.row - pos.row
		colDiff := first.col - pos.col

		firstAntinode := &Day8Pos{
			row: first.row + rowDiff,
			col: first.col + colDiff,
		}

		secondAntinode := &Day8Pos{
			row: pos.row - rowDiff,
			col: pos.col - colDiff,
		}

		for day8inbound(*firstAntinode, rowBound, colBound) {
			(*uniq)[*firstAntinode] = true
			res += 1
			firstAntinode.row = firstAntinode.row + rowDiff
			firstAntinode.col = firstAntinode.col + colDiff
		}

		for day8inbound(*secondAntinode, rowBound, colBound) {
			res += 1
			(*uniq)[*secondAntinode] = true
			secondAntinode.row = secondAntinode.row - rowDiff
			secondAntinode.col = secondAntinode.col - colDiff
		}
	}

	return res + day8evaluate2(acc[1:], rowBound, colBound, uniq)
}

func Day8P2(content string) string {
	c := strings.Split(content, "\n")

	count := 0
	acc := map[rune][]Day8Pos{}
	rowBound := len(c)
	colBound := len(strings.Split(c[0], ""))

	for row, line := range c {
		for col, ch := range line {
			if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
				_, ok := acc[ch]

				if ok {
					acc[ch] = append(acc[ch], Day8Pos{row: row, col: col})
				} else {
					acc[ch] = []Day8Pos{{row: row, col: col}}
				}
			}
		}
	}

	uniq := map[Day8Pos]bool{}
	for _, acc := range acc {
		day8evaluate2(acc, rowBound, colBound, &uniq)
	}

	for row, line := range c {
		for col, ch := range line {
			if uniq[Day8Pos{row: row, col: col}] {
				// fmt.Printf("#")
				count += 1
			} else {
				if ch != '.' {
					count += 1
				}
				// fmt.Printf(string(ch))
			}
		}

		// fmt.Printf("\n")
	}

	return fmt.Sprintf("%d", count)
}

package day10

import (
	"strings"
)

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.

type Signature = byte

type Node struct {
	row int
	col int

	signature Signature
}

var visitedNodes = make(map[Node]bool)

func findStartingEndpoint(rows []string) Node {
	for r_index, row := range rows {
		for c_index, col := range row {
			if col == 'S' {
				return Node{row: r_index, col: c_index, signature: 'S'}
			}
		}
	}

	return Node{row: 0, col: 0}
}

func isValidIndex(row int, col int, rows []string) bool {
	rowsLength := len(rows)
	colsLength := len(rows[0])

	return row > -1 && row < rowsLength && col > -1 && col < colsLength
}

func findConnection(node Node, rows []string) []Node {
	result := []Node{}

	check_top := node.signature == 'S' || node.signature == '|' || node.signature == 'L' || node.signature == 'J'

	if check_top && !visitedNodes[node] && isValidIndex(node.row-1, node.col, rows) {
		char := rows[node.row-1][node.col]

		if char == '|' || char == '7' || char == 'F' {
			result = append(result, Node{row: node.row - 1, col: node.col, signature: char})
		}
	}

	check_bottom := node.signature == 'S' || node.signature == '|' || node.signature == 'F' || node.signature == '7'

	if check_bottom && !visitedNodes[node] && isValidIndex(node.row+1, node.col, rows) {
		char := rows[node.row+1][node.col]

		if char == '|' || char == 'L' || char == 'J' {
			result = append(result, Node{row: node.row + 1, col: node.col, signature: char})
		}
	}

	check_left := node.signature == 'S' || node.signature == '-' || node.signature == 'J' || node.signature == '7'

	if check_left && !visitedNodes[node] && isValidIndex(node.row, node.col-1, rows) {
		char := rows[node.row][node.col-1]

		if char == '-' || char == 'L' || char == 'F' {
			result = append(result, Node{row: node.row, col: node.col - 1, signature: char})
		}
	}

	check_right := node.signature == 'S' || node.signature == '-' || node.signature == 'F' || node.signature == 'L'

	if check_right && !visitedNodes[node] && isValidIndex(node.row, node.col+1, rows) {
		char := rows[node.row][node.col+1]

		if char == '-' || char == 'J' || char == '7' {
			result = append(result, Node{row: node.row, col: node.col + 1, signature: char})
		}
	}

	return result
}

func dfs(node Node, rows []string, deep int) int {
	connections := findConnection(node, rows)

	visitedNodes[node] = true

	m := deep

	for _, connection := range connections {
		res := dfs(connection, rows, deep+1)

		m = max(m, res)
	}

	return m
}

func Part1(content string) int {
	rows := strings.Split(content, "\r\n")

	startingNode := findStartingEndpoint(rows)

	connections := findConnection(startingNode, rows)

	result := 0

	for _, connection := range connections {
		res := dfs(connection, rows, 1)

		result = max(res, result)
	}

	return result / 2
}

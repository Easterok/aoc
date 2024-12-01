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

type Signature2 = byte

type Node2 struct {
	row int
	col int

	signature Signature2
}

var visitedNode2s = make(map[Node2]bool)

var maybesSChar = []byte{'|', '-', 'L', 'J', '7', 'F'}

func intersection(maybes []byte, chars []byte) {
	seen := make(map[byte]bool)

	for _, value := range chars {
		seen[value] = true
	}

	res := []byte{}

	for _, value := range maybes {
		if seen[value] {
			res = append(res, value)
		}
	}

	maybesSChar = res
}

func findStartingEndpoint2(rows []string) Node2 {
	for r_index, row := range rows {
		for c_index, col := range row {
			if col == 'S' {
				return Node2{row: r_index, col: c_index, signature: 'S'}
			}
		}
	}

	return Node2{row: 0, col: 0}
}

func findConnection2(node2 Node2, rows []string) []Node2 {
	result := []Node2{}

	check_top := node2.signature == 'S' || node2.signature == '|' || node2.signature == 'L' || node2.signature == 'J'

	if check_top && !visitedNode2s[node2] && isValidIndex(node2.row-1, node2.col, rows) {
		char := rows[node2.row-1][node2.col]

		if char == '|' || char == '7' || char == 'F' {
			result = append(result, Node2{row: node2.row - 1, col: node2.col, signature: char})
		}

		if char == 'S' {
			intersection(maybesSChar, []byte{'|', '7', 'F'})
		}
	}

	check_bottom := node2.signature == 'S' || node2.signature == '|' || node2.signature == 'F' || node2.signature == '7'

	if check_bottom && !visitedNode2s[node2] && isValidIndex(node2.row+1, node2.col, rows) {
		char := rows[node2.row+1][node2.col]

		if char == '|' || char == 'L' || char == 'J' {
			result = append(result, Node2{row: node2.row + 1, col: node2.col, signature: char})
		}

		if char == 'S' {
			intersection(maybesSChar, []byte{'|', 'L', 'J'})
		}
	}

	check_left := node2.signature == 'S' || node2.signature == '-' || node2.signature == 'J' || node2.signature == '7'

	if check_left && !visitedNode2s[node2] && isValidIndex(node2.row, node2.col-1, rows) {
		char := rows[node2.row][node2.col-1]

		if char == '-' || char == 'L' || char == 'F' {
			result = append(result, Node2{row: node2.row, col: node2.col - 1, signature: char})
		}

		if char == 'S' {
			intersection(maybesSChar, []byte{'-', 'L', 'F'})
		}
	}

	check_right := node2.signature == 'S' || node2.signature == '-' || node2.signature == 'F' || node2.signature == 'L'

	if check_right && !visitedNode2s[node2] && isValidIndex(node2.row, node2.col+1, rows) {
		char := rows[node2.row][node2.col+1]

		if char == '-' || char == 'J' || char == '7' {
			result = append(result, Node2{row: node2.row, col: node2.col + 1, signature: char})
		}

		if char == 'S' {
			intersection(maybesSChar, []byte{'-', 'J', '7'})
		}
	}

	return result
}

func dfs2(node2 Node2, rows []string, deep int) int {
	connections := findConnection2(node2, rows)

	visitedNode2s[node2] = true

	m := deep

	for _, connection := range connections {
		res := dfs2(connection, rows, deep+1)

		m = max(m, res)
	}

	return m
}

func unionLength(a [][]bool, b [][]bool) int {
	result := 0

	for index, i := range a {
		for j := range i {
			if i[j] || b[index][j] {
				result += 1
			}
		}
	}

	return result
}

func Part2(content string) int {
	rows := strings.Split(content, "\r\n")

	startingNode2 := findStartingEndpoint2(rows)

	connections := findConnection2(startingNode2, rows)

	result := 0

	for _, connection := range connections {
		res := dfs2(connection, rows, 1)

		result = max(res, result)
	}

	correctSignature := maybesSChar[0]

	m := [][]string{}

	asd := [][]bool{}

	for _, row := range rows {
		withoutS := strings.Replace(row, "S", string(correctSignature), 1)
		splitted := strings.Split(withoutS, "")

		m = append(m, splitted)

		asd = append(asd, make([]bool, len(splitted)))
	}

	for node := range visitedNode2s {
		asd[node.row][node.col] = true
	}

	asd[startingNode2.row][startingNode2.col] = true

	for rowIndex, k := range m {
		for colIndex := range k {
			if !asd[rowIndex][colIndex] {
				m[rowIndex][colIndex] = "."
			}
		}
	}

	outside := [][]bool{}

	for _, row := range m {
		within := false
		up := false

		a := make([]bool, len(row))

		for colIndex, ch := range row {
			if ch == "|" {
				within = !within
			} else if ch == "-" {
			} else if ch == "L" || ch == "F" {
				up = ch == "L"
			} else if ch == "7" || ch == "J" {
				if up {
					if ch != "J" {
						within = !within
					}
				} else {
					if ch != "7" {
						within = !within
					}
				}

				up = false
			}

			if !within {
				a[colIndex] = true
			}
		}

		outside = append(outside, a)
	}

	// for r, k := range outside {
	// 	for c, i := range k {
	// 		if i && !asd[r][c] {
	// 			fmt.Printf("#")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}

	// 	fmt.Println("")
	// }

	return len(asd)*len(asd[0]) - unionLength(outside, asd)
}

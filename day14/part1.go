package day14

import (
	"fmt"
	"strings"
)

func Part1(content string) int {
	rows := strings.Split(content, "\r\n")

	result := 0
	length := len(rows)

	c := []int{}

	for i := 0; i < len(rows[0]); i++ {
		c = append(c, length+1)
	}

	acc := [][]int{c}

	for r := range rows {
		a := []int{}

		for c, ch := range rows[r] {
			prev := acc[r]

			if ch == '#' {
				a = append(a, 0)
			} else if ch == 'O' {
				if prev[c] == 0 {
					a = append(a, length-r)
				} else {
					a = append(a, prev[c]-1)
				}
			} else if ch == '.' {
				if prev[c] == 0 {
					a = append(a, length-r+1)
				} else {
					a = append(a, prev[c])
				}
			}
		}

		acc = append(acc, a)
	}

	for r := range acc[1:] {
		fmt.Println(acc[r+1])
	}

	for r, k := range rows {
		for c := range k {
			if k[c] == 'O' {
				result += acc[r+1][c]
			}
		}
	}

	return result
}

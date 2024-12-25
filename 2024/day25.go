package aoc_2024

import (
	"fmt"
	"strings"
)

func Day25P1(content string) string {
	keys := [][]int{}
	locs := [][]int{}

	for _, el := range strings.Split(content, "\n\n") {
		asd := strings.Split(el, "\n")

		acc := []int{}

		for col := range asd[0] {
			a := -1

			for row := range asd {
				if asd[row][col] == '#' {
					a += 1
				}
			}

			acc = append(acc, a)
		}

		if asd[0][0] == '#' {
			locs = append(locs, acc)
		} else {
			keys = append(keys, acc)
		}
	}

	count := 0

	for _, loc := range locs {
		for _, key := range keys {
			overlap := false

			for i := 0; i < 5; i++ {
				overlap = overlap || (loc[i]+key[i] > 5)
			}

			if !overlap {
				count += 1
			}
		}

	}

	return fmt.Sprintf("%d", count)
}

package aoc_2024

import (
	"fmt"
	"math"
	"strings"
)

func day22preudorandom(num int) int {
	value := num * 64
	value = num ^ value
	value = value % 16777216

	val := int(math.Floor(float64(value / 32)))
	val = value ^ val
	val = val % 16777216

	nval := val * 2048
	nval = val ^ nval
	nval = nval % 16777216

	return nval
}

func Day22P1(content string) string {
	nums := toInts(strings.Split(content, "\n")...)

	count := 0

	for _, num := range nums {
		for i := 0; i < 2_000; i++ {
			next := day22preudorandom(num)
			num = next
		}

		count += num
	}

	return fmt.Sprintf("%d", count)
}

func Day22P2(content string) string {
	nums := toInts(strings.Split(content, "\n")...)

	variants := map[string]bool{}
	acc := []map[string]int{}

	for _, num := range nums {
		m := map[string]int{}
		key := []string{}

		for i := 0; i < 2_000; i++ {
			next := day22preudorandom(num)
			key = append(key, fmt.Sprintf("%d", next%10-num%10))

			if len(key) == 4 {
				kkey := strings.Join(key, "")
				if _, ok := m[kkey]; !ok {
					m[kkey] = next % 10
				}
				variants[kkey] = true
				key = key[1:]
			}

			num = next
		}

		acc = append(acc, m)
	}

	count := 0

	for variant := range variants {
		c := 0

		for _, a := range acc {
			c += a[variant]
		}

		count = max(count, c)
	}

	return fmt.Sprintf("%d", count)
}

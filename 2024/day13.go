package aoc_2024

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Day13Config struct {
	A     GridPosition
	B     GridPosition
	Prize GridPosition
}

func day13parse(g string) Day13Config {
	s := strings.Split(g, "\n")
	axy := strings.Split(strings.Replace(s[0], "Button A: ", "", 1), ", ")
	bxy := strings.Split(strings.Replace(s[1], "Button B: ", "", 1), ", ")
	prizexy := strings.Split(strings.Replace(s[2], "Prize: ", "", 1), ", ")

	return Day13Config{
		A:     GridPosition{row: toInts(axy[0][1:])[0], col: toInts(axy[1][1:])[0]},
		B:     GridPosition{row: toInts(bxy[0][1:])[0], col: toInts(bxy[1][1:])[0]},
		Prize: GridPosition{row: toInts(prizexy[0][2:])[0], col: toInts(prizexy[1][2:])[0]},
	}
}

func Day13P1(content string) string {
	count := 0

	groups := strings.Split(content, "\n\n")

	for _, g := range groups {
		cfg := day13parse(g)

		total := math.MaxInt

		for a := 0; a < 101; a++ {
			for b := 0; b < 101; b++ {
				if cfg.A.row*a+cfg.B.row*b == cfg.Prize.row && cfg.A.col*a+cfg.B.col*b == cfg.Prize.col {
					total = min(total, a*3+b)
				}
			}
		}

		if total != math.MaxInt {
			count += total
		}
	}

	return fmt.Sprintf("%d", count)
}

func Day13P2(content string) string {
	count := 0

	groups := strings.Split(content, "\n\n")

	for _, g := range groups {
		cfg := day13parse(g)

		cfg.Prize.row = 10000000000000 + cfg.Prize.row
		cfg.Prize.col = 10000000000000 + cfg.Prize.col

		// ax * s + bx * t = cx
		// ay * s + by * t = cy

		// ax * ay * s + bx * ay * t = cx * ay
		// ay * ax * s + by * ax * t = cy * ax

		// bx * ay * t - by * ax * t = cx * ay - cy * ax
		// t = (cx * ay - cy * ax) / (bx * ay - by * ax)
		// s = (cx - bx * t) / ax

		t1 := cfg.Prize.row*cfg.A.col - cfg.Prize.col*cfg.A.row
		t2 := cfg.B.row*cfg.A.col - cfg.B.col*cfg.A.row

		if t2 == 0 {
			fmt.Println("unluck. division by 0")
			os.Exit(1)
		}
		if t1%t2 != 0 {
			continue
		}

		t := t1 / t2

		s1 := cfg.Prize.row - cfg.B.row*t
		s2 := cfg.A.row
		if s1%s2 != 0 {
			continue
		}

		s := s1 / s2
		count += (s*3 + t)
	}

	return fmt.Sprintf("%d", count)
}

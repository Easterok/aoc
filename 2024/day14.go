package aoc_2024

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type Day14Robot struct {
	Pos GridPosition
	Vel GridPosition
}

func day14parse(lines []string) []Day14Robot {
	acc := []Day14Robot{}

	for _, line := range lines {
		s := strings.Split(line, " ")
		p := toInts(strings.Split(s[0][2:], ",")...)

		pos := GridPosition{
			row: p[1],
			col: p[0],
		}

		v := toInts(strings.Split(s[1][2:], ",")...)
		vel := GridPosition{
			row: v[1],
			col: v[0],
		}

		acc = append(acc, Day14Robot{
			Pos: pos,
			Vel: vel,
		})
	}

	return acc
}

func Day14P1(content string) string {
	robots := day14parse(strings.Split(content, "\n"))

	ticks := 100

	rowsCount := 103
	colsCount := 101

	middleR := (rowsCount - 1) / 2
	middleC := (colsCount - 1) / 2

	quadA := 0
	quadB := 0
	quadC := 0
	quadD := 0

	for _, robot := range robots {
		p := robot.Pos.row + robot.Vel.row*ticks
		row := (rowsCount - (rowsCount-p)%rowsCount) % rowsCount
		c := robot.Pos.col + robot.Vel.col*ticks
		col := (colsCount - (colsCount-c)%colsCount) % colsCount

		if middleR == row || middleC == col {
			continue
		}

		if row >= 0 && row < middleR {
			if col >= 0 && col < middleC {
				quadA += 1
			} else {
				quadB += 1
			}
		} else if row > middleR && row < rowsCount {
			if col >= 0 && col < middleC {
				quadC += 1
			} else {
				quadD += 1
			}
		}
	}

	count := quadA * quadB * quadC * quadD

	return fmt.Sprintf("%d", count)
}

func Day14P2(content string) string {
	robots := day14parse(strings.Split(content, "\n"))

	rowsCount := 103
	colsCount := 101

	for ticks := 1; ticks < 10_000; ticks++ {
		img := image.NewRGBA(image.Rect(0, 0, colsCount, rowsCount))
		file, _ := os.Create(fmt.Sprintf("day14/%d.png", ticks))

		defer file.Close()

		for _, robot := range robots {
			p := robot.Pos.row + robot.Vel.row*ticks
			row := (rowsCount - (rowsCount-p)%rowsCount) % rowsCount
			c := robot.Pos.col + robot.Vel.col*ticks
			col := (colsCount - (colsCount-c)%colsCount) % colsCount

			img.Set(row, col, color.RGBA{0, 0, 0, 255})
		}

		png.Encode(file, img)
	}

	return fmt.Sprintf("%d", 0)
}

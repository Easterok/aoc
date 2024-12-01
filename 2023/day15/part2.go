package day15

import (
	"aoc/utils"
	"strings"
)

type Node struct {
	label string
	value int
}

func Part2(content string) int {
	seq := strings.Split(content, ",")
	boxes := make([][]Node, 256)

	for _, s := range seq {
		if strings.Contains(s, "-") {
			// remove
			label := strings.Split(s, "-")[0]
			boxIndex := Part1(label)

			if boxes[boxIndex] != nil {
				for index, a := range boxes[boxIndex] {
					if a.label == label {
						boxes[boxIndex] = append(boxes[boxIndex][:index], boxes[boxIndex][index+1:]...)

						continue
					}
				}
			}
		} else if strings.Contains(s, "=") {
			// add
			splitted := strings.Split(s, "=")
			label := splitted[0]
			value := utils.LineToNums(splitted[1])[0]

			boxIndex := Part1(label)

			if boxes[boxIndex] == nil {
				boxes[boxIndex] = []Node{
					{label: label, value: value},
				}
			} else {
				hasOld := false

				for i, k := range boxes[boxIndex] {
					if k.label == label {
						boxes[boxIndex][i].value = value

						hasOld = true
					}
				}

				if !hasOld {
					boxes[boxIndex] = append(boxes[boxIndex], Node{label: label, value: value})
				}
			}
		}
	}

	total := 0

	for i, box := range boxes {
		for j, node := range box {
			t := (i + 1) * (j + 1) * node.value

			total += t
		}
	}

	return total
}

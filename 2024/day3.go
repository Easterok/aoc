package aoc_2024

import (
	"strconv"
	"unicode"
)

func valid(a, b string) bool {
	alen := len(a)
	blen := len(b)

	return alen > 0 && blen > 0 && alen < 4 && blen < 4
}

func Day3P1(content string) int {
	result := 0

	index := 0

	for index+4 < len(content) {
		s := content[index : index+4]

		if s == "mul(" {
			a := ""
			b := ""
			toA := true
			safe := false
			lastIndex := index + 4

			for i, r := range content[index+4:] {
				str := string(r)
				lastIndex = index + 4 + i
				if str == "," {
					toA = false
				} else if unicode.IsDigit(r) {
					if toA {
						a += str
					} else {
						b += str
					}
				} else if str == ")" {
					safe = true
					break
				} else {
					break
				}
			}

			if safe && valid(a, b) {
				anum, _ := strconv.Atoi(a)
				bnum, _ := strconv.Atoi(b)

				result += anum * bnum
			}

			index = lastIndex
		} else {
			index += 1
		}
	}

	return result
}

func Day3P2(content string) int {
	result := 0

	enabled := true
	index := 0

	for index+7 < len(content) {
		if content[index:index+7] == "don't()" {
			enabled = false
			index = index + 7
			continue
		} else if content[index:index+4] == "do()" {
			enabled = true
			index = index + 4
			continue
		}

		s := content[index : index+4]

		if enabled && s == "mul(" {
			a := ""
			b := ""
			toA := true
			safe := false
			lastIndex := index + 4
			for i, r := range content[index+4:] {
				str := string(r)
				lastIndex = index + 4 + i
				if str == "," {
					toA = false
				} else if unicode.IsDigit(r) {
					if toA {
						a += str
					} else {
						b += str
					}
				} else if str == ")" {
					safe = true
					break
				} else {
					break
				}
			}

			if safe && valid(a, b) {
				anum, _ := strconv.Atoi(a)
				bnum, _ := strconv.Atoi(b)

				result += anum * bnum
			}

			index = lastIndex
		} else {
			index += 1
		}
	}

	return result
}

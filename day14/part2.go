package day14

import (
	"sort"
	"strings"
)

func StringReverse(s string) string {
	result := ""

	for _, ch := range s {
		result = string(ch) + result
	}

	return result
}

type SortByRocks []byte

func (a SortByRocks) Len() int           { return len(a) }
func (a SortByRocks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByRocks) Less(i, j int) bool { return a[i] > a[j] }

func zip(group string) string {
	splitted := strings.Split(group, "\r\n")

	acc := []string{}

	col := len(splitted[0])
	row := len(splitted)

	for c := 0; c < col; c++ {
		a := make([]string, row)

		for r := 0; r < row; r++ {
			a = append(a, string(splitted[r][c]))
		}

		acc = append(acc, strings.Join(a, ""))
	}

	return strings.Join(acc, "\r\n")
}

func toString(b []byte) string {
	r := ""

	for _, k := range b {
		r = r + string(k)
	}

	return r
}

func rocksToTheLeft(content string) string {
	s := strings.Split(content, "\r\n")

	res := []string{}

	for _, row := range s {
		c := strings.Split(row, "#")
		acc := []string{}

		for _, k := range c {
			b := []byte(k)

			sort.Sort(SortByRocks(b))

			acc = append(acc, toString(b))
		}

		res = append(res, strings.Join(acc, "#"))
	}

	return strings.Join(res, "\r\n")
}

func moveToNorth(content string) string {
	return rocksToTheLeft(zip(content))
}

func transform(content string) string {
	r := content

	for i := 0; i < 4; i++ {
		k := moveToNorth(r)

		n := []string{}
		splitted := strings.Split(k, "\r\n")

		for _, k := range splitted {
			n = append(n, StringReverse(k))
		}

		r = strings.Join(n, "\r\n")
	}

	return r
}

func Part2(content string) int {
	qwe := strings.Split(content, "\r\n")

	l := strings.Join(qwe, "\r\n")

	visited := map[string]bool{l: true}
	arr := []string{l}
	index := 0

	for {
		index += 1
		l = transform(l)

		if visited[l] {
			break
		}

		visited[l] = true
		arr = append(arr, l)
	}

	first := 0

	for i, c := range arr {
		if c == l {
			first = i

			break
		}
	}

	c := strings.Split(arr[(1000000000-first)%(index-first)+first], "\r\n")

	result := 0

	for i, line := range c {
		zeros := 0

		for _, k := range line {
			if k == 'O' {
				zeros++
			}
		}

		result += zeros * (len(c) - i)
	}

	return result
}

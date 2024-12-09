package aoc_2024

import (
	"fmt"
	"slices"
	"strings"
)

func Day9P1(content string) string {
	s := strings.Split(content, "")
	nums := toInts(s...)

	fileId := 0

	res := []string{}

	for index, num := range nums {
		add := "."
		if index%2 == 0 {
			add = fmt.Sprintf("%d", fileId)
			fileId += 1
		}

		for i := 0; i < num; i++ {
			res = append(res, add)
		}
	}

	count := 0
	rightPointer := len(res) - 1

	for index, ch := range res {
		if index > rightPointer {
			break
		}

		if ch == "." {
			count += index * toInts(res[rightPointer])[0]

			for i := rightPointer - 1; i >= 0; i-- {
				if res[i] == "." {
					continue
				}

				rightPointer = i
				break
			}
		} else {
			count += index * toInts(ch)[0]
		}
	}

	return fmt.Sprintf("%d", count)
}

func Day9P2(content string) string {
	s := strings.Split(content, "")
	nums := toInts(s...)

	type File struct {
		id   int
		size int
	}

	all := []File{}
	fileId := 0

	for index, num := range nums {
		id := -1
		if index%2 == 0 {
			id = fileId
			fileId++
		}

		all = append(all, File{id: id, size: num})
	}

	for file := len(all) - 1; file >= 0; file-- {
		for free := 0; free < file; free++ {
			if all[file].id != -1 && all[free].id == -1 && all[free].size >= all[file].size {
				all = slices.Insert(all, free, all[file])
				all[file+1].id = -1
				all[free+1].size = all[free+1].size - all[file+1].size
			}
		}
	}

	count := 0

	a := []int{}

	for _, file := range all {
		a = append(a, slices.Repeat([]int{file.id}, file.size)...)
	}

	for index, value := range a {
		if value != -1 {
			count += value * index
		}
	}

	return fmt.Sprintf("%d", count)
}

package day7

import (
	"aoc/utils"
	"sort"
	"strings"
)

type HandType = int

const (
	HIGH HandType = iota + 1
	ONE_PAIR
	TWO_PAIR
	THREE_OAK
	FULL_HOUSE
	FOUR_OAK
	FIVE_OAK
)

var CARD_RATING = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type Eval struct {
	value  int
	origin string
	hand   HandType
}

type EvalArr []Eval

func compareSameTypes(e1 Eval, e2 Eval) bool {
	i := 0

	origin_a := e1.origin
	origin_b := e2.origin

	for i < 5 {
		a := CARD_RATING[origin_a[i]]
		b := CARD_RATING[origin_b[i]]

		if a < b {
			return false
		} else if a > b {
			break
		}

		i += 1
	}

	return true
}

func (a EvalArr) Len() int      { return len(a) }
func (a EvalArr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a EvalArr) Less(i, j int) bool {
	if a[i].hand < a[j].hand {
		return false
	}

	if a[i].hand == a[j].hand {
		return compareSameTypes(a[i], a[j])
	}

	return true
}

func eval(str string) Eval {
	dict := map[rune]int{}

	_type := HIGH

	j_count := 0

	for _, i := range str {
		r := rune(i)

		if r == 'J' {
			j_count += 1

			continue
		}

		value, ok := dict[r]

		if ok {
			dict[r] = value + 1
		} else {
			dict[r] = 1
		}
	}

	dictlen := len(dict)

	keys := []rune{}

	for key := range dict {
		keys = append(keys, key)
	}

	if j_count > 0 {
		if j_count == 5 {
			return Eval{value: 0, origin: str, hand: FIVE_OAK}
		}

		k := keys[0]
		m := dict[k]

		for ke, ve := range dict {
			if ve > m {
				k = ke
				m = ve
			}
		}

		dict[k] = m + j_count
	}

	if dictlen == 1 {
		_type = FIVE_OAK
	} else if dictlen == 2 {
		f := dict[keys[0]]

		if f == 4 || f == 1 {
			_type = FOUR_OAK
		} else {
			_type = FULL_HOUSE
		}
	} else if dictlen == 3 {
		a := dict[keys[0]]
		b := dict[keys[1]]
		c := dict[keys[2]]

		if a == 3 || b == 3 || c == 3 {
			_type = THREE_OAK
		} else {
			_type = TWO_PAIR
		}

	} else if dictlen == 4 {
		_type = ONE_PAIR
	}

	return Eval{value: 0, origin: str, hand: _type}
}

func Part1(content string) int {
	handsAndBid := strings.Split(content, "\r\n")

	acc := []Eval{}

	for _, i := range handsAndBid {
		c := strings.Split(i, " ")

		hand := strings.TrimSpace(c[0])
		bid := utils.LineToNums(c[1])[0]

		a := eval(hand)
		a.value = bid

		acc = append(acc, a)
	}

	sort.Sort(EvalArr(acc))

	l := len(acc)
	result := 0

	for index, item := range acc {
		rank := l - index

		result += item.value * rank
	}

	return result
}

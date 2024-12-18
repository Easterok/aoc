package aoc_2024

import (
	"fmt"
	"math"
	"strings"
)

func day17run(opcode, operand int, regA, regB, regC *int, program []int, i *int) int {
	combo := operand

	if operand == 4 {
		combo = *regA
	} else if operand == 5 {
		combo = *regB
	} else if operand == 6 {
		combo = *regC
	}

	out := -1

	switch opcode {
	case 0:
		*regA = *regA / int(math.Pow(2, float64(combo)))
		*i += 2
	case 1:
		*regB = *regB ^ operand
		*i += 2
	case 2:
		*regB = combo % 8
		*i += 2
	case 3:
		if *regA != 0 {
			*i = program[*i+1]
		} else {
			*i += 2
		}
	case 4:
		*regB = *regB ^ *regC
		*i += 2
	case 5:
		out = combo % 8
		*i += 2
	case 6:
		*regB = *regA / int(math.Pow(2, float64(combo)))
		*i += 2
	case 7:
		*regC = *regA / int(math.Pow(2, float64(combo)))
		*i += 2
	}

	return out
}

func Day17P1(content string) string {
	c := strings.Split(content, "\n\n")
	registers := strings.Split(c[0], "\n")

	regA := toInts(registers[0][len("Register A: "):])[0]
	regB := toInts(registers[1][len("Register B: "):])[0]
	regC := toInts(registers[2][len("Register C: "):])[0]

	program := toInts(strings.Split(c[1][len("Program: "):], ",")...)

	count := []string{}

	i := 0

	for i < len(program) {
		opcode := program[i]
		operand := program[i+1]

		res := day17run(opcode, operand, &regA, &regB, &regC, program, &i)

		if res != -1 {
			count = append(count, fmt.Sprintf("%d", res))
		}
	}

	return strings.Join(count, ",")
}

func day17back(program []int, pos int, init int) int {
	for idx := 0; idx < 8; idx++ {
		regA := init*8 + idx
		regB := 0
		regC := 0

		i := 0

		output := []int{}

		for i < len(program) {
			opcode := program[i]
			operand := program[i+1]

			res := day17run(opcode, operand, &regA, &regB, &regC, program, &i)

			if res != -1 {
				output = append(output, res)
			}
		}

		ok := true

		for i := pos; i < len(program); i++ {
			if output[i-pos] != program[i] {
				ok = false
				break
			}
		}

		if ok {
			if pos == 0 {
				return init*8 + idx
			}

			val := day17back(program, pos-1, init*8+idx)

			if val != -1 {
				return val
			}
		}
	}

	return -1
}

func Day17P2(content string) string {
	c := strings.Split(content, "\n\n")

	program := toInts(strings.Split(c[1][len("Program: "):], ",")...)

	res := day17back(program, len(program)-1, 0)

	return fmt.Sprintf("%d", res)
}

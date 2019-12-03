package part1

import (
	"fmt"
)

func RunProgram(input []int) error {
	i := 0
	var op int
	var a int
	var b int
	var pos int

	for {
		op = input[i]

		if op == 99 {
			break
		}

		a = input[i+1]
		b = input[i+2]
		pos = input[i+3]
		i += 4

		if op == 1 {
			input[pos] = input[a] + input[b]
		} else if op == 2 {
			input[pos] = input[a] * input[b]
		} else {
			return fmt.Errorf("invalid opcode: %d", op)
		}
	}

	return nil
}

func Solution(input []int) (int, error) {
	input[1] = 12
	input[2] = 2
	err := RunProgram(input)
	return input[0], err

}

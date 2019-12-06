package part1

import (
	"github.com/bope/adventofcode2019-go/day5"
)

func Solution(program, input []int) (int, error) {
	e := day5.NewEmulator(program, input)
	output, err := e.Run()
	if err != nil {
		return 0, err
	}
	return output[len(output)-1], nil
}

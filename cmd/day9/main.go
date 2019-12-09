package main

import (
	"fmt"
	"os"

	"github.com/bope/adventofcode2019-go/day9/part1"
	"github.com/bope/adventofcode2019-go/intcode"
)

func main() {
	program, err := intcode.Parse(os.Stdin)
	program_cpy := make([]int, len(program))

	if err != nil {
		panic(err)
	}

	copy(program_cpy, program)

	ret, err := part1.Solution(program_cpy, []int{1})
	if err != nil {
		panic(err)
	}
	fmt.Printf("part1: %d\n", ret[len(ret)-1])

	ret, err = part1.Solution(program_cpy, []int{2})
	if err != nil {
		panic(err)
	}
	fmt.Printf("part2: %d\n", ret[len(ret)-1])

}

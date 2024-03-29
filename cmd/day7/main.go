package main

import (
	"fmt"
	"os"

	"github.com/bope/adventofcode2019-go/day7/part1"
	"github.com/bope/adventofcode2019-go/day7/part2"
	"github.com/bope/adventofcode2019-go/intcode"
)

func main() {
	program, err := intcode.Parse(os.Stdin)
	program_cpy := make([]int, len(program))

	if err != nil {
		panic(err)
	}

	copy(program_cpy, program)

	ret := part1.Solution(program_cpy)
	if err != nil {
		panic(err)
	}
	fmt.Printf("part1: %d\n", ret)

	ret = part2.Solution(program_cpy)
	if err != nil {
		panic(err)
	}
	fmt.Printf("part2: %d\n", ret)

}

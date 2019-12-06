package main

import (
	"fmt"
	"os"

	"github.com/bope/adventofcode2019-go/day5"
	"github.com/bope/adventofcode2019-go/day5/part1"
	"github.com/bope/adventofcode2019-go/day5/part2"
)

func main() {
	program, err := day5.Parse(os.Stdin)
	program_cpy := make([]int, len(program))

	if err != nil {
		panic(err)
	}

	copy(program_cpy, program)

	ret, err := part1.Solution(program_cpy, []int{1})
	if err != nil {
		panic(err)
	}
	fmt.Printf("part1: %d\n", ret)

	ret, err = part2.Solution(program, []int{5})
	if err != nil {
		panic(err)
	}
	fmt.Printf("part2: %d\n", ret)

}

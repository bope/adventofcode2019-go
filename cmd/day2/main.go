package main

import (
	"fmt"
	"github.com/bope/adventofcode2019-go/day2"
	"github.com/bope/adventofcode2019-go/day2/part1"
	"github.com/bope/adventofcode2019-go/day2/part2"
	"os"
)

func main() {
	input, err := day2.Parse(os.Stdin)

	if err != nil {
		panic(err)
	}

	input1 := make([]int, len(input))
	copy(input1, input)

	result, err := part1.Solution(input1)

	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %v\n", result)

	result, err = part2.Solution(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part2: %v\n", result)
}

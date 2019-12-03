package main

import (
	"fmt"
	"github.com/bope/adventofcode2019-go/day1"
	"github.com/bope/adventofcode2019-go/day1/part1"
	"github.com/bope/adventofcode2019-go/day1/part2"
	"os"
)

func main() {
	input, err := day1.Parse(os.Stdin)

	if err != nil {
		panic(err)
	}

	result1 := part1.Solution(input)
	fmt.Printf("part1: %v\n", result1)

	result2 := part2.Solution(input)
	fmt.Printf("part2: %v\n", result2)
}

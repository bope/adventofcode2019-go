package main

import (
	"fmt"
	"os"

	"github.com/bope/adventofcode2019-go/day6"
	"github.com/bope/adventofcode2019-go/day6/part1"
	"github.com/bope/adventofcode2019-go/day6/part2"
)

func main() {
	input, err := day6.Parse(os.Stdin)

	if err != nil {
		panic(err)
	}

	ret := part1.Solution(input)
	fmt.Printf("part1: %d\n", ret)

	ret = part2.Solution(input)
	fmt.Printf("part2: %d\n", ret)
}

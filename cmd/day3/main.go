package main

import (
	"fmt"
	"github.com/bope/adventofcode2019-go/day3"
	"github.com/bope/adventofcode2019-go/day3/part1"
	"github.com/bope/adventofcode2019-go/day3/part2"
	"os"
)

func main() {
	a, b, err := day3.Parse(os.Stdin)

	if err != nil {
		panic(err)
	}
	ret := part1.Solution(a, b)
	fmt.Printf("part1: %d\n", ret)

	ret = part2.Solution(a, b)
	fmt.Printf("part2: %d\n", ret)

}

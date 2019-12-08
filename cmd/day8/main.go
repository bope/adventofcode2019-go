package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bope/adventofcode2019-go/day8"
	"github.com/bope/adventofcode2019-go/day8/part1"
	"github.com/bope/adventofcode2019-go/day8/part2"
)

func main() {
	input, err := day8.Parse(os.Stdin)

	if err != nil {
		panic(err)
	}

	ret := part1.Solution(input, 25, 6)
	fmt.Printf("part1: %d\n", ret)

	sret := part2.Solution(input, 25, 6)
	sret = strings.Replace(sret, "2", " ", -1)
	sret = strings.Replace(sret, "0", ".", -1)
	sret = strings.Replace(sret, "1", "X", -1)

	fmt.Printf("part2:\n%s", sret)

}

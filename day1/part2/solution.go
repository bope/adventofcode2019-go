package part2

import (
	"github.com/bope/adventofcode2019-go/day1/part1"
)

func CalcMass(mass int) int {
	total := 0
	for {
		mass = part1.CalcMass(mass)
		if mass <= 0 {
			break
		}

		total += mass

	}
	return total
}

func Solution(input []int) int {
	total := 0
	for _, mass := range input {
		total += CalcMass(mass)
	}
	return total
}

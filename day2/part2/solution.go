package part2

import "github.com/bope/adventofcode2019-go/day2/part1"

import "fmt"

func Solution(input []int) (int, error) {
	inputCopy := make([]int, len(input))
	var err error

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(inputCopy, input)
			inputCopy[1] = noun
			inputCopy[2] = verb
			if err = part1.RunProgram(inputCopy); err != nil {
				return 0, err
			}

			if inputCopy[0] == 19690720 {
				return 100*noun + verb, nil
			}
		}
	}

	return 0, fmt.Errorf("no solution found")
}

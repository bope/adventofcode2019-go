package part1

import (
	"github.com/bope/adventofcode2019-go/intcode"
	"sync"
)

func Solution(program, input []int) ([]int, error) {
	inputc := make(chan int)
	outputc := make(chan int)
	output := make([]int, 0)
	var wg sync.WaitGroup
	e := intcode.New(program, inputc, outputc)

	wg.Add(1)
	go func() {
		for _, i := range input {
			inputc <- i
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for o := range outputc {
			output = append(output, o)
		}
		wg.Done()
	}()

	if err := e.Run(); err != nil {
		return nil, err
	}

	close(inputc)
	close(outputc)

	wg.Wait()
	return output, nil
}

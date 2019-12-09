package part1

import (
	"sync"

	"github.com/bope/adventofcode2019-go/day7"
	"github.com/bope/adventofcode2019-go/intcode"
)

func run_with_phases(program, phases []int) int {
	emulators := make([]*intcode.Emulator, len(phases))

	for i, phase := range phases {
		p := make([]int, len(program))
		copy(p, program)

		var input chan int
		if i != 0 {
			input = emulators[i-1].Output
		} else {
			input = make(chan int, 2)
		}
		output := make(chan int, 2)

		emulators[i] = intcode.New(p, input, output)
		input <- phase
	}

	emulators[0].Input <- 0
	var wg sync.WaitGroup
	for _, e := range emulators {
		wg.Add(1)
		go func(e *intcode.Emulator) {
			defer func() {
				wg.Done()
			}()
			if err := e.Run(); err != nil {
				panic(err)
			}

		}(e)
	}
	wg.Wait()
	return <-emulators[len(emulators)-1].Output

}

func solution(program []int) (int, []int) {
	var max_phases []int
	var max_signal int

	for _, phases := range day7.Permutations([]int{0, 1, 2, 3, 4}) {
		signal := run_with_phases(program, phases)
		if signal > max_signal {
			max_signal = signal
			max_phases = phases
		}
	}
	return max_signal, max_phases
}

func Solution(program []int) int {
	ret, _ := solution(program)
	return ret
}

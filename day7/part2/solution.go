package part2

import (
	"sync"

	"github.com/bope/adventofcode2019-go/day7"
)

func run_with_phases(program, phases []int) int {
	emulators := make([]*day7.Emulator, len(phases))

	for i := range phases {
		p := make([]int, len(program))
		copy(p, program)

		var input chan int
		var output chan int

		if i == 0 {
			input = make(chan int, 10)
			output = make(chan int, 10)
		} else if i == len(phases)-1 {
			input = emulators[i-1].Output
			output = emulators[0].Input
		} else {
			input = emulators[i-1].Output
			output = make(chan int, 10)

		}
		emulators[i] = day7.NewEmulator(p, input, output)
	}

	for i, phase := range phases {
		emulators[i].Input <- phase
	}
	emulators[0].Input <- 0

	var wg sync.WaitGroup
	for _, e := range emulators {
		wg.Add(1)
		go func(e *day7.Emulator) {
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

	for _, phases := range day7.Permutations([]int{5, 6, 7, 8, 9}) {
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

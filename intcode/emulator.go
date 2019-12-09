package intcode

import "fmt"

type Emulator struct {
	memory  []int
	ptr     int
	rptr    int
	running bool

	Input  chan int
	Output chan int
}

func ParseInstruction(inst int) (int, [3]int) {
	op := inst % 100
	modes := [3]int{
		inst / 100 % 10,
		inst / 1000 % 10,
		inst / 10000 % 10,
	}
	return op, modes
}

func New(program []int, input, output chan int) *Emulator {
	return &Emulator{
		memory:  program,
		ptr:     0,
		rptr:    0,
		running: true,
		Input:   input,
		Output:  output,
	}
}

func (e *Emulator) checkMemory(ptr int) {
	if ptr >= len(e.memory) {
		e.memory = append(e.memory, make([]int, ptr-len(e.memory)+1)...)
	}
}

func (e *Emulator) resolvePointer(mode, val int) (int, error) {
	var ptr int
	switch mode {
	case 0:
		ptr = e.memory[val]
	case 1:
		ptr = val
	case 2:
		ptr = e.memory[val] + e.rptr
	default:
		return ptr, fmt.Errorf("invalid mode: %d", mode)
	}

	if ptr < 0 {
		return ptr, fmt.Errorf("negative pointer: %d", ptr)
	}

	return ptr, nil
}

func (e *Emulator) get(mode, ptr int) (*int, error) {
	var err error
	if ptr, err = e.resolvePointer(mode, ptr); err != nil {
		return nil, err
	}

	e.checkMemory(ptr)
	return &e.memory[ptr], nil
}

func (e *Emulator) params(count int, modes [3]int) ([]*int, error) {
	params := make([]*int, count)
	var err error
	for i := 0; i < count; i++ {
		if params[i], err = e.get(modes[i], e.ptr+i); err != nil {
			return nil, err
		}
	}
	e.ptr += count
	return params, nil
}

func (e *Emulator) Run() error {
	var err error

	for e.running {
		if err = e.Step(); err != nil {
			return err
		}
	}
	return nil
}

func (e *Emulator) runOp(f func([]*int), count int, modes [3]int) error {
	params, err := e.params(count, modes)
	if err != nil {
		return err
	}
	f(params)
	return nil
}

func (e *Emulator) Step() error {
	op, modes := ParseInstruction(e.memory[e.ptr])
	e.ptr++
	var err error

	switch op {
	case 99:
		e.running = false
	case 1:
		err = e.runOp(e.add, 3, modes)
	case 2:
		err = e.runOp(e.mul, 3, modes)
	case 3:
		err = e.runOp(e.input, 1, modes)
	case 4:
		err = e.runOp(e.output, 1, modes)
	case 5:
		err = e.runOp(e.jumptrue, 2, modes)
	case 6:
		err = e.runOp(e.jumpfalse, 2, modes)
	case 7:
		err = e.runOp(e.lessthan, 3, modes)
	case 8:
		err = e.runOp(e.equals, 3, modes)
	case 9:
		err = e.runOp(e.rel, 1, modes)
	default:
		err = fmt.Errorf("invalid opcode: %d", op)
	}

	return err
}

func (e *Emulator) add(params []*int) {
	*params[2] = *params[0] + *params[1]
}

func (e *Emulator) mul(params []*int) {
	*params[2] = *params[0] * *params[1]
}

func (e *Emulator) input(params []*int) {
	*params[0] = <-e.Input
}

func (e *Emulator) output(params []*int) {
	e.Output <- *params[0]
}

func (e *Emulator) jumptrue(params []*int) {
	if *params[0] != 0 {
		e.ptr = *params[1]
	}
}

func (e *Emulator) jumpfalse(params []*int) {
	if *params[0] == 0 {
		e.ptr = *params[1]
	}
}

func (e *Emulator) lessthan(params []*int) {
	if *params[0] < *params[1] {
		*params[2] = 1
	} else {
		*params[2] = 0
	}
}

func (e *Emulator) equals(params []*int) {
	if *params[0] == *params[1] {
		*params[2] = 1
	} else {
		*params[2] = 0
	}
}

func (e *Emulator) rel(params []*int) {
	e.rptr = e.rptr + (*params[0])
}

package day7

import "fmt"

type Emulator struct {
	program []int
	ptr     int
	iptr    int
	Input   chan int
	Output  chan int
}

type Instruction struct {
	OpCode int
	Modes  [3]int
}

func ParseInstruction(inst int) Instruction {
	return Instruction{
		OpCode: inst % 100,
		Modes: [3]int{
			inst / 100 % 10,
			inst / 1000 % 10,
			inst / 10000 % 10,
		},
	}
}

func NewEmulator(program []int, input, output chan int) *Emulator {
	return &Emulator{
		program: program,
		ptr:     0,
		iptr:    0,
		Input:   input,
		Output:  output,
	}
}

func (e *Emulator) params(count int, modes [3]int) []int {
	params := make([]int, count)
	for i := 0; i < count; i++ {
		if modes[i] == 1 {
			params[i] = e.program[e.ptr+i]
		} else {
			params[i] = e.program[e.program[e.ptr+i]]
		}
	}
	e.ptr += count
	return params
}

func (e *Emulator) Run() error {
	var inst Instruction

	for {
		inst = ParseInstruction(e.program[e.ptr])
		e.ptr++

		switch inst.OpCode {
		case 99:
			return nil
		case 1:
			e.add(inst.Modes)
		case 2:
			e.mul(inst.Modes)
		case 3:
			e.input(inst.Modes)
		case 4:
			e.output(inst.Modes)
		case 5:
			e.jumptrue(inst.Modes)
		case 6:
			e.jumpfalse(inst.Modes)
		case 7:
			e.lessthan(inst.Modes)
		case 8:
			e.equals(inst.Modes)
		default:
			return fmt.Errorf("invalid opcode: %d", inst.OpCode)
		}
	}
}

func (e *Emulator) add(modes [3]int) {
	modes[2] = 1
	params := e.params(3, modes)
	e.program[params[2]] = params[0] + params[1]
}

func (e *Emulator) mul(modes [3]int) {
	modes[2] = 1
	params := e.params(3, modes)
	e.program[params[2]] = params[0] * params[1]
}

func (e *Emulator) input(modes [3]int) {
	modes[0] = 1
	param := e.params(1, modes)
	e.program[param[0]] = <-e.Input
}

func (e *Emulator) output(modes [3]int) {
	params := e.params(1, modes)
	e.Output <- params[0]
}

func (e *Emulator) jumptrue(modes [3]int) {
	params := e.params(2, modes)
	if params[0] != 0 {
		e.ptr = params[1]
	}
}

func (e *Emulator) jumpfalse(modes [3]int) {
	params := e.params(2, modes)
	if params[0] == 0 {
		e.ptr = params[1]
	}
}

func (e *Emulator) lessthan(modes [3]int) {
	modes[2] = 1
	params := e.params(3, modes)
	if params[0] < params[1] {
		e.program[params[2]] = 1
	} else {
		e.program[params[2]] = 0
	}
}

func (e *Emulator) equals(modes [3]int) {
	modes[2] = 1
	params := e.params(3, modes)
	if params[0] == params[1] {
		e.program[params[2]] = 1
	} else {
		e.program[params[2]] = 0
	}
}

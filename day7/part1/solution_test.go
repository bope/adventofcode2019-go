package part1

import (
	"strings"
	"testing"

	"github.com/bope/adventofcode2019-go/intcode"
)

type testCase struct {
	program         string
	expected_signal int
	expected_phases []int
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestSolution(t *testing.T) {
	tests := []testCase{
		testCase{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", 43210, []int{4, 3, 2, 1, 0}},
		testCase{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", 54321, []int{0, 1, 2, 3, 4}},
		testCase{"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", 65210, []int{1, 0, 4, 3, 2}},
	}

	for _, test := range tests {
		program, err := intcode.Parse(strings.NewReader(test.program))
		if err != nil {
			t.Errorf("parse error: %s", err.Error())
		}
		sig, phases := solution(program)
		if sig != test.expected_signal {
			t.Errorf("program: %v signal != %d: %d", program, sig, test.expected_signal)
		}

		if !equalSlice(phases, test.expected_phases) {
			t.Errorf("program: %v phases != %d: %d", program, phases, test.expected_phases)
		}
	}
}

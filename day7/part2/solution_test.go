package part2

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
		testCase{"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", 139629729, []int{9, 8, 7, 6, 5}},
		testCase{"3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", 18216, []int{9, 7, 8, 5, 6}},
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

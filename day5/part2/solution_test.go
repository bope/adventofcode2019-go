package part2

import (
	"strings"
	"testing"

	"github.com/bope/adventofcode2019-go/day5"
)

type testCase struct {
	program  string
	input    []int
	expected int
}

func TestSolution(t *testing.T) {
	tests := []testCase{
		testCase{"3,9,8,9,10,9,4,9,99,-1,8", []int{8}, 1},
		testCase{"3,9,8,9,10,9,4,9,99,-1,8", []int{9}, 0},

		testCase{"3,9,7,9,10,9,4,9,99,-1,8", []int{7}, 1},
		testCase{"3,9,7,9,10,9,4,9,99,-1,8", []int{8}, 0},

		testCase{"3,3,1108,-1,8,3,4,3,99", []int{8}, 1},
		testCase{"3,3,1108,-1,8,3,4,3,99", []int{7}, 0},

		testCase{"3,3,1107,-1,8,3,4,3,99", []int{7}, 1},
		testCase{"3,3,1107,-1,8,3,4,3,99", []int{8}, 0},

		testCase{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", []int{0}, 0},
		testCase{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", []int{1}, 1},

		testCase{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", []int{7}, 999},
		testCase{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", []int{8}, 1000},
		testCase{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", []int{9}, 1001},
	}

	for _, test := range tests {
		program, err := day5.Parse(strings.NewReader(test.program))

		if err != nil {
			t.Errorf("parse error: %s", err.Error())
		}

		ret, err := Solution(program, test.input)
		if err != nil {
			t.Errorf("error: %s", err.Error())
		}

		if ret != test.expected {
			t.Errorf("program: %s input: %v != %d: %d", test.program, test.input, test.expected, ret)
		}
	}
}

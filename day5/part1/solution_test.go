package part1

import (
	"strings"
	"testing"

	"github.com/bope/adventofcode2019-go/intcode"
)

type testCase struct {
	program  string
	input    []int
	expected int
}

func TestSolution(t *testing.T) {
	tests := []testCase{
		testCase{"3,0,4,0,99", []int{1}, 1},
	}

	for _, test := range tests {
		program, err := intcode.Parse(strings.NewReader(test.program))

		if err != nil {
			t.Errorf("parse error: %s", err.Error())
		}

		ret, err := Solution(program, test.input)
		if err != nil {
			t.Errorf("error: %s", err.Error())
		}

		if ret != test.expected {
			t.Errorf("%s != %d: %d", test.program, test.expected, ret)
		}
	}
}

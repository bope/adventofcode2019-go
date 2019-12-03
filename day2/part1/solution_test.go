package part1

import (
	"strings"
	"testing"

	"github.com/bope/adventofcode2019-go/day2"
)

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

type testProgram struct {
	input    string
	expected []int
}

func TestProgram(t *testing.T) {
	testPrograms := []testProgram{
		testProgram{"1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		testProgram{"2,3,0,3,99", []int{2, 3, 0, 6, 99}},
		testProgram{"2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
		testProgram{"1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, test := range testPrograms {
		input, err := day2.Parse(strings.NewReader(test.input))

		if err != nil {
			t.Errorf("parse failed: %s", err)
		}

		if err = RunProgram(input); err != nil {
			t.Errorf("program error: %s", err)
		}

		if !equalSlice(input, test.expected) {
			t.Errorf("%v != %v", input, test.expected)
		}
	}
}

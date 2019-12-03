package part1

import (
	"strings"
	"testing"

	"github.com/bope/adventofcode2019-go/day1"
)

type testCase struct {
	input    string
	expected int
}

func TestSolution(t *testing.T) {
	tests := []testCase{
		testCase{"12", 2},
		testCase{"14", 2},
		testCase{"1969", 654},
		testCase{"100756", 33583},
	}

	for _, test := range tests {
		input, err := day1.Parse(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("parse error: %s", err)
		}

		if res := Solution(input); res != test.expected {
			t.Errorf("%v != %v", input, test.expected)
		}
	}
}

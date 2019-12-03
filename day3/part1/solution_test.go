package part1

import (
	"github.com/bope/adventofcode2019-go/day3"
	"strings"
	"testing"
)

type testCase struct {
	input    string
	expected int
}

func TestSolution(t *testing.T) {
	tests := []testCase{
		testCase{
			input:    "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
			expected: 159,
		},
		testCase{
			input:    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			expected: 135,
		},
	}

	for _, test := range tests {
		w1, w2, err := day3.Parse(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("parse failed: %s", err.Error())
		}

		if ret := Solution(w1, w2); ret != test.expected {
			t.Errorf("%v != %v", ret, test.expected)
		}
	}
}

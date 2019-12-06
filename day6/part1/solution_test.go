package part1

import (
	"strings"
	"testing"

	"github.com/bope/adventofcode2019-go/day6"
)

func TestSolution(t *testing.T) {
	t1i := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	input, err := day6.Parse(strings.NewReader(t1i))
	if err != nil {
		t.Errorf("parse error: %s", err.Error())
	}

	if ret := Solution(input); ret != 42 {
		t.Errorf("%d != 42", ret)
	}
}

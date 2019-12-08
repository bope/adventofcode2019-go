package part2

import "testing"

func TestSolution(t *testing.T) {
	if ret := Solution([]rune("0222112222120000"), 2, 2); ret != "01\n10" {
		t.Errorf("0222112222120000 != \n 01\n10\n: %s", ret)
	}
}

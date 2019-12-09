package part1

import (
	"strconv"
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
	t1i := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	t1p, err := intcode.Parse(strings.NewReader(t1i))
	if err != nil {
		t.Errorf("parse error: %s", err.Error())
	}

	ret, err := Solution(t1p, []int{})
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if !equalSlice(t1p, ret) {
		t.Errorf("%v != %v", t1p, ret)
	}

	t2i := "1102,34915192,34915192,7,4,7,99,0"
	t2p, err := intcode.Parse(strings.NewReader(t2i))
	if err != nil {
		t.Errorf("parse error: %s", err.Error())
	}

	ret, err = Solution(t2p, []int{})
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	rets := strconv.Itoa(ret[len(ret)-1])
	if len(rets) != 16 {
		t.Errorf("len(%s) != 16", rets)
	}

	t3i := "104,1125899906842624,99"
	t3p, err := intcode.Parse(strings.NewReader(t3i))
	if err != nil {
		t.Errorf("parse error: %s", err.Error())
	}

	ret, err = Solution(t3p, []int{})
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if ret[len(ret)-1] != 1125899906842624 {
		t.Errorf("%d != 1125899906842624", ret)
	}
}

package part1

import "testing"

func TestIsValid(t *testing.T) {

	if !IsValid(111111) {
		t.Errorf("111111 is not valid")
	}

	if IsValid(223450) {
		t.Errorf("223450 is valid")
	}

	if IsValid(123789) {
		t.Errorf("123789 is valid")
	}
}

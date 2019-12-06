package part2

import "testing"

func TestIsValid(t *testing.T) {

	if !IsValid(112233) {
		t.Errorf("112233 is not valid")
	}

	if IsValid(123444) {
		t.Errorf("123444 is valid")
	}

	if !IsValid(111122) {
		t.Errorf("111122 is not valid")
	}
}

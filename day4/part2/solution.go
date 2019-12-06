package part2

import "strconv"

func IsValid(num int) bool {
	pwr := []rune(strconv.Itoa(num))
	l := len(pwr)
	adj := 0
	found := false

	for i := 0; i < l-1; i++ {
		if pwr[i] > pwr[i+1] {
			return false
		}
		if pwr[i] == pwr[i+1] {
			adj += 1
		} else {
			if adj == 1 {
				found = true
			}
			adj = 0
		}
	}

	if adj == 1 {
		found = true
	}

	return found

}

func Solution(minVal, maxVal int) int {
	var check bool
	ret := 0
	for i := minVal; i <= maxVal; i++ {
		if check = IsValid(i); check {
			ret += 1
		}
	}
	return ret
}

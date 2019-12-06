package part1

import "strconv"

func IsValid(num int) bool {
	pwr := []rune(strconv.Itoa(num))
	l := len(pwr)
	adj := false

	for i := 0; i < l-1; i++ {
		if pwr[i] > pwr[i+1] {
			return false
		}

		if pwr[i] == pwr[i+1] {
			adj = true
		}
	}

	return adj

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

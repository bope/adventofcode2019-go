package day3

type Vec [2]int

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (v Vec) Add(other Vec) Vec {
	return Vec{v[0] + other[0], v[1] + other[1]}
}

func (v Vec) Sub(other Vec) Vec {
	return Vec{v[0] - other[0], v[1] - other[1]}
}

func (v Vec) Distance() int {
	return abs(v[0]) + abs(v[1])
}

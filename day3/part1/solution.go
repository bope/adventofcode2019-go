package part1

import (
	"github.com/bope/adventofcode2019-go/day3"
)

func mapPath(path []day3.Step) map[day3.Vec]bool {
	ret := make(map[day3.Vec]bool)
	pos := day3.Vec{0, 0}
	for _, step := range path {
		for i := 0; i < step.Length; i++ {
			pos = pos.Add(step.Direction)
			ret[pos] = true
		}
	}
	return ret
}

func min(vs []int) int {
	m := vs[0]
	for _, v := range vs {
		if v < m {
			m = v
		}
	}
	return m
}

func Solution(wp1, wp2 []day3.Step) int {
	wp1map := mapPath(wp1)
	wp2map := mapPath(wp2)

	start := day3.Vec{0, 0}
	dists := make([]int, 0)

	for w1, _ := range wp1map {
		_, ok := wp2map[w1]
		if !ok {
			continue
		}
		dists = append(dists, start.Sub(w1).Distance())
	}
	return min(dists)
}

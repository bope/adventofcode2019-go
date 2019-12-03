package part2

import (
	"github.com/bope/adventofcode2019-go/day3"
)

func mapPath(path []day3.Step) map[day3.Vec]int {
	ret := make(map[day3.Vec]int)
	pos := day3.Vec{0, 0}
	count := 0
	for _, step := range path {
		for i := 0; i < step.Length; i++ {
			count += 1
			pos = pos.Add(step.Direction)
			_, found := ret[pos]
			if found {
				continue
			}
			ret[pos] = count
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

	dists := make([]int, 0)

	for w1, v1 := range wp1map {
		v2, ok := wp2map[w1]
		if !ok {
			continue
		}
		dists = append(dists, v1+v2)
	}
	return min(dists)
}

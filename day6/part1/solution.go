package part1

import (
	"github.com/bope/adventofcode2019-go/day6"
)

func Solution(input [][2]string) int {
	nodes := make(map[string]*day6.Node)

	for _, rel := range input {
		if _, ok := nodes[rel[0]]; !ok {
			nodes[rel[0]] = day6.NewNode(rel[0], nil)
		}

		if _, ok := nodes[rel[1]]; !ok {
			nodes[rel[1]] = day6.NewNode(rel[1], nil)
		}

		nodes[rel[1]].Parent = nodes[rel[0]]
	}

	ret := 0
	for _, node := range nodes {
		ret += node.Depth()
	}

	return ret

}

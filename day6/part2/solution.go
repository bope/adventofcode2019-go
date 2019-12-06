package part2

import "github.com/bope/adventofcode2019-go/day6"

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

	me := nodes["YOU"]
	santa := nodes["SAN"]

	p1 := me.Path()
	p2 := santa.Path()

	steps := me.Depth() + santa.Depth()

	for i := 0; ; i++ {
		if p1[len(p1)-i-1] != p2[len(p2)-i-1] {
			break
		}
		steps -= 2
	}

	return steps

}

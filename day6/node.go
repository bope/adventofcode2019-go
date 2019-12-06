package day6

type Node struct {
	Parent *Node
	Name   string
}

func NewNode(name string, parent *Node) *Node {
	return &Node{
		Name:   name,
		Parent: parent,
	}
}

func (n *Node) Depth() int {
	curr := n
	depth := 0
	for {
		curr = curr.Parent
		if curr == nil {
			break
		}
		depth++
	}
	return depth
}

func (n *Node) Path() []*Node {
	curr := n
	ret := make([]*Node, 0)
	for {
		if curr.Parent == nil {
			break
		}
		curr = curr.Parent
		ret = append(ret, curr)
	}
	return ret
}

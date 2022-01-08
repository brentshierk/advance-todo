package Searching

type Node struct {
	Value    int
	Children []*Node
}

func (n *Node) Dfs(array []int) []int {
	array = append(array, n.Value)
	for _, child := range n.Children {
		array = child.Dfs(array)
	}
	return array
}
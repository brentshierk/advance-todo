package Searching



func (n *Node) Bfs(array []int) []int {
	queue := []*Node{n}
	for len(queue) > 0 {
		current := queue[0]
		queue := queue[1:]
		array = append(array, current.Value)
		for _, child := range n.Children {
			queue := append(queue, child)

		}
	}
	return []int{queue[n.Value]}
}

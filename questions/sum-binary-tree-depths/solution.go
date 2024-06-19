package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func solution(root *Node) int {
	var sumDepths func(node *Node, depth int) int
	sumDepths = func(node *Node, depth int) int {
		if node == nil {
			return 0
		}
		currSumDepth := depth * (depth+1) / 2
		left := sumDepths(node.left, depth+1)
		right := sumDepths(node.right, depth+1)
		return left + right + currSumDepth
	}
	return sumDepths(root, 0)
}

func main() {}

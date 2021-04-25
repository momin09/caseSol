package Leet104

type BinTreenode2 struct {
	Val   int
	Left  *BinTreenode2
	Right *BinTreenode2
}

var md int

func maxDepth(root *BinTreenode2) int {
	if root == nil {
		return 0
	}

	md = 1

	if root.Left != nil {
		dfs(root.Left, 2)
	}

	if root.Right != nil {
		dfs(root.Right, 2)
	}

	return md
}

func dfs(node *BinTreenode2, depth int) {
	if depth > md {
		md = depth
	}

	if node.Left != nil {
		dfs(node.Left, depth+1)
	}
	if node.Right != nil {
		dfs(node.Right, depth+1)
	}
}

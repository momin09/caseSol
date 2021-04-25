package Leet100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		}
		return false
	} else if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	return isSameTree(p.Right, q.Right)
}

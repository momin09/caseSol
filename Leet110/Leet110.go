package Leet110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//높이 변수 추가
	lh := height(root.Left)
	rh := height(root.Right)

	if absDiff(lh, rh) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func absDiff(a, b int) int {
	//양 서브트리의 높이를 비교
	if a >= b {
		return a - b
	}
	return b - a
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//1을 더하는 이유는 자기 자신의 위치에 있을 때 높이가 하나 높아지기 때문임
	return 1 + maxHeight(height(root.Left), height(root.Right))
}

func maxHeight(a, b int) int {
	//양쪽 트리 높이를 비교 후 높은 값을 반환.
	if a >= b {
		return a
	}
	return b
}

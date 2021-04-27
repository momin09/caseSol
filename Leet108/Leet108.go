package Leet108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	tree := &TreeNode{}
	mid := len(nums) / 2
	another := nums

	if tree == nil {
		return tree
	}

	for len(nums) > 0 {
		val := nums[0]
		another = nums[1:]
		if val < mid {
			if tree.Left == nil {
				tree.Left = &TreeNode{Val: val}
				return tree.Left
			} else {
				return sortedArrayToBST(another)
			}
		} else {
			if tree.Right == nil {
				tree.Right = &TreeNode{Val: val}
				return tree.Right
			} else {
				return sortedArrayToBST(another)
			}
		}
	}
	return tree
}

package Baekjoon1991

import "fmt"

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func inorder(node *TreeNode) {
	if node != nil {
		inorder(node.left)
		fmt.Printf("%d", node.val)
		inorder(node.right)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type TreeNode struct {
	alph  rune
	left  *TreeNode
	right *TreeNode
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func inorder(node *TreeNode) {
	defer writer.Flush()

	var t int

	fmt.Fscan(reader, &t)

	if node != nil {
		inorder(node.left)
		fmt.Printf("%c", node.alph)
		inorder(node.right)
	}
}

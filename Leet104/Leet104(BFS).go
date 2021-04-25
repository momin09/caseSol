package Leet104

type BinTreenode1 struct {
	Val   int
	Left  *BinTreenode1
	Right *BinTreenode1
}

type NodewithDepth1 struct {
	node  *BinTreenode1
	depth int
}

func maxDepth1(root *BinTreenode1) int {
	if root == nil {
		return 0
	}

	q := make([]NodewithDepth1, 0, 1024)
	q = append(q, NodewithDepth1{root, 1})

	maxDepth := 0
	for len(q) > 0 {
		nd := q[0]
		if nd.depth > maxDepth {
			maxDepth = nd.depth
		}
		if nd.node.Left != nil {
			q = append(q, NodewithDepth1{nd.node.Left, nd.depth + 1})
		}
		if nd.node.Right != nil {
			q = append(q, NodewithDepth1{nd.node.Right, nd.depth + 1})
		}
		q = q[1:]
	}

	return maxDepth
}

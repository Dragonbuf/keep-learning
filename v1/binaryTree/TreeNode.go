package binaryTree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func DemoNode() *TreeNode {
	left := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	right := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	root := TreeNode{
		Val:   2,
		Left:  &left,
		Right: &right,
	}

	// 1 2 3
	return &root
}

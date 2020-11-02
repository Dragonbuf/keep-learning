package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树的前序遍历
// push in stack
// travel left
// when left is nil, pop then travel
// turn to right
func preorderTraversalV1(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	retVal := make([]int, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			retVal = append(retVal, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	return retVal
}

func inorderTraversal(root *TreeNode) []int {
	retVal := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		retVal = append(retVal, node.Val)
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	return retVal
}

// in stack again
func postorderTraversal(root *TreeNode) (res []int) {

	var stk []*TreeNode
	var prev *TreeNode

	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}

		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}

	}

	return
}

func levelOrder(root *TreeNode) (res [][]int) {

	if root == nil {
		return nil
	}

	q := []*TreeNode{root}

	for len(q) > 0 {

		tmp := make([]int, 0)
		var p []*TreeNode

		for i := 0; i < len(q); i++ {
			node := q[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}

			if node.Right != nil {
				p = append(p, node.Right)
			}
		}

		q = p
		res = append(res, tmp)
	}

	return
}

package binaryTree

import (
	"github.com/keep-learning/tool"
	"math"
)

//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1
func isBalanced(root *tool.TreeNode) bool {
	if root == nil {
		return true
	}

	isBalance := tool.Abs(maxDepthV1(root.Left)-maxDepthV1(root.Right)) <= 1
	return isBalanced(root.Left) && isBalanced(root.Right) && isBalance
}

func maxDepthV1(root *tool.TreeNode) int {
	if root == nil {
		return 0
	}
	return tool.Max(maxDepthV1(root.Left), maxDepthV1(root.Right)) + 1
}

//给定一个非空二叉树，返回其最大路径和。
func maxPathSum(root *tool.TreeNode) int {
	// 节点 = self.val + left + right
	// if left < 0 || right < 0  = 0
	//
	maxSum := math.MinInt32

	var maxSumFun func(*tool.TreeNode) int
	maxSumFun = func(node *tool.TreeNode) int {
		if node == nil {
			return 0
		}

		left := tool.Max(0, maxSumFun(node.Left))
		right := tool.Max(0, maxSumFun(node.Right))

		maxSum = tool.Max(maxSum, node.Val+left+right)

		return node.Val + tool.Max(left, right)
	}

	maxSumFun(root)

	return maxSum
}

//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// p q in root ,return root
// find left right
// if left != nil && right != nil , return root
// return left || right

func lowestCommonAncestorV1(root, p, q *tool.TreeNode) *tool.TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestorV1(root.Left, p, q)
	right := lowestCommonAncestorV1(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

func insertIntoBST(root *tool.TreeNode, val int) *tool.TreeNode {
	if root == nil {
		return &tool.TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

package binaryTree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	result := make([]int, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Right
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Left
	}
	return result
}

//核心就是：根节点必须在右节点弹出之后，再弹出
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var visitNode *TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == visitNode {
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			visitNode = node
		} else {
			root = node.Right
		}

	}

	return result
}

func preorderTraversalDfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	dfs(root, &result)
	return result
}

//DFS 深度搜索-从上到下
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

func divideAndConquer(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

//BFS 层次遍历
func leverOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

package binaryTree

// this is BFS

func minDepth(root *TreeNode) int {
	var step int
	// 入队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	step++
	// 记录步数

	// 队列不为空
	for len(queue) > 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				return step
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

		}

		step++
	}
	// 判断队列长度

	return step
}

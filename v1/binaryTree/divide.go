package binaryTree

import "math"

/**
  1  终止条件
  2 divide
  3  conquer


*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			} else {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, list)
	}

	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil
}

func divideAndConquerB(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	left := divideAndConquerB(root.Left)
	right := divideAndConquerB(root.Right)

	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[mid:])
	right := mergeSort(nums[:mid])

	return merge(left, right)
}

func merge(left []int, right []int) (result []int) {
	var l, r int

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}

	data[head] = mid
	QuickSort(data[:head])
	QuickSort(data[head+1:])
	return data
}

func quicksort(nums []int, start int, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quicksort(nums, 0, pivot-1)
		quicksort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start int, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			//swap
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// 把中间的值换为用于比较的基准值
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func isBalance(root *TreeNode) bool {
	return maxDepth2(root) != -1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || right-left > 1 || left-right > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}

	return right + 1
}

func maxPath(root *TreeNode) int {
	return helper(root).maxPath
}

type resultType struct {
	singlePath int
	maxPath    int
}

func helper(root *TreeNode) resultType {
	res := resultType{0, math.MinInt32}
	if root == nil {
		return res
	}

	left := helper(root.Left)
	right := helper(root.Right)

	if left.singlePath > right.singlePath {
		res.singlePath = max(left.singlePath+root.Val, 0)
	} else {
		res.singlePath = max(right.singlePath+root.Val, 0)
	}

	maxPath := max(left.maxPath, right.maxPath)
	res.maxPath = max(maxPath, left.singlePath+right.singlePath+root.Val)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

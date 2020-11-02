package sort

import "math"

// todo 非递归的方法
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		// 为什么每次从 0 开始？？ todo
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start int, end int) int {
	pivot := nums[end]
	i, j := start, start
	for ; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]
	return i
}

func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	// 找到中点
	mid := len(nums) / 2
	// 左边 递归
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	// 谁小合并谁
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	// 处理剩余的
	return result
}

func HeapSort(nums []int) []int {

	minHeap := make([]int, len(nums)+1)
	minHeap[0] = math.MinInt64
	copy(minHeap[1:], nums)

	for i := (len(minHeap) - 1) / 2; i > 0; i-- {
		down(minHeap, i)
	}

	return minHeap
}

func down(heap []int, i int) {
	temp := heap[i]

	var parent, child int
	for parent = i; parent*2 <= (len(heap) - 1); parent = child {
		child = parent * 2
		if child != (len(heap)-1) && heap[child] > heap[child+1] {
			child++
		}
		if temp < heap[child] {
			break
		}

		heap[parent] = heap[child]
	}

	heap[parent] = temp
}

package binaarySearch

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

//5, 7, 7, 8, 8, 10 => 8    [3,4]
//给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。 如果目标值不在数组中，则返回[-1, -1]
func searchRange(A []int, target int) []int {
	//两次二分查找，找到第一次和最后一次的位置

	returnData := []int{-1, -1}
	if len(A) <= 0 {
		return returnData
	}

	start := 0
	end := len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] == target {
			// 如果相等，向左查找，找到起始位置
			end = mid
		} else if A[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}

	if A[start] == target {
		returnData[0] = start
	} else if A[end] == target {
		returnData[0] = end
	} else {
		return returnData
	}

	start = 0
	end = len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] == target {
			// 如果相等，向右查找，找到结束位置
			start = mid
		} else if A[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}

	if A[start] == target {
		returnData[1] = start
	} else if A[end] == target {
		returnData[1] = end
	} else {
		return returnData
	}

	return returnData
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置
// 找到第一个 >= target 的元素位置
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	//  最后的情况  start---end   ,所以判断 target 即可
	if nums[start] > target {
		return start
	} else if nums[end] < target {
		return end + 1
	} else {
		return end
	}

}

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。

//处理成二分法即可
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	start := 0
	row := len(matrix)
	col := len(matrix[0])
	end := row*col - 1

	for start <= end {
		mid := start + (end-start)/2
		if matrix[mid/col][mid%col] == target {
			return true
		} else if matrix[mid/col][mid%col] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

//假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
//你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。
//实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数

func firstBadVersion(n int) int {
	start := 1
	end := n
	for start <= end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			return mid
		} else {
			start = mid + 1
		}
	}

	return -1
}

func isBadVersion(version int) bool {
	return version == 4
}

//假设按照升序排序的数组在预先未知的某个点上进行了旋转( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 请找出其中最小的元素。
// [3,4,5,1,2]
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	start := 0
	end := len(nums)
	target := nums[0]

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

//假设按照升序排序的数组在预先未知的某个点上进行了旋转 ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//请找出其中最小的元素。(包含重复元素)

func findMin2(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end--
		}
	}

	return nums[start]
}

//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 你可以假设数组中不存在重复的元素

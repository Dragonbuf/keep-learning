package minWindow

import "math"

//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串
func minWindow(s, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, v := range []byte(t) {
		need[v] = 1
	}

	var left, right, start, valid int
	targetLen := math.MaxInt64

	for right < len(s) {

		tmp := s[right]
		right++
		if _, ok := need[tmp]; ok {
			window[tmp]++
			if window[tmp] == need[tmp] {
				valid++
			}
		}

		for valid == len(need) {

			if right-left < targetLen {
				start = left
				targetLen = right - left
			}

			tmp := s[left]
			left++
			if _, ok := need[tmp]; ok {
				if window[tmp] == need[tmp] {
					valid--
				}
				window[tmp]--
			}
		}

	}

	if targetLen == math.MaxInt64 {
		return ""
	}

	return s[start : start+targetLen]
}

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, v := range []byte(s1) {
		need[v] = 1
	}
	var left, right, valid int

	for right < len(s2) {
		tmp := s2[right]
		right++

		if _, ok := need[tmp]; ok {
			window[tmp]++
			if window[tmp] == need[tmp] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			tmp := s2[left]
			left++
			if _, ok := need[tmp]; ok {
				if need[tmp] == window[tmp] {
					valid--
				}
				window[tmp]--
			}
		}
	}

	return false
}

//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
func findAnagrams(s string, p string) []int {
	// 处理方法和上面一样，只不过需要记录下 []int res
	return nil
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	var left, right, res int
	for right < len(s) {
		tmp := s[right]
		right++
		window[tmp]++

		for window[tmp] > 1 {
			tmp := s[left]
			left++
			window[tmp]--
		}
		res = max(res, right-left)
	}
	return res
}

func max(res int, i int) int {
	if res > i {
		return res
	}
	return i
}

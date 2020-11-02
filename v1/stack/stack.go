package stack

import (
	"github.com/keep-learning/v1/binaryTree"
	"math"
	"strconv"
	"strings"
)

type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

// 必须保证 同时 push ，不然 pop 可能有问题
func (s *MinStack) Push(x int) {
	min := s.GetMin()
	if min > x {
		s.min = append(s.min, x)
	} else {
		s.min = append(s.min, min)
	}

	s.stack = append(s.stack, x)
}

func (s *MinStack) Pop() {
	if len(s.stack) < 0 {
		return
	}

	s.stack = s.stack[:len(s.stack)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}

	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.min) < 1 {
		return math.MaxInt64
	}

	return s.min[len(s.min)-1]
}

func evalRPN(tokens []string) int {

	stack := make([]int, 0)
	var first, second int

	for _, token := range tokens {
		switch token {
		case "+":
			first, second, stack = findStackLastTwoValue(stack)
			stack = append(stack, first+second)
		case "-":
			first, second, stack = findStackLastTwoValue(stack)
			stack = append(stack, first-second)
		case "*":
			first, second, stack = findStackLastTwoValue(stack)
			stack = append(stack, first*second)
		case "/":
			first, second, stack = findStackLastTwoValue(stack)
			stack = append(stack, first/second)
		default:
			value, _ := strconv.Atoi(token)
			stack = append(stack, value)
		}
	}

	return stack[len(stack)-1]
}

func findStackLastTwoValue(stack []int) (first, second int, chuckedStack []int) {
	second = stack[len(stack)-1]
	first = stack[len(stack)-2]
	chuckedStack = stack[:len(stack)-2]
	return
}

func decodeString(s string) string {
	stack := make([]string, 0)
	ptr := 0

	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			stack = append(stack, extractNumber(s, &ptr))
		} else if cur == ']' {
			// 出栈 组成 sub
			ptr++
			var sub []string
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			// 反转 sub
			for i, j := 0, len(sub)-1; i < j; i, j = i+1, j-1 {
				sub[i], sub[j] = sub[j], sub[i]
			}

			repTime, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]

			// 反转拼接 sub ，入栈
			t := strings.Repeat(getString(sub), repTime)
			stack = append(stack, t)
		} else {
			stack = append(stack, string(cur))
			ptr++
		}
	}

	return getString(stack)
}

func getString(sub []string) string {
	tmp := ""
	for _, v := range sub {
		tmp += v
	}
	return tmp
}

func extractNumber(s string, ptr *int) string {

	tmp := ""
	if s[*ptr] <= '9' && s[*ptr] >= '0' {
		tmp += string(s[*ptr])
		*ptr++
	}

	return tmp
}

//给定一个二叉树，返回它的中序遍历。
func inorderTraversal(root *binaryTree.TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*binaryTree.TreeNode
	var result []int
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}

	return result
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func getInputNode() *Node {
	first := Node{
		Val:       1,
		Neighbors: nil,
	}

	second := Node{
		Val:       2,
		Neighbors: []*Node{&first},
	}
	first.Neighbors = []*Node{&second}

	return &second
}

func getOutputNode() *Node {
	first := Node{
		Val:       1,
		Neighbors: nil,
	}

	second := Node{
		Val:       2,
		Neighbors: []*Node{&first},
	}
	first.Neighbors = []*Node{&second}

	return &second
}

//给你无向连通图中一个节点的引用，请你返回该图的深拷贝（克隆）
func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := visited[node]; ok {
		return v
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}

	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}

	return newNode
}

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}
func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		// 标记已经访问过(每一个点只需要访问一次)
		grid[i][j] = 0
		return dfs(grid, i-1, j) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + 1
	}
	return 0
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	stack := make([]int, 0)
	max := 0

	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}

		for len(stack) != 0 && heights[stack[len(stack)-1]] >= cur {

			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			w := i
			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			max = findMax(max, h*w)

		}

		stack = append(stack, i)

	}
	return max
}

func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

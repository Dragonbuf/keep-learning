package binaryTree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
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

// inorder traversal
var answer = 0

func maximumDepth(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if root.Left == nil || root.Right == nil {
		answer = Max(answer, depth)
	}

	maximumDepth(root.Left, depth+1)
	maximumDepth(root.Right, depth+1)
}

//recur

//给定一个二叉树，检查它是否是镜像对称的。
/*
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
*/

func isSymmetricRecur(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recur(root.Left, root.Right)
}

func recur(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil || l.Val != r.Val {
		return false
	}

	return recur(l.Left, r.Right) && recur(l.Right, r.Left)
}

func isSymmetricQueue(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left)
		queue = append(queue, right.Right)
		queue = append(queue, left.Right)
		queue = append(queue, right.Left)
	}

	return recur(root.Left, root.Right)
}

//回溯问题：
var res [][]int

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	backTrack(root, []int{}, sum)
	return len(res) > 0
}

func backTrack(root *TreeNode, path []int, sum int) {

	if root == nil {
		return
	}

	path = append(path, root.Val)

	if sum == root.Val && root.Left == nil && root.Right == nil {
		var tmp []int
		copy(tmp, path)
		res = append(res, tmp)
	}

	backTrack(root.Left, path, sum-root.Val)
	backTrack(root.Right, path, sum-root.Val)
	path = path[:len(path)-1]
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	helper := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		helper[inorder[i]] = i
	}

	var build func(int, int) *TreeNode

	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		index := helper[val]

		root.Right = build(index+1, right)
		root.Left = build(left, index+1)
		return root
	}

	return build(0, len(inorder)-1)
}

func connect(root *Node) *Node {
	q := []*Node{root}

	for len(q) > 0 {
		// has left  has right append,
		p := make([]*Node, 0)
		for i := 0; i < len(q); i++ {
			if i+1 < len(q) {
				q[i].Next = q[i+1]
			}
			if q[i].Left != nil {
				p = append(p, q[i].Left)
			}

			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		p = q
	}
	return root
}

func connectV2(root *Node) *Node {
	start := root

	for start != nil {
		var nextStart, last *Node

		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}

		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}

	return root
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/**
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if(root == null || root == p || root == q) return root;
        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        if(left == null && right == null) return null; // 1.
        if(left == null) return right; // 3.
        if(right == null) return left; // 4.
        return root; // 2. if(left != null and right != null)
    }
}

*/

type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return this.rserialize(root, "")
}

func (this *Codec) rserialize(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = this.rserialize(root.Left, str)
		str = this.rserialize(root.Right, str)
	}
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		if l[i] != "" {
			this.l = append(this.l, l[i])
		}
	}
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.l[0])
	this.l = this.l[1:]

	root := &TreeNode{Val: val}
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}

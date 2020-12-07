package tool

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

type ListNode struct {
	Val  int
	Next *ListNode
}

type Codec struct {
	l []string
}

func Constructor() *Codec {
	return &Codec{}
}

//你可以将以下二叉树：
//
//    1
//   / \
//  2   3
//     / \
//    4   5
//
//序列化为 "[1,2,3,null,null,4,5]"
//1,2,3,null,null,4,5
//1,2,null,null,3,4,null,null,5,null,null
// Serializes a tree to a single string.
func (c *Codec) Serialize(root *TreeNode) string {
	c.printAll(root)
	return strings.TrimRight(strings.Join(c.l, ","), ",")
}

func (c *Codec) printAll(root *TreeNode) {
	if root == nil {
		c.l = append(c.l, "null")
		return
	}
	c.l = append(c.l, strconv.Itoa(root.Val))
	c.printAll(root.Left)
	c.printAll(root.Right)
}

// Deserializes your encoded data to tree.
func (c *Codec) Deserialize(data string) *TreeNode {

	c.l = strings.Split(data, ",")

	return c.rdeserialize()
}

func (c *Codec) rdeserialize() *TreeNode {
	if c.l[0] == "null" {
		c.l = c.l[1:]
		return nil
	}
	v, _ := strconv.Atoi(c.l[0])
	node := &TreeNode{Val: v}
	c.l = c.l[1:]
	node.Left = c.rdeserialize()
	node.Right = c.rdeserialize()
	return node
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

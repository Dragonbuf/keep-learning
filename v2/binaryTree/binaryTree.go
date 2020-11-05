package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func hasPathSum(root *TreeNode, sum int) bool {

	return true
}

/**
class Solution {
    LinkedList<List<Integer>> res = new LinkedList<>();

    LinkedList<Integer> path = new LinkedList<>();

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        recur(root, sum);
        return res;
    }

    void recur(TreeNode root, int tar) {
        if(root == null) return;
        path.add(root.val);
        tar -= root.val;
        if(tar == 0 && root.left == null && root.right == null)
            res.add(new LinkedList(path));


        recur(root.left, tar);
        recur(root.right, tar);


        path.removeLast();
    }
}

/**
func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
        return nil
    }
    var ret [][]int
    dfs(root,sum,[]int{},&ret)
    return ret
}
func dfs(root *TreeNode,sum int,arr []int,ret *[][]int){
    if root == nil{
        return
    }
    arr = append(arr,root.Val)
    if root.Val == sum && root.Left == nil && root.Right == nil {
        tmp := make([]int,len(arr))
        copy(tmp,arr)
        *ret = append(*ret,tmp)
    }
    dfs(root.Left,sum - root.Val,arr,ret)
    dfs(root.Right,sum - root.Val,arr,ret)
    arr = arr[:len(arr)-1]
}
*/

func backTrack(root *TreeNode, current, sum int) bool {

	if root == nil {
		return false
	}

	if root.Val+current <= sum {
		current += root.Val
		if current == sum {
			return true
		}
	}

	for i := 0; i < 2; i++ {

		if backTrack(root.Left, current, sum) {
			return true
		}

		if backTrack(root.Right, current, sum) {
			return true
		}

	}

	return false
}

/**
了解递归并利用递归解决问题并不容易。

当遇到树问题时，请先思考一下两个问题：

你能确定一些参数，从该节点自身解决出发寻找答案吗？
你可以使用这些参数和节点本身的值来决定什么应该是传递给它子节点的参数吗？
如果答案都是肯定的，那么请尝试使用 “自顶向下” 的递归来解决此问题。

或者你可以这样思考：
对于树中的任意一个节点，如果你知道它子节点的答案，你能计算出该节点的答案吗？
如果答案是肯定的，那么 “自底向上” 的递归可能是一个不错的解决方法。
*/

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

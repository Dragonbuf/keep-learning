package binaryTree

import (
	"reflect"
	"testing"
)

func Test_preorderTraversalV1(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	r1 := &TreeNode{
		Val: 1,
	}
	r2 := &TreeNode{
		Val: 2,
	}
	r3 := &TreeNode{
		Val: 3,
	}
	r1.Right = r2
	r2.Left = r3

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see 1 2 3",
			args{r1},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalV1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	r1 := &TreeNode{
		Val: 1,
	}
	r2 := &TreeNode{
		Val: 2,
	}
	r3 := &TreeNode{
		Val: 3,
	}
	r1.Right = r2
	r2.Left = r3

	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			"can see 1  3 2",
			args{r1},
			[]int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	r1 := &TreeNode{
		Val: 1,
	}
	r2 := &TreeNode{
		Val: 2,
	}
	r3 := &TreeNode{
		Val: 3,
	}
	r1.Right = r2
	r2.Left = r3
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			"can see  3 2 1",
			args{r1},
			[]int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := postorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("postorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	//    3
	//   / \
	//  9  20
	//    /  \
	//   15   7

	r3 := &TreeNode{
		Val: 3,
	}
	r9 := &TreeNode{
		Val: 9,
	}
	r20 := &TreeNode{
		Val: 20,
	}

	r15 := &TreeNode{
		Val: 15,
	}
	r7 := &TreeNode{
		Val: 7,
	}
	r3.Left = r9
	r3.Right = r20
	r20.Left = r15
	r20.Right = r7

	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{
			"can see 3 9 20 15 7",
			args{root: r3},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := levelOrder(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("levelOrder() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

/**
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
*/
func Test_isSymmetricRecur(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	r1 := &TreeNode{
		Val: 1,
	}
	r2l := &TreeNode{
		Val: 2,
	}
	r2r := &TreeNode{
		Val: 2,
	}

	r1.Left = r2l
	r1.Right = r2r

	r3l1 := &TreeNode{
		Val: 4,
	}
	r3r2 := &TreeNode{
		Val: 3,
	}
	r2l.Left = r3l1
	r2l.Right = r3r2

	r3r1 := &TreeNode{
		Val: 3,
	}
	r3l2 := &TreeNode{
		Val: 4,
	}

	r2r.Left = r3r1
	r2r.Right = r3l2

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see true",
			args{root: r1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetricRecur(tt.args.root); got != tt.want {
				t.Errorf("isSymmetricRecur() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymmetricQueue(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	r1 := &TreeNode{
		Val: 1,
	}
	r2l := &TreeNode{
		Val: 2,
	}
	r2r := &TreeNode{
		Val: 2,
	}

	r1.Left = r2l
	r1.Right = r2r

	r3l1 := &TreeNode{
		Val: 4,
	}
	r3r2 := &TreeNode{
		Val: 3,
	}
	r2l.Left = r3l1
	r2l.Right = r3r2

	r3r1 := &TreeNode{
		Val: 3,
	}
	r3l2 := &TreeNode{
		Val: 4,
	}

	r2r.Left = r3r1
	r2r.Right = r3l2

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see true",
			args{root: r1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetricQueue(tt.args.root); got != tt.want {
				t.Errorf("isSymmetricQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see false",
			args{
				root: nil,
				sum:  1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

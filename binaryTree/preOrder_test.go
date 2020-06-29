package binaryTree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see nil",
			args{root: nil},
			nil,
		},
		{
			"can see 213",
			args{root: DemoNode()},
			[]int{2, 1, 3},
		},
		{
			"can see 0",
			args{root: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			} else {
				for k, v := range tt.want {
					fmt.Println(k, v)
				}
			}
		})
	}
}

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see nil",
			args{root: nil},
			nil,
		},
		{
			"can see 231",
			args{root: DemoNode()},
			[]int{2, 3, 1},
		},
		{
			"can see 0",
			args{root: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			} else {
				for k, v := range tt.want {
					fmt.Println(k, v)
				}
			}
		})
	}
}

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see nil",
			args{root: nil},
			nil,
		},
		{
			"can see 132",
			args{root: DemoNode()},
			[]int{1, 3, 2},
		},
		{
			"can see 0",
			args{root: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversalDfs(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see nil",
			args{root: nil},
			nil,
		},
		{
			"can see 213",
			args{root: DemoNode()},
			[]int{2, 1, 3},
		},
		{
			"can see 0",
			args{root: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalDfs(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalDfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divideAndConquer(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see nil",
			args{root: nil},
			nil,
		},
		{
			"can see 213",
			args{root: DemoNode()},
			[]int{2, 1, 3},
		},
		{
			"can see 0",
			args{root: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideAndConquer(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideAndConquer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leverOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see nil",
			args{root: nil},
			nil,
		},
		{
			"can see 213",
			args{root: DemoNode()},
			[]int{2, 1, 3},
		},
		{
			"can see 0",
			args{root: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			}},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leverOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leverOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

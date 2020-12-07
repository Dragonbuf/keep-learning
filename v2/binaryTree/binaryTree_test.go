package binaryTree

import (
	"github.com/keep-learning/tool"
	"reflect"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *tool.TreeNode
	}
	constructor := tool.Constructor()
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see false",
			args{root: constructor.Deserialize("1,2,3,4,null,null,4,null,null,3,null,null,2,null,null")},
			false,
		},
		{
			"can see true",
			args{root: constructor.Deserialize("3,9,null,null,20,15,null,null,7,null,null")},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepthV1(t *testing.T) {
	type args struct {
		root *tool.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 3",
			args{root: tool.Constructor().Deserialize("3,9,null,null,20,15,null,null,7,null,null")},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthV1(tt.args.root); got != tt.want {
				t.Errorf("maxDepthV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *tool.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 6",
			args{root: tool.Constructor().Deserialize("1,2,null,null,3,null,null")},
			6,
		},
		{
			"can see 42",
			args{root: tool.Constructor().Deserialize("-10,9,null,null,20,15,null,null,7,null,null")},
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *tool.TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *tool.TreeNode
	}{
		{
			"can see ",
			args{
				root: tool.Constructor().Deserialize("4,2,1,null,null,3,null,null,7,null,null"),
				val:  5,
			},
			tool.Constructor().Deserialize("4,2,1,null,null,3,null,null,7,5,null,null,null"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

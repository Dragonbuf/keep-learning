package binaryTree

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"nil is nil",
			args{nums: nil},
			nil,
		}, {
			"can sort 3 4 1",
			args{nums: []int{3, 4, 1}},
			[]int{1, 3, 4},
		}, {
			"can sort 3",
			args{nums: []int{3}},
			[]int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	"nil is nil",
		//	args{nums: nil},
		//	nil,
		//},
		{
			"can sort 3 4 1",
			args{nums: []int{3, 4, 1}},
			[]int{1, 3, 4},
		},
		//{
		//	"can sort 1 3 7 11 245",
		//	args{nums: []int{7,3,11,245,1}},
		//	[]int{1,3,7,11,245},
		//},{
		//	"can sort 3",
		//	args{nums: []int{3}},
		//	[]int{3},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get 1",
			args: args{root: DemoNode()},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

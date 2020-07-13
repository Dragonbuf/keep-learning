package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see 1 2 3 4 5 6",
			args{[]int{6, 4, 3, 1, 5, 2}},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see 1 2 3 4 5 6",
			args{[]int{6, 4, 3, 1, 5, 2}},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// 因为是堆排序，所以只要成为
			//     1
			//  4     2
			// 6 5   3
			// 即可排序成功，而且最小堆中的第一个哨兵为最小的数值
			"can see 1 2 3 4 5 6",
			args{[]int{6, 4, 3, 1, 5, 2}},
			[]int{-9223372036854775808, 1, 4, 2, 6, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

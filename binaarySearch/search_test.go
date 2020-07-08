package binaarySearch

import (
	"reflect"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 4",
			args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange(t *testing.T) {
	type args struct {
		A      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"want -1,-1",
			args{
				A:      nil,
				target: 9,
			},
			[]int{-1, -1},
		},
		{
			"want 3,4",
			args{
				A:      []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			[]int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.A, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 2",
			args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			2,
		},
		{
			"can see 2",
			args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			1,
		},
		{
			"can see 2",
			args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			4,
		},
		{
			"can see 2",
			args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see true",
			args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50}},
				target: 3,
			},
			true,
		},
		{
			"can see true",
			args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50}},
				target: 13,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 4",
			args{n: 5},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 1",
			args{nums: []int{3, 4, 5, 1, 2}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMin2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 1",
			args{nums: []int{1, 3, 5}},
			1,
		},
		{
			"can see 0",
			args{nums: []int{2, 2, 2, 0, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin2(tt.args.nums); got != tt.want {
				t.Errorf("findMin2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 4",
			args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			4,
		},
		{
			"can see -1",
			args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search3(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	"can see false",
		//	args{
		//		nums:   []int{},
		//		target: 0,
		//	},
		//	false,
		//},
		//{
		//	"can see false",
		//	args{
		//		nums:   []int{},
		//		target: 0,
		//	},
		//	false,
		//},
		//{
		//	"can see false",
		//	args{
		//		nums:   []int{1,1},
		//		target: 0,
		//	},
		//	false,
		//},
		//{
		//	"can see false",
		//	args{
		//		nums:   []int{3,1,1},
		//		target: 3,
		//	},
		//	true,
		//},
		{
			"can see false",
			args{
				nums:   []int{3, 5, 1},
				target: 1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search3(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see 3",
			args{
				nums:   []int{3, 1, 1},
				target: 3,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search4() = %v, want %v", got, tt.want)
			}
		})
	}
}

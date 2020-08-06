package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"want 3 6",
			args{
				nums:   []int{5, 3, 1, 6},
				target: 9,
			},
			[]int{3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{
			"want see 1,3 and 22",
			args{
				nums:   []int{1, 3, 1, 2, 2, 3},
				target: 4,
			},
			[][]int{{1, 3}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := twoSum2(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("twoSum2() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

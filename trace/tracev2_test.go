package trace

import (
	"reflect"
	"testing"
)

func Test_permuteV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"can see track",
			args{[]int{1, 2, 3}},
			[][]int{{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteV1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

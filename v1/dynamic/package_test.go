package dynamic

import "testing"

func Test_knapsack(t *testing.T) {
	type args struct {
		N   int
		W   int
		wt  []int
		val []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"want 6",
			args{
				N:   3,
				W:   4,
				wt:  []int{2, 1, 3},
				val: []int{4, 2, 3},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack(tt.args.N, tt.args.W, tt.args.wt, tt.args.val); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

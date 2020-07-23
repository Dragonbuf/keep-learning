package dynamic

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see -1",
			args{
				coins:  []int{3, 4, 8, 5},
				amount: 10,
			},
			-1,
		},
		{
			"can see 12",
			args{
				coins:  []int{2, 3, 5, 7},
				amount: 12,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

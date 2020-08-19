package dynamic

import "testing"

func Test_changeV1(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"want 4",
			args{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeV1(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("changeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

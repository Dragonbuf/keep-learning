package dynamic

import "testing"

func Test_coinChangev1(t *testing.T) {
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
				coins:  []int{3, 4, 8, 10},
				amount: 2,
			},
			-1,
		},
		{
			"can see 2",
			args{
				coins:  []int{2, 3, 5, 7},
				amount: 12,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangev1(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChangev1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChangev2(t *testing.T) {
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
				coins:  []int{3, 4, 8, 10},
				amount: 2,
			},
			-1,
		},
		{
			"can see 2",
			args{
				coins:  []int{2, 3, 5, 7},
				amount: 12,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangev2(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChangev2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dpv3(t *testing.T) {
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
			"can see -1",
			args{
				coins:  []int{3, 4, 8, 10},
				amount: 2,
			},
			-1,
		},
		{
			"can see 2",
			args{
				coins:  []int{2, 3, 5, 7},
				amount: 12,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dpv3(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("dpv3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package buyStock

import "testing"

func Test_buy1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"want 2",
			args{[]int{2, 4, 1}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buy1(tt.args.prices); got != tt.want {
				t.Errorf("buy1() = %v, want %v", got, tt.want)
			}
		})
	}
}

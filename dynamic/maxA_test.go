package dynamic

import "testing"

func Test_maxA(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 3",
			args{n: 3},
			3,
		},
		{
			"can see 9",
			args{n: 7},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxA(tt.args.n); got != tt.want {
				t.Errorf("maxA() = %v, want %v", got, tt.want)
			}
		})
	}
}

package dynamic

import "testing"

func Test_fibv1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 701408733",
			args{n: 44},
			701408733,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibv1(tt.args.n); got != tt.want {
				t.Errorf("fibv1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibv2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 701408733",
			args{n: 44},
			701408733,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibv2(tt.args.n); got != tt.want {
				t.Errorf("fibv2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibv3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 701408733",
			args{n: 44},
			701408733,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibv3(tt.args.n); got != tt.want {
				t.Errorf("fibv3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibv4(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 701408733",
			args{n: 44},
			701408733,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibv4(tt.args.n); got != tt.want {
				t.Errorf("fibv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

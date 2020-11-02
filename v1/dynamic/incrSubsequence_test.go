package dynamic

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"want 4",
			args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"want 4",
			args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISV2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLISV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

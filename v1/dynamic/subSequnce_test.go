package dynamic

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 3",
			args{
				str1: "abcde",
				str2: "ace",
			},
			3,
		},
		{
			"can see 3",
			args{
				str1: "abc",
				str2: "abc",
			},
			3,
		},
		{
			"can see 0",
			args{
				str1: "abc",
				str2: "def",
			},
			0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 4",
			args{s: "bbbab"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}

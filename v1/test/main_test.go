package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
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
			args{[]int{100, 4, 200, 1, 3, 2}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 2",
			args{[][]int{{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0}}},
			2,
		},
		{
			"can see 0",
			args{[][]int{{0, 0}, {1, 1}, {0, 0}}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see true",
			args{[]int{2, 3, 1, 1, 4}},
			true,
		},

		{
			"can see false",
			args{[]int{3, 2, 1, 0, 4}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canJump2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 1",
			args{[]int{2, 3, 1, 1, 4}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump2(tt.args.nums); got != tt.want {
				t.Errorf("canJump2() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			"can see 4",
			args{[]int{10, 9, 2, 5, 3, 7, 101, 18}},
			4,
		},

		{
			"can see 3",
			args{[]int{10, 9, 2, 5, 3, 4}},
			3,
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

func Test_minCut(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 1",
			args{s: "aab"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"can see true",
			args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			true,
		},
		{
			"can see false",
			args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 3",
			args{
				a: "abcde",
				b: "ace",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			"can see 3",
			args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			3,
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

func Test_backPack(t *testing.T) {
	type args struct {
		m int
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 9",
			args{
				m: 10,
				A: []int{3, 4, 8, 5},
			},
			9,
		},
		{
			"can see 12",
			args{
				m: 12,
				A: []int{2, 3, 5, 7},
			},
			12,
		},
		{
			"can see 83",
			args{
				m: 90,
				A: []int{12, 3, 7, 4, 5, 13, 2, 8, 4, 7, 6, 5, 7},
			},
			83,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPack(tt.args.m, tt.args.A); got != tt.want {
				t.Errorf("backPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_revertString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"can see cba",
			args{s: []byte("abc")},
			"cba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revertString(tt.args.s); got != tt.want {
				t.Errorf("revertString() = %v, want %v", got, tt.want)
			}
		})
	}
}

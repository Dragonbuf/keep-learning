package binary

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
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
			args{[]int{1, 2, 2, 3, 3}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberInThree(t *testing.T) {
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
			args{[]int{1, 2, 2, 2, 3, 3, 3}},
			1,
		},
		{
			"can see 1",
			args{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 999}},
			999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberInThree(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberInThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see 1",
			args{[]int{1, 2, 3, 3}},
			[]int{1, 2},
		},
		{
			"can see 999 1",
			args{[]int{1, 2, 2, 3, 3, 999}},
			[]int{999, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 2",
			args{num: 3},
			2,
		},
		{
			"can see 1",
			args{num: 1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"can see 0,1,1",
			args{num: 2},
			[]int{0, 1, 1},
		},
		{
			"can see 0,1,1,2,1,2",
			args{num: 5},
			[]int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			"can see 964176192",
			args{num: 43261596},
			964176192,
		},
		{
			"can see 3221225471",
			args{num: 4294967293},
			3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 4",
			args{
				m: 5,
				n: 7,
			},
			4,
		},
		{
			"can see 5",
			args{
				m: 5,
				n: 5,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

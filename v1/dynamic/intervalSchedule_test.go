package dynamic

import "testing"

func Test_intervalSchedule(t *testing.T) {
	type args struct {
		intvs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 2",
			args{[][]int{{1, 2}, {1, 2}, {1, 2}}},
			2,
		},
		{
			"can see 1",
			args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			1,
		},
		{
			"can see 0",
			args{[][]int{{1, 2}, {2, 3}}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalSchedule(tt.args.intvs); got != tt.want {
				t.Errorf("intervalSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

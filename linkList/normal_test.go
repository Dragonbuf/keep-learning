package linkList

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates2(t *testing.T) {
	type args struct {
		head *LinkList
	}
	tests := []struct {
		name string
		args args
		want *LinkList
	}{
		{
			"dumplicates",
			args{head: DemoDuplicates()},
			&LinkList{
				Val:  3,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *LinkList
	}
	tests := []struct {
		name string
		args args
		want *LinkList
	}{
		{
			"reverse",
			args{head: DemoDuplicates()},
			DemoReverse(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *LinkList
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *LinkList
	}{
		{
			"12345, m = 2, n = 4  can see 1->4->3->2->5   ",
			args{head: DemoDuplicatesBewteen(), m: 2, n: 4},
			DemoDuplicatesBewteenReverse(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetweenOthers(t *testing.T) {
	type args struct {
		head *LinkList
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *LinkList
	}{
		{
			"12345, m = 2, n = 4  can see 1->4->3->2->5   ",
			args{head: DemoDuplicatesBewteen(), m: 2, n: 4},
			DemoDuplicatesBewteenReverse(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetweenOthers(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetweenOthers() = %v, want %v", got, tt.want)
			}
		})
	}
}

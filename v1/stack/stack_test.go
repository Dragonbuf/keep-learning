package stack

import (
	"reflect"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	type fields struct {
		min   []int
		stack []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"can see -2",
			fields{
				min:   []int{-2, -2},
				stack: []int{-2, 1},
			},
			-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				min:   tt.fields.min,
				stack: tt.fields.stack,
			}
			if got := this.GetMin(); got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Top(t *testing.T) {
	type fields struct {
		min   []int
		stack []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"can see 1",
			fields{
				min:   []int{-2, -2},
				stack: []int{-2, 1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				min:   tt.fields.min,
				stack: tt.fields.stack,
			}
			if got := this.Top(); got != tt.want {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 9",
			args{tokens: []string{"2", "1", "+", "3", "*"}},
			9,
		},
		{
			"can see 6",
			args{tokens: []string{"4", "13", "5", "/", "+"}},
			6,
		},
		{
			"can see 22",
			args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"can see aaabcbc",
			args{s: "3[a]2[bc]"},
			"aaabcbc",
		},
		{
			"can see accaccacc",
			args{s: "3[a2[c]]"},
			"accaccacc",
		},
		{
			"can see abcabccdcdcdef",
			args{s: "2[abc]3[cd]ef"},
			"abcabccdcdcdef",
		},
		{
			"can see abccdcdcdxyz",
			args{s: "abc3[cd]xyz"},
			"abccdcdcdxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"can see true",
			args{node: getInputNode()},
			getOutputNode(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	var grid [][]byte
	grid = append(grid, []byte("11110"))
	grid = append(grid, []byte("11010"))
	grid = append(grid, []byte("11000"))
	grid = append(grid, []byte("00000"))
	//[["1","1","0","0","0"],
	//["1","1","0","0","0"],
	//["0","0","1","0","0"],
	//["0","0","0","1","1"]]

	var grid3 [][]byte
	grid3 = append(grid3, []byte("11000"))
	grid3 = append(grid3, []byte("11000"))
	grid3 = append(grid3, []byte("00100"))
	grid3 = append(grid3, []byte("00011"))

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 1",
			args{grid: grid},
			1,
		},
		{
			"can see 3",
			args{grid: grid3},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"can see 10",
			args{heights: []int{2, 1, 5, 6, 2, 3}},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

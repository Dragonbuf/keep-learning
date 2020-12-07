package tool

import (
	"reflect"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	type fields struct {
		l []string
	}
	type args struct {
		root *TreeNode
	}

	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 2}
	r3 := &TreeNode{Val: 3}
	r4 := &TreeNode{Val: 4}
	r5 := &TreeNode{Val: 5}
	r1.Left = r2
	r1.Right = r3
	r3.Left = r4
	r3.Right = r5

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"want see 1,2,null,null,3,4,null,null,5,null,null",
			fields{},
			args{r1},
			"1,2,null,null,3,4,null,null,5,null,null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Codec{
				l: tt.fields.l,
			}
			if got := c.Serialize(tt.args.root); got != tt.want {
				t.Errorf("Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_Deserialize(t *testing.T) {
	type fields struct {
		l []string
	}
	type args struct {
		data string
	}
	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 2}
	r3 := &TreeNode{Val: 3}
	r4 := &TreeNode{Val: 4}
	r5 := &TreeNode{Val: 5}
	r1.Left = r2
	r1.Right = r3
	r3.Left = r4
	r3.Right = r5
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TreeNode
	}{
		{
			"can see same tree",
			fields{},
			args{data: "1,2,null,null,3,4,null,null,5,null,null"},
			r1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Codec{
				l: tt.fields.l,
			}
			if got := c.Deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

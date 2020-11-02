package queue

import (
	"testing"
)

func TestMyQueue_Pop(t *testing.T) {

	type fields struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"can see 1",
			fields{
				s1: []int{1},
				s2: nil,
			},
			1,
		},
		{
			"can see 1",
			fields{
				s1: nil,
				s2: []int{1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyQueue{
				s1: tt.fields.s1,
				s2: tt.fields.s2,
			}
			if got := this.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

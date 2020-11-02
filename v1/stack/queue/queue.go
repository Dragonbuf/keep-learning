package queue

type MyQueue struct {
	s1 []int
	s2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		this.sync()
	}
	v := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return v
}

func (this *MyQueue) sync() {
	for len(this.s1) != 0 {
		v := this.s1[len(this.s1)-1]
		this.s2 = append(this.s2, v)
		this.s1 = this.s1[:len(this.s1)-1]
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		this.sync()
	}
	v := this.s2[len(this.s2)-1]
	return v
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s2) == 0 && len(this.s1) == 0
}

func updateMatrix(matrix [][]int) [][]int {
	// 把 0 放入队列

	q := make([][]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}

	// 循环队列
	dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

	for len(q) > 0 {
		point := q[0]
		q = q[1:]
		for _, v := range dirs {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				q = append(q, []int{x, y})
			}
		}
	}

	// 找上下左右
	return matrix
}

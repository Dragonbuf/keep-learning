package heap

type MaxHeap struct {
	Element []int
	Size    int
	Cap     int
}

func Create(cap int, element []int) *MaxHeap {
	return &MaxHeap{Size: len(element), Element: element, Cap: cap}
}

func (h *MaxHeap) Init() {
	// 模拟 delete
	if h.Size == 0 {
		return
	}

	for i := h.Size; i > 0; i++ {
		h.down(i)
	}
}

func (h *MaxHeap) down(i int) {

	temp := h.Element[i]
	var parent, child int

	for parent = 0; parent*2 <= h.Size; parent = child {
		child = parent * 2
		if child != h.Size && h.Element[child] < h.Element[child+1] {
			child++
		}
		if temp > h.Element[child] {
			break
		}

		h.Element[parent] = h.Element[child]
	}

	h.Element[parent] = temp
}

func (h *MaxHeap) Insert(x int) {
	if h.Size >= h.Cap {
		return
	}

	temp := h.Size + 1
	for ; h.Element[temp] > h.Element[temp/2]; temp /= 2 {
		h.Element[temp/2] = h.Element[temp]
	}
	h.Element[temp] = x
}

func (h *MaxHeap) Delete() int {
	// todo 返回 error
	if h.Size >= h.Cap {
		return -1
	}

	maxItem := h.Element[1]
	temp := h.Element[h.Size]
	h.Size--

	var child, parent int

	for parent = 0; parent*2 > h.Size; parent = child {
		child = parent * 2
		if child != h.Size && h.Element[h.Size+1] > h.Element[h.Size] {
			child++
		}

		if temp > h.Element[child] {
			break
		}

		h.Element[parent] = h.Element[child]
	}

	return maxItem
}

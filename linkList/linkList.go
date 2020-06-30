package linkList

type LinkList struct {
	Val  int
	Next *LinkList
}

func DemoDuplicates() *LinkList {

	tail := &LinkList{
		Val:  3,
		Next: nil,
	}

	mid := &LinkList{
		Val:  1,
		Next: tail,
	}

	head := &LinkList{
		Val:  1,
		Next: mid,
	}

	return head
}

func DemoDuplicatesBewteen() *LinkList {
	five := &LinkList{
		Val:  5,
		Next: nil,
	}
	four := &LinkList{
		Val:  4,
		Next: five,
	}
	three := &LinkList{
		Val:  3,
		Next: four,
	}
	two := &LinkList{
		Val:  2,
		Next: three,
	}

	one := &LinkList{
		Val:  1,
		Next: two,
	}

	return one
}

// 1->4->3->2->5    m = 2, n = 4
func DemoDuplicatesBewteenReverse() *LinkList {

	five := &LinkList{
		Val:  5,
		Next: nil,
	}

	two := &LinkList{
		Val:  2,
		Next: five,
	}

	three := &LinkList{
		Val:  3,
		Next: two,
	}
	four := &LinkList{
		Val:  4,
		Next: three,
	}

	one := &LinkList{
		Val:  1,
		Next: four,
	}

	return one
}

func DemoReverse() *LinkList {

	tail := &LinkList{
		Val:  1,
		Next: nil,
	}

	mid := &LinkList{
		Val:  1,
		Next: tail,
	}

	head := &LinkList{
		Val:  3,
		Next: mid,
	}

	return head
}

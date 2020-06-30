package linkList

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
func deleteDuplicates(head *LinkList) *LinkList {
	current := head
	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}

	return head
}

func deleteDuplicates2(head *LinkList) *LinkList {
	if head == nil {
		return nil
	}

	tmp := &LinkList{}
	tmp.Next = head
	head = tmp

	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal := head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}

		} else {
			head = head.Next
		}
	}

	return tmp.Next
}

func reverseList(head *LinkList) *LinkList {
	if head == nil {
		return nil
	}

	var pre *LinkList
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 1. 头插法反转链表, 每次将遍历的当前节点插入开始反转的位置。
 * 2. 通过哨兵节点处理 m == 1 的情况
 * 3. 通过m的值确定头插法的头节点位置
 * 4. 通过n-m的值确定执行几次头插操作
 */
func reverseBetweenOthers(head *LinkList, m int, n int) *LinkList {
	dummy, j := &LinkList{Next: head}, n-m
	d := dummy
	for m > 1 {
		d = d.Next
		m--
	}
	cur := d.Next.Next
	pre := d.Next
	for j > 0 {
		pre.Next = cur.Next
		cur.Next = d.Next
		d.Next = cur
		cur = pre.Next
		j--
	}
	return dummy.Next
}

// 反转 m n 之间的链表
func reverseBetween(head *LinkList, m int, n int) *LinkList {
	left := head
	right := head
	reverse(right, left, m, n)
	return head
}

func reverse(right *LinkList, left *LinkList, m, n int) {
	if n == 1 {
		return
	}

	right = right.Next

	if m > 1 {
		left = left.Next
	}

	reverse(right, left, m-1, n-1)

	if !(left == right || right.Next == left) {
		t := left.Val
		left.Val = right.Val
		right.Val = t
		left = left.Next
	}
}

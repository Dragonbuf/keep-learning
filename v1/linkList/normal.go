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

func mergeTwoLists(l1 *LinkList, l2 *LinkList) *LinkList {

	preHead := &LinkList{Val: -1}
	pre := preHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}

	return preHead.Next
}

func partition(head *LinkList, x int) *LinkList {
	beforeHead := &LinkList{}
	before := beforeHead
	afterHead := &LinkList{}
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}

		head = head.Next
	}

	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

func sortList(head *LinkList) *LinkList {
	if head == nil || head.Next == nil {
		return head
	}
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil
	left := sortList(tail)
	right := sortList(head)

	return mergeTwoLists2(left, right)
}

func mergeTwoLists2(left *LinkList, right *LinkList) *LinkList {
	tmp := &LinkList{}
	head := tmp.Next
	for left != nil && right != nil {
		if left.Val > right.Val {
			head.Next = right
			right = right.Next
		} else {
			head.Next = left
			left = left.Next
		}
		head = head.Next
	}

	if left != nil {
		head.Next = left
		head = head.Next
		left = left.Next
	}

	if right != nil {
		head.Next = right
		head = head.Next
		right = right.Next
	}

	return tmp.Next
}

func findMiddle(head *LinkList) *LinkList {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

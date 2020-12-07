package linklist

import "github.com/keep-learning/tool"

func deleteDuplicates(head *tool.ListNode) *tool.ListNode {
	p := head
	for p != nil {
		if p.Next != nil && p.Next.Val == p.Val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}

func deleteDuplicatesV2(head *tool.ListNode) *tool.ListNode {
	pre := &tool.ListNode{Next: head}

	for pre.Next != nil {
		if pre.Next.Next != nil && pre.Next.Val == pre.Next.Val {
			pre.Next.Next = pre.Next.Next.Next
		}
		pre.Next = pre.Next.Next
	}

	return pre.Next
}

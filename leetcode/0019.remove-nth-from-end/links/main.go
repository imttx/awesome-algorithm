package main

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func reverce(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next = prev
		prev = cur
		cur = cur.Next
	}

	return prev
}

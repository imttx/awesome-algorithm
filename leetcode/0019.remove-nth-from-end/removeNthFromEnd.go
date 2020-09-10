package leetcode

import "github.com/imttx/awesome-algorithm/structures"

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := new(ListNode)
	result.Next = head

	var pre, cur *ListNode
	cur = result

	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}

	pre.Next = cur.Next
	return result.Next
}

package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 25. 合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

限制：
0 <= 链表长度 <= 1000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structures.ListNode

// 解法一：双指针法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	l := h
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}

	if l1 != nil {
		l.Next = l1
	}

	if l2 != nil {
		l.Next = l2
	}

	return h.Next
}

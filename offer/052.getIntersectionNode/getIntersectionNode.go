package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 52. 两个链表的第一个公共节点
输入两个链表，找出它们的第一个公共节点。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

// 双指针法
// 解题思路：
// 我们使用两个指针 node1，node2 分别指向两个链表 headA，headB 的头结点，然后同时分别逐结点遍历，当 node1 到达链表 headA 的末尾时，重新定位到链表 headB 的头结点；当 node2 到达链表 headB 的末尾时，重新定位到链表 headA 的头结点。
// 这样，当它们相遇时，所指向的结点就是第一个公共结点。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	node1, node2 := headA, headB
	count := 0

	for node1 != node2 {
		node1 = node1.Next
		if node1 == nil {
			count++
			node1 = headB
		}

		node2 = node2.Next
		if node2 == nil {
			count++
			node2 = headA
		}

		if count > 2 {
			return nil
		}
	}
	return node1
}

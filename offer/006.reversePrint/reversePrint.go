package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：
输入：head = [1,3,2]
输出：[2,3,1]

限制：
0 <= 链表长度 <= 10000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

// 解法一：栈
func reversePrint(head *ListNode) []int {
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	for i, len := 0, len(stack); i < len/2; i++ {
		stack[i], stack[len-1-i] = stack[len-1-i], stack[i]
	}
	return stack
}

// 解法二：递归法
func reversePrintV2(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrintV2(head.Next), head.Val)
}

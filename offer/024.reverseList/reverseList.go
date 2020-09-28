package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 24. 反转链表
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

限制：
0 <= 节点个数 <= 5000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

// 解法一：双指针
// 定义两个指针： pre 和 cur，pre 在前 cur 在后。
// 每次让 pre 的 next 指向 cur ，实现一次局部反转
// 局部反转完成之后， pre 和 cur 同时往前移动一个位置
// 循环上述过程，直至 pre 到达链表尾部

func reverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	pre = head

	// 注意：操作顺序很重要！！！
	for pre != nil {
		next = pre.Next // 提前保存正常的下一个节点
		pre.Next = cur  // 节点局部反转
		cur = pre       // cur 向前移动
		pre = next      // pre 向前移动
	}

	return cur
}

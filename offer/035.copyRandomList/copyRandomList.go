package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 35. 复杂链表的复制
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node = structures.ListNodeWithRandom

// 算法：深度优先搜索
// 从头结点 head 开始拷贝；
// 由于一个结点可能被多个指针指到，因此如果该结点已被拷贝，则不需要重复拷贝；
// 如果还没拷贝该结点，则创建一个新的结点进行拷贝，并将拷贝过的结点保存在哈希表中；
// 使用递归拷贝所有的 next 结点，再递归拷贝所有的 random 结点。

// var nodeHash map[*Node]*Node

// func copyRandomList(head *Node) *Node {
// 	nodeHash = make(map[*Node]*Node)
// 	return dfs(head)
// }

// func dfs(head *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}

// 	if val, ok := nodeHash[head]; ok {
// 		return val
// 	}

// 	newNode := &Node{Val: head.Val}
// 	newNode.Next = dfs(head.Next)
// 	newNode.Random = dfs(head.Random)
// 	return newNode
// }

// 解法一：哈希表实现
func copyRandomList(head *Node) *Node {
	// 创建哈希表
	nodeHash := make(map[*Node]*Node)

	// 顺序遍历，存储老结点和新结点(先存储新创建的结点值)
	cur := head
	for cur != nil {
		nodeHash[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	// 复制新节点的 Next 和 Random 指向
	cur = head
	for cur != nil {
		newNode := nodeHash[cur]
		newNode.Next = nodeHash[cur.Next]     // 新结点next指向同旧结点的next指向
		newNode.Random = nodeHash[cur.Random] // 新结点random指向同旧结点的random指向
		cur = cur.Next
	}

	// 返回复制的链表
	return nodeHash[head]
}

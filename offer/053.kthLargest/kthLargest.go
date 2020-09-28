package offer

import (
	"container/list"

	"github.com/imttx/awesome-algorithm/structures"
)

/*
剑指 Offer 54. 二叉搜索树的第k大节点
给定一棵二叉搜索树，请找出其中第k大的节点。

示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4

示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4

限制：
1 ≤ k ≤ 二叉搜索树元素个数
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
本文解法基于此性质：二叉搜索树的中序遍历为 递增序列 。
中序遍历：左、根、右

根据以上性质，易得二叉搜索树的 中序遍历倒序 为 递减序列 。
因此，求 “二叉搜索树第 k 大的节点” 可转化为求 “此树的中序遍历倒序的第 k 个节点”。
*/

type TreeNode = structures.TreeNode

// 中序倒序遍历
func kthLargest(root *TreeNode, k int) int {
	stack := list.New()
	insert(root, stack, k)
	return stack.Back().Value.(int)
}

func insert(root *TreeNode, stack *list.List, k int) {
	if root == nil {
		return
	}

	insert(root.Right, stack, k)
	if stack.Len() < k {
		stack.PushBack(root.Val)
	}
	insert(root.Left, stack, k)
}

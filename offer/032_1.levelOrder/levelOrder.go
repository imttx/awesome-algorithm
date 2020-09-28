package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 32 - I. 从上到下打印二叉树
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回：
[3,9,20,15,7]

提示：
节点总数 <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structures.TreeNode

// 解法一：广度优先搜索（BFS） + 辅助队列
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := make([]int, 0)

	for len(queue) > 0 {
		var tmp []*TreeNode
		for _, node := range queue {
			result = append(result, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}

	return result
}

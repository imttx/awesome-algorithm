package offer

import (
	"github.com/imttx/awesome-algorithm/structures"
)

/*
剑指 Offer 32 - III. 从上到下打印二叉树 III
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [20,9],
  [15,7]
]

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

// 解法一：广度优先算法（BFS） + 辅助队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var tmp []*TreeNode
		var res []int
		for _, node := range queue {
			switch len(result) & 1 {
			case 0: // 偶数行
				res = append(res, node.Val)
			case 1: // 奇数行
				res = append([]int{node.Val}, res...)
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
		result = append(result, res)
	}

	return result
}

package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 55 - I. 二叉树的深度
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

提示：
节点总数 <= 10000
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
树的遍历方式总体分为两类：深度优先搜索（DFS）、广度优先搜索（BFS）；

常见的 DFS ： 先序遍历、中序遍历、后序遍历；
常见的 BFS ： 层序遍历（即按层遍历）；
*/

type TreeNode = structures.TreeNode

// 解法一：层序遍历（BFS）
// 树的层序遍历 / 广度优先搜索往往利用 队列 实现。
// 关键点： 每遍历一层，则计数器 +1 ，直到遍历完成，则可得到树的深度。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	count := 0

	for len(queue) > 0 {
		var tmp []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}

			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		queue = tmp
		count++
	}

	return count
}

// 方法一：后序遍历（DFS）
// 算法解析：
// 终止条件： 当 root​ 为空，说明已越过叶节点，因此返回 深度 00 。
// 递推工作： 本质上是对树做后序遍历。
// 计算节点 root​ 的 左子树的深度 ，即调用 maxDepth(root.left)；
// 计算节点 root​ 的 右子树的深度 ，即调用 maxDepth(root.right)；
// 返回值： 返回 此树的深度 ，即 max(maxDepth(root.left), maxDepth(root.right)) + 1。
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth2(root.Left)
	rightMax := maxDepth2(root.Right)

	if leftMax > rightMax {
		return leftMax + 1
	}
	return rightMax + 1
}

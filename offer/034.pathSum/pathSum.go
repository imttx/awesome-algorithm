package offer

import (
	"github.com/imttx/awesome-algorithm/structures"
)

/*
剑指 Offer 34. 二叉树中和为某一值的路径
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]

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

type TreeNode = structures.TreeNode

// 回溯法
var path []int
var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	path = make([]int, 0)
	res = make([][]int, 0)

	recur(root, sum)
	return res
}

func recur(root *TreeNode, target int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	target -= root.Val

	if target == 0 && root.Left == nil && root.Right == nil {
		res = append(res, append(path[:0:0], path...))
	}

	recur(root.Left, target)
	recur(root.Right, target)
	path = path[:len(path)-1]
}

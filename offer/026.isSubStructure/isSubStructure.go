package offer

import "github.com/imttx/awesome-algorithm/structures"

/*
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：
输入：A = [1,2,3], B = [3,1]
输出：false

示例 2：
输入：A = [3,4,5,1,2], B = [4,1]
输出：true

限制：
0 <= 节点个数 <= 10000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解题思路：
// 若树 B 是树 A 的子结构，则子结构的根节点可能为树 A 的任意一个节点。因此，判断树 B 是否是树 A 的子结构，需完成以下两步工作：
// 	1. 先序遍历树 A 中的每个节点（对应函数 isSubStructure(A, B)）
// 	2. 判断树 A 中的子树是否包含树 B 。（对应函数 recur(A, B)）

type TreeNode = structures.TreeNode

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return A != nil && B != nil && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

// 递归遍历a b 的每个子节点是否相等
func recur(a, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil || a.Val != b.Val {
		return false
	}

	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}

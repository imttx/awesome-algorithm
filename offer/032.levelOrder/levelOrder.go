package offer

import (
	"container/list"

	"github.com/imttx/awesome-algorithm/structures"
)

/*
剑指 Offer 32 - II. 从上到下打印二叉树 II
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

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
  [9,20],
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

/*
I. 按层打印： 题目要求的二叉树的 从上至下 打印（即按层打印），又称为二叉树的 广度优先搜索（BFS）。BFS 通常借助 队列 的先入先出特性来实现。
II. 每层打印到一行： 将本层全部节点打印到一行，并将下一层全部节点加入队列，以此类推，即可分为多行打印。

算法流程：
1. 特例处理： 当根节点为空，则返回空列表 [] ；
2. 初始化： 打印结果列表 res = [] ，包含根节点的队列 queue = [root] ；
3. BFS 循环： 当队列 queue 为空时跳出；
	3.1 新建一个临时列表 tmp ，用于存储当前层打印结果；
	3.2 当前层打印循环： 循环次数为当前层节点数（即队列 queue 长度）；
		出队： 队首元素出队，记为 node；
		打印： 将 node.val 添加至 tmp 尾部；
 	3.3 添加子节点： 若 node 的左（右）子节点不为空，则将左（右）子节点加入队列 queue ；
	3.4 将当前层结果 tmp 添加入 res 。
4. 返回值： 返回打印结果列表 res 即可。
*/

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		tmp := make([]int, 0)
		num := queue.Len()
		for i := 0; i < num; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}

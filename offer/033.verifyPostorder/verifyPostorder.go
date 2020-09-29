package offer

/*
剑指 Offer 33. 二叉搜索树的后序遍历序列
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3

示例 1：
输入: [1,6,3,2,5]
输出: false

示例 2：
输入: [1,3,2,6,5]
输出: true
*/

// 解法一：递归分治
// 解题思路：
// 后序遍历定义： [ 左子树 | 右子树 | 根节点 ] ，即遍历顺序为 “左、右、根” 。
// 二叉搜索树定义：
// 	1. 左子树中所有节点的值 << 根节点的值；
//  2. 右子树中所有节点的值 >> 根节点的值；
//  3. 其左、右子树也分别为二叉搜索树。

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}

	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p

	for postorder[p] > postorder[j] {
		p++
	}

	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}

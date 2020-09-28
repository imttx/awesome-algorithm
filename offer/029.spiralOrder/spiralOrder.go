package offer

/*
剑指 Offer 29. 顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

限制：
0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/

/*
算法流程：
1. 空值处理： 当 matrix 为空时，直接返回空列表 [] 即可。
2. 初始化： 矩阵 左、右、上、下 四个边界 l , r , t , b ，用于打印的结果列表 res 。
3. 循环打印： “从左向右、从上向下、从右向左、从下向上” 四个方向循环，每个方向打印中做以下三件事 （各方向的具体信息见下表） ；
	3.1 根据边界打印，即将元素按顺序添加至列表 res 尾部；
	3.2 边界向内收缩 1 （代表已被打印）；
	3.3 判断是否打印完毕（边界是否相遇），若打印完毕则跳出。
4. 返回值： 返回 res 即可。

打印方向	1. 根据边界打印	2. 边界向内收缩	3. 是否打印完毕
从左向右	左边界 l ，右边界 r	上边界 t 加 1	是否 t > b
从上向下	上边界 t ，下边界 b	右边界 r 减 1	是否 l > r
从右向左	右边界 r ，左边界 l	下边界 b 减 1	是否 t > b
从下向上	下边界 b ，上边界 t	左边界 l 加 1	是否 l > r
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1

	res := make([]int, 0, len(matrix)*len(matrix[0]))
	for {
		// 从左向右
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t > b {
			break
		}

		// 从上向下
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if l > r {
			break
		}

		// 从右向左
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		if b < t {
			break
		}

		// 从下向上
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}

	return res
}

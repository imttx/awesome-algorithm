package offer

/*
剑指 Offer 64. 求1+2+…+n
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

示例 1：
输入: n = 3
输出: 6

示例 2：
输入: n = 9
输出: 45

限制：
1 <= n <= 10000
*/

/*
解题思路：
本题在简单问题上做了许多限制，需要使用排除法一步步导向答案。
1+2+...+(n-1)+n 的计算方法主要有三种：平均计算、迭代、递归。

方法一： 平均计算
问题： 此计算必须使用 乘除法 ，因此本方法不可取，直接排除。

方法二： 迭代
问题： 循环必须使用 while 或 for ，因此本方法不可取，直接排除。

方法三： 递归
问题： 终止条件需要使用 if ，因此本方法不可取。
思考： 除了 if 和 switch 等判断语句外，是否有其他方法可用来终止递归
*/

// 逻辑运算符的短路计算法
func sumNums(n int) int {
	var res int
	var sum func(n int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return res
}

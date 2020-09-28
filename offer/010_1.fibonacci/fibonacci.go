package offer

/*
剑指 Offer 10- I. 斐波那契数列
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：

输入：n = 2
输出：1
示例 2：

输入：n = 5
输出：5

提示：
0 <= n <= 100
*/

// 动态规划： f(n) = f(n-1) + f(n-2)
// 复杂度分析：
// 时间复杂度 O(N)O(N) ： 计算 f(n)f(n) 需循环 nn 次，每轮循环内计算操作使用 O(1)O(1) 。
// 空间复杂度 O(1)O(1) ： 几个标志变量使用常数大小的额外空间。

const maxVal = 1000000007

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	prepre := 0
	pre := 1

	for i := 2; i <= n; i++ {
		cur := (pre + prepre) % maxVal
		prepre = pre
		pre = cur
	}

	return pre
}

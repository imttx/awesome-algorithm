package offer

import "math"

/*
剑指 Offer 14- I. 剪绳子
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36

提示：
2 <= n <= 58
*/

// 解题思路：
// 	① 当所有绳段长度相等时，乘积最大。
// 	② 最优的绳段长度为 3 。
//
// 切分规则：
// 最优： 3 。把绳子尽可能切为多个长度为 3 的片段，留下的最后一段绳子的长度可能为 0,1,2 三种情况。
// 次优： 2 。若最后一段绳子长度为 2 ；则保留，不再拆为 1+1 。
// 最差： 1 。若最后一段绳子长度为 1 ；则应把一份 3+1 替换为 2+2，因为 2×2>3×1。
// 算法流程：
// 当 n≤3 时，按照规则应不切分，但由于题目要求必须剪成 m>1 段，因此必须剪出一段长度为 1 的绳子，即返回 n−1 。
// 当 n>3 时，求 n 除以 3 的 整数部分 a 和 余数部分 b （即 n=3a+b ），并分为以下三种情况：
// 		1. 当 b = 0 时，直接返回 3^a
// 		2. 当 b = 1 时，要将一个 1+3 转换为 2+2，因此返回 3^(a-1) * 4
// 		3. 当 b = 2 时，返回 3^a * 2
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b := n/3, n%3
	switch b {
	case 0:
		return int(math.Pow(float64(3), float64(a)))
	case 1:
		return int(math.Pow(float64(3), float64(a-1))) * 4
	case 2:
		return int(math.Pow(float64(3), float64(a))) * 2
	}
	return 0
}

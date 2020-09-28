package offer

/*
剑指 Offer 62. 圆圈中最后剩下的数字
0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

示例 1：
输入: n = 5, m = 3
输出: 3

示例 2：
输入: n = 10, m = 17
输出: 2

限制：
1 <= n <= 10^5
1 <= m <= 10^6
*/

// 约瑟夫环问题
// 参考解法：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solution/huan-ge-jiao-du-ju-li-jie-jue-yue-se-fu-huan-by-as/
func lastRemaining(n int, m int) int {
	pos := 0 // 最终活下来的位置
	for i := 2; i <= n; i++ {
		pos = (pos + m) % i // 每次循环右移
	}
	return pos
}

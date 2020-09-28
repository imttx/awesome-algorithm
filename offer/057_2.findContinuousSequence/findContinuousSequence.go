package offer

/*
剑指 Offer 57 - II. 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]

示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：
1 <= target <= 10^5
*/

// 解法一：滑动窗口
// 如何用滑动窗口解这道题
// 要用滑动窗口解这道题，我们要回答两个问题：

// 第一个问题，窗口何时扩大，何时缩小？
// 第二个问题，滑动窗口能找到全部的解吗？
// 对于第一个问题，回答非常简单：

// 当窗口的和小于 target 的时候，窗口的和需要增加，所以要扩大窗口，窗口的右边界向右移动
// 当窗口的和大于 target 的时候，窗口的和需要减少，所以要缩小窗口，窗口的左边界向右移动
// 当窗口的和恰好等于 target 的时候，我们需要记录此时的结果。设此时的窗口为 [i, j)，那么我们已经找到了一个 i 开头的序列，也是唯一一个 i 开头的序列，接下来需要找 i+1 开头的序列，所以窗口的左边界要向右移动
func findContinuousSequence(target int) [][]int {
	var res [][]int
	sum := 0 // 数字和
	i := 1   // 左边界
	j := 1   // 右边界

	for i <= target/2 {
		if sum < target { // 右边界右移
			sum += j
			j++
		} else if sum > target { // 左边界右移
			sum -= i
			i++
		} else {
			// 记录结果
			var tmp []int
			for ii := i; ii < j; ii++ {
				tmp = append(tmp, ii)
			}
			// 左边界继续右移
			res = append(res, tmp)
			sum -= i
			i++
		}
	}

	return res
}

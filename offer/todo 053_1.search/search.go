package offer

/*
剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2

示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

限制：
0 <= 数组长度 <= 50000
*/

// 解法一：二分查找法
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] > target {
			r = m
		} else if nums[m] < target {
			l = m
		} else {

		}
	}
}

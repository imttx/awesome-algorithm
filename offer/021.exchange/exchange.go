package offer

/*
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

示例：
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。


提示：
1 <= nums.length <= 50000
1 <= nums[i] <= 10000
*/

// 解法一：双指针
func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]&1 == 1 { // 奇数，继续右移
			l++
			continue
		}

		if nums[r]&1 == 0 { // 偶数，继续左移
			r--
			continue
		}

		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	return nums
}

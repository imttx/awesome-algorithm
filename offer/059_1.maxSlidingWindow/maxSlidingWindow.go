package offer

/*
剑指 Offer 59 - I. 滑动窗口的最大值
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
*/

// 解法一：单调队列
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	res := make([]int, 0, len(nums)-k+1)

	for l, r := 1-k, 0; r < len(nums); l++ {
		if l > 0 && queue[0] == nums[l-1] {
			queue = queue[1:] // 删除队首对应的nums[l-1]
		}

		for len(queue) > 0 && queue[len(queue)-1] < nums[r] {
			queue = queue[:len(queue)-1] // 保持队列单调递减
		}
		queue = append(queue, nums[r])

		if l >= 0 {
			res = append(res, queue[0]) // 记录当前窗口最大值
		}

		r++
	}

	return res
}

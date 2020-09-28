package offer

/*
剑指 Offer 03. 数组中重复的数字

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

限制：
2 <= n <= 100000
*/

// 解法一： 遍历 + 哈希表
func findRepeatNumberV1(nums []int) int {
	numsHash := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := numsHash[num]; ok {
			return num
		}
		numsHash[num] = struct{}{}
	}
	return -1
}

// 解法二：本地置换
func findRepeatNumberV2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			tmp := nums[i]
			nums[i] = nums[tmp]
			nums[tmp] = tmp
		}
	}
	return -1
}

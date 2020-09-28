package offer

/*
剑指 Offer 57. 和为s的两个数字
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]

示例 2：
输入：nums = [10,26,30,31,47,60], target = 40
输出：[10,30] 或者 [30,10]

限制：
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^6
*/

// 解法一：双指针
func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] > target {
			r--
		} else if nums[l]+nums[r] < target {
			l++
		} else {
			return []int{nums[l], nums[r]}
		}
	}

	return nil
}

// 解法二：哈希表
func twoSum2(nums []int, target int) []int {
	numsHash := make(map[int]int, len(nums))
	for k, v := range nums {
		if _, ok := numsHash[target-v]; ok {
			return []int{target - v, v}
		}

		numsHash[v] = k
	}
	return nil
}

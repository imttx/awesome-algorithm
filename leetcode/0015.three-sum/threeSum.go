package leetcode

import "sort"

/*
15. 三数之和

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

// 排序 + 双指针
func threeSum(nums []int) [][]int {
	numsLen := len(nums)

	// 特判，对于数组长度小于 3，返回 nil
	if numsLen < 3 {
		return nil
	}

	// 排序
	sort.Ints(nums)

	var res [][]int

	// 遍历排序后的数组
	for i := range nums {
		// 如果 nums[i] > 0，并且已经排好序，说明后面的元素都大于0，三数之和不可能为0
		if nums[i] > 0 {
			return res
		}

		// 过滤重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := numsLen - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

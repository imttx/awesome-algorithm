package offer

import "sort"

/*
剑指 Offer 61. 扑克牌中的顺子
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

示例 1:
输入: [1,2,3,4,5]
输出: True

示例 2:
输入: [0,0,1,2,5]
输出: True

限制：
数组长度为 5
数组的数取值为 [0, 13] .
*/

/*
解题思路：
根据题意，此 5 张牌是顺子的 充分条件 如下：
	1. 除大小王外，所有牌 无重复 ；
	2. 设此 5 张牌中最大的牌为 max ，最小的牌为 min （大小王除外），则需满足：
		max - min < 5
*/

// 解法一：哈希去重 + 遍历
func isStraight(nums []int) bool {
	numsHash := make(map[int]struct{}, 5)
	min, max := 14, -1
	for _, v := range nums {
		// 跳过大小王
		if v == 0 {
			continue
		}

		// 若有重复，提前返回 false
		if _, ok := numsHash[v]; ok {
			return false
		}
		numsHash[v] = struct{}{}

		// 最大牌
		if v > max {
			max = v
		}

		// 最小牌
		if v < min {
			min = v
		}
	}

	return max-min < 5
}

// 方法二：排序 + 遍历
// 先对数组执行排序。
// 判别重复： 排序数组中的相同元素位置相邻，因此可通过遍历数组，判断 nums[i]=nums[i+1] 是否成立来判重。
// 获取最大 / 最小的牌： 排序后，数组末位元素 nums[4] 为最大牌；元素 nums[joker] 为最小牌，其中 joker 为大小王的数量。
func isStraight2(nums []int) bool {
	// 排序
	sort.Ints(nums)

	// 统计大小王数量
	joker := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			joker++
			continue
		}

		if nums[i] == nums[i+1] {
			return false
		}
	}

	return nums[4]-nums[joker] < 5
}

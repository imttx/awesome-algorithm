package offer

/*
剑指 Offer 11. 旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：
输入：[3,4,5,1,2]
输出：1

示例 2：
输入：[2,2,2,0,1]
输出：0
*/

// 解法一：暴力法
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	min := numbers[0]
	for _, val := range numbers {
		if val < min {
			min = val
		}
	}
	return min
}

// 解法二：二分法
// 特性：左排序数组的人与元素 >= 右排序的数组元素

/*
寻找旋转数组的最小元素即为寻找右排序数组的首个元素 nums[x] ，称 x 为 旋转点 。
排序数组的查找问题首先考虑使用 二分法 解决，其可将 遍历法 的 线性级别 时间复杂度降低至 对数级别 。

算法流程：
初始化： 声明 i, j 双指针分别指向 nums 数组左右两端；
循环二分： 设 m=(i+j)/2 为每次二分的中点（ "/" 代表向下取整除法，因此恒有 i≤m<j ），可分为以下三种情况：
当 nums[m] > nums[j] 时： m 一定在 左排序数组 中，即旋转点 x 一定在 [m + 1, j] 闭区间内，因此执行 i = m + 1；
当 nums[m] < nums[j] 时： m 一定在 右排序数组 中，即旋转点 x 一定在[i, m] 闭区间内，因此执行 j = m；
当 nums[m] = nums[j] 时： 无法判断 m 在哪个排序数组中，即无法判断旋转点 x 在 [i, m] 还是 [m + 1, j] 区间中。解决方案： 执行 j = j - 1 缩小判断范围，分析见下文。
返回值： 当 i = j 时跳出二分循环，并返回 旋转点的值 nums[i] 即可。
*/
func minArray2(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		m := (i + j) / 2
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] {
			j = m
		} else {
			j--
		}
	}
	return numbers[j]
}

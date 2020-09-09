package leetcode

/*
题目189: 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

通过观察我们可以得到，我们要得到最终的结果。我们只需要将所有元素反转，然后反转前 k 个元素，再反转后面l-k个元素，就能得到想要的结果。
*/
func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

// 反转数组
func reverse(nums []int) {
	numsLen := len(nums)
	for i := 0; i < numsLen/2; i++ {
		nums[i], nums[numsLen-i-1] = nums[numsLen-i-1], nums[i]
	}
}

// 使用额外数组
func rotate2(nums []int, k int) {
	tmpNums := make([]int, len(nums))
	numsLen := len(nums)
	for i, v := range nums {
		tmpNums[(i+k)%numsLen] = v
	}

	copy(nums, tmpNums)
}

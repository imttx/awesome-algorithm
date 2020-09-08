package leetcode

func maxSubArray(nums []int) int {
	if l := len(nums); l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

package leetcode

// 方法1：双指针
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

/**
方法2：双指针 —— 当要删除的元素很少时

思考：使用“元素顺序可以更改”这一属性

当我们遇到 nums[i] = val 时，我们可以将当前元素与最后一个元素进行交换，并释放最后一个元素。这实际上使数组的大小减少了 1。
请注意，被交换的最后一个元素可能是您想要移除的值。但是不要担心，在下一次迭代中，我们仍然会检查这个元素。
*/
func removeElementV2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		} else {
			i++
		}
	}

	return j
}

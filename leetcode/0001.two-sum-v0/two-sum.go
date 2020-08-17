//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//示例: 给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9, 所以返回 [0, 1]

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 13
	fmt.Println(twoSum01(nums, target))
	fmt.Println(twoSum02(nums, target))
	fmt.Println(twoSum03(nums, target))
}

//暴力法
//遍历每一个元素x，并查找是否存在一个值与target-x相等的目标元素
//时间复杂度：O(n^2)
//空间复杂度：O(1)
func twoSum01(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

//两遍哈希表
//通过以空间换取速度的方式，我们可以将查找时间从 O(n) 降低到 O(1)。
//一个简单的实现使用了两次迭代。在第一次迭代中，我们将每个元素的值和它的索引添加到表中。
//在第二次迭代中，我们将检查每个元素所对应的目标元素（target−nums[i]）是否存在于表中。
//注意，该目标元素不能是 nums[i]本身！
//时间复杂度：O(n)，空间复杂度：O(n)
func twoSum02(nums []int, target int) []int {
	numsHash := make(map[int]int, len(nums))
	for key, value := range nums {
		numsHash[value] = key
	}

	for k, v := range nums {
		if val, ok := numsHash[target-v]; ok && k != val {
			return []int{k, val}
		}
	}
	return nil
}

//一遍哈希表
//两遍哈希表可以优化为一遍，在进行迭代并将元素插入哈希表时，我们会检查对应的目标元素（target-nums[i]）是否已经存在，
//如果已经存在，即可直接返回
//时间复杂度：O(n)，空间复杂度：O(n)
func twoSum03(nums []int, target int) []int {
	numsHash := make(map[int]int, len(nums))
	for k, v := range nums {
		if val, ok := numsHash[target-v]; ok && k != val {
			return []int{val, k}
		} else {
			numsHash[v] = k
		}
	}

	return nil
}

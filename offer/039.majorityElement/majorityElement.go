package offer

/*
剑指 Offer 39. 数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

限制：
1 <= 数组长度 <= 50000
*/

// 解法一： 哈希表统计法 或 数组排序法
// 哈希表统计法： 遍历数组 nums ，用 HashMap 统计各数字的数量，最终超过数组长度一半的数字则为众数。此方法时间和空间复杂度均为 O(N)。
// 数组排序法： 将数组 nums 排序，由于众数的数量超过数组长度一半，因此 数组中点的元素 一定为众数。此方法时间复杂度 O(Nlog_2N)。
func majorityElement(nums []int) int {
	numsHash := make(map[int]int)
	for _, val := range nums {
		if _, ok := numsHash[val]; !ok {
			numsHash[val] = 0
		}
		numsHash[val]++
	}

	for k, v := range numsHash {
		if v > len(nums)/2 {
			return k
		}
	}

	return -1
}

/*
解法二：摩尔投票法：
1. 票数和： 由于众数出现的次数超过数组长度的一半；若记 众数 的票数为 +1 ，非众数 的票数为 -1 ，则一定有所有数字的 票数和 >0 。
2. 票数正负抵消： 设数组 nums 中的众数为 x ，数组长度为 n 。若 nums 的前 a 个数字的 票数和 =0 ，则 数组后 (n−a) 个数字的 票数和一定仍 >0 （即后 (n−a) 个数字的 众数仍为 x ）。

算法流程:
1. 初始化： 票数统计 votes=0 ， 众数 x；
2. 循环抵消： 遍历数组 nums 中的每个数字 num ；
	2.1 当 票数 votes 等于 0 ，则假设 当前数字 num 为 众数 x ；
	2.2 当 num=x 时，票数 votes 自增 1 ；否则，票数 votes 自减 1。
3. 返回值： 返回 众数 x 即可。

局限性：
	摩尔投票法找的其实不是众数，而是占一半以上的数。当数组没有超过一半的数，则可能返回非众数，比如[1, 1, 2, 2, 2, 3, 3]，最终返回3。
	投票法简单来说就是不同则抵消，占半数以上的数字必然留到最后。
*/
func majorityElement2(nums []int) int {
	var votes, x int
	for _, v := range nums {
		if votes == 0 {
			x = v
		}

		if v == x {
			votes++
		} else {
			votes--
		}
	}

	return x
}

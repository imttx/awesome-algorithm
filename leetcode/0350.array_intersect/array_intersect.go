package leetcode

import "sort"

// 两数组交集
func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int, len(nums1))
	for _, n := range nums1 {
		nums1Map[n] += 1
	}

	var i int
	for _, n := range nums2 {
		if nums1Map[n] > 0 {
			nums1Map[n]--
			nums2[i] = n
			i++
		}
	}

	return nums2[:i]
}

// 如果给定的数组已经排好序呢？你将如何优化你的算法？
func intersectV2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var i, j, k int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}

	return nums1[:k]
}

package leetcode

func twoSum(nums []int, target int) []int {
	h := make(map[int]int, len(nums))
	for k, v := range nums {
		vv := target - v
		if kk, exist := h[vv]; exist {
			return []int{kk, k}
		}

		h[v] = k
	}
	return nil
}

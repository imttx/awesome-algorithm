package offer

/*
剑指 Offer 48. 最长不含重复字符的子字符串
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

// 解法一：双指针 + 哈希表
func lengthOfLongestSubstring(s string) int {
	hash := make(map[rune]int)

	i := -1
	j := 0
	c := 0

	for _, sv := range s {
		if k, ok := hash[sv]; ok {
			if k > i {
				i = k // 更新左指针
			}
		}

		hash[sv] = j    // 更新哈希表
		c = max(c, j-i) // 更新最大值
		j++
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

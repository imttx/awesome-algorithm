package offer

/*
剑指 Offer 50. 第一个只出现一次的字符
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:
s = "abaccdeff"
返回 "b"

s = ""
返回 " "

限制：
0 <= s 的长度 <= 50000
*/

// 解法一：哈希表统计
func firstUniqChar(s string) byte {
	strhash := make(map[rune]int, len(s))
	for _, ss := range s {
		if _, ok := strhash[ss]; !ok {
			strhash[ss] = 0
		}
		strhash[ss]++
	}

	for _, ss := range s {
		if v, ok := strhash[ss]; ok && v == 1 {
			return byte(ss)
		}
	}

	return ' '
}

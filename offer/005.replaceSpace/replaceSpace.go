package offer

import "strings"

/*
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."

限制：
0 <= s 的长度 <= 10000
*/

// 解法一： 遍历添加
func replaceSpace(s string) string {
	finalBuilder := new(strings.Builder)
	finalBuilder.Grow(len(s) * 3)
	for _, c := range s {
		if c == ' ' {
			finalBuilder.WriteByte('%')
			finalBuilder.WriteByte('2')
			finalBuilder.WriteByte('0')
		} else {
			finalBuilder.WriteByte(byte(c))
		}
	}
	return finalBuilder.String()
}

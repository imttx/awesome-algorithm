package offer

import "strings"

/*
剑指 Offer 58 - II. 左旋转字符串
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"

限制：
1 <= k < s.length <= 10000
*/

// 解法一：字符串切片
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// 方法二：列表遍历拼接
// 若面试规定不允许使用 切片函数 ，则使用此方法。
// 算法流程：
// 新建一个 list(Python)、StringBuilder(Java) ，记为 res ；
// 先向 res 添加 “第n+1 位至末位的字符” ；
// 再向 res 添加 “首位至第 n 位的字符” ；
// 将 res 转化为字符串并返回。
func reverseLeftWords2(s string, n int) string {
	res := new(strings.Builder)
	for i := n; i < len(s); i++ {
		res.WriteByte(s[i])
	}
	for i := 0; i < n; i++ {
		res.WriteByte(s[i])
	}
	return res.String()
}

// 方法三：字符串遍历拼接
// 此方法与 方法二 思路一致，区别是使用字符串代替列表。
func reverseLeftWords3(s string, n int) string {
	res := ""
	for i := n; i < (n + len(s)); i++ {
		res += string(s[i%len(s)])
	}
	return res
}

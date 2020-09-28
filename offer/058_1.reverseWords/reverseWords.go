package offer

import "strings"

/*
剑指 Offer 58 - I. 翻转单词顺序
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

说明：
无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
*/

// 解法一：双指针
// 算法解析：
// 倒序遍历字符串 s ，记录单词左右索引边界 i , j ；
// 每确定一个单词的边界，则将其添加至单词列表 res ；
// 最终，将单词列表拼接为字符串，并返回即可。
func reverseWords(s string) string {
	res := &strings.Builder{}
	// l, r := len(s)-1, len(s)-1
	for l := len(s) - 1; l >= 0; l-- {
		if s[l] == ' ' {
			continue
		}

		r := l
		for l >= 0 && s[l] != ' ' {
			l--
			continue
		}
		res.WriteString(s[l+1 : r+1])
		res.WriteString(" ")
	}

	return strings.TrimSuffix(res.String(), " ")
}

// 方法二：分割 + 倒序
// 利用 “字符串分割”、“列表倒序” 的内置函数 （面试时不建议使用） ，可简便地实现本题的字符串翻转要求。

// 算法解析：
// Python ： 由于 split() 方法将单词间的 “多个空格看作一个空格” （参考自 split()和split(' ')的区别 ），因此不会出现多余的 “空单词” 。因此，直接利用 reverse() 方法翻转单词列表 strs ，拼接为字符串并返回即可。

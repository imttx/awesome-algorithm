package offer

import (
	"sort"
	"strconv"
	"strings"
)

/*
剑指 Offer 45. 把数组排成最小的数
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:
输入: [10,2]
输出: "102"

示例 2:
输入: [3,30,34,5,9]
输出: "3033459"

提示:
0 < nums.length <= 100
*/

// 解题思路：
// 此题求拼接起来的 “最小数字” ，本质上是一个排序问题。
// 排序判断规则： 设 nums 任意两数字的字符串格式 x 和 y ，则
// 若拼接字符串 x+y>y+x ，则 x>y ；
// 反之，若 x+y<y+x ，则 x<y ；
// 根据以上规则，套用任何排序方法对 nums 执行排序即可。
//
// 举例：x=3 y=30
// 拼接：x+y = 330 ，y+x=303
// 此时：303 < 330 ，y 应该排在 x 前面
//
// 算法流程：
// 初始化： 字符串列表 strs ，保存各数字的字符串格式；
// 列表排序： 应用以上 “排序判断规则” ，对 strs 执行排序；
// 返回值： 拼接 strs 中的所有字符串，并返回。
func minNumber(nums []int) string {
	numsSlice := make([]string, 0, len(nums))
	for _, v := range nums {
		numsSlice = append(numsSlice, strconv.Itoa(v))
	}

	sort.Slice(numsSlice, func(i, j int) bool {
		return numsSlice[i]+numsSlice[j] < numsSlice[j]+numsSlice[i]
	})

	return strings.Join(numsSlice, "")
}

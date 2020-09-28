package offer

/*
剑指 Offer 42. 连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

提示：
1 <= arr.length <= 10^5
-100 <= arr[i] <= 100
*/

/*
解法一：动态规划解析：
状态定义： 设动态规划列表 dp ，dp[i] 代表以元素 nums[i] 为结尾的连续子数组最大和。

为何定义最大和 dp[i] 中必须包含元素 nums[i] ：保证 dp[i] 递推到 dp[i+1] 的正确性；如果不包含 nums[i] ，递推时则不满足题目的 连续子数组 要求。
转移方程： 若 dp[i−1]≤0 ，说明 dp[i - 1] 对 dp[i] 产生负贡献，即 dp[i-1] + nums[i] 还不如 nums[i] 本身大。
	当 dp[i−1]>0 时：执行 dp[i]=dp[i−1]+nums[i] ；
	当 dp[i−1]≤0 时：执行 dp[i]=nums[i] ；

初始状态：
dp[0]=nums[0]，即以 nums[0] 结尾的连续子数组最大和为 nums[0]。

返回值：
返回 dp 列表中的最大值，代表全局最大值。

空间复杂度降低：
由于 dp[i] 只与 dp[i-1] 和 nums[i] 有关系，因此可以将原数组 nums 用作 dp 列表，即直接在 nums 上修改即可。
*/

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

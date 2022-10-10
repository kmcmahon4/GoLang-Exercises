package leetcode

//https://leetcode.com/problems/two-sum/description/

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// This answer is suboptimal and does not pass all test cases

func TwoSum(nums []int, target int) []int {
	result := make([]int, 2)

	if len(nums) == 0 {
		return result
	}

	for idx, value := range nums {
		diff := target - value
		for i, val := range nums {
			if val == diff {
				result[0] = i
				result[1] = idx
				break
			}
		}
	}
	return result
}

package leetcode

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice
//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
//Example:
//Given nums = [2, 7, 11, 15], target = 9,
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
func TwoSum(nums []int, target int) []int {
	// nums[i] -> i
	container := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		_, ok := container[complement]
		if ok {
			return []int{ container[complement], i }
		}
		container[nums[i]] = i
	}
	return []int{}
}
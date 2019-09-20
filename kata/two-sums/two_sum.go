/* leetcode.com kata
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

TODO: version implementing a 2 pass hashmap
TODO: version implementing a 1 pass hashmap
*/

package main

func twoSum(nums []int, target int) []int {
	arr := []int{0, 1}

	for arr[0] < len(nums)-2 {
		for arr[1] < len(nums) {
			sum := nums[arr[0]] + nums[arr[1]]
			if sum == target {
				return arr
			}
			arr[1] = arr[1] + 1
		}
		arr[0] = arr[0] + 1
		arr[1] = arr[0] + 1
	}

	return arr
}

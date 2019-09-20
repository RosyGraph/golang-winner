/* leetcode.com kata
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	checkSum := func(got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	tc := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "positive case 1", nums: []int{2, 7, 11, 15},
			target: 9, want: []int{0, 1}},
		{name: "positive case 2", nums: []int{2, 7, 11, 15},
			target: 18, want: []int{1, 2}},
		{name: "negative case 1", nums: []int{-1, -2, -3, -4, -5},
			target: -8, want: []int{2, 4}},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := twoSum(c.nums, c.target)
			checkSum(got, c.want)
		})
	}
}

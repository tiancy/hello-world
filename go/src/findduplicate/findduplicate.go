/*
Source  https://leetcode.com/problems/find-the-duplicate-number/
Author  Tian
Date    2018/11/17

287. Find the Duplicate Number

Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive),
prove that at least one duplicate number must exist. Assume that there is only one duplicate number,
find the duplicate one.

Example 1:

	Input: [1,3,4,2,2]
	Output: 2

Example 2:

	Input: [3,1,3,4,2]
	Output: 3

*/

package findduplicate

import (
	"fmt"
	"sort"
)

// solution 1
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// solution 2
func findDuplicate2(nums []int) int {
	tmp := make(map[int]bool)
	for _, value := range nums {
		if tmp[value] {
			return value
		}
		tmp[value] = true
	}
	return -1
}

// solution 3
func findDuplicate3(nums []int) int {
	tortoise := nums[0]
	hare := nums[0]

	for {
		tortoise = nums[tortoise] // step
		hare = nums[nums[hare]]   // two steps, first step was nums[hare], second step was nums[nums[hare]]
		fmt.Printf("t: %v, h: %v\n", tortoise, hare)
		if tortoise == hare {
			break
		}
	}
	return -1
}

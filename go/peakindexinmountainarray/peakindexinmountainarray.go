/*
Source  https://leetcode.com/problems/peak-index-in-a-mountain-array/description/
Author  tian
Created 2018/6/24

852. Peak Index in a Mountain Array

Let's call an array A a mountain if the following properties hold:
	A.length >= 3
	There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]

Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:
	Input: [0,1,0]
	Output: 1
Example 2:
	Input: [0,2,1,0]
	Output: 1

Note:
	3 <= A.length <= 10000
	0 <= A[i] <= 10^6
	A is a mountain, as defined above.
*/
package main

import (
	"fmt"
)

// 实现思路，二分查找，找到最大数
func peakIndexInMountainArray(A []int) int {
	if len(A) < 3 || len(A) > 10000 {
		return -1
	}

	start := 0
	end := len(A) - 1

	for start+1 < end {
		mid := (start + end) / 2
		if A[mid] > A[mid-1] && A[mid] > A[mid+1] {
			return mid
		}
		if A[mid] > A[mid-1] {
			start = mid
		} else {
			end = mid
		}
	}
	return -1
}

func main() {
	A := []int{1, 6, 7, 2}
	i := peakIndexInMountainArray(A)
	fmt.Println(i)
}

package main

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

func main() {
	nums := []int{1, 2, 4, 2, 3}
	r := findDuplicate3(nums)
	fmt.Println(r)
}

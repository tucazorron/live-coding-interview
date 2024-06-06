/*

problem:
reduce the amplitude of the array using only 3 movements for removing some edge element

input:
integer array

ouput:
sum of subarray

*/

package main

import (
	"fmt"
	"sort"
)

func min(nums []int) int {
	min_num := nums[0]
	for _, num := range nums[1:] {
		if num < min_num {
			min_num = num
		}
	}
	return min_num
}

func solution(nums []int) int {
	size := len(nums)
	if size <= 4 {
		return 0
	}
	sort.Ints(nums)
	return min([]int{nums[size-4] - nums[0], nums[size-3] - nums[1], nums[size-2] - nums[2], nums[size-1], nums[3]})
}

func main() {
	tests := [][]int{
		{1, 2, 3, 4, 5}, // 1
	}
	for _, test := range tests {
		result := solution(test)
		fmt.Printf("test: %v => %d\n", test, result)
	}
}

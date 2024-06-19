package main

import (
	"fmt"
	"sort"
)

type input struct {
	array  []int
	target int
}

// time n2
func solution1(nums []int, target int) []int {
	for i, n1 := range nums {
		for j, n2 := range nums {
			if i != j && n1+n2 == target {
				return []int{n1, n2}
			}
		}
	}
	return []int{}
}

// time nlogn
func solution2(nums []int, target int) []int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{}
}

// time n
func solution3(nums []int, target int) []int {
	diffSet := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := diffSet[num]; exists {
			return []int{target - num, num}
		}
		diffSet[target-num] = struct{}{}
	}
	return []int{}
}

func main() {
	tests := []input{
		{[]int{1, 1}, 2},       // true
		{[]int{0, 9, 0}, 0},    // false
		{[]int{1}, 1},          // false
		{[]int{1, 2, 4, 3}, 4}, // false

	}
	for _, test := range tests {
		result := solution3(test.array, test.target)
		fmt.Printf("test: %v => %v\n", test, result)
	}
}

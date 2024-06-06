/*

problem:
verify if two values can sum to a target

input:
integer array
target (integer)

ouput:
boolean

*/

package main

import "fmt"

type input struct {
	array  []int
	target int
}

func sumOfTwo(nums []int, target int) bool {
	if len(nums) < 2 {
		return false
	}
	diffs := make(map[int]struct{})
	for _, num := range nums {
		complement := target - num
		if _, exists := diffs[complement]; exists {
			return true
		}
		diffs[complement] = struct{}{}
	}
	return false
}

func main() {
	tests := []input{
		{[]int{1, 1}, 2},    // true
		{[]int{}, 0},        // false
		{[]int{1}, 1},       // false
		{[]int{1, 2, 4}, 4}, // false

	}
	for _, test := range tests {
		result := sumOfTwo(test.array, test.target)
		fmt.Printf("test: %v => %t\n", test, result)
	}
}

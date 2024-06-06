/*

problem:
after removing one element from the array, verify if it can be divided into two subarrays with the same total sum

input:
integer array

ouput:
boolean

*/

package main

import "fmt"

func sum(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	return totalSum
}

// time: o(n^2), space: o(n)
func solution(nums []int) bool {
	size := len(nums)
	if size < 3 {
		return false
	}
	for idx := range nums {
		copyArray := make([]int, size)
		copy(copyArray, nums)
		subArray := append(copyArray[:idx], copyArray[idx+1:]...)
		sumSubArray := sum(subArray)
		if sumSubArray%2 != 0 {
			continue
		}
		halfSum := sumSubArray / 2
		for _, num := range subArray {
			sumSubArray -= num
			if sumSubArray == halfSum {
				return true
			}
		}
	}
	return false
}

func main() {
	tests := [][]int{
		{1, 2, 3, 3, 3, 20}, // true
		{},                  // false
		{1},                 // false
		{1, 1},              // false
	}
	for _, test := range tests {
		result := solution(test)
		fmt.Printf("test: %v => %t\n", test, result)
	}
}

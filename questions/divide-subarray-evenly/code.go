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

// sum all the array integers
func sum(nums []int) int {
	total_sum := 0
	for _, num := range nums {
		total_sum += num
	}
	return total_sum
}

// time: o(n^2), space: o(n)
func solution(nums []int) bool {
	size := len(nums)
	if size < 3 {
		return false
	}
	total_sum := sum(nums)
	for idx, num := range nums {
		if (total_sum - num) % 2 != 0 {
			continue
		}
		copyArray := make([]int, size)
		copy(copyArray, nums)
		subArray := append(copyArray[:idx], copyArray[idx+1:]...)
		sumSubArray := sum(subArray)
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

// time: O(n), space: O(n)
func solution2(nums []int) bool {

	// cant split evenly array less than 3 elements
	if len(nums) < 3 {
		return false
	}

	// create a hashset for elements to find
	nums_to_find := make(map[int]struct{})

	// create a hashset for already read elements
	nums_read := make(map[int]struct{})
	
	// create our candidate value
	candidate := sum(nums)

	// loop all elements
	for _, num := range nums {

		// if you found a number we were looking for
		if _, exists := nums_to_find[num]; exists {
			return true
		}

		// we update our candidate
		candidate -= 2 * num

		// if we found a candidate that is the complement of a read number
		if _, exists := nums_read[-candidate]; exists {
			return true
		}

		// update our hashsets
		nums_to_find[candidate] = struct{}{}
		nums_read[num] = struct{}{}
	}
	return false
}

func main() {
	tests := [][]int{
		{1, 2, 3, 3, 3, 20}, 	// true
		{},                  	// false
		{1},                 	// false
		{1, 1},              	// false
		{9, -8, 3, 6, -1, -10}, 	// true
	}
	for _, test := range tests {
		result := solution2(test)
		fmt.Printf("test: %v => %t\n", test, result)
	}
}

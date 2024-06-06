/*

problem:
find the most common chars from string

input:
string

ouput:
list of chars

*/

package main

import "fmt"

func solution(array string) []rune {
	counter := make(map[rune]int)
	for _, char := range array {
		if _, exists := counter[char]; !exists {
			counter[char] = 0
		}
		counter[char]++
	}
	result := []rune{}
	maxReps := 0
	for char, reps := range counter {
		if reps > maxReps {
			maxReps = reps
			result = []rune{char}
		} else if reps == maxReps {
			result = append(result, char)
		}
	}
	return result
}

func main() {
	tests := []string{
		"a",     // [a]
		"ab",    // [a, b]
		"aabcb", // [a, b]
	}
	for _, test := range tests {
		result := solution(test)
		fmt.Printf("test: %s => %c\n", test, result)
	}
}

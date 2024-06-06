/*

problem:
find the first char duplicated in the string

input:
string

ouput:
first duplicated char

*/

package main

import "fmt"

func solution(array string) rune {
	found := make(map[rune]struct{})
	for _, char := range array {
		if _, exists := found[char]; exists {
			return char
		}
		found[char] = struct{}{}
	}
	return ' '
}

func main() {
	tests := []string{
		"abb",   // b
		"abcba", // b
	}
	for _, test := range tests {
		result := solution(test)
		fmt.Printf("test: %s => %c\n", test, result)
	}
}

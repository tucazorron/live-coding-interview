package main

import (
	"fmt"
	"strings"
)

// func solution2(words []string) []string {
// 	max_counter := make(map[rune]int)
// 	for _, word := range words {
// 		frequency := strings.Map(func(r rune) rune {return r}, word)
// 		for letter, count := range frequency {

// 		}
// 	}
// }


// time: O(n**2)
// space: O(n)
func solution(words []string) []string {
	max_counter := make(map[rune]int)
	for _, word := range words {
		curr_counter := make(map[rune]int)
		rune_words := []rune(word)
		for _, letter := range rune_words {
			if _, exists := curr_counter[letter]; !exists {
				curr_counter[letter] = 0
			}
			curr_counter[letter]++
		}
		for letter, count := range curr_counter {
			max_count, exists := max_counter[letter]
			if (exists && max_count < count) || !exists {
				max_counter[letter] = count
			}
		}
	}
	runes := []rune{}
	for letter := range max_counter {
		repeated_string := strings.Repeat(string(letter), max_counter[letter])
		runes = append(runes, []rune(repeated_string)...)
	}
	answer := []string{}
	for _, letter := range runes {
		answer = append(answer, string(letter))
	}
	return answer
}

func main() {
	fmt.Println(solution([]string{"your", "you", "or", "u"}))
	fmt.Println(solution([]string{"this", "that", "did", "deed", "them", "a"}))
	fmt.Println(solution([]string{"a", "abc", "ab", "boo"}))
	fmt.Println(solution([]string{"a"}))
	fmt.Println(solution([]string{"abc", "ab", "b", "bac", "c"}))
	fmt.Println(solution([]string{"!!!2", "234", "222", "432"}))
}

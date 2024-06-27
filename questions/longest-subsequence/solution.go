package main

import (
	"fmt"
)

func solution(s string, words []string) []string {
	maxsubs, maxlen := []string{}, 0
	for _, word := range words {
		si, wi := 0, 0
		for si < len(s) && wi < len(word) {
			if word[wi] == s[si] {
				wi++
			}
			si++
		}
		if wi == len(word) {
			if len(word) > maxlen {
				maxsubs = []string{word}
				maxlen = len(word)
			} else if len(word) == maxlen {
				maxsubs = append(maxsubs, word)
			}
		}
	}
	return maxsubs
}

func main() {
	fmt.Println(solution(
		"arturzorron",
		[]string{
			"azn",
			"aton",
			"zorran",
			"auro",
		},
	))
}

package main

import (
	"fmt"
)

// time: O(w) (w = len of words)
// space: O(26) (26 = len of alphabet)
func solution1(words []string) []string {
	result := []string{}
	set := make(map[rune]bool)
	for _, word := range words {
		repeatedChar := false
		for _, char := range word {
			if set[char] {
				repeatedChar = true
				break
			}
		}
		if !repeatedChar {
			result = append(result, word)
			for _, char := range word {
				set[char] = true
			}
		}
		if len(result) == 5 {
			break
		}
	}
	return result
}

// type TrieNode struct {
// 	children map[rune]*TrieNode
// 	isWord   bool
// }

// type Trie struct {
// 	root *TrieNode
// }

// func newTrie() *Trie {
// 	return &Trie{root: newTrieNode()}
// }

// func newTrieNode() *TrieNode {
// 	return &TrieNode{children: make(map[rune]*TrieNode)}
// }

// func (t *Trie) insert(word string) {
// 	node := t.root
// 	for _, char := range word {
// 		if _, exists := node.children[char]; !exists {
// 			node.children[char] = newTrieNode()
// 		}
// 		node = node.children[char]
// 	}
// 	node.isWord = true
// }

// func (t *Trie) search(usedLetters *map[rune]bool, words *[]string) bool {
// 	node := t.root

// 	for _, char := range word {
// 		if _, exists := *usedLetters[char]; exists {
// 			return false
// 		}
// 	}
// }

// func solution2(words []string) []string {
// 	trie := newTrie()
// }

func main() {
	fmt.Println(solution1([]string{"afkpu", "abcde", "fghij", "klmno", "pqrst", "uvwxy"}))
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	var decoded []rune
	letters := [][]rune{}
	digits := [][]rune{}
	for i, char := range s {
		if char == '[' {
			letters = append(letters, []rune{})
		} else if char == ']' {
			mult, _ := strconv.Atoi(string(digits[len(digits)-1]))
			if len(digits) > 1 {
				letters[len(letters)-2] = append(letters[len(letters)-2], []rune(strings.Repeat(string(letters[len(letters)-1]), mult))...)
			} else {
				decoded = append(decoded, []rune(strings.Repeat(string(letters[len(letters)-1]), mult))...)
			}
			letters = letters[:len(letters)-1]
			digits = digits[:len(digits)-1]
		} else if unicode.IsLetter(char) {
			if len(digits) == 0 {
				decoded = append(decoded, char)
			} else {
				letters[len(letters)-1] = append(letters[len(letters)-1], char)
			}
		} else {
			if i == 0 || !unicode.IsDigit(rune(s[i-1])) {
				digits = append(digits, []rune{char})
			} else {
				digits[len(digits)-1] = append(digits[len(digits)-1], char)
			}
		}
		fmt.Println("i", i, "char", char, "decoded", decoded, "letters", letters, "digits", digits)
	}
	return string(decoded)
}
func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}

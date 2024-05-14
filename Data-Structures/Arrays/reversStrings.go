package main

import "fmt"

// reverseString reverses a string by iterating backwards.
func reverseString(word string) string {
	var newWord string
	for i := len(word) - 1; i >= 0; i-- {
		newWord += fmt.Sprintf("%c", word[i])
	}
	return newWord
}

// reverseStringTwo reverses a string by iterating forwards and appending characters in reverse order.
func reverseStringTwo(str string) string {
	var newStr string
	for _, v := range str {
		newStr = string(v) + newStr // Add the current character backwards
	}
	return newStr
}

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	var text_reverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		text_reverse += string(text[i])
	}

	if text_reverse == text {
		return true
	} else {
		return false
	}
}

func main() {
	slice := []string{"hi", "what", "are", "you", "doing", "?"}

	for i := range slice {
		fmt.Printf("\ni=%d\n", i)
		fmt.Println(slice[i])
	}

	var text string = "AvA"
	// This print a ASCII code, not a letter; for this we need convert it with string(arg)
	fmt.Println(text[0])
	// Convert to string
	fmt.Println(string(text[0]))

	// This will return a bool if the text is a palindrome
	// true = is a palindrome
	// false = is not a palindrome
	fmt.Printf("The text is a palindrome = %v", isPalindrome(text))
}

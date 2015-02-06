/*
 * Checks if the string entered by the user is a palindrome. That is that it reads the same forwards as backwards like “racecar”
*/

package main 

import (
	"fmt"
)

func Reverse(text string) string {
	runes := []rune(text)
	length := len(runes)
	step := length / 2
	for i:=0; i<step; i++ {
		runes[i],runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func IsPalindrome(word string) bool {
	return word == Reverse(word)
}

func main() {
	fmt.Print("Enter a string: ")
	word := ""
	fmt.Scanln(&word)
	if IsPalindrome(word) {
		fmt.Printf("%s is a palindrome\n", word)
	} else {
		fmt.Printf("%s is not a palindrome\n", word)
	}
}
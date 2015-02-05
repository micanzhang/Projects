/*
* Enter a string and the program will reverse it and print it out.
*/

package main 

import (
	"fmt"
)

func reverse(text string) string {
	runes := []rune(text)
	length := len(runes)
	step := length / 2
	for i:=0; i<step; i++ {
		runes[i],runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	text := ""
	fmt.Println("Enter a string to reverse: ")
	fmt.Scanln(&text)
	fmt.Println(reverse(text))
}
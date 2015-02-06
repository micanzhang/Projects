/*
* Counts the number of individual words in a string. 
* For added complexity read these strings in from a text file and generate a summary.
*/

package main 

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func WordCounter(msg string) map[string]int {
	m := make(map[string]int)
	words := strings.Split(msg, " ")
	for _,word := range words {
		m[word] += 1
	}
	return m
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string:")
	msg, _ := reader.ReadString('\n')
	fmt.Println(WordCounter(msg))
}
